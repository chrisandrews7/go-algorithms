package main

import "math"

const (
	MAX_HASH_ITERATIONS = 100.00
	PRIME_MULTIPLIER    = 8192
)

type hashNode []string

type HashMap struct {
	values [][]hashNode
	size   int
}

func (hashmap *HashMap) hash(key string) (hashValue int) {
	hashLoopIterations := math.Min(float64(len(key)), MAX_HASH_ITERATIONS)

	for i := 0; i < int(hashLoopIterations); i++ {
		hashValue = (hashValue*PRIME_MULTIPLIER + int(key[i])) % hashmap.size
	}

	return hashValue
}

func (hashmap *HashMap) getNestedIndex(topIndex int, key string) int {
	for index, value := range hashmap.values[topIndex] {
		if value[0] == key {
			return index
		}
	}

	return -1
}

func (hashmap *HashMap) Set(key string, value string) {
	index := hashmap.hash(key)
	node := hashNode{key, value}

	if nestedIndex := hashmap.getNestedIndex(index, key); nestedIndex >= 0 {
		hashmap.values[index][nestedIndex][1] = value
		return
	}

	hashmap.values[index] = append(hashmap.values[index], node)
}

func (hashmap *HashMap) Get(key string) string {
	index := hashmap.hash(key)

	if nestedIndex := hashmap.getNestedIndex(index, key); nestedIndex >= 0 {
		return hashmap.values[index][nestedIndex][1]
	}

	return ""
}

func NewHashMap(size int) HashMap {
	return HashMap{
		size:   size,
		values: make([][]hashNode, size),
	}
}
