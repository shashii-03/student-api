package students

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/te-shashikant/student-api/internal/types"
	"github.com/te-shashikant/student-api/internal/utils/response"
)


func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var student types.Student

		err:=json.NewDecoder(r.Body).Decode(&student)
		 if errors.Is(err, io.EOF){
			response.WriteJson(w, http.StatusBadRequest, err.Error())
			return
		 }

		 response.WriteJson(w, http.StatusCreated, map[string]string{"status":"ok"})

	}
}