package rpc

import (
	"bytes"
	"io"
	"log"
	"sync"
)

type RWCloseCombiner struct {
	mtx         sync.Mutex
	w           io.Writer
	Buf         *bytes.Buffer
	newDataChan chan bool
}

func NewRWCloseCombiner(w io.Writer) *RWCloseCombiner {
	var buf bytes.Buffer

	return &RWCloseCombiner{
		w:   w,
		Buf: &buf,
	}
}

func (rwc *RWCloseCombiner) Read(p []byte) (int, error) {
	rwc.mtx.Lock()
	defer rwc.mtx.Unlock()
	if rwc.Buf.Len() == 0 {
		return 0, nil
	} else {
		log.Println(rwc.Buf.String())
	}
	return rwc.Buf.Read(p)
}

func (rwc *RWCloseCombiner) Write(p []byte) (int, error) {
	rwc.mtx.Lock()
	defer rwc.mtx.Unlock()
	return rwc.w.Write(p)
}

func (rwc *RWCloseCombiner) Close() error {
	return nil
}
