// Package vkid ...
package vkid

// SilentAuthPayload struct.
type SilentAuthPayload struct {
	Auth              int        `json:"auth"`
	Token             string     `json:"token"`
	LoadExternalUsers bool       `json:"loadExternalUsers"`
	TTL               int        `json:"ttl"`
	Type              string     `json:"type"`
	User              SilentUser `json:"user"`
	UUID              string     `json:"uuid"`

	OAuthProvider  string       `json:"oauthProvider,omitempty"`
	ExternalUser   ExternalUser `json:"external_user,omitempty"`
	IsRegistration bool         `json:"is_registration,omitempty"`
}

// SilentUser struct.
type SilentUser struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Avatar    string `json:"avatar"`
	Phone     string `json:"phone"`
}

// ExternalUser struct.
type ExternalUser struct {
	ID          string      `json:"id,omitempty"`
	Avatar      string      `json:"avatar,omitempty"`
	FirstName   string      `json:"first_name"`
	LastName    string      `json:"last_name"`
	Phone       string      `json:"phone"`
	BorderColor string      `json:"borderColor,omitempty"`
	Payload     interface{} `json:"payload,omitempty"`
}
