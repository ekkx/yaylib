from ..config import Endpoints as ep


def follow_user(self, user_id):
    resp = self._post(f'{ep.USER_v2}/{user_id}/follow')
    return resp


def unfollow_user(self, user_id):
    resp = self._post(f'{ep.USER_v2}/{user_id}/unfollow')
    return resp


def accept_follow_request(self, user_id):
    resp = self._post(
        f'{ep.USER_v2}/{user_id}/follow_request?action=accept')
    return resp


def reject_follow_request(self, user_id):
    resp = self._post(
        f'{ep.USER_v2}/{user_id}/follow_request?action=reject')
    return resp


def send_letter(self, user_id, message):
    data = {'comment': message}
    resp = self._post(
        f'{ep.USER_v1}/reviews/{user_id}', data)
    return resp


def block_user(self, user_id):
    resp = self._post(
        f'{ep.USER_v1}/{user_id}/block')
    return resp


def unblock_user(self, user_id):
    resp = self._post(
        f'{ep.USER_v1}/{user_id}/unblock')
    return resp
