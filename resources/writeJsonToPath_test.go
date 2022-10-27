package resources_test

import (
	"github.com/andygeiss/utils/assert"
	"github.com/andygeiss/utils/resources"
	"os"
	"path/filepath"
	"testing"
)

func TestWriteJsonToPath(t *testing.T) {
	path := filepath.Join("testdata", "foo.json")
	_ = os.Remove("testdata")
	type Data struct {
		Name string `json:"name"`
	}
	err := resources.WriteJsonToPath(Data{Name: "bar"}, path)
	assert.That("err should be nil", t, err, nil)
}
