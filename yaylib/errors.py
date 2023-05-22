class YayError(Exception):
    pass


class HTTPException(YayError):
    pass


class BadRequestError(HTTPException):
    pass


class AuthenticationError(HTTPException):
    pass


class ForbiddenError(HTTPException):
    pass


class NotFoundError(HTTPException):
    pass


class RateLimitError(HTTPException):
    pass


class YayServerError(HTTPException):
    pass


class ExceedCallQuotaError(HTTPException):
    pass


class InvalidSignedInfoError(HTTPException):
    pass
