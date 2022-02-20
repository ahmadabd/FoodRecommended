package validation

type Validator interface {
	Validate(i interface{}) error
}