---
page_title: "zitadel_org_idp_jwt Resource - terraform-provider-zitadel"
subcategory: ""
description: |-
  Resource representing a generic JWT IdP of the organization.
---

# zitadel_org_idp_jwt (Resource)

Resource representing a generic JWT IdP of the organization.

## Example Usage

```terraform
resource "zitadel_org_idp_jwt" "default" {
  org_id        = data.zitadel_org.default.id
  name          = "jwtidp"
  styling_type  = "STYLING_TYPE_UNSPECIFIED"
  jwt_endpoint  = "https://jwtendpoint.com/jwt"
  issuer        = "https://google.com"
  keys_endpoint = "https://jwtendpoint.com/keys"
  header_name   = "x-auth-token"
  auto_register = false
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `auto_register` (Boolean) auto register for users from this idp
- `header_name` (String) the name of the header where the JWT is sent in, default is authorization
- `issuer` (String) the issuer of the jwt (for validation)
- `jwt_endpoint` (String) the endpoint where the jwt can be extracted
- `keys_endpoint` (String) the endpoint to the key (JWK) which are used to sign the JWT with
- `name` (String) Name of the IDP
- `styling_type` (String) Some identity providers specify the styling of the button to their login, supported values: STYLING_TYPE_UNSPECIFIED, STYLING_TYPE_GOOGLE

### Optional

- `org_id` (String) ID of the organization

### Read-Only

- `id` (String) The ID of this resource.

## Import

```terraform
# The resource can be imported using the ID format `<id[:org_id]>`, e.g.
terraform import org_idp_jwt.imported '123456789012345678:123456789012345678'
```
