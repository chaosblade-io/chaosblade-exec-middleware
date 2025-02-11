// Package redis-----------------------
// @author:  xiejunqiao
// @contact: xiejunqiao@wps.cn
// @since:   2024/11/25
// @desc: //TODO
// ----------------------------------------
package redis

import (
	"github.com/go-redis/redis/v8"
	"log"
	"testing"
	"time"
)

func Test_redis(t *testing.T) {
	clientList := make([]*redis.Client, 0)
	for i := 0; i < 100; i++ {
		cli := redis.NewClient(&redis.Options{
			Addr: "127.0.0.1:6379",
		})
		_, err := cli.Ping(cli.Context()).Result()
		if err != nil {
			log.Println("redis ping error: " + err.Error())
			continue
		}
		clientList = append(clientList, cli)
	}
	log.Println("clientList len: ", len(clientList))
	//定时
	tik := time.NewTicker(time.Second * 10)
	for {
		select {
		case <-tik.C:
			for _, cli := range clientList {
				_, err := cli.Ping(cli.Context()).Result()
				if err != nil {
					log.Println("redis ping error: " + err.Error())
					continue
				}
			}
		default:

		}
	}
}
