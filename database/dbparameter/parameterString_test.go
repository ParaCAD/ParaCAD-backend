package dbparameter

import "testing"

func TestStringParameter_VerifyValue(t *testing.T) {
	type fields struct {
		Name         string
		DisplayName  string
		DefaultValue string
		MinLength    int
		MaxLength    int
	}
	type args struct {
		value string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Valid string length",
			fields: fields{
				Name:         "test_string",
				DisplayName:  "Test String",
				DefaultValue: "default",
				MinLength:    3,
				MaxLength:    10,
			},
			args: args{
				value: "valid",
			},
			wantErr: false,
		},
		{
			name: "String too short",
			fields: fields{
				Name:         "test_string",
				DisplayName:  "Test String",
				DefaultValue: "default",
				MinLength:    3,
				MaxLength:    10,
			},
			args: args{
				value: "no",
			},
			wantErr: true,
		},
		{
			name: "String too long",
			fields: fields{
				Name:         "test_string",
				DisplayName:  "Test String",
				DefaultValue: "default",
				MinLength:    3,
				MaxLength:    10,
			},
			args: args{
				value: "this is too long",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := StringParameter{
				Name:         tt.fields.Name,
				DisplayName:  tt.fields.DisplayName,
				DefaultValue: tt.fields.DefaultValue,
				MinLength:    tt.fields.MinLength,
				MaxLength:    tt.fields.MaxLength,
			}
			if err := p.VerifyValue(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("StringParameter.VerifyValue() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
