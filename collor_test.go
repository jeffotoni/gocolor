package gcolor

import (
	"testing"
)

func TestColor_Cyan(t *testing.T) {
	type args struct {
		msg string
	}
	tests := []struct {
		name string
		c    Color
		args args
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Color{}
			if got := c.Cyan(tt.args.msg); got != tt.want {
				t.Errorf("Color.Cyan() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestColor_Purple(t *testing.T) {
	type args struct {
		msg string
	}
	tests := []struct {
		name string
		c    Color
		args args
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Color{}
			if got := c.Purple(tt.args.msg); got != tt.want {
				t.Errorf("Color.Purple() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestColor_Blue(t *testing.T) {
	type args struct {
		msg string
	}
	tests := []struct {
		name string
		c    Color
		args args
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Color{}
			if got := c.Blue(tt.args.msg); got != tt.want {
				t.Errorf("Color.Blue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestColor_Yellow(t *testing.T) {
	type args struct {
		msg string
	}
	tests := []struct {
		name string
		c    Color
		args args
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Color{}
			if got := c.Yellow(tt.args.msg); got != tt.want {
				t.Errorf("Color.Yellow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestColor_Green(t *testing.T) {
	type args struct {
		msg string
	}
	tests := []struct {
		name string
		c    Color
		args args
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Color{}
			if got := c.Green(tt.args.msg); got != tt.want {
				t.Errorf("Color.Green() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestColor_Red(t *testing.T) {
	type args struct {
		msg string
	}
	tests := []struct {
		name string
		c    Color
		args args
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Color{}
			if got := c.Red(tt.args.msg); got != tt.want {
				t.Errorf("Color.Red() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestColor_Black(t *testing.T) {
	type args struct {
		msg string
	}
	tests := []struct {
		name string
		c    Color
		args args
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Color{}
			if got := c.Black(tt.args.msg); got != tt.want {
				t.Errorf("Color.Black() = %v, want %v", got, tt.want)
			}
		})
	}
}
