package main

import (
	"fmt"
	"math"
	"strings"

	"golang.org/x/tour/pic"
)

func sample1() {
	i, j := 42, 2701

	p := &i         // i へのポインタを p に入れる
	fmt.Println(*p) // i のポインタを通して値を表示（42が表示される）

	*p = 21        // i のポインタを通して値を更新（ i の値を直接更新）
	fmt.Println(i) //21が表示される

	p = &j       // j へのポインタを p に入れる
	*p = *p / 37 // ポインタを通して j の値を取得して37で割るった値に更新する（j の値が73に直接更新）
	fmt.Println(j)
}

type Vertex struct {
	X int
	Y int
}

func sample2() {
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)
}

func sample3() {
	v := Vertex{1, 2}
	fmt.Println(v)
	p := &v
	p.X = 1e9
	fmt.Println(v)
}

var (
	v1 = Vertex{1, 2}
	v2 = Vertex{X: 1}
	v3 = Vertex{}
	p  = &Vertex{1, 2}
)

func sample4() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"

	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}

func sample5() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}

	fmt.Println(names)

	a := names[0:2]   // 0 and 1 john paul
	b := names[1:3]   // 1 and 2 paul george
	fmt.Println(a, b) //[john paul] [paul george]

	b[0] = "XXX"
	fmt.Println(a, b)  // [john paul] [XXX george]
	fmt.Println(names) // [John XXX George Ringo]
}

func sample6() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}

var (
	arr   [3]int = [3]int{1, 2, 3}
	slice []int  = []int{1, 2, 3}
)

func sample7() {
	arr1 := [3]int{1, 2, 3}
	arr2 := [5]string{"a", "b"}
	arr3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(len(arr1))
	fmt.Println(len(arr2))
	fmt.Println(len(arr3))
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func sample8() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s) // len=6 cap=6 [2][3][5][7][11][13]

	s = s[:0]     // len=0 cap=6
	printSlice(s) // []

	s = s[:4]     // len=4(インデックス0から3まで) cap=6(インデックス0から元の終わりまで)
	printSlice(s) // [2][3][5][7]

	s = s[2:]     //len=2(インデックス2から1個上の終わりまで) cap=4(インデックス2から元の終わりまで)
	printSlice(s) //[5][7]

	s = s[1:]     //len=1(インデックス1から1個上の終わりまで) cap=3(2インデックス1から元の終わりまで)
	printSlice(s) //[7]
}

func printSlice2(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}

func sample10() {
	a := make([]int, 5)
	printSlice2("a", a)

	b := make([]int, 0, 5)
	printSlice2("b", b)

	c := b[:2]
	printSlice2("c", c)

	d := b[2:5]
	printSlice2("d", d)
}

func sample11() {
	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}

	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func printSlice3(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func sample12() {
	var s []int
	printSlice3(s)

	s = append(s, 0)
	printSlice3(s)

	s = append(s, 1)
	printSlice3(s)

	s = append(s, 2, 3, 4)
	printSlice3(s)
}

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func sample13() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}

func sample14() {
	pow := make([]int, 11)
	for i := range pow {
		pow[i] = 1 << uint(i) //ビットシフト演算
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}

func Pic(dx, dy int) [][]uint8 {
	image := make([][]uint8, dy)

	for y := 0; y < dy; y++ {
		image[y] = make([]uint8, dx)

		for x := 0; x < dx; x++ {
			image[y][x] = uint8(x + y)
		}
	}

	return image
}

func sample15() {
	pic.Show(Pic)
}

type Vertex2 struct {
	Lat, Long float64
}

var m = map[string]Vertex2{
	"Bell Labs": {
		40.68433, -74.39967,
	},
	"Google": {
		37.42202, -122.08408,
	},
}

func sample16() {
	fmt.Println(m)
}

func WordCount(s string) map[string]int {
	var words = strings.Fields(s)
	var m = make(map[string]int)

	for _, word := range words {
		m[word]++
	}
	return m
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func sample17() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))

}

func fibona() func() int {
	a, b := 0, 1
	return func() int {
		result := a
		a, b = b, a+b
		return result
	}
}

func main() {
	f := fibona()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
