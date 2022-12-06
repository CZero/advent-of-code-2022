// https://adventofcode.com/2022/day/
// Day x:
package main

import (
	"testing"
)

func Test_findStartMarkers(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				input: "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			},
			want: 7,
		},
		{
			args: args{
				input: "bvwbjplbgvbhsrlpgdmjqwftvncz",
			},
			want: 5,
		},
		{
			args: args{
				input: "nppdvjthqldpwncqszvftbrmjlhg",
			},
			want: 6,
		},
		{
			args: args{
				input: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			},
			want: 10,
		},
		{
			args: args{
				input: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			},
			want: 11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findStartMarkers(tt.args.input); got != tt.want {
				t.Errorf("findStartMarkers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_containsDoubles(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				input: []string{"a", "a", "b", "c"},
			},
			want: true,
		},
		{
			args: args{
				input: []string{"a", "d", "b", "c"},
			},
			want: false,
		},
		{
			args: args{
				input: []string{"a", "b", "a", "c"},
			},
			want: true,
		},
		{
			args: args{
				input: []string{"b", "d", "c", "b"},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containsDoubles(tt.args.input); got != tt.want {
				t.Errorf("containsDoubles() = %v, want %v", got, tt.want)
			}
		})
	}
}
