package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/yeinermart/proyecto/models"
	repositorio "github.com/yeinermart/proyecto/repository"
)

var (
	updateQuery = "UPDATE estudiantes SET %s WHERE id=:id;"
	deleteQuery = "DELETE FROM estudiantes WHERE id=$1;"
	selectQuery = "SELECT id, nombre, edad, grado, jornada, direccion, telefono, correo FROM estudiantes WHERE id=$1;"
	listQuery   = "SELECT id, nombre, edad, grado, jornada, direccion, telefono, correo FROM estudiantes limit $1 offset $2"
	createQuery = "INSERT INTO estudiantes (nombre, edad, grado, jornada, direccion, telefono, correo) VALUES (:nombre, :edad, :grado, :jornada, :direccion, :telefono, :correo) returning id;"
)

type Controller struct {
	repo repositorio.Repository[models.Estudiante]
}

func NewController(repo repositorio.Repository[models.Estudiante]) (*Controller, error) {
	if repo == nil {
		return nil, fmt.Errorf("para instanciar un controlador se necesita un repositorio no nulo")
	}
	return &Controller{
		repo: repo,
	}, nil
}

func (c *Controller) ActualizarUnAmigo(reqBody []byte, id string) error {
	nuevosValoresAmigo := make(map[string]any)
	err := json.Unmarshal(reqBody, &nuevosValoresAmigo)
	if err != nil {
		log.Printf("fallo al actualizar un amigo, con error: %s", err.Error())
		return fmt.Errorf("fallo al actualizar un amigo, con error: %s", err.Error())
	}

	if len(nuevosValoresAmigo) == 0 {
		log.Printf("fallo al actualizar un amigo, con error: %s", err.Error())
		return fmt.Errorf("fallo al actualizar un amigo, con error: %s", err.Error())
	}

	query := construirUpdateQuery(nuevosValoresAmigo)
	nuevosValoresAmigo["id"] = id
	err = c.repo.Update(context.TODO(), query, nuevosValoresAmigo)
	if err != nil {
		log.Printf("fallo al actualizar un amigo, con error: %s", err.Error())
		return fmt.Errorf("fallo al actualizar un amigo, con error: %s", err.Error())
	}
	return nil
}

func construirUpdateQuery(nuevosValores map[string]any) string {
	columns := []string{}
	for key := range nuevosValores {
		columns = append(columns, fmt.Sprintf("%s=:%s", key, key))
	}
	columnsString := strings.Join(columns, ",")
	return fmt.Sprintf(updateQuery, columnsString)
}

func (c *Controller) EliminarUnAmigo(id string) error {
	err := c.repo.Delete(context.TODO(), deleteQuery, id)
	if err != nil {
		log.Printf("fallo al eliminar un amigo, con error: %s", err.Error())
		return fmt.Errorf("fallo al eliminar un amigo, con error: %s", err.Error())
	}
	return nil
}

func (c *Controller) LeerUnAmigo(id string) ([]byte, error) {
	amigo, err := c.repo.Read(context.TODO(), selectQuery, id)
	if err != nil {
		log.Printf("fallo al leer un amigo, con error: %s", err.Error())
		return nil, fmt.Errorf("fallo al leer un amigo, con error: %s", err.Error())
	}

	amigoJson, err := json.Marshal(amigo)
	if err != nil {
		log.Printf("fallo al leer un amigo, con error: %s", err.Error())
		return nil, fmt.Errorf("fallo al leer un amigo, con error: %s", err.Error())
	}
	return amigoJson, nil
}

func (c *Controller) LeerAmigos(limit, offset int) ([]byte, error) {
	amigos, _, err := c.repo.List(context.TODO(), listQuery, limit, offset)
	if err != nil {
		log.Printf("fallo al leer amigos, con error: %s", err.Error())
		return nil, fmt.Errorf("fallo al leer amigos, con error: %s", err.Error())
	}

	jsonAmigos, err := json.Marshal(amigos)
	if err != nil {
		log.Printf("fallo al leer amigos, con error: %s", err.Error())
		return nil, fmt.Errorf("fallo al leer amigos, con error: %s", err.Error())
	}
	return jsonAmigos, nil
}

func (c *Controller) CrearAmigo(reqBody []byte) (int64, error) {
	nuevoAmigo := &models.Estudiante{}
	err := json.Unmarshal(reqBody, nuevoAmigo)
	if err != nil {
		log.Printf("fallo al crear un nuevo amigo, con error: %s", err.Error())
		return 0, fmt.Errorf("fallo al crear un nuevo amigo, con error: %s", err.Error())
	}

	valoresColumnasNuevoAmigo := map[string]any{
		"nombre":    nuevoAmigo.Nombre,
		"edad":      nuevoAmigo.Edad,
		"grado":     nuevoAmigo.Grado,
		"jornada":   nuevoAmigo.Jornada,
		"direccion": nuevoAmigo.Direccion,
		"telefono":  nuevoAmigo.Telefono,
		"correo":    nuevoAmigo.Correo,
	}

	nuevoId, err := c.repo.Create(context.TODO(), createQuery, valoresColumnasNuevoAmigo)
	if err != nil {
		log.Printf("fallo al crear un nuevo amigo, con error: %s", err.Error())
		return 0, fmt.Errorf("fallo al crear un nuevo amigo, con error: %s", err.Error())
	}
	return nuevoId, nil
}
