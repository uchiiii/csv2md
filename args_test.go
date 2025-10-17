package main

import (
	"testing"
)

func TestArgs(t *testing.T) {

	t.Run("[test 1] should pass", func(t *testing.T) {
		args := &Args{
			Files: []string{"test1.csv", "test2.csv"},
			Delim: ",",
			Pad:   2,
		}
		err := args.ValidateAll()
		if err != nil {
			t.Error(err.Error())
			t.Errorf("Input Args: %+v\n", args)
		}
	})

	t.Run("[test 2] should pass", func(t *testing.T) {
		args := &Args{
			Files: []string{"test1.csv"},
			Delim: ";",
			Pad:   5,
		}
		err := args.ValidateAll()
		if err != nil {
			t.Error(err.Error())
			t.Errorf("Input Args: %+v\n", args)
		}
	})

	t.Run("[test 3] should fail (lack of filename)", func(t *testing.T) {
		args := &Args{
			Files: []string{},
			Delim: ";",
			Pad:   2,
		}
		err := args.ValidateAll()
		if err == nil {
			t.Errorf("Input Args: %+v\n", args)
		}
	})

	t.Run("[test 4] should fail (invalid delimiter)", func(t *testing.T) {
		args := &Args{
			Files: []string{"test.csv"},
			Delim: "!",
			Pad:   2,
		}
		err := args.ValidateAll()
		if err == nil {
			t.Errorf("Input Args: %+v\n", args)
		}
	})

	t.Run("[test 5] should fail (negative padding value)", func(t *testing.T) {
		args := &Args{
			Files: []string{"test.csv"},
			Delim: ",",
			Pad:   -1,
		}
		err := args.ValidateAll()
		if err == nil {
			t.Errorf("Input Args: %+v\n", args)
		}
	})
}
