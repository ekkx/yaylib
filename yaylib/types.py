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


class PostType(object):
    """投稿の種類"""

    text = "text"
    """ `text`: テキストのみ投稿タイプ"""
    media = "media"
    """ `media`: メディアを含める投稿タイプ"""
    image = "image"
    """ `image`: 画像を含める投稿タイプ"""
    video = "video"
    """ `video`: ビデオを含める投稿タイプ"""
    survey = "survey"
    """ `survey`: アンケートを含める投稿タイプ"""
    call = "call"
    """ `call`: 通話用の投稿タイプ"""
    shareable_url = "shareable_url"
    """ `shareable_url`: サークルやスレッド共有用の投稿タイプ"""


class CallType(object):
    """通話の種類"""

    voice = "voice"
    """ `voice`: 音声通話用の通話タイプ"""
    video = "vdo"
    """ `vdo`: ビデオ通話用の通話タイプ"""


class ImageType(object):
    """画像の種類"""

    post = "post"
    """ `post`: 投稿に画像をアップロードする際の画像タイプ"""
    chat_message = "chat_message"
    """ `chat_message`: 個人チャットに画像をアップロードする際の画像タイプ"""
    chat_background = "chat_background"
    """ `chat_background`: 個人チャットの背景用に画像をアップロードする際の画像タイプ"""
    report = "report"
    """ `report`: 通報用の画像をアップロードする際の画像タイプ"""
    user_avatar = "user_avatar"
    """ `user_avatar`: プロフィール画像をアップロードする際の画像タイプ"""
    user_cover = "user_cover"
    """ `user_cover`: プロフィールの背景画像をアップロードする際の画像タイプ"""
    group_cover = "group_cover"
    """ `group_cover`: グループの背景画像をアップロードする際の画像タイプ"""
    group_thread_icon = "group_thread_icon"
    """ `group_thread_icon`: グループ内のスレッド用アイコンをアップロードする際の画像タイプ"""
    group_icon = "group_icon"
    """ `group_icon`: グループのアイコンをアップロードする際の画像タイプ"""


class ShareableType(object):
    """共有の種類"""

    group = "group"
    """ `group`: サークル用の共有タイプ"""
    thread = "thread"
    """ `thread`: スレッド用の共有タイプ"""


class PolicyType(object):
    """利用規約の種類"""

    privacy_policy = "privacy_policy"
    """ `privacy_policy`: プライバシーポリシー"""
    terms_of_use = "terms_of_use"
    """ `privacy_policy`: 利用規約"""
