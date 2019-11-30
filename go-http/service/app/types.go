package main

type requestBody struct {
	Input       map[string]interface{} `json:"input"`
	SessionVars map[string]interface{} `json:"session_vars"`
}

type responseBody struct {
	Data   map[string]interface{} `json:"data"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}
