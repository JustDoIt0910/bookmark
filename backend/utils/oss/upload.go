package oss

import (
	"bookmark/config"
	"context"
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"io/fs"
	"path/filepath"
	"regexp"
)

func Upload() bool {
	bucket := config.GlobalConfig.GetString("oss.bucket")
	accessKey := config.GlobalConfig.GetString("oss.accessKey")
	secretKey := config.GlobalConfig.GetString("oss.secretKey")

	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)
	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Zone = &storage.ZoneHuadong
	// 是否使用https域名
	cfg.UseHTTPS = false
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false
	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{}

	var uploadSuccess = true
	_ = filepath.Walk("upload", func(path string, info fs.FileInfo, err error) error {
		if !info.IsDir() {
			ok, _ := regexp.MatchString("category_[a-zA-Z]+_bg[1-9]+\\.[jpg|png|jpeg]", info.Name())
			if !ok {
				return nil
			}
			err := formUploader.PutFile(context.Background(), &ret, upToken, info.Name(), path, &putExtra)
			if err != nil {
				fmt.Println(err)
				uploadSuccess = false
			}
		}
		return nil
	})
	return uploadSuccess
}
