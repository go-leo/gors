package gors

type Validator interface {
	Validate() error
}

func Validate(req interface{}) error {
	switch v := req.(type) {
	case Validator:
		return v.Validate()
	default:
		return nil
	}
}
