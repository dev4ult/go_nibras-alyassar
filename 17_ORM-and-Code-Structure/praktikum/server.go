package main

import (
	config "praktikum/config"
	route "praktikum/routes"
)

func main() {
	config.InitDB()

    e := route.New()
	e.Logger.Fatal(e.Start(":8000"))
}