package main

import "math/rand"

type RandomizedSet struct {
	dic map[int]int
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	dic := make(map[int]int)
	res := RandomizedSet{
		dic: dic,
	}
	return res
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	_, ok := this.dic[val]
	if !ok {
		this.dic[val]++
		return true
	}
	return false
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	_, ok := this.dic[val]
	if ok {
		delete(this.dic, val)
		return true
	}
	return false
}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	num := rand.Intn(len(this.dic))
	i := 0
	var a int
	for k, _ := range this.dic {
		a = k
		if i == num {
			return k
		}
		i++
	}
	return a
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */

func main() {
}
