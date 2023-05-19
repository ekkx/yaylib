import os
import json
import time
import requests

from ..config import Endpoints as ep
from ..utils import console_print
from ..exceptions import ForbiddenError


def follow_user(self, user_id):
    resp = self._post(f'{ep.USER_v2}/{user_id}/follow')
    return resp


def follow_users(self, user_ids):
    # self.logger.info(f'Going to follow {len(user_ids)} users.')
    return


def unfollow_user(self, user_id):
    resp = self._post(f'{ep.USER_v2}/{user_id}/unfollow')
    return resp


def unfollow_non_followers(self):
    # TODO: get user's followings and followers to get non followers
    # TODO: unfollow each users you got in the previous step
    # self.logger.info(f'Going to unfollow {user_count} users.')
    return


def accept_follow_request(self, user_id):
    resp = self._post(
        f'{ep.USER_v2}/{user_id}/follow_request?action=accept')
    return resp


def reject_follow_request(self, user_id):
    resp = self._post(
        f'{ep.USER_v2}/{user_id}/follow_request?action=reject')
    return resp


def send_letter(self, user_id, message):
    data = {'comment': message}
    resp = self._post(
        f'{ep.USER_v1}/reviews/{user_id}', data)
    return resp


def block_user(self, user_id):
    resp = self._post(
        f'{ep.USER_v1}/{user_id}/block')
    return resp


def unblock_user(self, user_id):
    resp = self._post(
        f'{ep.USER_v1}/{user_id}/unblock')
    return resp


def create_account(
    self,
    username,
    email,
    password,
    birth_date,
    biography,
    prefecture,
    gender,
    country_code,
    photo,
    base_path,
    save_login_info,
):
    console_print(
        '[!] アカウントの生成は、IPをBanされる可能性が高いです。プロキシを設定することをおすすめします。', 'yellow')

    try:
        filename = 'login_info.json'
        data = {
            'device_uuid': self.UUID,
            'locale': 'ja',
            'email': email
        }
        resp = requests.post(
            f'{ep.API_URL}/v1/email_verification_urls', data
        )

        if (resp.status_code != 201):
            console_print(resp.json()['message'], 'red')
            return {'result': 'error'}

        data = {'locale': 'ja', 'email': email}
        resp = requests.post(resp.json()['url'], data)

        if resp.status_code != 200:
            console_print(resp.json()['message'], 'red')
            return {'result': 'error'}

        console_print('メールアドレスにコードを送信しました。', 'green')
        code = input('コードを入力してください >> ')

        data = {'code': code, 'email': email}
        resp = requests.post(ep.GET_EMAIL_GRANT_TOKEN, data)

        if resp.status_code != 200:
            console_print('予期しないエラーが発生しました。', 'red')
            return {'result': 'error'}

        email_grant_token = resp.json()['email_grant_token']

        if photo and 'https://yay-cdn.com/file/yay-space/uploads/user_avatar/' not in photo:
            data = self.upload_image('user_avatar', photo)
            photo = data['filename']
        else:
            avatar_url_prefix = 'https://yay-cdn.com/file/yay-space/uploads/'
            photo = photo.replace(avatar_url_prefix, '')

        data = {
            "gender": gender,
            "timestamp": int(time.time()),
            "nickname": username,
            "prefecture": prefecture,
            "email_grant_token": email_grant_token,
            "country_code": country_code,
            "password": password,
            "birth_date": birth_date,
            "uuid": self.UUID,
            "api_key": self.api_key,
            "biography": biography,
            "profile_icon_filename": photo,
            "email": email
        }

        resp = requests.post(f'{ep.USER_v3}/register', data)

        if resp.status_code == 201:
            console_print('アカウントを生成しました。', 'green')

            if save_login_info is True:
                info = {'users': []}
                if not os.path.exists(base_path):
                    os.makedirs(base_path)

                fname = os.path.join(base_path, filename)

                if os.path.exists(fname):
                    with open(fname, 'r', encoding='utf-8') as f:
                        file_contents = f.read()
                        if file_contents.strip() == "":
                            with open(fname, 'w', encoding='utf-8') as f:
                                json.dump(info, f, ensure_ascii=False)
                        else:
                            info = json.loads(file_contents)
                else:
                    with open(fname, 'w', encoding='utf-8') as f:
                        json.dump(info, f, ensure_ascii=False)

                new_user = {
                    'id': resp.json()['id'],
                    'email': email,
                    'password': password
                }
                info['users'].append(new_user)

                with open(fname, 'w', encoding='utf-8') as f:
                    json.dump(info, f, ensure_ascii=False)

                print(f'(保存先: {fname})\n')

            print(f"ユーザーID: {resp.json()['id']}")
            print(f"ユーザー名: {username}")
            print(f"メールアドレス: {email}")
            print(f"パスワード: {password}")
            print(f"URL: https://yay.space/user/{resp.json()['id']}")

            return resp.json()

    except ForbiddenError:
        console_print('IP Banned', 'red')
        return {'result': 'error'}


def edit_profile(
        self,
        data_dict,
        profile_image=None,
        cover_image=None,
):
    # TODO: invalidSignedInfo
    if not data_dict:
        raise ValueError("data_dict must not be empty")
    data = {}
    if 'username' in data_dict:
        data['nickname'] = data_dict['username']
    if 'bio' in data_dict:
        data['biography'] = data_dict['bio']
    if 'country_code' in data_dict:
        data['country_code'] = data_dict['country_code']
    if profile_image:
        image_data = self.upload_image('user_avatar', profile_image)
        if 'filename' in image_data:
            data['profile_icon_filename'] = image_data['filename']
        else:
            raise ValueError("Failed to upload profile image")
    if cover_image:
        image_data = self.upload_image('user_avatar', cover_image)
        if 'filename' in image_data:
            data['cover_image_filename'] = image_data['filename']
        else:
            raise ValueError("Failed to upload cover image")
    try:
        resp = self._post(f'{ep.USER_v3}/edit', data)
        return resp
    except Exception as e:
        raise ValueError(f"Failed to edit profile: {str(e)}")
