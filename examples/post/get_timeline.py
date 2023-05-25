import yaylib

api = yaylib.Client()

data = api.get_timeline(number=5)

print(data.posts)
