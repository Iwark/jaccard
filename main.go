package main

import (
	"fmt"
	"math/rand"
	"time"

	set "github.com/deckarep/golang-set"
)

func main() {
	rand.Seed(time.Now().Unix())
	s1 := slice(1000000)
	s2 := slice(1000000)
	t1 := time.Now()
	d1 := setJaccard(s1, s2)
	t2 := time.Now()
	fmt.Println("setJaccard: ", d1, " spent:", t2.Sub(t1))
	d2 := jaccard(s1, s2)
	t3 := time.Now()
	fmt.Println("jaccard: ", d2, " spent:", t3.Sub(t2))
	fmt.Println(float32(t2.Sub(t1).Nanoseconds())/float32(t3.Sub(t2).Nanoseconds()), "times faster than set")
}

func slice(size int) []interface{} {
	s := make([]interface{}, 0, size)
	for i := 0; i < size; i++ {
		s = append(s, rand.Intn(100000000))
	}
	return s
}

func jaccard(s1, s2 []interface{}) float32 {
	tmp := make(map[int]int, len(s1)+len(s2))
	unionNum := 0
	intersectNum := 0

	for _, s := range s1 {
		tmp[s.(int)] = 1
		unionNum++
	}
	for _, s := range s2 {
		switch tmp[s.(int)] {
		case 1:
			intersectNum++
		case 2:
			continue
		}
		unionNum++
		tmp[s.(int)] = 2
	}
	return float32(intersectNum) / float32(unionNum)
}

func setJaccard(s1, s2 []interface{}) float32 {
	ss1 := set.NewSetFromSlice(s1)
	ss2 := set.NewSetFromSlice(s2)
	return float32(ss1.Intersect(ss2).Cardinality()) / float32(ss1.Union(ss2).Cardinality())
}
