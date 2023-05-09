import os
import json
import time
import uuid
import requests

from ..config import Endpoints as ep
from ..utils import console_print


def follow_user(self, user_id):
    resp = self._post(f'{ep.USER_v2}/{user_id}/follow')
    return resp


def unfollow_user(self, user_id):
    resp = self._post(f'{ep.USER_v2}/{user_id}/unfollow')
    return resp


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
    UUID = str(uuid.uuid4())

    data = {
        'device_uuid': UUID,
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

    if photo:
        data = self.upload_photo(self, photo)
        photo = data['filename']

    data = {
        "gender": gender,
        "timestamp": int(time.time()),
        "nickname": username,
        "prefecture": prefecture,
        "email_grant_token": email_grant_token,
        "country_code": country_code,
        "password": password,
        "birth_date": birth_date,
        "uuid": UUID,
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

            if os.path.exists(base_path + '/login_info.json'):
                with open(base_path + '/login_info.json', 'r', encoding='utf-8') as f:
                    file_contents = f.read()
                    if file_contents.strip() == "":
                        with open(base_path + '/login_info.json', 'w', encoding='utf-8') as f:
                            json.dump(info, f, ensure_ascii=False)
                    else:
                        info = json.loads(file_contents)
            else:
                with open(base_path + '/login_info.json', 'w', encoding='utf-8') as f:
                    json.dump(info, f, ensure_ascii=False)

            new_user = {
                'id': resp.json()['id'],
                'username': username,
                'email': email,
                'password': password
            }
            info['users'].append(new_user)

            with open(base_path + 'login_info.json', 'w', encoding='utf-8') as f:
                json.dump(info, f, ensure_ascii=False)
            print('(保存先: /config/login_info.json)\n')

        print(f"ユーザーID: {resp.json()['id']}")
        print(f"ユーザー名: {username}")
        print(f"メールアドレス: {email}")
        print(f"パスワード: {password}")
    return resp.json()
