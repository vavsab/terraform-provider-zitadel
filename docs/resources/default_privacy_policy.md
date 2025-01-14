---
page_title: "zitadel_default_privacy_policy Resource - terraform-provider-zitadel"
subcategory: ""
description: |-
  Resource representing the default privacy policy.
---

# zitadel_default_privacy_policy (Resource)

Resource representing the default privacy policy.

## Example Usage

```terraform
resource "zitadel_default_privacy_policy" "default" {
  tos_link      = "https://example.com/tos"
  privacy_link  = "https://example.com/privacy"
  help_link     = "https://example.com/help"
  support_email = "support@example.com"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `help_link` (String)
- `privacy_link` (String)
- `support_email` (String)
- `tos_link` (String)

### Read-Only

- `id` (String) The ID of this resource.

## Import

```terraform
# The resource can be imported using the ID format `<>`, e.g.
terraform import default_privacy_policy.imported ''
```
