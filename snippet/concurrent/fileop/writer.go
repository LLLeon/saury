/*************************************************************************
+Author   : chenhuijia@deepglint.com
+Data     : 2019-07-03
+************************************************************************/

package fileop

import (
	"io"
	"sync"
)

type SyncWriter struct {
	mux    sync.Mutex
	Writer io.Writer
}

func NewSyncWriter(writer io.Writer) *SyncWriter {
	return &SyncWriter{
		Writer: writer,
	}
}

func (sw *SyncWriter) Write(b []byte) (n int, err error) {
	sw.mux.Lock()
	defer sw.mux.Unlock()
	return sw.Writer.Write(b)
}
