package models

//UsuarioSeguido es la relacion entre seguidores
type UsuarioSeguido struct {
	UsuarioID      string `bson:"usuario_id" json:"usuario_id"`
	UsuarioSeguido string `bson:"usuarioseguido_id" json:"usuarioseguido_id"`
}
