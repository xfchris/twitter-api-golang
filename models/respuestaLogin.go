package models

//RespuestaLogin es el modelo del token que se genera
type RespuestaLogin struct {
	Token string `json:"token,omitempty"`
}