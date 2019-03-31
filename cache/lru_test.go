package cache

import (
	"reflect"
	"testing"
)

func TestLRUCache(t *testing.T) {
	lru := NewLRUCache(3)

	// case 1: now cache is ["value1"]
	lru.Set("k1", "value1")
	v1, exist := lru.Get("k1")
	if !exist {
		t.Errorf("failed to set value, v: %v\n", v1)
	} else if value, ok := v1.(string); !ok {
		t.Errorf("got a value type error, get: %v, want: %v\n",
			reflect.TypeOf(v1), reflect.TypeOf("s"))
	} else {
		t.Logf("case 1 success, get value: %v\n", value)
	}

	// case 2: now cache is ["value3", "value2", "value1"]
	lru.Set("k2", "value2")
	lru.Set("k3", "value3")
	t.Logf("case 2 success, now cache is full, size: %v, count: %v\n",
		lru.GetSize(), lru.GetCount())

	// case 3: now cache is still ["value3", "value2", "value1"]
	v3, exists := lru.Get("k3")
	if !exists {
		t.Errorf("get value error, want: %v, get: %v\n", "value3", v3)
	}
	headValue := lru.GetHeadValue()
	if headValue == nil {
		t.Error("there is no data in the cache")
	}
	if v3 != headValue {
		t.Errorf("cache logic error, want: %v, get: %v\n", headValue, v3)
	}
	t.Logf("case 3 success, want: %v, get: %v\n", headValue, v3)

	// case 4: now cache should be ["newValue2", "value3", "value1"]
	const newValue = "newValue2"
	lru.Set("k2", newValue)
	headValue = lru.GetHeadValue()
	if headValue == nil {
		t.Error("there is no data in the cache")
	}
	if headValue != newValue {
		t.Errorf("cache logic error, want: %v, get: %v\n", newValue, headValue)
	}
	t.Logf("case 4 success, want: %v, get: %v\n", newValue, headValue)

	// case 5: now cache should be ["value4", "newValue2", "value3"], "value1" is ejected.
	const value4 = "value4"
	lru.Set("k4", value4)
	v4 := lru.GetHeadValue()
	if v4 == nil {
		t.Error("there is no data in the cache")
	}
	if v4 != value4 {
		t.Errorf("cache logic error, want: %v, get: %v\n", value4, v4)
	}
	t.Logf("case 5 success, want: %v, get: %v\n", value4, v4)
}
