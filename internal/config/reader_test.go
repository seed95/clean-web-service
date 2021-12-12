package config

import "testing"

func TestParse(t *testing.T) {
	type args struct {
		path string
		cfg  *Config
	}

	tests := []struct {
		name     string
		args     args
		wantErrs bool
	}{
		{
			name: "parse config.yaml",
			args: args{
				path: "testdata/config.yaml",
				cfg:  &Config{},
			},
			wantErrs: false,
		},
		{
			name: "parse config.yml",
			args: args{
				path: "testdata/config.yml",
				cfg:  &Config{},
			},
			wantErrs: false,
		},
		{
			name: "parse config.json",
			args: args{
				path: "testdata/config.json",
				cfg:  &Config{},
			},
			wantErrs: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Parse(tt.args.path, tt.args.cfg); (err != nil) != tt.wantErrs {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErrs)
			}
		})
	}
}
