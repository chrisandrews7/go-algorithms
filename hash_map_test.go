package main

import (
	"testing"
)

func TestHashMap(t *testing.T) {
	primeNumber := 17
	hashmap := NewHashMap(primeNumber)

	hashmap.Set("fish", "chips")
	hashmap.Set("mushy", "peas")
	hashmap.Set("salt", "vinegar")

	if hashmap.Get("fish") != "chips" {
		t.Error("Expected the value chips")
	}
	if hashmap.Get("mushy") != "peas" {
		t.Error("Expected the value peas")
	}
	if hashmap.Get("salt") != "vinegar" {
		t.Error("Expected the value vinegar")
	}
}

func TestHashMapOverride(t *testing.T) {
	primeNumber := 17
	hashmap := NewHashMap(primeNumber)

	hashmap.Set("mushy", "pea")
	hashmap.Set("mushy", "peas")

	if hashmap.Get("mushy") != "peas" {
		t.Error("Expected the value peas")
	}
}

func TestHashMapCollisions(t *testing.T) {
	primeNumber := 17
	hashmap := NewHashMap(primeNumber)

	hashmap.Set("salt", "vinegar")
	hashmap.Set("salt2", "pepper")

	if hashmap.Get("salt") != "vinegar" {
		t.Error("Expected the value vinegar")
	}
	if hashmap.Get("salt2") != "pepper" {
		t.Error("Expected the value pepper")
	}
}
