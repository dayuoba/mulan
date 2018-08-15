package mux

import (
	"fmt"

	ml "github.com/dayuoba/mulan"
)

// optimize for real world rest api
// we often use router paths like:
// GET github.com/api/user/dayuoba HTTP/1.1
// when we define a routes we my has two types
// static: /api/user?id=dayuoba
// static means no param in request path
// dynamic: /api/user/:id
// here dynamic means we has a dynamically request path
// as in mulan we define a router like
// s.Get("/api/user/:id")
// this will match request like: /api/user/foo
// but will not match /api/uer/foo/....
// here is the tree should like
//      					[root]: "/"
//       					   |
//  		[node]: "/api"   		 [node]: "/foo"
//

//  	[node]: "/user" 				 [node]: "/bar"

// [node]: "/:id" 								[node]: "/baz"

// Trie data structure for muxing
type Trie struct {
	Root *TNode
}

// PrintTree ...
func (t *Trie) PrintTree() {
	PrintEages(t.Root)
}

// PrintEages ...
func PrintEages(n *TNode) {
	fmt.Println("****PRINTING****")
	if len(n.eages) > 0 {
		for key, t := range n.eages {
			fmt.Println(key)
			if t.tag != leaf {
				PrintEages(t)
			}
		}
	}
}

//

// 请求的路径
// /api/user/123
// 注册的路由
// /api/user/:id
// /api/book/:category/:id

// TNode trie's node
type TNode struct {
	Path       RoutePath
	tag        PathType
	Method     Methods
	handler    *ml.RouterHandler
	eages      Eages
	eageLevels int
	key        string
}

// PathType ...
type PathType uint8

// Eages ...
type Eages map[string]*TNode

const (
	leaf = iota
	eage = iota
	node = iota
)

// RoutePath ...
type RoutePath string

// Methods type alias
type Methods string

// Each ...
func Each(s string, iter func(item string, index int)) {
	len := len(s)
	for i := 0; i < len; i++ {
		iter(string(s[i]), i)
	}
}

// New ...
func (t *Trie) New() {
	t.Root = new(TNode)
	t.Root.eages = make(Eages)
	t.Root.eages["/"] = new(TNode)
	t.Root.eages["/"].tag = leaf
}

// Find ...
func (t *Trie) Find(path string) {
	for key := range t.Root.eages {
		lcp := LCP(key, path)
		llcp := len(lcp)

		if llcp == 0 {
			continue
		}
	}
}

// Insert ...
func (t *Trie) Insert(s string) {
	for key := range t.Root.eages {
		lcp := LCP(s, key)
		llcp := len(lcp)
		if llcp == 0 {
			continue
		} else {
			if len(s) < len(key) {
				// we has this
				break
			} else if llcp < len(s) && llcp < len(key) {
				cuts := Cut(s, llcp)
				cutk := Cut(key, llcp)
				fmt.Println(" lcp: ", lcp, " s: ", s, " key: ", key, " cuts: ", cuts, " cutk: ", cutk)

				t.Root.eages[lcp] = new(TNode)
				t.Root.eages[lcp].eages = make(Eages)
				t.Root.eages[lcp].eages[cutk] = new(TNode)
				t.Root.eages[lcp].eages[cutk].eages = make(Eages)
				t.Root.eages[lcp].eages[cutk].eages = t.Root.eages[key].eages
				delete(t.Root.eages, key)
				// add two new eages
				tr := new(Trie)
				tr.Root = t.Root.eages[lcp]
				tr.Insert(cuts)
				return

			} else {
				o := Cut(s, len(key))

				tr := new(Trie)
				tr.Root = t.Root.eages[key]
				tr.Insert(o)
				return
			}
		}
	}
	// has none in this level
	// create a new one in this
	if t.Root.eages == nil {
		t.Root.eages = make(Eages)
	}
	t.Root.eages[s] = new(TNode)
	t.Root.eages[s].tag = leaf
}

// Cut ...
func Cut(s string, scope int) string {
	o := ""
	sl := len(s)
	if scope >= sl {
		return ""
	}

	for i := scope; i < sl; i++ {
		o += string(s[i])
	}

	return o
}

// LCP longest common prefix
func LCP(s1, s2 string) string {
	s := ""
	if len(s1) == 0 || len(s2) == 0 {
		return s
	}
	l1 := len(s1)
	l2 := len(s2)
	for i := 0; i < l1; i++ {
		if i < l1 && i < l2 && s1[i] == s2[i] {
			s += string(s1[i])
		} else {
			return s
		}
	}
	return s
}
