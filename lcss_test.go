package lcss

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCharNodeAdd(t *testing.T) {
	node := &charNode{}
	node.Add([]byte("abc"))
	assert.Equal(t, byte(0), node.char)
	assert.Equal(t, 0, node.used)
	assert.Len(t, node.children, 1)
	assert.Equal(t, byte('a'), node.children[0].char)
	assert.Equal(t, 1, node.children[0].used)
	assert.Len(t, node.children[0].children, 1)
	assert.Equal(t, byte('b'), node.children[0].children[0].char)
	assert.Equal(t, 1, node.children[0].children[0].used)
	assert.Len(t, node.children[0].children[0].children, 1)
	assert.Equal(t, byte('c'), node.children[0].children[0].children[0].char)
	assert.Equal(t, 1, node.children[0].children[0].children[0].used)
	assert.Len(t, node.children[0].children[0].children[0].children, 0)
	node.Add([]byte{})
	assert.Equal(t, byte(0), node.char)
	assert.Equal(t, 0, node.used)
	assert.Len(t, node.children, 1)
	node.Add([]byte("abd"))
	assert.Equal(t, byte(0), node.char)
	assert.Equal(t, 0, node.used)
	assert.Len(t, node.children, 1)
	assert.Equal(t, byte('a'), node.children[0].char)
	assert.Equal(t, 2, node.children[0].used)
	assert.Len(t, node.children[0].children, 1)
	assert.Equal(t, byte('b'), node.children[0].children[0].char)
	assert.Equal(t, 2, node.children[0].children[0].used)
	assert.Len(t, node.children[0].children[0].children, 2)
	assert.Equal(t, byte('c'), node.children[0].children[0].children[0].char)
	assert.Equal(t, 1, node.children[0].children[0].children[0].used)
	assert.Len(t, node.children[0].children[0].children[0].children, 0)
	assert.Equal(t, byte('d'), node.children[0].children[0].children[1].char)
	assert.Equal(t, 1, node.children[0].children[0].children[1].used)
	assert.Len(t, node.children[0].children[0].children[1].children, 0)
	node.Add([]byte("abc"))
	assert.Equal(t, byte(0), node.char)
	assert.Equal(t, 0, node.used)
	assert.Len(t, node.children, 1)
	assert.Equal(t, byte('a'), node.children[0].char)
	assert.Equal(t, 3, node.children[0].used)
	assert.Len(t, node.children[0].children, 1)
	assert.Equal(t, byte('b'), node.children[0].children[0].char)
	assert.Equal(t, 3, node.children[0].children[0].used)
	assert.Len(t, node.children[0].children[0].children, 2)
	assert.Equal(t, byte('c'), node.children[0].children[0].children[0].char)
	assert.Equal(t, 2, node.children[0].children[0].children[0].used)
	assert.Len(t, node.children[0].children[0].children[0].children, 0)
	assert.Equal(t, byte('d'), node.children[0].children[0].children[1].char)
	assert.Equal(t, 1, node.children[0].children[0].children[1].used)
	assert.Len(t, node.children[0].children[0].children[1].children, 0)
}

func TestCharNodeRemove(t *testing.T) {
	node := &charNode{}
	node.Add([]byte("abc"))
	node.Remove([]byte("abc"))
	assert.Equal(t, byte(0), node.char)
	assert.Equal(t, 0, node.used)
	assert.Len(t, node.children, 0)
	node.Add([]byte("abc"))
	node.Add([]byte("abd"))
	node.Remove([]byte("abc"))
	assert.Equal(t, byte(0), node.char)
	assert.Equal(t, 0, node.used)
	assert.Len(t, node.children, 1)
	assert.Equal(t, byte('a'), node.children[0].char)
	assert.Equal(t, 1, node.children[0].used)
	assert.Len(t, node.children[0].children, 1)
	assert.Equal(t, byte('b'), node.children[0].children[0].char)
	assert.Equal(t, 1, node.children[0].children[0].used)
	assert.Len(t, node.children[0].children[0].children, 1)
	assert.Equal(t, byte('d'), node.children[0].children[0].children[0].char)
	assert.Equal(t, 1, node.children[0].children[0].children[0].used)
	assert.Len(t, node.children[0].children[0].children[0].children, 0)
	node.Remove([]byte{})
	assert.Equal(t, byte(0), node.char)
	assert.Equal(t, 0, node.used)
	assert.Len(t, node.children, 1)
	assert.Equal(t, byte('a'), node.children[0].char)
	assert.Equal(t, 1, node.children[0].used)
	assert.Len(t, node.children[0].children, 1)
	assert.Equal(t, byte('b'), node.children[0].children[0].char)
	assert.Equal(t, 1, node.children[0].children[0].used)
	assert.Len(t, node.children[0].children[0].children, 1)
	assert.Equal(t, byte('d'), node.children[0].children[0].children[0].char)
	assert.Equal(t, 1, node.children[0].children[0].children[0].used)
	assert.Len(t, node.children[0].children[0].children[0].children, 0)
	node.Add([]byte("ab"))
	node.Remove([]byte("ab"))
	assert.Equal(t, byte(0), node.char)
	assert.Equal(t, 0, node.used)
	assert.Len(t, node.children, 1)
	assert.Equal(t, byte('a'), node.children[0].char)
	assert.Equal(t, 1, node.children[0].used)
	assert.Len(t, node.children[0].children, 1)
	assert.Equal(t, byte('b'), node.children[0].children[0].char)
	assert.Equal(t, 1, node.children[0].children[0].used)
	assert.Len(t, node.children[0].children[0].children, 1)
	assert.Equal(t, byte('d'), node.children[0].children[0].children[0].char)
	assert.Equal(t, 1, node.children[0].children[0].children[0].used)
	assert.Len(t, node.children[0].children[0].children[0].children, 0)
	node.Add([]byte("ab"))
	node.Remove([]byte("abd"))
	assert.Equal(t, byte(0), node.char)
	assert.Equal(t, 0, node.used)
	assert.Len(t, node.children, 1)
	assert.Equal(t, byte('a'), node.children[0].char)
	assert.Equal(t, 1, node.children[0].used)
	assert.Len(t, node.children[0].children, 1)
	assert.Equal(t, byte('b'), node.children[0].children[0].char)
	assert.Equal(t, 1, node.children[0].children[0].used)
	assert.Len(t, node.children[0].children[0].children, 0)
}

func TestCharNodeLongestCommonPrefixLength(t *testing.T) {
	node := &charNode{}
	assert.Equal(t, 0, node.LongestCommonPrefixLength())
	node.Add([]byte("abc"))
	assert.Equal(t, 3, node.LongestCommonPrefixLength())
	node.Add([]byte("abd"))
	assert.Equal(t, 2, node.LongestCommonPrefixLength())
	node.Remove([]byte("abd"))
	assert.Equal(t, 3, node.LongestCommonPrefixLength())
	node.Add([]byte("ab"))
	assert.Equal(t, 2, node.LongestCommonPrefixLength())
}

func TestCharNodeLongestCommonPrefix(t *testing.T) {
	node := &charNode{}
	assert.Equal(t, []byte{}, node.LongestCommonPrefix())
	node.Add([]byte("abc"))
	assert.Equal(t, []byte("abc"), node.LongestCommonPrefix())
	node.Add([]byte("abd"))
	assert.Equal(t, []byte("ab"), node.LongestCommonPrefix())
	node.Remove([]byte("abd"))
	assert.Equal(t, []byte("abc"), node.LongestCommonPrefix())
	node.Add([]byte("ab"))
	assert.Equal(t, []byte("ab"), node.LongestCommonPrefix())
}

func TestCharNodeBug1(t *testing.T) {
	node := &charNode{}
	node.Add([]byte("a"))
	node.Add([]byte("a"))
	node.Remove([]byte("a"))
	node.Add([]byte("abbara"))
	node.Add([]byte("abr"))

	assert.Equal(t, byte(0), node.char)
	assert.Equal(t, 0, node.used)
	assert.Len(t, node.children, 1)
	assert.Equal(t, byte('a'), node.children[0].char)
	assert.Equal(t, 3, node.children[0].used)
	assert.Len(t, node.children[0].children, 1)
	assert.Equal(t, byte('b'), node.children[0].children[0].char)
	assert.Equal(t, 2, node.children[0].children[0].used)
	assert.Len(t, node.children[0].children[0].children, 2)
	assert.Equal(t, byte('b'), node.children[0].children[0].children[0].char)
	assert.Equal(t, 1, node.children[0].children[0].children[0].used)
	assert.Len(t, node.children[0].children[0].children[0].children, 1)
	assert.Equal(t, byte('r'), node.children[0].children[0].children[1].char)
	assert.Equal(t, 1, node.children[0].children[0].children[1].used)
	assert.Len(t, node.children[0].children[0].children[1].children, 0)
	assert.Equal(t, 1, node.LongestCommonPrefixLength())
	assert.Equal(t, []byte("a"), node.LongestCommonPrefix())
}

func TestCharNodeBug2(t *testing.T) {
	node := &charNode{}
	node.Add([]byte("habrahabr"))
	node.Add([]byte("bbara"))
	node.Add([]byte("mraja"))
	node.Remove([]byte("habrahabr"))
	node.Add([]byte("r"))
	node.Remove([]byte("bbara"))
	node.Add([]byte("ra"))
	node.Remove([]byte("r"))
	node.Add([]byte("rahabr"))
	node.Remove([]byte("bbara"))
	node.Add([]byte("ra"))
	node.Remove([]byte("mraja"))
	node.Add([]byte("raja"))

	assert.Equal(t, byte(0), node.char)
	assert.Equal(t, 0, node.used)
	assert.Len(t, node.children, 1)
	assert.Equal(t, byte('r'), node.children[0].char)
	assert.Equal(t, 3, node.children[0].used)
	assert.Len(t, node.children[0].children, 1)
	assert.Equal(t, byte('a'), node.children[0].children[0].char)
	assert.Equal(t, 3, node.children[0].children[0].used)
	assert.Len(t, node.children[0].children[0].children, 2)
	assert.Equal(t, byte('h'), node.children[0].children[0].children[0].char)
	assert.Equal(t, 1, node.children[0].children[0].children[0].used)
	assert.Len(t, node.children[0].children[0].children[0].children, 1)
	assert.Equal(t, byte('j'), node.children[0].children[0].children[1].char)
	assert.Equal(t, 1, node.children[0].children[0].children[1].used)
	assert.Len(t, node.children[0].children[0].children[1].children, 1)
	assert.Equal(t, []byte("ra"), node.LongestCommonPrefix())
}

func TestLongestCommonSubstring(t *testing.T) {
	assert.Equal(t, []byte{}, LongestCommonSubstring())
	assert.Equal(t, []byte("abc"), LongestCommonSubstring([]byte("abc")))
	assert.Equal(t, []byte{}, LongestCommonSubstring([]byte("abc"), []byte{}))
	assert.Equal(t, []byte("ab"), LongestCommonSubstring([]byte("abc"), []byte("abd")))
	assert.Equal(t, []byte("bc"), LongestCommonSubstring([]byte("abc"), []byte("dbc")))
	assert.Equal(t, []byte("ab"), LongestCommonSubstring([]byte("ab"), []byte("abd")))
	assert.Equal(t, []byte("ABC"), LongestCommonSubstring(
		[]byte("ABABC"), []byte("BABCA"), []byte("ABCBA")))
	assert.Equal(t, []byte("ra"), LongestCommonSubstring(
		[]byte("habrahabr"),
		[]byte("abbara"),
		[]byte("humraja")))
	assert.Equal(t, []byte("abcdez"), LongestCommonSubstring(
		[]byte("zxabcdezy"),
		[]byte("yzabcdezx"),
		[]byte("abcdez"),
		[]byte("zyzxabcdez")))
}

func TestLongestCommonSubstringWithSuffixArrays(t *testing.T) {
	assert.Panics(t, func() {
		LongestCommonSubstringWithSuffixArrays([][]byte{[]byte("abc")}, [][]int{})
	})
	assert.Panics(t, func() {
		LongestCommonSubstringWithSuffixArrays(
			[][]byte{[]byte("abc"), []byte("cba")},
			[][]int{{1, 2}, {1, 2, 3}})
	})
	assert.Equal(t, []byte{}, LongestCommonSubstringWithSuffixArrays(nil, nil))
	assert.Equal(t, []byte("abc"), LongestCommonSubstringWithSuffixArrays(
		[][]byte{[]byte("abc")}, [][]int{{1, 2, 3}}))
}
