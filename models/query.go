package models

type Query struct {
	// Fields to be selected
	Fields []string
	// Associations to be loaded (e.g. "User", "User.Posts")
	Associations []string
}
