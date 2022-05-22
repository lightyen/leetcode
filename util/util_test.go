package util_test

import (
	"app/util"
	"testing"
)

func c(t *testing.T, callback func(), excepted error) {
	defer func() {
		err := recover()
		if excepted != err {
			t.Log("util excepted:", excepted)
			t.Log("util received:", err)
			t.FailNow()
		}
	}()
	callback()
}

func TestExcept(t *testing.T) {
	c(t, func() {
		util.Expect(t, 0, 0)
	}, nil)
	c(t, func() {
		util.Expect(t, 0, 1)
	}, util.ErrFailed)
	c(t, func() {
		util.ExpectArray(t, []int{1, 2, 3}, []int{1, 2, 3})
	}, nil)
	c(t, func() {
		util.ExpectArray(t, []int{1, 2, 3}, []int{1, 2, 4})
	}, util.ErrFailed)
}
