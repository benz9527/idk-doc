// @Author Ben.Zheng
// @DateTime 2022/8/15 15:23

package oss

import (
	"context"
	"net"
	"testing"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"

	"github.com/benz9527/idk-doc/internal/pkg/file"
)

func Test_minio_client_creation(t *testing.T) {
	asserter := assert.New(t)

	// Private address.
	expectedAddr := "172.20.3.79"
	expectedPort := "9000"
	addr := net.JoinHostPort(expectedAddr, expectedPort)
	_, err := net.DialTimeout("tcp", addr, 2*time.Second)
	if err != nil {
		t.SkipNow() // Skip for other test environments.
	}

	v := viper.New()
	v.Set("oss.endpoint", expectedAddr+":"+expectedPort)
	v.Set("oss.security.enabled", false)
	v.Set("oss.security.access_key_id", "ben")
	v.Set("oss.security.secret_access_key", "benz2121")

	reader := file.NewSimpleReader(v)
	var client *minio.Client
	asserter.NotPanics(func() {
		client = newMinIOClient(reader)
	})
	bks, err := client.ListBuckets(context.TODO())
	asserter.Nil(err)
	asserter.Equal(1, len(bks))
}
