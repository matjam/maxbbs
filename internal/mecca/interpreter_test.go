package mecca_test

import (
	"bytes"
	"testing"

	"github.com/matjam/maxbbs/internal/mecca"
	"github.com/matjam/maxbbs/internal/system"
)

func TestInterpreter(t *testing.T) {
	bbs := system.NewBBS()

	template := `This is just a string you know. System Name: [sys_name] [sysop_name]`

	var in bytes.Buffer
	var out bytes.Buffer

	interpreter := mecca.NewInterpreter(bbs)
	session := interpreter.NewSession(&in, &out)
	err := session.Compile("test_template", template)
	if err != nil {
		t.Fatalf(err.Error())
	}

	err = session.Exec("test_template")
	if err != nil {
		t.Fatalf(err.Error())
	}

	t.Fatalf(out.String())
}
