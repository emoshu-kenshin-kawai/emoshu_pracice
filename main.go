package main

import (
	"emoshu_practice/infrastructure"
	"fmt"

	_ "github.com/swaggo/echo-swagger/example/docs"
)

func main() {
	infrastructure.InitDB()
	fmt.Println("heldddlo")
}
