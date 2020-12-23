package cups

import (
	"strconv"
	"strings"
)

type Cup struct {
	prev  *Cup
	next  *Cup
	value int
}

type CupRing struct {
	front         *Cup
	cupValueToCup map[int]*Cup
}

func NewCupRingFromValues(values []int) CupRing {
	cupValueToCup := make(map[int]*Cup)
	front := &Cup{
		prev:  nil,
		next:  nil,
		value: values[0],
	}
	cupValueToCup[values[0]] = front
	prev := front
	for _, val := range values[1:] {
		cur := &Cup{
			prev:  prev,
			next:  nil,
			value: val,
		}
		cupValueToCup[val] = cur
		prev.next = cur
		prev = cur
	}
	prev.next = front
	front.prev = prev

	return CupRing{
		front:         front,
		cupValueToCup: cupValueToCup,
	}
}

func (r *CupRing) DoRound() {
	// remove 3 cups
	cupStr := r.front.next
	r.front.next = r.front.next.next.next.next
	r.front.next.prev = r.front
	cupStr.prev = nil
	cupStr.next.next.next = nil

	// find dest
	destValue := r.front.value
	for ok := true; ok; ok = destValue == cupStr.value ||
		destValue == cupStr.next.value ||
		destValue == cupStr.next.next.value {

		destValue--
		if destValue <= 0 {
			destValue += len(r.cupValueToCup)
		}
	}
	dest := r.cupValueToCup[destValue]

	// insert 3 cups after dest
	cupStr.prev = dest
	cupStr.next.next.next = dest.next
	dest.next.prev = cupStr.next.next
	dest.next = cupStr

	// advance current
	r.front = r.front.next
}

func (r *CupRing) ComputeOrderStr() string {
	start := r.cupValueToCup[1]
	b := strings.Builder{}
	for cur := start.next; cur != start; cur = cur.next {
		b.WriteString(strconv.Itoa(cur.value))
	}
	return b.String()
}

func (r *CupRing) GetValuesAfterOne() (int, int) {
	one := r.cupValueToCup[1]
	return one.next.value, one.next.next.value
}
