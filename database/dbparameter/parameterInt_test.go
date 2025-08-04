package dbparameter

import "testing"

func TestIntParameter_VerifyValue(t *testing.T) {
	type fields struct {
		Name         string
		DisplayName  string
		DefaultValue int
		MinValue     int
		MaxValue     int
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
			name: "Valid int value",
			fields: fields{
				Name:         "test_int",
				DisplayName:  "Test Int",
				DefaultValue: 42,
				MinValue:     1,
				MaxValue:     100,
			},
			args: args{
				value: "50",
			},
			wantErr: false,
		},
		{
			name: "Int value below min",
			fields: fields{
				Name:         "test_int",
				DisplayName:  "Test Int",
				DefaultValue: 42,
				MinValue:     1,
				MaxValue:     100,
			},
			args: args{
				value: "0",
			},
			wantErr: true,
		},
		{
			name: "Int value above max",
			fields: fields{
				Name:         "test_int",
				DisplayName:  "Test Int",
				DefaultValue: 42,
				MinValue:     1,
				MaxValue:     100,
			},
			args: args{
				value: "101",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := IntParameter{
				Name:         tt.fields.Name,
				DisplayName:  tt.fields.DisplayName,
				DefaultValue: tt.fields.DefaultValue,
				MinValue:     tt.fields.MinValue,
				MaxValue:     tt.fields.MaxValue,
			}
			if err := p.VerifyValue(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("IntParameter.VerifyValue() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
