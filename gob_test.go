package goblin

import (
	"os"
	"testing"
)

func Test_GOBRead(t *testing.T) {
	var err error
	var testData struct {
		Hello []string
	} = struct{ Hello []string }{
		[]string{"World", "world"},
	}

	err = GOBWrite("?", testData)
	if err == nil {
		t.Error("fail")
	}

	err = GOBWrite("test.gob", testData)
	if err != nil {
		t.Error(err)
	}

	err = GOBRead("test.gob", &testData)
	if err != nil {
		t.Error(err)
	}

	os.Remove("test.gob")
}
