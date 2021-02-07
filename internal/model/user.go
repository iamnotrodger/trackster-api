package model

// User Struct
type User struct {
	UserID     string `json:"user_id"`
	ProviderID string `json:"id"`
	Email      string `json:"email"`
	Name       string `json:"name"`
	GivenName  string `json:"given_name"`
	FamilyName string `json:"family_name"`
	Picture    string `json:"picture"`
}
