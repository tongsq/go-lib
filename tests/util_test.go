package tests

import (
	"github.com/tongsq/go-lib/util"
	"testing"
)

func TestAdd(t *testing.T) {
	if util.Add(1, 2, 3) == 6 {
		t.Log("test Add success")
	} else {
		t.Fatal("test add fail")
	}
}
