package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ArrayReq struct {
	Numbers []int `json:"numbers" binding:"required"`
	Target  int   `json:"target" binding:"required"`
}

type Response struct {
	Solutions [][]int `json:"solutions"`
}

func findPairs(numbers []int, target int) [][]int {
	pairs := make(map[int]int)
	var result [][]int

	for i, num := range numbers {
		if index, found := pairs[target-num]; found {
			result = append(result, []int{index, i})
		}
		pairs[num] = i
	}
	return result
}

func findPairsHandler(ctx *gin.Context) {
	var request ArrayReq
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
	}

	solutions := findPairs(request.Numbers, request.Target)

	ctx.JSON(http.StatusOK, gin.H{"solutions": solutions})
}

func main() {
	router := gin.Default()

	router.POST("find-pairs", findPairsHandler)

	router.Run(":8080")
}
