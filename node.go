package suffix

type Node struct {
	tree *Tree

	startIdx int
	endIdx   int

	father   *Node
	sonMap   map[rune]*Node
	linkNode *Node

	leafCount int
}

func (node *Node) sonNode(c rune) *Node {
	if son, ok := node.sonMap[c]; ok {
		return son
	} else {
		return nil
	}
}

func (node *Node) addNode(startIdx, endIdx int) *Node {
	startChar := node.tree.content[startIdx]
	newNode := &Node{
		tree:     node.tree,
		startIdx: startIdx,
		endIdx:   endIdx,
		father:   node,
		sonMap:   make(map[rune]*Node),
		linkNode: nil,
	}
	node.sonMap[startChar] = newNode
	return newNode
}

func (node *Node) split(newEndIdx int) *Node {
	// replace old node with new node
	newNode := &Node{
		tree:     node.tree,
		startIdx: node.startIdx,
		endIdx:   newEndIdx,
		father:   node.father,
		sonMap:   make(map[rune]*Node),
		linkNode: nil,
	}
	startChar := node.tree.content[node.startIdx]
	endChar := node.tree.content[newEndIdx]
	node.father.sonMap[startChar] = newNode

	// shorten old node, add old node to the new node
	node.startIdx = newEndIdx
	node.father = newNode
	newNode.sonMap[endChar] = node
	return newNode
}

func (node *Node) setLeafCount() {
	if node.endIdx == -1 && node.startIdx > 0 {
		node.leafCount = 1
		return
	}

	for _, son := range node.sonMap {
		son.setLeafCount()
		node.leafCount += son.leafCount
	}

}
