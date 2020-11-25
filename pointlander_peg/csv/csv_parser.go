package pointlander_peg_csv

import (
	"errors"
	"fmt"
)

type Row []string

type CsvDoc struct {
	header []string
	rows []Row
}

func (p *CSV) Doc() (CsvDoc, error) {
	result := CsvDoc{}
	root := p.AST()

	rows := make([]Row, 0)

	for n := root.up; n != nil; n = n.next {
		switch n.pegRule {
		case ruleheader: {
			if h, err := p.header(n); err != nil {
				return result, err
			} else {
				result.header = h
			}
		}
		case rulerow: {
			if r, err := p.row(n); err != nil {
				return result, err
			} else {
				rows = append(rows, r)
			}
		}
		}
	}
	return result, nil
}

func (p *CSV) header(n *node32) (Row, error) {
	if !checkNodeType(n, ruleheader) {
		return nil, errors.New(fmt.Sprintf("not in rule:%v", ruleheader))
	}
	rowNode := n.up
	return p.row(rowNode)
}

func (p *CSV) row(n *node32) (Row, error) {
	result := make([]string, 0)
	if !checkNodeType(n, rulerow) {
		return result, errors.New(fmt.Sprintf("not in rule:%v", rulerow))
	}

	for i := n.up; i != nil; i = n.next {
		if i.pegRule == rulecol {
			if c, err := p.col(i); err != nil {
				return result, err
			} else {
				result = append(result, c)
			}
		}
	}
	return result, nil
}

func (p *CSV) col(n *node32) (string, error) {
	if !checkNodeType(n, rulecol) {
		return "", errors.New(fmt.Sprintf("not in rule:%v", rulecol))
	}
	return string(p.buffer[n.begin:n.end]), nil
}

func checkNodeType(n *node32, ruleType pegRule) bool {
	if n != nil && n.pegRule == ruleType {
		return true
	}
	return false
}
