package common

import "errors"

const (
	RED bool = true
	BLACK bool = false
)

const (
	LEFTROTATE bool = true
	RIGHTROTATE bool = false
)

type RBNode struct {
	value				int
	color				bool
	left, right, parent	*RBNode
}

func NewRBNode(data int) *RBNode {
	return &RBNode{value:data,color:RED}
}

func (rbnode *RBNode) getGrandParent() *RBNode {
	if rbnode.parent == nil {
		return nil
	}
	return rbnode.parent.parent
}

func (rbnode *RBNode) getSibling() *RBNode {
	if rbnode.parent == nil {
		return nil
	}
	if rbnode == rbnode.parent.left {
		return rbnode.parent.right
	} else {
		return rbnode.parent.left
	}
}

func (rbnode *RBNode) getUncle() *RBNode {
	if rbnode.getGrandParent() == nil {
		return nil
	}
	if rbnode.parent == rbnode.getGrandParent().right {
		return rbnode.getGrandParent().left
	} else {
		return rbnode.getGrandParent().right
	}
}

func (rbnode *RBNode) rotate(isRotateLeft bool) (*RBNode, error) {
	var root *RBNode
	if rbnode == nil {
		return root, nil
	}
	if !isRotateLeft && rbnode.left == nil {
		return root, errors.New("右旋左节点不能为空")
	} else if isRotateLeft && rbnode.right == nil {
		return root, errors.New("左旋右节点不能为空")
	}

	parent := rbnode.parent
	var isleft bool // 判断当前节点是不是父节点的左节点
	if parent != nil {
		isleft = parent.left == rbnode
	}
	if isRotateLeft {
		grandson := rbnode.right.left
		rbnode.right.left = rbnode
		rbnode.parent = rbnode.right
		rbnode.right = grandson
	} else {
		grandson := rbnode.left.right
		rbnode.left.right = rbnode
		rbnode.parent = rbnode.left
		rbnode.left = grandson
	}
	// 判断是否换了根节点
	if parent == nil {
		rbnode.parent.parent = nil
		root = rbnode.parent
	} else {
		if isleft {
			parent.left = rbnode.parent
		} else {
			parent.right = rbnode.parent
		}
		rbnode.parent.parent = parent
	}
	return root, nil
}

// 获取某节点最左侧的叶子，删除有2个孩子的节点时用
func (rbnode *RBNode) getLeftMostChild() *RBNode {
	if rbnode.left == nil {
		return rbnode
	}
	return rbnode.left.getLeftMostChild()
}


