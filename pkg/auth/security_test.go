package auth

import "testing"

func TestComparePassword(t *testing.T) {
	type args struct {
		plainTextPassword string
		encryptedPassword string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Validate that matching passwords return true",
			args: args{
				plainTextPassword: "admin123456",
				encryptedPassword: HashAndSalt("admin123456"),
			},
			want: true,
		},
		{
			name: "Validate that the wrong password does not work",
			args: args{
				plainTextPassword: "admin123456",
				encryptedPassword: HashAndSalt("notTheRightPassword90"),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ComparePassword(tt.args.plainTextPassword, tt.args.encryptedPassword); got != tt.want {
				t.Errorf("ComparePassword() = %v, want %v", got, tt.want)
			}
		})
	}
}
