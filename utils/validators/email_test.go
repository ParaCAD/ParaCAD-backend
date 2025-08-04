package validators

import "testing"

func TestEmail(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Valid email",
			args: args{
				email: "test@test.org",
			},
			wantErr: false,
		},
		{
			name: "No @ symbol",
			args: args{
				email: "testtest.org",
			},
			wantErr: true,
		},
		{
			name: "No domain",
			args: args{
				email: "test@",
			},
			wantErr: true,
		},
		{
			name: "no user",
			args: args{
				email: "@test.org",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Email(tt.args.email); (err != nil) != tt.wantErr {
				t.Errorf("Email() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
