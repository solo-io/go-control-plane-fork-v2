// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/config/trace/v2/datadog.proto

package tracev2

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on DatadogConfig with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *DatadogConfig) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DatadogConfig with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in DatadogConfigMultiError, or
// nil if none found.
func (m *DatadogConfig) ValidateAll() error {
	return m.validate(true)
}

func (m *DatadogConfig) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(m.GetCollectorCluster()) < 1 {
		err := DatadogConfigValidationError{
			field:  "CollectorCluster",
			reason: "value length must be at least 1 bytes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(m.GetServiceName()) < 1 {
		err := DatadogConfigValidationError{
			field:  "ServiceName",
			reason: "value length must be at least 1 bytes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return DatadogConfigMultiError(errors)
	}

	return nil
}

// DatadogConfigMultiError is an error wrapping multiple validation errors
// returned by DatadogConfig.ValidateAll() if the designated constraints
// aren't met.
type DatadogConfigMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DatadogConfigMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DatadogConfigMultiError) AllErrors() []error { return m }

// DatadogConfigValidationError is the validation error returned by
// DatadogConfig.Validate if the designated constraints aren't met.
type DatadogConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DatadogConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DatadogConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DatadogConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DatadogConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DatadogConfigValidationError) ErrorName() string { return "DatadogConfigValidationError" }

// Error satisfies the builtin error interface
func (e DatadogConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDatadogConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DatadogConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DatadogConfigValidationError{}
