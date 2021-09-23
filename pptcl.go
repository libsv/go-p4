package p4

// ClientError defines an error type that can be returned to handle client errors.
type ClientError struct {
	ID      string `json:"id" example:"e97970bf-2a88-4bc8-90e6-2f597a80b93d"`
	Code    string `json:"code" example:"N01"`
	Title   string `json:"title" example:"not found"`
	Message string `json:"message" example:"unable to find foo when loading bar"`
}
