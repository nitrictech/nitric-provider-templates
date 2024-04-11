

locals {
  is_bucket = var.resource_type == "bucket"
}

# Apply the IAM policy to the resource
resource "google_storage_bucket_iam_member" "bucket_iam_member" {
  count  = local.is_bucket ? 1 : 0
  bucket = var.resource_name
  role   = "roles/storage.objectViewer"
  member = "serviceAccount:${var.service_account.email}"
}

resource "google_pubsub_topic_iam_member" "topic_iam_member" {
  count  = local.is_bucket ? 0 : 1
  topic  = var.resource_name
  role   = "roles/pubsub.publisher"
  member = "user:email@example.com"
}