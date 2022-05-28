package algs

import (
	"fmt"
	"testing"

	"github.com/yanghl22/playgo/utils"
)

type question1 struct {
	para1
	anw1
}
type para1 struct {
	nums   []int
	target int
}
type anw1 struct {
	anw []int
}

func Test_Problem1(t *testing.T) {
	qs := []question1{{para1{[]int{1, 3, 5, 6, 8}, 9}, anw1{[]int{1, 3}}}}

	fmt.Printf("------------------------Leetcode Problem 1------------------------\n")

	for _, q := range qs {
		a, p := q.anw1, q.para1
		result := twoSum(p.nums, p.target)

		if utils.IsSliceEqual(result, a.anw) {

			t.Logf("Successful!		【input】:%v       【Expected】: %v			【output】:%v\n", p, a.anw, result)
		} else {

			t.Errorf("Failed!		【input】:%v 		【Expected】: %v		【output】: %v\n", p, a.anw, result)
		}
	}
	fmt.Printf("\n\n\n")
}
