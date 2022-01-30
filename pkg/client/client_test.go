package client

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	token = os.Getenv("PEXELS_API_TOKEN")
)

func TestSearchPhotos(t *testing.T) {
	clinet := New(token)
	res, err := clinet.SearchPhotos(&SearchPhotosRequest{
		Query:       "Ocean",
		Orientation: "landscape",
		Size:        "large",
		Color:       "red",
		Local:       "en-US",
		Page:        1,
		PerPage:     1,
	})
	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Greater(t, len(res.Photos), 0)
}
