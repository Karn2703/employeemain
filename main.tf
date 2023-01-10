provider "employee" {
  address = "http://localhost"
  port    = "5000"
  token   = "superSecretToken"
}

resource "employee_profile" "test" {
  name = "karanparmar"
  email = "anymain@gmail.com"
}
