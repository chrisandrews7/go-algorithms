package main

import (
	"testing"
)

func TestHashMap(t *testing.T) {
	primeNumber := 17
	hashmap := NewHashMap(primeNumber)

	hashmap.Set("fish", "chips")
	hashmap.Set("mushy", "pea")
	hashmap.Set("mushy", "peas") // Test overriding
	hashmap.Set("salt", "vinegar")
	hashmap.Set("salt2", "pepper") // Same key hash - Testing collisions

	if hashmap.Get("fish") != "chips" {
		t.Error("Expected the value chips")
	}
	if hashmap.Get("mushy") != "peas" {
		t.Error("Expected the value peas")
	}
	if hashmap.Get("salt") != "vinegar" {
		t.Error("Expected the value vinegar")
	}
	if hashmap.Get("salt2") != "pepper" {
		t.Error("Expected the value pepper")
	}
}
