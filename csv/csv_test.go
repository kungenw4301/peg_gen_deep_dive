package csv

import "testing"

//go:generate pigeon -o ./csv_parser.go ./csv.peg
//go:generate goimports -w ./csv_parser.go


func Test_Csv_Parser(t *testing.T) {
	s := `key,value,desc
a,0,00
b,1,11
`
	got, err := Parse("", []byte(s), Debug(false))
	if err != nil {
		t.Errorf("err:%v", err)
	} else {
		t.Logf("got: %v", got)
	}
}
