package util

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// GetQuery assigns query values from the context to the provided pointers,
// based on the provided query parameter names and pointers.
func GetQuery(c *gin.Context, nameAndArgs ...any) error {
	if len(nameAndArgs)%2 != 0 {
		return errors.New("must pass pairs of query name and pointer")
	}

	for i := 0; i < len(nameAndArgs); i += 2 {
		fieldName, ok := nameAndArgs[i].(string)
		if !ok {
			return errors.New("query parameter name must be a string")
		}

		arg := nameAndArgs[i+1]
		val := reflect.ValueOf(arg)
		if val.Kind() != reflect.Ptr || val.IsNil() {
			return fmt.Errorf("must pass a pointer for %s", fieldName)
		}

		valElem := val.Elem()
		// Fetch the query value based on the parameter name
		queryValue, isOk := c.GetQuery(fieldName)
		if !isOk {
			return fmt.Errorf("%s(%s) is required", fieldName, valElem.Kind())
		}

		// Set the appropriate type based on the kind of the argument

		switch valElem.Kind() {
		case reflect.Bool:
			parsedBool, err := strconv.ParseBool(queryValue)
			if err != nil {
				return fmt.Errorf("%s must be a boolean", fieldName)
			}
			valElem.SetBool(parsedBool)
		case reflect.String:
			valElem.SetString(queryValue)
		case reflect.Int:
			parsedInt, err := strconv.Atoi(queryValue)
			if err != nil {
				return fmt.Errorf("%s must be an integer", fieldName)
			}
			valElem.SetInt(int64(parsedInt))
		default:
			return fmt.Errorf("%s has unsupported type", fieldName)
		}
	}
	return nil
}

type ParamConstraint interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 |
		float32 | float64 |
		bool |
		string | uuid.UUID
}

func ParseParam[T ParamConstraint](key, paramValue string, validator func(T) error) (T, error) {
	var output T
	var err error
	switch any(output).(type) {
	case bool:
		var parsed bool
		parsed, err = strconv.ParseBool(paramValue)
		output = any(parsed).(T)

	case string:
		output = any(paramValue).(T)

	case int:
		var parsed int64
		parsed, err = strconv.ParseInt(paramValue, 10, 0)
		output = any(parsed).(T)

	case int8:
		var parsed int64
		parsed, err = strconv.ParseInt(paramValue, 10, 8)
		output = any(parsed).(T)

	case int16:
		var parsed int64
		parsed, err = strconv.ParseInt(paramValue, 10, 16)
		output = any(parsed).(T)

	case int32:
		var parsed int64
		parsed, err = strconv.ParseInt(paramValue, 10, 32)
		output = any(parsed).(T)

	case int64:
		var parsed int64
		parsed, err = strconv.ParseInt(paramValue, 10, 64)
		output = any(parsed).(T)

	case uint:
		parsed, err := strconv.ParseUint(paramValue, 10, 0)
		if err != nil {
			return output, err
		}
		output = any(uint(parsed)).(T)

	case uint8:
		var parsed uint64
		parsed, err = strconv.ParseUint(paramValue, 10, 8)
		output = any(parsed).(T)

	case uint16:
		var parsed uint64
		parsed, err = strconv.ParseUint(paramValue, 10, 16)
		output = any(parsed).(T)

	case uint32:
		var parsed uint64
		parsed, err = strconv.ParseUint(paramValue, 10, 32)
		output = any(parsed).(T)

	case uint64:
		var parsed uint64
		parsed, err = strconv.ParseUint(paramValue, 10, 64)
		output = any(parsed).(T)
	case float32:
		var parsed float64
		parsed, err = strconv.ParseFloat(paramValue, 32)
		output = any(parsed).(T)
	case float64:
		var parsed float64
		parsed, err = strconv.ParseFloat(paramValue, 64)
		output = any(parsed).(T)
	case uuid.UUID:
		var parsed uuid.UUID
		parsed, err = uuid.Parse(paramValue)
		output = any(parsed).(T)

	}
	if err != nil {
		return output, fmt.Errorf("path parameter %s is invalid: %w", key, err)
	}
	// Validate the parsed value
	if validator != nil {
		if err := validator(output); err != nil {
			return output, err
		}
	}
	return output, nil
}

func GetPathParam[T ParamConstraint](c *gin.Context, key string, validator func(T) error) (T, error) {

	paramValue, ok := c.Params.Get(key)
	if !ok {
		var output T
		return output, fmt.Errorf("path parameter %s is required", key)
	}

	return ParseParam(key, paramValue, validator)

}
func GetQueryParam[T ParamConstraint](c *gin.Context, key string, validator func(T) error) (T, error) {

	paramValue, ok := c.GetQuery(key)
	if !ok {
		var output T
		return output, fmt.Errorf("Query parameter %s is required", key)
	}

	return ParseParam(key, paramValue, validator)

}

// GetSHA256Hash takes an array of strings, joins them, and returns the SHA-256 hash
func GetSHA256Hash(words []string) []byte {
	// Join the strings (no separator, or use a custom one if you prefer)
	combined := strings.Join(words, "")

	// Calculate SHA-256
	hash := sha256.Sum256([]byte(combined))

	// Return as byte slice
	return hash[:]
}

func getPublicIP() (string, error) {
	resp, err := http.Get("https://api.ipify.org?format=text")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	ip, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(ip), nil
}
func IsValidIP(ip string) bool {
	return net.ParseIP(ip) != nil
}

func GetObjectFromContext[T any](c *gin.Context, key string) (T, bool) {
	value, exists := c.Get(key)
	if !exists {
		var zero T
		return zero, false
	}
	obj, ok := value.(T)
	return obj, ok
}

func Ptr[T any](in T) *T {
	var o T = in
	return &o
}

func GetAuthorizationToken(w http.ResponseWriter, r *http.Request) (string, error) {
	w.Header().Add("Vary", "Authorization")
	token := ""
	if r.URL.Query().Get("jwt") != "" {
		token = r.URL.Query().Get("jwt")

	} else if cookie, err := r.Cookie("SAT"); err == nil {
		// get from cookies
		token = cookie.Value
	} else {
		// get auth header
		authHeader := r.Header.Get("Authorization")

		// sanity check
		if authHeader == "" {
			return "", errors.New("no auth header")
		}

		// split the header on spaces
		headerParts := strings.Split(authHeader, " ")
		if len(headerParts) != 2 {
			fmt.Println(headerParts)
			return "", errors.New("invalid auth header.")
		}

		// check to see if we have the word Bearer
		if headerParts[0] != "Bearer" {
			return "", errors.New("invalid auth header")
		}

		return headerParts[1], nil
	}
	return token, nil
}

func GenerateRandomString(n int) string {
	// Define the character set to use for the random string
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"

	// Seed the random number generator to ensure different results each time
	rand.Seed(time.Now().UnixNano())

	// Create a byte slice of the specified length
	b := make([]byte, n)

	// Fill the byte slice with random characters from the charset
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}

	// Convert the byte slice to a string and return it
	return string(b)
}

// --- PatchStruct functions (copy your existing code here) ---
func PatchStruct(dst any, src any) error {
	dstVal := reflect.ValueOf(dst)
	if dstVal.Kind() != reflect.Ptr || dstVal.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("dst must be pointer to struct")
	}

	srcVal := reflect.ValueOf(src)
	if srcVal.Kind() != reflect.Ptr || srcVal.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("src must be pointer to struct")
	}

	return patch(dstVal.Elem(), srcVal.Elem())
}

func patch(dst reflect.Value, src reflect.Value) error {
	for i := 0; i < src.NumField(); i++ {
		srcField := src.Field(i)
		srcTypeField := src.Type().Field(i)

		// Only handle pointer fields in input
		if srcField.Kind() != reflect.Ptr || srcField.IsNil() {
			continue
		}

		dstField := dst.FieldByName(srcTypeField.Name)
		if !dstField.IsValid() || !dstField.CanSet() {
			continue
		}

		srcElem := srcField.Elem()

		switch srcElem.Type() {
		case reflect.TypeOf(time.Time{}):
			if dstField.Kind() == reflect.Ptr {
				dstField.Set(reflect.New(reflect.TypeOf(time.Time{})))
				dstField.Elem().Set(srcElem)
			} else {
				dstField.Set(srcElem)
			}
		default:
			switch srcElem.Kind() {
			case reflect.Struct:
				if dstField.Kind() == reflect.Ptr {
					if dstField.IsNil() {
						dstField.Set(reflect.New(dstField.Type().Elem()))
					}
					err := patch(dstField.Elem(), srcElem)
					if err != nil {
						return err
					}
				} else if dstField.Kind() == reflect.Struct {
					err := patch(dstField, srcElem)
					if err != nil {
						return err
					}
				}
			default:
				if dstField.Kind() == srcElem.Kind() {
					dstField.Set(srcElem)
				}
			}
		}

	}

	return nil
}
