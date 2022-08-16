// @Author Ben.Zheng
// @DateTime 2022/8/15 17:27

package cache

// References:
// https://github.com/go-redis/redis
// https://redis.uptrace.dev/guide/universal.html

// Preferring to use redis V7.

import (
	"errors"
	"strings"
	"time"

	"github.com/go-redis/redis/v9"
	mapUtils "github.com/mitchellh/mapstructure"

	"github.com/benz9527/idk-doc/internal/pkg/intf"
)

const (
	REDIS_SINGLE   = "single"
	REDIS_CLUSTER  = "cluster"
	REDIS_FAILOVER = "failover"
)

type redisCfg struct {
	Mode       string        `mapstructure:"mode"`
	Addresses  []string      `mapstructure:"addresses"`
	Auth       authCfg       `mapstructure:"auth"`
	Additional additionalCfg `mapstructure:"additional"`
}

type authCfg struct {
	Name   string `mapstructure:"name"`
	Secret string `mapstructure:"secret"`
}

type additionalCfg struct {
	MasterName string `mapstructure:"master_name"`
}

func NewRedisClient(cfgReader intf.IConfigurationReader) redis.UniversalClient {
	m, err := cfgReader.GetMap("redis")
	if err != nil {
		panic(err)
	}

	opts, err := modeVerify(m)
	if err != nil {
		panic(err)
	}

	client := redis.NewUniversalClient(opts)
	return client
}

func modeVerify(m map[string]any) (*redis.UniversalOptions, error) {
	var out redisCfg
	err := mapUtils.Decode(m, &out)
	if err != nil {
		return nil, err
	}

	var opts = redis.UniversalOptions{}
	opts.DB = 0
	if len(out.Auth.Secret) > 0 {
		opts.Password = out.Auth.Secret
		if len(out.Auth.Name) == 0 {
			opts.Username = "default"
		} else {
			opts.Username = out.Auth.Name
		}
	}

	switch strings.ToLower(out.Mode) {
	case REDIS_FAILOVER:
		if len(out.Additional.MasterName) == 0 {
			opts.MasterName = "idk-master"
		} else {
			opts.MasterName = out.Additional.MasterName
		}
		fallthrough
	case REDIS_CLUSTER:
		if len(out.Addresses) > 1 {
			opts.Addrs = out.Addresses
			break
		}
		fallthrough
	case REDIS_SINGLE:
		fallthrough
	default:
		if len(out.Addresses) <= 0 {
			return nil, errors.New("empty address to connect")
		}
		opts.Addrs = []string{out.Addresses[0]}
	}

	// Redis connection optimized.
	opts.WriteTimeout = 10 * time.Second
	opts.MaxRetries = 3
	opts.MaxRedirects = 10
	opts.DialTimeout = 5 * time.Second
	opts.ConnMaxIdleTime = 8 * time.Second
	opts.ConnMaxLifetime = 60 * time.Second
	opts.MaxIdleConns = 16
	opts.MinIdleConns = 4
	opts.PoolSize = 32

	return &opts, nil
}
