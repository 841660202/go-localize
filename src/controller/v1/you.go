package v1

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// @Summary Add a new pet to the store
// @Description get string by ID
// @Accept  json
// @Produce  json
// @Param   some_id     path    int     true        "Some ID"
// @Success 200 {string} string	"ok"
// @Router /testapi/get-string-by-int/{some_id} [get]
func GetStringByInt(c *gin.Context) {
	fmt.Println("1212")
}

// @Summary  get struct array by ID
// @Description get struct array by ID
// @Accept  json
// @Produce  json
// @Param   some_id     path    string     true        "Some ID"
// @Param   offset     query    int     true        "Offset"
// @Param   limit      query    int     true        "Limit"
// @Success 200 {string} string	"ok"
// @Router /testapi/get-struct-array-by-string/{some_id} [get]
func GetStructArrayByString(c *gin.Context) {

}

type Pet3 struct {
	ID int `json:"id"`
}
