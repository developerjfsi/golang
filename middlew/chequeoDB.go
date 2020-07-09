package middlew

import (
	"net/http"

	"github.com/developerjfsi/golang/bd"
)

/*ChequeoDB es la función para chequear si existe*/
func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.ChequeoConnection() == 0 {
			http.Error(w, "Conexion perdida con la base de datos", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
