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
      repository: myapp
      tag: "1.0"
  app2:
    image:
      registry: docker.io
      repository: app2
      tag: "1.1"
  `)
	err := ii.UnmarshalYAML(yamlData)
	assert.NoError(t, err)
	assert.Equal(t, "docker.io/myapp:1.0", ii.Images["app.image"])
	assert.Equal(t, "docker.io/app2:1.1", ii.Images["app2.image"])
}

func TestImageInfoMarshalJSON(t *testing.T) {
	ii := &ImageInfo{
		Images: map[string]string{
			"images": "example.com/myapp:latest",
		},
	}
	jsonData, err := ii.MarshalJSON()
	assert.NoError(t, err)
	assert.JSONEq(t, `{"images":"example.com/myapp:latest"}`, string(jsonData))
}
