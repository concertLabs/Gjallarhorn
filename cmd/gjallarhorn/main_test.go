package main

import (
	"reflect"
	"testing"
)

func Test_defaultArgs(t *testing.T) {
	tests := []struct {
		name string
		want Args
	}{
		{"default", Args{config: "config.json"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := defaultArgs(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("defaultArgs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseArgs(t *testing.T) {
	tests := []struct {
		name string
		want Args
	}{
		{"default", Args{config: "config.json"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseArgs(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseArgs() = %v, want %v", got, tt.want)
			}
		})
	}
}
