package main

import (
	"bytes"
	"fmt"
	_ "io"
	_ "io/ioutil"
	"log"
	"os"
	"os/exec"
	"AutoDeploy/commons"
	"golang.org/x/crypto/ssh"
	_ "html/template"
	"encoding/json"
	"bufio"
	"io"
)


func main() {
	fmt.Println("gopath=", os.Getenv("gopath"))
	//demo.Route()
	//web.Service()
	build()
	//a := "org.apache.coyote.AbstractProtocol.start Starting ProtocolHandler [\"ajp-nio-8009\"]"
	//fmt.Println(strings.Contains(a,"startup in"))
	//db.Insert()
}

func build() {
	sshConfig := &commons.SSHConfig{"root", "1qaz@WSX", "192.168.104.141", 22}

	var h5_ticket_com  DeployJob;
	h5_ticket_com.config = sshConfig
	h5_ticket_com.name = "h5_ticket_com"
	h5_ticket_com.remoteBefore = append(h5_ticket_com.remoteBefore,&remoteComm{false,"rm -rf /export/App/h5.ticket.jd.com"})
	h5_ticket_com.localBefore = append(h5_ticket_com.localBefore,&localComm{false,"E:\\github\\ticket.h5\\web","mvn",
		[]string{"clean","package","-Dmaven.test.skip=true","-P","artifactory,development","-Dfile.encoding=UTF-8"}})
	h5_ticket_com.uploadJob = &uploadJob{"E:/github/ticket.h5/web/target/ticket-h5-web","/home"}
	h5_ticket_com.remoteAfter = append(h5_ticket_com.remoteAfter , &remoteComm{false,"mv /home/ticket-h5-web /export/App/h5.ticket.jd.com/"})
	h5_ticket_com.remoteAfter = append(h5_ticket_com.remoteAfter , &remoteComm{true,"sh /export/Shell/h5.ticket.jd.com/restart"})
	//h5_ticket_com.deploy()

	var piao_jd_web  DeployJob;
	piao_jd_web.config = sshConfig
	piao_jd_web.name = "piao_jd_web"
	piao_jd_web.remoteBefore = append(piao_jd_web.remoteBefore,&remoteComm{false,"rm -rf /export/App/piao.jd.com"})
	piao_jd_web.localBefore = append(piao_jd_web.localBefore,&localComm{false,"E:\\git_source\\piao-web\\jd-ticket-web","mvn",
		[]string{"clean","package","-Dmaven.test.skip=true","-P","artifactory,development"}})
	piao_jd_web.uploadJob = &uploadJob{"E:/git_source/piao-web/jd-ticket-web/target/jd-ticket-web","/home"}
	piao_jd_web.remoteAfter = append(piao_jd_web.remoteAfter , &remoteComm{false,"mv /home/jd-ticket-web  /export/App/piao.jd.com/"})
	piao_jd_web.remoteAfter = append(piao_jd_web.remoteAfter , &remoteComm{true,"sh /export/Shell/piao.jd.com/tomcat"})
	//piao_jd_web.deploy()

	job := make([]DeployJob,0)
	job = append(job,h5_ticket_com)
	job = append(job,piao_jd_web)
	save(job)
	jobs := read("data")
	for _,v := range jobs {
		log.Println(v)
	}

}
func checkErr(e error)  {
	if e != nil{
		panic(e)
	}
}

func save(job []DeployJob)  {
	file,err := os.Open("data")
	if os.IsNotExist(err){
		file,err = os.Create("data")
	}
	if err != nil {
		log.Println("save job err",err)
	}
	defer file.Close()
	for _,v := range job {
		if v.config.Password != ""{
			file.Write(v.byte())
			file.Write([]byte("\n"))
		}

	}
}

func read(datafile string) []DeployJob  {
	deploy := make([]DeployJob,0)
	file,err := os.Open(datafile)
	if err != nil {
		log.Println("save job err",err)
	}
	defer file.Close()
	read := bufio.NewReader(file)
	for{
		line,err := read.ReadString('\n')
		if err!=nil {
			if err == io.EOF{
				break
			}
		}
		//fmt.Printf(line)
		deploy = append(deploy,byteToDeploy([]byte(line)))
	}
	return deploy
}

type DeployJob  struct{
	name string
	config *commons.SSHConfig
	localBefore []*localComm
	remoteBefore []*remoteComm
	uploadJob *uploadJob
	localAfter []*localComm
	remoteAfter []*remoteComm
}

func (d *DeployJob) byte() []byte{
	data := make(map[string][]byte)
	b ,_ := json.Marshal(d.config)
	data["cf"] = b
	b ,_ = json.Marshal(d.localBefore)
	data["lb"] = b
	b ,_ = json.Marshal(d.remoteBefore)
	data["rb"] = b
	b ,_ = json.Marshal(d.uploadJob)
	data["up"] = b
	b ,_ = json.Marshal(d.localAfter)
	data["la"] = b
	b ,_ = json.Marshal(d.remoteAfter)
	data["ra"] = b
	s,_ := json.Marshal(data)
	return s;
}

func byteToDeploy(b []byte) DeployJob   {
	var deployJob DeployJob
	data := make(map[string][]byte)
	_ = json.Unmarshal(b,&data)

	d := data["cf"]
	var config commons.SSHConfig
	_ = json.Unmarshal(d,&config)
	deployJob.config = &config

	d = data["lb"]
	var lb []*localComm
	_ = json.Unmarshal(d,&lb)
	deployJob.localBefore = lb

	d = data["rb"]
	var rb []*remoteComm
	_ = json.Unmarshal(d,&rb)
	deployJob.remoteBefore = rb

	d = data["up"]
	var up uploadJob
	_ = json.Unmarshal(d,&up)
	deployJob.uploadJob = &up

	d = data["la"]
	var la []*localComm
	_ = json.Unmarshal(d,&la)
	deployJob.localAfter = la

	d = data["ra"]
	var ra []*remoteComm
	_ = json.Unmarshal(d,&ra)
	deployJob.remoteAfter = ra

	return deployJob
}

func (deploy *DeployJob) deploy()  {
	client, err := commons.GetSSHClient(deploy.config)//*ssh.Client
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	for _,lb := range deploy.localBefore {
		lb.run()
	}
	for _,rb := range deploy.remoteBefore {
		rb.run(client)
	}
	uploadJob := deploy.uploadJob
	if uploadJob != nil && deploy.uploadJob.LocalPath != ""{
		deploy.uploadJob.run(client)
	}
	for _,la := range deploy.localAfter{
		la.run()
	}
	for _,ra := range deploy.remoteAfter {
		ra.run(client)
	}
}

type localComm  struct{
	IsGo bool
	Path string
	Command string
	Args []string

}

func (job *localComm) run()  {
	if job.IsGo{
		go cmd(job.Path,job.Command,job.Args)
	}else {
		cmd(job.Path,job.Command,job.Args)
	}
}

type remoteComm  struct{
	IsGo bool
	Command string
}

func (r *remoteComm) run(client *ssh.Client)  {
	if r.IsGo{
		commons.ExecuteShellGo(client,r.Command)
	}else {
		commons.ExecuteShell(client,r.Command)
	}
}

type uploadJob  struct{
	LocalPath string
	RemotePath string //
}

func (u *uploadJob) run(client *ssh.Client)  {
	if u.RemotePath == ""{
		u.RemotePath = "/tmp"
	}
	commons.UploadPath(client,u.LocalPath,u.RemotePath)
}

/**
本地命令
 */
func cmd(path, command string, arg []string) {
	cmd := exec.Command(command)
	cmd.Args = append(cmd.Args, arg...)
	cmd.Dir = path
	log.Println("cmd.path=", cmd.Dir)
	log.Println("cmd.command=", cmd.Args)
	var b bytes.Buffer
	cmd.Stdout = &b
	cmd.Stdout = os.Stdout
	cmd.Run()
	//log.Println(b.String())
}