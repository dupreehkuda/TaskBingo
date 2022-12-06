package main

import "github.com/dupreehkuda/TaskBingo/user-data-service/internal/server"

func main() {
	srv := server.NewByConfig()
	srv.Run()
}
