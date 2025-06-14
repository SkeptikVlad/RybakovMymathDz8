package RybakovMymathDz8

import (
	"math"
	"testing"
)

func TestFactorialBig(t *testing.T) {
	tests := []struct {
		name string
		n    uint
		want string
	}{
		{"0!", 0, "1"},
		{"1!", 1, "1"},
		{"5!", 5, "120"},
		{"10!", 10, "3628800"},
		{"20!", 20, "2432902008176640000"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FactorialBig(tt.n).String()
			if got != tt.want {
				t.Errorf("FactorialBig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRTailorDegree(t *testing.T) {
	tests := []struct {
		name    string
		a       float64
		b       float64
		want    float64
		wantErr bool //выводит неопределенность или ошибку или все окаееей
	}{
		{"2^3", 2, 3, 3 * math.Log(2), false},
		{"5^0", 5, 0, 0 * math.Log(5), false},
		{"0^0", 0, 0, 0, true},
		{"0^-1", 0, -1, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RTailorDegree(tt.a, tt.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("RTailorDegree() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && math.Abs(got-tt.want) > 0.0000000001 {
				t.Errorf("RTailorDegree() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFuncRTailorDegree(t *testing.T) {
	tests := []struct {
		name    string
		a       float64
		n       float64
		m       float64
		want    float64
		wantErr bool // выводит неопределенность или ошибку или все окаееей
	}{
		{"8^(2/3)", 8, 2, 3, (2.0 / 3.0) * math.Log(8), false},
		{"-8^(1/3)", -8, 1, 3, (1.0 / 3.0) * math.Log(8), false},
		{"0^(0/0)", 0, 0, 0, 0, true},
		{"2^(1/-2)", 2, 1, -2, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FuncRTailorDegree(tt.a, tt.n, tt.m, RootDegree)
			if (err != nil) != tt.wantErr {
				t.Errorf("FuncRTailorDegree() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && math.Abs(got-tt.want) > 1e-9 {
				t.Errorf("FuncRTailorDegree() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLnLimitations(t *testing.T) {
	tests := []struct {
		name    string
		a       float64
		b       float64
		wantErr bool //выводит неопределенность или ошибку или все окаееей
	}{
		{"Log_2 10", 10, 2, false},
		{"Log_0.5 0.5", 0.5, 0.5, false},
		{"Log_2 -1", -1, 2, true},
		{"Log_1 10", 10, 1, true},
		{"Log_0 10", 10, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, _, err := LnLimitations(tt.a, tt.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("LnLimitations() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
