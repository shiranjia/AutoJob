package web

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"html/template"
	"github.com/pkg/errors"
)

func Service() {
	http.HandleFunc("/index",index)
	http.HandleFunc("/get", get)
	http.HandleFunc("/post", post)
	//dh := &defaultHandler{template}
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func index(res http.ResponseWriter, req *http.Request) {
	fmt.Println(req.Header)
	fmt.Println(req.Method)
	//t := template.New("welcome")
	t,err := template.ParseFiles("resources/views/welcome.tmpl")
	if err != nil{
		log.Fatal(err)
	}
	data :=  make(map[string]interface{})
	data["a"] = "string aaa"
	err = t.Execute(res,data)
	template.Must(t,errors.New("template has error"))
	if err != nil{
		log.Fatal(err)
	}
}

func get(res http.ResponseWriter, req *http.Request) {
	fmt.Println(req.Header)
	fmt.Println(req.Method)
	io.WriteString(res, "get methos")
}

func post(res http.ResponseWriter, req *http.Request) {
	fmt.Println(req.Header)
	io.WriteString(res, "post methos")
}

type defaultHandler struct {
	html string
}

func (d *defaultHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	log.Println(req.RequestURI)
	if req.RequestURI == "/" {
		res.Write(toByte(d.html))
	}
}

func toByte(o string) []byte {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()
	return []byte(o)

}
