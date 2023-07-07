class YayError(Exception):
    pass


class HTTPError(YayError):
    pass


class BadRequestError(HTTPError):
    pass


class AuthenticationError(HTTPError):
    pass


class ForbiddenError(HTTPError):
    pass


class NotFoundError(HTTPError):
    pass


class RateLimitError(HTTPError):
    pass


class YayServerError(HTTPError):
    pass
