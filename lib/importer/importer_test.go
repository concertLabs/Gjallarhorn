package importer

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/quiteawful/Gjallarhorn/lib/config"
)

func TestNew(t *testing.T) {
	type args struct {
		cfg config.ImportConfig
	}
	tests := []struct {
		name string
		args args
		want Importer
	}{
		{
			name: "default",
			args: args{
				config.ImportConfig{
					UseImporter: true,
					ScanDir:     ".",
				},
			},
			want: Importer{
				InputDir:    ".",
				UseImporter: true,
				Wait:        10,
				Processor:   make(chan string),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := New(tt.args.cfg)
			if got.InputDir != tt.want.InputDir ||
				got.UseImporter != tt.want.UseImporter ||
				got.Wait != tt.want.Wait ||
				got.Processor == nil {
				t.Errorf("NewImporter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRun(t *testing.T) {
	cfg := config.ImportConfig{
		UseImporter: false,
	}

	importer := New(cfg)
	importer.Run()
	// should return instantly
}

func Test_readFiles(t *testing.T) {

	files, _ := ioutil.ReadDir("testdata/")

	type args struct {
		dir string
	}
	tests := []struct {
		name    string
		args    args
		want    []os.FileInfo
		wantErr bool
	}{
		{
			name:    "default",
			args:    args{dir: "testdata/"},
			want:    files,
			wantErr: false,
		},
		{
			name:    "missing-folder",
			args:    args{dir: "notexistingfolder/"},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := readFiles(tt.args.dir)
			if (err != nil) != tt.wantErr {
				t.Errorf("readFiles() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				if len(got) != len(tt.want) {
					t.Errorf("readFiles() = len mismatch: got %d files, want %d files", len(got), len(tt.want))
				}
			}
		})
	}
}

func Test_run(t *testing.T) {
	testchan := make(chan string)
	type args struct {
		inputDir string
		ch       chan string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "default",
			args: args{inputDir: "testdata/", ch: testchan},
		},
		{
			name: "missing chan",
			args: args{inputDir: "testdata/", ch: nil},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			run(tt.args.inputDir, tt.args.ch)
		})
	}
}
