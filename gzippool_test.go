package main

import (
	"bytes"
	"compress/gzip"
	"sync"
	"testing"
)

var (
	gzippers = sync.Pool{New: func() interface{} {
		wr := gzip.NewWriter(nil)
		return wr
	}}
)

func writerPool() {
	buf := bytes.NewBuffer(make([]byte, 100))
	gz := gzippers.Get().(*gzip.Writer)
	defer gzippers.Put(gz)
	defer gz.Close()
	gz.Reset(buf)
	gz.Write([]byte("TESTTESTTEST"))
}

func writerNaive() {
	buf := bytes.NewBuffer(make([]byte, 100))
	gz, _ := gzip.NewWriterLevel(buf, 1)
	defer gz.Close()
	gz.Write([]byte("TESTTESTTEST"))
}

func writerMyPool() {
	buf := bytes.NewBuffer(make([]byte, 100))
	gz := NewWriter(buf)
	defer gz.Close()
	gz.Write([]byte("TESTTESTTEST"))
}

func BenchmarkWritePool(b *testing.B) {
	b.N = 3000
	for i := 0; i < b.N; i++ {
		writerPool()
	}
}

func BenchmarkWriteMyPool(b *testing.B) {
	b.N = 3000
	for i := 0; i < b.N; i++ {
		writerMyPool()
	}
}

func BenchmarkWriteNaive(b *testing.B) {
	b.N = 3000
	for i := 0; i < b.N; i++ {
		writerNaive()
	}
}
