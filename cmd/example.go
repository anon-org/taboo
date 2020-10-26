package main

import (
	"errors"
	"fmt"
	"github.com/anon-org/taboo/pkg/taboo"
)

func foo(a, b int) int {
	return a / b
}

func bar(a, b int) int {
	var result int

	taboo.Try(func() {
		result = foo(a, b)
	}).Catch(func(e *taboo.Exception) {
		e.Throw("broken")
	}).Do()

	return result
}

func baz() {
	taboo.Throw(errors.New("panicc"))
}

func main() {
	taboo.Try(func() {
		bar(1, 0)
	}).Catch(func(e *taboo.Exception) {
		fmt.Println(e.Error())
	}).Do()

	taboo.Try(func() {
		baz()
	}).Catch(func(e *taboo.Exception) {
		fmt.Println(e.Error())
	}).Do()
}
