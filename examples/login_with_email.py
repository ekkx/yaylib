import yaylib


api = yaylib.Client()

api.login(email="your_email", password="your_password")

api.create_post("Hello with yaylib!")
