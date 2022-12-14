// Code generated by ogen, DO NOT EDIT.

package api

import (
	"github.com/go-faster/errors"

	"github.com/ogen-go/ogen/validate"
)

func (s ErrorCircular) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if s.Code.Set {
			if err := func() error {
				if err := s.Code.Value.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "code",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s ErrorCircularCode) Validate() error {
	switch s {
	case "Circular":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}
func (s ErrorForbidden) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if s.Code.Set {
			if err := func() error {
				if err := s.Code.Value.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "code",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s ErrorForbiddenCode) Validate() error {
	switch s {
	case "Forbidden":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}
func (s ErrorNotFound) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if s.Code.Set {
			if err := func() error {
				if err := s.Code.Value.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "code",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s ErrorNotFoundCode) Validate() error {
	switch s {
	case "NotFound":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}
func (s ErrorUnauthorized) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if err := s.Code.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "code",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s ErrorUnauthorizedCode) Validate() error {
	switch s {
	case "Unauthorized":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}

func (s SignInWrongUsernameResponse) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if s.Errorrr.Set {
			if err := func() error {
				if err := s.Errorrr.Value.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "errorrr",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s SignInWrongUsernameResponseErrorrr) Validate() error {
	switch s {
	case "WrongUsername":
		return nil
	case "other":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}
func (s TaskDependencies) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if s.Dependencies == nil {
			return errors.New("nil is invalid value")
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "dependencies",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s TaskRelationsList) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if s.Dependencies == nil {
			return errors.New("nil is invalid value")
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "dependencies",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s TaskTodo) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if s.Path == nil {
			return errors.New("nil is invalid value")
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "path",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s TasksList) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if s.Tasks == nil {
			return errors.New("nil is invalid value")
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "tasks",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
