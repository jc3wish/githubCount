package server_test

import (
	"github.com/jc3wish/githubCount/server"
	"testing"
)

func TestAddProject(t *testing.T) {
	server.DoInit()
	data := server.ProjectInfo{
		Id:   0,
		Name: "brokercap/Bifrost",
	}
	id,err:=server.AddProject(data)
	if err!=nil{
		t.Fatal(err)
	}
	t.Log("id:",id)
}

func TestGetProjectList(t *testing.T) {
	server.DoInit()
	data := server.GetProjectList()

	if len(data) == 0{
		t.Fatal("GetProjectList empty")
	}
	t.Log(data)
}