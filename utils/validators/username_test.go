package validators

import "testing"

func TestUsername(t *testing.T) {
	type args struct {
		username string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Valid username",
			args: args{
				username: "validUser123",
			},
			wantErr: false,
		},
		{
			name: "Too short username",
			args: args{
				username: "ab",
			},
			wantErr: true,
		},
		{
			name: "Too long username",
			args: args{
				username: "thisusernameiswaytoolong",
			},
			wantErr: true,
		},
		{
			name: "Username with special characters",
			args: args{
				username: "invalid@user!",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Username(tt.args.username); (err != nil) != tt.wantErr {
				t.Errorf("Username() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
