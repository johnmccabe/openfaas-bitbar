package types

// Function TODO
type Function struct {
	Name            string `json:"name"`
	Image           string `json:"image"`
	InvocationCount int    `json:"invocationCount"`
	Replicas        int    `json:"replicas"`
}
