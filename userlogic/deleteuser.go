package userlogic

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"usersenaryo/databaseOP"
)

func HandleUserDelete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		id, _ := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/delete-user/"))
		sonuc, hata := databaseOP.Delete(id).RowsAffected()
		fmt.Println("bu hatamız", hata)
		if sonuc == 0 {
			w.WriteHeader(http.StatusBadRequest)
			notFound, _ := json.Marshal(map[string]string{"notFound": "kullanıcı bulunamadı"})
			fmt.Fprintf(w, string(notFound))
			return
		}

		ok, _ := json.Marshal(map[string]string{"status": "başarıyla silindi"})
		fmt.Fprintf(w, string(ok))
		w.WriteHeader(http.StatusOK)
		return
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

}
