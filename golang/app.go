package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	http.ListenAndServe(":80", router)
}
