package country

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCountryExists(t *testing.T) {
	require.True(t, Exists("SK"))
}

func TestCountryNotPresent(t *testing.T) {
	require.False(t, Exists("YY"))
}

func TestValidCountry(t *testing.T) {
	c, ok := Get("GB")
	require.True(t, ok)
	require.Equal(t, "GB", c.Alpha2Code)
	require.Equal(t, "GBR", c.Alpha3Code)
	require.Equal(t, "United Kingdom", c.Name)
	require.Equal(t, c.Name, c.String())
}

func TestInvalidCountry(t *testing.T) {
	c, ok := Get("XX")
	require.False(t, ok)
	require.Equal(t, "", c.Alpha2Code)
	require.Equal(t, "", c.Alpha3Code)
	require.Equal(t, "", c.Name)
	require.Equal(t, c.Name, c.String())
}

func TestValidBbanStructure(t *testing.T) {
	struc, ok := GetBbanStructure("FR")
	require.True(t, ok)
	require.Equal(t, 23, struc.Length())
}

func TestInvalidBbanStructure(t *testing.T) {
	struc, ok := GetBbanStructure("XXX")
	require.False(t, ok)
	require.Equal(t, 0, struc.Length())
}
