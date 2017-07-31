package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func Test_openFile(t *testing.T) {
	f, err := ioutil.TempFile("", "test")
	if err != nil {
		t.Errorf("openFile() error while creating tempfile")
	}
	defer os.Remove(f.Name()) // delete after tests

	type args struct {
		file string
	}
	tests := []struct {
		name    string
		args    args
		want    *os.File
		wantErr bool
		err     error
	}{
		{
			name:    "default",
			args:    args{file: f.Name()},
			want:    f,
			wantErr: false,
			err:     nil,
		},
		{
			name:    "missing-file",
			args:    args{file: ""},
			want:    nil,
			wantErr: true,
			err:     os.ErrNotExist,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := openFile(tt.args.file)
			if (err != nil) != tt.wantErr {
				t.Errorf("openFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				gotFileInfo, err := got.Stat()
				if err != nil {
					t.Errorf("openFile() error while get stat info: %v\n", err)
				}
				wantFileInfo, err := tt.want.Stat()
				if err != nil {
					t.Errorf("openFile() error while get stat info: %v\n", err)
				}
				if !os.SameFile(gotFileInfo, wantFileInfo) {
					t.Errorf("openFile() = %v, want %v", got.Fd(), tt.want.Fd())
				}
			}
		})
	}
}

func Test_parseConfig(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    *Config
		wantErr bool
	}{
		{
			name:    "default",
			args:    args{bytes.NewBufferString(`{"Api": { "Port" : 8080 }}`)},
			want:    &Config{API: APIConfig{Port: 8080}},
			wantErr: false,
		},
		{
			name:    "invalid-json",
			args:    args{bytes.NewBufferString(``)},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseConfig(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_loadConfig(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name    string
		args    args
		want    *Config
		wantErr bool
	}{
		{
			name: "default",
			args: args{file: "config_test.json"},
			want: &Config{
				Httpd: HttpdConfig{
					InternalMode: true,
					RootDir:      "data/web/",
				},
				Importer: ImportConfig{
					ScanDir:     "data/import/",
					UseImporter: true,
				},
				API: APIConfig{
					Port: 3000,
				},
			},
			wantErr: false,
		},
		{
			name:    "missing file",
			args:    args{file: "doesnotexists.json"},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "invalid json",
			args:    args{file: "configfail_test.json"},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := loadConfig(tt.args.file)
			if (err != nil) != tt.wantErr {
				t.Errorf("loadConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("loadConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
