package main_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	errors "github.com/henrmota/errors-handling-example"
)
func TestContext(t *testing.T) {

	err := errors.BadRequest.New("an_error")
	errWithContext := errors.AddErrorContext(err, "a_field", "the field is empty")

	expectedContext := map[string]string{"field": "a_field", "message": "the field is empty"}

	assert.Equal(t, errors.BadRequest, errors.GetType(errWithContext))
	assert.Equal(t, expectedContext, errors.GetErrorContext(errWithContext))
	assert.Equal(t, err.Error(), errWithContext.Error())
}

func TestContextInNoTypeError(t *testing.T) {
	err := errors.New("a custom error")

	errWithContext := errors.AddErrorContext(err, "a_field", "the field is empty")

	expectedContext := map[string]string{"field": "a_field", "message": "the field is empty"}

	assert.Equal(t, errors.NoType, errors.GetType(errWithContext))
	assert.Equal(t, expectedContext, errors.GetErrorContext(errWithContext))
	assert.Equal(t, err.Error(), errWithContext.Error())
}

func TestWrapf(t *testing.T) {
	err := errors.New("an_error")
	wrappedError := errors.BadRequest.Wrapf(err, "error %s", "1")

	assert.Equal(t, errors.BadRequest, errors.GetType(wrappedError))
	assert.EqualError(t, wrappedError, "error 1: an_error")
}

func TestWrapfInNoTypeError(t *testing.T) {
	err := errors.Newf("an_error %s", "2")
	wrappedError := errors.Wrapf(err, "error %s", "1")

	assert.Equal(t, errors.NoType, errors.GetType(wrappedError))
	assert.EqualError(t, wrappedError, "error 1: an_error 2")
}
