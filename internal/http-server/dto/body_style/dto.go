package body_style

type BodyStyleDTO struct {
	ID    *uint16 `json:"id,omitempty"`
	Name  *string `json:"name,omitempty"`
	Value *uint16 `json:"value,omitempty"`
}
