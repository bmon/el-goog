package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateSession(t *testing.T) {
	for i := 0; i < 10; i++ {
		sessionID := CreateSession(i)
		user_id, err := GetSessionUserID(sessionID)
		if assert.NoError(t, err, "no error should be returned") {
			assert.Equal(t, i, user_id, "user_ids should be the same")
		}
	}
}
