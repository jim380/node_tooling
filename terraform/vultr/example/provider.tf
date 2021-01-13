provider "vultr" {
  api_key     = var.api_key
  rate_limit  = 700
  retry_limit = 3
}

provider "github" {
  #   token = "${var.github_token}"
  #   owner = "${var.github_owner}"
}