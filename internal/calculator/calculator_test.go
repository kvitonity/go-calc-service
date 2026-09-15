package calculator

import "testing"

func TestCalculate(t *testing.T) {
	tests := []struct {
		name     string
		op1      float64
		operator string
		op2      float64
		want     float64
		wantErr  bool
	}{
		{"addition", 2, "+", 4, 6, false},
		{"subtraction", 10, "-", 3, 7, false},
		{"multiplication", 3, "*", 4, 12, false},
		{"division", 10, "/", 2, 5, false},
		{"division by zero", 10, "/", 0, 0, true},
		{"unsupported operator", 10, "%", 2, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Calculate(tt.op1, tt.operator, tt.op2)
			if (err != nil) != tt.wantErr {
				t.Fatalf("Calculate() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}
