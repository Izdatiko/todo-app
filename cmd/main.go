package main

import "log"

func main() {
	srv := new(todo.Server)
	if err := srv.Run("8080"); err != nil {
		log.Fatalf("Error occuped while running http  server: %s", err.Error())
	}
}
