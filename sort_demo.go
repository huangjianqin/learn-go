package main

import (
	"fmt"
	"sort"
)

type Tuple struct {
	first  int
	second int
}

type TupleSlice []Tuple

func (ts TupleSlice) Len() int {
	return cap(ts)
}

func (ts TupleSlice) Less(i, j int) bool {
	return ts[i].first <= ts[j].first || (ts[i].first == ts[j].first && ts[i].second <= ts[j].second)
}

func (ts TupleSlice) Swap(i, j int) {
	tmp := ts[i]
	ts[i] = ts[j]
	ts[j] = tmp
}

func main() {
	sort1()

	sort2()
}

func sort1() {
	ints := []int{2, 4, 5, 61, 2, 1, 98, 1, 2, 56}
	fmt.Printf("%v\n", ints)
	sort.Ints(ints)
	fmt.Printf("%v\n", ints)
}

func sort2() {
	ts := TupleSlice{Tuple{1, 2}, Tuple{3, 2}, Tuple{2, 3}}
	fmt.Printf("%v\n", ts)
	sort.Sort(ts)
	fmt.Printf("%v\n", ts)
}
