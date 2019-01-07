package rpc

import (
	"io"
)

type RWCloseCombiner struct {
	w io.Writer
	r io.Reader
}

func NewRWCloseCombiner(w io.Writer, r io.Reader) *RWCloseCombiner {
	return &RWCloseCombiner{
		w: w,
		r: r,
	}
}

func (rwc *RWCloseCombiner) Read(p []byte) (int, error) {
	return rwc.r.Read(p)
}

func (rwc *RWCloseCombiner) Write(p []byte) (int, error) {
	return rwc.w.Write(p)
}

func (rwc *RWCloseCombiner) Close() error {
	return nil
}
