package main

import (
	"fmt"
	"regexp"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
)

func main() {

	var err error
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.POST("/", post)
	err = router.Run("localhost:8080")
	if err != nil {
		panic(err)
	}

}

func post(c *gin.Context) {
	var form struct {
		Input   string `form:"input" binding:"required"`
		Pattern string `form:"pattern" binding:"required"`
	}
	if err := c.ShouldBind(&form); err != nil {
		c.String(400, "Failed to bind request: %v", err)
		return
	}

	machedtext := prossesinput(form.Input, form.Pattern)
	hilitedtext := bolt(form.Input, form.Pattern)
	result := fmt.Sprintf("Matched Lines:\n%s\n\nHighlighted Text:\n%s", strings.Join(machedtext, "\n"), hilitedtext)
	c.String(200, result)
}

func prossesinput(input, pattern string) []string {
	inputs := strings.Fields(input)
	var machedinput []string
	var wg sync.WaitGroup
	var mu sync.Mutex
	for _, word := range inputs {
		wg.Add(1)
		go func(word string) {
			defer wg.Done()
			if matchpattern(word, pattern) {
				mu.Lock()
				machedinput = append(machedinput, word)
				mu.Unlock()
			}
		}(word)

	}
	wg.Wait()
	return machedinput
}

func matchpattern(input, pattern string) bool {
	t, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Printf("Invalid pattern: %v\n", err)
		return false
	}
	return t.MatchString(input)
}

func bolt(input string, patern string) string {
	return strings.ReplaceAll(input, patern, fmt.Sprintf("<b>%s<b>", patern))
}
