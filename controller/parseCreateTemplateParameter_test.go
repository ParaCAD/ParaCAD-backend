package controller

import (
	"reflect"
	"testing"

	"github.com/ParaCAD/ParaCAD-backend/database/dbparameter"
)

func Test_parseFloatParameter(t *testing.T) {
	type args struct {
		parameter CreateTemplateRequestParameter
	}
	tests := []struct {
		name    string
		args    args
		want    dbparameter.FloatParameter
		wantErr bool
	}{
		{
			name: "Valid float parameter",
			args: args{
				parameter: CreateTemplateRequestParameter{
					ParameterName:         "test_float",
					ParameterDisplayName:  "Test Float",
					ParameterType:         string(dbparameter.ParameterTypeFloat),
					ParameterDefaultValue: "3.14",
					ParameterConstraints: []CreateTemplateRequestParameterConstraint{
						{
							Type:  "min_value",
							Value: "1.0",
						},
						{
							Type:  "max_value",
							Value: "10.0",
						},
						{
							Type:  "step",
							Value: "0.01",
						},
					},
				},
			},
			want: dbparameter.FloatParameter{
				Name:         "test_float",
				DisplayName:  "Test Float",
				DefaultValue: 3.14,
				MinValue:     1.0,
				MaxValue:     10.0,
				Step:         0.01,
			},
			wantErr: false,
		},
		{
			name: "No step",
			args: args{
				parameter: CreateTemplateRequestParameter{
					ParameterName:         "test_float",
					ParameterDisplayName:  "Test Float",
					ParameterType:         string(dbparameter.ParameterTypeFloat),
					ParameterDefaultValue: "3.14",
					ParameterConstraints: []CreateTemplateRequestParameterConstraint{
						{
							Type:  "min_value",
							Value: "1.0",
						},
						{
							Type:  "max_value",
							Value: "10.0",
						},
					},
				},
			},
			want:    dbparameter.FloatParameter{},
			wantErr: true,
		},
		{
			name: "No min",
			args: args{
				parameter: CreateTemplateRequestParameter{
					ParameterName:         "test_float",
					ParameterDisplayName:  "Test Float",
					ParameterType:         string(dbparameter.ParameterTypeFloat),
					ParameterDefaultValue: "3.14",
					ParameterConstraints: []CreateTemplateRequestParameterConstraint{
						{
							Type:  "max_value",
							Value: "10.0",
						},
						{
							Type:  "step",
							Value: "0.01",
						},
					},
				},
			},
			want:    dbparameter.FloatParameter{},
			wantErr: true,
		},
		{
			name: "No max",
			args: args{
				parameter: CreateTemplateRequestParameter{
					ParameterName:         "test_float",
					ParameterDisplayName:  "Test Float",
					ParameterType:         string(dbparameter.ParameterTypeFloat),
					ParameterDefaultValue: "3.14",
					ParameterConstraints: []CreateTemplateRequestParameterConstraint{
						{
							Type:  "min_value",
							Value: "1.0",
						},
						{
							Type:  "step",
							Value: "0.01",
						},
					},
				},
			},
			want:    dbparameter.FloatParameter{},
			wantErr: true,
		},
		{
			name: "Too small",
			args: args{
				parameter: CreateTemplateRequestParameter{
					ParameterName:         "test_float",
					ParameterDisplayName:  "Test Float",
					ParameterType:         string(dbparameter.ParameterTypeFloat),
					ParameterDefaultValue: "0.14",
					ParameterConstraints: []CreateTemplateRequestParameterConstraint{
						{
							Type:  "min_value",
							Value: "1.0",
						},
						{
							Type:  "max_value",
							Value: "10.0",
						},
						{
							Type:  "step",
							Value: "0.01",
						},
					},
				},
			},
			want:    dbparameter.FloatParameter{},
			wantErr: true,
		},
		{
			name: "Too big",
			args: args{
				parameter: CreateTemplateRequestParameter{
					ParameterName:         "test_float",
					ParameterDisplayName:  "Test Float",
					ParameterType:         string(dbparameter.ParameterTypeFloat),
					ParameterDefaultValue: "11.14",
					ParameterConstraints: []CreateTemplateRequestParameterConstraint{
						{
							Type:  "min_value",
							Value: "1.0",
						},
						{
							Type:  "max_value",
							Value: "10.0",
						},
						{
							Type:  "step",
							Value: "0.01",
						},
					},
				},
			},
			want:    dbparameter.FloatParameter{},
			wantErr: true,
		},
		{
			name: "Too big step",
			args: args{
				parameter: CreateTemplateRequestParameter{
					ParameterName:         "test_float",
					ParameterDisplayName:  "Test Float",
					ParameterType:         string(dbparameter.ParameterTypeFloat),
					ParameterDefaultValue: "3.14",
					ParameterConstraints: []CreateTemplateRequestParameterConstraint{
						{
							Type:  "min_value",
							Value: "1.0",
						},
						{
							Type:  "max_value",
							Value: "10.0",
						},
						{
							Type:  "step",
							Value: "1000.0",
						},
					},
				},
			},
			want:    dbparameter.FloatParameter{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseFloatParameter(tt.args.parameter)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseFloatParameter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseFloatParameter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseIntParameter(t *testing.T) {
	type args struct {
		parameter CreateTemplateRequestParameter
	}
	tests := []struct {
		name    string
		args    args
		want    dbparameter.IntParameter
		wantErr bool
	}{
		{
			name: "Valid int parameter",
			args: args{
				parameter: CreateTemplateRequestParameter{
					ParameterName:         "test_int",
					ParameterDisplayName:  "Test Int",
					ParameterType:         string(dbparameter.ParameterTypeInt),
					ParameterDefaultValue: "42",
					ParameterConstraints: []CreateTemplateRequestParameterConstraint{
						{
							Type:  "min_value",
							Value: "1",
						},
						{
							Type:  "max_value",
							Value: "100",
						},
					},
				},
			},
			want: dbparameter.IntParameter{
				Name:         "test_int",
				DisplayName:  "Test Int",
				DefaultValue: 42,
				MinValue:     1,
				MaxValue:     100,
			},
			wantErr: false,
		},
		{
			name: "No min",
			args: args{
				parameter: CreateTemplateRequestParameter{
					ParameterName:         "test_int",
					ParameterDisplayName:  "Test Int",
					ParameterType:         string(dbparameter.ParameterTypeInt),
					ParameterDefaultValue: "42",
					ParameterConstraints: []CreateTemplateRequestParameterConstraint{
						{
							Type:  "max_value",
							Value: "100",
						},
					},
				},
			},
			want:    dbparameter.IntParameter{},
			wantErr: true,
		},
		{
			name: "No max",
			args: args{
				parameter: CreateTemplateRequestParameter{
					ParameterName:         "test_int",
					ParameterDisplayName:  "Test Int",
					ParameterType:         string(dbparameter.ParameterTypeInt),
					ParameterDefaultValue: "42",
					ParameterConstraints: []CreateTemplateRequestParameterConstraint{
						{
							Type:  "min_value",
							Value: "1",
						},
					},
				},
			},
			want:    dbparameter.IntParameter{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseIntParameter(tt.args.parameter)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseIntParameter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseIntParameter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseStringParameter(t *testing.T) {
	type args struct {
		parameter CreateTemplateRequestParameter
	}
	tests := []struct {
		name    string
		args    args
		want    dbparameter.StringParameter
		wantErr bool
	}{
		{
			name: "Valid string parameter",
			args: args{
				parameter: CreateTemplateRequestParameter{
					ParameterName:         "test_string",
					ParameterDisplayName:  "Test String",
					ParameterType:         string(dbparameter.ParameterTypeString),
					ParameterDefaultValue: "default",
					ParameterConstraints: []CreateTemplateRequestParameterConstraint{
						{
							Type:  "min_length",
							Value: "3",
						},
						{
							Type:  "max_length",
							Value: "10",
						},
					},
				},
			},
			want: dbparameter.StringParameter{
				Name:         "test_string",
				DisplayName:  "Test String",
				DefaultValue: "default",
				MinLength:    3,
				MaxLength:    10,
			},
			wantErr: false,
		},
		{
			name: "Bad min length",
			args: args{
				parameter: CreateTemplateRequestParameter{
					ParameterName:         "test_string",
					ParameterDisplayName:  "Test String",
					ParameterType:         string(dbparameter.ParameterTypeString),
					ParameterDefaultValue: "default",
					ParameterConstraints: []CreateTemplateRequestParameterConstraint{
						{
							Type:  "min_length",
							Value: "13",
						},
						{
							Type:  "max_length",
							Value: "100",
						},
					},
				},
			},
			want:    dbparameter.StringParameter{},
			wantErr: true,
		},
		{
			name: "Bad max length",
			args: args{
				parameter: CreateTemplateRequestParameter{
					ParameterName:         "test_string",
					ParameterDisplayName:  "Test String",
					ParameterType:         string(dbparameter.ParameterTypeString),
					ParameterDefaultValue: "default",
					ParameterConstraints: []CreateTemplateRequestParameterConstraint{
						{
							Type:  "min_length",
							Value: "1",
						},
						{
							Type:  "max_length",
							Value: "3",
						},
					},
				},
			},
			want:    dbparameter.StringParameter{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseStringParameter(tt.args.parameter)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseStringParameter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseStringParameter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseBoolParameter(t *testing.T) {
	type args struct {
		parameter CreateTemplateRequestParameter
	}
	tests := []struct {
		name    string
		args    args
		want    dbparameter.BoolParameter
		wantErr bool
	}{
		{
			name: "Valid bool parameter",
			args: args{
				parameter: CreateTemplateRequestParameter{
					ParameterName:         "test_bool",
					ParameterDisplayName:  "Test Bool",
					ParameterType:         string(dbparameter.ParameterTypeBool),
					ParameterDefaultValue: "true",
				},
			},
			want: dbparameter.BoolParameter{
				Name:         "test_bool",
				DisplayName:  "Test Bool",
				DefaultValue: true,
			},
			wantErr: false,
		},
		{
			name: "Invalid bool parameter (bad default value)",
			args: args{
				parameter: CreateTemplateRequestParameter{
					ParameterName:         "test_bool",
					ParameterDisplayName:  "Test Bool",
					ParameterType:         string(dbparameter.ParameterTypeBool),
					ParameterDefaultValue: "Moim zdaniem to nie ma tak, że dobrze albo że nie dobrze.",
				},
			},
			want:    dbparameter.BoolParameter{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseBoolParameter(tt.args.parameter)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseBoolParameter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseBoolParameter() = %v, want %v", got, tt.want)
			}
		})
	}
}
