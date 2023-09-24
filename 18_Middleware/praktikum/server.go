package main

import (
	route "praktikum/routes"
)

func main() {
    // config.ConnectDB()

    e := route.New()

    e.Logger.Fatal(e.Start(":8000"))
}
