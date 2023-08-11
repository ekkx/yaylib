import yaylib

api = yaylib.Client()

api.login(email="your_email", password="your_password")

new_users = api.search_users(recently_created=True)

for new_user in new_users.users:
    api.follow_user(new_user.id)
