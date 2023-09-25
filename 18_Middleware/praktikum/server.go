package main

import (
	env "github.com/joho/godotenv"

	config "praktikum/config"
	route "praktikum/routes"
)

func main() {
    config.ConnectDB()
    env.Load(".env")

    e := route.New()

    e.Logger.Fatal(e.Start(":8000"))
}
