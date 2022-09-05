# Users

This page contains configuration for users and their roles in the Core Banking
system.

The user may create an API key with a scope that allows them to access the API
programmatically without OAuth2. This feature may be "disabled" by not granting
any users the permission to create API keys.

Each user may create multiple API consumers which may have multiple API keys.
For example, the user may create one consumer per project and assign a
different API key to each program inside the project, granting different scopes
of permissions.

An API key is not necessary to access the API but is recommended to limit the
capabilities of automated programs and reduce the risk through principle of
least privilege.

See the Authentication section in the Home page for more details on accessing
the API.
