package controller

import (
	"reflect"
	"testing"

	"github.com/ParaCAD/ParaCAD-backend/database"
	"github.com/ParaCAD/ParaCAD-backend/database/dbparameter"
	"github.com/ParaCAD/ParaCAD-backend/generator"
	"github.com/google/uuid"
)

func Test_validateCreateTemplateRequest(t *testing.T) {
	type args struct {
		request CreateTemplateRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Valid request",
			args: args{
				request: CreateTemplateRequest{
					TemplateName:        "Valid Template",
					TemplateDescription: "This is a valid template description.",
					TemplateContent:     "This is the content of the template.",
					Parameters: []CreateTemplateRequestParameter{
						{
							ParameterDisplayName:  "Valid Parameter",
							ParameterName:         "valid_parameter",
							ParameterType:         "string",
							ParameterDefaultValue: "default_value",
							ParameterConstraints: []CreateTemplateRequestParameterConstraint{
								{
									Type:  "minLen",
									Value: "3",
								},
								{
									Type:  "maxLen",
									Value: "100",
								},
							},
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Name too short",
			args: args{
				request: CreateTemplateRequest{
					TemplateName:        "Va",
					TemplateDescription: "This is a valid template description.",
					TemplateContent:     "This is the content of the template.",
					Parameters: []CreateTemplateRequestParameter{
						{
							ParameterDisplayName:  "Valid Parameter",
							ParameterName:         "valid_parameter",
							ParameterType:         "string",
							ParameterDefaultValue: "default_value",
							ParameterConstraints: []CreateTemplateRequestParameterConstraint{
								{
									Type:  "minLen",
									Value: "3",
								},
								{
									Type:  "maxLen",
									Value: "100",
								},
							},
						},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "Description too short",
			args: args{
				request: CreateTemplateRequest{
					TemplateName:        "Valid Template",
					TemplateDescription: "This",
					TemplateContent:     "This is the content of the template.",
					Parameters: []CreateTemplateRequestParameter{
						{
							ParameterDisplayName:  "Valid Parameter",
							ParameterName:         "valid_parameter",
							ParameterType:         "string",
							ParameterDefaultValue: "default_value",
							ParameterConstraints: []CreateTemplateRequestParameterConstraint{
								{
									Type:  "minLen",
									Value: "3",
								},
								{
									Type:  "maxLen",
									Value: "100",
								},
							},
						},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "No content",
			args: args{
				request: CreateTemplateRequest{
					TemplateName:        "Valid Template",
					TemplateDescription: "This is a valid template description.",
					TemplateContent:     "",
					Parameters: []CreateTemplateRequestParameter{
						{
							ParameterDisplayName:  "Valid Parameter",
							ParameterName:         "valid_parameter",
							ParameterType:         "string",
							ParameterDefaultValue: "default_value",
							ParameterConstraints: []CreateTemplateRequestParameterConstraint{
								{
									Type:  "minLen",
									Value: "3",
								},
								{
									Type:  "maxLen",
									Value: "100",
								},
							},
						},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "No parameters",
			args: args{
				request: CreateTemplateRequest{
					TemplateName:        "Valid Template",
					TemplateDescription: "This is a valid template description.",
					TemplateContent:     "This is the content of the template.",
					Parameters:          []CreateTemplateRequestParameter{},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateCreateTemplateRequest(tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("validateCreateTemplateRequest() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_createDefaultFilledTemplate(t *testing.T) {
	type args struct {
		template database.Template
	}
	tests := []struct {
		name string
		args args
		want generator.FilledTemplate
	}{
		{
			name: "Empty template",
			args: args{
				template: database.Template{
					UUID:        uuid.Nil,
					Name:        "Empty Template",
					Description: "This is an empty template.",
					Template:    "",
					Parameters:  []dbparameter.Parameter{},
				},
			},
			want: generator.FilledTemplate{
				UUID:     uuid.Nil,
				Template: []byte(""),
			},
		},
		{
			name: "Filled template",
			args: args{
				template: database.Template{
					UUID:        uuid.Nil,
					Name:        "Filled Template",
					Description: "This is a filled template.",
					Template:    "This is the content of the filled template.",
					Parameters: []dbparameter.Parameter{
						dbparameter.StringParameter{
							DisplayName:  "param1",
							Name:         "param1",
							DefaultValue: "default_value",
						},
					},
				},
			},
			want: generator.FilledTemplate{
				UUID:     uuid.Nil,
				Template: []byte("This is the content of the filled template."),
				Params: []generator.Parameter{
					{
						Key:   "param1",
						Value: "default_value",
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createDefaultFilledTemplate(tt.args.template); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createDefaultFilledTemplate() = %v, want %v", got, tt.want)
			}
		})
	}
}
