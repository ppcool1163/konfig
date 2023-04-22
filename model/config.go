package model

type Config struct {
	Version     int        `json:"version"`
	Name        string     `json:"name"`
	Environment string     `json:"environment"`
	Variables   []Variable `json:"environmentVariables"`
}

type Variable struct {
	Name  string `json:"name"`
	Value string `json:"value"`
	Type  string `json:"type"`
}
