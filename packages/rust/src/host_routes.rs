// Code generated; DO NOT EDIT.

use std::collections::HashMap;
use std::sync::LazyLock;

/// Maps an upper-cased "METHOD path" to a symbolic auxiliary host.
/// The transport resolves the symbol to a configurable base URL;
/// an unlisted request uses the primary host.
pub static HOST_ROUTES: LazyLock<HashMap<&'static str, &'static str>> = LazyLock::new(|| {
    let mut m = HashMap::new();
    m.insert("GET /api/user_activities", "cassandra");
    m.insert("GET /api/v2/user_activities", "cassandra");
    m
});
