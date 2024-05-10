package log

import (
	"fmt"
	"testing"
)

func TestSuger(t *testing.T) {
	userLog := NewLogMap().WithOptionPath(LoggerOptions{Path: "log/user"})
	userLog.Info("this will print in log/user.log")

	orderLog := NewLogMap().WithOptionPath(LoggerOptions{Path: "log/order"})
	orderLog.Info("this will print in log/order.log")

	a1 := args{Name: "a1"}
	// 打印a1地址
	fmt.Printf("Pointer value: %p\n", &a1)
	fmt.Println(a1.Name)
	copy2(&a1)

	fmt.Printf("Pointer value: %p\n", &a1)
	fmt.Println(a1.Name)
}

type args struct {
	Name string
}

func copy(aa *args) {
	b1 := args{Name: "b1"}
	aa = &b1
}

func copy2(aa *args) {
	b1 := args{Name: "b1"}
	*aa = b1
}
