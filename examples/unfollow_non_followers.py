# フォロー整理用のサンプルコード
# フォロワー以外のアカウントをフォローから外します


email = "your_email"
password = "your_password"


import time
import yaylib
from yaylib.errors import BadRequestError


def get_all_followings(api):
    data = api.get_user_followings(api.user_id)
    followings = []

    while data.last_follow_id is not None:
        followings.extend(data.users)
        data = api.get_user_followings(api.user_id, from_follow_id=data.last_follow_id)

    return followings


def unfollow_user(api, user_id, max_retries=7):
    # フォロー解除が成功するまで同じユーザーに対して繰り返しリトライ
    retries = 0
    success = False
    while not success and retries < max_retries:
        try:
            api.unfollow_user(user_id)
            success = True
        except BadRequestError:
            print("この機能の上限回数に達しました。10分間待機しています...")
            time.sleep(60 * 10)
            retries += 1

    if not success:
        print(f"フォロー解除ができませんでした。ユーザーID: {user_id}")


if __name__ == "__main__":
    api = yaylib.Client()

    api.login(email, password)

    followings = get_all_followings(api)

    for following in followings:
        if not following.is_followed_by:
            unfollow_user(api, following.id)
