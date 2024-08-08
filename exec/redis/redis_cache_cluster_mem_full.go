package redis

import (
	"bytes"
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
	CacheMemFullBin           = "chaos_cacheMemFull"
	DefaultKeySize            = "10" //kb
	DefaultTotalKeys          = "100000"
	DefaultWorks              = "20"
	DefaultRunDuration        = "10m"
	DefaultExpire             = "10m"
	MemActionCommandName      = "mem-full"
	MemShortActionCommandName = "full"
	KeyPrefix                 = "chaos:"
	SingleNodePrefix          = "{00000}:"
)

type CacheMemFullActionCommandSpec struct {
	spec.BaseExpActionCommandSpec
}

func NewCacheMemFullActionSpec() spec.ExpActionCommandSpec {
	return &CacheMemFullActionCommandSpec{
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
					Name:    "total-keys",
					Desc:    "The total number of keys, default is 100000",
					Default: DefaultTotalKeys,
				},
				&spec.ExpFlag{
					Name:    "works",
					Desc:    "The number of works default is 20",
					Default: DefaultWorks,
				},
				&spec.ExpFlag{
					Name: "single-node-prefix",
					Desc: "The prefix of single node key({0000},{1111})",
				},
				&spec.ExpFlag{
					Name:    "duration",
					Desc:    "duration time, default is 10m(10m,5m,600s)",
					Default: DefaultRunDuration,
				},
			},
			ActionExecutor: &CacheMemFullExecutor{},
			ActionExample: `
`,
			ActionPrograms:    []string{CacheMemFullBin},
			ActionCategories:  []string{category.SystemTime},
			ActionProcessHang: true,
		},
	}
}

func (*CacheMemFullActionCommandSpec) Name() string {
	return MemActionCommandName
}

func (*CacheMemFullActionCommandSpec) Aliases() []string {
	return []string{MemShortActionCommandName}
}

func (*CacheMemFullActionCommandSpec) ShortDesc() string {
	return "Cache Mem Full"
}

func (k *CacheMemFullActionCommandSpec) LongDesc() string {
	if k.ActionLongDesc != "" {
		return k.ActionLongDesc
	}
	return "Memory Injection"
}

func (*CacheMemFullActionCommandSpec) Categories() []string {
	return []string{category.SystemProcess}
}

type CacheMemFullExecutor struct {
	channel spec.Channel
}

func (cbe *CacheMemFullExecutor) Name() string {
	return MemActionCommandName
}

func (cbe *CacheMemFullExecutor) Exec(uid string, ctx context.Context, model *spec.ExpModel) *spec.Response {

	addrStr := model.ActionFlags["addr"]

	passwordStr := model.ActionFlags["password"]
	keyPrefixStr := model.ActionFlags["key-prefix"]
	singleNodePrefix := model.ActionFlags["single-node-prefix"]
	threadCountStr := model.ActionFlags["works"]
	keyExpire := model.ActionFlags["key-expire"]
	keySize := model.ActionFlags["key-size"]
	totalKeys := model.ActionFlags["total-keys"]
	durationStr := model.ActionFlags["duration"]
	params := Params{
		Addr:             addrStr,
		Password:         passwordStr,
		KeyPrefix:        keyPrefixStr,
		Works:            threadCountStr,
		SingleNodePrefix: singleNodePrefix,
		KeyExpire:        keyExpire,
		KeySize:          keySize,
		TotalKeys:        totalKeys,
		Duration:         durationStr,
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

func (cbe *CacheMemFullExecutor) SetChannel(channel spec.Channel) {
	cbe.channel = channel
}

func (cbe *CacheMemFullExecutor) start(ctx context.Context, params Params) *spec.Response {
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
	if size > 1024 {
		errMsg := "size can not greater than 1M"
		log.Errorf(ctx, errMsg)
		return spec.ResponseFailWithFlags(spec.ActionNotSupport, errMsg)
	}
	totalKeys, err := strconv.ParseInt(params.TotalKeys, 10, 32)
	if err != nil {
		errMsg := "parse total keys error: " + err.Error()
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
		go func(num int) {
			cli := redis.NewClient(&redis.Options{
				Addr:     params.Addr,
				Password: params.Password,
			})
			for j := 0; j < int(totalKeys/threadCount); j++ {
				key := fmt.Sprintf("%s%skey%d:%d", params.SingleNodePrefix, params.KeyPrefix, num, j)
				if time.Now().Before(endTime) {
					cli.Set(context.Background(), key, valueForCache, expire)
				} else {
					break
				}
			}
			wg.Done()
		}(i)
	}
	log.Infof(ctx, "success start full memory command")
	wg.Wait()
	return spec.ReturnSuccess("finished full memory command success")
}

func (cbe *CacheMemFullExecutor) stop(ctx context.Context, uid string, params Params) *spec.Response {
	compile, _ := regexp.Compile(fmt.Sprintf(".*create.*%s.*", uid))
	processes, err := process.Processes()
	if err != nil {
		return spec.ResponseFailWithFlags(spec.ActionNotSupport, "failed to get processes")
	}

	totalKeys, err := strconv.ParseInt(params.TotalKeys, 10, 32)
	if err != nil {
		errMsg := "parse total keys error: " + err.Error()
		log.Errorf(ctx, errMsg)
		return spec.ResponseFailWithFlags(spec.ActionNotSupport, errMsg)
	}
	threadCount, err := strconv.ParseInt(params.Works, 10, 32)
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
			return spec.ResponseFailWithFlags(spec.ActionNotSupport, " failed to stop mem-full   command. find the process success, but failed to kill.")
		}
		errChan := make(chan error, threadCount)
		doneChan := make(chan bool)
		for i := 0; i < int(threadCount); i++ {
			go func(num int) {
				cli := redis.NewClient(&redis.Options{
					Addr:     params.Addr,
					Password: params.Password,
				})
				for j := 0; j < int(totalKeys/threadCount); j++ {
					key := fmt.Sprintf("%s%skey%d:%d", params.SingleNodePrefix, params.KeyPrefix, num, j)
					err = cli.Del(context.Background(), key).Err()
					if err != nil {
						errChan <- err
						return
					}
				}
				doneChan <- true
			}(i)
		}
		for i := 0; i < int(threadCount); i++ {
			select {
			case err = <-errChan:
				return spec.ResponseFailWithFlags(spec.ActionNotSupport, "failed to stop mem-full command. find the process success, but failed to kill.")
			case <-doneChan:
			}
		}
		return spec.ReturnSuccess("stop mem-full command success")

	}
	return spec.ResponseFailWithFlags(spec.ActionNotSupport, "failed to stop mem-full command. can not find the process success.")
}

func generateStrForKB(size int) string {
	var buffer bytes.Buffer
	for buffer.Len() < size {
		buffer.WriteString("C")
	}

	return buffer.String()
}

type Params struct {
	Addr             string `json:"addr"`
	Password         string `json:"password"`
	Works            string `json:"works"`
	KeyPrefix        string `json:"key-prefix"`
	SingleNodePrefix string `json:"single-node-prefix"`
	KeyExpire        string `json:"key-expire"`
	KeySize          string `json:"key-size"`
	TotalKeys        string `json:"total-keys"`
	Duration         string `json:"duration"`
}
