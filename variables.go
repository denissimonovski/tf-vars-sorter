package tf_vars_sorter

type Var struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description"`

	// Default is an approximate representation of the default value in
	// the native Go type system. The conversion from the value given in
	// configuration may be slightly lossy. Only values that can be
	// serialized by json.Marshal will be included here.
	Default  interface{} `json:"default,omitempty"`
	Required bool        `json:"required,omitempty"`
}
