package cache

import (
	"sync"
)

const (
	// linkedHead and linkedTail represents the head
	// and tail of a bidirectional linked list without
	// occupying the linked list space.
	linkedHead = "linkedHead"
	linkedTail = "linkedTail"
)

// LinkedNode represents a node in a bidirectional linked list.
type LinkedNode struct {
	key   string
	value interface{}
	pre   *LinkedNode
	next  *LinkedNode
}

// NewLinkedNode returns a new LinkedNode.
func NewLinkedNode(k string, v interface{}) *LinkedNode {
	return &LinkedNode{
		key:   k,
		value: v,
	}
}

// LRUCache represents a LRU cache implemented with hash
// tables and bidirectional linked list.
type LRUCache struct {
	count int
	size  int
	cache map[string]*LinkedNode
	lock  sync.Mutex
}

// NewLRUCache returns a LRUCache with a specified capacity.
func NewLRUCache(size int) *LRUCache {
	head := NewLinkedNode(linkedHead, nil)
	head.pre = nil

	tail := NewLinkedNode(linkedTail, nil)
	tail.next = nil

	head.next = tail
	tail.pre = head

	lc := &LRUCache{
		count: 0,
		size:  size,
		cache: make(map[string]*LinkedNode),
	}

	lc.cache[linkedHead] = head
	lc.cache[linkedTail] = tail

	return lc
}

func (lc *LRUCache) Set(k string, v interface{}) bool {
	lc.lock.Lock()
	defer lc.lock.Unlock()

	if k == linkedHead || k == linkedTail {
		return false
	}

	lhead := lc.cache[linkedHead]
	ltail := lc.cache[linkedTail]
	head := NewLinkedNode(k, v)

	// Special handling is required when there are no
	// nodes in the linked list.
	if lc.count == 0 {
		lhead.next = head
		head.pre = lhead
		head.next = ltail
		ltail.pre = head

		lc.cache[k] = head
		lc.count++
		return true
	}

	// When the cache is full, delete the tail node and
	// put the new node in the head.
	if lc.count == lc.size {
		nodeRemove := ltail.pre
		nodeRemove.pre.next = ltail
		ltail.pre = nodeRemove.pre
		nodeRemove = nil

		lc.count--
	}

	// If the key you want to set exists, remove it from its
	// original location and put it in the head of the linked list.
	if originalNode, exist := lc.cache[k]; exist {
		originalPre := originalNode.pre
		originalPre.next = originalNode.next
		originalNode.next.pre = originalPre
		originalNode = nil
	}

	oldHead := lhead.next
	lhead.next = head
	head.pre = lhead
	head.next = oldHead.next
	oldHead.next.pre = head
	oldHead = nil

	lc.cache[k] = head
	lc.count++

	return true
}

// Get returns the specified value in a bidirectional linked list.
func (lc *LRUCache) Get(k string) (interface{}, bool) {
	lc.lock.Lock()
	defer lc.lock.Unlock()

	if lc.count == 0 || k == linkedHead || k == linkedTail {
		return nil, false
	}

	dstNode, ok := lc.cache[k]
	if !ok {
		return nil, false
	}

	v := dstNode.value
	lhead := lc.cache[linkedHead]

	// If the target value is in the head of the list,
	// return it directly.
	if dstNode.pre == lhead {
		return v, true
	}

	// Delete the target node from its original location.
	preNode := dstNode.pre
	preNode.next = dstNode.next
	dstNode.next.pre = preNode

	// Move the target node to the head of the linked list.
	oldHead := lhead.next
	lhead.next = dstNode
	dstNode.pre = lhead
	dstNode.next = oldHead
	oldHead.pre = dstNode

	return v, true
}

// GetCount returns the number of nodes in the cache.
func (lc *LRUCache) GetCount() int {
	lc.lock.Lock()
	defer lc.lock.Unlock()
	return lc.count
}

// GetSize returns the size of the cache.
func (lc *LRUCache) GetSize() int {
	return lc.size
}
