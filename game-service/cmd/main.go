package main

import "github.com/dupreehkuda/TaskBingo/game-service/internal/server"

func main() {
	srv := server.NewByConfig()
	srv.Run()
}
