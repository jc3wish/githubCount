package manager

import (
	"encoding/json"
	"github.com/jc3wish/githubCount/server"
	"html/template"
	"log"
	"net/http"
)

func init()  {
	addRoute("/",index_controller)
	addRoute("/flow/get",get_flow_controller)
	addRoute("/project/add",add_project_controller)
}

func index_controller(w http.ResponseWriter,req *http.Request){
	type indexDetail struct {
		TemplateHeader
		ProjectList []server.ProjectInfo
	}
	req.ParseForm()
	var Result indexDetail
	Result = indexDetail{ProjectList:server.GetProjectList()}
	Result.Title = " home "
	t, err := template.ParseFiles(TemplatePath("manager/template/flow.html"),TemplatePath("manager/template/header.html"),TemplatePath("manager/template/footer.html"))
	if err != nil{
		log.Fatal(err)
	}
	t.Execute(w, Result)
}

func get_flow_controller(w http.ResponseWriter,req *http.Request)  {
	req.ParseForm()
	projectId := GetFormInt(req,"project_id")
	data := server.GetProjectStarList(int64(projectId))
	body,_:=json.Marshal(data)
	w.Write(body)
}


func add_project_controller(w http.ResponseWriter,req *http.Request)  {
	req.ParseForm()
	Name := req.Form.Get("project_name")
	if Name == ""{
		w.Write(returnDataResult(false,"project_name not be empty",""))
		return
	}
	_,err := server.GetGithubStarCountByName(Name)
	if err != nil{
		w.Write(returnDataResult(false,err.Error(),""))
		return
	}
	projectInfo := server.ProjectInfo{
		Id:   0,
		Name: Name,
	}
	server.AddProject(projectInfo)
	w.Write(returnDataResult(true,"success",""))
	return
}