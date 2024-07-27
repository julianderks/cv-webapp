package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Define a struct to represent the JSON object
type Description struct {
	Introduction string `json:"introduction"`
	UsedWhere    string `json:"used_where"`
}

type Tool struct {
	Name                    string      `json:"name"`
	IconifyIconName         string      `json:"iconify_icon_name,omitempty"`
	CustomIconLink          string      `json:"custom_icon_link,omitempty"`
	ExperienceLevel         int         `json:"experience_level"`           // Experience level field
	IconifyIconHeightFactor float64     `json:"iconify_icon_height_factor"` // Icon height factor field
	Area                    string      `json:"area"`                       // Categorization area
	Description             Description `json:"description"`                // Description field
	DocsLink                string      `json:"docs_link"`                  // Documentation link
}

func readTools() ([]Tool, error) {
	var tools []Tool

	// Read the tools directory
	files, err := os.ReadDir("tools")
	if err != nil {
		return nil, fmt.Errorf("failed to read tools directory: %w", err)
	}

	// Iterate over each file in the directory
	for _, file := range files {
		if file.IsDir() {
			continue // Skip directories
		}

		filePath := filepath.Join("tools", file.Name())
		// Read the file content
		data, err := ioutil.ReadFile(filePath)
		if err != nil {
			return nil, fmt.Errorf("failed to read file %s: %w", file.Name(), err)
		}

		// Decode the JSON content into a Tool struct
		var tool Tool
		if err := json.Unmarshal(data, &tool); err != nil {
			return nil, fmt.Errorf("failed to unmarshal JSON from file %s: %w", file.Name(), err)
		}

		// Append the Tool struct to the array
		tools = append(tools, tool)
	}

	return tools, nil
}

func main() {
	e := echo.New()

	// Enable CORS middleware
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:5173"},
		AllowMethods: []string{http.MethodGet, http.MethodPost},
	}))

	// Serve static files from the /logos directory
	e.Static("/logos", "logos")
	// Serve static files from the /tools/custom_icons directory
	e.Static("/tools/custom_icons", "tools/custom_icons")

	// Route for the root path
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// Route for the /api path
	e.GET("/api", func(c echo.Context) error {
		// Create a slice of Tool objects with URLs to the logos, experience levels, height factors, and areas
		_, err := readTools()
		if err != nil {
			log.Fatal(err)
		}

		// Return the slice as a JSON response
		return c.String(http.StatusOK, "Hello, World from /api!")
	})

	// Start the server on port 8000
	e.Logger.Fatal(e.Start(":8000"))
}
