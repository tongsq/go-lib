package redis_client

import (
	"github.com/gomodule/redigo/redis"
	"github.com/tongsq/go-lib/logger"
	"time"
)

type RedisClient struct {
	pool         *redis.Pool
	MaxIdle      int
	MaxActive    int
	Network      string
	Address      string
	Password     string
	Db           int
	Username     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
}

func (c *RedisClient) GetPool() *redis.Pool {
	if c.pool == nil {
		var options []redis.DialOption
		if c.Password != "" {
			options = append(options, redis.DialPassword(c.Password))
		}
		if c.Db > 0 {
			options = append(options, redis.DialDatabase(c.Db))
		}
		if c.Username != "" {
			options = append(options, redis.DialUsername(c.Username))
		}
		if c.ReadTimeout > 0 {
			options = append(options, redis.DialReadTimeout(c.ReadTimeout))
		}
		if c.WriteTimeout > 0 {
			options = append(options, redis.DialWriteTimeout(c.WriteTimeout))
		}
		c.pool = &redis.Pool{
			MaxIdle:     c.MaxIdle,
			MaxActive:   c.MaxActive,
			Wait:        true,
			IdleTimeout: c.IdleTimeout,
			Dial: func() (redis.Conn, error) {
				r, err := redis.Dial(c.Network, c.Address, options...)
				if err != nil {
					return nil, err
				}
				return r, nil
			},
		}
	}
	return c.pool
}

func (c *RedisClient) Get(key string) (string, error) {
	conn := c.GetPool().Get()

	defer conn.Close()
	res, err := conn.Do("GET", key)
	if err != nil {
		logger.Error("redis GET error", logger.Fields{"key": key, "res": res, "err": err})
	}
	if err != nil || res == nil {
		return "", err
	}
	return redis.String(res, err)
}

func (c *RedisClient) SetEx(key string, value string, ttl int) (bool, error) {
	var res interface{}
	var err error
	conn := c.GetPool().Get()
	defer conn.Close()
	if ttl <= 0 {
		res, err = conn.Do("SET", key, value)
	} else {
		res, err = conn.Do("SET", key, value, "EX", ttl)
	}
	if err != nil {
		logger.Error("redis SET error", logger.Fields{"key": key, "res": res, "err": err})
	}
	if err != nil || res == nil {
		return false, err
	}
	ok, err := redis.String(res, err)
	if ok == "OK" && err == nil {
		return true, nil
	}
	logger.Warning("redis SET fail", logger.Fields{"key": key, "res": res, "err": err, "ok": ok})
	return false, err
}

func (c *RedisClient) Del(key ...string) (int, error) {
	if len(key) <= 0 {
		return 0, nil
	}
	conn := c.GetPool().Get()
	defer conn.Close()
	res, err := conn.Do("DEL", key)
	if err != nil {
		logger.Error("redis DEL error", logger.Fields{"key": key, "res": res, "err": err})
	}
	if err != nil || res == nil {
		return 0, err
	}
	return redis.Int(res, err)
}

func (c *RedisClient) SAdd(key string, member ...interface{}) (int, error) {
	if len(member) <= 0 {
		return 0, nil
	}
	member = append([]interface{}{key}, member...)
	conn := c.GetPool().Get()
	defer conn.Close()
	res, err := conn.Do("SADD", member...)
	if err != nil {
		logger.Error("redis SADD error", logger.Fields{"key": key, "res": res, "err": err, "member": member})
	}
	if err != nil || res == nil {
		return 0, err
	}
	return redis.Int(res, err)
}

func (c *RedisClient) SRem(key string, member ...interface{}) (int, error) {
	if len(member) <= 0 {
		return 0, nil
	}
	member = append([]interface{}{key}, member...)
	conn := c.GetPool().Get()
	defer conn.Close()
	res, err := conn.Do("SREM", member...)
	if err != nil {
		logger.Error("redis SREM error", logger.Fields{"key": key, "res": res, "err": err, "member": member})
	}
	if err != nil || res == nil {
		return 0, err
	}
	return redis.Int(res, err)
}

func (c *RedisClient) SMembers(key string) ([]string, error) {
	conn := c.GetPool().Get()
	defer conn.Close()
	res, err := conn.Do("SMEMBERS", key)
	if err != nil {
		logger.Error("redis smembers error", logger.Fields{"key": key, "res": res, "err": err})
	}
	if err != nil || res == nil {
		return nil, err
	}
	return redis.Strings(res, err)
}

func (c *RedisClient) HMSet(key string, fields ...HMDto) (bool, error) {
	if len(fields) <= 0 {
		return false, nil
	}
	args := []interface{}{key}
	for _, field := range fields {
		args = append(args, field.Field, field.Value)
	}
	conn := c.GetPool().Get()
	defer conn.Close()
	res, err := conn.Do("HMSET", args...)
	if err != nil {
		logger.Error("redis HMSET error", logger.Fields{"key": key, "res": res, "err": err, "fields": fields})
	}
	if err != nil || res == nil {
		return false, err
	}
	ok, err := redis.String(res, err)
	if ok == "OK" && err == nil {
		return true, nil
	}
	logger.Warning("redis HMSET fail", logger.Fields{"key": key, "res": res, "err": err, "ok": ok})
	return false, err
}

func (c *RedisClient) HMGet(key string, fields ...interface{}) (map[string]string, error) {
	if len(fields) <= 0 {
		return nil, nil
	}
	args := append([]interface{}{key}, fields...)
	conn := c.GetPool().Get()
	defer conn.Close()
	res, err := conn.Do("HMGET", args...)
	if err != nil {
		logger.Error("redis HMGET error", logger.Fields{"key": key, "res": res, "err": err, "fields": fields})
	}
	if err != nil || res == nil {
		return nil, err
	}
	results, err := redis.Strings(res, err)
	if err != nil {
		return nil, err
	}
	result := map[string]string{}
	for i, v := range results {
		k := fields[i]
		result[k.(string)] = v
	}
	return result, err
}
func (c *RedisClient) HMGetOne(key string, field string) (string, error) {
	conn := c.GetPool().Get()
	defer conn.Close()
	res, err := conn.Do("HMGET", key, field)
	if err != nil {
		logger.Error("redis SREM error", logger.Fields{"key": key, "res": res, "err": err, "field": field})
	}
	if err != nil || res == nil {
		return "", err
	}
	results, err := redis.Strings(res, err)
	if err != nil || results == nil {
		return "", err
	}
	return results[0], err
}
func (c *RedisClient) HDel(key string, fields ...interface{}) (int, error) {
	if len(fields) <= 0 {
		return 0, nil
	}
	args := append([]interface{}{key}, fields...)
	conn := c.GetPool().Get()
	defer conn.Close()
	res, err := conn.Do("HDEL", args...)
	if err != nil {
		logger.Error("redis HDEL error", logger.Fields{"key": key, "res": res, "err": err, "fields": fields})
	}
	if err != nil || res == nil {
		return 0, err
	}
	return redis.Int(res, err)
}
