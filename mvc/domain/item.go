package domain

type Item struct {
	Name        uint64 `json:"name"`
	Description string `json:"description"`
	Homepage    string `json:"homepage"`
	Email       string `json:"email"`
}
