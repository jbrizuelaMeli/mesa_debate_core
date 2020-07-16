package tests

import (
	"fmt"
	"testing"
)

type Foo struct {
	foo int64
	bar int64
	baz int64
}

type Bar struct {
	foo int64
	bar int64
	baz int64
	bah int64
}

var bts []byte

func printMsg(b *Bar) {
	fmt.Println(b)
}

func BenchmarkAllocateFooStack(b *testing.B) {
	for i := 0; i < b.N; i++ {
		func() Foo {
			return Foo{}
		}()
	}
}

func BenchmarkAllocateBarStack(b *testing.B) {
	for i := 0; i < b.N; i++ {
		func() Bar {
			return Bar{}
		}()
	}
}

func BenchmarkAllocateFooHeap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		func() *Foo {
			return new(Foo)
		}()
	}
}

func BenchmarkAllocateBarHeap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		func() *Bar {
			b := new(Bar)
			//printMsg(b)
			return b
		}()
	}
}

func BenchmarkAllocateSliceHeapNoEscape(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bts := make([]byte, 1024)
		bts[0] = 1
	}
}

func BenchmarkAllocateSliceHeapEscape(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bts = make([]byte, 1024)
		bts[0] = 1
	}
}
