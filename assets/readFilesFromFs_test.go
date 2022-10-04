package assets_test

import (
	"embed"
	"github.com/andygeiss/utils/assert"
	"github.com/andygeiss/utils/assets"
	"testing"
)

//go:embed testembed
var testembed embed.FS

func TestReadFilesFromFs(t *testing.T) {
	filesMap, err := assets.ReadFilesFromFs(testembed)
	assert.That("err should be nil", t, err, nil)
	assert.That("filesMap should have file bar.txt", t, filesMap["testembed/bar.txt"] != nil, true)
	assert.That("file bar.txt content should be correct", t, filesMap["testembed/bar.txt"], []byte("bar"))
	assert.That("filesMap should have file foo.txt", t, filesMap["testembed/foo.txt"] != nil, true)
	assert.That("file foo.txt content should be correct", t, filesMap["testembed/foo.txt"], []byte("foo"))
}
