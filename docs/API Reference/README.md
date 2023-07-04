<!-- 参考: https://developer.twitter.com/en/docs/api-reference-index -->

# API reference index

<p align="center">
    <img src="https://github.com/qvco/yaylib/assets/77382767/5265b956-55b7-466c-8cdb-cf0f3abed946" alt="Logo" height="300px">
    <h3 align="center">yaylib</h3>
    <p align="center">
        <a href="https://github.com/qvco/yaylib/issues">Report Bug</a>
        ·
        <a href="https://github.com/qvco/yaylib/issues">Request Feature</a>
        ·
        <a href="https://discord.gg/MEuBfNtqRN">Join the discord</a>
    </p>
</p>

## Posts

### Bookmarks

- <a href="">DELETE /v1/users/:user_id/bookmarks/:post_id</a>
- <a href="">GET /v1/users/:user_id/bookmarks</a>
- <a href="">PUT /v1/users/:user_id/bookmarks/:post_id</a>

### Likes

- <a href="">GET /v1/posts/:post_id/likers</a>
- <a href="">POST /v2/posts/like</a>
- <a href="">POST /v1/posts/:post_id/unlike</a>

### Timelines

- <a href="">GET /v2/posts/timeline</a>
- <a href="">GET /v2/posts/noreply_timeline</a>
- <a href="">GET /v2/posts/call_timeline</a>
- <a href="">GET /v2/posts/call_followers_timeline</a>
- <a href="">GET /v2/posts/get_following_timeline</a>
- <a href="">GET /v2/posts/group_timeline</a>
- <a href="">GET /v2/posts/recommended_timeline</a>
- <a href="">GET /v2/posts/user_timeline</a>

### Search posts

- <a href="">GET /v2/posts/tags/:hashtag</a>
- <a href="">GET /v2/posts/search</a>
- <a href="">GET /v2/groups/:group_id/posts/search</a>

## Threads

- <a href="">PUT /v3/posts/:id/move_to_thread/:thread_id</a>
<!-- addPostToThread -->
- <a href="">POST /v3/posts/:id/move_to_thread</a>
<!-- convertPostToThread -->
- <a href="">POST /v1/threads</a>
<!-- createThread -->
- <a href="">GET /v1/threads</a>
<!-- getGroupThreadList -->
- <a href="">GET /v1/threads/joined_statuses</a>
<!-- getJoinedStatuses -->
- <a href="">GET /v1/threads/:id</a>
<!-- getThread -->
- <a href="">POST /v1/threads/:id/posts</a>
<!-- getThreadPosts -->
- <a href="">POST /v1/threads/:thread_id/members/:id</a>
<!-- leaveThread -->
- <a href="">DELETE /v1/threads/:thread_id/members/:id</a>
<!-- joinThread -->
- <a href="">DELETE /v1/threads/:id</a>
<!-- removeThread -->
- <a href="">PUT /v1/threads/:id</a>
<!-- updateThread -->

## Users

### Blocks

- <a href="">GET /v1/users/block_ids</a>
- <a href="">POST /v1/users/:user_id/block</a>
- <a href="">POST /v2/users/:user_id/unblock</a>
- <a href="">PUT /v2/users/blocked</a>

### Follows

- <a href="">POST /v2/users/:user_id/follow</a>
- <a href="">POST /v2/users/follow</a>
- <a href="">GET /v1/users/active_followings</a>
- <a href="">GET /v1/friends</a>
- <a href="">GET /v2/users/follow_requests</a>
- <a href="">GET /v2/users/follow_requests_count</a>
- <a href="">GET /v1/users/following_born_today</a>
- <a href="">GET /v2/users/:user_id/followers</a>
- <a href="">GET /v2/users/:user_id/list_followings</a>
- <a href="">POST /v2/users/:target_id/follow_request</a>
- <a href="">POST /v2/users/:user_id/unfollow</a>

### Mutes

- <a href="">GET /v1/hidden/users</a>
- <a href="">POST /v1/hidden/users</a>
- <a href="">DELETE /v1/hidden/users</a>

## Groups

- <a href="">PUT /v1/groups/:id/deputize</a>
<!-- acceptModeratorOffer -->
- <a href="">PUT /v1/groups/:id/transfer</a>
<!-- acceptOwnershipOffer -->
- <a href="">POST /v1/groups/:id/accept/:userId</a>
<!-- acceptUserRequest -->
- <a href="">PUT /v1/groups/:id/related</a>
<!-- addRelatedGroups -->
- <a href="">POST /v1/groups/:id/ban/:userId</a>
<!-- addRelatedGroups -->
- <a href="">GET /v1/groups/unread_status</a>
<!-- checkUnreadStatus -->
- <a href="">POST /v3/groups/new</a>
<!-- create -->
- <a href="">POST /v1/pinned/groups</a>
<!-- createPinGroup -->

## Chats

## Reviews

## Login

## Notifications

## Misc
