package pointlander_peg_date

import (
	"errors"
	"fmt"
	"strconv"
)

type DateWrapper struct {
	year int32
	month int32
	day int32
}

func (p *Date) Pack() (DateWrapper, error) {
	result := DateWrapper{}
	root := p.AST()
	for n := root.up; n != nil; n = n.next {
		switch n.pegRule {
		case ruleyear: {
			if year, err := p.eYear(n); err != nil {
				return result, err
			} else {
				result.year = year
			}
		}
		case rulemonth: {
			if m, err := p.eMonth(n); err != nil {
				return result, err
			} else {
				result.month = m
			}
		}
		case ruleday: {
			if day, err := p.eDay(n); err != nil {
				return result, err
			} else {
				result.day = day
			}
		}
		}
	}
	return result, nil
}

func (p *Date) eYear(n *node32) (int32, error) {
	if b := checkNodeType(n, ruleyear); !b {
		return 0, errors.New(fmt.Sprintf("not in rule:%v", ruleyear))
	}
	if r, err := strconv.ParseInt(string(p.buffer[n.begin:n.end]), 10, 32); err != nil {
		return 0, err
	} else {
		return int32(r), nil
	}
}

func (p *Date) eMonth(n *node32) (int32, error) {
	if b := checkNodeType(n, rulemonth); !b {
		return 0, errors.New(fmt.Sprintf("not in rule:%v", rulemonth))
	}
	if r, err := strconv.ParseInt(string(p.buffer[n.begin:n.end]), 10, 32); err != nil {
		return 0, err
	} else {
		return int32(r), nil
	}
}

func (p *Date) eDay(n *node32) (int32, error) {
	if b := checkNodeType(n, ruleday); !b {
		return 0, errors.New(fmt.Sprintf("not in rule:%v", ruleday))
	}
	if r, err := strconv.ParseInt(string(p.buffer[n.begin:n.end]), 10, 32); err != nil {
		return 0, err
	} else {
		return int32(r), nil
	}
}

func checkNodeType(n *node32, ruleType pegRule) bool {
	if n != nil && n.pegRule == ruleType {
		return true
	}
	return false
}


