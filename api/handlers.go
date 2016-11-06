package main

import "github.com/kataras/iris"

func index(ctx *iris.Context) {
	ctx.Write("Welcome to the API!\n")
}

func hi(ctx *iris.Context) {
	ctx.Write("Hi %s\n", "iris")
}
