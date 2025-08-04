package dbparameter

import "testing"

func TestFloatParameter_VerifyValue(t *testing.T) {
	type fields struct {
		Name         string
		DisplayName  string
		DefaultValue float64
		MinValue     float64
		MaxValue     float64
		Step         float64
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
			name: "Valid float value",
			fields: fields{
				Name:         "test_float",
				DisplayName:  "Test Float",
				DefaultValue: 3.14,
				MinValue:     1.0,
				MaxValue:     10.0,
				Step:         0.01,
			},
			args: args{
				value: "5.0",
			},
			wantErr: false,
		},
		{
			name: "Float value below min",
			fields: fields{
				Name:         "test_float",
				DisplayName:  "Test Float",
				DefaultValue: 3.14,
				MinValue:     1.0,
				MaxValue:     10.0,
				Step:         0.01,
			},
			args: args{
				value: "0.5",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := FloatParameter{
				Name:         tt.fields.Name,
				DisplayName:  tt.fields.DisplayName,
				DefaultValue: tt.fields.DefaultValue,
				MinValue:     tt.fields.MinValue,
				MaxValue:     tt.fields.MaxValue,
				Step:         tt.fields.Step,
			}
			if err := p.VerifyValue(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("FloatParameter.VerifyValue() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
