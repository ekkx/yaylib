import yaylib

api = yaylib.Client()

api.login(email="", password="")

api.create_post("Hello with yaylib!")
