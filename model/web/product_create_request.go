package web

type ProductCreateRequest struct {
	Name string `validate:"required" json:"name"`
}
