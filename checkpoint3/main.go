package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
	"log"
)

func main() {
	passport := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjb2RlIjoiMTIwIiwiaWF0IjoxNjM3MTQ2NDM1LCJuYmYiOjE2MzcxNDY0MzV9.mlggHuQMg4eooV1KBB9scFQE-J7018S5RLpXl-boWX4"
	info := "c2VjcmV0X2tleTpNdXhpU3R1ZGlvMjAzMzA0LCBlcnJvcl9jb2RlOmZvciB7Z28gZnVuYygpe3RpbWUuU2xlZXAoMSp0aW1lLkhvdXIpfSgpfQ=="
	decode,err :=base64.StdEncoding.DecodeString(info)
	if err !=nil {
		log.Fatalln(err)
	}
	fmt.Println(string(decode),"\t")
	//记下来secret_key:MuxiStudio203304, error_code:for {go func(){time.Sleep(1*time.Hour)}()} 
	secret_key :="MuxiStudio203304"
	error_code :="for {go func(){time.Sleep(1*time.Hour)}()}"
	//Encryption error code
	origData :=[]byte(error_code)
	k :=[]byte(secret_key)
	ivT := make([]byte, aes.BlockSize+len(origData))
	iv := ivT[:aes.BlockSize]
	//分组密钥
	block,_ :=aes.NewCipher(k)
	//获取密钥块长度
	blockSize := block.BlockSize()
	// 补全码
	padding := blockSize -len(origData)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)},padding)
	origData = append(origData,padText...)
    
    // 加密模式
    blockMode := cipher.NewCBCEncrypter(block, iv)
    // 创建数组
    cryted := make([]byte, len(origData))
    // 加密
    blockMode.CryptBlocks(cryted, origData)
    error:=  base64.StdEncoding.EncodeToString(cryted)
	fmt.Println("加密后的error：",error)
	//获得加密的error_code,成功进入
	//Post 发送json 格式
	url := "http://http-theft-bank.gtainccnu.muxixyz.com/api/v1/bank/gate"
	method := "PUT"
	type Body struct{
		Content string
	}
	BODY := Body {
		Content: error,
	}//输入新结构体
	configData,_ :=json.Marshal((BODY))//JSON序列化
	payload := strings.NewReader(string(configData))//封装
	request,_:= http.NewRequest(method,url,payload)
	request.Header.Add("Code","250")
	request.Header.Add("Passport",passport)
	response,err := client.Do(request)
	if err !=nil{
		panic(err)
	}
// 	fmt.Printf("Header3\n")
// 	fmt.Println(response.Header)
	body,_ := ioutil.ReadAll(response.Body)
	fmt.Println(string(body))
}
