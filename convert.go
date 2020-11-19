package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

const (
	VERTICAL_DEVIDER   = "|"
	HORIZONTAL_DEVIDER = "-"
)

func ConvertAll(args *Args) error {
	for _, file := range args.Files {
		md, err := Convert(file, args)
		if err != nil {
			return err
		}
		// print markdown
		fmt.Println(md)
	}
	return nil
}

func Convert(file string, args *Args) (string, error) {
	records, err := CsvToArray(file, args.Delim)
	if err != nil {
		return "", err
	}
	if len(records) < 1 { // when csv is empty
		return "", nil
	}

	// modify content
	records = Modify(records)

	// array to markdown
	md, err := ArrayToMd(records, args)
	if err != nil {
		return "", err
	}

	return md, nil
}

func CsvToArray(file, delim string) ([][]string, error) {
	fp, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer fp.Close()

	reader := csv.NewReader(fp)
	reader.Comma = []rune(delim)[0]
	reader.LazyQuotes = true

	return reader.ReadAll()
}

func ArrayToMd(records [][]string, args *Args) (string, error) {
	colSizes := colMaxSize(records)

	records, err := padCells(records, colSizes)
	if err != nil {
		return "", err
	}

	horiz := createHorizontalDivder(colSizes, args.Pad)
	rows := concateWithSep(records, args.Pad)

	md := []string{}
	md = append(md, rows[0]) // header
	md = append(md, horiz)   // horizontal devider
	for _, v := range rows[1:] {
		md = append(md, v)
	}

	return strings.Join(md, "\n"), nil
}

func Modify(records [][]string) [][]string {
	for i, row := range records {
		for j, ele := range row {
			records[i][j] = strings.ReplaceAll(ele, "\n", "<br/>")
		}
	}
	return records
}

func padCells(records [][]string, colSizes []int) ([][]string, error) {
	for i, row := range records {
		for j, v := range row {
			if diff := colSizes[j] - utf8.RuneCountInString(v); diff >= 0 {
				records[i][j] = v + strings.Repeat(" ", diff)
			} else {
				fmt.Println(v)
				fmt.Println(colSizes[j])
				fmt.Println(utf8.RuneCountInString(v))
				return nil, fmt.Errorf("Internal error: column size is bigger than max.")
			}
		}
	}
	return records, nil
}

func colMaxSize(records [][]string) []int {
	var sizes []int = make([]int, len(records[0]))
	for _, row := range records {
		for j, v := range row {
			if cur := utf8.RuneCountInString(v); sizes[j] < cur {
				sizes[j] = cur
			}
		}
	}
	return sizes
}

func createHorizontalDivder(colSizes []int, pad int) string {
	var dividers []string
	for _, v := range colSizes {
		dividers = append(dividers, strings.Repeat(HORIZONTAL_DEVIDER, v))
	}
	sep := strings.Repeat(HORIZONTAL_DEVIDER, pad) + VERTICAL_DEVIDER + strings.Repeat(HORIZONTAL_DEVIDER, pad)
	leftside := VERTICAL_DEVIDER + strings.Repeat("-", pad)
	rightside := strings.Repeat("-", pad) + VERTICAL_DEVIDER

	return leftside + strings.Join(dividers, sep) + rightside
}

func concateWithSep(records [][]string, pad int) []string {
	var rows []string
	sep := strings.Repeat(" ", pad) + VERTICAL_DEVIDER + strings.Repeat(" ", pad)
	leftside := VERTICAL_DEVIDER + strings.Repeat(" ", pad)
	rightside := strings.Repeat(" ", pad) + VERTICAL_DEVIDER
	for _, row := range records {
		rows = append(rows, leftside+strings.Join(row, sep)+rightside)
	}
	return rows
}
