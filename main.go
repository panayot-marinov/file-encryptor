package main

//TODO:
//CODE TO DIFFERENT FILES!	(50%)
//Refactoring				(50%)
//DB:
//Original file name		(checked)
//Key 						(checked)
//Timestamp 				(checked)
//Unit Tests				(checked) //TODO!!!
//UUID GoLang 				(checked)
//Environment variables		(checked)
//new tasks
//run app in docker			(checked)
//connect db				//host docker internal?
//export db in docker (skip)
//migrate in kubernetes (may be locally, read about ingres contr)
//* create a shellscript which pushes the image to dockerhub
//read about non-relational dbs

//Shell script for deployment

import (
	"file-encryptor/sources"
	"fmt"

	_ "github.com/lib/pq" //postgres drivers for initialization
)

type Upload struct {
	id       int
	fileId   string
	filePath string
}

func getData(id int) Upload { //Test function
	db := sources.ConnectToDb()
	defer db.Close()

	fmt.Println("Sucessfully connected")

	var upload Upload
	row := db.QueryRow("SELECT id, file_id, file_path FROM \"Uploads\" where id=$1", id)

	if err := row.Scan(&upload.id, &upload.fileId, &upload.filePath); err != nil {
		fmt.Println("ERROR! Cannot execute query!")
		fmt.Println(err)
	}

	//rows, err := db.Query(query, qteNum, limit)
	//defer rows.Close()
	// if err != nil {
	// 	log.Println(err)
	// }
	return upload
}

func main() {
	//os.Setenv("CONNSTR", "user=postgres password=parolazabaza host=127.0.0.1 port=5432 dbname=MainDB connect_timeout=20 sslmode=disable")

	// album := getData(1)
	// fmt.Println("ID: " + strconv.Itoa(album.id))
	// fmt.Println("fileID: " + album.fileId)
	// fmt.Println("filePath: " + album.filePath)

	sources.SetupRoutes()
}
