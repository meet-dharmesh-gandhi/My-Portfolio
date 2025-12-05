package routes

import (
	"net/http"

	"github.com/meet-dharmesh-gandhi/my-portfolio/backend/utils"
)

func healthCheck(w http.ResponseWriter, r *http.Request) {
	utils.JSON(w, 200, map[string]string{
		"status": "looks ok",
	})
}

func dummy(w http.ResponseWriter, r *http.Request) {
	utils.JSON(w, 200, map[string]interface{}{
		"message": "Some Message",
		"status":  true,
	})
}

func checkBuilder(w http.ResponseWriter, r *http.Request) {
	utils.JSON(w, 200, map[string]interface{}{
		"builder": "Works!!",
	})
}
