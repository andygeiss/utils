package assets_test

import (
	"github.com/andygeiss/utils/assert"
	"github.com/andygeiss/utils/assets"
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
	err := assets.WriteGobToPath(gob, path)
	assert.That("err should be nil", t, err, nil)
}
