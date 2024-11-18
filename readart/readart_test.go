package readart

import "testing"

func TestReadArt(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		wantArt string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotArt := ReadArt(tt.args.filename); gotArt != tt.wantArt {
				t.Errorf("ReadArt() = %v, want %v", gotArt, tt.wantArt)
			}
		})
	}
}
