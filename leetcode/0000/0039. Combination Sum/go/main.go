package main

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	return NewPermutator(candidates, target).Make()
}

type Permutator struct {
	permutations [][]int
	permutation  []int
	nums         []int
	target       int
}

func NewPermutator(nums []int, target int) *Permutator {
	return &Permutator{
		permutations: make([][]int, 0, 64),
		permutation:  make([]int, 40),
		nums:         nums,
		target:       target,
	}
}

func (p *Permutator) Make() [][]int {
	p.make(0, 0, 0)
	return p.permutations
}

func (p *Permutator) make(sum, nIndex, pIndex int) {
	for ; nIndex < len(p.nums); nIndex++ {
		switch nextSum := sum + p.nums[nIndex]; {
		case nextSum < p.target:
			p.permutation[pIndex] = p.nums[nIndex]
			p.make(nextSum, nIndex, pIndex+1)
		case nextSum == p.target:
			p.permutation[pIndex] = p.nums[nIndex]
			p.savePermutation(pIndex + 1)
			return
		default:
			return
		}
	}
}

func (p *Permutator) savePermutation(length int) {
	x := make([]int, length)
	copy(x, p.permutation)

	p.permutations = append(p.permutations, x)
}
