package v2

import "testing"

func TestConstructor(t *testing.T) {
	t.Skip()
	c := New(3)
	c.Set("one", "apple")
	v, ok := c.Get("one")
	if !ok {
		t.Fatal("value should exist")
	}
	if v != "apple" {
		t.Errorf("got %q, want %q", v, "apple")
	}
}

func TestMultiGetSet(t *testing.T) {
	c := New(2)
	if _, ok := c.Get("one"); ok {
		t.Errorf("should not be able to get a value when not set")
	}
	t.Log("## ll empty: ", c.linkedList)
	c.Set("one", "apple")
	t.Log("## ll set(one): ", c.linkedList)

	c.Set("two", "banana")
	t.Log("## ll set(two): ", c.linkedList)

	c.Set("three", "cucumber")
	t.Log("## ll set(three): ", c.linkedList)

	if _, ok := c.Get("one"); ok {
		t.Errorf("item should have been evicted")
	}
	t.Log("## ll get(one): ", c.linkedList)

	v, ok := c.Get("three")
	if v != "cucumber" {
		t.Errorf("got %q, want %q for entry %q", v, "cucumber", "three")
	}
	t.Log("## ll get(three): ", c.linkedList)

	v, ok = c.Get("two")
	if v != "banana" {
		t.Errorf("got %q, want %q for entry %q", v, "banana", "two")
	}
	t.Log("## ll get(two): ", c.linkedList)

	// evict the oldest, "three"
	c.Set("four", "dragonfruit")
	t.Log("## ll set(four): ", c.linkedList)

	if _, ok = c.Get("three"); ok {
		t.Errorf("%q should have evicted", "three")
	}
}
