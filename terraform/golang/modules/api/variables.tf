variable "project_id" {
  description = "The location of the pushed docker image to use"
  type        = string
}

variable "api_name" {
  description = "The location of the pushed docker image to use"
  type        = string
}

variable "region" {
  description = "The location to deploy the API Gateway into"
  type        = string
}

variable "openapi_spec" {
  description = "The contents of the OpenAPI spec to use in yaml/json format"
  type        = string
}

variable "labels" {
  description = "Labels for the API resource"
  type        = map(string)
  default     = {}
}