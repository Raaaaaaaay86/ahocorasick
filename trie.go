package ahocorasick

func NewTrie(texts []string) *Node {
	root := NewRootNode()

	for _, text := range texts {
		current := root

		runes := []rune(text)

		for i, r := range runes {
			next, exists := current.Next[r]
			if !exists {
				next = NewNode(r, current)
				current.Next[r] = next
			}

			current = next

			if i == len(runes)-1 {
				current.Exists = append(current.Exists, len(runes))
			}
		}
	}

	queue := []*Node{root}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.Parent != nil {
			if current.Parent == root {
				current.Fail = root
			} else {
				for parentFail := current.Parent.Fail; parentFail != nil; parentFail = parentFail.Fail {
					if next, exists := parentFail.Next[current.Rune]; exists {
						current.Fail = next

						if next.Exists != nil {
							current.Exists = append(current.Exists, next.Exists...)
						}

						break
					}
				}

				if current.Fail == nil {
					current.Fail = root
				}
			}
		}

		for _, next := range current.Next {
			queue = append(queue, next)
		}
	}

	return root
}
