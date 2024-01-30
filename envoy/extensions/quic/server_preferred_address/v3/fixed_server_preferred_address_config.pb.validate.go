//go:build !disable_pgv
// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/extensions/quic/server_preferred_address/v3/fixed_server_preferred_address_config.proto

package server_preferred_addressv3

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

// Validate checks the field values on FixedServerPreferredAddressConfig with
// the rules defined in the proto definition for this message. If any rules
// are violated, the first error encountered is returned, or nil if there are
// no violations.
func (m *FixedServerPreferredAddressConfig) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on FixedServerPreferredAddressConfig
// with the rules defined in the proto definition for this message. If any
// rules are violated, the result is a list of violation errors wrapped in
// FixedServerPreferredAddressConfigMultiError, or nil if none found.
func (m *FixedServerPreferredAddressConfig) ValidateAll() error {
	return m.validate(true)
}

func (m *FixedServerPreferredAddressConfig) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	switch v := m.Ipv4Type.(type) {
	case *FixedServerPreferredAddressConfig_Ipv4Address:
		if v == nil {
			err := FixedServerPreferredAddressConfigValidationError{
				field:  "Ipv4Type",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		// no validation rules for Ipv4Address
	default:
		_ = v // ensures v is used
	}
	switch v := m.Ipv6Type.(type) {
	case *FixedServerPreferredAddressConfig_Ipv6Address:
		if v == nil {
			err := FixedServerPreferredAddressConfigValidationError{
				field:  "Ipv6Type",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		// no validation rules for Ipv6Address
	default:
		_ = v // ensures v is used
	}

	if len(errors) > 0 {
		return FixedServerPreferredAddressConfigMultiError(errors)
	}

	return nil
}

// FixedServerPreferredAddressConfigMultiError is an error wrapping multiple
// validation errors returned by
// FixedServerPreferredAddressConfig.ValidateAll() if the designated
// constraints aren't met.
type FixedServerPreferredAddressConfigMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m FixedServerPreferredAddressConfigMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m FixedServerPreferredAddressConfigMultiError) AllErrors() []error { return m }

// FixedServerPreferredAddressConfigValidationError is the validation error
// returned by FixedServerPreferredAddressConfig.Validate if the designated
// constraints aren't met.
type FixedServerPreferredAddressConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FixedServerPreferredAddressConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FixedServerPreferredAddressConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FixedServerPreferredAddressConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FixedServerPreferredAddressConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FixedServerPreferredAddressConfigValidationError) ErrorName() string {
	return "FixedServerPreferredAddressConfigValidationError"
}

// Error satisfies the builtin error interface
func (e FixedServerPreferredAddressConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFixedServerPreferredAddressConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FixedServerPreferredAddressConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FixedServerPreferredAddressConfigValidationError{}
