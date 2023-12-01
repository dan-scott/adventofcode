package day01

type TrieTestResult = uint

type Trie[R comparable, V any] interface {
	Test(r []R) (V, bool)
}

type T[R comparable, V any] struct {
	r []R
	v V
}

func ti[R comparable, V any](r []R, v V) T[R, V] {
	return T[R, V]{r, v}
}

func NewTrie[R comparable, V any](rvs ...T[R, V]) Trie[R, V] {
	root := &trieNode[R, V]{
		end:  false,
		next: make(map[R]*trieNode[R, V]),
	}

	for _, rv := range rvs {
		cr := root
		for _, r := range rv.r {
			if n, ok := cr.next[r]; ok {
				cr = n
			} else {
				next := &trieNode[R, V]{
					end:  false,
					next: make(map[R]*trieNode[R, V]),
				}
				cr.next[r] = next
				cr = next
			}
		}
		cr.v = rv.v
		cr.end = true
	}

	return root
}

type trieNode[R comparable, V any] struct {
	end  bool
	v    V
	next map[R]*trieNode[R, V]
}

func (t *trieNode[R, V]) Test(r []R) (V, bool) {
	if len(r) == 0 {
		return t.v, t.end
	}
	c := r[0]
	if n, ok := t.next[c]; ok {
		return n.Test(r[1:])
	}
	return t.v, t.end
}
