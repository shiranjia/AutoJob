package web

import (
	"testing"
	"AutoJob/job"
	"encoding/json"
	"log"
)

func TestService(t *testing.T) {
	s := `{"Name":"test","Config":{"User":"root","Password":"1qaz@WSX","Ip":"192.168.104.141","Port":0},"LocalBefore":[{"IsGo":false,"Path":"E:\\github\\","Command":"java","Param":"","Args":["clean"]}],"RemoteBefore":[{"IsGo":false,"Command":"unzip -d /home /home/samClub-man.war"}],"UploadJob":{"LocalPath":"","RemotePath":""},"LocalAfter":null,"RemoteAfter":null,"Show":false,"show":true}`
	t.Log(s)
	var _job *job.DeployJob = new(job.DeployJob)
	err := json.Unmarshal([]byte(s),_job)
	if err != nil{
		log.Println("err:",err)
	}
}
