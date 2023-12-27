# フォロー整理用のサンプルコード
# フォロワー以外のアカウントをフォローから外します


email = "your_email"
password = "your_password"


import yaylib


def get_all_followings(client):
    data = client.get_user_followings(client.user_id)
    followings = []

    while data.last_follow_id is not None:
        followings.extend(data.users)
        data = client.get_user_followings(
            client.user_id, from_follow_id=data.last_follow_id
        )

    return followings


if __name__ == "__main__":
    client = yaylib.Client()

    client.login(email, password)

    followings = get_all_followings(client)

    for following in followings:
        if not following.is_followed_by:
            client.unfollow_user(following.id)
