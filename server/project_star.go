package server

import (
	"log"
	"time"
)

type ProjectStarInfo struct {
	Id int64
	Project_id int64
	Add_time int64
	Subscribers_count int64
	Stargazers_count int64
	Forks_count	int64
}

func GetProjectStarList(projext_id int64) (data []ProjectStarInfo){
	dbConn := GetDbConn()
	defer BackDbConn(dbConn)
	sql :=`SELECT id,add_time,subscribers_count,stargazers_count,forks_count FROM x_project_star WHERE project_id=?`
	stmt,err := dbConn.Prepare(sql)
	if err != nil{
		log.Println(err)
		return
	}
	defer stmt.Close()
	rows,err := stmt.Query(projext_id)
	if err != nil{
		log.Println(err)
		return
	}
	defer rows.Close()
	var id int64
	var add_time int64
	var subscribers_count int64
	var stargazers_count int64
	var forks_count int64
	for rows.Next() {
		err = rows.Scan(&id, &add_time,&subscribers_count,&stargazers_count,&forks_count)
		data = append(data,ProjectStarInfo{
			Id:id,
			Add_time:add_time,
			Subscribers_count:subscribers_count,
			Stargazers_count:stargazers_count,
			Forks_count:forks_count,
		})
	}
	return data
}

func AddProjectStar(data ProjectStarInfo)  (int64,error) {
	sql := `INSERT INTO x_project_star (project_id,add_time,subscribers_count,stargazers_count,forks_count) VALUES (?,?,?,?,?)`
	dbConn := GetDbConn()
	defer BackDbConn(dbConn)
	stmt,err := dbConn.Prepare(sql)
	if err != nil{
		log.Println(err)
		return 0,err
	}
	defer stmt.Close()
	result,err := stmt.Exec(data.Project_id,time.Now().Unix(),data.Subscribers_count,data.Stargazers_count,data.Forks_count)
	return result.LastInsertId()
}