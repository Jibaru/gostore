package shared

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateHttpUrl(t *testing.T) {
	urlGenerator := NewUrlGenerator(
		"localhost",
		80,
		"storage",
		false,
	)

	expectedUrl := "http://localhost:80/storage/folder-name/object.png"
	actualUrl := urlGenerator.GenerateUrlFromObjectPath("/folder-name/object.png")

	assert.Equal(t, expectedUrl, actualUrl)
}

func TestGenerateHttpsUrl(t *testing.T) {
	urlGenerator := NewUrlGenerator(
		"localhost",
		80,
		"storage",
		true,
	)

	expectedUrl := "https://localhost:80/storage/folder-name/object.png"
	actualUrl := urlGenerator.GenerateUrlFromObjectPath("/folder-name/object.png")

	assert.Equal(t, expectedUrl, actualUrl)
}
