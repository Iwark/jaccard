package jaccard

import "github.com/deckarep/golang-set"

// CalcurateBySlices calcurate the jaccard score
func CalcurateBySlices(s1, s2 []int) float32 {

	tmp := make(map[int]int, len(s1)+len(s2))
	var s1ors2, s1ands2 int

	for _, s := range s1 {
		if tmp[s] == 0 {
			s1ors2++
			tmp[s] = 1
		}
	}
	for _, s := range s2 {
		switch tmp[s] {
		case 0:
			s1ors2++
		case 1:
			s1ands2++
		}
		tmp[s] = 2
	}
	return float32(s1ands2) / float32(s1ors2)
}

// CalcurateBySets calcurate the jaccard score
func CalcurateBySets(s1, s2 mapset.Set) float32 {
	s1ands2 := s1.Intersect(s2).Cardinality()
	s1ors2 := s1.Union(s2).Cardinality()
	return float32(s1ands2) / float32(s1ors2)
}
