package main

// https://golang.org/pkg/container/list/
// type Element
//     func (e *Element) Next() *Element
//     func (e *Element) Prev() *Element
// type Element struct {
//     // The value stored with this element.
//     Value interface{}
//     // contains filtered or unexported fields
// }
// type List
//     func New() *List
//     func (l *List) Back() *Element
//     func (l *List) Front() *Element
//     func (l *List) Init() *List
//     func (l *List) InsertAfter(v interface{}, mark *Element) *Element
//     func (l *List) InsertBefore(v interface{}, mark *Element) *Element
//     func (l *List) Len() int
//     func (l *List) MoveAfter(e, mark *Element)
//     func (l *List) MoveBefore(e, mark *Element)
//     func (l *List) MoveToBack(e *Element)
//     func (l *List) MoveToFront(e *Element)
//     func (l *List) PushBack(v interface{}) *Element
//     func (l *List) PushBackList(other *List)
//     func (l *List) PushFront(v interface{}) *Element
//     func (l *List) PushFrontList(other *List)
//     func (l *List) Remove(e *Element) interface{}
