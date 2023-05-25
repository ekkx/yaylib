import yaylib

api = yaylib.Client()

api.login_with_email(email="abcde@exmaple.com", password="pass?word123")

print(api.api_key)
