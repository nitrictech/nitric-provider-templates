output "roles" {
  value = {
    base_compute_role = google_project_iam_custom_role.base_compute_role.id
    bucket_read = google_project_iam_custom_role.bucket_reader_role.id
    bucket_write = google_project_iam_custom_role.bucket_writer_role.id
    bucket_delete = google_project_iam_custom_role.bucket_deleter_role.id
  }
  description = "Nitric custom IAM roles"
}