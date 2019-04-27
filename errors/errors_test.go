package errors

import (
	"testing"
	goerrors "errors"
	"fmt"
)

func TestWithCode(t *testing.T) {
	err0 := WithStack(goerrors.New("caused by xxxXxxx"))
	err1 := WithCode(err0, "0001", "第一个错误")

	fmt.Println("=======================")
	fmt.Printf("%+v\n", err0)
	fmt.Println("=======================")
	fmt.Printf("%+v\n", err1)

	fmt.Println("=======================")
	fmt.Println(Cause(err1))
}