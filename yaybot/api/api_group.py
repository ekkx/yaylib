from ..config import Endpoints as ep
from ..utils import handle_response, console_print, ObjectGenerator as gen


def create_group(
    self,
    group_name: str,
    description: str = None,
    guidelines: str = None,
    group_category_id=21,
    sub_category_id: int = None,
    is_private=False,
    call_timeline_display=True,
    hide_reported_posts=False,
    allow_ownership_transfer=True,
    allow_thread_creation_by='member',
    allow_members_to_post_image_and_video=True,
    allow_members_to_post_url=True,
    hide_conference_call=False,
    secret=False,
    only_verified_age=False,
    only_mobile_verified=False,
    gender=-1,
    generation_groups_limit=0
):
    data = {
        'topic': group_name,
        'description': description,
        'guidelines': guidelines,
        'group_category_id': group_category_id,
        'sub_category_id': sub_category_id,
        'is_private': is_private,
        'call_timeline_display': call_timeline_display,
        'hide_reported_posts': hide_reported_posts,
        'allow_ownership_transfer': allow_ownership_transfer,
        'allow_thread_creation_by': allow_thread_creation_by,
        'allow_members_to_post_image_and_video': allow_members_to_post_image_and_video,
        'allow_members_to_post_url': allow_members_to_post_url,
        'hide_conference_call': hide_conference_call,
        'secret': secret,
        'only_verified_age': only_verified_age,
        'only_mobile_verified': only_mobile_verified,
        'gender': gender,
        'generation_groups_limit': generation_groups_limit
    }
    resp = self._post(
        f'{ep.GROUP_v1}/new', data)
    return resp


def delete_group(self, group_id):
    data = {'groupId': group_id}
    resp = self._delete(
        f'{ep.GROUP_v1}/{group_id}/leave', data)
    return resp


def change_group_settings(
    self,
    group_id,
    group_name: str = None,
    description: str = None,
    guidelines: str = None,
    group_category_id=21,
    sub_category_id: int = None,
    is_private=None,
    call_timeline_display=True,
    hide_reported_posts=False,
    allow_ownership_transfer=True,
    allow_thread_creation_by='member',
    allow_members_to_post_image_and_video=True,
    allow_members_to_post_url=True,
    hide_conference_call=False,
    hide_from_game_eight=False,
    secret=False,
    only_verified_age=False,
    only_mobile_verified=False,
    gender=-1,
    generation_groups_limit=0
):
    data = {
        'topic': group_name,
        'description': description,
        'guidelines': guidelines,
        'group_category_id': group_category_id,
        'sub_category_id': sub_category_id,
        'is_private': is_private,
        'call_timeline_display': call_timeline_display,
        'hide_reported_posts': hide_reported_posts,
        'allow_ownership_transfer': allow_ownership_transfer,
        'allow_thread_creation_by': allow_thread_creation_by,
        'allow_members_to_post_image_and_video': allow_members_to_post_image_and_video,
        'allow_members_to_post_url': allow_members_to_post_url,
        'hide_conference_call': hide_conference_call,
        'hide_from_game_eight': hide_from_game_eight,
        'secret': secret,
        'only_verified_age': only_verified_age,
        'only_mobile_verified': only_mobile_verified,
        'gender': gender,
        'generation_groups_limit': generation_groups_limit,
        'uuid': ''
    }
    resp = self._put(
        f'https://yay.space/api/groups/{group_id}', data)
    return gen.group_object(self, resp.get('group'))


def transfer_group_ownership(self, group_id, user_id):
    # need a fix. json.decoder.JSONDecodeError: Expecting value: line 1 column 1 (char 0)
    # the resp: <Response [400]>, probably the uuid is invalid.
    data = {
        'uuid': '',
        'user_id': user_id
    }
    resp = self._post(
        f'https://yay.space/api/groups/{group_id}/transfer', data)
    return resp


def offer_group_sub_owner(self, group_id, user_id):
    # need a fix. json.decoder.JSONDecodeError: Expecting value: line 1 column 1 (char 0)
    # the resp: <Response [400]>, probably the uuid is invalid.
    data = {'uuid': '', 'user_ids[]': user_id}
    resp = self._post(
        f'https://yay.space/api/groups/{group_id}/deputize', data)
    return resp


def undo_group_ownership_transfer(self, group_id, user_id):
    data = {'user_id': user_id}
    resp = self._put(
        f'{ep.GROUP_v1}/{group_id}/transfer/withdraw', data)
    return resp


def undo_group_sub_owner_offer(self, group_id, user_id):
    resp = self._put(
        f'{ep.GROUP_v1}/{group_id}/deputize/{user_id}/withdraw')
    return resp


def fire_group_sub_owner(self, group_id, user_id):
    resp = self._post(
        f'{ep.GROUP_v1}/{group_id}/fire/{user_id}')
    return resp


def accept_group_join_request(self, group_id, user_id):
    data = {
        'groupId': group_id,
        'mode': 'accept',
        'userId': user_id
    }
    resp = self._post(
        f'{ep.GROUP_v1}/{group_id}/accept/{user_id}', data)
    return resp


def decline_group_join_request(self, group_id, user_id):
    data = {
        'groupId': group_id,
        'mode': 'decline',
        'userId': user_id
    }
    resp = self._post(
        f'{ep.GROUP_v1}/{group_id}/decline/{user_id}', data)
    return resp


def invite_user_to_group(self, group_id, user_id):
    data = {'user_ids[]': user_id}
    resp = self._post(
        f'{ep.GROUP_v1}/{group_id}/invite', data)
    return resp


def pin_group_post(self, group_id, post_id):
    data = {'group_id': group_id, 'post_id': post_id}
    resp = self._put(
        f'{ep.POST_v2}/group_pinned_post', data)
    return resp


def unpin_group_post(self, group_id):
    data = {'group_id': group_id}
    resp = self._delete(
        f'{ep.POST_v2}/group_pinned_post', data)
    return resp


def ban_user_from_group(self, group_id, user_id):
    resp = self._post(
        f'{ep.GROUP_v1}/{group_id}/ban/{user_id}')
    return resp


def unban_user_from_group(self, group_id, user_id):
    resp = self._post(
        f'{ep.GROUP_v1}/{group_id}/unban/{user_id}')
    return resp


def join_group(self, group_id):
    self._post(f'{ep.GROUP_v1}/{group_id}/join')
    return self.get_group(group_id)


def leave_group(self, group_id):
    data = {'groupId': group_id}
    resp = self._delete(
        f'{ep.GROUP_v1}/{group_id}/leave', data)
    return resp
