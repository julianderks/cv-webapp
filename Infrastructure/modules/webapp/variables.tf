variable "subscription_id" {
  type        = string
  description = "The Azure Subscription ID."
  nullable    = false
  sensitive   = true
}

variable "personal_name_short" {
  description = "Short version of your personal name"
  type        = string
}

variable "location_short" {
  description = "Short version of the location"
  type        = string
}

variable "project_name" {
  description = "Name of the project"
  type        = string
}

variable "location" {
  description = "Azure location"
  type        = string
}

variable "backend_docker_image_name" {
  description = "Docker image name with tag"
  type        = string
}


variable "github_repository_url" {
  description = "Github repository URL"
  type        = string
}

variable "github_PAT" {
  type        = string
  sensitive   = false
  description = "Personal Access Token for Github"
}