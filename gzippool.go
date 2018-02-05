package main

import (
	"compress/gzip"
	"io"
	"sync"
)

type gatew struct {
	*gzip.Writer
}

type Writer struct {
	gatew
}

var zippers = sync.Pool{New: func() interface{} {
	wr := gzip.NewWriter(nil)
	return &Writer{gatew{wr}}
}}

func NewWriter(w io.Writer) *Writer {
	wr := zippers.Get().(*Writer)
	wr.Reset(w)
	return wr
}

func (gw *Writer) Close() error {
	e := gw.gatew.Close()
	zippers.Put(gw)
	return e
}
