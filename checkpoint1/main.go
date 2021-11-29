package main
import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/base64"//base64 解密
	"log"
	"crypto/aes"
	"crypto/cipher"//加密
	"bytes"
	"strings"
	"encoding/json"
	"os"
	"io"
	"mime/multipart"
)


type Response struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Data Data `json:"data"`
}
type Data struct {
	Text string `json:"text"`
	ExtraInfo string `json:"extra_info"`
}


func main() {

	url := "http://http-theft-bank.gtainccnu.muxixyz.com/api/v1/organization/code"
	method := "GET"

  	client := &http.Client {
  }
  	request,err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
    	return
  	}
  	request.Header.Add("Code","250")//添加code
  	response, err := client.Do(request)
  	if err != nil {
    	fmt.Println(err)
    	return
  	}
	fmt.Printf("Header1\n")
	fmt.Println(response.Header)

	//passport
	passport := response.Header["Passport"][0]//保存passport
	url = "http://http-theft-bank.gtainccnu.muxixyz.com/api/v1/organization/secret_key"
	
	request,err =http.NewRequest(method,url,nil)
	request.Header.Add("Code","250")
	request.Header.Add("Passport",passport)//将passport加入请求头
	if err != nil{
		fmt.Println(err)
		return
	}
	response,err = client.Do(request)
	if err != nil {
		panic(err)
	}
 	body, err := ioutil.ReadAll(response.Body)
  	if err != nil {
    	fmt.Println(err)
    	return
  	}
	info := "c2VjcmV0X2tleTpNdXhpU3R1ZGlvMjAzMzA0LCBlcnJvcl9jb2RlOmZvciB7Z28gZnVuYygpe3RpbWUuU2xlZXAoMSp0aW1lLkhvdXIpfSgpfQ=="
	fmt.Printf("Header2\n")
	fmt.Println(response.Header)
  	fmt.Println(string(body))
	fmt.Println("密文内容：\t")
	fmt.Println(info)
	//解密
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
	url = "http://http-theft-bank.gtainccnu.muxixyz.com/api/v1/bank/gate"
	method = "PUT"
	type Body struct{
		Content string
	}
	BODY := Body {
		Content: error,
	}//输入新结构体
	configData,_ :=json.Marshal((BODY))//JSON序列化
	payload := strings.NewReader(string(configData))//封装
	request,_= http.NewRequest(method,url,payload)
	request.Header.Add("Code","250")
	request.Header.Add("Passport",passport)
	response,err = client.Do(request)
	if err !=nil{
		panic(err)
	}
	fmt.Printf("Header3\n")
	fmt.Println(response.Header)
	body,_ = ioutil.ReadAll(response.Body)
	fmt.Println(string(body))
	//第二关，冲啊啊啊
	url ="http://http-theft-bank.gtainccnu.muxixyz.com/api/v1/bank/iris_recognition_gate"
	method = "GET"
	request,err = http.NewRequest(method,url,nil)
	if err != nil{
		panic(err)
	}
	request.Header.Add("Code","250")
	request.Header.Add("Passport",passport)
	response,_ =client.Do(request)
	body,err = ioutil.ReadAll(response.Body)
	fmt.Printf("Header4\n")
	fmt.Println(response.Header)
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
	fmt.Printf("Header6\n")
	fmt.Println(response.Header)
	body, _ = ioutil.ReadAll(response.Body)
	fmt.Println(string(body))
	//最后一道门了！
	method = "GET"
	url="http://http-theft-bank.gtainccnu.muxixyz.com/api/v1/muxi/backend/computer/examination"
	request,_ = http.NewRequest(method,url,nil)
	request.Header.Add("Code","250")
	request.Header.Add("Passport",passport)
	response,_ = client.Do(request)
	body ,_ = ioutil.ReadAll(response.Body)
	fmt.Println(string(body))
	//最后一关
	bodyBuf = &bytes.Buffer{}
	bodyWriter = multipart.NewWriter(bodyBuf)
	writer, _ = bodyWriter.CreateFormFile("file", "permute.go")
	file, _ = os.Open("./permute.go")
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
