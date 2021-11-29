package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://http-theft-bank.gtainccnu.muxixyz.com/api/v1/organization/code", nil)
	req.Header.Add("Code", "120")
	res, _ := client.Do(req)
	by, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(by))
	
}
