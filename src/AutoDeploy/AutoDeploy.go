package main

import (
	"bytes"
	"errors"
	"fmt"
	_ "io"
	_ "io/ioutil"
	"log"
	"os"
	"os/exec"
	"AutoDeploy/commons"

	"AutoDeploy/db"
)

const A = "a"
const B = "asd"
const (
	i = iota
	j
	h
)

func main() {
	fmt.Println("gopath=", os.Getenv("gopath"))
	//demo.Route()
	//web.Service()
	//build()
	db.Insert()
}

func build() {
	config := &commons.SSHConfig{"root", "1qaz@WSX", "192.168.104.141", 22}
	client, err := commons.GetSSHClient(config)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	arg := []string{"clean","package","-Dmaven.test.skip=true","-P","artifactory,development","-Dfile.encoding=UTF-8"}
	cmd("E:\\github\\ticket.h5\\web","mvn",arg)
	/*go commons.ExecuteShell(client,"rm -rf /export/App/h5.ticket.jd.com*/ /*");



	commons.UploadPath(client,"E:/github/ticket.h5/web/target/ticket-h5-web","/home")
	commons.ExecuteShell(client,"mv /home/ticket-h5-web*/ /* /export/App/h5.ticket.jd.com/");*/
	//commons.ExecuteShellGo(client, "sh /export/Shell/h5.ticket.jd.com/restart")
}

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

func equals(a string) (bool, error) {
	if a == "a" {
		return true, nil
	} else {
		return false, errors.New("not equals")
	}
}

func testSlice() {
	fmt.Println(B)
	fmt.Print(i)
	fmt.Print(j)
	fmt.Print(h)

	arr := [20]int{1, 2}
	reset(arr)
	fmt.Println(arr)
	a := []int{1, 2, 3}
	fmt.Println(a)
	b := make([]int, 3)
	//reset(b)
	b[1] = 12
	fmt.Println(b)
	b = arr[1:3]
	_ = append(b, 25)
	b[0] = 34
	fmt.Println(arr)
	copy(b, a)
	fmt.Println(a)
}

func testMap() {
	m := make(map[string]string)
	m["a"] = "b"
	m["b"] = "c"
	for k, v := range m {
		fmt.Println(k + " = " + v)
	}
	delete(m, "a")
	fmt.Println(m)
	a := make(map[int]int)
	a[12] = 21
	_ = a
	fmt.Println(a)
}

func reset(sli [20]int) {
	sli[0] = 12
	for v := range sli {
		defer fmt.Println("defer..", v)
	}
}
