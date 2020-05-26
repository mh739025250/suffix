package suffix

type Active struct {
	tree *Tree

	node     *Node
	startIdx int
	length   int
}

func (active *Active) activeSon() *Node {
	startChar := active.tree.content[active.startIdx]
	return active.node.sonNode(startChar)
}

func (active *Active) next() {
	if active.node == active.tree.root {
		active.startIdx++
		if active.length > 0 {
			active.length--
		}
	} else {
		if active.node.linkNode != nil {
			active.node = active.node.linkNode
		} else {
			active.node = active.tree.root
			active.startIdx = len(active.tree.content) - active.tree.remain
			active.length = active.tree.remain - 1
		}
	}
}

func (active *Active) canonize() {
	if active.length > 0 {
		startChar := active.tree.content[active.startIdx]
		sonNode := active.node.sonNode(startChar)
		sonLength := sonNode.endIdx - sonNode.startIdx
		// sonLength < 0 means sonNode is a leaf node
		if sonLength > 0 && active.length >= sonLength {
			active.node = sonNode
			active.startIdx += sonLength
			active.length -= sonLength
			active.canonize()
		}
	}
}
