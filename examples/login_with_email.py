import yaylib

api = yaylib.Client()

api.login_with_email(email="", password="")

api.create_post("Hello with yaylib!")
