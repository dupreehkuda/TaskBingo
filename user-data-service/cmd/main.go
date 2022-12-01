package main

import "github.com/dupreehkuda/TaskBingo/internal/api"

func main() {
	srv := api.NewByConfig()
	srv.Run()
}
