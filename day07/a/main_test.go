// https://adventofcode.com/2022/day/7
// Day 7: No Space Left On Device
package main

import (
	"aoc/libaoc"
	"testing"
)

func Test_getDirSize(t *testing.T) {
	type args struct {
		dirname string
	}
	input, err := libaoc.ReadLines("example.txt")
	// input, err := libaoc.ReadLines("input.txt")
	if err != nil {
		panic("No input!")
	}
	initialize(input)
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "e",
			args: args{
				dirname: "e",
			},
			want: 584,
		},
		{
			name: "a",
			args: args{
				dirname: "a",
			},
			want: 94853,
		},
		{
			name: "d",
			args: args{
				dirname: "d",
			},
			want: 24933642,
		},
		{
			name: "/",
			args: args{
				dirname: "/",
			},
			want: 48381165,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getDirSize(tt.args.dirname); got != tt.want {
				t.Errorf("getDirSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getSumDirsUnder100k(t *testing.T) {
	input, err := libaoc.ReadLines("example.txt")
	// input, err := libaoc.ReadLines("input.txt")
	if err != nil {
		panic("No input!")
	}
	initialize(input)
	tests := []struct {
		name    string
		wantSum int
	}{
		{
			name:    "All dirs under 100k",
			wantSum: 95437,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSum := getSumDirsUnder100k(); gotSum != tt.wantSum {
				t.Errorf("getSumDirsUnder100k() = %v, want %v", gotSum, tt.wantSum)
			}
		})
	}
}
