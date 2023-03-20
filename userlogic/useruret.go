package userlogic

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"usersenaryo/databaseOP"
	"usersenaryo/models"
)

func HandleUserCreate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		//UNMARSALL İŞLEMİ
		var user models.User
		hata := json.Unmarshal(body, &user)
		if hata != nil {
			log.Println("satır 29 : json okunamadı", hata)
			w.WriteHeader(http.StatusBadRequest)
			wrongData, _ := json.Marshal(map[string]string{"message": "hatalı json verisi tespit edildi"})
			fmt.Fprintf(w, string(wrongData))
			return
		}

		if databaseOP.Exist(user.Username) {
			w.WriteHeader(http.StatusBadRequest)
			wrongData, _ := json.Marshal(map[string]string{"message": "kullanıcı zaten var"})
			fmt.Fprintf(w, string(wrongData))
			return
		}

		result := databaseOP.Create(user.Username, user.Username, user.Password)

		fmt.Println("bu sonuç:", result)

		w.WriteHeader(http.StatusOK) // başarılı durumda 200 kodu dön
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed) // metot uygun değilse 405 kodu dön
	}
}
