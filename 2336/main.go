package main

import "github.com/emirpasic/gods/sets/treeset"

type SmallestInfiniteSet struct {
	thres int
	s     *treeset.Set
}

func Constructor() SmallestInfiniteSet {
	return SmallestInfiniteSet{thres: 1,
		s: treeset.NewWithIntComparator()}
}

func (this *SmallestInfiniteSet) PopSmallest() int {
	if this.s.Empty() {
		ans := this.thres
		this.thres++
		return ans
	}
	it := this.s.Iterator()
	it.Next()
	ans := it.Value().(int)
	this.s.Remove(ans)
	return ans
}

func (this *SmallestInfiniteSet) AddBack(num int) {
	if num < this.thres {
		this.s.Add(num)
	}
}
func main() {

}
