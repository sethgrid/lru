package linkedlist

import "testing"

func TestLinkedList(t *testing.T) {
	ll := New()
	if got, want := ll.String(), `[]`; got != want {
		t.Errorf("empty list: got %q, want %q", got, want)
	}
	node := ll.First()

	node = node.InsertNext("a", "apple")
	if got, want := ll.String(), `[] :: a[apple]`; got != want {
		t.Errorf("one item: got %q, want %q", got, want)
	}

	node = node.InsertNext("b", "banana")
	if got, want := ll.String(), `[] :: a[apple] :: b[banana]`; got != want {
		t.Errorf("two items: got %q, want %q", got, want)
	}

	node = node.InsertNext("c", "cucumber")
	if got, want := ll.String(), `[] :: a[apple] :: b[banana] :: c[cucumber]`; got != want {
		t.Errorf("three items: got %q, want %q", got, want)
	}

	if got, want := ll.First().IsRoot(), true; got != want {
		t.Errorf("got %v, want %v for first data node isRoot", got, want)
	}

	if got, want := ll.Last().Key(), "c"; got != want {
		t.Errorf("got %q, want %q for first data node", got, want)
	}

	// delete the middle node
	node = ll.First()  // root, empty
	node = node.Next() // a
	node = node.Next() // b

	if got, want := node.Key(), "b"; got != want {
		t.Errorf("got %q, want %q for second node", got, want)
	}

	node.Delete()

	node = ll.First()  // root, empty
	node = node.Next() // a
	node = node.Next() // c

	if got, want := node.Key(), "c"; got != want {
		t.Errorf("got %q, want %q for second node", got, want)
	}

	if got, want := ll.String(), "[] :: a[apple] :: c[cucumber]"; got != want {
		t.Errorf("post delete: got %q, want %q", got, want)
	}

	// delete the last node
	node = ll.First()  // root, empty
	node = node.Next() // a
	node = node.Next() // c
	node.Delete()

	if got, want := ll.String(), "[] :: a[apple]"; got != want {
		t.Errorf("post delete: got %q, want %q", got, want)
	}
}
