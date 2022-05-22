package main

import (
	"app/util"
	"testing"
)

func TestCases(t *testing.T) {
	util.ExpectArray(t, util.ListToArray(addTwoNumbers(util.ArrayToList([]int{2, 4, 3}), util.ArrayToList([]int{5, 6, 4}))), []int{7, 0, 8})
}
