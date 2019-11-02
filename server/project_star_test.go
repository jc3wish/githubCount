package server_test


import (
	"github.com/jc3wish/githubCount/server"
	"testing"
	"time"
)

func TestAddProjectStar(t *testing.T) {
	server.DoInit()
	data := server.ProjectStarInfo{
		Id:   0,
		Project_id:1,
		Add_time:time.Now().Unix(),
		Subscribers_count:13,
		Stargazers_count:154,
		Forks_count:47,
	}
	id,err:=server.AddProjectStar(data)
	if err!=nil{
		t.Fatal(err)
	}
	t.Log("id:",id)
}

func TestGetProjectStarList(t *testing.T) {
	server.DoInit()
	data := server.GetProjectStarList(1)
	if len(data) == 0{
		t.Fatal("GetProjectStarList empty")
	}
	t.Log(data)
}
