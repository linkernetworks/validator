package validator

type ValidationMap map[string]FieldValidation

func (v *ValidationMap) Add(vld FieldValidation) {
	(*v)[vld.Field] = vld
}

func (v ValidationMap) HasError() bool {
	for _, it := range v {
		if it.Error {
			return true
		}
	}
	return false
}

type FieldValidation struct {
	Error   bool   `json:"error"`
	Field   string `json:"field"`
	Message string `json:"message"`
}
