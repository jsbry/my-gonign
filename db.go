package main

import "github.com/gin-gonic/gin"

// Db データベース設計画面
func Db(c *gin.Context) {
	id := c.Param("id")
	var projectID, projectName string

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

	if err := db.QueryRow("SELECT project_id, project_name FROM projects WHERE project_id = ? AND is_deleted = 0 LIMIT 1", id).Scan(&projectID, &projectName); err != nil {
		c.HTML(200, "Error", gin.H{
			"title":   "エラー",
			"message": "データベースエラーが発生しました",
		})
		return
	}

	c.HTML(200, "Db", gin.H{
		"title":        "DB設計書",
		"project_id":   projectID,
		"project_name": projectName,
		"vuejs":        "db.js",
	})
	return
}

// Dbinfo データベース詳細
func Dbinfo(c *gin.Context) {
	id := c.Param("id")
	var updated string
	var dbName, dbEngine, dbCharset string

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

	if err := db.QueryRow("SELECT db_name, db_engine, db_charset FROM dbs WHERE project_id = ? AND is_deleted = 0 LIMIT 1", id).Scan(&dbName, &dbEngine, &dbCharset); err != nil {
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
			"db_name":    dbName,
			"db_engine":  dbEngine,
			"db_charset": dbCharset,
		},
	})
	return
}
