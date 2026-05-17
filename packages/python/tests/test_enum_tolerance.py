# Enum-typed wire fields are declared with a fixed set of values, but
# the server may start sending a value the SDK does not know yet.
# from_dict must accept the unknown string verbatim instead of
# raising, and the typed constants must stay usable for comparisons.

from yaylib.models.post import Post
from yaylib.models.post_type import PostType


def test_unknown_enum_value_is_kept_as_string():
    post = Post.from_dict({"post_type": "__mock_unknown_enum__"})
    assert post.post_type == "__mock_unknown_enum__"


def test_known_enum_value_decodes_to_typed_constant():
    post = Post.from_dict({"post_type": "image"})
    assert post.post_type == PostType.IMAGE
    # PostType is a str-backed Enum, so the typed constant compares
    # equal to its wire string in both directions.
    assert post.post_type == "image"
    assert PostType.IMAGE == "image"
    assert "image" == PostType.IMAGE


def test_known_value_recognized_unknown_value_not():
    # There is no dedicated IsValid()-style predicate on the enum in
    # this port; the membership of a value among the enumerated
    # constants is the equivalent introspection.
    known = {member.value for member in PostType}
    assert "image" in known
    assert "__mock_unknown_enum__" not in known
