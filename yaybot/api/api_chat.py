from ..config import Endpoints as ep
from ..utils import handle_response, console_print


def send_message(self, message, user_id=None, chat_room_id=None):
    if user_id:
        chat_room_id = self.get_chat_room_id_from_user(user_id)
    data = {'message_type': 'text', 'text': message}
    resp = self._post(
        f'{ep.CHATROOM_v1}/{chat_room_id}/messages/new', data)
    return resp


def accept_chat_request(self, chat_room_id):
    data = {'chat_room_ids[]': chat_room_id}
    resp = self._post(
        f'{ep.CHATROOM_v1}/accept_chat_request', data)
    return resp


def delete_chat_room(self, chat_room_id):
    data = {'chat_room_ids[]': chat_room_id}
    resp = self._post(
        f'{ep.CHATROOM_v1}/mass_destroy', data)
    return resp
