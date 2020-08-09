package main

import (
	"fmt"
	"hash/crc32"
	"sort"
	"strconv"
	"sync"
)

func main() {
	h := &HashRing{}
	node := []string{"node1","node2"}
	h.New(node)
	fmt.Println(h.Keys.Len())
	fmt.Println(h.GetNode("no"))
	fmt.Println(h.Nodes)
}

const DEFAULT_REPLICAS = 12

type SortKeys []uint32

func (sk SortKeys) Len() int {
	return len(sk)
}

func (sk SortKeys) Less(i, j int) bool {
	return sk[i] < sk[j]
}

func (sk SortKeys) Swap(i, j int) {
	sk[i], sk[j] = sk[j], sk[i]
}

type HashRing struct {
	Nodes map[uint32]string
	Keys  SortKeys
	sync.RWMutex
}

func (hr *HashRing) New(nodes []string) {
	if nodes == nil {
		return
	}

	hr.Nodes = make(map[uint32]string)
	hr.Keys = SortKeys{}
	for _, node := range (nodes) {
		for i := 0; i < DEFAULT_REPLICAS; i++ {
			str := node + strconv.Itoa(i)
			hr.Nodes[hr.hashStr(str)] = node
			hr.Keys = append(hr.Keys, hr.hashStr(str))
		}
	}
	sort.Sort(hr.Keys)
}

func (hr *HashRing) hashStr(key string) uint32 {
	return crc32.ChecksumIEEE([]byte(key))
}

func (hr *HashRing) GetNode(key string) string {
	hr.RLock()
	defer hr.RUnlock()
	hash := hr.hashStr(key)
	i := hr.get_position(hash)
	return hr.Nodes[hr.Keys[i]]
}

func (hr *HashRing) get_position(hash uint32) int {
	i := sort.Search(len(hr.Keys), func(i int) bool {
		return hr.Keys[i] >= hash
	})

	if i < len(hr.Keys) {
		if i == len(hr.Keys)-1 {
			return 0
		} else {
			return i
		}
	} else {
		return len(hr.Keys) - 1
	}
}
