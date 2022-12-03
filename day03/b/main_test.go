// https://adventofcode.com/2022/day/3
// Day 3: Rucksack Reorganization
package main

import (
	"testing"
)

func Test_cutInHalf(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name  string
		args  args
		wantA string
		wantB string
	}{
		{
			name: "a",
			args: args{
				"12345678",
			},
			wantA: "1234",
			wantB: "5678",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotA, gotB := cutInHalf(tt.args.input)
			if gotA != tt.wantA {
				t.Errorf("cutInHalf() gotA = %v, want %v", gotA, tt.wantA)
			}
			if gotB != tt.wantB {
				t.Errorf("cutInHalf() gotB = %v, want %v", gotB, tt.wantB)
			}
		})
	}
}

func Test_value(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "a",
			args: args{
				"a",
			},
			want: 1,
		},
		{
			name: "A",
			args: args{
				"A",
			},
			want: 27,
		},
		{
			name: "B",
			args: args{
				"B",
			},
			want: 28,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := value(tt.args.input); got != tt.want {
				t.Errorf("value() = %v, want %v", got, tt.want)
			}
		})
	}
}
