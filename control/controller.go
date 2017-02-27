package controller

import (
	"fmt"
	"time"
)

func Dffsample() {
	defer fmt.Println("Hoge")
	defer fmt.Println("huga")
	//run on out function
	//stack run

	sum := 0

	// 1 2 (loop) 3 2 (loop) ...
	for i:=0; i<20; i++{
		sum += i
	}

	for ; sum<1000; {
		sum += sum
	}

	for sum < 3000 {
		sum += sum
	}

	for {
		sum--
		if sum<1000 {
			break
		}
	}
	for {
		sum--
		if v:=5; sum<v{
			break
		} else if sum<v*10{
			break
		}
	}
	switch os := 5; os{
		case 1:
			fmt.Println("OS X.")
			//autobreak
			//if non break
			//fallthrough
		case 2:
			fmt.Println("Linux.")
		default:
			fmt.Printf("%s.", os)
	}
	//place many if-else
	t := time.Now()
   switch{
	   case t.Hour() < 12:
		   fmt.Println("Good morning")
	   case t.Hour() < 17:
		   fmt.Println("Good afternoon")
	   default:
		   fmt.Println("Good evening")
   }
}
