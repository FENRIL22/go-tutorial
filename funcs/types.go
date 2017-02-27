package types

import "fmt"

func returnstring() string{
	//return string on program on Printf rule
	return fmt.Sprint("hoge")
}

func Printer() int {
	var ps int = 10

	//print as is type
	//c printf like
	fmt.Printf("First point is %v\n", ps)
	//insert space at "," point
	//for non space, use "+"
	fmt.Println("Second point is",ps)
	//and more...
	fmt.Println(returnstring())

	return 0
}

func Printtype() (hoge, huga int) {
	hoge, huga = 10, 20
	var inu, neko string = "neow", "bug"
	nugat := 13

	fmt.Println(hoge,huga)
	fmt.Printf("%v, %q\n", inu, neko)
	fmt.Printf("%T\n", nugat)

	return
}

func Add(x int, y int) (int){
	return x+y
}

func Div(x , y int) int{
	return x/y
}

func Casting() int{
	var tora float64 = 1.23
	//can't run
	//var hoge int = tora
	var torakiti int = int(tora)
	torako := int(tora)
	fmt.Printf("%T, %T, %T\n", tora, torakiti, torako)
	return 0
}

func Constants() {
	const hoge = 10
	// hoge := 10
	//only string, boolean, num
	//cant rewrite
	fmt.Println(hoge)
}
