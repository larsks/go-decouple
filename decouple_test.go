package decouple

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite
}

func TestDecouple(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

func (t *TestSuite) TestGetStringExists() {
	expected := "This is a test"
	t.NoError(os.Setenv("TEST_VAR_EXISTS", expected))

	have, exists := GetString("TEST_VAR_EXISTS", "")
	t.True(exists)
	t.Equal(have, expected)
}

func (t *TestSuite) TestGetStringNotExists() {
	expected := "This is a test"

	have, exists := GetString("TEST_VAR_NOT_EXISTS", expected)
	t.False(exists)
	t.Equal(have, expected)
}

func (t *TestSuite) TestGetIntExists() {
	expected := 42
	t.NoError(os.Setenv("TEST_VAR_EXISTS", fmt.Sprintf("%d", 42)))
	have, exists := GetInt("TEST_VAR_EXISTS", 0)
	t.True(exists)
	t.Equal(have, expected)
}

func (t *TestSuite) TestGetIntNotExists() {
	expected := 42
	have, exists := GetInt("TEST_VAR_NOT_EXISTS", 42)
	t.False(exists)
	t.Equal(have, expected)
}

func (t *TestSuite) TestGetIntInRangeExists() {
	expected := 42
	t.NoError(os.Setenv("TEST_VAR_EXISTS", fmt.Sprintf("%d", 42)))
	have, exists := GetIntInRange("TEST_VAR_EXISTS", 25, 10, 50)
	t.True(exists)
	t.Equal(have, expected)
}

func (t *TestSuite) TestGetIntInRangeExistsMax() {
	expected := 50
	t.NoError(os.Setenv("TEST_VAR_EXISTS", fmt.Sprintf("%d", 100)))
	have, exists := GetIntInRange("TEST_VAR_EXISTS", 25, 10, 50)
	t.True(exists)
	t.Equal(have, expected)
}

func (t *TestSuite) TestGetIntInRangeExistsMin() {
	expected := 10
	t.NoError(os.Setenv("TEST_VAR_EXISTS", fmt.Sprintf("%d", 0)))
	have, exists := GetIntInRange("TEST_VAR_EXISTS", 25, 10, 50)
	t.True(exists)
	t.Equal(have, expected)
}

func (t *TestSuite) TestGetIntInRangeNotExists() {
	expected := 25
	have, exists := GetIntInRange("TEST_VAR_NOT_EXISTS", 25, 10, 50)
	t.False(exists)
	t.Equal(have, expected)
}

func (t *TestSuite) TestGetCSVStringExists() {
	expected := []string{"one", "two", "three"}

	t.NoError(os.Setenv("TEST_VAR_EXISTS", "one,two,three"))
	have, exists := GetCSVString("TEST_VAR_EXISTS", []string{})
	t.True(exists)
	t.Equal(have, expected)
}

func (t *TestSuite) TestGetCSVStringNotExists() {
	expected := []string{"one", "two", "three"}

	have, exists := GetCSVString("TEST_VAR_NOT_EXISTS", expected)
	t.False(exists)
	t.Equal(have, expected)
}
