output "container_app_identity" {
  value       = azurerm_container_app.app.identity.0.principal_id
  description = "The managed identity of this container app"
}