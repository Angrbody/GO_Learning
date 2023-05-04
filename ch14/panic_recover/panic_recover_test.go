package panicrecover_test

import (
	"errors"
	"fmt"
	"testing"
)

func TestPanicVxExit(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recovered from ", err)
		}
	}()
	// defer func() {
	// 	fmt.Println("Finally!")
	// }()
	fmt.Println("Start")
	panic(errors.New("Something wrong!"))
	// os.Exit(-1)
}
