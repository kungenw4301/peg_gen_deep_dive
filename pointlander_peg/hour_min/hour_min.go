package pointlander_peg_hour_min

import (
	"fmt"
	"strconv"
)

type HM struct {
	hour int32
	min int32
}


func (t *HourMin) pack() (int32, int32) {
	n := t.AST()
	hm := HM{}
	if n.pegRule == ruletime {
		for node := n; node != nil; node = node.next {
			switch node.pegRule {
			case rulehour: {
				hm.hour = t.walkHour(node)
			}
			case rulemin: {
				hm.min = t.walkMin(node)
			}
			}
		}
	} else {
		fmt.Printf("error, not time")
		return -1, -1
	}
	return hm.hour, hm.min
}

func (t *HourMin) walkHour(n *node32) int32 {
	switch n.pegRule {
	case rulehour: {
		s := string(t.buffer[n.begin:n.end])
		if parseInt, err := strconv.ParseInt(s, 10, 32); err != nil {
			fmt.Printf("walkHour error")
			return -1
		} else {
			return int32(parseInt)
		}
	}
	}
	return -1
}

func (t *HourMin) walkMin(n *node32) int32 {
	switch n.pegRule {
	case rulemin: {
		s := string(t.buffer[n.begin:n.end])
		if parseInt, err := strconv.ParseInt(s, 10, 32); err != nil {
			fmt.Printf("walkHour error")
			return -1
		} else {
			return int32(parseInt)
		}
	}
	}
	return -1
}