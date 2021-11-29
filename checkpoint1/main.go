package main
import(
"fmt"
"net/http"
"io/ioutil"
)
func main(){
	url := "http://http-theft-bank.gtainccnu.muxixyz.com/api/v1/organization/code"
	method := "GET"

  	client := &http.Client {
  }
  	request,_ := http.NewRequest(method, url, nil)
  	request.Header.Add("Code","250")//添加code
  	response, _ := client.Do(request)
	body,_:=ioutil.ReadAll(response.Body)
	fmt.Println(string(body))
}
