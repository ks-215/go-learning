package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
)

func hello() {
	fmt.Println("Hello Go!")
}

func packages() {
	fmt.Println(rand.Intn(10))
}

func export_names() {
	fmt.Println(math.Pi)
}

func add(x int, y int) int {
	return x + y
} // 2つ以上の引数が同じ場合は(x, y int)といった感じでまとめて書ける

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

var c, python, java bool

func variables() {
	var i int
	fmt.Println(i, c, python, java)
}

var i, j int = 1, 2

func with_init() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func basic_types() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}

const (
	Big   = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }

func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
