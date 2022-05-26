package util

import (
	"errors"
	"runtime/debug"
	"testing"
)

type Value interface {
	uint | int | float64 | string | byte | bool
}

var ErrFailed = errors.New("failed")

func Expect[T Value](t *testing.T, received T, expected T) {
	if received != expected {
		t.Log("expected:", expected)
		t.Log("received:", received)
		t.Log(string(debug.Stack()))
		panic(ErrFailed)
	}
}

func ExpectArray[T Value](t *testing.T, received []T, expected []T) {
	if len(received) != len(expected) {
		t.Log("expected:", expected)
		t.Log("received:", received)
		t.Log(string(debug.Stack()))
		panic(ErrFailed)
	}
	for i := 0; i < len(received); i++ {
		if received[i] != expected[i] {
			t.Log("expected:", expected)
			t.Log("received:", received)
			t.Log(string(debug.Stack()))
			panic(ErrFailed)
		}
	}
}

type ListNode[T Value] struct {
	Val  T
	Next *ListNode[T]
}

type TreeNode[T Value] struct {
	Val   T
	Left  *TreeNode[T]
	Right *TreeNode[T]
}

func ListToArray[T Value](list *ListNode[T]) []T {
	var arr []T
	for p := list; p != nil; p = p.Next {
		arr = append(arr, p.Val)
	}
	return arr
}

func ArrayToList[T Value](arr []T) *ListNode[T] {
	var h = &ListNode[T]{}
	var p = h
	for _, v := range arr {
		p.Next = &ListNode[T]{
			Val: v,
		}
		p = p.Next
	}
	return h.Next
}
