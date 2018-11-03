provider "restapi" {
  uri                  = "https://whoami.lmhd.me"
  debug                = true
  write_returns_object = true
}

# This will make information about the user named "John Doe" available by finding him by first name
data "restapi_object" "name" {
  path         = "/name"
  search_key   = ""
  search_value = ""
  results_key  = "full_name"
  id_attribute = "preferred"
}

output "name" {
  value = "${data.restapi_object.name.api_data.preferred}"
}
