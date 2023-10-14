package main

import (
	router "EdmundsBankai/golang-intro/gin-tutorial/router"
)

func main() {
	r := router.SetupRouter()
	r.Run("localhost:8080")
}
