package main

import (
	//基本的に1ファイル内において1依存まで
	"fmt"
	"github.com/fenril22/template/funcs"
	"github.com/fenril22/template/control"
	"github.com/fenril22/template/moretype"
	"github.com/fenril22/template/templateandmethod"
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
	fmt.Println("--controller-")
	controller.Dffsample()
	fmt.Println("--pointer-")
	moretype.Pointer()
	fmt.Println("--Struct-")
	moretype.Mkstruct()
	moretype.Strpointer()
	moretype.Globalvprint()
	fmt.Println("--Slices--")
	moretype.Printarray()
	moretype.Printsliceabnormal()
	moretype.Printslicepat1()
	moretype.Printslicenormal()
	moretype.Printsliceofslice()
	moretype.Printsliceappend()
	moretype.Printrange()
	moretype.Printmap()
	fmt.Println("--methods--")
	hoge := templateandmethod.NewVertex(3.2,2.4)
	hoge.Scale(1.0)
	fmt.Println(hoge)
	fmt.Println("--interface--")
	fuga := templateandmethod.NewI()
	fuga.Say()
	piyo := templateandmethod.NewII()
	piyo.Say()
	templateandmethod.Printsaya("yes")
}
