type Stack []int

func (s *Stack) Push(v int) {
  *s = append(*s, v)
}

func (s *Stack) Pop() (int, bool) {
	if (*s).IsEmpty() {
        return 0, false
	}
	index := len(*s) - 1
	element := (*s)[index]
	*s = (*s)[:index]
	return element, true
}

func (s *Stack) IsEmpty() bool {
  return len(*s) == 0
}

func (s *Stack) Peak() (int, bool) {
	if (*s).IsEmpty() {
        return 0, false
	}

	index := len(*s) - 1
	return (*s)[index], true
}

func nextGreaterElement(nums1 []int, nums2 []int) []int {
    var stack Stack
    maps := make(map[int]int)
    result := make([]int, len(nums1))

    for _, v := range nums2 {
        for !stack.IsEmpty() {
            peak, _ := stack.Peak()
            if peak < v {
                top, _ := stack.Pop()
                maps[top] = v
            } else {
                break
            }
		}
        stack.Push(v)
    }

    for i, v := range nums1 {
        if val, ok := maps[v]; ok {
            result[i] = val
        } else {
            result[i] = -1
        }
    }

    return result
}
