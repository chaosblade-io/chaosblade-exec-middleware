package redis

import (
	"context"
	"fmt"
	"github.com/chaosblade-io/chaosblade-exec-os/exec/category"
	"github.com/chaosblade-io/chaosblade-spec-go/log"
	"github.com/chaosblade-io/chaosblade-spec-go/spec"
	"github.com/go-redis/redis/v8"
	"github.com/shirou/gopsutil/process"
	"regexp"
	"strconv"
	"sync"
	"time"
)

const (
	CacheNetBin               = "chaos_cacheNet"
	NetActionCommandName      = "network"
	NetShortActionCommandName = "net"
)

type CacheNetworkActionCommandSpec struct {
	spec.BaseExpActionCommandSpec
}

func NewCacheNetworkActionSpec() spec.ExpActionCommandSpec {
	return &CacheNetworkActionCommandSpec{
		spec.BaseExpActionCommandSpec{
			ActionMatchers: []spec.ExpFlagSpec{},
			ActionFlags: []spec.ExpFlagSpec{
				&spec.ExpFlag{
					Name: "addr",
					Desc: "The address of redis server",
				},
				&spec.ExpFlag{
					Name: "password",
					Desc: "The password of redis server",
				},
				&spec.ExpFlag{
					Name:    "key-prefix",
					Desc:    "The prefix of key",
					Default: KeyPrefix,
				},
				&spec.ExpFlag{
					Name:    "key-expire",
					Desc:    "The expire of key ,default is 10m(10m,300s)",
					Default: DefaultExpire,
				},
				&spec.ExpFlag{
					Name:    "key-size",
					Desc:    "The size of key, default is 10kb（10,20）",
					Default: DefaultKeySize,
				},
				&spec.ExpFlag{
					Name:    "works",
					Desc:    "The number of works default is 20",
					Default: DefaultWorks,
				},
				&spec.ExpFlag{
					Name:    "duration",
					Desc:    "duration time, default is 10m(10m,5m,600s)",
					Default: DefaultRunDuration,
				},
				&spec.ExpFlag{
					Name:    "direction",
					Desc:    "The direction of network, default is all(in/out/all)",
					Default: "all",
				},
				&spec.ExpFlag{
					Name:    "single-node-prefix",
					Desc:    "The prefix of single node key({0000},{1111})",
					Default: SingleNodePrefix,
				},
			},
			ActionExecutor: &CacheNetworkExecutor{},
			ActionExample: `
`,
			ActionPrograms:    []string{CacheNetBin},
			ActionCategories:  []string{category.SystemTime},
			ActionProcessHang: true,
		},
	}
}

func (*CacheNetworkActionCommandSpec) Name() string {
	return NetActionCommandName
}

func (*CacheNetworkActionCommandSpec) Aliases() []string {
	return []string{NetShortActionCommandName}
}

func (*CacheNetworkActionCommandSpec) ShortDesc() string {
	return "Network"
}

func (k *CacheNetworkActionCommandSpec) LongDesc() string {
	if k.ActionLongDesc != "" {
		return k.ActionLongDesc
	}
	return "Network Injection"
}

func (*CacheNetworkActionCommandSpec) Categories() []string {
	return []string{category.SystemProcess}
}

type CacheNetworkExecutor struct {
	channel spec.Channel
}

func (cbe *CacheNetworkExecutor) Name() string {
	return NetActionCommandName
}

func (cbe *CacheNetworkExecutor) Exec(uid string, ctx context.Context, model *spec.ExpModel) *spec.Response {

	addrStr := model.ActionFlags["addr"]
	passwordStr := model.ActionFlags["password"]
	keyPrefixStr := model.ActionFlags["key-prefix"]
	singleNodePrefix := model.ActionFlags["single-node-prefix"]
	threadCountStr := model.ActionFlags["works"]
	keyExpire := model.ActionFlags["key-expire"]
	keySize := model.ActionFlags["key-size"]
	durationStr := model.ActionFlags["duration"]
	direction := model.ActionFlags["direction"]
	params := NetworkParams{
		Addr:             addrStr,
		Password:         passwordStr,
		KeyPrefix:        keyPrefixStr,
		Works:            threadCountStr,
		SingleNodePrefix: singleNodePrefix,
		KeyExpire:        keyExpire,
		KeySize:          keySize,
		Duration:         durationStr,
		Direction:        direction,
	}
	if _, ok := spec.IsDestroy(ctx); ok {
		return cbe.stop(ctx, uid, params)
	}

	cli := redis.NewClient(&redis.Options{
		Addr:     addrStr,
		Password: passwordStr,
	})

	_, err := cli.Ping(cli.Context()).Result()
	if err != nil {
		errMsg := "redis ping error: " + err.Error()
		log.Errorf(ctx, errMsg)
		return spec.ResponseFailWithFlags(spec.ActionNotSupport, errMsg)
	}
	_ = cli.Close()
	return cbe.start(ctx, params)
}

func (cbe *CacheNetworkExecutor) SetChannel(channel spec.Channel) {
	cbe.channel = channel
}

func (cbe *CacheNetworkExecutor) start(ctx context.Context, params NetworkParams) *spec.Response {
	duration, err := time.ParseDuration(params.Duration)
	if err != nil {
		errMsg := "parse duration-1 error: " + err.Error()
		log.Errorf(ctx, errMsg)
		return spec.ResponseFailWithFlags(spec.ActionNotSupport, errMsg)
	}

	threadCount, err := strconv.ParseInt(params.Works, 10, 32)
	if err != nil {
		errMsg := "parse thread count error: " + err.Error()
		log.Errorf(ctx, errMsg)
		return spec.ResponseFailWithFlags(spec.ActionNotSupport, errMsg)
	}
	size, err := strconv.ParseInt(params.KeySize, 10, 32)
	if err != nil {
		errMsg := "parse size error: " + err.Error()
		log.Errorf(ctx, errMsg)
		return spec.ResponseFailWithFlags(spec.ActionNotSupport, errMsg)
	}
	if size > 1024*5 {
		errMsg := "size can not greater than 5M"
		log.Errorf(ctx, errMsg)
		return spec.ResponseFailWithFlags(spec.ActionNotSupport, errMsg)
	}
	expire, err := time.ParseDuration(params.KeyExpire)
	if err != nil {
		errMsg := "parse expire error: " + err.Error()
		log.Errorf(ctx, errMsg)
		return spec.ResponseFailWithFlags(spec.ActionNotSupport, errMsg)
	}

	endTime := time.Now().Add(duration)
	valueForCache := generateStrForKB(int(size * 1024))
	var wg sync.WaitGroup
	for i := 0; i < int(threadCount); i++ {
		wg.Add(1)
		key := fmt.Sprintf("%s%snetwork:key%d", params.SingleNodePrefix, params.KeyPrefix, i)
		if params.Direction == "out" {
			_ = write(params, key, valueForCache, expire)
		}
		go func() {
			for {
				if time.Now().Before(endTime) {
					if params.Direction == "out" {
						_ = read(params, key)
					}
					if params.Direction == "in" {
						_ = write(params, key, valueForCache, time.Second*30)
					}
					if params.Direction == "all" {
						_ = readWrite(params, key, valueForCache, time.Second*30)
					}
				} else {
					break
				}
			}
			wg.Done()
		}()
	}
	log.Infof(ctx, "success start network command")
	wg.Wait()
	return spec.ReturnSuccess("finished network command success")
}

func (cbe *CacheNetworkExecutor) stop(ctx context.Context, uid string, params NetworkParams) *spec.Response {
	compile, _ := regexp.Compile(fmt.Sprintf(".*create.*%s.*", uid))
	processes, err := process.Processes()
	if err != nil {
		return spec.ResponseFailWithFlags(spec.ActionNotSupport, "failed to get processes")
	}

	works, err := strconv.ParseInt(params.Works, 10, 32)
	if err != nil {
		errMsg := "parse thread count error: " + err.Error()
		log.Errorf(ctx, errMsg)
		return spec.ResponseFailWithFlags(spec.ActionNotSupport, errMsg)
	}
	for _, proc := range processes {
		cmdline, err := proc.Cmdline()
		if err != nil {
			continue
		}

		if !compile.MatchString(cmdline) {
			continue
		}

		err = proc.Kill()
		if err != nil {
			return spec.ResponseFailWithFlags(spec.ActionNotSupport, " failed to stop network   command. find the process success, but failed to kill.")
		}
		err = stop(works, params)
		if err != nil {
			return spec.ResponseFailWithFlags(spec.ActionNotSupport, " failed to stop network   command. find the process success, but failed to kill.")
		}
		return spec.ReturnSuccess("stop network command success")

	}
	err = stop(works, params)
	if err != nil {
		return spec.ResponseFailWithFlags(spec.ActionNotSupport, " failed to stop network   command. find the process success, but failed to kill.")
	}
	return spec.ReturnSuccess("stop network command success")
}

func write(params NetworkParams, key, value string, exp time.Duration) error {
	cli := redis.NewClient(&redis.Options{
		Addr:     params.Addr,
		Password: params.Password,
	})
	err := cli.Set(context.Background(), key, value, exp).Err()
	return err
}

func read(params NetworkParams, key string) error {
	cli := redis.NewClient(&redis.Options{
		Addr:     params.Addr,
		Password: params.Password,
	})
	return cli.Get(context.Background(), key).Err()
}

func readWrite(params NetworkParams, key, v string, exp time.Duration) error {
	err := write(params, key, v, exp)
	if err != nil {
		return err
	}
	err = read(params, key)
	if err != nil {
		return err
	}
	return nil
}

func stop(works int64, params NetworkParams) error {
	threadCount := works
	errChan := make(chan error, threadCount)
	doneChan := make(chan bool)
	for i := 0; i < int(threadCount); i++ {
		key := fmt.Sprintf("%s%snetwork:key%d", params.SingleNodePrefix, params.KeyPrefix, i)
		go func() {
			cli := redis.NewClient(&redis.Options{
				Addr:     params.Addr,
				Password: params.Password,
			})
			err := cli.Del(context.Background(), key).Err()
			if err != nil {
				errChan <- err
				return
			}
			doneChan <- true
		}()
	}
	for i := 0; i < int(threadCount); i++ {
		select {
		case err := <-errChan:
			return err
		case <-doneChan:
		}
	}
	return nil
}

type NetworkParams struct {
	Addr             string `json:"addr"`
	Password         string `json:"password"`
	Works            string `json:"works"`
	KeyPrefix        string `json:"key-prefix"`
	SingleNodePrefix string `json:"single-node-prefix"`
	KeyExpire        string `json:"key-expire"`
	KeySize          string `json:"key-size"`
	Duration         string `json:"duration"`
	Direction        string `json:"direction"`
}
