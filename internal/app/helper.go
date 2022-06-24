package app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"github.com/gorilla/schema"
	"github.com/nsf/jsondiff"
	"github.com/pkg/errors"
)

const (
	iso8601Format     = "2006-01-02T15:04:05.000Z"
	iso8601DATEFormat = "2006-01-02"
)

type ISO8601 struct {
	time time.Time
}

var decoder = schema.NewDecoder()

// ToEpoch converts an ISO8601 time to epoch time
func (t ISO8601) ToEpoch() Epoch {
	return Epoch(t.time.UnixNano())
}

//Epoch is a representation of NANO seconds past epoch (unix) time
type Epoch int64

// NewISO8601 creates a new ISO8601 time
func NewISO8601(v string) (ISO8601, error) {
	t, err := time.Parse(iso8601DATEFormat, string(v))
	if err != nil {
		t, err := time.Parse(iso8601Format, string(v))
		if err != nil {
			return ISO8601{}, fmt.Errorf("time - unable to parse value. Format should be either :%s or %s ", iso8601DATEFormat, iso8601Format)
		}
		return ISO8601{time: t}, nil
	}
	return ISO8601{time: t}, nil
}

// GetQueryParams maps the query params from an http request into an interface
func GetQueryParams(value interface{}, r *http.Request) error {
	// decoder lookup for values on the json tag, instead of the default schema tag
	decoder.SetAliasTag("json")

	var globalErr error

	// Decoder Register for custom type ISO8601
	decoder.RegisterConverter(ISO8601{}, func(input string) reflect.Value {
		ISOTime, errISO := NewISO8601(input)

		if errISO != nil {
			globalErr = errors.Wrapf(errISO, "handler - invalid iso time provided")
			return reflect.ValueOf(ISO8601{})
		}

		return reflect.ValueOf(ISOTime)
	})

	// Decoder Register for custom type Epoch
	decoder.RegisterConverter(Epoch(0), func(input string) reflect.Value {
		ISOTime, errISO := NewISO8601(input)

		if errISO != nil {
			globalErr = errors.Wrapf(errISO, "handler - invalid iso time provided")
			return reflect.ValueOf(ISO8601{}.ToEpoch())
		}

		return reflect.ValueOf(ISOTime.ToEpoch())
	})

	if err := decoder.Decode(value, r.URL.Query()); err != nil {
		return errors.Wrapf(err, "handler - failed to decode query params")
	}

	if globalErr != nil {
		return globalErr
	}

	return nil
}

func ServeJSON(res interface{}, w http.ResponseWriter, statusCode int) {
	w.WriteHeader(statusCode)
	bb, err := json.Marshal(res)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(bb)
}

// FileToStruct reads a json file to a struct
// As an additive utility a reader for the file bytes is returned to the caller
func FileToStruct(filepath string, s interface{}) io.Reader {
	bb, _ := ioutil.ReadFile(filepath)
	json.Unmarshal(bb, s)
	return bytes.NewReader(bb)
}

func NewTestServer(h http.HandlerFunc) (string, func()) {
	ts := httptest.NewServer(h)
	return ts.URL, func() { ts.Close() }
}

// AssertResBodyEqual asserts that the response body matches the results specified in a file path
func AssertResBodyEqual(t *testing.T, resFilePath string, r *http.Response) {
	body, _ := ioutil.ReadAll(r.Body)
	expected, _ := ioutil.ReadFile(resFilePath)
	opts := jsondiff.DefaultConsoleOptions()
	diff, diffStr := jsondiff.Compare(expected, body, &opts)
	if diff != jsondiff.FullMatch {
		t.Errorf(fmt.Sprintf("Diff=%s", diffStr))
	}
}
