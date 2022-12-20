package main

type NestedInteger struct{}

func (n *NestedInteger) IsInteger() bool           { return true }
func (n *NestedInteger) GetInteger() int           { return 0 }
func (n *NestedInteger) GetList() []*NestedInteger { return nil }

func depthSum(list []*NestedInteger) int {
	return sum(list, 1)
}

func sum(list []*NestedInteger, depth int) (result int) {
	for _, x := range list {
		if x.IsInteger() {
			result += x.GetInteger() * depth
			continue
		}

		result += sum(x.GetList(), depth+1)
	}

	return result
}
