package helper

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

// 发送GET请求
// url：         请求地址
// response：    请求返回的内容
func Get(url string) (result string, err error) {
	// 超时时间：5秒
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	var buffer [512]byte
	byteSlice := bytes.NewBuffer(nil)
	for {
		n, err := resp.Body.Read(buffer[0:])
		byteSlice.Write(buffer[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			return "", err
		}
	}

	result = byteSlice.String()

	return
}

// 发送POST请求
// url：         请求地址
// data：        POST请求提交的数据
// contentType： 请求体格式，如：application/json
// content：     请求放回的内容
func Post(url string, data interface{}, contentType string) (result string, err error) {
	// 超时时间：5秒
	client := &http.Client{Timeout: 5 * time.Second}
	jsonStr, _ := json.Marshal(data)
	resp, err := client.Post(url, contentType, bytes.NewBuffer(jsonStr))
	if err != nil {
		return
	}
	defer resp.Body.Close()

	res, err := ioutil.ReadAll(resp.Body)
	result = string(res)
	return
}
