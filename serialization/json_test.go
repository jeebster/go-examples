package serialization

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestStruct struct {
	id          int    `json:"id"`
	description string `json:"description"`
}

func TestToJSON(t *testing.T) {
	testCases := []TestStruct{
		{1, "hello"},
		{2, "world"},
	}

	for _, testCase := range testCases {
		b := make([]byte, 0, 512)
		buf := bytes.NewBuffer(b)
		err := ToJSON(testCase, buf)
		assert.Nil(t, err)
	}
}

func TestFromJSON(t *testing.T) {
	testCases := []string{
		"{\"id\":1,\"description\":\"hello\"}",
		"{\"id\":2,\"description\":\"hello\"}",
	}

	for _, testCase := range testCases {
		reader := strings.NewReader(testCase)
		iface := &TestStruct{}
		err := FromJSON(iface, reader)
		assert.Nil(t, err)
	}
}
