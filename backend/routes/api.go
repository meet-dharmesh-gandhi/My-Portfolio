package routes

import (
	"net/http"

	"github.com/my-portfolio/backend/utils"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	utils.JSON(w, 200, map[string]string{
		"status": "looks ok",
	})
}

func Dummy(w http.ResponseWriter, r *http.Request) {
	utils.JSON(w, 200, map[string]interface{}{
		"message": "Some Message",
		"status":  true,
	})
}
