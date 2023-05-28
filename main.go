package main

import (
	"assignment2/routes"
	"assignment2/utils/postgres"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		HttpMainHandler()
		defer wg.Done()
	}()
	wg.Wait()
}

func HttpMainHandler() {
	g := gin.Default()
	db := postgres.NewConnection(postgres.BaseConfig()).Database

	routes.InitHtttpRoute(g, db)

	g.Run()
}
