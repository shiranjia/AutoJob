package main

import (
	"AutoJob/job"
	"AutoJob/web"
	"os"
	"log"
	"path/filepath"
)

func main() {
	args := os.Args
	jobName := ""
	path := ""
	if len(args) != 0{
		for i ,arg := range args{
			if arg == "-j" {
			   jobName = args[i+1]
			}
			if arg == "-p" {
				path = args[i+1]
				break
			}
		}
	}
	if jobName != ""{
		jobs := job.ReadFrom(path + string(filepath.Separator) + "data")
		log.Println("run job ",jobName,"path:",path)
		for _ , j := range jobs {
			//log.Println(" job name : ",j.Name)
			if j.Name == jobName{
				j.Deploy()
				//log.Println(" job name : ",j.Name)
				break
			}
		}
	}else {
		web.Service()
	}

}



