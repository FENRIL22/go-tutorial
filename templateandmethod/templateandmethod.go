package templateandmethod

import "fmt"

//普通に型を作る
type Vertex struct {
	X, Y float64
}
//structだけでなく普通の型を利用して定義することもできる
type Myfloat float64

// Method宣言 //
//func宣言の前にレシーバ引数を定義することでメソッド定義が可能
//レシーバ宣言は同一パッケージ内で行う
//2種類の宣言方法があるがどちらかに統一して記載することが望ましい

//変数レシーバ
//構造体内の値を変更しないなどの場合は使える
func (f Myfloat) Abs() float64{
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}
//ポインタレシーバ
//作った後に構造体などの内部を操作する必要がある時
//おおよそはこちらを使っておけばいい
//(コピーのコストが問題になるため)
//float64ならコピーが早いとか色々ある
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

//上の構造体や下のインタフェースは外部に公開することはできるものの，以下の条件を満たさないかぎり他パッケージからアクセスできない
//・package名が同じものになっている
//・かつ同一ファイル内(同一階層上にいる)こと
//(基本ファイル名=package名で機能ごとにフォルダを分け実装を行うのが普通)
//・もしくはimport時にpathの前に". "を付ける
//(その箇所にそのpackageを全て展開する．あまりしない方が良い)
//これらを利用する際は，New~という関数を作り，その関数が構造体やinterfaceそのものやポインタを返させ，その結果を利用することが一般的である
func NewVertex(x, y float64) (Vertex) {
	v := Vertex{x,y}
	return v
}


/*tips
構造体などを生成する時，ポインタを使うか構造体そのものを使うかを選ぶ事ができるが，
メソッドについては特段(*V.METHOD)をせず(V.METHOD)としてもアクセスできる
そのためメソッド数が多い(10個とか)場合ポインタ宣言すると良い(ヒープ領域)．
immutableなものはそのものとして宣言し，mutableなものはポインタとして宣言したりする例もある
*/

/*// interface //
goでかなり汎用性の高いもの．役割としては以下．
1. 実装の分離
interfaceはメソッド列挙で成り立つ．必要に応じて型を併記する．
ここにメソッドを挙げることで，処理を一覧できるほか，実装とデータを分離できる
2. 構造体の隠蔽
基本的には，全ての型は空interfaceを持つ．
上のメソッド宣言は一種これに対して実装を行っていると言っていい．
またinterfaceだけを公開するとstructに対しての操作方法はinterfaceのみになり，直接の操作が行われないなるという安全点もある．
(structの子の頭を小文字にするというのも対策)
3.ダックタイピンク
直接メソッドを定義してもよいが，interfaceによる宣言を行うことで，そのメソッドを実装していることを保証できる(実装しなくても見た目上表現できる)
4.全ての型の許容
空interfaceは全型を入れることができる．そのため未知の入力への対応が可能．
> Ue は整理中
> interfaceは実装，structはデータとして捉える．
> 実装の際に，共通項として括れる実装はinterfaceにレシーバを，各個で処理が異なる場合はstruct自体にという実装とする
> そして，処理を呼び出す時はinterfaceに入れて呼ぶと，同じ名前のmethodが呼ばれる
( 適宜interfaceに定義されているもの，structに定義されているものが良い感じに選ばれて実行されると思う)*/

//>> まとめ
//>> templateはメソッドのシグニチャの集合として示す．
//>> これはそのメソッドを定義した型を作るようなもの．
type I interface {
	Say()
}
//>> そしてこれと同じメソッドが定義されている場合に代入して実行が可能．
//>> (余分なものが宣言されていても問題ないが，足りない場合弾かれる)
//>> この時レシーバがポインタレシーバであるとその型のポインタに対しての処理となるため，構造体宣言はポインタで行わなければ実行できない
//>> 変数レシーバである時も同様
type hoge struct{
	dat int
}
type nugat int
//メソッド(interfaceにない．interfaceのみを返す方式にすると，ここにはアクセスできないため，内部処理だけで用いることができる(処理の隠蔽))
func (h hoge) M(){
	fmt.Println(h.dat)
}
//メソッド(interfaceに書いてるもの．これだけが公開される)
//ここで，同じメソッド名による宣言を行っている
//これにより，同じinterfaceに代入した時に同じメソッド名を実行することで，違う処理を行うことができる
//これでポリフォーニズムを表現できるほか，共通処理の括り出しにも使える
//一種，メソッドをそれぞれ宣言して作り，公開するものだけをinterfaceで選び出して公開するという選択用のツールとしてinterfaceを使うのもありかもしれない
func (h hoge) Say() {
	fmt.Println("Hello")
}
func (n nugat) Say() {
	fmt.Println(n)
}
//ここで返すものも，操作が必要であれば全てポインタで統一し，そうでないなら変数で統一する
//interfaceはその名前そのものを返すものとして記載
//なお，何も代入されていないinterfaceはnilをレシーバとして呼び出す
//これを処理しないとnullptrするためt == nilのようなチェックを各メソッドに持たせておくのが通例
//(型も値も持っていないため)
func NewI() I {
	var i I
	i = hoge{1}
	return i
}
func NewII() I {
	var i I
	i = nugat(1)
	return i
}

//空interface{}を用いると，全ての型を許容可能な関数を作成可能である
//fmtのprintは内部ではこれを用いている
func Printsaya(i interface{}){
	//型分岐
	//typeとすることで型による分岐を行える
	//あまり好まれない仕組みの様子
	switch v := i.(type) {
	case int:
		//型アサーション
		//指定の型に変換することができる
		//value, okで構成され，1つ目が値，2つ目が変換できたかの有無である
		//2つ目を取らなかった時かつその型がinterfaceに入っていなかった時，panicする
		//2つ目を取り型が異なる場合は0値とfalseが返り，処理は続行される
		s := i.(int)
		fmt.Println(s)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

/*>> 実装手法 <<
>> てっとり早く実装したい or 速度を重視したい -> struct + method
>> 共通の処理を括り出したい -> interface + struct + method(for interface + for struct)
>> 実装を隠したい -> interface + struct + method(for interface)
>>>> interfaceへの実装とstructへの実装がぶつかった時にどういう挙動をするか
*/

//goのダックタイピング//
//goはダックタイピングを採用しているため，その項目が実装されてさえいれば，そのインタフェースとしてみなすということができる
//つまり，以下のように標準で設定されているメソッドを定義することにより，その標準でのインタフェースと同様の振舞いができ，その結果デバッグや処理の入出力において利便性が上がる．
//方法は，普通にその型にメソッドを実装するだけであるが，interfaceの時同様pointerレシーバと変数宣言時の宣言方法(ポインタなど)に注意

//Stringer (fmt interface)
//interface宣言は次のとおり
//type Stringer interface{
//	String() string
//}
//これを以下のように用意すると自動的にその形式で出力が行われる．
//なおこの時上のようなinterfaceを宣言する必要はない
type S1 struct{
	x,y int
}
func (s S1) String() string {
	return fmt.Sprintf("x: %v y: %v, x*y : %v", s.x, s.y, s.x*s.y)
}
func NewS1(x, y int) *S1 {
	return &S1{x,y}
}

//Error
//type error interface {
//    Error() string
//}
//基本的に成功=nil/失敗=それ以外
//それ以外の時の値にエラー対応を割り当てたり表示を分かりやすくするのに用いる
type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprint("cannot Sqrt negative number: ",float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x<0.{
		return 0, ErrNegativeSqrt(x)
	}
	z := 1.

	for i:= 0 ; i<=10; i++ {
		z = z - (z*z-x)/(2*z)
	}
	return z, nil
}

//io.Reader
//A tour of go21(Readers)
