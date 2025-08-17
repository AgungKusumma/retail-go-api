package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.POST("/login", loginHandler)
		v1.POST("/report/:context", reportHandler)
	}

	r.Run(":8080")
}

func loginHandler(c *gin.Context) {
	var loginData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid JSON"})
		return
	}

	if loginData.Username == "fajar" && loginData.Password == "1234" {
		c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Login success"})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "error", "message": "Invalid credentials"})
	}
}

func reportHandler(c *gin.Context) {
	contextType := c.Param("context")

	switch contextType {
	case "attendance":
		var data struct {
			Status string `json:"status"`
		}
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid attendance data"})
			return
		}

		statusLower := strings.ToLower(data.Status)
		if statusLower != "hadir" && statusLower != "tidak hadir" {
			c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid attendance data"})
			return
		}

		data.Status = statusLower
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "Attendance report received",
			"data":    map[string]string{"status": data.Status},
		})

	case "product":
		var data struct {
			Status string `json:"status"`
		}
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid product data"})
			return
		}

		statusLower := strings.ToLower(data.Status)
		if statusLower != "tersedia" && statusLower != "tidak tersedia" {
			c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid product data"})
			return
		}

		data.Status = statusLower
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "Product report received",
			"data":    map[string]string{"status": data.Status},
		})

	case "promo":
		var data struct {
			Nama        string  `json:"nama"`
			HargaNormal float64 `json:"harga_normal"`
			HargaPromo  float64 `json:"harga_promo"`
		}
		if err := c.ShouldBindJSON(&data); err != nil || data.Nama == "" || data.HargaNormal <= 0 || data.HargaPromo <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid promo data"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "Promo report received",
			"data": map[string]interface{}{
				"nama":         data.Nama,
				"harga_normal": data.HargaNormal,
				"harga_promo":  data.HargaPromo,
			},
		})

	default:
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid context type"})
	}
}
