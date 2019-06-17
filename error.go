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

import "fmt"

// ContextError interface describes the simple type that is able to provide a
// textual context as well as the cause explaining the underlying error.
type ContextError interface {
	Context() string
	Cause() error
}

// wrappedError describes an error with added context information
type wrappedError struct {
	context string
	cause   error
}

func (e *wrappedError) Error() string {
	return fmt.Sprintf("%s: %v", e.context, e.cause)
}

func (e *wrappedError) Context() string {
	return e.context
}

func (e *wrappedError) Cause() error {
	return e.cause
}

// Error creates an error with additional context
func Error(err error, context string) error {
	return &wrappedError{context, err}
}

// Errorf creates an error with additional formatted context
func Errorf(err error, format string, a ...interface{}) error {
	return &wrappedError{fmt.Sprintf(format, a...), err}
}
