// @Author Ben.Zheng
// @DateTime 2022/8/16 11:02

package cache

import (
	"context"
	"net"
	"testing"
	"time"

	"github.com/go-redis/redis/v9"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"go.uber.org/fx"

	"github.com/benz9527/idk-doc/internal/pkg/file"
	"github.com/benz9527/idk-doc/internal/pkg/intf"
)

func Test_redis_conn(t *testing.T) {
	asserter := assert.New(t)

	// Private address.
	expectedAddr := "172.20.3.79"
	expectedPort := "6380"
	addr := net.JoinHostPort(expectedAddr, expectedPort)
	_, err := net.DialTimeout("tcp", addr, 2*time.Second)
	if err != nil {
		t.SkipNow() // Skip for other test environments.
	}

	v := viper.New()
	v.Set("redis", map[string]any{
		"mode": "single",
		"addresses": []string{
			addr,
		},
		"auth": map[string]any{
			"name":   "default",
			"secret": "benz2121",
		},
		"additional": map[string]any{
			"master_name": "",
		},
	})

	reader := file.NewSimpleReader(v)
	var client redis.UniversalClient
	asserter.NotPanics(func() {
		client = NewRedisClient(reader)
	})
	result, err := client.Do(context.TODO(), "acl", "list").Slice()
	asserter.Nil(err)
	asserter.True(len(result) > 0)
}

func Test_redis_client_addr_cmp(t *testing.T) {
	asserter := assert.New(t)

	// Private address.
	expectedAddr := "172.20.3.79"
	expectedPort := "6380"
	addr := net.JoinHostPort(expectedAddr, expectedPort)
	_, err := net.DialTimeout("tcp", addr, 2*time.Second)
	if err != nil {
		t.SkipNow() // Skip for other test environments.
	}

	v := viper.New()
	v.Set("redis", map[string]any{
		"mode": "single",
		"addresses": []string{
			addr,
		},
		"auth": map[string]any{
			"name":   "default",
			"secret": "benz2121",
		},
		"additional": map[string]any{
			"master_name": "",
		},
	})

	reader := file.NewSimpleReader(v)
	var c1, c2 redis.UniversalClient
	var opts []fx.Option
	opts = append(opts, fx.Provide(func() intf.IConfigurationReader {
		return reader
	}))
	opts = append(opts, fx.Provide(NewRedisClient))
	opts = append(opts, fx.Invoke(func(client redis.UniversalClient) {
		c1 = client
	}))
	opts = append(opts, fx.Invoke(func(client redis.UniversalClient) {
		c2 = client
	}))
	app := fx.New(opts...)
	err = app.Start(context.Background())
	asserter.Nil(err)
	defer app.Stop(context.Background())
	asserter.Equal(&c1, &c2)
}
