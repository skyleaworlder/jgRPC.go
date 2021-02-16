package jgrpcserver

import (
	"fmt"
	"reflect"
	"strconv"
)

// RPC is an interface support numeric calculation
type RPC interface {
	Init()
	Call(FuncName string, Param ...interface{}) (res []reflect.Value)
}

// Calculator is a struct implement numeric
type Calculator struct {
	Config    map[string]string
	PTB       map[string]interface{}
	ptbMaxNum int
}

// Init is a method to initialize Calculator
// self-defined
func (c *Calculator) Init() {
	c.PTB = make(map[string]interface{}, c.ptbMaxNum)
	c.PTB["AddInt8"] = c.AddInt8
}

// Call is a method to call a interface method
// since result's type might be various, so return []reflect.Value
func (c *Calculator) Call(FuncName string, Param ...interface{}) (res []reflect.Value) {
	// if FuncName in PTB
	if _, ok := c.PTB[FuncName]; !ok {
		res = []reflect.Value{}
		return
	}
	// function descriptor
	// get the function handler identified by FuncName
	fd := reflect.ValueOf(c.PTB[FuncName])

	fmt.Println("About function " + FuncName + ":")
	// process parameters input
	param := make([]reflect.Value, len(Param))
	for i, val := range Param {
		param[i] = reflect.ValueOf(val)
		fmt.Println("No."+strconv.Itoa(i)+" Parameter is:", val)
	}

	// for debug
	// fmt.Println("Calculator.Call:", len(Param), Param)

	// use function descriptor & parameters
	res = fd.Call(param)
	return
}

// AddInt8 is a function to calculate a + b
func (c *Calculator) AddInt8(a, b int8) int8 {
	return addint8(c, a, b)
}
