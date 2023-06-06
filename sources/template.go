package sources

import (
	"embed"
	"html/template"
)

var tpl *template.Template

//go:embed templates/*
var templatesData embed.FS

func init() {
	//os.Setenv("CONNSTR", "user=postgres password=parolazabaza host=127.0.0.1 port=5432 dbname=MainDB connect_timeout=20 sslmode=disable")

	// fmt.Println("filepath = ", filepath.Join(".", "./sources/templates/styles", "mainStyle.css"))
	// _, err := os.Stat(filepath.Join(".", "./sources/templates/styles", "mainStyle.css"))
	// if err != nil {
	// 	fmt.Println("file not found")
	// } else {
	// 	fmt.Println("file found")
	// }

	// var err error
	// tpl, err = template.ParseFS(templatesData, "static/templates/*")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	tpl = template.Must(template.ParseFS(templatesData, "templates/*.html"))

	//TODO: See articles about golang embed static files
	//tpl = template.Must(template.ParseGlob("templates/*.html"))
}
