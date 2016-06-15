package main

import (
	"fmt"
	"os"
	"AutoDeploy/commons"
	"AutoDeploy/job"
	"AutoDeploy/web"
)


func main() {
	fmt.Println("gopath=", os.Getenv("gopath"))
	//demo.Route()
	web.Service()
	//build()
	//a := "org.apache.coyote.AbstractProtocol.start Starting ProtocolHandler [\"ajp-nio-8009\"]"
	//fmt.Println(strings.Contains(a,"startup in"))
	//db.Insert()
}


func build() {
	sshConfig := &commons.SSHConfig{"root", "1qaz@WSX", "192.168.104.141", 22}

	var h5_ticket_com  job.DeployJob;
	h5_ticket_com.Config = sshConfig
	h5_ticket_com.Name = "h5_ticket_com"
	h5_ticket_com.RemoteBefore = append(h5_ticket_com.RemoteBefore,&job.RemoteComm{false,"rm -rf /export/App/h5.ticket.jd.com"})
	h5_ticket_com.LocalBefore = append(h5_ticket_com.LocalBefore,&job.LocalComm{false,"E:\\github\\ticket.h5\\web","mvn",
		[]string{"clean","package","-Dmaven.test.skip=true","-P","artifactory,development","-Dfile.encoding=UTF-8"}})
	h5_ticket_com.UploadJob = &job.UploadJob{"E:/github/ticket.h5/web/target/ticket-h5-web","/home"}
	h5_ticket_com.RemoteAfter = append(h5_ticket_com.RemoteAfter , &job.RemoteComm{false,"mv /home/ticket-h5-web /export/App/h5.ticket.jd.com/"})
	h5_ticket_com.RemoteAfter = append(h5_ticket_com.RemoteAfter , &job.RemoteComm{true,"sh /export/Shell/h5.ticket.jd.com/restart"})
	h5_ticket_com.Deploy()

	var piao_jd_web  job.DeployJob;
	piao_jd_web.Config = sshConfig
	piao_jd_web.Name = "piao_jd_web"
	piao_jd_web.RemoteBefore = append(piao_jd_web.RemoteBefore,&job.RemoteComm{false,"rm -rf /export/App/piao.jd.com"})
	piao_jd_web.LocalBefore = append(piao_jd_web.LocalBefore,&job.LocalComm{false,"E:\\git_source\\piao-web\\jd-ticket-web","mvn",
		[]string{"clean","package","-Dmaven.test.skip=true","-P","artifactory,development"}})
	piao_jd_web.UploadJob = &job.UploadJob{"E:/git_source/piao-web/jd-ticket-web/target/jd-ticket-web","/home"}
	piao_jd_web.RemoteAfter = append(piao_jd_web.RemoteAfter , &job.RemoteComm{false,"mv /home/jd-ticket-web  /export/App/piao.jd.com/"})
	piao_jd_web.RemoteAfter = append(piao_jd_web.RemoteAfter , &job.RemoteComm{true,"sh /export/Shell/piao.jd.com/tomcat"})
	//piao_jd_web.deploy()

	/*jobs := make([]job.DeployJob,0)
	jobs = append(jobs,h5_ticket_com)
	jobs = append(jobs,piao_jd_web)
	job.Save(jobs)
	jobs = job.Read("data")
	for _,v := range jobs {
		log.Println(v.Name)
	}*/

}

func a(){
	web.Service()
}