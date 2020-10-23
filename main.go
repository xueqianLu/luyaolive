package main

import (
	"database/sql"

	"github.com/labstack/echo"
	"github.com/xueqianLu/luyaolive/handler"
)

func main() {
	db := initDb("storage.db")
	migrate(db)

	e := echo.New()
	// 注册路由
	e.Static("/", "public")
	e.GET("/tasks", handler.GetTasks(db))
	e.POST("/task", handler.PostTask(db))
	e.PUT("/task", handler.PutTask(db))
	e.DELETE("/task/:id", handler.DeleteTask(db))
	//e.GET("/", handler.GetIndex)
	// 开启 HTTP Server
	e.Logger.Fatal(e.Start(":80"))
}

func initDb(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)

	if err != nil {
		panic(err)
	}
	if db == nil {
		panic("db nil")
	}

	return db
}

func migrate(db *sql.DB) {
	sql := `
		CREATE TABLE IF NOT EXISTS tasks (
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			name VARCHAR NOT NULL,
			done INTEGER NOT NULL
		);
	`

	_, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
}
