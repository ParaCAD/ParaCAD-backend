package utils

import (
	"reflect"
	"testing"
)

func TestGetPtr(t *testing.T) {
	type args struct {
		value int
	}
	tests := []struct {
		name string
		args args
		want *int
	}{
		{
			name: "Get pointer of int",
			args: args{
				value: 42,
			},
			want: func() *int {
				i := 42
				return &i
			}(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetPtr(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPtr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValueOrEmpty(t *testing.T) {
	type args struct {
		value *int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Value is nil",
			args: args{
				value: nil,
			},
			want: 0,
		},
		{
			name: "Value is set",
			args: args{
				value: func() *int {
					i := 100
					return &i
				}(),
			},
			want: 100,
		},
		{
			name: "Value is zero",
			args: args{
				value: func() *int {
					i := 0
					return &i
				}(),
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValueOrEmpty(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ValueOrEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValueOrDefault(t *testing.T) {
	type args struct {
		value *int
		def   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Value is nil",
			args: args{
				value: nil,
				def:   42,
			},
			want: 42,
		},
		{
			name: "Value is set",
			args: args{
				value: func() *int {
					i := 100
					return &i
				}(),
				def: 42,
			},
			want: 100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValueOrDefault(tt.args.value, tt.args.def); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ValueOrDefault() = %v, want %v", got, tt.want)
			}
		})
	}
}
