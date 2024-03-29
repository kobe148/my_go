package algo

import (
	"errors"
	"fmt"
)

/**
 * 1.数组的插入，删除，按照下标随机访问操作；
 * 2.数组中的数据类型均为int类型
 *
 * Author rlf
 */

type Array struct {
	data []int
	length uint
}

//为数组初始化内存
func NewArray(capacity uint) *Array {
	if capacity == 0 {
		return nil
	}

	return &Array{
		data: make([]int, capacity, capacity),
		length: 0,
	}
}

//数组长度
func (this *Array) Len() uint {
	return this.length
}

//判断索引是否越界
func (this *Array) isIndexOutOfRange(index uint) bool {
	if index >= uint(cap(this.data)) {
		return true
	}

	return false
}
//通过索引查找数组，索引范围为[0, n-1]
func (this *Array) Find(index uint) (int, error) {
	if this.isIndexOutOfRange(index) {
		return 0, errors.New("数组下标越界")
	}

	return this.data[index], nil
}
//插入数值到索引index上
func (this *Array) Insert(index uint, v int) error {
	if this.length == uint(cap(this.data)) {
		return errors.New("数组满了")
	}
	if index != this.length && this.isIndexOutOfRange(index) {
		return errors.New("数组下标越界")
	}
	for i := this.length; i > index; i-- {
		this.data[i] = this.data[i-1]
	}
	this.data[index] = v
	this.length++
	return nil
}
func (this *Array) InsertToTail(v int) error {
	return this.Insert(this.Len(), v)
}
//删除索引的值
func (this *Array) Delete(index uint) (int, error) {
	if this.isIndexOutOfRange(index) {
		return 0, errors.New("out of index range")
	}
	v := this.data[index]
	for i := index; i < this.Len() - 1 ; i++ {
		this.data[i] = this.data[i+1]
	}
	this.length--
	return v, nil
}

//打印数组
func (this *Array) Print() {
	var format string
	for i := uint(0); i < this.Len(); i++ {
		format += fmt.Sprintf("|%+v", this.data[i])
	}

	fmt.Println(format)
}