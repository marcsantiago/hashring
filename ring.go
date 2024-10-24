package go_hashring

import (
	"hash/crc32"
	"sort"
	"strconv"
	"unsafe"
)

type hash func(data []byte) uint32

type HashRing struct {
	hash         hash
	replicaCount int
	keys         []int
	hashMap      map[int]string
}

func New(nodes []string) *HashRing {
	m := &HashRing{
		replicaCount: len(nodes),
		hash:         crc32.ChecksumIEEE,
		hashMap:      make(map[int]string),
	}
	m.AddNode(nodes...)
	return m
}

func (h *HashRing) AddNode(keys ...string) {
	h.replicaCount += len(keys)
	for i, key := range keys {
		hashValue := int(h.hash(unsafeGetBytes(strconv.Itoa(i) + key)))
		h.keys = append(h.keys, hashValue)
		h.hashMap[hashValue] = key
	}
	sort.Ints(h.keys)
}

func (h *HashRing) GetTargetNode(key string) (string, error) {
	if len(h.keys) == 0 {
		return "", ErrNoNodes
	}
	hashValue := int(h.hash(unsafeGetBytes(key)))
	idx := sort.Search(len(h.keys), func(i int) bool { return h.keys[i] >= hashValue })
	if idx == len(h.keys) {
		idx = 0
	}
	return h.hashMap[h.keys[idx]], nil
}

func unsafeGetBytes(s string) (b []byte) {
	if s == "" {
		return nil
	}
	return unsafe.Slice(unsafe.StringData(s), len(s))
}
