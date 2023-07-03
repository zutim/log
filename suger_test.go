package log

import "testing"

func TestSuger(t *testing.T) {
	userLog := NewLogMap().WithOptionPath("log/user.log")
	userLog.Info("this will print in log/user.log")

	orderLog := NewLogMap().WithOptionPath("log/order.log")
	orderLog.Info("this will print in log/order.log")
}
