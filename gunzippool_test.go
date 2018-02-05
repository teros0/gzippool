package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"sync"
	"testing"
)

var (
	gunzippers = sync.Pool{New: func() interface{} {
		abuf := []byte{31, 139, 8, 0, 0, 0, 0, 0, 0, 255}
		arbuf := bytes.NewReader(abuf)
		wr, _ := gzip.NewReader(arbuf)
		return wr
	}}
)

func readerNaive() {
	data := []byte{31, 139, 8, 0, 0, 0, 0, 0, 0,
		255, 203, 72, 205, 201, 201, 7, 0, 134, 166, 16, 54, 5, 0, 0, 0}
	d := bytes.NewReader(data)

	gz, e := gzip.NewReader(d)
	if e != nil {
		fmt.Println("NIL READER", e)
	}

	gz.Reset(d)
	defer gz.Close()
	//io.Copy(os.Stderr, gz)
}

func readerPool() {
	data := []byte{31, 139, 8, 0, 0, 0, 0, 0, 0,
		255, 203, 72, 205, 201, 201, 7, 0, 134, 166, 16, 54, 5, 0, 0, 0}
	d := bytes.NewReader(data)

	gz := gunzippers.Get().(*gzip.Reader)
	gz.Reset(d)

	defer gunzippers.Put(gz)
	defer gz.Close()
}

func readerMyPool() {
	data := []byte{31, 139, 8, 0, 0, 0, 0, 0, 0,
		255, 203, 72, 205, 201, 201, 7, 0, 134, 166, 16, 54, 5, 0, 0, 0}
	d := bytes.NewReader(data)

	gz, _ := NewReader(d)
	gz.Reset(d)
	defer gz.Close()
}

func BenchmarkReadNaive(b *testing.B) {
	b.N = 3000
	for i := 0; i < b.N; i++ {
		readerNaive()
	}
}

func BenchmarkReadPool(b *testing.B) {
	b.N = 3000
	for i := 0; i < b.N; i++ {
		readerPool()
	}
}

func BenchmarkReadMyPool(b *testing.B) {
	b.N = 3000
	for i := 0; i < b.N; i++ {
		readerMyPool()
	}
}
