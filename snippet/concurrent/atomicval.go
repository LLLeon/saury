package concurrent

import (
	"fmt"
	"sync/atomic"
)

type Value struct {
	Key string
	Val interface{}
}

type BoxOffice struct {
	Movies atomic.Value
	Total  atomic.Value
}

func NewBoxOffice() *BoxOffice {
	bo := &BoxOffice{}
	bo.Movies.Store(&Value{Key: "movie", Val: "Star Wars"})
	bo.Total.Store("$25,539,306")
	return bo
}

func StoreAndLoad() {
	bo := NewBoxOffice()

	value := bo.Movies.Load().(*Value)
	total := bo.Total.Load().(string)

	fmt.Printf("Movies %v domestic total as of Aug. 27, 2017: %v \n", value.Val, total)
}
