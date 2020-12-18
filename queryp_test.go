package queryp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueryParser(t *testing.T) {

	q, err := ParseQuery("field=value&(another<=value|yet=another1|limit=weee)|third=value&limit=10&option=beans&sort=test,-another")
	assert.Nil(t, err)
	assert.Equal(t, `field=value&(another<=value|yet=another1|limit=weee)|third=value&limit=10&option=beans&sort=test,-another`, q.String())

}
