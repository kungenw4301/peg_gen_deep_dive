package pointlander_peg_date

import "testing"

func Test_ParseDate(t *testing.T) {
	d := Date{
		Pretty: true,
		Buffer: "2020-11-25",
	}
	d.Init()
	if err := d.Parse(); err != nil {
		t.Fatalf("parse failed, err:%v", err)
	}

	d.PrintSyntaxTree()

	if dateWrapper, err := d.Pack(); err != nil {
		t.Fatalf("err: %v", err)
	} else {
		t.Logf("date: %+v", dateWrapper)
	}
}
