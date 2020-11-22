package hour_min

import "testing"

//go:generate pigeon -o hour_min_parser.go ./hour_min.peg
//go:generate goimports -w ./hour_min_parser.go

func Test_ParseHour(t *testing.T) {
	s := "23:10"
	debugOption := Debug(true)
	got, err := Parse("", []byte(s), debugOption)
	if err != nil {
		t.Errorf("err:%v", err)
	} else {
		hourMin := got.(HourMin)
		t.Logf("got hour: %v, got minute: %v", hourMin.hour, hourMin.min)
	}
}

