
terraform {
  required_providers {
    vultr = {
      source  = "vultr/vultr"
      version = "2.1.2"
    }
    github = {
      source = "hashicorp/github"
    }
  }
  required_version = ">= 0.13"
}
