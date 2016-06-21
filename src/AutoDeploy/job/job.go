package job

import (
	"os"
	"log"
	"bufio"
	"io"
	"AutoDeploy/commons"
	"encoding/json"
	"os/exec"
	"bytes"
	"golang.org/x/crypto/ssh"
)

const DataFile  = "data"

/**
保存
 */
func save(job []*DeployJob)  {
	file,err := os.OpenFile(DataFile,os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if os.IsNotExist(err){
		file,err = os.Create(DataFile)
	}
	if err != nil {
		log.Println("save job err",err)
	}
	defer file.Close()
	for _,v := range job {
		if v.Config.Password != ""{
			file.Write(v.byte())
			file.Write([]byte("\n"))
		}

	}
}

/**
从文件中读取数据
 */
func Read() []*DeployJob  {
	deploy := make([]*DeployJob,0)
	file,err := os.Open(DataFile)
	if os.IsNotExist(err){
		return deploy
	}
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
		deploy = append(deploy,byteToDeploy([]byte(line)))
	}
	return deploy
}

func SaveOrUpdate(job *DeployJob) []*DeployJob {
	jobs := Read()
	exit := false
	for i,j := range jobs {
		if j.Name == job.Name{
			jobs[i] = job
			exit = true
			break
		}
	}
	if !exit{
		jobs = append(jobs,job)
	}
	save(jobs)
	return jobs
}

func Delete(j DeployJob) []*DeployJob {
	jobs := Read()
	for i,v := range jobs {
		if v.Name == j.Name{
			a := i + 1
			jobs = append(jobs[0:i] ,jobs[a:]...)
		}
	}
	save(jobs)
	return jobs
}

type DeployJob  struct{
	Name string
	Config *commons.SSHConfig
	LocalBefore []*LocalComm
	RemoteBefore []*RemoteComm
	UploadJob *UploadJob
	LocalAfter []*LocalComm
	RemoteAfter []*RemoteComm
}

func (d *DeployJob) byte() []byte{
	data := make(map[string][]byte)
	data["0"] = []byte(d.Name)
	b ,_ := json.Marshal(d.Config)
	data["cf"] = b
	b ,_ = json.Marshal(d.LocalBefore)
	data["lb"] = b
	b ,_ = json.Marshal(d.RemoteBefore)
	data["rb"] = b
	b ,_ = json.Marshal(d.UploadJob)
	data["up"] = b
	b ,_ = json.Marshal(d.LocalAfter)
	data["la"] = b
	b ,_ = json.Marshal(d.RemoteAfter)
	data["ra"] = b
	s,_ := json.Marshal(data)
	return s;
}

func byteToDeploy(b []byte) *DeployJob   {
	var deployJob DeployJob
	data := make(map[string][]byte)
	_ = json.Unmarshal(b,&data)

	deployJob.Name = string(data["0"])

	d := data["cf"]
	var config commons.SSHConfig
	_ = json.Unmarshal(d,&config)
	deployJob.Config = &config

	d = data["lb"]
	var lb []*LocalComm
	_ = json.Unmarshal(d,&lb)
	deployJob.LocalBefore = lb

	d = data["rb"]
	var rb []*RemoteComm
	_ = json.Unmarshal(d,&rb)
	deployJob.RemoteBefore = rb

	d = data["up"]
	var up UploadJob
	_ = json.Unmarshal(d,&up)
	deployJob.UploadJob = &up

	d = data["la"]
	var la []*LocalComm
	_ = json.Unmarshal(d,&la)
	deployJob.LocalAfter = la

	d = data["ra"]
	var ra []*RemoteComm
	_ = json.Unmarshal(d,&ra)
	deployJob.RemoteAfter = ra

	return &deployJob
}

func (deploy *DeployJob) Deploy()  {
	client, err := commons.GetSSHClient(deploy.Config)//*ssh.Client
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	for _,lb := range deploy.LocalBefore {
		lb.run()
	}
	for _,rb := range deploy.RemoteBefore {
		rb.run(client)
	}
	uploadJob := deploy.UploadJob
	if uploadJob != nil && deploy.UploadJob.LocalPath != ""{
		deploy.UploadJob.run(client)
	}
	for _,la := range deploy.LocalAfter{
		la.run()
	}
	for _,ra := range deploy.RemoteAfter {
		ra.run(client)
	}
}

type LocalComm  struct{
	IsGo bool
	Path string
	Command string
	Args []string

}

func (job *LocalComm) run()  {
	if job.IsGo{
		go cmd(job.Path,job.Command,job.Args)
	}else {
		cmd(job.Path,job.Command,job.Args)
	}
}

type RemoteComm  struct{
	IsGo bool
	Command string
}

func (r *RemoteComm) run(client *ssh.Client)  {
	if r.IsGo{
		commons.ExecuteShellGo(client,r.Command)
	}else {
		commons.ExecuteShell(client,r.Command)
	}
}

type UploadJob  struct{
	LocalPath string
	RemotePath string //
}

func (u *UploadJob) run(client *ssh.Client)  {
	if u.RemotePath == ""{
		u.RemotePath = "/tmp"
	}
	commons.UploadPath(client,u.LocalPath,u.RemotePath)
}

/**
本地命令
 */
func cmd(path, command string, arg []string) {
	if path=="" || command==""{
		return
	}
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
