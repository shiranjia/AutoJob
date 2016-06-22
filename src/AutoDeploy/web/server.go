package web

import (
	"io"
	"log"
	"net/http"
	"html/template"
	"AutoDeploy/job"
	"AutoDeploy/commons"
	"strings"
	"github.com/astaxie/session"
	_ "github.com/astaxie/session/providers/memory"
)

const(
	temp  = "resources/views/welcome.html"
)

var globalSessions *session.Manager

//然后在init函数中初始化
func init() {
	globalSessions, _ = session.NewManager("memory", "gosessionid", 3600)
	go globalSessions.GC()
}


func Service() {
	http.HandleFunc("/",index)
	http.HandleFunc("/saveOrUpdate", saveOrUpdate)
	http.HandleFunc("/delete", delete)
	http.HandleFunc("/deploy",deploy)
	http.HandleFunc("/loading",loading)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func index(res http.ResponseWriter, req *http.Request) {
	//t := getTemplateFromFile()
	t := getTemplateFromString()
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
	err := t.Execute(res,jobs)
	if err != nil{
		log.Fatal(err)
	}
}

func getTemplateFromFile() *template.Template  {
	t,err := template.ParseFiles(temp)
	if err != nil{
		log.Println(err)
		return nil
	}
	return t
}

func getTemplateFromString() *template.Template {
	t := template.New("deploy")
	t.Parse(commons.Html)
	return t;
}

func saveOrUpdate(res http.ResponseWriter, req *http.Request) {
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

func delete(res http.ResponseWriter, req *http.Request) {
	j := toJob(req)
	job.Delete(j)
	http.Redirect(res,req,"/",http.StatusMovedPermanently)
}

func deploy(res http.ResponseWriter, req *http.Request) {
	deploy := toJob(req)
	go deploy.Deploy()
	//log.Println(req.Form)
	http.Redirect(res,req,"/loading",http.StatusMovedPermanently)
}

func loading(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "<html><head><title>deploying...</title></head><body><img src='http://img12.360buyimg.com/piao/jfs/t2971/357/809642283/282981/3ea2914a/576926baNf53c63d4.gif'><a href='/'>return</a></body></html>")
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
		//log.Println("remote shell:",c)
		c = strings.Replace(c,"\r","",-1)
		if "" != c {
			remoteBeforeCom = append(remoteBeforeCom,&job.RemoteComm{false,c})
		}
	}

	localBefore := req.Form.Get("LocalBefore.Command")
	var localBeforeCom []*job.LocalComm
	lb := strings.Split(localBefore,"\n")
	for _,c := range lb {
		//log.Println(c)
		com := strings.Split(c,";")
		var beforeCom job.LocalComm;
		beforeCom.IsGo = false
		for i,v := range com {
			if v == ""{
				continue
			}
			if i == 0 {
				beforeCom.Path = v
			}else if i == 1 {
				beforeCom.Command = v
			}else if i == 2 {
				args := strings.Split(v," ")
				for _,v := range args {
					if "" != v && "\r" != v{
						beforeCom.Args = append(beforeCom.Args,v)
					}
				}
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
		//log.Println("remote shell:",c)
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
			if v == ""{
				continue
			}
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

