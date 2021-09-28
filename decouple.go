// Package go-decouple is inspired by the python-decouple package
// (https://github.com/henriquebastos/python-decouple). It provides a
// layuer above gotdotenv (https://github.com/joho/godotenv) that
// handles defaults and type conversion.
//
// For example, if you want to read an integer value from an
// environment variable, you can call:
//
//	decouple.Load()
//	value, exists = decouple.GetInt("MY_INT_VAR", 0)
//
// If MY_INT_VAR exists, 'value' will be set to the integer value and
// 'exists' will be true. If MY_INT_VAR does not exist, 'value' will
// be set to 0 and 'exists' will be false. If MY_INT_VAR exists but
// cannot be converted into the requested type, 'value' will be set to
// 0 and 'exists' will be false.
package decouple

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

var nameprefix string

// SetPrefix sets a prefix that will be applied when looking for
// variables. If you call:
//
//	decouple.SetPrefix("FOO_")
//	decouple.GetString("CONFIG")
//
// Then decouple will look for a variable named "FOO_CONFIG".
func SetPrefix(prefix string) {
	nameprefix = prefix
}

// LookupEnv is a proxy for os.LookupEnv that applies the prefix
// configured with SetPrefix.
func LookupEnv(name string) (string, bool) {
	name = fmt.Sprintf("%s%s", nameprefix, name)
	return os.LookupEnv(name)
}

// GetString returns the value of an environment variable as a string.
//
// If the named variable exists, return the tuple (value, true). If
// the named variable does not exist, reutrn the tuple (defval,
// false).
//
// Example:
//
//	os.Setenv("CONFIG_PATH", "/etc/sharedconfig.yaml")
//	configpath, _ := decouple.GetString("CONFIG_PATH", "/home/.config/myconfig.yaml")
func GetString(name, defval string) (string, bool) {
	val, exists := LookupEnv(name)
	if !exists {
		return defval, false
	}

	return val, true
}

// GetStringChoices returns the value of an environment as a string if
// it is a valid choice. Otherwise, returns a default value.
//
// If the named variable exists and is a valid choice, return the
// tuple (value, true). If the named varible exists but is not a valid
// choice, reutrn (defval, true). If the named variable does not
// exist, return (defval, false).
//
// Example:
//
//	os.SetEnv("WIDGET_SIZE", "small")
//	widget_size := GetStringChoices("WIDGET_SIZE", "small", []string{"small", "medium", "large"})
func GetStringChoices(name, defval string, choices []string) (string, bool) {
	val, exists := GetString(name, defval)

	for _, choice := range choices {
		if val == choice {
			return val, exists
		}
	}

	return defval, exists
}

// GetInt returns the value of an environment variable as an int.
//
// If the named variable exists, attempt to convert it to an integer.
// If the conversion is successful, return (value, true). If the
// conversion fails or if the named variable does not exist, return
// (defval, false).
//
// Example:
//
//	os.Setenv("WIDGET_COUNT", 2)
//	widgetCount, _ := decouple.GetInt("WIDGET_COUNT", 10)
func GetInt(name string, defval int) (int, bool) {
	val, exists := LookupEnv(name)
	if !exists {
		return defval, false
	}

	ret, err := strconv.ParseInt(val, 0, 0)
	if err != nil {
		return defval, false
	}

	return int(ret), true
}

// GetIntInRange returns the value of environment variable as an int,
// clamped to an explicit range.
//
// If the named variable exists, attempt to convert it to an integer.
// If the conversion is successful and the value falls within the
// given range, return (value, true). If value > maxval, return
// (maxval, true). If value < minval, return (minval, true). If the
// conversion fails or if the named variable does not exist, return
// (defval, false).
//
// Example:
//
//	os.Setenv("LOG_LEVEL", 2)
//	logLevel, _ := decouple.GetIntInRange("LOG_LEVEL", 1, -1, 5)
func GetIntInRange(name string, defval, minval, maxval int) (int, bool) {
	ret, exists := GetInt(name, defval)

	switch {
	case ret < minval:
		ret = minval
	case ret > maxval:
		ret = maxval
	}

	return int(ret), exists
}

// GetBool returns the value of an environment varilable as a
// boolean.
//
// If the named variable exists, attempt to convert it to a bool. If
// the conversion is successful, return (value, true). If the
// conversion fails or if the named variable does not exist, return
// (defval, false).
//
// Example:
//
//	os.Setenv("DEBUG_MODE", "true")
// 	debugMode, _ := decouple.GetBool("DEBUG_MODE")
func GetBool(name string, defval bool) (bool, bool) {
	val, exists := LookupEnv(name)
	if !exists {
		return defval, false
	}

	ret, err := strconv.ParseBool(val)
	if err != nil {
		return defval, false
	}

	return ret, true
}

// GetCSVString parses an environment variable as a single row in a
// CSV document and returns a list of strings.
//
// Example:
//
//	os.Setenv("LIST_OF_NAMES", "alice,bob,carol")
//	names, _ := decouple.GetCSVString("LIST_OF_NAMES", []string{})
func GetCSVString(name string, defval []string) ([]string, bool) {
	val, exists := GetString(name, "")
	if !exists {
		return defval, false
	}

	r := strings.NewReader(val)
	csvr := csv.NewReader(r)
	rec, err := csvr.Read()
	if err != nil {
		return defval, false
	}

	return rec, true
}

// Load is a proxy for godotenv.Load. It will load environment
// variables from the named files, or from '.env' if no filenames are
// provided.
//
// Load variables from '.env':
//
//	decouple.Load()
//
// Load variables from 'production.env':
//
//	decouple.Load("production.env")
func Load(filenames ...string) error {
	return godotenv.Load(filenames...)
}
