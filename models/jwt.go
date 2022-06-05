package models

type FeelgoodJWTPayload struct {
	UID       string   `json:"uid"`
	Email     string   `json:"email"`
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	Type      string   `json:"type"`
	Roles     []string `json:"roles"`
}

type FeelgoodJWTCustomClaims struct {
	Email    string `json:"email"`
	Type     string `json:"type"`
	Provider string `json:"provider"`
	UID      string `json:"uid"`
}

func (f FeelgoodJWTCustomClaims) Valid() error {
	//TODO implement me
	panic("implement me")
}
