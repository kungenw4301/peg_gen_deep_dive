package pointlander_peg_hour_min

import "testing"

func Test_HourMin(t *testing.T) {
	h := HourMin{Buffer: "99:60"}
	h.Init()
	if err := h.Parse(); err != nil {
		t.Fatal(err)
	}
	hour, min := h.pack()
	t.Logf("hour:%v, min:%v", hour, min)
}
