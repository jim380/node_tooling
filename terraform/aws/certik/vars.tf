variable "AWS_ACCESS_KEY" {
  default = ""
}

variable "AWS_SECRET_KEY" {
  default = ""
}

variable "AWS_REGION" {
  default = "us-east-2"
}

variable "AMIS" {
  type = map(string)
  default = {
    us-east-2 = "ami-0a040c35ca945058a"
    // eu-west-1 = ""
  }
}

variable "PATH_TO_PRIVATE_KEY" {
  default = "ctk"
}

variable "PATH_TO_PUBLIC_KEY" {
  default = "ctk.pub"
}

variable "INSTANCE_USERNAME" {
  default = "ubuntu"
}

variable "NODE_INSTANCE_MODEL" {
  default = "t3.medium"
}

variable "NODE_INSTANCE_VOLUME" {
  description = "EBS volume initiated on the node"
  default     = 500 # 500 GB
}


variable "protocol" {
  description = "name of the protocol (in lowercase)"
  type        = string
  default     = "certik"
}

variable "chain_id" {
  type        = string
  default     = ""
}

variable "moniker" {
  type        = string
  default     = "ctk-test"
}