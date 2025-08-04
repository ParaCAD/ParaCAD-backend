package openscadgenerator

import (
	"reflect"
	"testing"
)

func Test_buildArgs(t *testing.T) {
	type args struct {
		format       string
		outputFile   string
		scadFilePath string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "STL format",
			args: args{
				format:       FormatSTL,
				outputFile:   "output.stl",
				scadFilePath: "model.scad",
			},
			want: []string{"--export-format=binstl", "-ooutput.stl", "model.scad"},
		},
		{
			name: "PNG format",
			args: args{
				format:       FormatPNG,
				outputFile:   "output.png",
				scadFilePath: "model.scad",
			},
			want: []string{"--export-format=png", "--colorscheme=BeforeDawn", "-ooutput.png", "model.scad"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildArgs(tt.args.format, tt.args.outputFile, tt.args.scadFilePath); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildArgs() = %v, want %v", got, tt.want)
			}
		})
	}
}
