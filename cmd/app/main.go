package main

import "memo/internal/server"

func main() {
	s := server.New()
	s.Init()
	s.Run()
}
