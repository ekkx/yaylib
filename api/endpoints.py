class Endpoints:
    BASE_URL = 'https://api.yay.space'
    CAS_BASE_URL = 'https://cas.yay.space'

    USER_v1 = BASE_URL + '/v1/users'
    POST_v1 = BASE_URL + '/v1/posts'
    GROUP_v1 = BASE_URL + '/v1/groups'
    CHATROOM_v1 = BASE_URL + '/v1/chat_rooms'
    PIN_v1 = BASE_URL + '/v1/pinned'

    USER_v2 = BASE_URL + '/v2/users'
    POST_v2 = BASE_URL + '/v2/posts'
    CONVERSATION_v2 = BASE_URL + '/v2/conversations'
    GROUP_v2 = BASE_URL + '/v2/groups'
    CHATROOM_v2 = BASE_URL + '/v2/chat_rooms'

    USER_v3 = BASE_URL + '/v3/users'
    POST_v3 = BASE_URL + '/v3/posts'

    GET_TIMELINE = POST_v2 + '/timeline'
    GET_USER_TIMELINE = POST_v2 + '/user_timeline'
    GET_FOLLOWING_TIMELINE = POST_v2 + '/following_timeline'
    GET_TIMELINE_BY_KEYWORD = POST_v2 + '/search'
    GET_TIMELINE_BY_HASHTAG = POST_v2 + '/tags'
