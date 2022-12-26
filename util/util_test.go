package util

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRoundWithPrecision(t *testing.T) {
	require.Equal(t, 12.312, RoundWithPrecision(12.3124, 3))
	require.Equal(t, 0.32112, RoundWithPrecision(0.32112, 5))
	require.Equal(t, 0.3211, RoundWithPrecision(0.32112, 4))
}

func TestTickToPrecision(t *testing.T) {
	require.Equal(t, 0, TickToPrecision(0.))
	require.Equal(t, 1, TickToPrecision(0.1))
	require.Equal(t, 3, TickToPrecision(0.005))
	require.Equal(t, 5, TickToPrecision(0.00002))
}
