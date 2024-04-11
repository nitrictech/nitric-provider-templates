// A Google Cloud Api Gateway Resource

resource "google_api_gateway_api" "api" {
  provider = google
  project  = var.project_id
  api_id   = var.api_name
  display_name = "example-api"
}

resource "google_api_gateway_api_config" "api_config" {
  provider = google
  api      = google_api_gateway_api.api.api_id
  api_config_id = "${var.api_name}-conf"
  openapi_documents {
    document {
      path     = "spec.yaml"
      contents =  base64encode(var.openapi_spec)
    }
  }
  dynamic "labels" {
        for_each = var.labels

        content {
          name  = env.key
          value = env.value
        }
      }
}

resource "google_api_gateway_gateway" "gateway" {
  provider = google
  project  = var.project_id
  gateway_id = "${var.api_name}-gw"
  api_config = google_api_gateway_api_config.api_config.id
  location   = var.region
}
