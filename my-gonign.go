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
	r.AddFromFiles("index", "./templates/base.html", "./templates/index.html")
	r.AddFromFiles("project", "./templates/base.html", "./templates/project.html")
	//r.AddFromFiles("article", "./templates/base.html", "templates/index.html", "templates/article.html")
	return r
}

func main() {
	router := gin.Default()

	router.Static("/css", "./css")
	router.Static("/js", "./js")
	router.Static("/fonts", "./fonts")

	router.HTMLRender = createMyRender()
	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index", gin.H{
			"title": "ダッシュボード",
		})
	})
	router.GET("/project", func(c *gin.Context) {
		c.HTML(200, "project", gin.H{
			"title": "プロジェクト",
		})
	})
	router.Run(":8080")
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
var db sql.DB

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
	db, err := sql.Open("mysql", connect)

	if err != nil {
		return 500, err
	}
	defer db.Close()

	return 0, nil
}
