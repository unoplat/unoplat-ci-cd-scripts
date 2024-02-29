package dockerimageprocessing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestImageInfoUnmarshalYAML(t *testing.T) {
	ii := NewImageInfo()
	yamlData := []byte(`
  resource:
    hello:
      enabled: "true"
  app:
    image:
      registry: docker.io
      repository: myapp4
      tag: "1.1"
  app2:
    image:
      registry: docker.io
      repository: app2
      tag: "1.1"
  `)
	err := ii.UnmarshalYAML(yamlData)
	assert.NoError(t, err)
	assert.Contains(t, ii.Images["docker.io/myapp4:1.1"], "app.image.tag")
	assert.Contains(t, ii.Images["docker.io/app2:1.1"], "app2.image.tag")
}

func TestImageInfoMarshalJSON(t *testing.T) {
	ii := NewImageInfo()
	ii.Images = map[string][]string{
		"example.com/myapp:latest": {"images"},
	}
	jsonData, err := ii.MarshalJSON()
	assert.NoError(t, err)
	assert.JSONEq(t, `{"example.com/myapp:latest":["images"]}`, string(jsonData))

}
