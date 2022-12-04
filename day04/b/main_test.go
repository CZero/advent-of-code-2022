// https://adventofcode.com/2022/day/4
// Day 4: Camp Cleanup
package main

import "testing"

func Test_isContained(t *testing.T) {
	type args struct {
		elves []elf
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				[]elf{
					{section{2, 4}},
					{section{6, 8}},
				},
			},
			want: false,
		},
		{
			name: "1r",
			args: args{
				[]elf{
					{section{6, 8}},
					{section{2, 4}},
				},
			},
			want: false,
		},
		{
			name: "2",
			args: args{
				[]elf{
					{section{2, 3}},
					{section{4, 5}},
				},
			},
			want: false,
		},
		{
			name: "3",
			args: args{
				[]elf{
					{section{5, 7}},
					{section{7, 9}},
				},
			},
			want: true,
		},
		{
			name: "3r",
			args: args{
				[]elf{
					{section{7, 9}},
					{section{5, 7}},
				},
			},
			want: true,
		},
		{
			name: "4",
			args: args{
				[]elf{
					{section{2, 8}},
					{section{3, 6}},
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isContained(tt.args.elves); got != tt.want {
				t.Errorf("isContained() = %v, want %v", got, tt.want)
			}
		})
	}
}
