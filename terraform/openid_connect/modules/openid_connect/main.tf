resource "aws_iam_openid_connect_provider" "oid_provider" {
  for_each = local.enabled_providers

  url = format("https://%s", each.value.identity_provider_url)

  client_id_list = [
    each.value.audience
  ]

  thumbprint_list = [
    data.tls_certificate.oid_provider[each.key].certificates.0.sha1_fingerprint
  ]
}

data "aws_iam_policy_document" "assuming_role" {

  dynamic "statement" {
    for_each = local.enabled_providers

    content {
      actions = ["sts:AssumeRoleWithWebIdentity"]

      principals {
        type = "Federated"
        identifiers = [
          aws_iam_openid_connect_provider.oid_provider[statement.key].arn
        ]
      }

      condition {
        test     = "StringLike"
        variable = format("%s:aud", statement.value.identity_provider_url)

        values = [
          statement.value.audience
        ]
      }

      condition {
        test     = "StringLike"
        variable = format("%s:sub", statement.value.identity_provider_url)

        values = compact([
          # Github
          statement.key == "github" ? format(
            "repo:%s/%s:*",
            statement.value.workspace_name,
            statement.value.repository
          ) : null,
          # Gitlab
          statement.key == "gitlab" ? format(
            "*:%s:*:*:*:*",
            join("/",
              slice(
                split("/", replace(statement.value.group_url, "https://", "")),
                1,
                length(
                  split("/", replace(statement.value.group_url, "https://", ""))
                )
              )
            )
          ) : null,
          # Bitbucket
          statement.key == "bitbucket" ? format(
            "%s:*",
            statement.value.repository_uuid
          ) : null
        ])
      }

    }
  }
}

resource "aws_iam_role" "assuming_role" {
  name               = "identity_provider_assume_role"
  assume_role_policy = data.aws_iam_policy_document.assuming_role.json

  inline_policy {
    name = "identity_provider_permissions"

    policy = jsonencode({
      Version = "2012-10-17",
      Statement = var.identity_provider_permissions
    })
  }
}
