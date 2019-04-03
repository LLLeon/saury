package workpool

import (
	"sync"
	"testing"
)

func TestWorker_Pool(t *testing.T) {
	dispatcher := NewDispatcher(4)
	dispatcher.Run()
	counter := &counter{}

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go dispatcher.PushToJobQ(counter)
	}

	wg.Wait()

	if count != 20 {
		t.Errorf("count error: want: 20, get: %v\n", count)
	} else {
		t.Logf("count: %v\n", count)
	}

}

var (
	count int
	wg    = sync.WaitGroup{}
)

type counter struct{}

func (c *counter) Do() error {
	defer wg.Done()
	count++
	return nil
}
