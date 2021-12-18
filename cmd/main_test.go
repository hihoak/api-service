package main

import (
	"github.com/golang/mock/gomock"
	"os"
	"testing"
)

func TestGetString(t *testing.T) {
	t.Parallel()
	got := getString()
	expect := "Hello, World!"

	if ! gomock.Eq(got).Matches(expect) {
		t.Errorf("Got %v, expect %v", got, expect)
	}
}

func TestMain(m *testing.M) {
	main()
	os.Exit(m.Run())
}
