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
  default = "terra"
}

variable "PATH_TO_PUBLIC_KEY" {
  default = "terra.pub"
}

variable "INSTANCE_USERNAME" {
  default = "ubuntu"
}

