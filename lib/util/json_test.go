package util

import "testing"

func TestInterfaceToJSONString(t *testing.T) {
	type test struct {
		A int `json:"a"`
	}
	type testfail struct {
		A int
	}
	type args struct {
		data interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "default",
			args:    args{test{A: 1}},
			want:    `{"a":1}`,
			wantErr: false,
		},
		{
			name:    "default",
			args:    args{make(chan int)},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := InterfaceToJSONString(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("InterfaceToJSONString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("InterfaceToJSONString() = %v, want %v", got, tt.want)
			}
		})
	}
}
