variable "project_id" {
  description = "The location of the pushed docker image to use"
  type        = string
}

variable "image_uri" {
  description = "The location of the pushed docker image to use"
  type        = string
}

variable "service_name" {
  description = "The name of the cloud run service"
  type        = string
}

variable "region" {
  description = "The region of the cloud run service"
  type        = string
}

variable "cmd" {
  description = "The command that will be executed in the container"
  type        = string
}