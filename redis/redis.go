package redis

import (
	goRedis "github.com/go-redis/redis/v8"
)

// RedisClient is a redis client which wraps single or cluster client
type RedisClient interface {
	Keys(pattern string) *goRedis.StringSliceCmd
	Scan(cursor uint64, match string, count int64) *goRedis.ScanCmd
	Type(key string) *goRedis.StatusCmd
	TTL(key string) *goRedis.DurationCmd
	Get(key string) *goRedis.StringCmd
	LRange(key string, start, stop int64) *goRedis.StringSliceCmd
	SMembers(key string) *goRedis.StringSliceCmd
	ZRangeWithScores(key string, start, stop int64) *goRedis.ZSliceCmd
	HKeys(key string) *goRedis.StringSliceCmd
	HGet(key, field string) *goRedis.StringCmd
	Process(cmd goRedis.Cmder) error
	Do(args ...interface{}) *goRedis.Cmd
	Info(section ...string) *goRedis.StringCmd
}
