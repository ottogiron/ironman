package ironman

import "testing"

func TestInitIronmanHome(t *testing.T) {
	type args struct {
		ironmanHome string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := InitIronmanHome(tt.args.ironmanHome); (err != nil) != tt.wantErr {
				t.Errorf("InitIronmanHome() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
