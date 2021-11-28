package main

import (
	"fmt"
	"net/http"
)

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
	fmt.Printf("Header\n")
	fmt.Println(response.Header)
}
