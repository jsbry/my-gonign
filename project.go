package main

import "github.com/gin-gonic/gin"

func Project(c *gin.Context) {
	c.HTML(200, "Project", gin.H{
		"title": "プロジェクト",
		"vuejs": "project.js",
	})
	return
}

func Projectlist(c *gin.Context) {
	var project_id, project_name, created, updated string
	var list []gin.H
	InitDB()
	defer db.Close()

	err := db.Ping()
	if err != nil {
		c.JSON(200, gin.H{
			"title":   "エラー",
			"message": "データベースエラーが発生しました",
		})
		return
	}

	rows, err := db.Query("SELECT project_id, project_name, DATE_FORMAT(created, '%Y年%m月%d日 %H:%i:%s') AS created, DATE_FORMAT(updated, '%Y年%m月%d日 %H:%i:%s') AS updated FROM projects WHERE is_deleted = 0")
	if err != nil {
		c.JSON(200, gin.H{
			"code": 500,
			"error": gin.H{
				"message": "データベースエラーが発生しました",
			},
		})
		return
	}
	defer rows.Close()

	for i := 0; rows.Next(); i++ {
		if err := rows.Scan(&project_id, &project_name, &created, &updated); err != nil {
			c.JSON(200, gin.H{
				"code": 500,
				"error": gin.H{
					"message": "データベースエラーが発生しました",
				},
			})
			return
		}
		data := gin.H{
			"project_id":   project_id,
			"project_name": project_name,
			"created":      created,
			"updated":      updated,
		}
		list = append(list, data)
	}

	c.JSON(200, gin.H{
		"code": 200,
		"result": gin.H{
			"projectlist": list,
		},
	})
	return
}
