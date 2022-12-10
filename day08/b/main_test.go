// https://adventofcode.com/2022/day/8
// Day 8: Treetop Tree House
package main

import (
	"oac/libaoc"
	"testing"
)

func Test_calcScenicScores(t *testing.T) {
	type args struct {
		coord Coord
	}
	input, err := libaoc.ReadLines("example.txt")
	if err != nil {
		panic("No input!")
	}
	gridheight = len(input)
	gridwidth = len(input[0])
	buildGrid(input)
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Boom 1",
			args: args{
				Coord{2, 3},
			},
			want: 4,
		},
		{
			name: "Boom 2",
			args: args{
				Coord{2, 1},
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calcScenicScores(tt.args.coord); got != tt.want {
				t.Errorf("calcScenicScores() = %v, want %v", got, tt.want)
			}
		})
	}
}
