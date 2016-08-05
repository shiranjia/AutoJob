package web

import (
	"net/http"
	"log"
	"io/ioutil"
	"strings"
	j "encoding/json"
	"net/url"
)

func httpDial(url_,method string,head,json map[string]string) string  {
	var client http.Client
	request ,err := http.NewRequest(method,url_,nil)
	checkErr(err)
	for k,v := range head {
		request.Header.Set(k,v)
	}

	request.PostForm = url.Values{"cityId":[]string{"36"}}

	j,_ := j.Marshal(json)
	b := strings.NewReader(string(j))
	request.Body =  ioutil.NopCloser(b)

	log.Println("Header:",request.Header)
	log.Println("Body:",request.Body)
	response,err := client.Do(request)
	checkErr(err)
	log.Println("response.Status:",response.Status)
	if response.Status == "200 OK" {
		res,err := ioutil.ReadAll(response.Body)
		checkErr(err)
		return string(res)
	}
	return ""
}

func httpDialBody(url_,method string,head map[string]string,bo string) string  {
	var client http.Client
	request ,err := http.NewRequest(method,url_,nil)
	checkErr(err)
	for k,v := range head {
		request.Header.Set(k,v)
	}

	b := strings.NewReader(bo)
	request.Body =  ioutil.NopCloser(b)

	log.Println("Header:",request.Header)
	log.Println("Body:",request.Body)
	response,err := client.Do(request)
	checkErr(err)
	log.Println("response.Status:",response.Status)
	if response.Status == "200 OK" {
		res,err := ioutil.ReadAll(response.Body)
		checkErr(err)
		return string(res)
	}
	return ""
}

func checkErr(err error)  {
	if err != nil{
		log.Println(err)
	}
}
