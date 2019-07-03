/*************************************************************************
+Author   : chenhuijia@deepglint.com
+Data     : 2019-07-03
+************************************************************************/

package fileop

import (
	"fmt"
	"os"
	"sync"
	"testing"
)

func TestSyncWriter_Write(t *testing.T) {
	f, err := os.Create("test.txt")
	if err != nil {
		t.Error(err)
	}

	wg := &sync.WaitGroup{}
	sw := NewSyncWriter(f)
	data := []string{"a", "b", "c", "d"}

	for _, v := range data {
		wg.Add(1)
		go func(s string) {
			defer wg.Done()
			fmt.Fprintln(sw, s)
		}(v)
	}

	wg.Wait()

	t.Log("success")
}
