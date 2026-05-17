# Code generated; DO NOT EDIT.

# Maps an upper-cased "METHOD path" to a symbolic auxiliary host.
# The transport resolves the symbol to a configurable base URL;
# an unlisted request uses the primary host.
HOST_ROUTES: dict[str, str] = {
    "GET /api/user_activities": "cassandra",
    "GET /api/v2/user_activities": "cassandra",
}
