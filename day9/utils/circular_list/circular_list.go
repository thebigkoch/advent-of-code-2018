package circular_list

import "errors"

type CircleNode struct {
	Value int
	Prev  *CircleNode
	Next  *CircleNode
}

func (n *CircleNode) InsertNext(value int) error {

	if n == nil {
		return errors.New("Cannot insert a value into a nil circle")
	}

	newNode := CircleNode{Value: value}
	newNode.Next = n.Next
	newNode.Prev = n
	n.Next.Prev = &newNode
	n.Next = &newNode

	return nil
}

func (n *CircleNode) GetPosition(relativePosition int) (*CircleNode, error) {

	if n == nil {
		return nil, errors.New("Cannot get the position in a nil circle")
	}

	resultNode := n
	if relativePosition < 0 {
		for counter := 0; counter > relativePosition; counter-- {
			resultNode = resultNode.Prev
		}
	} else if relativePosition > 0 {
		for counter := 0; counter < relativePosition; counter++ {
			resultNode = resultNode.Next
		}
	}

	return resultNode, nil
}
