package main

import (
	//"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"k8s.io/api/admission/v1"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/json"
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
			var sts appsv1.StatefulSet
			if err := json.Unmarshal(ar.Request.Object.Raw, &sts); err != nil {
				panic(err)
			}
			jsonPatch := v1.PatchTypeJSONPatch
			patch, err := updatePatch(&sts)
			if err != nil {
				panic(err)
			}
			fmt.Println(string(patch))
			//encodeString := base64.StdEncoding.EncodeToString(patch)

			admissionResponse := &v1.AdmissionResponse{
				UID:       ar.Request.UID,
				Allowed:   true,
				PatchType: &jsonPatch,
				Patch:     patch,
			}
			admissionReview := v1.AdmissionReview{}
			admissionReview.TypeMeta = ar.TypeMeta

			admissionReview.Response = admissionResponse
			c.JSON(200, admissionReview)
		}

	})
	_ = r.RunTLS(":5678", "./secret/server.crt", "./secret/server.key")

}
func updatePatch(sts *appsv1.StatefulSet) ([]byte, error) {
	annotations := sts.GetAnnotations()
	annotations["aaaaa"] = "666"
	return json.Marshal([]patchOperation{{
		Op:    "add",
		Path:  "/metadata/annotations",
		Value: annotations,
	}})
}

type patchOperation struct {
	Op    string      `json:"op"`
	Path  string      `json:"path"`
	Value interface{} `json:"value,omitempty"`
}
