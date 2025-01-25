package main

import (
	"fmt"
	"internal/pokecache"
	"testing"
	"time"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "",
			expected: []string{},
		},
		{
			input:    "SingleWord",
			expected: []string{"singleword"},
		},
		{
			input:    "A SentANCE WITH mixed CASE",
			expected: []string{"a", "sentance", "with", "mixed", "case"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Len of actual doesn't match Expected")
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("%v does not match %v", word, expectedWord)
			}
		}
	}
}

func TestCacheAddGet(t *testing.T) {
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://example.com",
			val: []byte("fakedata"),
		},
		{
			key: "https://example.com/path",
			val: []byte("extrafakedata"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case: %v", i), func(t *testing.T) {
			cache := pokecache.NewCache(5 * time.Second)
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("couldn't get newly added item")
				return
			}
			if string(val) != string(c.val) {
				t.Errorf("Value returned with Get doesn't match what was added")
				return
			}
		})
	}
}

func TestReapLoop(t *testing.T) {

	key := "https://example.com"
	val := []byte("fakedata")

	cache := pokecache.NewCache(5 * time.Millisecond)
	cache.Add(key, val)

	if _, ok := cache.Get(key); !ok {
		t.Errorf("%v :not found when should have", key)
		return
	}

	time.Sleep(10 * time.Millisecond)

	if _, ok := cache.Get(key); ok {
		t.Errorf("%v :found when shouldn't have", key)
		return
	}

}
