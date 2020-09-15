package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"k8s.io/api/admission/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/webhook", func(c *gin.Context) {
		var admissionResponse *v1.AdmissionResponse

		ar := v1.AdmissionReview{}
		err := c.BindJSON(&ar)
		if err != nil {
			fmt.Println(err.Error())
			admissionResponse = &v1.AdmissionResponse{
				Result: &metav1.Status{
					Message: err.Error(),
				},
			}
			c.JSON(400, admissionResponse)
		} else {
			admissionResponse := &v1.AdmissionResponse{
				UID:     ar.Request.UID,
				Allowed: true,
			}

			admissionReview := v1.AdmissionReview{}
			admissionReview.TypeMeta = ar.TypeMeta
			admissionReview.Response = admissionResponse
			c.JSON(200, admissionReview)
		}

	})
	_ = r.RunTLS(":80", "./secret/server.crt", "./secret/server.key")

}
