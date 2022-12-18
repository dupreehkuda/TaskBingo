package main

import "github.com/dupreehkuda/TaskBingo/task-data-service/internal/server"

func main() {
	srv := server.NewByConfig()
	srv.Run()
}
