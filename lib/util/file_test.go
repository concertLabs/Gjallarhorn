package util

import (
	"os"
	"testing"
)

func TestStringToFile(t *testing.T) {
	type args struct {
		file string
		data string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "default",
			args:    args{file: "testdata/tmp.txt", data: "hello world"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := StringToFile(tt.args.file, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("StringToFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}

	err := os.Remove("testdata/tmp.txt")
	if err != nil {
		t.Logf("could not delete testdata: tmp.txt: %v\n", err)
	}
}

func TestFileToString(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "default",
			args:    args{"testdata/readfiletest.txt"},
			want:    "hello world",
			wantErr: false,
		},
		{
			name:    "missing file",
			args:    args{"testdata/readfiletestmissing.txt"},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FileToString(tt.args.file)
			if (err != nil) != tt.wantErr {
				t.Errorf("FileToString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FileToString() = %v, want %v", got, tt.want)
			}
		})
	}
}
