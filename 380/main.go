package main

import "math/rand"

type RandomizedSet struct {
	nums []int
	hmap map[int]int
}

func Constructor() RandomizedSet {
	return RandomizedSet{[]int{}, map[int]int{}}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.hmap[val]; ok {
		return false
	} else {
		this.hmap[val] = len(this.nums)
		this.nums = append(this.nums, val)
		return true
	}
}

func (this *RandomizedSet) Remove(val int) bool {
	if v, ok := this.hmap[val]; ok {
		last := len(this.nums) - 1
		this.nums[v] = this.nums[last]
		this.hmap[this.nums[v]] = v
		this.nums = this.nums[:last]
		delete(this.hmap, val)
		return true
	} else {
		return false
	}
}

func (this *RandomizedSet) GetRandom() int {
	return this.nums[rand.Intn(len(this.nums))]
}
func main() {

}
