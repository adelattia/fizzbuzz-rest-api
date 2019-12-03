package core

import (
	"errors"
	"reflect"
	"testing"
)

func TestFizzBuzzNaive(t *testing.T) {
	type args struct {
		int1    int
		int2    int
		string1 string
		string2 string
		limit   int
	}
	tests := []struct {
		name  string
		args  args
		want  []string
		want1 error
		want2 uint
	}{
		{
			name: "success", args: args{int1: 3, int2: 5, string1: "Fizz", string2: "Buzz", limit: 15},
			want:  []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"},
			want1: nil,
			want2: 0,
		},
		{
			name: "success", args: args{int1: 3, int2: 3, string1: "Fizz", string2: "Buzz", limit: 15},
			want:  []string{},
			want1: errors.New("You should provide different multiples"),
			want2: 1,
		},
		{
			name: "success", args: args{int1: 55, int2: 3, string1: "Fizz", string2: "Buzz", limit: 15},
			want:  []string{},
			want1: errors.New("Multiples should be less than limit"),
			want2: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := FizzBuzzNaive(tt.args.int1, tt.args.int2, tt.args.string1, tt.args.string2, tt.args.limit)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FizzBuzzNaive() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("FizzBuzzNaive() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("FizzBuzzNaive() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}
