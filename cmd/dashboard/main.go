package main

import (
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"strconv"
	"sort"

	"github.com/gin-gonic/gin"
)

// Helper function to extract the leading number from a folder name (e.g., "1071-greatest-...")
func getProblemID(dirName string) int {
	// Find where the first hyphen is
	hyphenIdx := strings.Index(dirName, "-")
	if hyphenIdx == -1 {
		return 0 // Fallback if there's no hyphen
	}

	// Extract the number substring
	numStr := dirName[:hyphenIdx]

	// Convert it to an actual integer
	id, err := strconv.Atoi(numStr)
	if err != nil {
		return 0 // Fallback if it's not a valid number
	}
	return id
}

func main() {
	r := gin.Default()

	// Load HTML templates from your custom structure
	r.LoadHTMLGlob("cmd/dashboard/templates/*")

	// 1. HOME: Scan the 'problems' directory and list them on the sidebar
	// r.GET("/", func(c *gin.Context) {
	// 	files, err := os.ReadDir("./problems")
	// 	if err != nil {
	// 		c.String(http.StatusInternalServerError, "Failed to read problems directory: %v", err)
	// 		return
	// 	}

	// 	var problemList []string
	// 	for _, f := range files {
	// 		if f.IsDir() && !strings.HasPrefix(f.Name(), ".") {
	// 			problemList = append(problemList, f.Name())
	// 		}
	// 	}

	// 	c.HTML(http.StatusOK, "index.html", gin.H{
	// 		"Questions": problemList,
	// 	})
	// })
    r.GET("/", func(c *gin.Context) {
		files, err := os.ReadDir("./problems")
		if err != nil {
			c.String(http.StatusInternalServerError, "Failed to read problems directory: %v", err)
			return
		}

		var problemList []string
		for _, f := range files {
			if f.IsDir() && !strings.HasPrefix(f.Name(), ".") {
				problemList = append(problemList, f.Name())
			}
		}

		// ✨ THE FIX: Sort the slice using our custom natural-number logic
		sort.Slice(problemList, func(i, j int) bool {
			return getProblemID(problemList[i]) < getProblemID(problemList[j])
		})

		c.HTML(http.StatusOK, "index.html", gin.H{
			"Questions": problemList,
		})
	})	

	// 2. GET PROBLEM CONTENTS: Read README.md and main.go for a selected question
	r.GET("/problem/:dir", func(c *gin.Context) {
		dir := c.Param("dir")
		
		readmePath := filepath.Join("problems", dir, "README.md")
		mainPath := filepath.Join("problems", dir, "main.go")

		readmeContent, _ := os.ReadFile(readmePath)
		mainContent, _ := os.ReadFile(mainPath)

		// Return an HTML fragment that HTMX swaps into our 2-pane workspace layout
		c.HTML(http.StatusOK, "workspace_fragment.html", gin.H{
			"Dir":    dir,
			"Readme": string(readmeContent),
			"Code":   string(mainContent),
		})
	})

	// 3. RUN TESTS: Execute tests inside the browser
	r.POST("/run/:dir", func(c *gin.Context) {
		dir := c.Param("dir")
		targetPath := filepath.Join("./problems", dir)

		// Run 'go test -v' specifically targeting that problem package
		cmd := exec.Command("go", "test", "-v", "./"+targetPath)
		output, _ := cmd.CombinedOutput()

		// Return raw terminal text to render in a pre block
		c.String(http.StatusOK, string(output))
	})

	// 4. NEW: EXECUTE MAIN PROGRAM (Runs go run main.go)
	r.POST("/execute/:dir", func(c *gin.Context) {
		dir := c.Param("dir")
		// Target the specific main.go inside the problem directory
		mainFile := filepath.Join("./problems", dir, "main.go")

		cmd := exec.Command("go", "run", mainFile)
		output, _ := cmd.CombinedOutput()

		c.String(http.StatusOK, string(output))
	})

	log.Println("Dashboard running on http://localhost:8080")
	r.Run(":8080")
}