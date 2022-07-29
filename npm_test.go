package ossinspector

import (
	"testing"

	"github.com/jarcoal/httpmock"
)

func Test_getRepoLink(t *testing.T) {
	type args struct {
		dep string
	}
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://registry.npmjs.com/react", httpmock.NewStringResponder(200, `{
		"repository": {
		  "type": "git",
		  "url": "git+https://github.com/facebook/react.git",
		  "directory": "packages/react"
		}
	  }`))

	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{name: "repo test", args: args{dep: "react"}, want: "https://github.com/facebook/react.git", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getRepoLink(tt.args.dep)
			if (err != nil) != tt.wantErr {
				t.Errorf("getRepo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getRepo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getRepo(t *testing.T) {
	type args struct {
		link string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{name:"link test", args: args{link: "https://github.com/facebook/react.git"}, want: "facebook/react", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getRepo(tt.args.link)
			if (err != nil) != tt.wantErr {
				t.Errorf("getRepo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getRepo() = %v, want %v", got, tt.want)
			}
		})
	}
}
