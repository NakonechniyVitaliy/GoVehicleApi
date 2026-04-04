package body_style

// BodyStylePayload is the request body for POST/PUT /body-style
type BodyStylePayload struct {
	BodyStyle struct {
		Name  *string `json:"name"`
		Value *uint16 `json:"value"`
	} `json:"BodyStyle"`
}
