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

func Project(c *gin.Context) {
	var project_id, project_name string
	var list []gin.H
	InitDB()
	defer db.Close()

	err := db.Ping()
	if err != nil {
		c.HTML(200, "Error", gin.H{
			"title":   "エラー",
			"message": "データベースエラーが発生しました",
		})
		return
	}

	rows, err := db.Query("SELECT project_id, project_name FROM projects WHERE is_deleted = 0")
	if err != nil {
		c.HTML(200, "Error", gin.H{
			"title":   "エラー",
			"message": "データベースエラーが発生しました",
		})
		return
	}
	defer rows.Close()

	for i := 0; rows.Next(); i++ {
		if err := rows.Scan(&project_id, &project_name); err != nil {
			c.HTML(200, "Error", gin.H{
				"title":   "エラー",
				"message": "データベースエラーが発生しました",
			})
			return
		}
		data := gin.H{
			"project_id":   project_id,
			"project_name": project_name,
		}
		list = append(list, data)
	}

	c.HTML(200, "Project", gin.H{
		"title": "プロジェクト",
		"list":  list,
	})
	return
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
