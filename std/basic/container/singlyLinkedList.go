package main

type Element struct {
	Value interface{}
	next  *Element
}

type List struct {
	e    Element
	size int
}

//初始化一个空链表
func New() *List {
	l := new(List)
        l.e.next = &l.e
        l.size = 0
	return l
}
// |v|next|   |v|next|
func (l *List) Add(e, at *Element) *Element {
       l.e.Value == at.Value
}

func (l *List) Remove() *Element {

}

func (l *List) Size() int {
	return l.size
}

func main() {
}
