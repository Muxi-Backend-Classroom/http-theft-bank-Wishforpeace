package main

import (
	"fmt"
	"net/http"
	"os"
	"io/ioutil"
	"mime/multipart"
	"bytes"
	"io"
)
func main() {
	client := &http.Client{}
	passport := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjb2RlIjoiMTIwIiwiaWF0IjoxNjM3MTQ2NDM1LCJuYmYiOjE2MzcxNDY0MzV9.mlggHuQMg4eooV1KBB9scFQE-J7018S5RLpXl-boWX4"
	method := "GET"
	url :="http://http-theft-bank.gtainccnu.muxixyz.com/api/v1/muxi/backend/computer/examination"
	request,_ := http.NewRequest(method,url,nil)
	request.Header.Add("Code","250")
	request.Header.Add("Passport",passport)
	response,_ := client.Do(request)
	body ,_ := ioutil.ReadAll(response.Body)
	fmt.Println(string(body))
	//最后一关
	//find the correct token
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)
	writer, _ := bodyWriter.CreateFormFile("file", "permute.go")
	file, _ := os.Open("./permute.go")
	io.Copy(writer, file)
	bodyWriter.Close()
	request, _ = http.NewRequest("POST", "http://http-theft-bank.gtainccnu.muxixyz.com/api/v1/muxi/backend/computer/examination", bodyBuf)
	request.Header.Add("Code", "120")
	request.Header.Add("Passport", passport)
	request.Header.Set("Content-Type", bodyWriter.FormDataContentType())
	response, _ = client.Do(request)
	contents, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(contents))
}
