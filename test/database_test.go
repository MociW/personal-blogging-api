package test

import (
	"personal-blogging-api/app"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDatabaseConnection(t *testing.T) {
	db := app.NewDB()
	assert.NotNil(t, db)
}
