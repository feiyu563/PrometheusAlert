package controllers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTimeDuration(t *testing.T) {
	source := "2023-08-04T02:51:54.972Z"
	duration := GetTimeDuration(source)
	if assert.NotEqual(t, "", duration) {
		t.Log(duration)
	}
}
