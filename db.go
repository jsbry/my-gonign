package main

import "github.com/gin-gonic/gin"

func Db(c *gin.Context) {
	id := c.Param("id")
	var project_id, project_name string

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

	if err := db.QueryRow("SELECT project_id, project_name FROM projects WHERE project_id = ? AND is_deleted = 0 LIMIT 1", id).Scan(&project_id, &project_name); err != nil {
		c.HTML(200, "Error", gin.H{
			"title":   "エラー",
			"message": "データベースエラーが発生しました",
		})
		return
	}

	c.HTML(200, "Db", gin.H{
		"title":        "DB設計書",
		"project_id":   project_id,
		"project_name": project_name,
		"vuejs":        "db.js",
	})
	return
}

func Dbinfo(c *gin.Context) {
	id := c.Param("id")
	var updated string
	var db_name, db_engine, db_charset string

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

	if err := db.QueryRow("SELECT DATE_FORMAT(updated, '%Y年%m月%d日 %H:%i:%s') AS updated FROM projects WHERE project_id = ? AND is_deleted = 0 LIMIT 1", id).Scan(&updated); err != nil {
		c.HTML(200, "Error", gin.H{
			"title":   "エラー",
			"message": "データベースエラーが発生しました",
		})
		return
	}

	if err := db.QueryRow("SELECT db_name, db_engine, db_charset FROM dbs WHERE project_id = ? AND is_deleted = 0 LIMIT 1", id).Scan(&db_name, &db_engine, &db_charset); err != nil {
		c.HTML(200, "Error", gin.H{
			"title":   "エラー",
			"message": "データベースエラーが発生しました",
		})
		return
	}

	c.JSON(200, gin.H{
		"code": 200,
		"result": gin.H{
			"updated":    updated,
			"db_name":    db_name,
			"db_engine":  db_engine,
			"db_charset": db_charset,
		},
	})
	return
}
