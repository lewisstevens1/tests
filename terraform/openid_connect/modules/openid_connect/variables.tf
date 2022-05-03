variable "github" {
  type = object({
    enabled        = bool
    workspace_name = string
    repository     = optional(string)
  })

  default = {
    enabled        = false
    workspace_name = ""
  }
}

variable "gitlab" {
  type = object({
    enabled   = bool
    group_url = string
  })

  default = {
    enabled   = false
    group_url = ""
  }
}

variable "bitbucket" {
  type = object({
    enabled         = bool
    workspace_name  = string
    workspace_uuid  = string
    repository_uuid = optional(string)
  })

  default = {
    enabled        = false
    workspace_name = ""
    workspace_uuid = ""
  }

  validation {
    condition = (
      var.bitbucket.enabled ? (
        length(var.bitbucket.workspace_name) > 0 &&
        length(var.bitbucket.workspace_uuid) > 0
      ) : true
    )
    error_message = "Workspace name and uuid is required. These can be found from OpenId Connect under the pipeline settings."
  }

  validation {
    condition = (
      var.bitbucket.repository_uuid != null ?
      substr(var.bitbucket.repository_uuid, 0, 1) == "{" &&
      substr(var.bitbucket.repository_uuid, -1, 1) == "}"
      : true
    )

    error_message = "The repository uuid is invalid, it must start with '{' and end with '}'."
  }
}

variable "identity_provider_permissions" {
  description = "A list of the actions allowed by the identity providers assumed role."

  type = list(any)
}
