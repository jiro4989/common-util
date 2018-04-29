package config

import (
	"fmt"
	"path/filepath"
	"testing"

	jos "github.com/jiro4989/common-util/os"
)

type testData struct {
	in   string
	got  string
	want string
}

func TestGetConfigDir(t *testing.T) {
	datas := []testData{
		testData{
			in:   "",
			got:  GetConfigDir(),
			want: fmt.Sprintf("%s/%s", jos.GetenvHome(), ConfigDir),
		},
	}

	for _, d := range datas {
		if d.got != d.want {
			t.Errorf("in:%v, got:%v, want:%v", d.in, d.got, d.want)
		}
	}
}

func TestGetAppConfigDir(t *testing.T) {
	datas := []testData{
		testData{
			in:   "a",
			got:  GetAppConfigDir("a"),
			want: filepath.Join(jos.GetenvHome(), ConfigDir, "a"),
		},
		testData{
			in:   "tmp",
			got:  GetAppConfigDir("tmp"),
			want: filepath.Join(jos.GetenvHome(), ConfigDir, "tmp"),
		},
	}

	for _, d := range datas {
		if d.got != d.want {
			t.Errorf("in:%v, got:%v, want:%v", d.in, d.got, d.want)
		}
	}
}

func TestMkAppConfigDir(t *testing.T) {
	datas := []string{"lib-go-tmp", "lib-go-tmp2"}
	for _, d := range datas {
		if err := MkAppConfigDir(d); err != nil {
			t.Errorf("%v", err)
		}
		cfgp := GetConfigDir() + "/" + d
		if !jos.Exists(cfgp) {
			t.Errorf("%vが存在しません。", cfgp)
		}
	}
}
