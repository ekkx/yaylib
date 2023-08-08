import yaylib

api = yaylib.Client()

# フォロワーを取得する対象のユーザーID
target_id = 0

data = api.get_user_followers(target_id)
followers = []

while data.last_follow_id is not None:
    followers.extend(data.users)
    data = api.get_user_followers(target_id, from_follow_id=data.last_follow_id)

print(len(followers))
