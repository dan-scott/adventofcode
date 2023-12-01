package day01

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTrie_SingleMatch(t *testing.T) {
	trie := NewTrie(ti([]rune("one"), 1))

	val, match := trie.Test([]rune("onel"))
	assert.True(t, match)
	assert.Equal(t, 1, val)
}

func TestTrie_MultiplePossibilities(t *testing.T) {
	trie := NewTrie[int32, int](ti([]rune("one"), 1), ti([]rune("two"), 2))
	val, match := trie.Test([]rune("twoone"))
	assert.True(t, match)
	assert.Equal(t, 2, val)
}
