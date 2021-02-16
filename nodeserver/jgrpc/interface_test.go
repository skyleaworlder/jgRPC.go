package jgrpcserver

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_CalcuInit(t *testing.T) {
	calcu := new(Calculator)
	calcu.Init()

	f := reflect.ValueOf(calcu.PTB["AddInt8"])
	in := make([]reflect.Value, 2)
	in[0], in[1] = reflect.ValueOf(int8(1)), reflect.ValueOf(int8(2))

	res := f.Call(in)
	arr := make([]int8, 16)
	for i, v := range res {
		arr[i] = (v.Interface().(int8))
	}
	fmt.Print(arr)
}

func Test_Call(t *testing.T) {
	calcu := new(Calculator)
	calcu.Init()

	res := calcu.Call("AddInt8", int8(1), int8(2))
	arr := make([]int8, 16)
	for i, v := range res {
		arr[i] = (v.Interface().(int8))
	}
	fmt.Print(arr, "\n")
}
