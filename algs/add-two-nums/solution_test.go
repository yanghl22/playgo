package algs

import (
	"fmt"
	"testing"

	"github.com/yanghl22/playgo/structures"
	"github.com/yanghl22/playgo/utils"
)

type question2 struct {
	para2
	ans2
}

// para 是参数
// one 代表第一个参数
type para2 struct {
	one     []int
	another []int
}

// ans 是答案
// one 代表第一个答案
type ans2 struct {
	one []int
}

func Test_Problem2(t *testing.T) {

	qs := []question2{

		{
			para2{[]int{}, []int{}},
			ans2{[]int{}},
		},

		{
			para2{[]int{1}, []int{1}},
			ans2{[]int{2}},
		},

		{
			para2{[]int{1, 2, 3, 4}, []int{1, 2, 3, 4}},
			ans2{[]int{2, 4, 6, 8}},
		},

		{
			para2{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
			ans2{[]int{2, 4, 6, 8, 0, 1}},
		},

		{
			para2{[]int{1}, []int{9, 9, 9, 9, 9}},
			ans2{[]int{0, 0, 0, 0, 0, 1}},
		},

		{
			para2{[]int{9, 9, 9, 9, 9}, []int{1}},
			ans2{[]int{0, 0, 0, 0, 0, 1}},
		},

		{
			para2{[]int{2, 4, 3}, []int{5, 6, 4}},
			ans2{[]int{7, 0, 8}},
		},

		{
			para2{[]int{1, 8, 3}, []int{7, 1}},
			ans2{[]int{8, 9, 3}},
		},
		{
			para2{[]int{1, 8, 3}, []int{7, 1}},
			ans2{[]int{8, 9, 3}},
		},
		// 如需多个测试，可以复制上方元素。
	}

	fmt.Printf("------------------------Leetcode Problem 2------------------------\n")

	for _, q := range qs {
		a, p := q.ans2, q.para2
		p1 := structures.Ints2List(p.one)
		p2 := structures.Ints2List(p.another)
		result := structures.List2Ints(addTwoNumbers(p1, p2))
		if utils.IsSliceEqual(result, a.one) {

			t.Logf("Successful!		【input】:%v       【Expected】: %v			【output】:%v\n", p, a.one, result)
		} else {

			t.Errorf("Failed!		【input】:%v 		【Expected】: %v		【output】: %v\n", p, a.one, result)
		}
		// fmt.Printf("【input】:%v       【expected】:%v 			output】:%v\n", p, a, structures.List2Ints(addTwoNumbers(structures.Ints2List(p.one), structures.Ints2List(p.another))))
	}
	fmt.Printf("\n\n\n")
}
