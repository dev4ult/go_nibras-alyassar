package main

func EquationXYZ(a, b, c int) {
	for x := 1; x <= 1000; x++ {
		for y := 1; y <= 1000; y++ {
			for z := 1; z <= 1000; z++ {
				if x+y+z == a && x*y*z == b {
					fmt.Println(x, y, z)
					return
				}
			}
		}
	}
	fmt.Println("no solution")
}

func main() {

	EquationXYZ(1, 2, 3) // no solution

	EquationXYZ(6, 6, 14) // 1 2 3

}
