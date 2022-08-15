// @Author Ben.Zheng
// @DateTime 2022/8/15 13:32

package oss

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"

	"github.com/benz9527/idk-doc/internal/pkg/intf"
)

// References:
// https://docs.min.io/docs/minio-quickstart-guide.html
// https://docs.min.io/minio/baremetal/reference/minio-mc.html
// https://docs.min.io/minio/baremetal/reference/minio-mc-admin.html
// https://github.com/minio/minio-go

func newMinIOClient(cfgReader intf.IConfigurationReader) *minio.Client {

	endpoint, err := cfgReader.GetString("oss.endpoint")
	if err != nil {
		panic(err)
	}
	accessKeyId, err := cfgReader.GetString("oss.security.access_key_id")
	if err != nil {
		panic(err)
	}
	secretAccessKey, err := cfgReader.GetString("oss.security.secret_access_key")
	if err != nil {
		panic(err)
	}
	useSSL, err := cfgReader.GetBoolean("oss.security.enabled")
	if err != nil {
		panic(err)
	}

	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyId, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		panic(err)
	}
	return client
}
