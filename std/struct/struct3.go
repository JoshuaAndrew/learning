/*
Method Expressions:
MethodExpr    = ReceiverType "." MethodName
ReceiverType  = TypeName | "(" "*" TypeName ")" | "(" ReceiverType ")"

*/
package main

import (
	"fmt"
)

//可以自由的向基本类型中添加方法
type WW int

func (w *WW) Print(flag string) {
	fmt.Println(flag)
}

func main() {
	var ww WW
	ww.Print("weiweiu")         //method is value
	ww.Print("Isoftstone Inc.") //method is value

	(*WW).Print(&ww, "weiweiu")         //Method is Expressions
	(*WW).Print(&ww, "Isoftstone Inc.") // Method is Expressions
}
