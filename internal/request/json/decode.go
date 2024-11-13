package json

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func DecodeRequestBody(w http.ResponseWriter, req *http.Request, res interface{}) error {
	decoder := json.NewDecoder(req.Body)
	decoder.DisallowUnknownFields() // Any extra fields added to the req.Body are not allowed

	err := decoder.Decode(res)

	var syntaxError *json.SyntaxError
	var unmarshalTypeError *json.UnmarshalTypeError
	var invalidUnmarshalError *json.InvalidUnmarshalError

	switch {
	case errors.As(err, &syntaxError):
		return fmt.Errorf("body contains badly-formed JSON (at character %d)", syntaxError.Offset)
	case errors.Is(err, io.ErrUnexpectedEOF):
		return errors.New("Body contains badly-formed JSON")
	case errors.As(err, &unmarshalTypeError):
		if unmarshalTypeError.Field != "" {
			return fmt.Errorf(
				"body contains incorrect JSON type for field %q",
				unmarshalTypeError.Field,
			)
		}
		return fmt.Errorf(
			"body contains incorrect JSON type (at character %d)",
			unmarshalTypeError.Offset,
		)
	// An io.EOF error will be returned by Decode() if the request body is empty. We
	// check for this with errors.Is() and return a plain-english error message
	// instead.
	case errors.Is(err, io.EOF):
		return errors.New("body must not be empty")
	// A json.InvalidUnmarshalError error will be returned if we pass a non-nil
	// pointer to Decode(). We catch this and panic, rather than returning an error
	// to our handler. At the end of this chapter we'll talk about panicking
	// versus returning errors, and discuss why it's an appropriate thing to do in // this specific situation.
	case errors.As(err, &invalidUnmarshalError):
		panic(err)

	default:
		return err
	}
}

