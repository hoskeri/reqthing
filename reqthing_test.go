package reqthing

import "testing"

func TestBasic(t *testing.T) {
	var r = New(Options{})
	t.Logf("reqthing instance: %+v", r)
}
