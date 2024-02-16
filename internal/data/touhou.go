package data

// Touhou represents a Touhou character.
type Touhou struct {
	ID        int64    `json:"id"`
	Name      string   `json:"name"`
	Species   string   `json:"species"`
	Abilities []string `json:"abilities"`
}
