package server

import (
	"log"
	"time"
)

type ProjectInfo struct {
	Id	int64
	Name string
}

func GetProjectList() (data []ProjectInfo) {
	dbConn := GetDbConn()
	defer BackDbConn(dbConn)
	sql :=`SELECT id,name,add_time FROM x_project`
	//data = append(data,ProjectInfo{Id:1,Name:"brokercap/Birost"})
	stmt,err := dbConn.Prepare(sql)
	if err != nil{
		log.Println(err)
		return
	}
	defer stmt.Close()
	rows,err := stmt.Query()
	if err != nil{
		log.Println(err)
		return
	}
	defer rows.Close()
	var id int64
	var name string
	var add_time int64
	for rows.Next() {
		err = rows.Scan(&id, &name, &add_time)
		data = append(data,ProjectInfo{Id:id,Name:name})
	}
	return data
}

func AddProject(data ProjectInfo) (int64,error) {
	sql := `INSERT INTO x_project (name,add_time) VALUES (?,?)`
	dbConn := GetDbConn()
	defer BackDbConn(dbConn)
	stmt,err := dbConn.Prepare(sql)
	if err != nil{
		log.Println(err)
		return 0,err
	}
	defer stmt.Close()
	result,err := stmt.Exec(data.Name,time.Now().Unix())
	if err != nil{
		log.Println(err)
		return 0,err
	}
	return result.LastInsertId()
}


