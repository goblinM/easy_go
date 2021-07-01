package main

import (
	"bytes"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"io/ioutil"
	"net/http"
	"os"
)

const ACCESSKEYID = ""
const ACCESSKEYSECRET = ""
const SOFTBUCKETNAME = ""
const LIVEBUCKETNAME = ""
const ENDPOINT = ""

// oss
type (
	OssInterface interface {
		UploadFile(key, localPath string) error   // 上传文件
		DownloadFile(key, localPath string) error // 下载文件
		UploadFileReader(key string, fileByte []byte) error
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

// 文件上传
func (d defaultOss) UploadFile(key, localPath string) error {
	//panic("implement me")
	//local_path := "D://working/pam/test2.txt"
	//filename := "test2.txt"
	//key := fmt.Sprintf("soft/custom/material/%s", filename)
	err := d.Bucket.PutObjectFromFile(key, localPath)
	if err != nil {
		fmt.Println("go oss 文件上传失败")
		return err
	} else {
		fmt.Println("go oss 文件上传成功")

		if err := os.Remove(localPath); err != nil {
			fmt.Println(err)
		}
	}
	return nil
}

// 二进制上传
func (d defaultOss) UploadFileReader(key string, fileByte []byte) error {
	err := d.Bucket.PutObject(key, bytes.NewReader([]byte(fileByte)))
	if err != nil {
		fmt.Println(err)
		fmt.Println("go oss 二进制文件上传失败")
		return err
	} else {
		fmt.Println("go oss 二进制文件上传成功")
	}
	return nil
}
func (d defaultOss) DownloadFile(key, localPath string) error {

	err := d.Bucket.DownloadFile(key, localPath, 10*1024*1024)
	if err != nil {
		return err
	} else {
		fmt.Println("文件从oss下载成功")
	}
	return nil
}

// 通过form-data上传文件，文件名：file
func handleUploadFile(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(100)
	mForm := r.MultipartForm
	fmt.Println("mForm.Value", mForm.Value)
	files := mForm.File
	parameters := mForm.Value // 不是file类型的其他表单参数
	fmt.Println(parameters)
	for k, _ := range files {
		// k is the key of file part
		file, fileHeader, err := r.FormFile(k)
		if err != nil {
			fmt.Println("inovke FormFile error:", err)
		}
		defer file.Close()
		fmt.Printf("the uploaded file: name[%s], size[%d], header[%#v]\n",
			fileHeader.Filename, fileHeader.Size, fileHeader.Header)
		key := fmt.Sprintf("soft/custom/material/%s", fileHeader.Filename)
		// 文件地址上传oss
		//localFileName := "D://working/pam/" + fileHeader.Filename
		//out, err := os.Create(localFileName)
		//if err != nil {
		//	fmt.Printf("failed to open the file %s for writing", localFileName)
		//	return nil,nil
		//}
		//defer out.Close()
		//_, err = io.Copy(out, file)
		//if err != nil {
		//	fmt.Printf("copy file err:%s\n", err)
		//	return nil,nil
		//}
		//fmt.Printf("file %s uploaded ok\n", fileHeader.Filename)
		//if err := NewSoftOssServer().UploadFile(key, localFileName); err!=nil {
		//	fmt.Println(err)
		//}
		// 二进制上传oss
		fileByte, err := ioutil.ReadAll(file) //获取上传文件字节流
		if err != nil {
			fmt.Println(err)
		}
		if err := NewSoftOssServer().UploadFileReader(key, fileByte); err != nil {
			fmt.Println(err)
		}
	}
}

// 上传的写法
func ossDemo() {
	// 文件路径上传
	local_path := "D://working/pam/test.txt"
	filename := "test2.txt"
	key := fmt.Sprintf("soft/custom/material/%s", filename)
	if err := NewSoftOssServer().UploadFile(key, local_path); err != nil {
		fmt.Println(err)
	}

	// 二进制上传
	file, err := os.Open(local_path)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()                    //关闭文件
	fileByte, err := ioutil.ReadAll(file) //获取上传文件字节流
	if err != nil {
		fmt.Println(err)
	}
	if err := NewSoftOssServer().UploadFileReader(key, fileByte); err != nil {
		fmt.Println(err)
	}

	// 文件下载
	if err := NewSoftOssServer().DownloadFile(key, local_path); err != nil {
		fmt.Println(err)
	}
}

func main() {
	http.HandleFunc("/upload", handleUploadFile)
	http.ListenAndServe(":8080", nil)
}
