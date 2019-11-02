package server_test

import (
	"github.com/jc3wish/githubCount/server"
	"github.com/robfig/cron"
	"log"
	"testing"
	"time"
)

func TestGetGithubStarCount(t *testing.T)  {
	server.DoInit()
	server.GetGithubStarCount()
	t.Log("over")
}

func TestGetGithubStarCountByName(t *testing.T) {
	Name := "brokercap/Bifrost"

	data ,err := server.GetGithubStarCountByName(Name)
	if err !=nil{
		t.Fatal(err)
	}

	if data.Stargazers_count == 0{
		t.Fatal("Stargazers_count error",data)
	}

	t.Log(data)
}

var c *cron.Cron

func TestCrond(t *testing.T)  {
	go TestCrondStart(t)
	for{
		time.Sleep(time.Duration(100) * time.Second)
		break
	}
	t.Log("test success")
}


func TestCrondStart(t *testing.T)  {
	c = cron.New()
	_,err := c.AddFunc("*/1 * * * *", func() {
		log.Println("crond:",time.Now().Format("2006-01-02 15:04:05"))
	})
	if err != nil{
		t.Fatal(err)
	}

	_,err = c.AddFunc("@daily", func() {
		log.Println("crond:",time.Now().Format("2006-01-02 15:04:05"))
	})
	if err != nil{
		t.Fatal(err)
	}
	c.Start()

	/*
	defer c.Stop()

	// 这是一个使用time包实现的定时器，与cron做对比
	t1 := time.NewTimer(time.Second * 10)
	for {
		select {
		case <-t1.C:
			t1.Reset(time.Second * 10)
		}
	}
	*
	 */
}