package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// 发送GET请求
// url：         请求地址
// response：    请求返回的内容
func GetController(url string) string {

	// 超时时间：5秒
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	var buffer [512]byte
	result := bytes.NewBuffer(nil)
	for {
		n, err := resp.Body.Read(buffer[0:])
		result.Write(buffer[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}

	return result.String()
}

// 发送POST请求
// url：         请求地址
// data：        POST请求提交的数据
// contentType： 请求体格式，如：application/json
// content：     请求放回的内容
// json
// contentType := "application/json"
// data := `{"name":"枯藤","age":18}`

func PostController(url string, data interface{}, contentType string) string {
	// 超时时间：5秒
	client := &http.Client{Timeout: 10 * time.Second}
	contentType = "application/json;charset=utf-8"
	jsonStr, _ := json.Marshal(data)
	resp, err := client.Post(url, contentType, bytes.NewBuffer(jsonStr))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	result, _ := ioutil.ReadAll(resp.Body)
	return string(result)
}

// 发送POST请求,header添加信息
// url：         请求地址
// data：        POST请求提交的数据
// contentType： 请求体格式，如：application/x-www-form-urlencoded
// content：     请求放回的内容
// 表单数据
//contentType := "application/x-www-form-urlencoded"
//data := "name=枯藤&age=18"

func PostFormController(url string, data string, contentType string) string {
	client := &http.Client{Timeout: 10 * time.Second}
	contentType = "application/x-www-form-urlencoded"
	req, err := http.NewRequest("POST", url, strings.NewReader(data))
	if err != nil {
		panic(err)
	}
	// 如果采用的是BasicAuth验证，可以这么设置header
	//req.SetBasicAuth(username, token)
	req.Header.Set("Content-Type", contentType)
	fmt.Println("header:", req.Header)
	resp, err := client.Do(req)
	defer resp.Body.Close()
	result, _ := ioutil.ReadAll(resp.Body)
	return string(result)
}
