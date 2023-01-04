connection "hibp" {
  plugin  = "hibp"

  # `api_key` - The API key to access the HIBP API. 
  # This is only required while querying `hibp_breached_account` and `hibp_paste` tables. 
  # This can also be set with the 'HIBP_API_KEY' environment variable
  # See https://haveibeenpwned.com/API/Key for more information on how to generate one
  # api_key = "03ef6bfxxxxxxxxxxxxxxx8ad568286b"
}
