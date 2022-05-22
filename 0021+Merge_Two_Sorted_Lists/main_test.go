package main

import (
	"app/util"
	"testing"
)

func TestCases(t *testing.T) {
	util.ExpectArray(t, util.ListToArray(mergeTwoLists(util.ArrayToList([]int{1, 2, 4}), util.ArrayToList([]int{1, 3, 4}))), []int{1, 1, 2, 3, 4, 4})
}
