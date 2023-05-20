class User:

    def __init__(self,
                 user_id: int,
                 biography: str,
                 birth_date: str,
                 cover_image: str,
                 gender: str,
                 cover_image_thumbnail: str,
                 followers_count: int,
                 followings_count: int,
                 groups_users_count: int,
                 is_age_verified: bool,
                 is_chat_request_on: bool,
                 is_dangerous: bool,
                 is_email_confirmed: bool,
                 is_facebook_connected: bool,
                 is_follow_pending: bool,
                 is_followed_by: bool,
                 is_following: bool,
                 is_from_different_generation_and_trusted: bool,
                 is_group_phone_on: bool,
                 is_hidden: bool,
                 is_hide_vip: bool,
                 is_interests_selected: bool,
                 is_line_connected: bool,
                 is_lobi_connected: bool,
                 is_new: bool,
                 is_phone_on: bool,
                 is_private: bool,
                 is_recently_penalized: bool,
                 is_registered: bool,
                 is_video_pn: bool,
                 is_vip: bool,
                 masked_email: str,
                 mobile_number: str,
                 nickname: str,
                 posts_count: int,
                 prefecture: str,
                 profile_icon: str,
                 profile_icon_thumbnail: str,
                 reviews_count: int,
                 ):
        self.id = user_id
        self.biography = biography
        self.birth_date = birth_date
        self.cover_image = cover_image
        self.gender = gender
        self.cover_image_thumbnail = cover_image_thumbnail
        self.followers_count = followers_count
        self.followings_count = followings_count
        self.groups_users_count = groups_users_count
        self.is_age_verified = is_age_verified
        self.is_chat_request_on = is_chat_request_on
        self.is_dangerous = is_dangerous
        self.is_email_confirmed = is_email_confirmed
        self.is_facebook_connected = is_facebook_connected
        self.is_follow_pending = is_follow_pending
        self.is_followed_by = is_followed_by
        self.is_following = is_following
        self.is_from_different_generation_and_trusted = is_from_different_generation_and_trusted
        self.is_group_phone_on = is_group_phone_on
        self.is_hidden = is_hidden
        self.is_hide_vip = is_hide_vip
        self.is_interests_selected = is_interests_selected
        self.is_line_connected = is_line_connected
        self.is_lobi_connected = is_lobi_connected
        self.is_new = is_new
        self.is_phone_on = is_phone_on
        self.is_private = is_private
        self.is_recently_penalized = is_recently_penalized
        self.is_registered = is_registered
        self.is_video_pn = is_video_pn
        self.is_vip = is_vip
        self.masked_email = masked_email
        self.mobile_number = mobile_number
        self.nickname = nickname
        self.posts_count = posts_count
        self.prefecture = prefecture
        self.profile_icon = profile_icon
        self.profile_icon_thumbnail = profile_icon_thumbnail
        self.reviews_count = reviews_count

    def __repr__(self):
        return f'User(username={self.nickname})'
