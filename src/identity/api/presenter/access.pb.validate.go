// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: identity/api/presenter/access.proto

package presenter

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
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
)

// Validate checks the field values on FabricCredentialsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *FabricCredentialsRequest) Validate() error {
	if m == nil {
		return nil
	}

	if err := m._validateEmail(m.GetEmail()); err != nil {
		return FabricCredentialsRequestValidationError{
			field:  "Email",
			reason: "value must be a valid email address",
			cause:  err,
		}
	}

	if len(m.GetPasscode()) < 8 {
		return FabricCredentialsRequestValidationError{
			field:  "Passcode",
			reason: "value length must be at least 8 bytes",
		}
	}

	return nil
}

func (m *FabricCredentialsRequest) _validateHostname(host string) error {
	s := strings.ToLower(strings.TrimSuffix(host, "."))

	if len(host) > 253 {
		return errors.New("hostname cannot exceed 253 characters")
	}

	for _, part := range strings.Split(s, ".") {
		if l := len(part); l == 0 || l > 63 {
			return errors.New("hostname part must be non-empty and cannot exceed 63 characters")
		}

		if part[0] == '-' {
			return errors.New("hostname parts cannot begin with hyphens")
		}

		if part[len(part)-1] == '-' {
			return errors.New("hostname parts cannot end with hyphens")
		}

		for _, r := range part {
			if (r < 'a' || r > 'z') && (r < '0' || r > '9') && r != '-' {
				return fmt.Errorf("hostname parts can only contain alphanumeric characters or hyphens, got %q", string(r))
			}
		}
	}

	return nil
}

func (m *FabricCredentialsRequest) _validateEmail(addr string) error {
	a, err := mail.ParseAddress(addr)
	if err != nil {
		return err
	}
	addr = a.Address

	if len(addr) > 254 {
		return errors.New("email addresses cannot exceed 254 characters")
	}

	parts := strings.SplitN(addr, "@", 2)

	if len(parts[0]) > 64 {
		return errors.New("email address local phrase cannot exceed 64 characters")
	}

	return m._validateHostname(parts[1])
}

// FabricCredentialsRequestValidationError is the validation error returned by
// FabricCredentialsRequest.Validate if the designated constraints aren't met.
type FabricCredentialsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FabricCredentialsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FabricCredentialsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FabricCredentialsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FabricCredentialsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FabricCredentialsRequestValidationError) ErrorName() string {
	return "FabricCredentialsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e FabricCredentialsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFabricCredentialsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FabricCredentialsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FabricCredentialsRequestValidationError{}

// Validate checks the field values on FabricCredentialsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *FabricCredentialsResponse) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetSecret()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return FabricCredentialsResponseValidationError{
				field:  "Secret",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for ApiAccessToken

	if v, ok := interface{}(m.GetUser()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return FabricCredentialsResponseValidationError{
				field:  "User",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// FabricCredentialsResponseValidationError is the validation error returned by
// FabricCredentialsResponse.Validate if the designated constraints aren't met.
type FabricCredentialsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FabricCredentialsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FabricCredentialsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FabricCredentialsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FabricCredentialsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FabricCredentialsResponseValidationError) ErrorName() string {
	return "FabricCredentialsResponseValidationError"
}

// Error satisfies the builtin error interface
func (e FabricCredentialsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFabricCredentialsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FabricCredentialsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FabricCredentialsResponseValidationError{}

// Validate checks the field values on VaultSecret with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *VaultSecret) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Token

	// no validation rules for Path

	return nil
}

// VaultSecretValidationError is the validation error returned by
// VaultSecret.Validate if the designated constraints aren't met.
type VaultSecretValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e VaultSecretValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e VaultSecretValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e VaultSecretValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e VaultSecretValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e VaultSecretValidationError) ErrorName() string { return "VaultSecretValidationError" }

// Error satisfies the builtin error interface
func (e VaultSecretValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sVaultSecret.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = VaultSecretValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = VaultSecretValidationError{}

// Validate checks the field values on CertificateAuthRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CertificateAuthRequest) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetCertificate()) < 100 {
		return CertificateAuthRequestValidationError{
			field:  "Certificate",
			reason: "value length must be at least 100 bytes",
		}
	}

	if len(m.GetSigningKey()) < 25 {
		return CertificateAuthRequestValidationError{
			field:  "SigningKey",
			reason: "value length must be at least 25 bytes",
		}
	}

	return nil
}

// CertificateAuthRequestValidationError is the validation error returned by
// CertificateAuthRequest.Validate if the designated constraints aren't met.
type CertificateAuthRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CertificateAuthRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CertificateAuthRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CertificateAuthRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CertificateAuthRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CertificateAuthRequestValidationError) ErrorName() string {
	return "CertificateAuthRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CertificateAuthRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCertificateAuthRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CertificateAuthRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CertificateAuthRequestValidationError{}

// Validate checks the field values on CertificateAuthResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CertificateAuthResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for ApiAccessToken

	if v, ok := interface{}(m.GetUser()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CertificateAuthResponseValidationError{
				field:  "User",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// CertificateAuthResponseValidationError is the validation error returned by
// CertificateAuthResponse.Validate if the designated constraints aren't met.
type CertificateAuthResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CertificateAuthResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CertificateAuthResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CertificateAuthResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CertificateAuthResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CertificateAuthResponseValidationError) ErrorName() string {
	return "CertificateAuthResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CertificateAuthResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCertificateAuthResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CertificateAuthResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CertificateAuthResponseValidationError{}
