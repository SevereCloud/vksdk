package vksdk_test

import (
	"errors"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"testing"

	"github.com/SevereCloud/vksdk"
)

type version struct {
	major, minor, patch int
}

func newVersion(name string) (*version, error) {
	re := regexp.MustCompile(`(?m)(\d+)\.(\d+)\.(\d+)`)

	v := &version{}

	s := re.FindStringSubmatch(name)
	if len(s) == 4 {
		v.major, _ = strconv.Atoi(s[1])
		v.minor, _ = strconv.Atoi(s[2])
		v.patch, _ = strconv.Atoi(s[3])
	} else {
		return nil, errors.New("not found version")
	}

	return v, nil
}

func TestVersion(t *testing.T) {
	shaOut, err := exec.Command("git", "rev-list", "--tags", "--max-count=1").Output()
	if err != nil {
		t.Error(string(shaOut))
		t.Fatal("git rev-list: ", err)
	}

	sha := strings.TrimSuffix(string(shaOut), "\n")

	out, err := exec.Command("git", "describe", "--tags", sha, "--always").Output()
	if err != nil {
		t.Error(string(out))
		t.Fatal("git describe: ", err)
	}

	tag, err := newVersion(string(out))
	if err != nil {
		t.Fatal(err)
	}

	code, err := newVersion(vksdk.Version)
	if err != nil {
		t.Fatal(err)
	}

	switch {
	case code.major > tag.major: // 2.*.* > 1.*.*
		break
	case code.major < tag.major: // 1.*.* < 2.*.* bad
		t.Error("vksdk.Version bad MAJOR version")
	case code.minor > tag.minor: // 1.1.* > 1.0.*
		break
	case code.minor < tag.minor: // 1.0.* < 1.1.* bad
		t.Error("vksdk.Version bad MINOR version")
	case code.patch < tag.patch: // 1.0.0 < 1.0.1 bad
		t.Error("vksdk.Version bad PATCH version")
	}
}
