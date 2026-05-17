# The pagination cursor is a string in the schema, but the server
# sends it as a bare JSON number on some endpoints and a JSON string
# on others. from_dict must accept both and yield the string form.

from yaylib.models.posts_response import PostsResponse


def test_cursor_tolerates_number_and_string():
    assert (
        PostsResponse.from_dict(
            {"next_page_value": 585851614, "posts": []}
        ).next_page_value
        == "585851614"
    )
    assert (
        PostsResponse.from_dict(
            {"next_page_value": "abc123", "posts": []}
        ).next_page_value
        == "abc123"
    )
    assert (
        PostsResponse.from_dict(
            {"next_page_value": None, "posts": []}
        ).next_page_value
        is None
    )
