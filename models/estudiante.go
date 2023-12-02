package models

/*
es mejor conservar un estándar entre las etiquetas de json y db para no tener problemas al parsear
de json a db en el método ActualizarUnAmigo
*/
type Estudiante struct { //Amigo==Estudiante
	Id        int    `db:"id" json:"id"` //etiquetas
	Nombre    string `db:"nombre" json:"nombre"`
	Edad      uint   `db:"edad" json:"edad"`
	Grado     string `db:"grado" json:"grado"`
	Jornada   string `db:"jornada" json:"jornada"`
	Direccion string `db:"direccion" json:"direccion"`
	Telefono  uint64 `db:"telefono" json:"telefono"` //solo numeror de 7 cifras
	Correo    string `db:"correo" json:"correo"`
}
