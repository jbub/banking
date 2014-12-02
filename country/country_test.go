package country

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type CountryTestSuite struct {
	suite.Suite
}

func (s *CountryTestSuite) TestCountryExists() {
	ok := Exists("GB")
	s.True(ok)
}

func (s *CountryTestSuite) TestCountryNotPresent() {
	ok := Exists("XX")
	s.False(ok)
}

func (s *CountryTestSuite) TestValidCountry() {
	c, ok := Get("GB")
	s.True(ok)
	s.Equal("GB", c.Alpha2Code)
	s.Equal("GBR", c.Alpha3Code)
	s.Equal("United Kingdom", c.Name)
	s.Equal(c.Name, c.String())
}

func (s *CountryTestSuite) TestInvalidCountry() {
	c, ok := Get("XX")
	s.False(ok)
	s.Equal("", c.Alpha2Code)
	s.Equal("", c.Alpha3Code)
	s.Equal("", c.Name)
	s.Equal(c.Name, c.String())
}

func (s *CountryTestSuite) TestValidBbanStructure() {
	structure, ok := GetBbanStructure("GB")
	s.True(ok)
	s.Equal(18, structure.Length())
}

func (s *CountryTestSuite) TestInvalidBbanStructure() {
	structure, ok := GetBbanStructure("XX")
	s.False(ok)
	s.Equal(0, structure.Length())
}

func TestCountryTestSuite(t *testing.T) {
	suite.Run(t, new(CountryTestSuite))
}
