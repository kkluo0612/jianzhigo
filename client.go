// package main

// import (
// 	"errors"
// 	"fmt"
// )

// type Array struct {
// 	data   []int
// 	length uint64
// }

// func NewArray(capacity uint64) *Array {
// 	if capacity == 0 {
// 		return nil
// 	}
// 	return &Array{
// 		data:   make([]int, capacity, capacity),
// 		length: 0,
// 	}
// }

// func (this *Array) len() uint64 {
// 	return this.length
// }

// func (this *Array) isIndexOutOfRange(index uint64) bool {
// 	if index >= uint64(cap(this.data)) {
// 		return true
// 	}
// 	return false
// }

// func (this *Array) Find(index uint64) (int, error) {
// 	if this.isIndexOutOfRange(index) {
// 		return 0, errors.New("index out of range")
// 	}
// 	return this.data[index], nil
// }

// func (this *Array) Insert(index uint64, v int) error {
// 	if this.len() == uint64(cap(this.data)) {
// 		return errors.New("full array")
// 	}
// 	if index != this.length && this.isIndexOutOfRange(index) {
// 		return errors.New("out if index range")
// 	}
// 	for i := this.length; i > index; i-- {
// 		this.data[i] = this.data[i-1]
// 	}
// 	this.data[index] = v
// 	this.length++
// 	return nil
// }

// func (this *Array) InsertToTail(v int) error {
// 	return this.Insert(this.len(), v)
// }

// func (this *Array) Delete(index uint64) (int, error) {
// 	if this.isIndexOutOfRange(index) {
// 		return 0, errors.New("out of index range")
// 	}
// 	v := this.data[index]
// 	for i := index; i < this.len()-1; i++ {
// 		this.data[i] = this.data[i+1]
// 	}
// 	this.length--
// 	return v, nil
// }

// func (this *Array) Print() {
// 	var format string
// 	for i := uint64(0); i < this.len(); i++ {
// 		format += fmt.Sprintf("|%+v", this.data[i])
// 	}
// 	fmt.Println(format)
// }

// func main() {
// 	data := make([]int, 5, 5)
// 	this := &Array{data, 0}
// 	this.Insert(0, 1)
// 	this.Insert(1, 2)
// 	this.Insert(2, 3)
// 	this.Insert(3, 4)
// 	this.Insert(4, 5)
// 	fmt.Println(this.Insert(5, 6))
// 	fmt.Println(this.Find(7))
// }