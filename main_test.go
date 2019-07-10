package main

import "testing"

func Test_fizzbuzz(t *testing.T) {
	type args struct {
		a int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "15",
			args: args{a: 15},
			want: "fizzbuzz",
		},
		{
			name: "5",
			args: args{a: 5},
			want: "buzz",
		},
		{
			name: "3",
			args: args{a: 3},
			want: "fizz",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fizzbuzz(tt.args.a); got != tt.want {
				t.Errorf("fizzbuzz() = %v, want %v", got, tt.want)
			}
		})
	}
}
