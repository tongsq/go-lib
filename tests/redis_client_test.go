package tests

import (
	"github.com/tongsq/go-lib/redis-client"
	"testing"
)

func getClient() *redis_client.RedisClient {
	return &redis_client.RedisClient{
		MaxIdle:   3,
		MaxActive: 3,
		Network:   "tcp",
		Address:   "127.0.0.1:6379",
	}
}

func TestSet(t *testing.T) {
	c := getClient()
	result, err := c.SetEx("test_set", "hello", 60)
	if !result || err != nil {
		t.Fatal("test RedisClient set fail", result, err)
	} else {
		t.Log("test RedisClient set success", result, err)
	}
}

func TestGet(t *testing.T) {
	c := getClient()
	result, err := c.Get("test_set")
	if err != nil {
		t.Fatal("test RedisClient get fail", result, err)
	} else {
		t.Log("test RedisClient get success", result, err)
	}
}

func TestDel(t *testing.T) {
	c := getClient()
	result, err := c.Del("test_set")
	if err != nil {
		t.Fatal("test RedisClient DEL fail", result, err)
	} else {
		t.Log("test RedisClient del success", result, err)
	}
}

func TestSadd(t *testing.T) {
	c := getClient()
	result, err := c.SAdd("test_sadd", "a", "b")
	if err != nil {
		t.Fatal("test RedisClient SADD fail", result, err)
	} else {
		t.Log("test RedisClient SADD success", result, err)
	}
}

func TestSmembers(t *testing.T) {
	c := getClient()
	result, err := c.SMembers("test_sadd")
	if err != nil {
		t.Fatal("test RedisClient smembers fail", result, err)
	} else {
		t.Log("test RedisClient smembers success", result, err)
	}
}

func TestSRem(t *testing.T) {
	c := getClient()
	result, err := c.SRem("test_sadd", "a", "b")
	if err != nil {
		t.Fatal("test RedisClient SRem fail", result, err)
	} else {
		t.Log("test RedisClient SRem success", result, err)
	}
}

func TestHMSet(t *testing.T) {
	c := getClient()
	fields := []redis_client.HMDto{redis_client.HMDto{Field: "k3", Value: "v3"}, redis_client.HMDto{Field: "k1", Value: "v1"}}
	result, err := c.HMSet("hmset", fields...)
	if err != nil || !result {
		t.Fatal("test RedisClient HMSet fail", result, err)
	} else {
		t.Log("test RedisClient HMSet success", result, err)
	}
}

func TestHMGet(t *testing.T) {
	c := getClient()

	result, err := c.HMGet("hmset", "k1", "k4")
	if err != nil || result == nil {
		t.Fatal("test RedisClient HMGet fail", result, err)
	} else {
		t.Log("test RedisClient HMGet success", result, err)
	}
}

func TestHMGetOne(t *testing.T) {
	c := getClient()

	result, err := c.HMGetOne("hmset", "k1")
	if err != nil {
		t.Fatal("test RedisClient HMGetOne fail", result, err)
	} else {
		t.Log("test RedisClient HMGetOne success", result, err)
	}
}

func TestHDel(t *testing.T) {
	c := getClient()

	result, err := c.HDel("hmset", "k1", "k4")
	if err != nil {
		t.Fatal("test RedisClient Hdel fail", result, err)
	} else {
		t.Log("test RedisClient Hdel success", result, err)
	}
}
