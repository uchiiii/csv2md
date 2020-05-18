package main

import (
	"encoding/csv"
	"io/ioutil"
	"strings"
	"testing"
)

// test for Convert()
func TestConvert(t *testing.T) {
	t.Run("should pass", func(t *testing.T) {
		testFile := "./testdata/test1.csv"
		args := &Args{
			Files: []string{testFile},
			Delim: ",",
			Pad:   2,
		}

		expectedTestFile := "./testdata/test1_expected.md"
		expected := readFile(expectedTestFile)

		md, err := Convert(testFile, args)
		if err != nil {
			t.Errorf(err.Error())
		}

		if md != expected {
			t.Errorf("Convert is something wrong. \n expected: %q \n output  : %q", expected, md)
		}
	})

	t.Run("should pass", func(t *testing.T) {
		testFile := "./testdata/test2.csv"
		args := &Args{
			Files: []string{testFile},
			Delim: ",",
			Pad:   2,
		}

		expectedTestFile := "./testdata/test2_expected.md"
		expected := readFile(expectedTestFile)

		md, err := Convert(testFile, args)
		if err != nil {
			t.Errorf(err.Error())
		}

		if md != expected {
			t.Errorf("Convert is something wrong. \n expected: %q \n output  : %q", expected, md)
		}
	})
}

// test for Modify function.
func TestModifyFunc(t *testing.T) {
	in := `first_name,last_name,username,comment
"Rob","Pike",rob,hello
Ken,Thompson,ken,"Hi.
I am good."
"Robert","Griesemer","gri","Nice to meet you.
Thank you"
`
	expected := `first_name,last_name,username,comment
"Rob","Pike",rob,hello
Ken,Thompson,ken,"Hi.<br/>I am good."
"Robert","Griesemer","gri","Nice to meet you.<br/>Thank you"
`

	r_in := csv.NewReader(strings.NewReader(in))
	records, _ := r_in.ReadAll()

	r_expected := csv.NewReader(strings.NewReader(expected))
	records_expected, _ := r_expected.ReadAll()

	out := Modify(records)

	if !isSame2dStringArray(out, records_expected) {
		t.Errorf("Modify is somethig wrong. \n expected: %q \n output  : %q", records_expected, out)
	}
}

// test for colMaxSize function.
func TestColMaxSize(t *testing.T) {
	t.Run("should pass", func(t *testing.T) {
		rows := [][]string{}
		row1 := []string{"hello", "ok", "see you soon."}
		rows = append(rows, row1)

		expected := []int{5, 2, 13}

		out := colMaxSize(rows)
		if !isSame1dIntArray(expected, out) {
			t.Errorf("colMaxSize is something wrong. \n expected: %v \n output  : %v", expected, out)
		}
	})

	t.Run("should pass", func(t *testing.T) {
		rows := [][]string{}
		row1 := []string{"はい", "こんにちは", "お願いします."}
		rows = append(rows, row1)

		expected := []int{2, 5, 7}

		out := colMaxSize(rows)
		if !isSame1dIntArray(expected, out) {
			t.Errorf("colMaxSize is something wrong. \n expected: %v \n output  : %v", expected, out)
		}
	})
}

func isSame1dIntArray(arr1, arr2 []int) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	for i, v := range arr1 {
		if v != arr2[i] {
			return false
		}
	}
	return true
}

func isSame2dStringArray(arr1, arr2 [][]string) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	for i, v := range arr1 {
		if len(v) != len(arr2[i]) {
			return false
		}
		for j, x := range v {
			if x != arr2[i][j] {
				return false
			}
		}
	}
	return true
}

func readFile(filename string) string {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(string(bytes))
}
