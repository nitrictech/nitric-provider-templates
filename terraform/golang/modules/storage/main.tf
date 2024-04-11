# A Google cloud storage bucket
resource "google_storage_bucket" "bucket" {
  name          = var.bucket_name
  location      = var.bucket_location
  project       = var.project_id
  storage_class = var.storage_class
}