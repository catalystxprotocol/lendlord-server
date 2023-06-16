package tools

import "testing"

func TestCheckFileExist(t *testing.T) {
	type args struct {
		filePath string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CheckFileExist(tt.args.filePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckFileExist() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CheckFileExist() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateFolder(t *testing.T) {
	type args struct {
		folderPath string
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
			if err := CreateFolder(tt.args.folderPath); (err != nil) != tt.wantErr {
				t.Errorf("CreateFolder() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
