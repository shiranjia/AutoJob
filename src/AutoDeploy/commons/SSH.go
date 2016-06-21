package commons

import (
	"bufio"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"log"
	"os"
	"strconv"
	"io"
	"time"
)

/**
SSH配置
*/
type SSHConfig struct {
	User     string
	Password string
	Ip       string
	Port     int
}

/**
创建ssh客户端
*/
func GetSSHClient(conf *SSHConfig) (*ssh.Client, error) {
	config := &ssh.ClientConfig{
		User: conf.User,
		Auth: []ssh.AuthMethod{
			ssh.Password(conf.Password),
		},
	}
	if 0 == conf.Port{
		conf.Port = 22
	}
	client, err := ssh.Dial("tcp", conf.Ip+":"+strconv.Itoa(conf.Port), config)
	if err != nil {
		//panic("Failed to dial: " + err.Error())
		return nil, err
	}
	return client, nil
}

/**
执行脚本
*/
func ExecuteShell(client *ssh.Client, shell string) {
	if shell == "" {
		log.Println("shell is nil")
		return
	}
	session, err := client.NewSession()
	if err != nil {
		panic("Failed to create session: " + err.Error())
	}
	defer session.Close()
	session.Setenv("LANG","zh_CN.UTF-8")
	// Once a Session is created, you can execute a single command on
	// the remote side using the Run method.
	//var b bytes.Buffer
	session.Stdout = os.Stdout

	if err := session.Run(shell); err != nil {
		//panic("Failed to run: " + err.Error() + "shell:" + shell)
		log.Println("Failed to run: " , err.Error() , "shell:" , shell)
	}
	//fmt.Println(b.String())
}

func ExecuteShellGo(client *ssh.Client, shell string){
	if shell == "" {
		log.Println("shell is nil")
		return
	}
	session, err := client.NewSession()
	if err != nil {
		panic("Failed to create session: " + err.Error())
	}
	defer session.Close()
	// Once a Session is created, you can execute a single command on
	// the remote side using the Run method.
	//var b bytes.Buffer
	//session.Stdout = &b

	out, err := session.StdoutPipe()
	if err != nil {
		log.Println("estart shell err:", err)
	}
	read := bufio.NewReader(out)
	session.Setenv("LANG","zh_CN.UTF-8")
	session.Start(shell)
	start := time.Now().Second()
	for {
		line, err := read.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}
		log.Print(line)
		if (time.Now().Second() - start) >= 10 {
			break
		}
	}

}

/**
创建sftp
*/
func GetSftp(client *ssh.Client) *sftp.Client {
	sftp, err := sftp.NewClient(client)
	if err != nil {
		log.Println("GetSftp.error",err)
	}
	return sftp
}


