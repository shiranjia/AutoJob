package web

import (
	"net/http"
	"AutoDeploy/job"
	"encoding/json"
	"log"
	"io"
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
	j := toJob(req)
	if j.Name == ""{
		io.WriteString(res, "name must not null")
	}
	job.SaveOrUpdate(&j)
	session := globalSessions.SessionStart(res, req)
	session.Set("name",j.Name)
	http.Redirect(res,req,"/" ,http.StatusMovedPermanently)
	return
}
