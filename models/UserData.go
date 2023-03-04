package models

// SessionData holds information about the current user/session in a transaction.
// It is used to pass information about the current session to different components, usually services.
// In the context of an HTTP request, the request could have user data coming from the session.
type SessionData struct {
	UID       string   `json:"uid"`
	Email     string   `json:"email"`
	Token     string   `json:"token"`
	UserType  string   `json:"user_type"`
	UserRoles []string `json:"user_roles"`
}
