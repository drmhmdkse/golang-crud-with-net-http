package userlogic

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"usersenaryo/databaseOP"
	"usersenaryo/models"
)

func HandleUpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == "POST" {
		id, _ := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/update-user/"))
		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Println("satır 17 :", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// UNMARSHALL işlemi #json verisini struc'a çevirme
		var user models.User
		hata := json.Unmarshal(body, &user)
		if hata != nil {
			log.Println("json verisi çevrilemedi", hata)
			w.WriteHeader(http.StatusBadRequest)
			wrongData, _ := json.Marshal(map[string]string{"message": "hatalı json verisi tespit edildi"})
			fmt.Fprintf(w, string(wrongData))
			return
		}

		if databaseOP.Exist(user.Username) {
			w.WriteHeader(http.StatusBadRequest)
			wrongData, _ := json.Marshal(map[string]string{"message": "bu kullanıcı adı zaten kullanılıyor"})
			fmt.Fprintf(w, string(wrongData))
			return
		}

		result := databaseOP.Update(id, user.Username, user.FirstName, user.Password)

		if i, _ := result.RowsAffected(); i == 0 {
			w.WriteHeader(http.StatusBadRequest)
			notFound, _ := json.Marshal(map[string]string{"notFound": "kullanıcı bulunamadı"})
			fmt.Fprintf(w, string(notFound))
			return
		}

		ok, _ := json.Marshal(map[string]string{"status": "başarıyla güncellendi"})
		fmt.Fprintf(w, string(ok))
		w.WriteHeader(http.StatusOK)

	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		notok, _ := json.Marshal(map[string]string{"status": "Method not allowed"})
		fmt.Fprintf(w, string(notok))
	}

}
