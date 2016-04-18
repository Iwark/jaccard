package jaccard

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/deckarep/golang-set"
	"github.com/stretchr/testify/suite"
)

const sliceSize = 200000

type JaccardTestSuite struct {
	suite.Suite
	s1  []int
	s2  []int
	st1 int64
	st2 int64
}

func (suite *JaccardTestSuite) SetupSuite() {
	rand.Seed(time.Now().Unix())
	suite.s1 = makeSlice(sliceSize)
	suite.s2 = makeSlice(sliceSize)
}

func (suite *JaccardTestSuite) TestCalcurateBySlices() {
	t1 := time.Now()
	d := CalcurateBySlices(suite.s1, suite.s2)
	fmt.Println("jaccard: ", d, " spent:", time.Now().Sub(t1))
	suite.st1 = time.Now().Sub(t1).Nanoseconds()
}

func (suite *JaccardTestSuite) TestCalcurateBySets() {
	t1 := time.Now()
	ss1 := mapset.NewSet()
	ss2 := mapset.NewSet()
	for _, i := range suite.s1 {
		ss1.Add(i)
	}
	for _, i := range suite.s2 {
		ss2.Add(i)
	}
	d := CalcurateBySets(ss1, ss2)
	fmt.Println("jaccard_by_set: ", d, " spent:", time.Now().Sub(t1))
	suite.st2 = time.Now().Sub(t1).Nanoseconds()
}

func (suite *JaccardTestSuite) TearDownSuite() {
	fmt.Println(float32(suite.st2)/float32(suite.st1), "times faster than set")
}

func makeSlice(size int) []int {
	s := make([]int, 0, size)
	for i := 0; i < size; i++ {
		s = append(s, rand.Intn(100000000))
	}
	return s
}

func TestRunSuite(t *testing.T) {
	suite.Run(t, new(JaccardTestSuite))
}
