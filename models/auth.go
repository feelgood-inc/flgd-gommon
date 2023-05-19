package models

type LoginData struct {
	User         User   `json:"user"`
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token" omitempty:"true"`
}
