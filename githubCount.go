package main

import (
	"flag"
	"fmt"
	"github.com/jc3wish/githubCount/config"
	"github.com/jc3wish/githubCount/manager"
	"github.com/jc3wish/githubCount/server"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

var Daemon bool
var Pid *string
var DataDir *string
var Port *int
var Version bool
var Help bool

func usage() {
	fmt.Fprintf(os.Stderr, `Bifrost version: `+config.VERSION+`
Usage: Bifrost [-hv] [-config ./etc/Bifrost.ini] [-pid Bifrost.pid] [-data_dir dir]

Options:
`)
	flag.PrintDefaults()
}

func ConfigInit()  {
	flag.BoolVar(&Version, "v", false, "this version")
	flag.BoolVar(&Daemon, "d", false, "Daemon")
	flag.BoolVar(&Help, "h", false, "this help")
	Pid = flag.String("pid", "", "pid file path")
	DataDir = flag.String("data_dir", "./data", "data dir")
	Port = flag.Int("port", 11036, "data dir")
	flag.Usage = usage
	flag.Parse()

	if Help{
		flag.Usage()
		os.Exit(0)
	}
	if Version {
		fmt.Println(config.VERSION)
		os.Exit(0)
	}
}

func initLog(){
	logsDir  := config.DataDir+"/logs"
	os.MkdirAll(logsDir,0700)
	t := time.Now().Format("2006-01-02")
	LogFileName := logsDir+"/"+t+".log"
	f, err := os.OpenFile(LogFileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0700) //打开文件
	if err != nil{
		log.Println("log init error:",err)
	}
	log.SetOutput(f)
	fmt.Println("log input to",LogFileName)
}

func initParam(){
	os.MkdirAll(config.DataDir,0700)
	config.DataDir = *DataDir
	if *Port > 0{
		config.Listen = "0.0.0.0:"+fmt.Sprint(*Port)
	}
}

func WritePid(){
	if *Pid == ""{
		return
	}
	f, err2 := os.OpenFile(*Pid, os.O_CREATE|os.O_RDWR, 0700) //打开文件
	if err2 !=nil{
		log.Println("Open Pid Error; File:",*Pid,"; Error:",err2)
		os.Exit(1)
		return
	}
	defer f.Close()
	io.WriteString(f, fmt.Sprint(os.Getpid()))
}

func main()  {
	initParam()
	if Daemon == true && runtime.GOOS != "windows"{
		if os.Getppid() != 1{
			filePath,_:=filepath.Abs(os.Args[0])  //将命令行参数中执行文件路径转换成可用路径
			args:=append([]string{filePath},os.Args[1:]...)
			os.StartProcess(filePath,args,&os.ProcAttr{Files:[]*os.File{os.Stdin,os.Stdout,os.Stderr}})
			return
		}
		initLog()
	}
	go manager.Start(config.Listen)
	server.DoInit()
	WritePid()
	for{
		time.Sleep(time.Duration(1) * time.Hour)
	}
}

