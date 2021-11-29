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
	 "strings"
)

func main() {
	//第二关，冲啊啊啊
	url :="http://http-theft-bank.gtainccnu.muxixyz.com/api/v1/bank/iris_recognition_gate"
	method := "GET"
	request,err := http.NewRequest(method,url,nil)
	if err != nil{
		panic(err)
	}
	request.Header.Add("Code","250")
	request.Header.Add("Passport",passport)
	response,_ :=client.Do(request)
	body,err := ioutil.ReadAll(response.Body)
	fmt.Println(string(body))
	//虹膜识别
	//下载图片
	url = "http://http-theft-bank.gtainccnu.muxixyz.com/api/v1/organization/iris_sample"
	method = "GET"
	request,err = http.NewRequest(method,url,nil)
	if err !=nil{
		fmt.Println(err)
		return
	}
	request.Header.Add("Code","250")
	request.Header.Add("Passport",passport)
	response,err = client.Do(request)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Header5\n")
	fmt.Println(response.Header)
	// type data struct{
	// 	extra_info string `json: "extra_info"`
	// }
	body,err = ioutil.ReadAll(response.Body)
	if err != nil{
		fmt.Println(err)
		return
	}
	defer response.Body.Close()

	var res Response
	if err = json.Unmarshal(body, &res); err != nil {
		fmt.Printf("unmarshal err, %s\n", err)
		return 
	}

	// fmt.Println(res.Data.ExtraInfo)
	content := res.Data.ExtraInfo
	// fmt.Printf("%v",content)
	// //图片转码
	datasource := []byte(content)
	ddd, _ := base64.StdEncoding.DecodeString(string(datasource))//成图片文件并把文件写入到buffer
	ioutil.WriteFile("./output.jpg", ddd, 0777)   //buffer输出到jpg文件中（不做处理，直接写到文件）
	//输入图片访问
	method = "POST"
	url = "http://http-theft-bank.gtainccnu.muxixyz.com/api/v1/bank/iris_recognition_gate"
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)//写入新的body
	writer,_:= bodyWriter.CreateFormFile("file","output.jpg")//生成的multipart主体。
	file,_ := os.Open("./output.jpg")
	io.Copy(writer, file)
	bodyWriter.Close()
	request,err =http.NewRequest(method,url,bodyBuf)
	if err != nil {
		fmt.Println(err)
		return
	}
	request.Header.Add("Code", "250")
	request.Header.Add("Passport", passport)
	request.Header.Set("Content-Type", bodyWriter.FormDataContentType())
	response, _ = client.Do(request)
// 	fmt.Printf("Header6\n")
// 	fmt.Println(response.Header)
	body, _ = ioutil.ReadAll(response.Body)
	fmt.Println(string(body))
}
