package slimlist

import (
	"strconv"
	"strings"
)

type Node struct {
	next *Node
	str  string
	list *SlimList
}

type SlimList struct {
	length int
	head   *Node
	tail   *Node
}

func Node_Create(str string) *Node {
	return &Node{str: str}
}

func SlimList_Create() *SlimList {
	return new(SlimList)
}

func (s *SlimList) CreateIterator() *Node {
	return s.head
}

func (s *SlimList) AddString(str string) {
	newNode := Node_Create(str)

	s.insertNode(newNode)
}

func (s *SlimList) insertNode(node *Node) {
	if s.length == 0 {
		s.head = node
	} else {
		s.tail.next = node
	}
	s.tail = node
	s.length++
}

func (s *SlimList) AddList(element *SlimList) {
	s.AddString(element.Serialize())
}

func (s *SlimList) PopHead() {
	if s.head == nil {
		panic("nil pointer")
	}
	previousHead := s.head
	s.head = previousHead.next

	if s.tail == previousHead {
		s.tail = nil
	}

	s.length--
}

func (s *SlimList) GetLength() int {
	return s.length
}

func (s *SlimList) Equals(other *SlimList) bool {
	if s.length != other.length {
		return false
	}

	for p, q := s.head, other.head; p != nil; p, q = p.next, q.next {
		if strings.Compare(p.str, q.str) != 0 {
			return false
		}
	}

	return true
}

func (s *SlimList) GetNodeAt(index int) *Node {
	node := s.head

	if index >= s.length {
		return nil
	}

	for i := 0; i < index; i++ {
		node = node.next
	}
	return node
}

func (s *SlimList) GetListAt(index int) *SlimList {
	node := s.GetNodeAt(index)
	if node == nil {
		return nil
	}
	return node.GetList()
}

func (s *SlimList) GetStringAt(index int) string {
	node := s.GetNodeAt(index)
	if node == nil {
		return ""
	}

	return node.GetString()
}

func (s *SlimList) GetDoubleAt(index int) float64 {
	i, _ := strconv.ParseFloat(s.GetStringAt(index), 64)
	return i
}

func (s *SlimList) GetFloatAt(index int) float32 {
	i, _ := strconv.ParseFloat(s.GetStringAt(index), 32)
	return float32(i)
}

func (s *SlimList) GetHashAt(index int) *SlimList {
	return SlimList_deserializeHash(s.GetStringAt(index))
}

func (s *SlimList) ReplaceAt(index int, replacementString string) {
	node := s.GetNodeAt(index)
	node.Replace(replacementString)
}

func (s *SlimList) GetTailAt(index int) *SlimList {
	tail := SlimList_Create()

	iterator := s.head
	iterator = iterator.AdvanceBy(index)

	for ; iterator != nil; iterator = iterator.next {
		tail.AddString(iterator.GetString())
	}

	return tail
}

func (s *SlimList) ToString() string {
	result := "["

	for iterator := s.head; iterator != nil; iterator = iterator.next {
		sublist := iterator.GetList()

		if sublist != nil {
			result += sublist.ToString()
		} else {
			result += "\""
			result += iterator.GetString()
			result += "\""
		}

		if iterator.next != nil {
			result += ", "
		}
	}
	result += "]"

	return result
}

func (n *Node) Advance() *Node {
	return n.next
}

func (n *Node) GetString() string {
	return n.str
}

func (n *Node) GetList() *SlimList {
	if n.list == nil {
		n.list = SlimList_Deserialize(n.str)
	}
	return n.list
}

func (n *Node) Replace(replacementString string) {
	n.str = replacementString
}

func (n *Node) AdvanceBy(amount int) *Node {
	nextNode := n
	for i := 0; i < amount; i++ {
		if nextNode != nil {
			nextNode = nextNode.next
		}
	}
	return nextNode
}

func (n *Node) GetStringWithNullAsANormalString() string {
	if n.str == "" {
		return "null"
	} else {
		return n.str
	}
}
