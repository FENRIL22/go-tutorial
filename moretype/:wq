package moretype

import (
	"fmt"
	"strings"
)

//struct
type Vertex struct {
	X int
	Y int
}

//初期化の説明のため記載
//varとして並べて表記できる
var(
	v1 = Vertex{1,2}//X=1,Y=2設定
	v3 = Vertex{}	//自動初期化(0/0)
	v2 = Vertex{X: 1}//Xのみ指定して初期化(X: 1)
	p = &Vertex{1,2} //structへのポインタを格納
	//ポインタはmain内で宣言したstructを他場所でも参照渡ししたいなどの時に使うかも
)

func Globalvprint(){
	fmt.Println(v1,v3,v2,p)
}

func Mkstruct(){
	v := Vertex{1,2}
	//Access
	v.X = 4
	fmt.Println(v)
}

func Strpointer(){
	//structのポインタ経由のアクセス
	//本来は(*p).Xが正しいが略記として以下が使える
	//これはメソッド定義に良く用いる
	v := Vertex{1,2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
}

func Pointer() {
	i := 42
	p := &i
	*p = *p * 3
	fmt.Println(*p)
	fmt.Println(i)
}

func Printarray(){
	//この配列は値渡しになる
	//そのため関数で操作させる場合はポインタを渡す
	//また配列長は型の一部のため不変
	//利便性なども含めスライスに軍配が挙がる
	primes := [3]int{2,3,4}
	primes[2] = 23232
	fmt.Println(primes)
}

func Printsliceabnormal(){
	primes := [6]int{2,3,4,5,6,7}
	//要素の0から4を切り出したスライス
	//後ろの数字は0からでなく1から始まる数値で数える
	//つまり0:0で長さ0のスライスを生成
	//またこれは元配列のポインタみたいなものなので
	//何らかの変更を行うと共有するスライスは全て影響を受ける
	//これにより関数に渡す時はそのまま渡せる
	//さらに型の指定が簡単になるので便利
	//より高度な型指定はinterfaceで行う
	//0値のスライス(nilスライス)
	s := primes[0:0]
	//4以降を切り出したもの
	v := primes[2:]
	//2より手前
	g := primes[:3]
	//全て
	all := primes[:]
	//これを足すと全ての4が90になる
	//all[2] = 90
	fmt.Println(s,v,g,all)
}

func Printslicepat1(){
	//初期化を伴うsliceの作り方
	//この場合は自動で適切な配列を作成してくれる
	q := []int{2,3,4}
	r := []bool{true,false}
	//構造体もこんな感じに
	s := []struct{
		i int
		b bool
	}{
		{2,true},
		{3,false},
		{4,true},
	}
	//削り落とすこともできる
	//また[:]とすることで再びアクセスできるようになる(再スライス)
	//ただし頭を切り落とすと([1:]など)その箇所は参照不可能になる
	//詳細はmakeによるsliceで
	s = s[:2]

	fmt.Println(q,r,s)
}

func Printslicenormal(){
	//sliceは長さlenと容量capを持つ
	//lenはそのsliceが持つ要素数
	//capはその配列が持ちうる量(元となる配列の要素数)
	//なので5個のintが入るのを使いたければ次をsliceに入れるとできる
	//[5]int
	//より便利な宣言方法はmakeを用いる
	//len,capを指定(capは任意)
	s := make([]int, 3, 5)
	//capが変化していなければ再スライスによりその分が使えるようになる
	//ただ頭については削るとアクセス不能になる模様
	//s = s[:cap(s)]
	fmt.Printf("len=%d, cap=%v, list %v\n",len(s),cap(s),s)
	//絶対的な容量が必要ならスライス後のを別のスライスとして保存するか
	//頭を切らないか、配列を作ってからスライスするかを取る
	//またはlenだけを指定してappendなどを用いて容量をあまり気にしないで使えるようにするか
}


func Printsliceofslice(){
	//スライスは複数の同じ型?を内蔵できる。
	//もしかしたら他の型も内包できるのかも
	//配列を内包する配列みたいな感じで[][]
	board := [][]string{
		[]string{"-","-","-"},
		[]string{"-","-","-"},
	}


	for i:=0; i<len(board); i++ {
		fmt.Printf("%s\n",strings.Join(board[i], " "))
	}
}

func Printsliceappend(){
	var s []int
	//appendすることで要素を追加できる
	//容量が不足する時は自動で再生成される
	s = append(s, 0)
	//複数の値をappendできる
	//ただスライス外に足される形になるらしい
	//速度等が必要な時はcapを任意に、lenを0にしておけばいける？
	s = append(s,1,2)

	f:= []int{1,2,3,4}
	//...を付けるとスライス同士の結合もできる
	//ただ1つずつしか結合できない模様
	s = append(s, f...)
	fmt.Println(s)
}

func Printrange() {
	//rangeは配列などから要素を1つずつ取り出すことができる。
	//pythonのrangeみたいな
	//forの終了条件の書き型を用いて書く
	//(foreachにあたるらしい)
	//配列を食わせる時は[]を省略した形で入れる
	//([][]の時は[]が1つ付く)
	//戻り値は1つ目がindex/2つ目が値
	//1つだけ書くとindexが得られる
	//2つ目だけが欲しい時は＿を書くと値を捨てられる
	var pow = []int{1,2,4,8}
	//以下のどれか
	//for i := range pow
	//for _,v := range pow
	for i,v := range pow{
		fmt.Println(i,v)
	}
}

type fertex struct{
	Lat,Long float64
}
func Printmap() {
	//[key]value の型を宣言
	//作るのはmake(宣言だけでは使えない)
	m := make(map[string]fertex)
	//もしくは初期化変数付き(keyの指定が必要)
	// hoge = map[string]Vertex{
	//	"Bell" : Vertex{
	//		32.3, -43.4,
	//	},
	//	"Google" : Vertex{
	//		23.4,-122.08,
	//	},
	//}
	//以下のように内部の型を省略して書く事もできる
	//(同じVertex型であると推測できるため)
	//hoge = map[string]Vertex{
	//	"Bell" : { 32.3, -43.4},
	//	"Google" : {23.4,-122.08},
	//}



	//入れるのは値に対して入れるだけ。
	//また値が存在する時は上書きされる
	m["bell"] = fertex{
		40.68433, -74.34967,
	}
	//参照はkeyの指定
	//第2戻り値が存在の有無を示す(省略可)
	//初期化は適当なkeyの戻り値を入れさせたら便利かな
	//elemは今回はfertex型,okは常にbool
	elem, ok := m["bell"]
	fmt.Println("The value:",elem,"Present?",ok)
	//削除はdelete
	//delete(m,key)
	fmt.Println(m)
	//m["bell"])
	//mapは参照型なので関数から直接操作が可能
}

/*func Printclosure(){
	//関数を引数に取らせることもできる
	//つまり関数を戻り値として返すものも作れる
	//これを利用してClosureを作ることもできる
	//つまり関数内で状態を保持しており呼び出すたびにその値を更新するような具合である。
	//ある種初期生成はコンストラクタみたいな動作と言えるかもしれない
	//これについては必要になった時に記載する
}*/
