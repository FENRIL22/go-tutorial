package main

import (
	"fmt"
	///githubのを入れるなら
	//github.com/user/package
	//下の書き方もできる
	"./funcs"
	//gopherjsとかの時は以下のようにして誤魔化す
	//"github.com/fenril22/template/funcs"
	//もしくはライブラリを機能として作り込んで、githubに入れる
	//多分トップを指定すると自動的に全体を指定してくれるはずなので
	//こちらの方が多分スマート
)

func main(){
	fmt.Println("--Printerfase-----")
	types.Printer()
	fmt.Println("--Printtype-------")
	types.Printtype()
	fmt.Println("--Add-------")
	fmt.Println(types.Add(1,2))
	fmt.Println("--Div-------")
	fmt.Println(types.Div(4,2))
	fmt.Println("--Casting---")
	types.Casting()
	fmt.Println("--Constants-")
	types.Constants()
}
