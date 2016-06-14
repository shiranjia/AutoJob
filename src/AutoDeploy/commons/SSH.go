package commons

import (
	"bufio"
	"bytes"
	"fmt"
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
	// Once a Session is created, you can execute a single command on
	// the remote side using the Run method.
	var b bytes.Buffer
	session.Stdout = &b

	if err := session.Run(shell); err != nil {
		panic("Failed to run: " + err.Error())
	}
	fmt.Println(b.String())
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
		log.Fatal("estart shell err:", err)
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
		log.Fatal(err)
	}
	return sftp
}

/**
上传文件
*/
func UploadFile(sftp *sftp.Client, localFile, remotePath string) {
	log.Println(localFile, ",", remotePath)
	// leave your mark
	inputFile, inputError := os.Open(localFile)
	//fileInfo , err := inputFile.Stat();
	defer inputFile.Close()

	f, err := sftp.Create(remotePath)

	if err != nil {
		log.Fatal("sftp.Create.err", err)
	}

	if inputError != nil {
		fmt.Fprintf(os.Stderr, "File Error: %s\n", inputError)
	}

	fileReader := bufio.NewReader(inputFile)
	counter := 0
	for {
		buf := make([]byte, 20480)
		n, err := fileReader.Read(buf)
		if err == io.EOF{
			break
		}
		counter++
		//fmt.Printf("%d,%s", n, string(buf))
		if n == 0 {
			break
		}
		//fmt.Println(string(buf))
		if _, err := f.Write(buf[0:n]); err != nil {
			log.Fatal(err)
		}

	}
	// check it's there
	fi, err := sftp.Lstat(remotePath)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(fi)

}

/**
删除文件
*/
func RemoveFile(remoateFile string, sftp *sftp.Client) {
	err := sftp.Remove(remoateFile)
	if err != nil {
		log.Println(err)
	}
}

/**
查看文件列表
*/
func ListPath(sftp *sftp.Client, remotePath string) {
	//defer sftp.Close()
	// walk a directory
	w := sftp.Walk(remotePath)
	for w.Step() {
		if w.Err() != nil {
			continue
		}
		log.Println(w.Path())
	}
}
