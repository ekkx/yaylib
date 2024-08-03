"""
MIT License

Copyright (c) 2023 ekkx

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
"""

from datetime import datetime

# from .. import client
from .. import config
from ..errors import ForbiddenError, NotFoundError
from ..models import Post, SharedUrl
from ..responses import (
    BookmarkPostResponse,
    CreatePostResponse,
    PostResponse,
    PostsResponse,
    PostLikersResponse,
    PostTagsResponse,
    LikePostsResponse,
    VoteSurveyResponse,
)
from ..utils import build_message_tags, get_post_type, md5


class PostApi(object):
    def __init__(self, client) -> None:
        self.__client = client

    async def get_timeline(self, **params) -> PostsResponse:
        endpoint = f"/v2/posts/timeline"
        if "noreply_mode" in params and params["noreply_mode"] is True:
            endpoint = f"/v2/posts/noreply_timeline"
        return await self.__client.request(
            "GET",
            config.API_HOST + endpoint,
            params=params,
            return_type=PostsResponse,
        )
