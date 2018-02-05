package main

import (
	"bytes"
	"compress/gzip"
	"io"
	"sync"
)

type gater struct {
	*gzip.Reader
}

type Reader struct {
	gater
	e error
}

var (
	unzippers = sync.Pool{New: func() interface{} {
		buf := []byte{31, 139, 8, 0, 0, 0, 0, 0, 0, 255}
		rbuf := bytes.NewReader(buf)
		wr, err := gzip.NewReader(rbuf)
		return &Reader{gater{wr}, err}
	}}
)

func NewReader(r io.Reader) (*Reader, error) {
	gr := unzippers.Get().(*Reader)
	if gr.e != nil {
		return nil, gr.e
	}
	gr.Reset(r)
	return gr, nil
}

func (gr *Reader) Close() error {
	e := gr.gater.Close()
	unzippers.Put(gr)
	return e
}
