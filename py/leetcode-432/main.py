class AllOne:
    # every operation should be O(1) average time
    # since we only ever increase or decrease the count of a key by 1
    # the operations in a heap are bounded

    def __init__(self):
        self.minHeap = Heap() 
        self.maxHeap = Heap(min_heap=False) 

    def inc(self, key: str) -> None:
        self.minHeap.update(key, 1)
        self.maxHeap.update(key, 1)

    def dec(self, key: str) -> None:
        self.minHeap.update(key, -1)
        self.maxHeap.update(key, -1)
        if self.minHeap.get_value(key) == 0:
            self.minHeap.remove(key)
            self.maxHeap.remove(key)

    def getMaxKey(self) -> str:
        x = self.maxHeap.first() 
        if x is None:
            return ""
        else:
            return x[0]

    def getMinKey(self) -> str:
        x = self.minHeap.first() 
        if x is None:
            return ""
        else:
            return x[0]

class Heap:
    def __init__(self, min_heap=True) -> None:
        self.key_to_index = {}
        self.size = 0
        self.heap = []
        self.min_heap = min_heap

    def first(self):
        if self.size == 0:
            return None
        else:
            return self.heap[0]

    def update(self, key, value):
        if key in self.key_to_index:
            i = self.key_to_index[key]
            e = self.heap[i]
            self.heap[i] = (key, e[1]+value)
            self._siftUp(i)
            self._siftDown(i)
        else:
            self._push(key, value)

    def get_value(self, key):
        if key in self.key_to_index:
            return self.heap[self.key_to_index[key]][1]
        else:
            return None

    def remove(self, key):
        if key not in self.key_to_index:
            return
        
        i = self.key_to_index[key]
        if self.size == 1:
            self.heap = []
            self.size = 0
            del self.key_to_index[key]
        elif i == self.size-1:
            self.heap = self.heap[:self.size-1]
            self.size -= 1
            del self.key_to_index[key]
        else:
            self._swap(i, self.size-1)
            self.heap = self.heap[:self.size-1]
            self.size -= 1
            del self.key_to_index[key]
            self._siftUp(i)
            self._siftDown(i)

    def _push(self, key, value):
        self.heap.append((key, value))
        self.key_to_index[key] = self.size
        self.size += 1
        self._siftUp(self.size-1)

    def _siftUp(self, i):
        while i > 0:
            parent = (i-1)//2
            if self._less(i, parent):
                self._swap(i, parent)
                i = parent
            else:
                break
    
    def _siftDown(self, i):
        while 2*i+1 < self.size:
            minChild = 2*i+1
            if 2*i+2 < self.size and self._less(2*i+2, minChild):
                minChild = 2*i+2
            if self._less(minChild, i):
                self._swap(i, minChild)
                i = minChild
            else:
                break

    def _swap(self, i, j):
        key1, key2 = self.heap[i][0], self.heap[j][0]
        self.key_to_index[key1], self.key_to_index[key2] = j, i
        self.heap[i], self.heap[j] = self.heap[j], self.heap[i]

    def _less(self, i, j):
        if self.min_heap:
            return self.heap[i][1] < self.heap[j][1]
        else:
            return self.heap[i][1] > self.heap[j][1]
