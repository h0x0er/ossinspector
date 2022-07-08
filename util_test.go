package ossinspector

import "testing"

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
		// {name: "test greater", args: args{val: ">100"}, want: GREATER_THAN, want1: 100},
		{name: "test lesser", args: args{val: "<100"}, want: LESSER_THAN, want1: 100},
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
