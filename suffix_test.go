package suffix

import "testing"

func TestSuffixTree(t *testing.T) {
	content := "abcabxabcd"
	tree := NewSuffixTree()
	tree.BuildFromStr(content)

	count := tree.CountStr("a")
	if count != 3 {
		t.Errorf("a count should be 3, but %d", count)
	}
	count = tree.CountStr("bc")
	if count != 2 {
		t.Errorf("bc count should be 2, but %d", count)
	}
	count = tree.CountStr("ac")
	if count != 0 {
		t.Errorf("ac count should be 0, but %d", count)
	}
	if tree.ExistStr("ae") {
		t.Errorf("ae doesn't exist!")
	}

}
