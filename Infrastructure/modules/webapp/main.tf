#### Create a resource group
resource "azurerm_resource_group" "rg" {
  name     = "${var.personal_name_short}-${var.location_short}-${var.project_name}-rg"
  location = var.location
}

#### Create a container registry
resource "azurerm_container_registry" "acr" {
  name                = "${var.personal_name_short}${var.location_short}${var.project_name}acr"
  resource_group_name      = azurerm_resource_group.rg.name
  location                 = azurerm_resource_group.rg.location
  sku                      = "Standard"
  admin_enabled            = true
}

#### Create ACR Task for CI
resource "azurerm_container_registry_task" "backend_build" {
  name                = "build-backend-image"
  container_registry_id = azurerm_container_registry.acr.id

  platform {
    os = "Linux"
  }

  docker_step {
    dockerfile_path      = "Dockerfile"
    context_path         = "${var.github_repository_url}#main:backend"
    context_access_token = var.github_PAT
    image_names          = [var.backend_docker_image_name]
  }

  source_trigger {
    name           = "trigger-on-commit-to-image-source"
    events         = ["commit"]
    repository_url = var.github_repository_url
    source_type    = "Github"
    branch         = "main"
    enabled        = true
    
    authentication {
      # somehow if this is 0 the azdo token is considered instantly invalid in now and future runs while it does actually work..
      expire_in_seconds = 9999999
      token_type        = "PAT"
      token             = var.github_PAT
    }
  }
}

#### Create a service plan
resource "azurerm_service_plan" "asp" {
  name                = "${var.personal_name_short}-${var.location_short}-${var.project_name}-splan"
  location            = azurerm_resource_group.rg.location
  resource_group_name = azurerm_resource_group.rg.name
  os_type             = "Linux"
  sku_name            = "B1"
}

#### Create a web app
resource "azurerm_linux_web_app" "linux_webapp" {
  name                          = "${var.personal_name_short}-${var.location_short}-${var.project_name}-wapp"
  location                      = azurerm_resource_group.rg.location
  resource_group_name           = azurerm_resource_group.rg.name
  service_plan_id               = azurerm_service_plan.asp.id
  public_network_access_enabled = true

  site_config {
    application_stack {
      docker_image_name        = var.backend_docker_image_name
      docker_registry_url      = "https://${azurerm_container_registry.acr.login_server}"
      docker_registry_username = azurerm_container_registry.acr.admin_username
      docker_registry_password = azurerm_container_registry.acr.admin_password

    }
  }

  # Optional: Environment Variables
  app_settings = {
    MY_ENV_VARIABLE = "value"
    DOCKER_ENABLE_CI = "true" # This will allow us to configure webhook for CD: redeploy the app when a new image lands in ACR
  }
}

resource "azurerm_container_registry_webhook" "webhook" {
  service_uri         = "https://${azurerm_linux_web_app.linux_webapp.site_credential.0.name}:${azurerm_linux_web_app.linux_webapp.site_credential.0.password}@${azurerm_linux_web_app.linux_webapp.name}.scm.azurewebsites.net/api/registry/webhook"
  location            = azurerm_resource_group.rg.location
  resource_group_name = azurerm_resource_group.rg.name
  registry_name       = azurerm_container_registry.acr.name
  name                = "${replace(azurerm_linux_web_app.linux_webapp.name, "/-|_|\\W/", "")}hook" 
  actions             = ["push"]
}
