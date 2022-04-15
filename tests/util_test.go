package tests

import (
	"fmt"
	"testing"

	"github.com/tongsq/go-lib/logger"
	"github.com/tongsq/go-lib/util"
)

func TestAdd(t *testing.T) {
	if util.Add(1, 2, 3) == 6 {
		t.Log("test Add success")
	} else {
		t.Fatal("test add fail")
	}
}

func TestMax(t *testing.T) {
	if util.Max(1, 3, 2) == 3 {
		t.Log("test Max success")
	} else {
		t.Fatal("test Max fail")
	}
}

func TestMin(t *testing.T) {
	if util.Min(1, 3, 2) == 1 {
		t.Log("test Min success")
	} else {
		t.Fatal("test Min fail")
	}
}

func printStackFunc() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			logger.PrintStack()
		}
	}()
	var m map[string]struct{}
	m["a"] = struct{}{}
}

func TestPrintStack(t *testing.T) {
	printStackFunc()
}
