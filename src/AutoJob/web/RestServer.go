package web

import (
	"net/http"
	"AutoDeploy/job"
	"encoding/json"
	"log"
	"io"
	"fmt"
	"strings"
)

type Rest struct {

}

func (r *Rest) init()  {
	http.HandleFunc("/rest",restIndex)
	http.HandleFunc("/rest/saveOrUpdate", restSaveOrUpdate)
	//http.HandleFunc("/rest/delete", restdelete)
	//http.HandleFunc("/rest/deploy",restdeploy)
	//http.HandleFunc("/rest/loading",restloading)
}

type Jobs struct {
	JS map[string][]*job.DeployJob
}

func restIndex(res http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	name := globalSessions.SessionStart(res, req).Get("name")
	log.Println(req)
	for k, v := range req.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	jobs := job.Read()
	set := false
	for _,v := range jobs {
		if name == v.Name{
			v.Show = true
			set = true
		}
	}
	if !set && len(jobs) > 0 {
		jobs[0].Show = true
	}

	res.Header().Add("Access-Control-Allow-Origin","*")
	j := Jobs{}
	j.JS = make(map[string][]*job.DeployJob)
	j.JS["jobs"] = jobs
	bytes,err := json.Marshal(j)
	if err != nil{
		log.Println(err)
	}

	io.WriteString(res,string(bytes))
}

func restSaveOrUpdate(res http.ResponseWriter, req *http.Request) {
	res.Header().Add("Access-Control-Allow-Origin","*")
	defer func(){
		if err:=recover();err!=nil{
			log.Println("restSaveOrUpdate.err:",err)
			io.WriteString(res,"err")
		}
	}()
	req.ParseForm()
	jobJson := req.Form.Get("job")
	//log.Println("jobJson:",jobJson)
	var _job *job.DeployJob = new(job.DeployJob)
	err := json.Unmarshal([]byte(jobJson),_job)
	if err != nil{
		log.Println("err:",err)
	}
	logJob(_job)
	job.SaveOrUpdate(_job)
	io.WriteString(res,"ok")
}

func logJob(job *job.DeployJob)  {
	log.Println("Name:",job.Name)
	log.Println("Config:",job.Config)
	localBefore := job.LocalBefore
	for _,l:=range localBefore {
		log.Println("localBefore:",l)
	}
	RemoteBefore := job.RemoteBefore
	for _,l:=range RemoteBefore {
		log.Println("RemoteBefore:",l)
	}
	UploadJob := job.UploadJob
	log.Println("UploadJob:",UploadJob)
	LocalAfter := job.LocalAfter
	for _,l:=range LocalAfter {
		log.Println("LocalAfter:",l)
	}
	RemoteAfter := job.RemoteAfter
	for _,l:=range RemoteAfter {
		log.Println("RemoteAfter:",l)
	}
}
