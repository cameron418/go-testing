package main

import "github.com/kataras/iris"

func main() {
    api := iris.New()
    api.Get("/", index)
    api.Get("/hi", hi)

    api.Listen(":8080")
}