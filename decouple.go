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
	"fmt"
	"os"
	"strconv"

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
//	configpath, _ := decouple.GetString("CONFIG_PATH", "/home/.config/myconfig.yaml")
func GetString(name, defval string) (string, bool) {
	val, exists := LookupEnv(name)
	if !exists {
		return defval, false
	}

	return val, true
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
