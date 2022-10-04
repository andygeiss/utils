package assets_test

import (
	"github.com/andygeiss/utils/assert"
	"github.com/andygeiss/utils/assets"
	"os"
	"path/filepath"
	"testing"
)

func TestReadGobFromPath(t *testing.T) {
	path := filepath.Join("testdata", "foo.gob")
	_ = os.Remove("testdata")
	type Data struct{ Name string }
	_ = assets.WriteGobToPath(Data{Name: "foo"}, path)
	got, err := assets.ReadGobFromPath[Data](path)
	assert.That("err should be nil", t, err, nil)
	assert.That("Name should be foo", t, got.Name, "foo")
}
