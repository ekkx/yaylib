// Code generated from the Yay! API schema; DO NOT EDIT.

use serde::{Deserialize, Deserializer, de};

/// Deserializes an optional string field that may arrive as a JSON
/// number on some endpoints. Accepts null, string, or number and
/// coerces numbers to their string form.
pub fn deserialize_option_string_or_number<'de, D>(
    deserializer: D,
) -> Result<Option<String>, D::Error>
where
    D: Deserializer<'de>,
{
    use serde_json::Value;
    match Option::<Value>::deserialize(deserializer)? {
        None | Some(Value::Null) => Ok(None),
        Some(Value::String(s)) => Ok(Some(s)),
        Some(Value::Number(n)) => Ok(Some(n.to_string())),
        Some(v) => Err(de::Error::custom(format!(
            "cursor: cannot decode {} into string",
            v
        ))),
    }
}
