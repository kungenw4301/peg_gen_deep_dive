package date

import "testing"

//go:generate pigeon -o ./date_parser.go ./date.peg
//go:generate goimports -w ./date_parser.go

func Test_ParseDate(t *testing.T) {
	s := "2020-10-11"
	s = "0000-01-01"
	got, err := Parse("", []byte(s), Debug(true))
	if err != nil {
		t.Errorf("err:%v", err)
	} else {
		t.Logf("got: %v", got)
	}
}

