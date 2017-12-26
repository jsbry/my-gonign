package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

func createMyRender() multitemplate.Render {
	r := multitemplate.New()
	r.AddFromFiles("Index", "./templates/Base.html", "./templates/Index.html")
	r.AddFromFiles("Project", "./templates/Base.html", "./templates/Project.html")
	r.AddFromFiles("Db", "./templates/Base.html", "./templates/Db.html")
	r.AddFromFiles("Error", "./templates/Base.html", "./templates/Error.html")
	r.AddFromFiles("NoRoute", "./templates/Base.html", "./templates/NoRoute.html")
	return r
}

func main() {
	router := gin.Default()

	router.Static("/css", "./css")
	router.Static("/js", "./js")
	router.Static("/fonts", "./fonts")
	router.StaticFile("/favicon.ico", "./favicon.ico")

	router.HTMLRender = createMyRender()

	router.GET("/", Index)
	router.GET("/project", Project)
	router.GET("/projectlist", Projectlist)
	router.GET("/db/:id", Db)
	router.GET("/dbinfo/:id", Dbinfo)

	router.NoRoute(NoRoute)
	router.Run(":8080")
}

func Index(c *gin.Context) {
	c.HTML(200, "Index", gin.H{
		"title": "ダッシュボード",
	})
}

func NoRoute(c *gin.Context) {
	c.HTML(404, "NoRoute", gin.H{
		"title": "Not Found",
	})
}

/**
 * データベース接続
 */
type DbConfig struct {
	Dsn      string `json:"dsn"`
	Username string `json:"username"`
	Password string `json:"password"`
	Server   string `json:"server"`
	Database string `json:"database"`
	Charset  string `json:"charset"`
}

var conf DbConfig
var db *sql.DB

func InitDB() (int, error) {
	jsonString, err := ioutil.ReadFile("./config/database.json")
	if err != nil {
		return 500, err
	}
	err = json.Unmarshal(jsonString, &conf)
	if err != nil {
		return 500, err
	}

	connect := fmt.Sprintf(conf.Dsn, conf.Username, conf.Password, conf.Server, conf.Database, conf.Charset)
	db, err = sql.Open("mysql", connect)

	if err != nil {
		return 500, err
	}

	return 0, nil
}
