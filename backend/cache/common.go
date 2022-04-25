package cache

import (
	"bookmark/logger"
	"errors"
	"github.com/go-redis/redis"
	"strconv"
	"strings"
	"time"
)

func (c *RedisClient)CommonKeyGen(fields []string) string {
	return strings.Join(fields, ":")
}

// BookmarkKeyGen 为bookmark数据生成key, 若传入num参数，生成形如 bookmark:{uid}:{cid}:{num}的key，用于获取redis中
// 的一条bookmark记录，若不传num参数，则生成形如 bookmark:{uid}:{cid}:count的key，用于获取指定uid和cid下的数据条数
func (c *RedisClient)BookmarkKeyGen(uid uint, cid uint, num ...int) (string, error) {
	if len(num) > 1 {
		return "", errors.New("Invalid field format: num. ")
	}
	fields := []string{"bookmark", strconv.Itoa(int(uid)), strconv.Itoa(int(cid))}
	if len(num) != 0 {
		fields = append(fields, strconv.Itoa(num[0]))
	} else {
		fields = append(fields, "count")
	}
	return c.CommonKeyGen(fields), nil
}

func (c *RedisClient) SetHash(key string, fields map[string]interface{}) (string, error) {
	val, err := c.client.HMSet(key, fields).Result()
	if err != nil {
		logger.SugarLogger.Errorf("Redis set hash failed. error: %s", err.Error())
	}
	return val, err
}

func (c *RedisClient) BatchSetHash(keys []string, objs []map[string]interface{}) error {
	for index, key := range keys {
		c.pipeline.HMSet(key, objs[index])
	}
	_, err := c.pipeline.Exec()
	return err
}

func (c *RedisClient) BatchGetHash(keys []string) ([]map[string]string, error) {
	for _, key := range keys {
		c.pipeline.HGetAll(key)
	}
	cmds, err := c.pipeline.Exec()
	if err != nil {
		logger.SugarLogger.Errorf("Redis get hashes failed. error: %s", err.Error())
		return nil, err
	}
	var result []map[string]string
	for _, cmd := range cmds {
		if stringStringMap, ok := cmd.(*redis.StringStringMapCmd); ok {
			res, err := stringStringMap.Result()
			if err != nil {
				logger.SugarLogger.Errorf("Redis get hashes failed. error: %s", err.Error())
				return nil, err
			}
			result = append(result, res)
		} else {
			logger.SugarLogger.Errorf("Redis get hashes failed.")
		}
	}
	return result, nil
}

func (c *RedisClient)BatchDelete(keys []string) error {
	for _, key := range keys {
		c.pipeline.Del(key)
	}
	_, err := c.pipeline.Exec()
	return err
}

func (c *RedisClient)GetKeys(pattern string) ([]string, error) {
	var (
		res, keys []string
		cursor uint64
		err error
	)
	for {
		res, cursor, err = c.client.Scan(cursor, pattern, 20).Result()
		if err != nil {
			return nil, err
		}
		keys = append(keys, res...)
		if cursor == 0 {
			break
		}
	}
	return keys, nil
}

func (c *RedisClient)GetString(key string) (string, error) {
	return c.client.Get(key).Result()
}

func (c *RedisClient)SetString(key string, value string, expiration time.Duration) {
	c.client.Set(key, value, expiration)
}

func (c *RedisClient)Delete(key string) {
	c.client.Del(key)
}

