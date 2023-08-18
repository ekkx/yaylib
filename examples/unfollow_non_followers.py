# フォロー整理用のサンプルコード
# フォロワー以外のアカウントをフォローから外します


email = "your_email"
password = "your_password"


import yaylib


def get_all_followings(api):
    data = api.get_user_followings(api.user_id)
    followings = []

    while data.last_follow_id is not None:
        followings.extend(data.users)
        data = api.get_user_followings(api.user_id, from_follow_id=data.last_follow_id)

    return followings


if __name__ == "__main__":
    api = yaylib.Client()

    api.login(email, password)

    followings = get_all_followings(api)

    for following in followings:
        if not following.is_followed_by:
            api.unfollow_user(following.id)
