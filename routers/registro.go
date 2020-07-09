package routers

import (
	"encoding/json"
	"net/http"

	"github.com/developerjfsi/golang/bd"
	"github.com/developerjfsi/golang/models"
)

/*Registro es para crear */
func Registro(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "El email del usuario es requerido "+err.Error(), 400)
		return
	}

	if len(t.Email) < 6 {
		http.Error(w, "Debe contener al menos 6 caracteres la contraseña "+err.Error(), 400)
		return
	}

	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)
	if encontrado == true {
		http.Error(w, "Ya existe un usuario registrado con ese email "+err.Error(), 400)
		return
	}

	_, status, err := bd.InsertoRegistro(t)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar realizar el registro de usuario "+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado insertar el registro del usuario "+err.Error(), 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
