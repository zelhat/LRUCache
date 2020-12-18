package cache

import (
	"testing"
)

func getItems() []item {
	return []item{
		{key: "1", value: "123"},
		{key: "2", value: "234"},
		{key: "3", value: "345"},
		{key: "4", value: "456"},
		{key: "5", value: "567"},
		{key: "6", value: "678"},
		{key: "7", value: "789"},
	}
}

func TestLRUcacheAdd(t *testing.T) {
	cache := NewLRUCache(5)
	items := getItems()
	for _, item := range items {
		ok := cache.Add(item.key, item.value)
		if !ok {
			t.Fail()
		}
	}

	if cache.Add("7", "111") != false {
		t.Fail()
	}
}

func TestLRUcacheGet(t *testing.T) {
	cache := NewLRUCache(5)
	items := getItems()
	for _, item := range items {
		cache.Add(item.key, item.value)
	}

	for i := 2; i < len(items); i++ {
		val, ok := cache.Get(items[i].key)
		if val != items[i].value || !ok {
			t.Fail()
		}
	}

	_, ok := cache.Get("8")
	if ok {
		t.Fail()
	}
}

func TestLRUcacheRemove(t *testing.T) {
	cache := NewLRUCache(5)
	cache.Add("key", "value")
	if cache.Remove("key") != true {
		t.Fail()
	}
	if cache.Remove("key") != false {
		t.Fail()
	}
}
