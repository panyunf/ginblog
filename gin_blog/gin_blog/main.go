package main

import (
	"gin_blog/model"
	"gin_blog/routes"
)

func main() {
	model.InitDb()
	routes.InitRouter()
}
