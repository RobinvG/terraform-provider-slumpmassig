terraform {
  required_providers {
    slumpmassig = {
      source = "registry.terraform.io/robing/slumpmassig"
    }
  }
}

provider "slumpmassig" {}

resource "slumpmassig_products" "example" {
}
