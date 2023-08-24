package main

import "fmt"

func main() {
	LuasTrapesium(50, 75, 60)
}

func LuasTrapesium(alasAtas float64, alasBawah float64, tinggi float64) {
	var luas float64 = 0.5 * (alasAtas + alasBawah) * tinggi
	fmt.Println("Luas Trapesium :", luas)
}