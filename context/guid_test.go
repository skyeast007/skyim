package context

import (
	"fmt"
	"testing"
)

func Test_NextID(t *testing.T) {
	g, err := NewGUID(int64(1))
	if err != nil {
		t.Error(err)
	}
	var ID int64
	for i := 0; i < 100; i++ {
		ID, err = g.NextID()
		if err != nil {
			t.Error(err)
		}
		println(ID)
	}
}

func Test_GetIncreaseID(t *testing.T) {
	g, err := NewGUID(int64(1))
	if err != nil {
		t.Error(err)
	}
	var x = uint64(0)
	for i := 0; i < 100; i++ {
		ID := g.GetIncreaseID(&x)
		fmt.Printf("%v\n", ID)
	}
}
