package tests

import (
	"github.com/tongsq/go-lib/component"
	"testing"
	"time"
)

func TestTaskPool(t *testing.T) {
	pool := component.NewTaskPool(3)
	num := 0
	for i := 0; i < 5; i++ {
		pool.RunTask(func() {
			num++
		})
	}
	time.Sleep(time.Second)
	if num == 5 {
		t.Log("TestTaskPool success", num)
	} else {
		t.Fatal("TestTaskPool fail", num)
	}
}
