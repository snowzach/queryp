package queryp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueryParser(t *testing.T) {

	q, err := ParseQuery(`field=value&(another<=value|yet=another1|limit=weee)|bob=(a,"\"b",null,"null")|third=value&limit=10&option=beans&option[soup]=tomato&sort=test,-another`)
	assert.Nil(t, err)
	assert.Equal(t, "tomato", q.Options.Get("soup"))
	assert.Equal(t, `field=value&(another<=value|yet=another1|limit=weee)|bob=(a,"\"b",null,"null")|third=value&limit=10&option=beans&option[soup]=tomato&sort=test,-another`, q.String())

	q, err = ParseQuery(`class="htechpipe"&sort=id&offset=0&limit=10`)
	assert.Nil(t, err)
	assert.Equal(t, `class=htechpipe&limit=10&sort=id`, q.String())

	q, err = ParseQuery(`id=(1)`)
	assert.Nil(t, err)
	assert.Equal(t, `id=(1)`, q.String())

}
