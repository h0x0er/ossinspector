package ossinspector

import (
	"testing"
)

func Test_evaluate(t *testing.T) {
	type args struct {
		val string
	}
	tests := []struct {
		name  string
		args  args
		want  ConstraintType
		want1 int
	}{
		// TODO: Add test cases.
		{name: "test greater", args: args{val: ">100"}, want: GREATER_THAN, want1: 100},
		{name: "test lesser", args: args{val: "<100"}, want: LESSER_THAN, want1: 100},
		{name: "month lesser", args: args{val: "<100m"}, want: MONTHS_LESSER_THAN, want1: 100},
		{name: "year lesser", args: args{val: "<100y"}, want: YEARS_LESSER_THAN, want1: 100},
		{name: "days greater", args: args{val: ">100d"}, want: DAYS_GREATER_THAN, want1: 100},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := evaluate(tt.args.val)
			if got != tt.want {
				t.Errorf("evaluate() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("evaluate() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_isAge(t *testing.T) {
	type args struct {
		expr string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{name: "test for days", args: args{expr: ">10d"}, want: true},
		{name: "test for years", args: args{expr: ">10y"}, want: true},
		{name: "test for months", args: args{expr: ">10m"}, want: true},
		{name: "test for days", args: args{expr: ">10a"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAge(tt.args.expr); got != tt.want {
				t.Errorf("isAge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_trim(t *testing.T) {
	type args struct {
		expr string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{name: "Extract from month", args: args{expr: "<10m"}, want: 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trim(tt.args.expr); got != tt.want {
				t.Errorf("trim() = %v, want %v", got, tt.want)
			}
		})
	}
}
