class YayError(Exception):
    pass


class AuthenticationError(YayError):
    pass


class RateLimitError(YayError):
    pass
