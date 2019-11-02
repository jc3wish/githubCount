package manager

import (
	"encoding/json"
	"github.com/jc3wish/githubCount/manager/xgo"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime/debug"
	"strconv"
	"strings"
)


var execDir string

func init(){
	execPath, _ := exec.LookPath(os.Args[0])
	execDir = filepath.Dir(execPath)+"/"
}

func TemplatePath(fileName string) string{
	return execDir+fileName
}

type TemplateHeader struct {
	Title string
}

func (TemplateHeader *TemplateHeader) setTile(title string){
	TemplateHeader.Title = title
}

type resultStruct struct {
	Status bool `json:"status"`
	Msg string `json:"msg"`
}

type resultDataStruct struct {
	Status bool `json:"status"`
	Msg string `json:"msg"`
	Data interface{} `json:"data"`
}

func returnResult(r bool,msg string)[]byte{
	b,_:=json.Marshal(resultStruct{Status:r,Msg:msg})
	return  b
}

func returnDataResult(r bool,msg string,data interface{})[]byte{
	b,_:=json.Marshal(resultDataStruct{Status:r,Msg:msg,Data:data})
	return  b
}

func GetFormInt(req *http.Request,key string) int{
	v := strings.Trim(req.Form.Get(key),"")
	intv,err:=strconv.Atoi(v)
	if err != nil{
		return 0
	}else{
		return intv
	}
}

func addRoute(route string, callbackFUns func(http.ResponseWriter,*http.Request) ){
	xgo.AddRoute(route,callbackFUns)
}

var writeRequestOp = []string{"/add","/del","/start","/stop","/close","/deal","/update","/export","/import"}
//判断是否为写操作
func checkWriteRequest(uri string) bool {
	for _,v := range writeRequestOp{
		if strings.Contains(uri,v){
			return true
		}
	}
	return false
}

func controller_FirstCallback(w http.ResponseWriter,req *http.Request) bool {
	return true
}

func Start(IpAndPort string){
	defer func() {
		if err:=recover(); err!= nil{
			debug.PrintStack()
		}
	}()
	xgo.AddStaticRoute("/css/",TemplatePath("manager/public/"))
	xgo.AddStaticRoute("/js/",TemplatePath("manager/public/"))
	xgo.AddStaticRoute("/fonts/",TemplatePath("manager/public/"))
	xgo.AddStaticRoute("/img/",TemplatePath("manager/public/"))
	xgo.AddStaticRoute("/plugin/",TemplatePath("/"))
	xgo.SetFirstCallBack(controller_FirstCallback)
	var err error
	err = xgo.Start(IpAndPort)
	if err != nil{
		log.Println("Manager Start Err:",err)
	}else{
		log.Println("http server :",IpAndPort," success")
	}
}