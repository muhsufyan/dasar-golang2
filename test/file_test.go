package test

import (
	"testing"

	"github.com/muhsufyan/dasar-golang2/simple"
	"github.com/stretchr/testify/assert"
)

func TestConnection(t *testing.T) {
	connection, cleanup := simple.InitializedConnection("Databaes")
	assert.NotNil(t, connection)

	cleanup()
}
