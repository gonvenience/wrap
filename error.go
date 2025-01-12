// Copyright © 2019 The Homeport Team
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package wrap

import (
	"errors"
	"fmt"
)

// ContextError interface describes the simple type that is able to provide a
// textual context as well as the cause explaining the underlying error.
//
// Deprecated: Discontinued, use fmt.Errorf() instead.
type ContextError interface {
	Context() string
	Cause() error
}

// ListOfErrors interface describes a list of errors with additional context
// information with an explanation.
//
// Deprecated: Discontinued, use errors.Join() instead.
type ListOfErrors interface {
	Context() string
	Errors() []error
}

// Error creates an error with additional context
//
// Deprecated: Use fmt.Errorf() instead using the `%w` format specifier.
func Error(err error, context string) error {
	if err == nil {
		return errors.New(context)
	}

	return fmt.Errorf("%s: %w", context, err)
}

// Errorf creates an error with additional formatted context
//
// Deprecated: Use fmt.Errorf() instead using the `%w` format specifier.
func Errorf(err error, format string, a ...interface{}) error {
	return Error(err, fmt.Sprintf(format, a...))
}

// Errors creates a list of errors with additional context
//
// Deprecated: Use fmt.Errorf() and errors.Join() instead.
func Errors(errs []error, context string) error {
	switch len(errs) {
	case 0:
		return errors.New(context)

	case 1:
		return fmt.Errorf("%s: %w", context, errs[0])

	default:
		return fmt.Errorf("%s:\n%w",
			context,
			errors.Join(errs...),
		)
	}
}

// Errorsf creates a list of errors with additional formatted context
//
// Deprecated: Use fmt.Errorf() and errors.Join() instead.
func Errorsf(errors []error, format string, a ...interface{}) error {
	return Errors(errors, fmt.Sprintf(format, a...))
}
