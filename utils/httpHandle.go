package utils

import (
	"net/http"
	"strconv"

	"github.com/ParaCAD/ParaCAD-backend/utils/logging"
)

func HandleErr(r *http.Request, w http.ResponseWriter, status int, err error) {
	w.WriteHeader(status)
	w.Header().Add("Content-Type", "text/plain")
	w.Write([]byte(strconv.Itoa(status) + " "))
	w.Write([]byte(err.Error()))
	logging.AddErrorToRequestCtx(r, err, status)
}

func SendFile(w http.ResponseWriter, file []byte, fileName string) {
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", "attachment; filename="+fileName)
	w.Header().Set("Content-Transfer-Encoding", "binary")
	w.Header().Set("Expires", "0")
	w.Write(file)
}
