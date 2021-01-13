data "github_release" "node_exporter" {
  repository  = "node_exporter"
  owner       = "prometheus"
  retrieve_by = "latest"
  # retrieve_by = "tag"
  # release_tag = "v1.0.0"
}

data "github_release" "prometheus" {
  repository  = "prometheus"
  owner       = "prometheus"
  retrieve_by = "latest"
}