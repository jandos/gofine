package gofine_test

import (
	"testing"

	"github.com/jandos/gofine"
)

var (
	env = &gofine.Environment{}
)

func TestProcess(t *testing.T) {
	// TODO write parallel tests
	err := env.InitDefault()
	if err != nil {
		t.Fatal(err)
	}

	lgoreCount := env.LgoreCount()
	if lgoreCount <= 0 {
		t.Fatal("lgore count should be greater than zero")
	}

	lgoreId := 0
	err = env.Occupy(lgoreId)
	if err != nil {
		t.Fatal(err)
	}

	state, err := env.GetLgoreState(lgoreId)
	if err != nil {
		t.Fatal(err)
	}
	if state != gofine.Busy {
		t.Fatal("lgore state should be Busy")
	}

	err = env.Release(lgoreId)
	if err != nil {
		t.Fatal(err)
	}

	state, err = env.GetLgoreState(lgoreId)
	if err != nil {
		t.Fatal(err)
	}
	if state != gofine.Available {
		t.Fatal("lgore state should be Available")
	}
}
