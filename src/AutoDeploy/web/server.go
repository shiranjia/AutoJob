package web

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"html/template"
	"AutoDeploy/job"
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
	t,err := template.ParseFiles("resources/views/welcome.html")
	if err != nil{
		log.Fatal(err)
	}
	jobs := job.Read("data")
	err = t.Execute(res,jobs)
	//template.Must(t,errors.New("template has error"))
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
