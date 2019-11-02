package server

import (
	"encoding/json"
	"github.com/robfig/cron"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func StartCron()  {
	go doStartCron()
}

var crondOjb *cron.Cron

func doStartCron()  {
	crondOjb:= cron.New()
	crondOjb.AddFunc("@daily", GetGithubStarCount)
	crondOjb.Start()
}

type gitHubCount struct {
	Subscribers_count int64 `json:"subscribers_count"`
	Stargazers_count int64 `json:"stargazers_count"`
	Forks_count int64 `json:"forks_count"`
}

func GetGithubStarCount()  {
	NowTime := time.Now().Unix()
	for _,p := range GetProjectList(){
		d,err :=GetGithubStarCountByName(p.Name)
		if err!=nil{
			log.Println(err)
			continue
		}
		data := ProjectStarInfo{
			Project_id:p.Id,
			Add_time:NowTime,
			Subscribers_count:d.Subscribers_count,
			Stargazers_count:d.Stargazers_count,
			Forks_count:d.Forks_count,
		}
		AddProjectStar(data)
	}
}

func GetGithubStarCountByName(name string) (d gitHubCount,err error){
	url:="https://api.github.com/repos/"+name
	var body []byte
	body,err = GetHttpResult(url)
	if err!=nil{
		log.Println(err)
		return
	}
	err = json.Unmarshal(body,&d)
	if err!=nil{
		log.Println(err)
		return
	}
	return
}

func GetHttpResult(url string) (body []byte,err error) {
	request, _ := http.NewRequest("GET", url, nil)
	//fmt.Println(url)
	client := &http.Client{}
	response, _ := client.Do(request)
	body, err = ioutil.ReadAll(response.Body)
	return
}