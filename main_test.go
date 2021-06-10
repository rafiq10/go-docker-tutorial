package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEcho(t *testing.T) {
	// test funny path

	_, err := echo([]string{"bin-name", "hello", "world!"})

	require.NoError(t, err)
}

func TestErrorNoArgs(t *testing.T) {
	// test empt arguments
	_, err := echo([]string{})
	require.Error(t, err)
}
