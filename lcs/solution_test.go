package lcs

import "testing"

func Test_longestConsecutive(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "[100,4,200,1,3,2]",
			args: args{nums: []int{100, 4, 200, 1, 3, 2}},
			want: 4,
		},
		{
			name: "[0,3,7,2,5,8,4,6,0,1]",
			args: args{nums: []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}},
			want: 9,
		},
		{
			name: "[0]",
			args: args{nums: []int{0}},
			want: 1,
		},
		{
			name: "[0, 0]",
			args: args{nums: []int{0, 0}},
			want: 1,
		},
		{
			name: "[0, -1]",
			args: args{nums: []int{0, -1}},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestConsecutive(tt.args.nums); got != tt.want {
				t.Errorf("longestConsecutive() = %v, want %v", got, tt.want)
			}
		})
	}
}
