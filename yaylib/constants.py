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

from dataclasses import dataclass


@dataclass
class Color:
    """カラー定数"""

    HEADER = "\033[95m"
    OKBLUE = "\033[94m"
    OKCYAN = "\033[96m"
    OKGREEN = "\033[92m"
    WARNING = "\033[93m"
    FAIL = "\033[91m"
    RESET = "\033[0m"
    BOLD = "\033[1m"
    UNDERLINE = "\033[4m"


@dataclass
class PostType:
    """投稿の種類"""

    TEXT = "text"
    """ `text`: テキストのみ投稿タイプ"""
    MEDIA = "media"
    """ `media`: メディアを含める投稿タイプ"""
    IMAGE = "image"
    """ `image`: 画像を含める投稿タイプ"""
    VIDEO = "video"
    """ `video`: ビデオを含める投稿タイプ"""
    SURVEY = "survey"
    """ `survey`: アンケートを含める投稿タイプ"""
    CALL = "call"
    """ `call`: 通話用の投稿タイプ"""
    SHAREABLE_URL = "shareable_url"
    """ `shareable_url`: サークルやスレッド共有用の投稿タイプ"""


@dataclass
class CallType:
    """通話の種類"""

    VOICE = "voice"
    """ `voice`: 音声通話用の通話タイプ"""
    VIDEO = "vdo"
    """ `vdo`: ビデオ通話用の通話タイプ"""


@dataclass
class ImageType:
    """画像の種類"""

    POST = "post"
    """ `post`: 投稿に画像をアップロードする際の画像タイプ"""
    CHAT_MESSAGE = "chat_message"
    """ `chat_message`: 個人チャットに画像をアップロードする際の画像タイプ"""
    CHAT_BACKGROUND = "chat_background"
    """ `chat_background`: 個人チャットの背景用に画像をアップロードする際の画像タイプ"""
    REPORT = "report"
    """ `report`: 通報用の画像をアップロードする際の画像タイプ"""
    USER_AVATAR = "user_avatar"
    """ `user_avatar`: プロフィール画像をアップロードする際の画像タイプ"""
    USER_COVER = "user_cover"
    """ `user_cover`: プロフィールの背景画像をアップロードする際の画像タイプ"""
    GROUP_COVER = "group_cover"
    """ `group_cover`: グループの背景画像をアップロードする際の画像タイプ"""
    GROUP_THREAD_ICON = "group_thread_icon"
    """ `group_thread_icon`: グループ内のスレッド用アイコンをアップロードする際の画像タイプ"""
    GROUP_ICON = "group_icon"
    """ `group_icon`: グループのアイコンをアップロードする際の画像タイプ"""


@dataclass
class ShareableType:
    """共有の種類"""

    GROUP = "group"
    """ `group`: サークル用の共有タイプ"""
    THREAD = "thread"
    """ `thread`: スレッド用の共有タイプ"""


@dataclass
class PolicyType:
    """利用規約の種類"""

    PRIVACY_POLICY = "privacy_policy"
    """ `privacy_policy`: プライバシーポリシー"""
    TERM_OF_USE = "terms_of_use"
    """ `terms_of_use`: 利用規約"""
