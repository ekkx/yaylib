class Endpoints:

    # base urls
    API_DOMAIN = 'api.yay.space'
    CAS_DOMAIN = 'cas.yay.space'
    API_URL = 'https://{domain}'.format(domain=API_DOMAIN)
    CAS_URL = 'https://{domain}'.format(domain=CAS_DOMAIN)

    # api v1
    USER_v1 = API_URL + '/v1/users'
    POST_v1 = API_URL + '/v1/posts'
    GROUP_v1 = API_URL + '/v1/groups'
    CHATROOM_v1 = API_URL + '/v1/chat_rooms'
    PIN_v1 = API_URL + '/v1/pinned'

    # api v2
    USER_v2 = API_URL + '/v2/users'
    POST_v2 = API_URL + '/v2/posts'
    SURVEY_v2 = API_URL + '/v2/surveys'
    CONVERSATION_v2 = API_URL + '/v2/conversations'
    GROUP_v2 = API_URL + '/v2/groups'
    CHATROOM_v2 = API_URL + '/v2/chat_rooms'

    # api v3
    USER_v3 = API_URL + '/v3/users'
    POST_v3 = API_URL + '/v3/posts'

    # api other
    GET_TIMELINE = POST_v2 + '/timeline'
    GET_USER_TIMELINE = POST_v2 + '/user_timeline'
    GET_FOLLOWING_TIMELINE = POST_v2 + '/following_timeline'
    GET_TIMELINE_BY_KEYWORD = POST_v2 + '/search'
    GET_TIMELINE_BY_HASHTAG = POST_v2 + '/tags'

    GET_EMAIL_GRANT_TOKEN = 'https://idcardcheck.com/apis/v1/apps/yay/email_grant_tokens'
