package filter

import (
	"reflect"
	"testing"
)

func TestAlgorithmsFilter_Apply(t *testing.T) {
	type args struct {
		filenames []string
	}
	tests := []struct {
		name         string
		args         args
		wantFiltered []string
	}{
		// TODO: Add test cases.
		{"multiple", args{[]string{"0430_name_x_y.go", "1000_name_z_y.go"}}, []string{"0430_name_x_y.go", "1000_name_z_y.go"}},
		{"none", args{[]string{"lcci_name_x_y.go", "_name_z_y.go"}}, []string(nil)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &AlgorithmsFilter{}
			gotFiltered := f.Apply(tt.args.filenames)
			if !reflect.DeepEqual(gotFiltered, tt.wantFiltered) {
				t.Errorf("AlgorithmsFilter.Apply() = %v, want %v", gotFiltered, tt.wantFiltered)
			}
		})
	}
}
