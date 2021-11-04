package foo

type Foo struct {
	Foo string `json:"foo" binding:"required"`
	Bar string `json:"bar" binding:"required"`
}
