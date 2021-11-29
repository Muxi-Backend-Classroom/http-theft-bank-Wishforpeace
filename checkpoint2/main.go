package main

import (
	"fmt"
// 	"github.com/Grand-Theft-Auto-In-CCNU-MUXI/hacker-support/httptool"
	"net/http"
	"io/ioutil"
	
)

func main() {
// 	req, err := httptool.NewRequest(
// 		"",
// 		"",
// 		"",
// 		httptool.DEFAULT, // 这里可能不是 DEFAULT，自己去翻阅文档
// 	)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	fmt.Println(req)

// 	// write your code below
// 	// ...
	client := &http.Client {
  }
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
  	fmt.Println(string(body))
}
