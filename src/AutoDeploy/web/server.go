package web

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"html/template"
	"AutoDeploy/job"
	"AutoDeploy/commons"
	"strings"
)

const temp  = "resources/views/welcome.html"

func Service() {
	http.HandleFunc("/index",index)
	http.HandleFunc("/saveOrUpdate", saveOrUpdate)
	http.HandleFunc("/delete", delete)
	http.HandleFunc("/deploy",deploy)

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func index(res http.ResponseWriter, req *http.Request) {
	t,err := template.ParseFiles(temp)
	if err != nil{
		log.Fatal(err)
	}
	jobs := job.Read()
	err = t.Execute(res,jobs)
	//template.Must(t,errors.New("template has error"))
	if err != nil{
		log.Fatal(err)
	}
}

func saveOrUpdate(res http.ResponseWriter, req *http.Request) {
	j := toJob(req)
	if j.Name == ""{
		io.WriteString(res, "name must not null")
	}
	job.SaveOrUpdate(&j)
	http.Redirect(res,req,"/index",http.StatusMovedPermanently)
	/*t,err := template.ParseFiles(temp)
	if err != nil{
		log.Fatal(err)
	}
	err = t.Execute(res,jobs)
	if err != nil{
		log.Fatal(err)
	}*/
	return
}

func delete(res http.ResponseWriter, req *http.Request) {
	j := toJob(req)
	job.Delete(j)
	http.Redirect(res,req,"/index",http.StatusMovedPermanently)
}

func deploy(res http.ResponseWriter, req *http.Request) {
	deploy := toJob(req)
	go deploy.Deploy()

	fmt.Println(req.Header)
	fmt.Println(req.Form)
	io.WriteString(res, "success")
}

func toJob(req *http.Request) job.DeployJob {
	defer func(){
		if err := recover();err != nil{
			log.Println("transfer to job error:",err)
		}
	}()

	req.ParseForm()
	var deploy job.DeployJob
	deploy.Name = req.Form.Get("deploy.Name")
	user := req.Form.Get("config.User")
	password := req.Form.Get("config.Password")
	ip := req.Form.Get("config.ip")
	var config commons.SSHConfig
	config.User = user
	config.Password = password
	config.Ip = ip

	remoteBefore := req.Form.Get("RemoteBefore.Command")
	rb := strings.Split(remoteBefore,"\n")
	var remoteBeforeCom []*job.RemoteComm
	for _,c := range rb {
		log.Println("remote shell:",c)
		c = strings.Replace(c,"\r","",-1)
		if "" != c {
			remoteBeforeCom = append(remoteBeforeCom,&job.RemoteComm{false,c})
		}
	}

	localBefore := req.Form.Get("LocalBefore.Command")
	var localBeforeCom []*job.LocalComm
	lb := strings.Split(localBefore,"\n")
	for _,c := range lb {
		com := strings.Split(c,";")
		var beforeCom job.LocalComm;
		beforeCom.IsGo = false
		for i,v := range com {
			if i == 0 {
				beforeCom.Path = v
			}else if i == 1 {
				beforeCom.Command = v
			}else if i == 2 {
				args := strings.Split(v," ")
				beforeCom.Args = args
			}
		}
		localBeforeCom = append(localBeforeCom,&beforeCom)
	}

	uploadLocalPath := req.Form.Get("Upload.Path")
	uploadRemotePath := req.Form.Get("Upload.RemotePath")
	uploadJob := &job.UploadJob{uploadLocalPath,uploadRemotePath}

	remoteAfter := req.Form.Get("RemoteAfter.Command")
	ra := strings.Split(remoteAfter,"\n")
	var remoteAfterCom []*job.RemoteComm
	for _,c := range ra {
		log.Println("remote shell:",c)
		c = strings.Replace(c,"\r","",-1)
		if "" != c {
			remoteAfterCom = append(remoteAfterCom,&job.RemoteComm{false,c})
		}
	}

	localAfter := req.Form.Get("LocalAfter.Command")
	var localAfterCom []*job.LocalComm
	la := strings.Split(localAfter,"\n")
	for _,c := range la {
		com := strings.Split(c,";")
		var afterCom job.LocalComm;
		afterCom.IsGo = false
		for i,v := range com {
			if i == 0 {
				afterCom.Path = v
			}else if i == 1 {
				afterCom.Command = v
			}else if i == 2 {
				args := strings.Split(v," ")
				afterCom.Args = args
			}
		}
		localAfterCom = append(localAfterCom,&afterCom)
	}

	deploy.Config = &config
	deploy.RemoteBefore = remoteBeforeCom
	deploy.LocalBefore = localBeforeCom
	deploy.UploadJob = uploadJob
	deploy.RemoteAfter = remoteAfterCom
	deploy.LocalAfter = localAfterCom

	return deploy
}

