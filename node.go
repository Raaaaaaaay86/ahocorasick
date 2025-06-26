package ahocorasick

type Node struct {
	Rune   rune
	Parent *Node
	Next   map[rune]*Node
	Fail   *Node
	Exists []int
}

type Match struct {
	StartIndex int
	EndIndex   int
}

func NewNode(r rune, parent *Node) *Node {
	return &Node{
		Rune:   r,
		Next:   make(map[rune]*Node),
		Parent: parent,
	}
}

func NewRootNode() *Node {
	return &Node{
		Next:   make(map[rune]*Node),
		Parent: nil,
	}
}

func (n *Node) IsRoot() bool {
	return n.Parent == nil
}

func (n *Node) FindLongestMatches(sentence string) []Match {
	return n.findMatches(sentence, true)
}

func (n *Node) FindAllMatches(sentence string) []Match {
	return n.findMatches(sentence, false)
}

func (n *Node) findMatches(sentence string, excludeSubMatches bool) []Match {
	matches := make([]Match, 0)

	characters := []rune(sentence)
	current := n

	i := 0
	for i < len(characters) {
		char := characters[i]

		next, ok := current.Next[char]
		if !ok {
			if !current.IsRoot() {
				current = current.Fail
				continue
			}

			i++
			continue
		}

		if len(next.Exists) > 0 {
			for _, length := range next.Exists {
				matches = append(matches, Match{
					StartIndex: (i + 1) - length,
					EndIndex:   i + 1,
				})

				if excludeSubMatches {
					break
				}
			}
		}

		i++
		current = next
	}

	return matches
}
