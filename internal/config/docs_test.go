package config

import (
	_ "embed"
	"encoding"
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/pelletier/go-toml"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/multierr"

	"github.com/smartcontractkit/chainlink/core/services/chainlink"
)

func TestDoc(t *testing.T) {
	var c chainlink.Config
	d := toml.NewDecoder(strings.NewReader(docsTOML))
	// Note: using v1 of go-toml since v2 provides no feedback about which keys
	d.Strict(true) // Ensure no extra fields
	err := d.Decode(&c)
	if err != nil && strings.Contains(err.Error(), "undecoded keys: ") {
		t.Errorf("Docs contain extra fields: %v", err)
	} else {
		require.NoError(t, err)
	}

	err = assertFieldsNotNil(t, "", reflect.ValueOf(c))
	assert.NoError(t, err, multiErrorList(multierr.Errors(err)))

	//TODO validate defaults? ensure non-zero examples?
}

func assertFieldsNotNil(t *testing.T, prefix string, s reflect.Value) (err error) {
	t.Helper()
	require.Equal(t, reflect.Struct, s.Kind())

	typ := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		key := prefix
		if tf := typ.Field(i); !tf.Anonymous {
			if key != "" {
				key += "."
			}
			key += tf.Name
		}
		err = multierr.Combine(err, assertValNotNil(t, key, f))
	}
	return
}

func assertValuesNotNil(t *testing.T, prefix string, m reflect.Value) (err error) {
	t.Helper()
	require.Equal(t, reflect.Map, m.Kind())
	if prefix != "" {
		prefix += "."
	}

	mi := m.MapRange()
	for mi.Next() {
		key := prefix + mi.Key().String()
		err = multierr.Combine(err, assertValNotNil(t, key, mi.Value()))
	}
	return
}

func assertElementsNotNil(t *testing.T, prefix string, s reflect.Value) (err error) {
	t.Helper()
	require.Equal(t, reflect.Slice, s.Kind())

	for i := 0; i < s.Len(); i++ {
		err = multierr.Combine(err, assertValNotNil(t, prefix, s.Index(i)))
	}
	return
}

var (
	textUnmarshaler     encoding.TextUnmarshaler
	textUnmarshalerType = reflect.TypeOf(&textUnmarshaler).Elem()
)

func assertValNotNil(t *testing.T, key string, val reflect.Value) error {
	t.Helper()
	k := val.Kind()
	switch k {
	case reflect.Ptr, reflect.Map, reflect.Slice:
		if val.IsNil() {
			return fmt.Errorf("%s: missing from documentation", key)
		}
	}
	if k == reflect.Ptr {
		if val.Type().Implements(textUnmarshalerType) {
			return nil // skip values unmarshaled from strings
		}
		val = val.Elem()
	}
	switch val.Kind() {
	case reflect.Struct:
		if val.Type().Implements(textUnmarshalerType) {
			return nil // skip values unmarshaled from strings
		}
		return assertFieldsNotNil(t, key, val)
	case reflect.Map:
		return assertValuesNotNil(t, key, val)
	case reflect.Slice:
		return assertElementsNotNil(t, key, val)
	default:
		return nil
	}
}
