package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Define a struct to represent the JSON object
type Tool struct {
	Tool            string `json:"tool"`
	IconPath        string `json:"icon_path"`
	ExperienceLevel int    `json:"experience_level"` // Add the experience level field
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

	// Route for the root path
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// Route for the /api path
	e.GET("/api", func(c echo.Context) error {
		// Create a slice of Tool objects with URLs to the logos and experience levels
		tools := []Tool{
			{Tool: "Terraform", IconPath: "http://localhost:8000/icons/terraform.svg", ExperienceLevel: 4},
			{Tool: "Terraform", IconPath: "http://localhost:8000/icons/terraform.svg", ExperienceLevel: 4},
			{Tool: "Terraform", IconPath: "http://localhost:8000/icons/terraform.svg", ExperienceLevel: 4},
			{Tool: "Terraform", IconPath: "http://localhost:8000/icons/terraform.svg", ExperienceLevel: 4},
			{Tool: "Terraform", IconPath: "http://localhost:8000/icons/terraform.svg", ExperienceLevel: 4},
			{Tool: "Terraform", IconPath: "http://localhost:8000/icons/terraform.svg", ExperienceLevel: 4},
			// {Tool: "Helm", IconPath: "http://localhost:8000/logos/helm.png", ExperienceLevel: 3},
			// // Add more tools with their experience levels here
			// {Tool: "Docker", IconPath: "http://localhost:8000/logos/docker.png", ExperienceLevel: 3},
			// {Tool: "Kubernetes", IconPath: "http://localhost:8000/logos/docker.png", ExperienceLevel: 3},
			// {Tool: "Python", IconPath: "http://localhost:8000/logos/python.png", ExperienceLevel: 4},
			// {Tool: "Golang", IconPath: "http://localhost:8000/logos/golang.png", ExperienceLevel: 1},
			// // Continue adding more tools as needed
			// {Tool: "Pyspark", IconPath: "http://localhost:8000/logos/golang.png", ExperienceLevel: 1},
			// {Tool: "Helmfile", IconPath: "http://localhost:8000/logos/helm.png", ExperienceLevel: 3},
			// {Tool: "Grafan", IconPath: "http://localhost:8000/logos/golang.png", ExperienceLevel: 1},
			// {Tool: "Prometheus", IconPath: "http://localhost:8000/logos/golang.png", ExperienceLevel: 1},
			// {Tool: "GCP", IconPath: "http://localhost:8000/logos/golang.png", ExperienceLevel: 1},
			// {Tool: "Azure", IconPath: "http://localhost:8000/logos/golang.png", ExperienceLevel: 1},
			// {Tool: "Git", IconPath: "http://localhost:8000/logos/golang.png", ExperienceLevel: 1},
			// {Tool: "Databricks", IconPath: "http://localhost:8000/logos/golang.png", ExperienceLevel: 1},
			// {Tool: "Power BI", IconPath: "http://localhost:8000/logos/golang.png", ExperienceLevel: 1},
			// {Tool: "PostgreSQL", IconPath: "http://localhost:8000/logos/golang.png", ExperienceLevel: 1},
			// {Tool: "SQL", IconPath: "http://localhost:8000/logos/golang.png", ExperienceLevel: 1},
			// {Tool: "Linux", IconPath: "http://localhost:8000/logos/golang.png", ExperienceLevel: 1},
			// {Tool: "HTML", IconPath: "http://localhost:8000/logos/golang.png", ExperienceLevel: 1},
			// {Tool: "CSS", IconPath: "http://localhost:8000/logos/golang.png", ExperienceLevel: 1},
			// {Tool: "Javascript", IconPath: "http://localhost:8000/logos/golang.png", ExperienceLevel: 1},
			// {Tool: "Svelte", IconPath: "http://localhost:8000/logos/golang.png", ExperienceLevel: 1},
		}

		// Return the slice as a JSON response
		return c.JSON(http.StatusOK, tools)
	})

	// Start the server on port 8000
	e.Logger.Fatal(e.Start(":8000"))
}
