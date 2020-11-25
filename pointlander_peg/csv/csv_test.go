package pointlander_peg_csv

import "testing"

func Test_csv_parser(t *testing.T) {
	s := `key,value,desc
a,0,00
b,1,11
`
	c := CSV{
		Pretty: true,
		Buffer: s,
	}
	c.Init()
	if err := c.Parse(); err != nil {
		t.Fatalf("parse csv error, err:%v", err)
	}

	c.PrintSyntaxTree()

	if doc, err := c.Doc(); err != nil {
		t.Fatalf("pack doc error, err:%v", err)
	} else {
		t.Logf("result csv document:%v", doc)
	}
}
