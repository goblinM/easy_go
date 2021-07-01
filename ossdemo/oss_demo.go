package main

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

const ACCESSKEYID = ""
const ACCESSKEYSECRET = ""
const SOFTBUCKETNAME = ""
const LIVEBUCKETNAME = ""
const ENDPOINT = ""

// oss
type (
	OssInterface interface {
		UploadFile() error   // 上传文件
		DownloadFile() error // 下载文件
	}

	defaultOss struct {
		Client *oss.Client
		Bucket *oss.Bucket
	}
)

func Init(accessKeyId, accessKeySecret, bucketName, endPoint string) (client *oss.Client, bucket *oss.Bucket, err error) {
	client, err = oss.New(endPoint, accessKeyId, accessKeySecret)
	if err != nil {
		return
	}
	if bucketName != "" {
		bucket, err = client.Bucket(bucketName)
		if err != nil {
			return
		}
	}
	return
}

func NewSoftOssServer() OssInterface {
	//ossInstance := GetOssInstance()
	client, bucket, _ := Init(ACCESSKEYID, ACCESSKEYSECRET, SOFTBUCKETNAME, ENDPOINT)
	return defaultOss{
		Client: client,
		Bucket: bucket,
	}
}

func NewLiveOssServer() OssInterface {
	//ossInstance := GetOssInstance()
	client, bucket, _ := Init(ACCESSKEYID, ACCESSKEYSECRET, LIVEBUCKETNAME, ENDPOINT)
	return defaultOss{
		Client: client,
		Bucket: bucket,
	}
}

func (d defaultOss) UploadFile() error {
	//panic("implement me")
	localPath := "D://xxx/xxx/test2.txt"
	filename := "test2.txt"
	key := fmt.Sprintf("soft/custom/material/%s", filename)
	err := d.Bucket.PutObjectFromFile(key, localPath)
	if err != nil {
		fmt.Println(err)
		return err
	} else {
		fmt.Println("go oss 文件上传成功")
	}
	return nil
}

func (d defaultOss) DownloadFile() error {
	localPath := "D://xxx/xxx/test.txt"
	filename := "test2.txt"
	key := fmt.Sprintf("soft/custom/material/%s", filename)
	err := d.Bucket.DownloadFile(key, localPath, 10*1024*1024)
	if err != nil {
		return err
	} else {
		fmt.Println("文件从oss下载成功")
	}
	return nil
}

func main() {
	if err := NewSoftOssServer().UploadFile(); err != nil {
		fmt.Println(err)
	}
	if err := NewSoftOssServer().DownloadFile(); err != nil {
		fmt.Println(err)
	}
}
