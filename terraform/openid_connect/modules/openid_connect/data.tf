data "tls_certificate" "oid_provider" {
  for_each = local.enabled_providers

  url = format("https://%s", each.value.identity_provider_url)
}
