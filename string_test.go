package csv

import (
	"testing"
)

type containsArgs struct {
	s string
	v string
}

var containsTests = []struct {
	name string
	args containsArgs
	want bool
}{
	{
		name: "empty_false",
		args: containsArgs{
			s: "",
			v: "",
		},
		want: false,
	},
	{
		name: "empty_true",
		args: containsArgs{
			s: ",", // 2 empty string values
			v: "",
		},
		want: true,
	},
}

func TestContains(t *testing.T) {
	for _, tt := range containsTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Contains(tt.args.s, tt.args.v, ","); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestString_Add(t *testing.T) {
	tests := []struct {
		name string
		s    String
		v    []string
		want String
	}{
		{
			name: "single_to_empty",
			s:    "",
			v:    []string{"a"},
			want: "a",
		},
		{
			name: "few_to_empty",
			s:    "",
			v:    []string{"a", "b", "c"},
			want: "a,b,c",
		},
		{
			name: "few_to_few",
			s:    "a,b",
			v:    []string{"c", "d", "e"},
			want: "a,b,c,d,e",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Add(tt.v...); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestString_Contains(t *testing.T) {
	for _, tt := range containsTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := String(tt.args.s).Contains(tt.args.v); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestString_Set(t *testing.T) {
	type args struct {
		i int
		v string
	}
	tests := []struct {
		name string
		s    String
		args args
		want String
	}{
		{
			name: "set 0",
			s:    "a,b,c",
			args: args{
				i: 0,
				v: "x",
			},
			want: "x,b,c",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Set(tt.args.i, tt.args.v); got != tt.want {
				t.Errorf("Set() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestString_Values(t *testing.T) {
	tests := []struct {
		name string
		s    String
		want []string
	}{
		{name: "empty", s: "", want: nil},
		{name: "abc", s: "a,b,c", want: []string{"a", "b", "c"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.s.Values()
			different := len(got) != len(tt.want)
			if !different {
				for i := range got {
					if got[i] != tt.want[i] {
						different = true
						break
					}
				}
			}
			if different {
				t.Errorf("want %+v, got %+v", tt.want, got)
			}
		})
	}
}

func TestString_Remove(t *testing.T) {
	tests := []struct {
		name string
		s    String
		v    []string
		want String
	}{
		{"empty", "", []string{}, ""},
		{"first", "a,b,c", []string{"a"}, "b,c"},
		{"last", "a,b,c", []string{"c"}, "a,b"},
		{"middle", "a,b,c", []string{"b"}, "a,c"},
		{"few", "a,b,c,a,a,d,e,f", []string{"a"}, "b,c,d,e,f"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Remove(tt.v...); got != tt.want {
				t.Errorf("Remove() = %v, want %v", got, tt.want)
			}
		})
	}
}
