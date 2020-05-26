package suffix

type Tree struct {
	root   *Node
	active *Active
	remain int

	content []rune
}

func NewSuffixTree() *Tree {
	node := &Node{
		startIdx: -1,
		endIdx:   -1,
		father:   nil,
		sonMap:   make(map[rune]*Node),
		linkNode: nil,
	}
	active := &Active{
		node:     node,
		startIdx: -1,
		length:   0,
	}

	tree := &Tree{
		root:   node,
		active: active,
		remain: 0,

		content: make([]rune, 0),
	}
	node.tree = tree
	active.tree = tree
	return tree
}

func (tree *Tree) AddChar(c rune) {
	tree.remain++
	currentIdx := len(tree.content)
	tree.content = append(tree.content, c)

	var lastNewNode *Node = nil

	for tree.remain > 0 {
		tree.active.canonize()
		if tree.active.length == 0 {
			if tree.active.node.sonNode(c) != nil {
				tree.active.startIdx = currentIdx
				tree.active.length++
				break
			} else {
				tree.active.node.addNode(currentIdx, -1)
				tree.remain--
				tree.active.next()
			}
		} else {
			activeSon := tree.active.activeSon()
			endIdx := activeSon.startIdx + tree.active.length
			if tree.content[endIdx] == c {
				tree.active.length++
				break
			} else {
				newNode := activeSon.split(endIdx)
				newNode.addNode(currentIdx, -1)
				tree.remain--
				if lastNewNode != nil {
					lastNewNode.linkNode = newNode
				}
				lastNewNode = newNode
				tree.active.next()
			}
		}
	}
}

func (tree *Tree) BuildFromStr(s string) {
	for _, c := range s {
		tree.AddChar(c)
	}
	tree.AddChar(0)
	tree.root.setLeafCount()
}

func (tree *Tree) ExistChars(chars []rune) bool {
	node := tree.root
	n := len(chars)
	for i := 0; i < n; {
		c := chars[i]
		son := node.sonNode(c)
		if son == nil {
			return false
		}
		endIdx := son.endIdx
		if endIdx < 0 {
			endIdx = len(tree.content)
		}
		for j := son.startIdx; j < endIdx && i < n; j++ {
			if chars[i] != tree.content[j] {
				return false
			}
			i++
		}
		node = son
	}
	return true
}

func (tree *Tree) ExistStr(s string) bool {
	return tree.ExistChars([]rune(s))
}

func (tree *Tree) CountChars(chars []rune) int {
	node := tree.root
	n := len(chars)
	for i := 0; i < n; {
		c := chars[i]
		son := node.sonNode(c)
		if son == nil {
			return 0
		}
		endIdx := son.endIdx
		if endIdx < 0 {
			endIdx = len(tree.content)
		}
		for j := son.startIdx; j < endIdx && i < n; j++ {
			if chars[i] != tree.content[j] {
				return 0
			}
			i++
		}
		node = son
	}
	return node.leafCount
}

func (tree *Tree) CountStr(s string) int {
	return tree.CountChars([]rune(s))
}
