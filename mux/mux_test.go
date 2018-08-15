package mux

import (
	"fmt"
	"testing"
)

type T *testing.T

func TestEcho(t *testing.T) {
	// var m *map[string]string
	a := 1
	b := 2
	m := make(map[string]*int)
	// *m = make(map[string]string)
	m["hi"] = &a
	m["hello"] = &a
	fmt.Println(m)

	s := make(map[string]*int)
	s = m
	fmt.Println(s)
	s["hi"] = &b

	for key, v := range m {
		fmt.Println(key, v)
	}

}

func TestNew(t *testing.T) {
	fmt.Println("hello testing mux")
}

func TestEach(t *testing.T) {
	s := "hello world"
	Each(s, func(item string, i int) {
		if i == 0 {
			fmt.Println(item == "h")
		}
		fmt.Println(item, i)
	})
}

func TestLCP(t *testing.T) {
	s1 := "hello world"
	s2 := "hoo"
	s := LCP(s1, s2)
	fmt.Println(s)

	s1 = "how are you"
	s2 = "how old are you"
	s = LCP(s1, s2)
	fmt.Println(s)
}

func TestInsert(t *testing.T) {
	path := "/ello"
	tr := new(Trie)
	tr.New()
	tr.Insert(path)

	path2 := "/ello/asd"
	tr.Insert(path2)

	path3 := "/api/user"
	tr.Insert(path3)

	path4 := "/api/book"
	tr.Insert(path4)

	path5 := "/apu/foo"
	tr.Insert(path5)

	fmt.Println(tr.Root.eages["/"].eages)
	fmt.Println(tr.Root.eages["/"].eages["ap"].eages)
	fmt.Println(tr.Root.eages["/"].eages["ap"].eages["i/"].eages)

	s := tr.Find("/api/book")
	fmt.Println(s)
	// fmt.Println(tr.Root.eages["/"].eages["api/"].eages)
	// fmt.Println(tr.Root.eages["/"].eages["ello"].eages)
	// fmt.Println(tr.Root.eages["/"].eages["/api/user"])
	// fmt.Println(tr.Root.eages["/"].eages["ello"].eages)
}

func TestCut(t *testing.T) {
	s := "hello world"
	fmt.Println(Cut(s, 10))

	s = "newyork"
	fmt.Println(Cut(s, 20))

}
