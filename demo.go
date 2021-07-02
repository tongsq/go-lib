package main

import (
	"github.com/tongsq/go-lib/logger"
	redis_client "github.com/tongsq/go-lib/redis-client"
	"time"
)

func main() {
	//logrus.SetFormatter(&logrus.TextFormatter{
	//	TimestampFormat: "2006-01-02 15:04:05",
	//	DisableTimestamp: false,
	//	FullTimestamp: true,
	//})
	//logrus.WithFields(logrus.Fields{"key":"value"}).Info("hello world")
	c := &redis_client.RedisClient{
		MaxIdle:     5,
		MaxActive:   5,
		Network:     "tcp",
		Address:     "127.0.0.1:6379",
		IdleTimeout: 5 * time.Second,
	}
	c.SetEx("test", "aa", 60)
	for i := 0; i < 1000; i++ {
		//component.TaskPool.RunTask(func() {
		//	test(c, i)
		//})
		test(c, i)
		time.Sleep(time.Second * 6)
	}

}

func test(c *redis_client.RedisClient, i int) {
	r, err := c.HMGetOne("proxy_info_map", "110.243.4.213:9999")
	if err != nil {
		logger.Error("set fail", logger.Fields{"err": err, "r": r})
	} else {
		logger.Success("set success", logger.Fields{"r": r, "i": i})
	}
}
