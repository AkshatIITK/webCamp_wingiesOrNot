package server

import (
	"net/http"
	model "wingiesOrNot/models"
	"wingiesOrNot/utils"

	"github.com/gin-gonic/gin"
)

func getReq2(c *gin.Context, groupedData map[string]model.Hall) {
	h := c.Param("hall")
	w := c.Param("wing")
	r := c.Param("room")

	if hall, ok := groupedData[h]; ok {
		if w == "" {
			c.JSON(200, hall)
		}
		if wing, ok := hall[w]; ok {
			if r == "" {
				c.JSON(200, wing)
			}
			if room, ok := wing[r]; ok {
				c.JSON(200, room)
				return
			}
		}
	}

	c.JSON(404, gin.H{"error": "Not Found"})
}

func postReq2(c *gin.Context, raw model.Students) {
	// Expected body struct of post req
	var reqBody model.WingiesOrNot
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	result, err := utils.WingiesOrNot(reqBody.Id1, reqBody.Id2, raw)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if result {
		c.String(http.StatusOK, "YES")
	} else {
		c.String(http.StatusOK, "NO")
	}
}

func getStudentDetailsByID(c *gin.Context, raw model.Students) model.Student {
	i := c.Param("id")
	var student model.Student
	for _, s := range raw {
		if s.Id == i {
			student = s
		}
	}
	if student.Id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id not found"})
		return model.Student{}
	}
	return student
}

func fetchWingies(c *gin.Context, groupedData map[string]model.Hall, student model.Student) {
	h := student.Hall
	w := student.Room[0:3]

	if hall, ok := groupedData[h]; ok {
		if w == "" {
			c.JSON(200, hall)
		}
		if wing, ok := hall[w]; ok {
			c.JSON(200, wing)
			return
		}
	}

	c.JSON(404, gin.H{"error": "Not Found"})
}
