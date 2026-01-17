package main

import (
	"fmt"
	"math"
	"runtime"
)

func sample1() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func sample2() {
	sum := 1
	for i := 1; sum < 1000; i++ {
		sum += sum
		fmt.Println(i, sum)
	}
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func sample3() {
	fmt.Println(sqrt(2), sqrt(-4))
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func sample4() {
	fmt.Println(pow(3, 2, 10))
	fmt.Println(pow(3, 3, 20))
}

func sample5() {
	for i := 0; i < 30; i++ {
		if i <= 10 || i >= 20 {
			fmt.Println(i)
			continue
		}

		if i%2 != 0 {
			fmt.Println(i)
			continue
		}

		if i/3 <= 5 {
			fmt.Println(i)
			continue
		}

		println(i, "★")
	}
}

func Sqrt(x float64) float64 { // ?????
	z := float64(1)
	for i := 0; i < 10; i++ {
		fmt.Println(z)
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func lesson1() { // ?????
	fmt.Println(Sqrt(2))
}

func sample6() {
	fmt.Println("Go runs on")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s.\n", os)
	}

	switch num := 3; num {
	case 1:
		fmt.Println("One")
	case 2:
		println("Two")
	case 3:
		fmt.Println("three")
		fallthrough
	case 4:
		println("Unknown number")
	}
}

func sample7() {
	defer fmt.Println("world")

	fmt.Println("hello")
}

func sample8() {
	fmt.Println("Counting")

	for i := 0; i <= 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("Done")
}

func main() {
	sample8()
}
