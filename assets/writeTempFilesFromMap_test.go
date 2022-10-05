package assets_test

import (
	"github.com/andygeiss/utils/assert"
	"github.com/andygeiss/utils/assets"
	"os"
	"path/filepath"
	"testing"
)

func TestWriteTempFilesFromMap(t *testing.T) {
	filesMap := map[string][]byte{
		"testdata/bar.txt": []byte("bar"),
		"testdata/foo.txt": []byte("foo"),
	}
	prefix, err := assets.WriteTempFilesFromMap(filesMap)
	_, err1 := os.Stat(filepath.Join(prefix, "testdata", "bar.txt"))
	_, err2 := os.Stat(filepath.Join(prefix, "testdata", "foo.txt"))
	assert.That("err should be nil", t, err, nil)
	assert.That("err1 should be nil", t, err1, nil)
	assert.That("err2 should be nil", t, err2, nil)
	assert.That("prefix should not be empty", t, prefix != "", true)
}
