package models

import (
	"testing"
)

func init() {
	BasePath = "/Users/wenjie/work/go/src/github.com/qiniu-share"
	test()
}

func TestCreate(t *testing.T) {
	s := MockSharing()
	t.Logf("%v\n", s)
	a := MockAttachment()
	t.Logf("%v\n", a)
	isSuccess = a.Create(*s)
}
