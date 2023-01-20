package main

func findSubsequences(nums []int) [][]int {
	return NewPermutator(nums).Make()
}

type Permutator struct {
	permutations [][]int
	permutation  []int
	nums         []int
}

func NewPermutator(nums []int) *Permutator {
	return &Permutator{permutations: make([][]int, 0, 100), permutation: make([]int, len(nums)), nums: nums}
}

func (p *Permutator) Make() [][]int {
	p.make(0, 0)
	return p.permutations
}

func (p *Permutator) make(nIndex, pIndex int) {
	if pIndex > 1 {
		p.savePermutation(pIndex)
	}

	if nIndex == len(p.nums) {
		return
	}

	skip, sIndex := [201]bool{}, 0
	permutation, nums := p.permutation, p.nums

	for ; nIndex < len(nums); nIndex++ {
		sIndex = nums[nIndex] + 100

		if skip[sIndex] || (pIndex > 0 && permutation[pIndex-1] > nums[nIndex]) {
			continue
		}

		skip[sIndex] = true
		permutation[pIndex] = nums[nIndex]
		p.make(nIndex+1, pIndex+1)
	}
}

func (p *Permutator) savePermutation(length int) {
	x := make([]int, length)
	copy(x, p.permutation)

	p.permutations = append(p.permutations, x)
}
