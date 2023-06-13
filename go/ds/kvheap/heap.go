package kvheap

type kv[K any, V any] struct {
	key   K
	value V
}

type KeyValueHeap[K comparable, V any] struct {
	h          []kv[K, V]
	size       int
	keyToIndex map[K]int
	lessFn     func(V, V) bool
}

func New[K comparable, V any](less func(V, V) bool) *KeyValueHeap[K, V] {
	return &KeyValueHeap[K, V]{
		h:          nil,
		size:       0,
		keyToIndex: make(map[K]int),
		lessFn:     less,
	}
}

func (kvh *KeyValueHeap[K, V]) Size() int {
	return kvh.size
}

func (kvh *KeyValueHeap[K, V]) Update(key K, value V) {
	if i, ok := kvh.keyToIndex[key]; ok {
		kvh.h[i].value = value
		kvh.siftUp(i)
		kvh.siftDown(i)
	} else {
		kvh.h = append(kvh.h, kv[K, V]{key: key, value: value})
		kvh.size += 1
		kvh.keyToIndex[key] = kvh.size - 1
		kvh.siftUp(kvh.size - 1)
	}
}

func (kvh *KeyValueHeap[K, V]) First() (K, V, bool) {
	if kvh.size == 0 {
		var emptyKey K
		var emptyValue V
		return emptyKey, emptyValue, false
	}

	m := kvh.h[0]
	return m.key, m.value, true
}

func (kvh *KeyValueHeap[K, V]) Pop() (K, V, bool) {
	if kvh.size == 0 {
		var emptyKey K
		var emptyValue V
		return emptyKey, emptyValue, false
	}

	min := kvh.h[0]

	if kvh.size == 1 {
		kvh.h = nil
		kvh.size = 0
		delete(kvh.keyToIndex, min.key)
	} else {
		kvh.swap(0, kvh.size-1)
		kvh.size -= 1
		kvh.h = kvh.h[0:kvh.size]
		delete(kvh.keyToIndex, min.key)
		kvh.siftDown(0)
	}

	return min.key, min.value, true
}

func (kvh *KeyValueHeap[K, V]) RemoveKey(k K) {
	i, ok := kvh.keyToIndex[k]
	if !ok {
		return
	}

	if kvh.size == 1 {
		kvh.size = 0
		kvh.h = nil
		delete(kvh.keyToIndex, k)
	} else if i == kvh.size-1 {
		kvh.size -= 1
		kvh.h = kvh.h[0:kvh.size]
		delete(kvh.keyToIndex, k)
	} else {
		kvh.swap(i, kvh.size-1)
		kvh.size -= 1
		kvh.h = kvh.h[0:kvh.size]
		delete(kvh.keyToIndex, k)
		kvh.siftDown(i)
	}
}

func (kvh *KeyValueHeap[K, V]) swap(i, j int) {
	a, b := kvh.h[i], kvh.h[j]
	kvh.h[i], kvh.h[j] = b, a
	kvh.setIndex(a.key, j)
	kvh.setIndex(b.key, i)
}

func (kvh *KeyValueHeap[K, V]) siftDown(i int) {
	for 2*i+1 < kvh.size {
		minChild := 2*i + 1
		if 2*i+2 < kvh.size && kvh.less(2*i+2, 2*i+1) {
			minChild = 2*i + 2
		}
		if kvh.less(minChild, i) {
			kvh.swap(i, minChild)
			i = minChild
		} else {
			break
		}
	}
}

func (kvh *KeyValueHeap[K, V]) siftUp(i int) {
	for i > 0 {
		parent := (i - 1) / 2
		if kvh.less(i, parent) {
			kvh.swap(i, parent)
			i = parent
		} else {
			break
		}
	}
}

func (kvh *KeyValueHeap[K, V]) less(i, j int) bool {
	return kvh.lessFn(kvh.h[i].value, kvh.h[j].value)
}

func (kvh *KeyValueHeap[K, V]) setIndex(key K, i int) {
	kvh.keyToIndex[key] = i
}

func (kvh *KeyValueHeap[K, V]) getIndex(key K) (int, bool) {
	i, ok := kvh.keyToIndex[key]
	return i, ok
}
