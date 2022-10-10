package resources_test

import (
	"github.com/andygeiss/utils/assert"
	"github.com/andygeiss/utils/resources"
	"os"
	"path/filepath"
	"testing"
)

func TestWriteGobToPath(t *testing.T) {
	path := filepath.Join("testdata", "foo.gob")
	_ = os.Remove("testdata")
	gob := struct {
		Bar string
	}{
		Bar: "bar",
	}
	err := resources.WriteGobToPath(gob, path)
	assert.That("err should be nil", t, err, nil)
}
