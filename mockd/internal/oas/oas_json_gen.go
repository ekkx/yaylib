// Code generated; DO NOT EDIT.

package oas

import (
	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
)

// Encode implements json.Marshaler.
func (s *ActiveFollowingsResponse) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *ActiveFollowingsResponse) encodeFields(e *jx.Encoder) {
	{
		if s.LastLoggedinAt.Set {
			e.FieldStart("last_loggedin_at")
			s.LastLoggedinAt.Encode(e)
		}
	}
	{
		if s.Users.Set {
			e.FieldStart("users")
			s.Users.Encode(e)
		}
	}
}

var jsonFieldsNameOfActiveFollowingsResponse = [2]string{
	0: "last_loggedin_at",
	1: "users",
}

// Decode decodes ActiveFollowingsResponse from json.
func (s *ActiveFollowingsResponse) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode ActiveFollowingsResponse to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "last_loggedin_at":
			if err := func() error {
				s.LastLoggedinAt.Reset()
				if err := s.LastLoggedinAt.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"last_loggedin_at\"")
			}
		case "users":
			if err := func() error {
				s.Users.Reset()
				if err := s.Users.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"users\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode ActiveFollowingsResponse")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *ActiveFollowingsResponse) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *ActiveFollowingsResponse) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *ActivitiesResponse) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *ActivitiesResponse) encodeFields(e *jx.Encoder) {
	{
		if s.Activities.Set {
			e.FieldStart("activities")
			s.Activities.Encode(e)
		}
	}
	{
		if s.LastTimestamp.Set {
			e.FieldStart("last_timestamp")
			s.LastTimestamp.Encode(e)
		}
	}
}

var jsonFieldsNameOfActivitiesResponse = [2]string{
	0: "activities",
	1: "last_timestamp",
}

// Decode decodes ActivitiesResponse from json.
func (s *ActivitiesResponse) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode ActivitiesResponse to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "activities":
			if err := func() error {
				s.Activities.Reset()
				if err := s.Activities.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"activities\"")
			}
		case "last_timestamp":
			if err := func() error {
				s.LastTimestamp.Reset()
				if err := s.LastTimestamp.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"last_timestamp\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode ActivitiesResponse")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *ActivitiesResponse) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *ActivitiesResponse) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *Activity) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *Activity) encodeFields(e *jx.Encoder) {
	{
		if s.BirthdayUsers.Set {
			e.FieldStart("birthday_users")
			s.BirthdayUsers.Encode(e)
		}
	}
	{
		if s.BirthdayUsersCount.Set {
			e.FieldStart("birthday_users_count")
			s.BirthdayUsersCount.Encode(e)
		}
	}
	{
		if s.CreatedAt.Set {
			e.FieldStart("created_at")
			s.CreatedAt.Encode(e)
		}
	}
	{
		if s.EmplAmount.Set {
			e.FieldStart("empl_amount")
			s.EmplAmount.Encode(e)
		}
	}
	{
		if s.Followers.Set {
			e.FieldStart("followers")
			s.Followers.Encode(e)
		}
	}
	{
		if s.FollowersCount.Set {
			e.FieldStart("followers_count")
			s.FollowersCount.Encode(e)
		}
	}
	{
		if s.FromPost != nil {
			e.FieldStart("from_post")
			s.FromPost.Encode(e)
		}
	}
	{
		if s.FromPostIds.Set {
			e.FieldStart("from_post_ids")
			s.FromPostIds.Encode(e)
		}
	}
	{
		if s.GiftItem.Set {
			e.FieldStart("gift_item")
			s.GiftItem.Encode(e)
		}
	}
	{
		if s.Group != nil {
			e.FieldStart("group")
			s.Group.Encode(e)
		}
	}
	{
		if s.ID.Set {
			e.FieldStart("id")
			s.ID.Encode(e)
		}
	}
	{
		if s.Metadata.Set {
			e.FieldStart("metadata")
			s.Metadata.Encode(e)
		}
	}
	{
		if s.ToPost != nil {
			e.FieldStart("to_post")
			s.ToPost.Encode(e)
		}
	}
	{
		if s.Type.Set {
			e.FieldStart("type")
			s.Type.Encode(e)
		}
	}
	{
		if s.User != nil {
			e.FieldStart("user")
			s.User.Encode(e)
		}
	}
	{
		if s.VipReward.Set {
			e.FieldStart("vip_reward")
			s.VipReward.Encode(e)
		}
	}
}

var jsonFieldsNameOfActivity = [16]string{
	0:  "birthday_users",
	1:  "birthday_users_count",
	2:  "created_at",
	3:  "empl_amount",
	4:  "followers",
	5:  "followers_count",
	6:  "from_post",
	7:  "from_post_ids",
	8:  "gift_item",
	9:  "group",
	10: "id",
	11: "metadata",
	12: "to_post",
	13: "type",
	14: "user",
	15: "vip_reward",
}

// Decode decodes Activity from json.
func (s *Activity) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Activity to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "birthday_users":
			if err := func() error {
				s.BirthdayUsers.Reset()
				if err := s.BirthdayUsers.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"birthday_users\"")
			}
		case "birthday_users_count":
			if err := func() error {
				s.BirthdayUsersCount.Reset()
				if err := s.BirthdayUsersCount.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"birthday_users_count\"")
			}
		case "created_at":
			if err := func() error {
				s.CreatedAt.Reset()
				if err := s.CreatedAt.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"created_at\"")
			}
		case "empl_amount":
			if err := func() error {
				s.EmplAmount.Reset()
				if err := s.EmplAmount.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"empl_amount\"")
			}
		case "followers":
			if err := func() error {
				s.Followers.Reset()
				if err := s.Followers.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"followers\"")
			}
		case "followers_count":
			if err := func() error {
				s.FollowersCount.Reset()
				if err := s.FollowersCount.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"followers_count\"")
			}
		case "from_post":
			if err := func() error {
				s.FromPost = nil
				var elem Post
				if err := elem.Decode(d); err != nil {
					return err
				}
				s.FromPost = &elem
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"from_post\"")
			}
		case "from_post_ids":
			if err := func() error {
				s.FromPostIds.Reset()
				if err := s.FromPostIds.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"from_post_ids\"")
			}
		case "gift_item":
			if err := func() error {
				s.GiftItem.Reset()
				if err := s.GiftItem.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"gift_item\"")
			}
		case "group":
			if err := func() error {
				s.Group = nil
				var elem Group
				if err := elem.Decode(d); err != nil {
					return err
				}
				s.Group = &elem
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"group\"")
			}
		case "id":
			if err := func() error {
				s.ID.Reset()
				if err := s.ID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"id\"")
			}
		case "metadata":
			if err := func() error {
				s.Metadata.Reset()
				if err := s.Metadata.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"metadata\"")
			}
		case "to_post":
			if err := func() error {
				s.ToPost = nil
				var elem Post
				if err := elem.Decode(d); err != nil {
					return err
				}
				s.ToPost = &elem
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"to_post\"")
			}
		case "type":
			if err := func() error {
				s.Type.Reset()
				if err := s.Type.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"type\"")
			}
		case "user":
			if err := func() error {
				s.User = nil
				var elem RealmUser
				if err := elem.Decode(d); err != nil {
					return err
				}
				s.User = &elem
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"user\"")
			}
		case "vip_reward":
			if err := func() error {
				s.VipReward.Reset()
				if err := s.VipReward.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"vip_reward\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode Activity")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *Activity) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *Activity) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *AdditionalSettingsResponse) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *AdditionalSettingsResponse) encodeFields(e *jx.Encoder) {
	{
		if s.Settings.Set {
			e.FieldStart("settings")
			s.Settings.Encode(e)
		}
	}
}

var jsonFieldsNameOfAdditionalSettingsResponse = [1]string{
	0: "settings",
}

// Decode decodes AdditionalSettingsResponse from json.
func (s *AdditionalSettingsResponse) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode AdditionalSettingsResponse to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "settings":
			if err := func() error {
				s.Settings.Reset()
				if err := s.Settings.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"settings\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode AdditionalSettingsResponse")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *AdditionalSettingsResponse) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *AdditionalSettingsResponse) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *Application) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *Application) encodeFields(e *jx.Encoder) {
	{
		if s.Settings.Set {
			e.FieldStart("settings")
			s.Settings.Encode(e)
		}
	}
}

var jsonFieldsNameOfApplication = [1]string{
	0: "settings",
}

// Decode decodes Application from json.
func (s *Application) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Application to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "settings":
			if err := func() error {
				s.Settings.Reset()
				if err := s.Settings.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"settings\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode Application")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *Application) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *Application) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *ApplicationConfigResponse) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *ApplicationConfigResponse) encodeFields(e *jx.Encoder) {
	{
		if s.App.Set {
			e.FieldStart("app")
			s.App.Encode(e)
		}
	}
}

var jsonFieldsNameOfApplicationConfigResponse = [1]string{
	0: "app",
}

// Decode decodes ApplicationConfigResponse from json.
func (s *ApplicationConfigResponse) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode ApplicationConfigResponse to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "app":
			if err := func() error {
				s.App.Reset()
				if err := s.App.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"app\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode ApplicationConfigResponse")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *ApplicationConfigResponse) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *ApplicationConfigResponse) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *BanWord) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *BanWord) encodeFields(e *jx.Encoder) {
	{
		if s.ID.Set {
			e.FieldStart("id")
			s.ID.Encode(e)
		}
	}
	{
		if s.Type.Set {
			e.FieldStart("type")
			s.Type.Encode(e)
		}
	}
	{
		if s.Word.Set {
			e.FieldStart("word")
			s.Word.Encode(e)
		}
	}
}

var jsonFieldsNameOfBanWord = [3]string{
	0: "id",
	1: "type",
	2: "word",
}

// Decode decodes BanWord from json.
func (s *BanWord) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode BanWord to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "id":
			if err := func() error {
				s.ID.Reset()
				if err := s.ID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"id\"")
			}
		case "type":
			if err := func() error {
				s.Type.Reset()
				if err := s.Type.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"type\"")
			}
		case "word":
			if err := func() error {
				s.Word.Reset()
				if err := s.Word.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"word\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode BanWord")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *BanWord) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *BanWord) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *BanWordsResponse) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *BanWordsResponse) encodeFields(e *jx.Encoder) {
	{
		if s.BanWords.Set {
			e.FieldStart("ban_words")
			s.BanWords.Encode(e)
		}
	}
}

var jsonFieldsNameOfBanWordsResponse = [1]string{
	0: "ban_words",
}

// Decode decodes BanWordsResponse from json.
func (s *BanWordsResponse) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode BanWordsResponse to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "ban_words":
			if err := func() error {
				s.BanWords.Reset()
				if err := s.BanWords.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"ban_words\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode BanWordsResponse")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *BanWordsResponse) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *BanWordsResponse) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *Bgm) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *Bgm) encodeFields(e *jx.Encoder) {
	{
		if s.ID.Set {
			e.FieldStart("id")
			s.ID.Encode(e)
		}
	}
	{
		if s.MusicURL.Set {
			e.FieldStart("music_url")
			s.MusicURL.Encode(e)
		}
	}
	{
		if s.Order.Set {
			e.FieldStart("order")
			s.Order.Encode(e)
		}
	}
	{
		if s.Title.Set {
			e.FieldStart("title")
			s.Title.Encode(e)
		}
	}
}

var jsonFieldsNameOfBgm = [4]string{
	0: "id",
	1: "music_url",
	2: "order",
	3: "title",
}

// Decode decodes Bgm from json.
func (s *Bgm) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Bgm to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "id":
			if err := func() error {
				s.ID.Reset()
				if err := s.ID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"id\"")
			}
		case "music_url":
			if err := func() error {
				s.MusicURL.Reset()
				if err := s.MusicURL.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"music_url\"")
			}
		case "order":
			if err := func() error {
				s.Order.Reset()
				if err := s.Order.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"order\"")
			}
		case "title":
			if err := func() error {
				s.Title.Reset()
				if err := s.Title.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"title\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode Bgm")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *Bgm) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *Bgm) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *BgmsResponse) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *BgmsResponse) encodeFields(e *jx.Encoder) {
	{
		if s.Bgm.Set {
			e.FieldStart("bgm")
			s.Bgm.Encode(e)
		}
	}
}

var jsonFieldsNameOfBgmsResponse = [1]string{
	0: "bgm",
}

// Decode decodes BgmsResponse from json.
func (s *BgmsResponse) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode BgmsResponse to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "bgm":
			if err := func() error {
				s.Bgm.Reset()
				if err := s.Bgm.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"bgm\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode BgmsResponse")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *BgmsResponse) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *BgmsResponse) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *BlockedUserIdsResponse) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *BlockedUserIdsResponse) encodeFields(e *jx.Encoder) {
	{
		if s.BlockIds.Set {
			e.FieldStart("block_ids")
			s.BlockIds.Encode(e)
		}
	}
}

var jsonFieldsNameOfBlockedUserIdsResponse = [1]string{
	0: "block_ids",
}

// Decode decodes BlockedUserIdsResponse from json.
func (s *BlockedUserIdsResponse) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode BlockedUserIdsResponse to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "block_ids":
			if err := func() error {
				s.BlockIds.Reset()
				if err := s.BlockIds.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"block_ids\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode BlockedUserIdsResponse")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *BlockedUserIdsResponse) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *BlockedUserIdsResponse) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *BlockedUsersResponse) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *BlockedUsersResponse) encodeFields(e *jx.Encoder) {
	{
		if s.BlockedCount.Set {
			e.FieldStart("blocked_count")
			s.BlockedCount.Encode(e)
		}
	}
	{
		if s.LastID.Set {
			e.FieldStart("last_id")
			s.LastID.Encode(e)
		}
	}
	{
		if s.Users.Set {
			e.FieldStart("users")
			s.Users.Encode(e)
		}
	}
}

var jsonFieldsNameOfBlockedUsersResponse = [3]string{
	0: "blocked_count",
	1: "last_id",
	2: "users",
}

// Decode decodes BlockedUsersResponse from json.
func (s *BlockedUsersResponse) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode BlockedUsersResponse to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "blocked_count":
			if err := func() error {
				s.BlockedCount.Reset()
				if err := s.BlockedCount.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"blocked_count\"")
			}
		case "last_id":
			if err := func() error {
				s.LastID.Reset()
				if err := s.LastID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"last_id\"")
			}
		case "users":
			if err := func() error {
				s.Users.Reset()
				if err := s.Users.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"users\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode BlockedUsersResponse")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *BlockedUsersResponse) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *BlockedUsersResponse) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *BookmarkPostResponse) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *BookmarkPostResponse) encodeFields(e *jx.Encoder) {
	{
		if s.IsBookmarked.Set {
			e.FieldStart("is_bookmarked")
			s.IsBookmarked.Encode(e)
		}
	}
}

var jsonFieldsNameOfBookmarkPostResponse = [1]string{
	0: "is_bookmarked",
}

// Decode decodes BookmarkPostResponse from json.
func (s *BookmarkPostResponse) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode BookmarkPostResponse to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "is_bookmarked":
			if err := func() error {
				s.IsBookmarked.Reset()
				if err := s.IsBookmarked.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"is_bookmarked\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode BookmarkPostResponse")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *BookmarkPostResponse) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *BookmarkPostResponse) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *CallActionSignatureResponse) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *CallActionSignatureResponse) encodeFields(e *jx.Encoder) {
	{
		if s.SignaturePayload.Set {
			e.FieldStart("signature_payload")
			s.SignaturePayload.Encode(e)
		}
	}
}

var jsonFieldsNameOfCallActionSignatureResponse = [1]string{
	0: "signature_payload",
}

// Decode decodes CallActionSignatureResponse from json.
func (s *CallActionSignatureResponse) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode CallActionSignatureResponse to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "signature_payload":
			if err := func() error {
				s.SignaturePayload.Reset()
				if err := s.SignaturePayload.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"signature_payload\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode CallActionSignatureResponse")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *CallActionSignatureResponse) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *CallActionSignatureResponse) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *CallGiftHistoryResponse) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *CallGiftHistoryResponse) encodeFields(e *jx.Encoder) {
	{
		if s.NextPageValue.Set {
			e.FieldStart("next_page_value")
			s.NextPageValue.Encode(e)
		}
	}
	{
		if s.SentGifts.Set {
			e.FieldStart("sent_gifts")
			s.SentGifts.Encode(e)
		}
	}
}

var jsonFieldsNameOfCallGiftHistoryResponse = [2]string{
	0: "next_page_value",
	1: "sent_gifts",
}

// Decode decodes CallGiftHistoryResponse from json.
func (s *CallGiftHistoryResponse) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode CallGiftHistoryResponse to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "next_page_value":
			if err := func() error {
				s.NextPageValue.Reset()
				if err := s.NextPageValue.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"next_page_value\"")
			}
		case "sent_gifts":
			if err := func() error {
				s.SentGifts.Reset()
				if err := s.SentGifts.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"sent_gifts\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode CallGiftHistoryResponse")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *CallGiftHistoryResponse) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *CallGiftHistoryResponse) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *CallStatusResponse) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *CallStatusResponse) encodeFields(e *jx.Encoder) {
	{
		if s.PhoneStatus.Set {
			e.FieldStart("phone_status")
			s.PhoneStatus.Encode(e)
		}
	}
	{
		if s.RoomURL.Set {
			e.FieldStart("room_url")
			s.RoomURL.Encode(e)
		}
	}
	{
		if s.VideoStatus.Set {
			e.FieldStart("video_status")
			s.VideoStatus.Encode(e)
		}
	}
}

var jsonFieldsNameOfCallStatusResponse = [3]string{
	0: "phone_status",
	1: "room_url",
	2: "video_status",
}

// Decode decodes CallStatusResponse from json.
func (s *CallStatusResponse) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode CallStatusResponse to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "phone_status":
			if err := func() error {
				s.PhoneStatus.Reset()
				if err := s.PhoneStatus.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"phone_status\"")
			}
		case "room_url":
			if err := func() error {
				s.RoomURL.Reset()
				if err := s.RoomURL.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"room_url\"")
			}
		case "video_status":
			if err := func() error {
				s.VideoStatus.Reset()
				if err := s.VideoStatus.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"video_status\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode CallStatusResponse")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *CallStatusResponse) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *CallStatusResponse) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *ChatInvitation) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *ChatInvitation) encodeFields(e *jx.Encoder) {
	{
		if s.FirstUser.Set {
			e.FieldStart("first_user")
			s.FirstUser.Encode(e)
		}
	}
	{
		if s.Inviter.Set {
			e.FieldStart("inviter")
			s.Inviter.Encode(e)
		}
	}
	{
		if s.UsersCount.Set {
			e.FieldStart("users_count")
			s.UsersCount.Encode(e)
		}
	}
}

var jsonFieldsNameOfChatInvitation = [3]string{
	0: "first_user",
	1: "inviter",
	2: "users_count",
}

// Decode decodes ChatInvitation from json.
func (s *ChatInvitation) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode ChatInvitation to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "first_user":
			if err := func() error {
				s.FirstUser.Reset()
				if err := s.FirstUser.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"first_user\"")
			}
		case "inviter":
			if err := func() error {
				s.Inviter.Reset()
				if err := s.Inviter.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"inviter\"")
			}
		case "users_count":
			if err := func() error {
				s.UsersCount.Reset()
				if err := s.UsersCount.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"users_count\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode ChatInvitation")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *ChatInvitation) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *ChatInvitation) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *ChatRoomLastMessage) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *ChatRoomLastMessage) encodeFields(e *jx.Encoder) {
	{
		if s.ConferenceCall.Set {
			e.FieldStart("conference_call")
			s.ConferenceCall.Encode(e)
		}
	}
	{
		if s.CreatedAt.Set {
			e.FieldStart("created_at")
			s.CreatedAt.Encode(e)
		}
	}
	{
		if s.ID.Set {
			e.FieldStart("id")
			s.ID.Encode(e)
		}
	}
	{
		if s.MessageType.Set {
			e.FieldStart("message_type")
			s.MessageType.Encode(e)
		}
	}
	{
		if s.RoomID.Set {
			e.FieldStart("room_id")
			s.RoomID.Encode(e)
		}
	}
	{
		if s.Text.Set {
			e.FieldStart("text")
			s.Text.Encode(e)
		}
	}
	{
		if s.UserID.Set {
			e.FieldStart("user_id")
			s.UserID.Encode(e)
		}
	}
}

var jsonFieldsNameOfChatRoomLastMessage = [7]string{
	0: "conference_call",
	1: "created_at",
	2: "id",
	3: "message_type",
	4: "room_id",
	5: "text",
	6: "user_id",
}

// Decode decodes ChatRoomLastMessage from json.
func (s *ChatRoomLastMessage) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode ChatRoomLastMessage to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "conference_call":
			if err := func() error {
				s.ConferenceCall.Reset()
				if err := s.ConferenceCall.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"conference_call\"")
			}
		case "created_at":
			if err := func() error {
				s.CreatedAt.Reset()
				if err := s.CreatedAt.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"created_at\"")
			}
		case "id":
			if err := func() error {
				s.ID.Reset()
				if err := s.ID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"id\"")
			}
		case "message_type":
			if err := func() error {
				s.MessageType.Reset()
				if err := s.MessageType.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"message_type\"")
			}
		case "room_id":
			if err := func() error {
				s.RoomID.Reset()
				if err := s.RoomID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"room_id\"")
			}
		case "text":
			if err := func() error {
				s.Text.Reset()
				if err := s.Text.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"text\"")
			}
		case "user_id":
			if err := func() error {
				s.UserID.Reset()
				if err := s.UserID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"user_id\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode ChatRoomLastMessage")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *ChatRoomLastMessage) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *ChatRoomLastMessage) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *ChatRoomResponse) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *ChatRoomResponse) encodeFields(e *jx.Encoder) {
	{
		if s.Chat != nil {
			e.FieldStart("chat")
			s.Chat.Encode(e)
		}
	}
}

var jsonFieldsNameOfChatRoomResponse = [1]string{
	0: "chat",
}

// Decode decodes ChatRoomResponse from json.
func (s *ChatRoomResponse) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode ChatRoomResponse to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "chat":
			if err := func() error {
				s.Chat = nil
				var elem RealmChatRoom
				if err := elem.Decode(d); err != nil {
					return err
				}
				s.Chat = &elem
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"chat\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode ChatRoomResponse")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *ChatRoomResponse) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *ChatRoomResponse) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *ChatRoomsResponse) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *ChatRoomsResponse) encodeFields(e *jx.Encoder) {
	{
		if s.ChatRooms.Set {
			e.FieldStart("chat_rooms")
			s.ChatRooms.Encode(e)
		}
	}
	{
		if s.NextPageValue.Set {
			e.FieldStart("next_page_value")
			s.NextPageValue.Encode(e)
		}
	}
	{
		if s.PinnedChatRooms.Set {
			e.FieldStart("pinned_chat_rooms")
			s.PinnedChatRooms.Encode(e)
		}
	}
}

var jsonFieldsNameOfChatRoomsResponse = [3]string{
	0: "chat_rooms",
	1: "next_page_value",
	2: "pinned_chat_rooms",
}

// Decode decodes ChatRoomsResponse from json.
func (s *ChatRoomsResponse) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode ChatRoomsResponse to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "chat_rooms":
			if err := func() error {
				s.ChatRooms.Reset()
				if err := s.ChatRooms.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"chat_rooms\"")
			}
		case "next_page_value":
			if err := func() error {
				s.NextPageValue.Reset()
				if err := s.NextPageValue.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"next_page_value\"")
			}
		case "pinned_chat_rooms":
			if err := func() error {
				s.PinnedChatRooms.Reset()
				if err := s.PinnedChatRooms.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"pinned_chat_rooms\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode ChatRoomsResponse")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *ChatRoomsResponse) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *ChatRoomsResponse) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *Choice) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *Choice) encodeFields(e *jx.Encoder) {
	{
		if s.ID.Set {
			e.FieldStart("id")
			s.ID.Encode(e)
		}
	}
	{
		if s.Label.Set {
			e.FieldStart("label")
			s.Label.Encode(e)
		}
	}
	{
		if s.VotesCount.Set {
			e.FieldStart("votes_count")
			s.VotesCount.Encode(e)
		}
	}
}

var jsonFieldsNameOfChoice = [3]string{
	0: "id",
	1: "label",
	2: "votes_count",
}

// Decode decodes Choice from json.
func (s *Choice) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Choice to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "id":
			if err := func() error {
				s.ID.Reset()
				if err := s.ID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"id\"")
			}
		case "label":
			if err := func() error {
				s.Label.Reset()
				if err := s.Label.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"label\"")
			}
		case "votes_count":
			if err := func() error {
				s.VotesCount.Reset()
				if err := s.VotesCount.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"votes_count\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode Choice")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *Choice) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *Choice) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *CommonIdsRequest) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *CommonIdsRequest) encodeFields(e *jx.Encoder) {
	{
		if s.Ids.Set {
			e.FieldStart("ids")
			s.Ids.Encode(e)
		}
	}
}

var jsonFieldsNameOfCommonIdsRequest = [1]string{
	0: "ids",
}

// Decode decodes CommonIdsRequest from json.
func (s *CommonIdsRequest) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode CommonIdsRequest to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "ids":
			if err := func() error {
				s.Ids.Reset()
				if err := s.Ids.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"ids\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode CommonIdsRequest")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *CommonIdsRequest) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *CommonIdsRequest) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *CommonUrlResponse) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *CommonUrlResponse) encodeFields(e *jx.Encoder) {
	{
		if s.URL.Set {
			e.FieldStart("url")
			s.URL.Encode(e)
		}
	}
}

var jsonFieldsNameOfCommonUrlResponse = [1]string{
	0: "url",
}

// Decode decodes CommonUrlResponse from json.
func (s *CommonUrlResponse) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode CommonUrlResponse to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "url":
			if err := func() error {
				s.URL.Reset()
				if err := s.URL.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"url\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode CommonUrlResponse")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *CommonUrlResponse) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *CommonUrlResponse) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *ConferenceCallBumpParams) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *ConferenceCallBumpParams) encodeFields(e *jx.Encoder) {
	{
		if s.ParticipantLimit.Set {
			e.FieldStart("participant_limit")
			s.ParticipantLimit.Encode(e)
		}
	}
}

var jsonFieldsNameOfConferenceCallBumpParams = [1]string{
	0: "participant_limit",
}

// Decode decodes ConferenceCallBumpParams from json.
func (s *ConferenceCallBumpParams) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode ConferenceCallBumpParams to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "participant_limit":
			if err := func() error {
				s.ParticipantLimit.Reset()
				if err := s.ParticipantLimit.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"participant_limit\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode ConferenceCallBumpParams")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *ConferenceCallBumpParams) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *ConferenceCallBumpParams) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *ConferenceCallResponse) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *ConferenceCallResponse) encodeFields(e *jx.Encoder) {
	{
		if s.ConferenceCall.Set {
			e.FieldStart("conference_call")
			s.ConferenceCall.Encode(e)
		}
	}
	{
		if s.ConferenceCallUserUUID.Set {
			e.FieldStart("conference_call_user_uuid")
			s.ConferenceCallUserUUID.Encode(e)
		}
	}
}

var jsonFieldsNameOfConferenceCallResponse = [2]string{
	0: "conference_call",
	1: "conference_call_user_uuid",
}

// Decode decodes ConferenceCallResponse from json.
func (s *ConferenceCallResponse) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode ConferenceCallResponse to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "conference_call":
			if err := func() error {
				s.ConferenceCall.Reset()
				if err := s.ConferenceCall.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"conference_call\"")
			}
		case "conference_call_user_uuid":
			if err := func() error {
				s.ConferenceCallUserUUID.Reset()
				if err := s.ConferenceCallUserUUID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"conference_call_user_uuid\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode ConferenceCallResponse")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *ConferenceCallResponse) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *ConferenceCallResponse) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *ConferenceCallUserRole) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *ConferenceCallUserRole) encodeFields(e *jx.Encoder) {
	{
		if s.ID.Set {
			e.FieldStart("id")
			s.ID.Encode(e)
		}
	}
	{
		if s.Role.Set {
			e.FieldStart("role")
			s.Role.Encode(e)
		}
	}
	{
		if s.UserID.Set {
			e.FieldStart("user_id")
			s.UserID.Encode(e)
		}
	}
}

var jsonFieldsNameOfConferenceCallUserRole = [3]string{
	0: "id",
	1: "role",
	2: "user_id",
}

// Decode decodes ConferenceCallUserRole from json.
func (s *ConferenceCallUserRole) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode ConferenceCallUserRole to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "id":
			if err := func() error {
				s.ID.Reset()
				if err := s.ID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"id\"")
			}
		case "role":
			if err := func() error {
				s.Role.Reset()
				if err := s.Role.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"role\"")
			}
		case "user_id":
			if err := func() error {
				s.UserID.Reset()
				if err := s.UserID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"user_id\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode ConferenceCallUserRole")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *ConferenceCallUserRole) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *ConferenceCallUserRole) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *Config) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *Config) encodeFields(e *jx.Encoder) {
	{
		if s.AdTesterUserIds.Set {
			e.FieldStart("ad_tester_user_ids")
			s.AdTesterUserIds.Encode(e)
		}
	}
	{
		if s.AndroidNewsVersion.Set {
			e.FieldStart("android_news_version")
			s.AndroidNewsVersion.Encode(e)
		}
	}
	{
		if s.Dapps.Set {
			e.FieldStart("dapps")
			s.Dapps.Encode(e)
		}
	}
	{
		if s.IsChatWebsocketEnabled.Set {
			e.FieldStart("is_chat_websocket_enabled")
			s.IsChatWebsocketEnabled.Encode(e)
		}
	}
	{
		if s.IsDirectVipPurchaseEnabled.Set {
			e.FieldStart("is_direct_vip_purchase_enabled")
			s.IsDirectVipPurchaseEnabled.Encode(e)
		}
	}
	{
		if s.IsGiftFeaturesEnabled.Set {
			e.FieldStart("is_gift_features_enabled")
			s.IsGiftFeaturesEnabled.Encode(e)
		}
	}
	{
		if s.IsIDCardAndPhoneVerificationCheckForReviewEnabled.Set {
			e.FieldStart("is_id_card_and_phone_verification_check_for_review_enabled")
			s.IsIDCardAndPhoneVerificationCheckForReviewEnabled.Encode(e)
		}
	}
	{
		if s.IsIDCardCheckRequiredBlockerViewEnabled.Set {
			e.FieldStart("is_id_card_check_required_blocker_view_enabled")
			s.IsIDCardCheckRequiredBlockerViewEnabled.Encode(e)
		}
	}
	{
		if s.IsMaintenanced.Set {
			e.FieldStart("is_maintenanced")
			s.IsMaintenanced.Encode(e)
		}
	}
	{
		if s.IsPhoneVerificationRequiredBlockerViewEnabled.Set {
			e.FieldStart("is_phone_verification_required_blocker_view_enabled")
			s.IsPhoneVerificationRequiredBlockerViewEnabled.Encode(e)
		}
	}
	{
		if s.IsSpeedTestEnabled.Set {
			e.FieldStart("is_speed_test_enabled")
			s.IsSpeedTestEnabled.Encode(e)
		}
	}
	{
		if s.IsSslPinningEnabled.Set {
			e.FieldStart("is_ssl_pinning_enabled")
			s.IsSslPinningEnabled.Encode(e)
		}
	}
	{
		if s.IsStarEnabled.Set {
			e.FieldStart("is_star_enabled")
			s.IsStarEnabled.Encode(e)
		}
	}
	{
		if s.LatestAndroidAppVersion.Set {
			e.FieldStart("latest_android_app_version")
			s.LatestAndroidAppVersion.Encode(e)
		}
	}
	{
		if s.LineOfficialAccountID.Set {
			e.FieldStart("line_official_account_id")
			s.LineOfficialAccountID.Encode(e)
		}
	}
	{
		if s.LocalizedAndroidNewsTitle.Set {
			e.FieldStart("localized_android_news_title")
			s.LocalizedAndroidNewsTitle.Encode(e)
		}
	}
	{
		if s.LocalizedAndroidNewsURL.Set {
			e.FieldStart("localized_android_news_url")
			s.LocalizedAndroidNewsURL.Encode(e)
		}
	}
	{
		if s.LocalizedMaintenanceURL.Set {
			e.FieldStart("localized_maintenance_url")
			s.LocalizedMaintenanceURL.Encode(e)
		}
	}
	{
		if s.LocalizedNewsTitle.Set {
			e.FieldStart("localized_news_title")
			s.LocalizedNewsTitle.Encode(e)
		}
	}
	{
		if s.LocalizedNewsURL.Set {
			e.FieldStart("localized_news_url")
			s.LocalizedNewsURL.Encode(e)
		}
	}
	{
		if s.MaintenanceFeatures.Set {
			e.FieldStart("maintenance_features")
			s.MaintenanceFeatures.Encode(e)
		}
	}
	{
		if s.MaxImageFrameCount.Set {
			e.FieldStart("max_image_frame_count")
			s.MaxImageFrameCount.Encode(e)
		}
	}
	{
		if s.MinimumAndroidAppVersionRequired.Set {
			e.FieldStart("minimum_android_app_version_required")
			s.MinimumAndroidAppVersionRequired.Encode(e)
		}
	}
	{
		if s.MinimumAndroidVersionSupported.Set {
			e.FieldStart("minimum_android_version_supported")
			s.MinimumAndroidVersionSupported.Encode(e)
		}
	}
	{
		if s.NewsVersion.Set {
			e.FieldStart("news_version")
			s.NewsVersion.Encode(e)
		}
	}
	{
		if s.NftInfos.Set {
			e.FieldStart("nft_infos")
			s.NftInfos.Encode(e)
		}
	}
	{
		if s.PromotionStickerPackIds.Set {
			e.FieldStart("promotion_sticker_pack_ids")
			s.PromotionStickerPackIds.Encode(e)
		}
	}
	{
		if s.ShouldAppendUserIDToNewsURL.Set {
			e.FieldStart("should_append_user_id_to_news_url")
			s.ShouldAppendUserIDToNewsURL.Encode(e)
		}
	}
	{
		if s.SupportEmailAddress.Set {
			e.FieldStart("support_email_address")
			s.SupportEmailAddress.Encode(e)
		}
	}
	{
		if s.TokenInfos.Set {
			e.FieldStart("token_infos")
			s.TokenInfos.Encode(e)
		}
	}
	{
		if s.UseRandomMessageRefreshRate.Set {
			e.FieldStart("use_random_message_refresh_rate")
			s.UseRandomMessageRefreshRate.Encode(e)
		}
	}
	{
		if s.Web3IsMaintenanced.Set {
			e.FieldStart("web3_is_maintenanced")
			s.Web3IsMaintenanced.Encode(e)
		}
	}
	{
		if s.Web3LocalizedMaintenanceURL.Set {
			e.FieldStart("web3_localized_maintenance_url")
			s.Web3LocalizedMaintenanceURL.Encode(e)
		}
	}
	{
		if s.Web3MaintenanceFeatures.Set {
			e.FieldStart("web3_maintenance_features")
			s.Web3MaintenanceFeatures.Encode(e)
		}
	}
	{
		if s.Web3MaintenanceURL.Set {
			e.FieldStart("web3_maintenance_url")
			s.Web3MaintenanceURL.Encode(e)
		}
	}
	{
		if s.XYayGlobalID.Set {
			e.FieldStart("x_yay_global_id")
			s.XYayGlobalID.Encode(e)
		}
	}
	{
		if s.XYayJpID.Set {
			e.FieldStart("x_yay_jp_id")
			s.XYayJpID.Encode(e)
		}
	}
	{
		if s.XYayUpdateID.Set {
			e.FieldStart("x_yay_update_id")
			s.XYayUpdateID.Encode(e)
		}
	}
}

var jsonFieldsNameOfConfig = [38]string{
	0:  "ad_tester_user_ids",
	1:  "android_news_version",
	2:  "dapps",
	3:  "is_chat_websocket_enabled",
	4:  "is_direct_vip_purchase_enabled",
	5:  "is_gift_features_enabled",
	6:  "is_id_card_and_phone_verification_check_for_review_enabled",
	7:  "is_id_card_check_required_blocker_view_enabled",
	8:  "is_maintenanced",
	9:  "is_phone_verification_required_blocker_view_enabled",
	10: "is_speed_test_enabled",
	11: "is_ssl_pinning_enabled",
	12: "is_star_enabled",
	13: "latest_android_app_version",
	14: "line_official_account_id",
	15: "localized_android_news_title",
	16: "localized_android_news_url",
	17: "localized_maintenance_url",
	18: "localized_news_title",
	19: "localized_news_url",
	20: "maintenance_features",
	21: "max_image_frame_count",
	22: "minimum_android_app_version_required",
	23: "minimum_android_version_supported",
	24: "news_version",
	25: "nft_infos",
	26: "promotion_sticker_pack_ids",
	27: "should_append_user_id_to_news_url",
	28: "support_email_address",
	29: "token_infos",
	30: "use_random_message_refresh_rate",
	31: "web3_is_maintenanced",
	32: "web3_localized_maintenance_url",
	33: "web3_maintenance_features",
	34: "web3_maintenance_url",
	35: "x_yay_global_id",
	36: "x_yay_jp_id",
	37: "x_yay_update_id",
}

// Decode decodes Config from json.
func (s *Config) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Config to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "ad_tester_user_ids":
			if err := func() error {
				s.AdTesterUserIds.Reset()
				if err := s.AdTesterUserIds.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"ad_tester_user_ids\"")
			}
		case "android_news_version":
			if err := func() error {
				s.AndroidNewsVersion.Reset()
				if err := s.AndroidNewsVersion.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"android_news_version\"")
			}
		case "dapps":
			if err := func() error {
				s.Dapps.Reset()
				if err := s.Dapps.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"dapps\"")
			}
		case "is_chat_websocket_enabled":
			if err := func() error {
				s.IsChatWebsocketEnabled.Reset()
				if err := s.IsChatWebsocketEnabled.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"is_chat_websocket_enabled\"")
			}
		case "is_direct_vip_purchase_enabled":
			if err := func() error {
				s.IsDirectVipPurchaseEnabled.Reset()
				if err := s.IsDirectVipPurchaseEnabled.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"is_direct_vip_purchase_enabled\"")
			}
		case "is_gift_features_enabled":
			if err := func() error {
				s.IsGiftFeaturesEnabled.Reset()
				if err := s.IsGiftFeaturesEnabled.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"is_gift_features_enabled\"")
			}
		case "is_id_card_and_phone_verification_check_for_review_enabled":
			if err := func() error {
				s.IsIDCardAndPhoneVerificationCheckForReviewEnabled.Reset()
				if err := s.IsIDCardAndPhoneVerificationCheckForReviewEnabled.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"is_id_card_and_phone_verification_check_for_review_enabled\"")
			}
		case "is_id_card_check_required_blocker_view_enabled":
			if err := func() error {
				s.IsIDCardCheckRequiredBlockerViewEnabled.Reset()
				if err := s.IsIDCardCheckRequiredBlockerViewEnabled.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"is_id_card_check_required_blocker_view_enabled\"")
			}
		case "is_maintenanced":
			if err := func() error {
				s.IsMaintenanced.Reset()
				if err := s.IsMaintenanced.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"is_maintenanced\"")
			}
		case "is_phone_verification_required_blocker_view_enabled":
			if err := func() error {
				s.IsPhoneVerificationRequiredBlockerViewEnabled.Reset()
				if err := s.IsPhoneVerificationRequiredBlockerViewEnabled.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"is_phone_verification_required_blocker_view_enabled\"")
			}
		case "is_speed_test_enabled":
			if err := func() error {
				s.IsSpeedTestEnabled.Reset()
				if err := s.IsSpeedTestEnabled.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"is_speed_test_enabled\"")
			}
		case "is_ssl_pinning_enabled":
			if err := func() error {
				s.IsSslPinningEnabled.Reset()
				if err := s.IsSslPinningEnabled.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"is_ssl_pinning_enabled\"")
			}
		case "is_star_enabled":
			if err := func() error {
				s.IsStarEnabled.Reset()
				if err := s.IsStarEnabled.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"is_star_enabled\"")
			}
		case "latest_android_app_version":
			if err := func() error {
				s.LatestAndroidAppVersion.Reset()
				if err := s.LatestAndroidAppVersion.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"latest_android_app_version\"")
			}
		case "line_official_account_id":
			if err := func() error {
				s.LineOfficialAccountID.Reset()
				if err := s.LineOfficialAccountID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"line_official_account_id\"")
			}
		case "localized_android_news_title":
			if err := func() error {
				s.LocalizedAndroidNewsTitle.Reset()
				if err := s.LocalizedAndroidNewsTitle.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"localized_android_news_title\"")
			}
		case "localized_android_news_url":
			if err := func() error {
				s.LocalizedAndroidNewsURL.Reset()
				if err := s.LocalizedAndroidNewsURL.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"localized_android_news_url\"")
			}
		case "localized_maintenance_url":
			if err := func() error {
				s.LocalizedMaintenanceURL.Reset()
				if err := s.LocalizedMaintenanceURL.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"localized_maintenance_url\"")
			}
		case "localized_news_title":
			if err := func() error {
				s.LocalizedNewsTitle.Reset()
				if err := s.LocalizedNewsTitle.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"localized_news_title\"")
			}
		case "localized_news_url":
			if err := func() error {
				s.LocalizedNewsURL.Reset()
				if err := s.LocalizedNewsURL.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"localized_news_url\"")
			}
		case "maintenance_features":
			if err := func() error {
				s.MaintenanceFeatures.Reset()
				if err := s.MaintenanceFeatures.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"maintenance_features\"")
			}
		case "max_image_frame_count":
			if err := func() error {
				s.MaxImageFrameCount.Reset()
				if err := s.MaxImageFrameCount.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"max_image_frame_count\"")
			}
		case "minimum_android_app_version_required":
			if err := func() error {
				s.MinimumAndroidAppVersionRequired.Reset()
				if err := s.MinimumAndroidAppVersionRequired.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"minimum_android_app_version_required\"")
			}
		case "minimum_android_version_supported":
			if err := func() error {
				s.MinimumAndroidVersionSupported.Reset()
				if err := s.MinimumAndroidVersionSupported.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"minimum_android_version_supported\"")
			}
		case "news_version":
			if err := func() error {
				s.NewsVersion.Reset()
				if err := s.NewsVersion.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"news_version\"")
			}
		case "nft_infos":
			if err := func() error {
				s.NftInfos.Reset()
				if err := s.NftInfos.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"nft_infos\"")
			}
		case "promotion_sticker_pack_ids":
			if err := func() error {
				s.PromotionStickerPackIds.Reset()
				if err := s.PromotionStickerPackIds.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"promotion_sticker_pack_ids\"")
			}
		case "should_append_user_id_to_news_url":
			if err := func() error {
				s.ShouldAppendUserIDToNewsURL.Reset()
				if err := s.ShouldAppendUserIDToNewsURL.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"should_append_user_id_to_news_url\"")
			}
		case "support_email_address":
			if err := func() error {
				s.SupportEmailAddress.Reset()
				if err := s.SupportEmailAddress.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"support_email_address\"")
			}
		case "token_infos":
			if err := func() error {
				s.TokenInfos.Reset()
				if err := s.TokenInfos.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"token_infos\"")
			}
		case "use_random_message_refresh_rate":
			if err := func() error {
				s.UseRandomMessageRefreshRate.Reset()
				if err := s.UseRandomMessageRefreshRate.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"use_random_message_refresh_rate\"")
			}
		case "web3_is_maintenanced":
			if err := func() error {
				s.Web3IsMaintenanced.Reset()
				if err := s.Web3IsMaintenanced.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"web3_is_maintenanced\"")
			}
		case "web3_localized_maintenance_url":
			if err := func() error {
				s.Web3LocalizedMaintenanceURL.Reset()
				if err := s.Web3LocalizedMaintenanceURL.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"web3_localized_maintenance_url\"")
			}
		case "web3_maintenance_features":
			if err := func() error {
				s.Web3MaintenanceFeatures.Reset()
				if err := s.Web3MaintenanceFeatures.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"web3_maintenance_features\"")
			}
		case "web3_maintenance_url":
			if err := func() error {
				s.Web3MaintenanceURL.Reset()
				if err := s.Web3MaintenanceURL.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"web3_maintenance_url\"")
			}
		case "x_yay_global_id":
			if err := func() error {
				s.XYayGlobalID.Reset()
				if err := s.XYayGlobalID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"x_yay_global_id\"")
			}
		case "x_yay_jp_id":
			if err := func() error {
				s.XYayJpID.Reset()
				if err := s.XYayJpID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"x_yay_jp_id\"")
			}
		case "x_yay_update_id":
			if err := func() error {
				s.XYayUpdateID.Reset()
				if err := s.XYayUpdateID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"x_yay_update_id\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode Config")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *Config) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *Config) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *CreateChatRoomResponse) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *CreateChatRoomResponse) encodeFields(e *jx.Encoder) {
	{
		if s.RoomID.Set {
			e.FieldStart("room_id")
			s.RoomID.Encode(e)
		}
	}
}

var jsonFieldsNameOfCreateChatRoomResponse = [1]string{
	0: "room_id",
}

// Decode decodes CreateChatRoomResponse from json.
func (s *CreateChatRoomResponse) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode CreateChatRoomResponse to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "room_id":
			if err := func() error {
				s.RoomID.Reset()
				if err := s.RoomID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"room_id\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode CreateChatRoomResponse")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *CreateChatRoomResponse) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *CreateChatRoomResponse) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *CreateConferenceCallPostReqMessageTags) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *CreateConferenceCallPostReqMessageTags) encodeFields(e *jx.Encoder) {
}

var jsonFieldsNameOfCreateConferenceCallPostReqMessageTags = [0]string{}

// Decode decodes CreateConferenceCallPostReqMessageTags from json.
func (s *CreateConferenceCallPostReqMessageTags) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode CreateConferenceCallPostReqMessageTags to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		default:
			return d.Skip()
		}
	}); err != nil {
		return errors.Wrap(err, "decode CreateConferenceCallPostReqMessageTags")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *CreateConferenceCallPostReqMessageTags) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *CreateConferenceCallPostReqMessageTags) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *CreateGroupQuota) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *CreateGroupQuota) encodeFields(e *jx.Encoder) {
	{
		if s.RemainingQuota.Set {
			e.FieldStart("remaining_quota")
			s.RemainingQuota.Encode(e)
		}
	}
	{
		if s.UsedQuota.Set {
			e.FieldStart("used_quota")
			s.UsedQuota.Encode(e)
		}
	}
}

var jsonFieldsNameOfCreateGroupQuota = [2]string{
	0: "remaining_quota",
	1: "used_quota",
}

// Decode decodes CreateGroupQuota from json.
func (s *CreateGroupQuota) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode CreateGroupQuota to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "remaining_quota":
			if err := func() error {
				s.RemainingQuota.Reset()
				if err := s.RemainingQuota.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"remaining_quota\"")
			}
		case "used_quota":
			if err := func() error {
				s.UsedQuota.Reset()
				if err := s.UsedQuota.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"used_quota\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode CreateGroupQuota")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *CreateGroupQuota) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *CreateGroupQuota) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *CreateGroupResponse) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *CreateGroupResponse) encodeFields(e *jx.Encoder) {
	{
		if s.GroupID.Set {
			e.FieldStart("group_id")
			s.GroupID.Encode(e)
		}
	}
}

var jsonFieldsNameOfCreateGroupResponse = [1]string{
	0: "group_id",
}

// Decode decodes CreateGroupResponse from json.
func (s *CreateGroupResponse) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode CreateGroupResponse to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "group_id":
			if err := func() error {
				s.GroupID.Reset()
				if err := s.GroupID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"group_id\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode CreateGroupResponse")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *CreateGroupResponse) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *CreateGroupResponse) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *CreateGroupThreadRequest) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *CreateGroupThreadRequest) encodeFields(e *jx.Encoder) {
	{
		if s.GroupID.Set {
			e.FieldStart("group_id")
			s.GroupID.Encode(e)
		}
	}
	{
		if s.ThreadIconFilename.Set {
			e.FieldStart("thread_icon_filename")
			s.ThreadIconFilename.Encode(e)
		}
	}
	{
		if s.Title.Set {
			e.FieldStart("title")
			s.Title.Encode(e)
		}
	}
}

var jsonFieldsNameOfCreateGroupThreadRequest = [3]string{
	0: "group_id",
	1: "thread_icon_filename",
	2: "title",
}

// Decode decodes CreateGroupThreadRequest from json.
func (s *CreateGroupThreadRequest) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode CreateGroupThreadRequest to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "group_id":
			if err := func() error {
				s.GroupID.Reset()
				if err := s.GroupID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"group_id\"")
			}
		case "thread_icon_filename":
			if err := func() error {
				s.ThreadIconFilename.Reset()
				if err := s.ThreadIconFilename.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"thread_icon_filename\"")
			}
		case "title":
			if err := func() error {
				s.Title.Reset()
				if err := s.Title.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"title\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode CreateGroupThreadRequest")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *CreateGroupThreadRequest) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *CreateGroupThreadRequest) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *CreateMuteKeywordResponse) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *CreateMuteKeywordResponse) encodeFields(e *jx.Encoder) {
	{
		if s.HiddenWord.Set {
			e.FieldStart("hidden_word")
			s.HiddenWord.Encode(e)
		}
	}
}

var jsonFieldsNameOfCreateMuteKeywordResponse = [1]string{
	0: "hidden_word",
}

// Decode decodes CreateMuteKeywordResponse from json.
func (s *CreateMuteKeywordResponse) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode CreateMuteKeywordResponse to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "hidden_word":
			if err := func() error {
				s.HiddenWord.Reset()
				if err := s.HiddenWord.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"hidden_word\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode CreateMuteKeywordResponse")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *CreateMuteKeywordResponse) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *CreateMuteKeywordResponse) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *CreatePostReqMessageTags) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *CreatePostReqMessageTags) encodeFields(e *jx.Encoder) {
}

var jsonFieldsNameOfCreatePostReqMessageTags = [0]string{}

// Decode decodes CreatePostReqMessageTags from json.
func (s *CreatePostReqMessageTags) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode CreatePostReqMessageTags to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		default:
			return d.Skip()
		}
	}); err != nil {
		return errors.Wrap(err, "decode CreatePostReqMessageTags")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *CreatePostReqMessageTags) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *CreatePostReqMessageTags) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *CreatePostReqSharedURL) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *CreatePostReqSharedURL) encodeFields(e *jx.Encoder) {
}

var jsonFieldsNameOfCreatePostReqSharedURL = [0]string{}

// Decode decodes CreatePostReqSharedURL from json.
func (s *CreatePostReqSharedURL) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode CreatePostReqSharedURL to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		default:
			return d.Skip()
		}
	}); err != nil {
		return errors.Wrap(err, "decode CreatePostReqSharedURL")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *CreatePostReqSharedURL) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *CreatePostReqSharedURL) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *CreatePostResponse) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *CreatePostResponse) encodeFields(e *jx.Encoder) {
	{
		if s.ConferenceCall.Set {
			e.FieldStart("conference_call")
			s.ConferenceCall.Encode(e)
		}
	}
	{
		if s.Post != nil {
			e.FieldStart("post")
			s.Post.Encode(e)
		}
	}
}

var jsonFieldsNameOfCreatePostResponse = [2]string{
	0: "conference_call",
	1: "post",
}

// Decode decodes CreatePostResponse from json.
func (s *CreatePostResponse) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode CreatePostResponse to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "conference_call":
			if err := func() error {
				s.ConferenceCall.Reset()
				if err := s.ConferenceCall.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"conference_call\"")
			}
		case "post":
			if err := func() error {
				s.Post = nil
				var elem Post
				if err := elem.Decode(d); err != nil {
					return err
				}
				s.Post = &elem
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"post\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode CreatePostResponse")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *CreatePostResponse) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *CreatePostResponse) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *CreateQuotaResponse) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *CreateQuotaResponse) encodeFields(e *jx.Encoder) {
	{
		if s.Create.Set {
			e.FieldStart("create")
			s.Create.Encode(e)
		}
	}
}

var jsonFieldsNameOfCreateQuotaResponse = [1]string{
	0: "create",
}

// Decode decodes CreateQuotaResponse from json.
func (s *CreateQuotaResponse) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode CreateQuotaResponse to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "create":
			if err := func() error {
				s.Create.Reset()
				if err := s.Create.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"create\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode CreateQuotaResponse")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *CreateQuotaResponse) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *CreateQuotaResponse) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *CreateSharePostReqMessageTags) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *CreateSharePostReqMessageTags) encodeFields(e *jx.Encoder) {
}

var jsonFieldsNameOfCreateSharePostReqMessageTags = [0]string{}

// Decode decodes CreateSharePostReqMessageTags from json.
func (s *CreateSharePostReqMessageTags) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode CreateSharePostReqMessageTags to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		default:
			return d.Skip()
		}
	}); err != nil {
		return errors.Wrap(err, "decode CreateSharePostReqMessageTags")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *CreateSharePostReqMessageTags) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *CreateSharePostReqMessageTags) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *CreateThreadPostReqMessageTags) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *CreateThreadPostReqMessageTags) encodeFields(e *jx.Encoder) {
}

var jsonFieldsNameOfCreateThreadPostReqMessageTags = [0]string{}

// Decode decodes CreateThreadPostReqMessageTags from json.
func (s *CreateThreadPostReqMessageTags) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode CreateThreadPostReqMessageTags to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		default:
			return d.Skip()
		}
	}); err != nil {
		return errors.Wrap(err, "decode CreateThreadPostReqMessageTags")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *CreateThreadPostReqMessageTags) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *CreateThreadPostReqMessageTags) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *CreateThreadPostReqSharedURL) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *CreateThreadPostReqSharedURL) encodeFields(e *jx.Encoder) {
}

var jsonFieldsNameOfCreateThreadPostReqSharedURL = [0]string{}

// Decode decodes CreateThreadPostReqSharedURL from json.
func (s *CreateThreadPostReqSharedURL) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode CreateThreadPostReqSharedURL to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		default:
			return d.Skip()
		}
	}); err != nil {
		return errors.Wrap(err, "decode CreateThreadPostReqSharedURL")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *CreateThreadPostReqSharedURL) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *CreateThreadPostReqSharedURL) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *DefaultSettingsResponse) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *DefaultSettingsResponse) encodeFields(e *jx.Encoder) {
	{
		if s.TimelineSettings.Set {
			e.FieldStart("timeline_settings")
			s.TimelineSettings.Encode(e)
		}
	}
}

var jsonFieldsNameOfDefaultSettingsResponse = [1]string{
	0: "timeline_settings",
}

// Decode decodes DefaultSettingsResponse from json.
func (s *DefaultSettingsResponse) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode DefaultSettingsResponse to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "timeline_settings":
			if err := func() error {
				s.TimelineSettings.Reset()
				if err := s.TimelineSettings.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"timeline_settings\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode DefaultSettingsResponse")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *DefaultSettingsResponse) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *DefaultSettingsResponse) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *FollowRequestCountResponse) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *FollowRequestCountResponse) encodeFields(e *jx.Encoder) {
	{
		if s.UsersCount.Set {
			e.FieldStart("users_count")
			s.UsersCount.Encode(e)
		}
	}
}

var jsonFieldsNameOfFollowRequestCountResponse = [1]string{
	0: "users_count",
}

// Decode decodes FollowRequestCountResponse from json.
func (s *FollowRequestCountResponse) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode FollowRequestCountResponse to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "users_count":
			if err := func() error {
				s.UsersCount.Reset()
				if err := s.UsersCount.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"users_count\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode FollowRequestCountResponse")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *FollowRequestCountResponse) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *FollowRequestCountResponse) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *FollowUsersResponse) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *FollowUsersResponse) encodeFields(e *jx.Encoder) {
	{
		if s.NextPageValue.Set {
			e.FieldStart("next_page_value")
			s.NextPageValue.Encode(e)
		}
	}
	{
		if s.Users.Set {
			e.FieldStart("users")
			s.Users.Encode(e)
		}
	}
}

var jsonFieldsNameOfFollowUsersResponse = [2]string{
	0: "next_page_value",
	1: "users",
}

// Decode decodes FollowUsersResponse from json.
func (s *FollowUsersResponse) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode FollowUsersResponse to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "next_page_value":
			if err := func() error {
				s.NextPageValue.Reset()
				if err := s.NextPageValue.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"next_page_value\"")
			}
		case "users":
			if err := func() error {
				s.Users.Reset()
				if err := s.Users.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"users\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode FollowUsersResponse")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *FollowUsersResponse) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *FollowUsersResponse) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *FootprintDTO) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *FootprintDTO) encodeFields(e *jx.Encoder) {
	{
		if s.ID.Set {
			e.FieldStart("id")
			s.ID.Encode(e)
		}
	}
	{
		if s.User.Set {
			e.FieldStart("user")
			s.User.Encode(e)
		}
	}
	{
		if s.VisitedAt.Set {
			e.FieldStart("visited_at")
			s.VisitedAt.Encode(e)
		}
	}
}

var jsonFieldsNameOfFootprintDTO = [3]string{
	0: "id",
	1: "user",
	2: "visited_at",
}

// Decode decodes FootprintDTO from json.
func (s *FootprintDTO) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode FootprintDTO to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "id":
			if err := func() error {
				s.ID.Reset()
				if err := s.ID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"id\"")
			}
		case "user":
			if err := func() error {
				s.User.Reset()
				if err := s.User.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"user\"")
			}
		case "visited_at":
			if err := func() error {
				s.VisitedAt.Reset()
				if err := s.VisitedAt.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"visited_at\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode FootprintDTO")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *FootprintDTO) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *FootprintDTO) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *FootprintsResponse) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *FootprintsResponse) encodeFields(e *jx.Encoder) {
	{
		if s.Footprints.Set {
			e.FieldStart("footprints")
			s.Footprints.Encode(e)
		}
	}
	{
		if s.NextPageValue.Set {
			e.FieldStart("next_page_value")
			s.NextPageValue.Encode(e)
		}
	}
}

var jsonFieldsNameOfFootprintsResponse = [2]string{
	0: "footprints",
	1: "next_page_value",
}

// Decode decodes FootprintsResponse from json.
func (s *FootprintsResponse) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode FootprintsResponse to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "footprints":
			if err := func() error {
				s.Footprints.Reset()
				if err := s.Footprints.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"footprints\"")
			}
		case "next_page_value":
			if err := func() error {
				s.NextPageValue.Reset()
				if err := s.NextPageValue.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"next_page_value\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode FootprintsResponse")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *FootprintsResponse) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *FootprintsResponse) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *GamesResponse) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *GamesResponse) encodeFields(e *jx.Encoder) {
	{
		if s.FromID.Set {
			e.FieldStart("from_id")
			s.FromID.Encode(e)
		}
	}
	{
		if s.Games.Set {
			e.FieldStart("games")
			s.Games.Encode(e)
		}
	}
}

var jsonFieldsNameOfGamesResponse = [2]string{
	0: "from_id",
	1: "games",
}

// Decode decodes GamesResponse from json.
func (s *GamesResponse) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode GamesResponse to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "from_id":
			if err := func() error {
				s.FromID.Reset()
				if err := s.FromID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"from_id\"")
			}
		case "games":
			if err := func() error {
				s.Games.Reset()
				if err := s.Games.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"games\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode GamesResponse")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *GamesResponse) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *GamesResponse) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *Gender) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *Gender) encodeFields(e *jx.Encoder) {
	{
		if s.APIIntValue.Set {
			e.FieldStart("api_int_value")
			s.APIIntValue.Encode(e)
		}
	}
	{
		if s.APIStringValue.Set {
			e.FieldStart("api_string_value")
			s.APIStringValue.Encode(e)
		}
	}
}

var jsonFieldsNameOfGender = [2]string{
	0: "api_int_value",
	1: "api_string_value",
}

// Decode decodes Gender from json.
func (s *Gender) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Gender to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "api_int_value":
			if err := func() error {
				s.APIIntValue.Reset()
				if err := s.APIIntValue.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"api_int_value\"")
			}
		case "api_string_value":
			if err := func() error {
				s.APIStringValue.Reset()
				if err := s.APIStringValue.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"api_string_value\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode Gender")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *Gender) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *Gender) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *GenresResponse) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *GenresResponse) encodeFields(e *jx.Encoder) {
	{
		if s.Genres.Set {
			e.FieldStart("genres")
			s.Genres.Encode(e)
		}
	}
	{
		if s.NextPageValue.Set {
			e.FieldStart("next_page_value")
			s.NextPageValue.Encode(e)
		}
	}
}

var jsonFieldsNameOfGenresResponse = [2]string{
	0: "genres",
	1: "next_page_value",
}

// Decode decodes GenresResponse from json.
func (s *GenresResponse) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode GenresResponse to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "genres":
			if err := func() error {
				s.Genres.Reset()
				if err := s.Genres.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"genres\"")
			}
		case "next_page_value":
			if err := func() error {
				s.NextPageValue.Reset()
				if err := s.NextPageValue.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"next_page_value\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode GenresResponse")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *GenresResponse) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *GenresResponse) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s GetJoinedGroupStatusesOK) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields implements json.Marshaler.
func (s GetJoinedGroupStatusesOK) encodeFields(e *jx.Encoder) {
	for k, elem := range s {
		e.FieldStart(k)

		e.Str(elem)
	}
}

// Decode decodes GetJoinedGroupStatusesOK from json.
func (s *GetJoinedGroupStatusesOK) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode GetJoinedGroupStatusesOK to nil")
	}
	m := s.init()
	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		var elem string
		if err := func() error {
			v, err := d.Str()
			elem = string(v)
			if err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return errors.Wrapf(err, "decode field %q", k)
		}
		m[string(k)] = elem
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode GetJoinedGroupStatusesOK")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s GetJoinedGroupStatusesOK) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *GetJoinedGroupStatusesOK) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s GetJoinedThreadStatusesOK) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields implements json.Marshaler.
func (s GetJoinedThreadStatusesOK) encodeFields(e *jx.Encoder) {
	for k, elem := range s {
		e.FieldStart(k)

		e.Str(elem)
	}
}

// Decode decodes GetJoinedThreadStatusesOK from json.
func (s *GetJoinedThreadStatusesOK) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode GetJoinedThreadStatusesOK to nil")
	}
	m := s.init()
	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		var elem string
		if err := func() error {
			v, err := d.Str()
			elem = string(v)
			if err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return errors.Wrapf(err, "decode field %q", k)
		}
		m[string(k)] = elem
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode GetJoinedThreadStatusesOK")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s GetJoinedThreadStatusesOK) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *GetJoinedThreadStatusesOK) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *GifImage) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *GifImage) encodeFields(e *jx.Encoder) {
	{
		if s.Height.Set {
			e.FieldStart("height")
			s.Height.Encode(e)
		}
	}
	{
		if s.ID.Set {
			e.FieldStart("id")
			s.ID.Encode(e)
		}
	}
	{
		if s.URL.Set {
			e.FieldStart("url")
			s.URL.Encode(e)
		}
	}
	{
		if s.Width.Set {
			e.FieldStart("width")
			s.Width.Encode(e)
		}
	}
}

var jsonFieldsNameOfGifImage = [4]string{
	0: "height",
	1: "id",
	2: "url",
	3: "width",
}

// Decode decodes GifImage from json.
func (s *GifImage) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode GifImage to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "height":
			if err := func() error {
				s.Height.Reset()
				if err := s.Height.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"height\"")
			}
		case "id":
			if err := func() error {
				s.ID.Reset()
				if err := s.ID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"id\"")
			}
		case "url":
			if err := func() error {
				s.URL.Reset()
				if err := s.URL.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"url\"")
			}
		case "width":
			if err := func() error {
				s.Width.Reset()
				if err := s.Width.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"width\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode GifImage")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *GifImage) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *GifImage) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *Gift) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *Gift) encodeFields(e *jx.Encoder) {
	{
		if s.Currency.Set {
			e.FieldStart("currency")
			s.Currency.Encode(e)
		}
	}
	{
		if s.Icon.Set {
			e.FieldStart("icon")
			s.Icon.Encode(e)
		}
	}
	{
		if s.IconThumbnail.Set {
			e.FieldStart("icon_thumbnail")
			s.IconThumbnail.Encode(e)
		}
	}
	{
		if s.ID.Set {
			e.FieldStart("id")
			s.ID.Encode(e)
		}
	}
	{
		if s.Price.Set {
			e.FieldStart("price")
			s.Price.Encode(e)
		}
	}
	{
		if s.Title.Set {
			e.FieldStart("title")
			s.Title.Encode(e)
		}
	}
}

var jsonFieldsNameOfGift = [6]string{
	0: "currency",
	1: "icon",
	2: "icon_thumbnail",
	3: "id",
	4: "price",
	5: "title",
}

// Decode decodes Gift from json.
func (s *Gift) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Gift to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "currency":
			if err := func() error {
				s.Currency.Reset()
				if err := s.Currency.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"currency\"")
			}
		case "icon":
			if err := func() error {
				s.Icon.Reset()
				if err := s.Icon.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"icon\"")
			}
		case "icon_thumbnail":
			if err := func() error {
				s.IconThumbnail.Reset()
				if err := s.IconThumbnail.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"icon_thumbnail\"")
			}
		case "id":
			if err := func() error {
				s.ID.Reset()
				if err := s.ID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"id\"")
			}
		case "price":
			if err := func() error {
				s.Price.Reset()
				if err := s.Price.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"price\"")
			}
		case "title":
			if err := func() error {
				s.Title.Reset()
				if err := s.Title.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"title\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode Gift")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *Gift) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *Gift) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *GiftCount) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *GiftCount) encodeFields(e *jx.Encoder) {
	{
		if s.ID.Set {
			e.FieldStart("id")
			s.ID.Encode(e)
		}
	}
	{
		if s.Quantity.Set {
			e.FieldStart("quantity")
			s.Quantity.Encode(e)
		}
	}
}

var jsonFieldsNameOfGiftCount = [2]string{
	0: "id",
	1: "quantity",
}

// Decode decodes GiftCount from json.
func (s *GiftCount) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode GiftCount to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "id":
			if err := func() error {
				s.ID.Reset()
				if err := s.ID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"id\"")
			}
		case "quantity":
			if err := func() error {
				s.Quantity.Reset()
				if err := s.Quantity.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"quantity\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode GiftCount")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *GiftCount) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *GiftCount) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *GiftHistory) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *GiftHistory) encodeFields(e *jx.Encoder) {
	{
		if s.Anonymous.Set {
			e.FieldStart("anonymous")
			s.Anonymous.Encode(e)
		}
	}
	{
		if s.GiftItem.Set {
			e.FieldStart("gift_item")
			s.GiftItem.Encode(e)
		}
	}
	{
		if s.GiftsCount.Set {
			e.FieldStart("gifts_count")
			s.GiftsCount.Encode(e)
		}
	}
	{
		if s.Receiver != nil {
			e.FieldStart("receiver")
			s.Receiver.Encode(e)
		}
	}
	{
		if s.Sender != nil {
			e.FieldStart("sender")
			s.Sender.Encode(e)
		}
	}
	{
		if s.SentAt.Set {
			e.FieldStart("sent_at")
			s.SentAt.Encode(e)
		}
	}
}

var jsonFieldsNameOfGiftHistory = [6]string{
	0: "anonymous",
	1: "gift_item",
	2: "gifts_count",
	3: "receiver",
	4: "sender",
	5: "sent_at",
}

// Decode decodes GiftHistory from json.
func (s *GiftHistory) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode GiftHistory to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "anonymous":
			if err := func() error {
				s.Anonymous.Reset()
				if err := s.Anonymous.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"anonymous\"")
			}
		case "gift_item":
			if err := func() error {
				s.GiftItem.Reset()
				if err := s.GiftItem.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"gift_item\"")
			}
		case "gifts_count":
			if err := func() error {
				s.GiftsCount.Reset()
				if err := s.GiftsCount.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"gifts_count\"")
			}
		case "receiver":
			if err := func() error {
				s.Receiver = nil
				var elem RealmUser
				if err := elem.Decode(d); err != nil {
					return err
				}
				s.Receiver = &elem
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"receiver\"")
			}
		case "sender":
			if err := func() error {
				s.Sender = nil
				var elem RealmUser
				if err := elem.Decode(d); err != nil {
					return err
				}
				s.Sender = &elem
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"sender\"")
			}
		case "sent_at":
			if err := func() error {
				s.SentAt.Reset()
				if err := s.SentAt.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"sent_at\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode GiftHistory")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *GiftHistory) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *GiftHistory) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *GiftReceivedResponse) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *GiftReceivedResponse) encodeFields(e *jx.Encoder) {
	{
		if s.NextPageValue.Set {
			e.FieldStart("next_page_value")
			s.NextPageValue.Encode(e)
		}
	}
	{
		if s.ReceivedGifts.Set {
			e.FieldStart("received_gifts")
			s.ReceivedGifts.Encode(e)
		}
	}
	{
		if s.TotalCount.Set {
			e.FieldStart("total_count")
			s.TotalCount.Encode(e)
		}
	}
}

var jsonFieldsNameOfGiftReceivedResponse = [3]string{
	0: "next_page_value",
	1: "received_gifts",
	2: "total_count",
}

// Decode decodes GiftReceivedResponse from json.
func (s *GiftReceivedResponse) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode GiftReceivedResponse to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "next_page_value":
			if err := func() error {
				s.NextPageValue.Reset()
				if err := s.NextPageValue.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"next_page_value\"")
			}
		case "received_gifts":
			if err := func() error {
				s.ReceivedGifts.Reset()
				if err := s.ReceivedGifts.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"received_gifts\"")
			}
		case "total_count":
			if err := func() error {
				s.TotalCount.Reset()
				if err := s.TotalCount.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"total_count\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode GiftReceivedResponse")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *GiftReceivedResponse) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *GiftReceivedResponse) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *GiftReceivedTransactionResponse) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *GiftReceivedTransactionResponse) encodeFields(e *jx.Encoder) {
	{
		if s.CreatedAt.Set {
			e.FieldStart("created_at")
			s.CreatedAt.Encode(e)
		}
	}
	{
		if s.Gifts.Set {
			e.FieldStart("gifts")
			s.Gifts.Encode(e)
		}
	}
	{
		if s.Receiver != nil {
			e.FieldStart("receiver")
			s.Receiver.Encode(e)
		}
	}
}

var jsonFieldsNameOfGiftReceivedTransactionResponse = [3]string{
	0: "created_at",
	1: "gifts",
	2: "receiver",
}

// Decode decodes GiftReceivedTransactionResponse from json.
func (s *GiftReceivedTransactionResponse) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode GiftReceivedTransactionResponse to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "created_at":
			if err := func() error {
				s.CreatedAt.Reset()
				if err := s.CreatedAt.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"created_at\"")
			}
		case "gifts":
			if err := func() error {
				s.Gifts.Reset()
				if err := s.Gifts.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"gifts\"")
			}
		case "receiver":
			if err := func() error {
				s.Receiver = nil
				var elem RealmUser
				if err := elem.Decode(d); err != nil {
					return err
				}
				s.Receiver = &elem
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"receiver\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode GiftReceivedTransactionResponse")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *GiftReceivedTransactionResponse) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *GiftReceivedTransactionResponse) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *GiftSendersResponse) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *GiftSendersResponse) encodeFields(e *jx.Encoder) {
	{
		if s.Senders.Set {
			e.FieldStart("senders")
			s.Senders.Encode(e)
		}
	}
	{
		if s.TotalSendersCount.Set {
			e.FieldStart("total_senders_count")
			s.TotalSendersCount.Encode(e)
		}
	}
}

var jsonFieldsNameOfGiftSendersResponse = [2]string{
	0: "senders",
	1: "total_senders_count",
}

// Decode decodes GiftSendersResponse from json.
func (s *GiftSendersResponse) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode GiftSendersResponse to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "senders":
			if err := func() error {
				s.Senders.Reset()
				if err := s.Senders.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"senders\"")
			}
		case "total_senders_count":
			if err := func() error {
				s.TotalSendersCount.Reset()
				if err := s.TotalSendersCount.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"total_senders_count\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode GiftSendersResponse")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *GiftSendersResponse) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *GiftSendersResponse) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *GiftSlugItem) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *GiftSlugItem) encodeFields(e *jx.Encoder) {
	{
		if s.Icon.Set {
			e.FieldStart("icon")
			s.Icon.Encode(e)
		}
	}
	{
		if s.ID.Set {
			e.FieldStart("id")
			s.ID.Encode(e)
		}
	}
	{
		if s.Quantity.Set {
			e.FieldStart("quantity")
			s.Quantity.Encode(e)
		}
	}
	{
		if s.Slug.Set {
			e.FieldStart("slug")
			s.Slug.Encode(e)
		}
	}
}

var jsonFieldsNameOfGiftSlugItem = [4]string{
	0: "icon",
	1: "id",
	2: "quantity",
	3: "slug",
}

// Decode decodes GiftSlugItem from json.
func (s *GiftSlugItem) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode GiftSlugItem to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "icon":
			if err := func() error {
				s.Icon.Reset()
				if err := s.Icon.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"icon\"")
			}
		case "id":
			if err := func() error {
				s.ID.Reset()
				if err := s.ID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"id\"")
			}
		case "quantity":
			if err := func() error {
				s.Quantity.Reset()
				if err := s.Quantity.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"quantity\"")
			}
		case "slug":
			if err := func() error {
				s.Slug.Reset()
				if err := s.Slug.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"slug\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode GiftSlugItem")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *GiftSlugItem) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *GiftSlugItem) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *GiftTransactionsResponse) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *GiftTransactionsResponse) encodeFields(e *jx.Encoder) {
	{
		if s.HideGiftsReceived.Set {
			e.FieldStart("hide_gifts_received")
			s.HideGiftsReceived.Encode(e)
		}
	}
	{
		if s.NextPageValue.Set {
			e.FieldStart("next_page_value")
			s.NextPageValue.Encode(e)
		}
	}
	{
		if s.SentGifts.Set {
			e.FieldStart("sent_gifts")
			s.SentGifts.Encode(e)
		}
	}
}

var jsonFieldsNameOfGiftTransactionsResponse = [3]string{
	0: "hide_gifts_received",
	1: "next_page_value",
	2: "sent_gifts",
}

// Decode decodes GiftTransactionsResponse from json.
func (s *GiftTransactionsResponse) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode GiftTransactionsResponse to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "hide_gifts_received":
			if err := func() error {
				s.HideGiftsReceived.Reset()
				if err := s.HideGiftsReceived.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"hide_gifts_received\"")
			}
		case "next_page_value":
			if err := func() error {
				s.NextPageValue.Reset()
				if err := s.NextPageValue.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"next_page_value\"")
			}
		case "sent_gifts":
			if err := func() error {
				s.SentGifts.Reset()
				if err := s.SentGifts.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"sent_gifts\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode GiftTransactionsResponse")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *GiftTransactionsResponse) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *GiftTransactionsResponse) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *GiftingAbility) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *GiftingAbility) encodeFields(e *jx.Encoder) {
	{
		if s.CanReceive.Set {
			e.FieldStart("can_receive")
			s.CanReceive.Encode(e)
		}
	}
	{
		if s.CanSend.Set {
			e.FieldStart("can_send")
			s.CanSend.Encode(e)
		}
	}
	{
		if s.Enabled.Set {
			e.FieldStart("enabled")
			s.Enabled.Encode(e)
		}
	}
	{
		if s.UserID.Set {
			e.FieldStart("user_id")
			s.UserID.Encode(e)
		}
	}
}

var jsonFieldsNameOfGiftingAbility = [4]string{
	0: "can_receive",
	1: "can_send",
	2: "enabled",
	3: "user_id",
}

// Decode decodes GiftingAbility from json.
func (s *GiftingAbility) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode GiftingAbility to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "can_receive":
			if err := func() error {
				s.CanReceive.Reset()
				if err := s.CanReceive.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"can_receive\"")
			}
		case "can_send":
			if err := func() error {
				s.CanSend.Reset()
				if err := s.CanSend.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"can_send\"")
			}
		case "enabled":
			if err := func() error {
				s.Enabled.Reset()
				if err := s.Enabled.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"enabled\"")
			}
		case "user_id":
			if err := func() error {
				s.UserID.Reset()
				if err := s.UserID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"user_id\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode GiftingAbility")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *GiftingAbility) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *GiftingAbility) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *GiftsResponse) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *GiftsResponse) encodeFields(e *jx.Encoder) {
	{
		if s.Gifts.Set {
			e.FieldStart("gifts")
			s.Gifts.Encode(e)
		}
	}
	{
		if s.NextPageValue.Set {
			e.FieldStart("next_page_value")
			s.NextPageValue.Encode(e)
		}
	}
	{
		if s.TotalCount.Set {
			e.FieldStart("total_count")
			s.TotalCount.Encode(e)
		}
	}
}

var jsonFieldsNameOfGiftsResponse = [3]string{
	0: "gifts",
	1: "next_page_value",
	2: "total_count",
}

// Decode decodes GiftsResponse from json.
func (s *GiftsResponse) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode GiftsResponse to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "gifts":
			if err := func() error {
				s.Gifts.Reset()
				if err := s.Gifts.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"gifts\"")
			}
		case "next_page_value":
			if err := func() error {
				s.NextPageValue.Reset()
				if err := s.NextPageValue.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"next_page_value\"")
			}
		case "total_count":
			if err := func() error {
				s.TotalCount.Reset()
				if err := s.TotalCount.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"total_count\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode GiftsResponse")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *GiftsResponse) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *GiftsResponse) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *Group) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *Group) encodeFields(e *jx.Encoder) {
	{
		if s.AllowMembersToPostImageAndVideo.Set {
			e.FieldStart("allow_members_to_post_image_and_video")
			s.AllowMembersToPostImageAndVideo.Encode(e)
		}
	}
	{
		if s.AllowMembersToPostURL.Set {
			e.FieldStart("allow_members_to_post_url")
			s.AllowMembersToPostURL.Encode(e)
		}
	}
	{
		if s.AllowOwnershipTransfer.Set {
			e.FieldStart("allow_ownership_transfer")
			s.AllowOwnershipTransfer.Encode(e)
		}
	}
	{
		if s.AllowThreadCreationBy.Set {
			e.FieldStart("allow_thread_creation_by")
			s.AllowThreadCreationBy.Encode(e)
		}
	}
	{
		if s.CallTimelineDisplay.Set {
			e.FieldStart("call_timeline_display")
			s.CallTimelineDisplay.Encode(e)
		}
	}
	{
		if s.CoverImage.Set {
			e.FieldStart("cover_image")
			s.CoverImage.Encode(e)
		}
	}
	{
		if s.CoverImageThumbnail.Set {
			e.FieldStart("cover_image_thumbnail")
			s.CoverImageThumbnail.Encode(e)
		}
	}
	{
		if s.Description.Set {
			e.FieldStart("description")
			s.Description.Encode(e)
		}
	}
	{
		if s.Gender.Set {
			e.FieldStart("gender")
			s.Gender.Encode(e)
		}
	}
	{
		if s.GenerationGroupsLimit.Set {
			e.FieldStart("generation_groups_limit")
			s.GenerationGroupsLimit.Encode(e)
		}
	}
	{
		if s.GroupCategoryID.Set {
			e.FieldStart("group_category_id")
			s.GroupCategoryID.Encode(e)
		}
	}
	{
		if s.GroupIcon.Set {
			e.FieldStart("group_icon")
			s.GroupIcon.Encode(e)
		}
	}
	{
		if s.GroupIconThumbnail.Set {
			e.FieldStart("group_icon_thumbnail")
			s.GroupIconThumbnail.Encode(e)
		}
	}
	{
		if s.GroupsUsersCount.Set {
			e.FieldStart("groups_users_count")
			s.GroupsUsersCount.Encode(e)
		}
	}
	{
		if s.Guidelines.Set {
			e.FieldStart("guidelines")
			s.Guidelines.Encode(e)
		}
	}
	{
		if s.HideConferenceCall.Set {
			e.FieldStart("hide_conference_call")
			s.HideConferenceCall.Encode(e)
		}
	}
	{
		if s.HideFromGameEight.Set {
			e.FieldStart("hide_from_game_eight")
			s.HideFromGameEight.Encode(e)
		}
	}
	{
		if s.HideReportedPosts.Set {
			e.FieldStart("hide_reported_posts")
			s.HideReportedPosts.Encode(e)
		}
	}
	{
		if s.HighlightedCount.Set {
			e.FieldStart("highlighted_count")
			s.HighlightedCount.Encode(e)
		}
	}
	{
		if s.ID.Set {
			e.FieldStart("id")
			s.ID.Encode(e)
		}
	}
	{
		if s.InvitedToJoin.Set {
			e.FieldStart("invited_to_join")
			s.InvitedToJoin.Encode(e)
		}
	}
	{
		if s.IsJoined.Set {
			e.FieldStart("is_joined")
			s.IsJoined.Encode(e)
		}
	}
	{
		if s.IsPending.Set {
			e.FieldStart("is_pending")
			s.IsPending.Encode(e)
		}
	}
	{
		if s.IsPrivate.Set {
			e.FieldStart("is_private")
			s.IsPrivate.Encode(e)
		}
	}
	{
		if s.IsRelated.Set {
			e.FieldStart("is_related")
			s.IsRelated.Encode(e)
		}
	}
	{
		if s.JoinedCommunityCampaign.Set {
			e.FieldStart("joined_community_campaign")
			s.JoinedCommunityCampaign.Encode(e)
		}
	}
	{
		if s.ModeratorIds.Set {
			e.FieldStart("moderator_ids")
			s.ModeratorIds.Encode(e)
		}
	}
	{
		if s.OnlyMobileVerified.Set {
			e.FieldStart("only_mobile_verified")
			s.OnlyMobileVerified.Encode(e)
		}
	}
	{
		if s.OnlyVerifiedAge.Set {
			e.FieldStart("only_verified_age")
			s.OnlyVerifiedAge.Encode(e)
		}
	}
	{
		if s.Owner != nil {
			e.FieldStart("owner")
			s.Owner.Encode(e)
		}
	}
	{
		if s.PendingCount.Set {
			e.FieldStart("pending_count")
			s.PendingCount.Encode(e)
		}
	}
	{
		if s.PendingDeputizeIds.Set {
			e.FieldStart("pending_deputize_ids")
			s.PendingDeputizeIds.Encode(e)
		}
	}
	{
		if s.PendingTransferID.Set {
			e.FieldStart("pending_transfer_id")
			s.PendingTransferID.Encode(e)
		}
	}
	{
		if s.PostsCount.Set {
			e.FieldStart("posts_count")
			s.PostsCount.Encode(e)
		}
	}
	{
		if s.RelatedCount.Set {
			e.FieldStart("related_count")
			s.RelatedCount.Encode(e)
		}
	}
	{
		if s.SafeMode.Set {
			e.FieldStart("safe_mode")
			s.SafeMode.Encode(e)
		}
	}
	{
		if s.Secret.Set {
			e.FieldStart("secret")
			s.Secret.Encode(e)
		}
	}
	{
		if s.Seizable.Set {
			e.FieldStart("seizable")
			s.Seizable.Encode(e)
		}
	}
	{
		if s.SeizableBefore.Set {
			e.FieldStart("seizable_before")
			s.SeizableBefore.Encode(e)
		}
	}
	{
		if s.SubCategoryID.Set {
			e.FieldStart("sub_category_id")
			s.SubCategoryID.Encode(e)
		}
	}
	{
		if s.ThreadsCount.Set {
			e.FieldStart("threads_count")
			s.ThreadsCount.Encode(e)
		}
	}
	{
		if s.Title.Set {
			e.FieldStart("title")
			s.Title.Encode(e)
		}
	}
	{
		if s.Topic.Set {
			e.FieldStart("topic")
			s.Topic.Encode(e)
		}
	}
	{
		if s.UnreadCounts.Set {
			e.FieldStart("unread_counts")
			s.UnreadCounts.Encode(e)
		}
	}
	{
		if s.UnreadThreadsCount.Set {
			e.FieldStart("unread_threads_count")
			s.UnreadThreadsCount.Encode(e)
		}
	}
	{
		if s.UpdatedAt.Set {
			e.FieldStart("updated_at")
			s.UpdatedAt.Encode(e)
		}
	}
	{
		if s.UserID.Set {
			e.FieldStart("user_id")
			s.UserID.Encode(e)
		}
	}
	{
		if s.ViewsCount.Set {
			e.FieldStart("views_count")
			s.ViewsCount.Encode(e)
		}
	}
	{
		if s.WalkthroughRequested.Set {
			e.FieldStart("walkthrough_requested")
			s.WalkthroughRequested.Encode(e)
		}
	}
}

var jsonFieldsNameOfGroup = [49]string{
	0:  "allow_members_to_post_image_and_video",
	1:  "allow_members_to_post_url",
	2:  "allow_ownership_transfer",
	3:  "allow_thread_creation_by",
	4:  "call_timeline_display",
	5:  "cover_image",
	6:  "cover_image_thumbnail",
	7:  "description",
	8:  "gender",
	9:  "generation_groups_limit",
	10: "group_category_id",
	11: "group_icon",
	12: "group_icon_thumbnail",
	13: "groups_users_count",
	14: "guidelines",
	15: "hide_conference_call",
	16: "hide_from_game_eight",
	17: "hide_reported_posts",
	18: "highlighted_count",
	19: "id",
	20: "invited_to_join",
	21: "is_joined",
	22: "is_pending",
	23: "is_private",
	24: "is_related",
	25: "joined_community_campaign",
	26: "moderator_ids",
	27: "only_mobile_verified",
	28: "only_verified_age",
	29: "owner",
	30: "pending_count",
	31: "pending_deputize_ids",
	32: "pending_transfer_id",
	33: "posts_count",
	34: "related_count",
	35: "safe_mode",
	36: "secret",
	37: "seizable",
	38: "seizable_before",
	39: "sub_category_id",
	40: "threads_count",
	41: "title",
	42: "topic",
	43: "unread_counts",
	44: "unread_threads_count",
	45: "updated_at",
	46: "user_id",
	47: "views_count",
	48: "walkthrough_requested",
}

// Decode decodes Group from json.
func (s *Group) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Group to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "allow_members_to_post_image_and_video":
			if err := func() error {
				s.AllowMembersToPostImageAndVideo.Reset()
				if err := s.AllowMembersToPostImageAndVideo.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"allow_members_to_post_image_and_video\"")
			}
		case "allow_members_to_post_url":
			if err := func() error {
				s.AllowMembersToPostURL.Reset()
				if err := s.AllowMembersToPostURL.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"allow_members_to_post_url\"")
			}
		case "allow_ownership_transfer":
			if err := func() error {
				s.AllowOwnershipTransfer.Reset()
				if err := s.AllowOwnershipTransfer.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"allow_ownership_transfer\"")
			}
		case "allow_thread_creation_by":
			if err := func() error {
				s.AllowThreadCreationBy.Reset()
				if err := s.AllowThreadCreationBy.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"allow_thread_creation_by\"")
			}
		case "call_timeline_display":
			if err := func() error {
				s.CallTimelineDisplay.Reset()
				if err := s.CallTimelineDisplay.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"call_timeline_display\"")
			}
		case "cover_image":
			if err := func() error {
				s.CoverImage.Reset()
				if err := s.CoverImage.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"cover_image\"")
			}
		case "cover_image_thumbnail":
			if err := func() error {
				s.CoverImageThumbnail.Reset()
				if err := s.CoverImageThumbnail.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"cover_image_thumbnail\"")
			}
		case "description":
			if err := func() error {
				s.Description.Reset()
				if err := s.Description.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"description\"")
			}
		case "gender":
			if err := func() error {
				s.Gender.Reset()
				if err := s.Gender.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"gender\"")
			}
		case "generation_groups_limit":
			if err := func() error {
				s.GenerationGroupsLimit.Reset()
				if err := s.GenerationGroupsLimit.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"generation_groups_limit\"")
			}
		case "group_category_id":
			if err := func() error {
				s.GroupCategoryID.Reset()
				if err := s.GroupCategoryID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"group_category_id\"")
			}
		case "group_icon":
			if err := func() error {
				s.GroupIcon.Reset()
				if err := s.GroupIcon.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"group_icon\"")
			}
		case "group_icon_thumbnail":
			if err := func() error {
				s.GroupIconThumbnail.Reset()
				if err := s.GroupIconThumbnail.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"group_icon_thumbnail\"")
			}
		case "groups_users_count":
			if err := func() error {
				s.GroupsUsersCount.Reset()
				if err := s.GroupsUsersCount.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"groups_users_count\"")
			}
		case "guidelines":
			if err := func() error {
				s.Guidelines.Reset()
				if err := s.Guidelines.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"guidelines\"")
			}
		case "hide_conference_call":
			if err := func() error {
				s.HideConferenceCall.Reset()
				if err := s.HideConferenceCall.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"hide_conference_call\"")
			}
		case "hide_from_game_eight":
			if err := func() error {
				s.HideFromGameEight.Reset()
				if err := s.HideFromGameEight.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"hide_from_game_eight\"")
			}
		case "hide_reported_posts":
			if err := func() error {
				s.HideReportedPosts.Reset()
				if err := s.HideReportedPosts.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"hide_reported_posts\"")
			}
		case "highlighted_count":
			if err := func() error {
				s.HighlightedCount.Reset()
				if err := s.HighlightedCount.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"highlighted_count\"")
			}
		case "id":
			if err := func() error {
				s.ID.Reset()
				if err := s.ID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"id\"")
			}
		case "invited_to_join":
			if err := func() error {
				s.InvitedToJoin.Reset()
				if err := s.InvitedToJoin.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"invited_to_join\"")
			}
		case "is_joined":
			if err := func() error {
				s.IsJoined.Reset()
				if err := s.IsJoined.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"is_joined\"")
			}
		case "is_pending":
			if err := func() error {
				s.IsPending.Reset()
				if err := s.IsPending.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"is_pending\"")
			}
		case "is_private":
			if err := func() error {
				s.IsPrivate.Reset()
				if err := s.IsPrivate.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"is_private\"")
			}
		case "is_related":
			if err := func() error {
				s.IsRelated.Reset()
				if err := s.IsRelated.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"is_related\"")
			}
		case "joined_community_campaign":
			if err := func() error {
				s.JoinedCommunityCampaign.Reset()
				if err := s.JoinedCommunityCampaign.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"joined_community_campaign\"")
			}
		case "moderator_ids":
			if err := func() error {
				s.ModeratorIds.Reset()
				if err := s.ModeratorIds.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"moderator_ids\"")
			}
		case "only_mobile_verified":
			if err := func() error {
				s.OnlyMobileVerified.Reset()
				if err := s.OnlyMobileVerified.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"only_mobile_verified\"")
			}
		case "only_verified_age":
			if err := func() error {
				s.OnlyVerifiedAge.Reset()
				if err := s.OnlyVerifiedAge.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"only_verified_age\"")
			}
		case "owner":
			if err := func() error {
				s.Owner = nil
				var elem RealmUser
				if err := elem.Decode(d); err != nil {
					return err
				}
				s.Owner = &elem
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"owner\"")
			}
		case "pending_count":
			if err := func() error {
				s.PendingCount.Reset()
				if err := s.PendingCount.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"pending_count\"")
			}
		case "pending_deputize_ids":
			if err := func() error {
				s.PendingDeputizeIds.Reset()
				if err := s.PendingDeputizeIds.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"pending_deputize_ids\"")
			}
		case "pending_transfer_id":
			if err := func() error {
				s.PendingTransferID.Reset()
				if err := s.PendingTransferID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"pending_transfer_id\"")
			}
		case "posts_count":
			if err := func() error {
				s.PostsCount.Reset()
				if err := s.PostsCount.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"posts_count\"")
			}
		case "related_count":
			if err := func() error {
				s.RelatedCount.Reset()
				if err := s.RelatedCount.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"related_count\"")
			}
		case "safe_mode":
			if err := func() error {
				s.SafeMode.Reset()
				if err := s.SafeMode.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"safe_mode\"")
			}
		case "secret":
			if err := func() error {
				s.Secret.Reset()
				if err := s.Secret.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"secret\"")
			}
		case "seizable":
			if err := func() error {
				s.Seizable.Reset()
				if err := s.Seizable.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"seizable\"")
			}
		case "seizable_before":
			if err := func() error {
				s.SeizableBefore.Reset()
				if err := s.SeizableBefore.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"seizable_before\"")
			}
		case "sub_category_id":
			if err := func() error {
				s.SubCategoryID.Reset()
				if err := s.SubCategoryID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"sub_category_id\"")
			}
		case "threads_count":
			if err := func() error {
				s.ThreadsCount.Reset()
				if err := s.ThreadsCount.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"threads_count\"")
			}
		case "title":
			if err := func() error {
				s.Title.Reset()
				if err := s.Title.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"title\"")
			}
		case "topic":
			if err := func() error {
				s.Topic.Reset()
				if err := s.Topic.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"topic\"")
			}
		case "unread_counts":
			if err := func() error {
				s.UnreadCounts.Reset()
				if err := s.UnreadCounts.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"unread_counts\"")
			}
		case "unread_threads_count":
			if err := func() error {
				s.UnreadThreadsCount.Reset()
				if err := s.UnreadThreadsCount.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"unread_threads_count\"")
			}
		case "updated_at":
			if err := func() error {
				s.UpdatedAt.Reset()
				if err := s.UpdatedAt.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"updated_at\"")
			}
		case "user_id":
			if err := func() error {
				s.UserID.Reset()
				if err := s.UserID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"user_id\"")
			}
		case "views_count":
			if err := func() error {
				s.ViewsCount.Reset()
				if err := s.ViewsCount.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"views_count\"")
			}
		case "walkthrough_requested":
			if err := func() error {
				s.WalkthroughRequested.Reset()
				if err := s.WalkthroughRequested.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"walkthrough_requested\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode Group")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *Group) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *Group) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *GroupCategoriesResponse) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *GroupCategoriesResponse) encodeFields(e *jx.Encoder) {
	{
		if s.GroupCategories.Set {
			e.FieldStart("group_categories")
			s.GroupCategories.Encode(e)
		}
	}
}

var jsonFieldsNameOfGroupCategoriesResponse = [1]string{
	0: "group_categories",
}

// Decode decodes GroupCategoriesResponse from json.
func (s *GroupCategoriesResponse) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode GroupCategoriesResponse to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "group_categories":
			if err := func() error {
				s.GroupCategories.Reset()
				if err := s.GroupCategories.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"group_categories\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode GroupCategoriesResponse")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *GroupCategoriesResponse) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *GroupCategoriesResponse) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *GroupCategory) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *GroupCategory) encodeFields(e *jx.Encoder) {
	{
		if s.Icon.Set {
			e.FieldStart("icon")
			s.Icon.Encode(e)
		}
	}
	{
		if s.ID.Set {
			e.FieldStart("id")
			s.ID.Encode(e)
		}
	}
	{
		if s.Name.Set {
			e.FieldStart("name")
			s.Name.Encode(e)
		}
	}
	{
		if s.Rank.Set {
			e.FieldStart("rank")
			s.Rank.Encode(e)
		}
	}
}

var jsonFieldsNameOfGroupCategory = [4]string{
	0: "icon",
	1: "id",
	2: "name",
	3: "rank",
}

// Decode decodes GroupCategory from json.
func (s *GroupCategory) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode GroupCategory to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "icon":
			if err := func() error {
				s.Icon.Reset()
				if err := s.Icon.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"icon\"")
			}
		case "id":
			if err := func() error {
				s.ID.Reset()
				if err := s.ID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"id\"")
			}
		case "name":
			if err := func() error {
				s.Name.Reset()
				if err := s.Name.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"name\"")
			}
		case "rank":
			if err := func() error {
				s.Rank.Reset()
				if err := s.Rank.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"rank\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode GroupCategory")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *GroupCategory) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *GroupCategory) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *GroupGiftHistory) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *GroupGiftHistory) encodeFields(e *jx.Encoder) {
	{
		if s.GiftsCount.Set {
			e.FieldStart("gifts_count")
			s.GiftsCount.Encode(e)
		}
	}
	{
		if s.ReceivedDate.Set {
			e.FieldStart("received_date")
			s.ReceivedDate.Encode(e)
		}
	}
	{
		if s.User != nil {
			e.FieldStart("user")
			s.User.Encode(e)
		}
	}
}

var jsonFieldsNameOfGroupGiftHistory = [3]string{
	0: "gifts_count",
	1: "received_date",
	2: "user",
}

// Decode decodes GroupGiftHistory from json.
func (s *GroupGiftHistory) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode GroupGiftHistory to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "gifts_count":
			if err := func() error {
				s.GiftsCount.Reset()
				if err := s.GiftsCount.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"gifts_count\"")
			}
		case "received_date":
			if err := func() error {
				s.ReceivedDate.Reset()
				if err := s.ReceivedDate.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"received_date\"")
			}
		case "user":
			if err := func() error {
				s.User = nil
				var elem RealmUser
				if err := elem.Decode(d); err != nil {
					return err
				}
				s.User = &elem
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"user\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode GroupGiftHistory")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *GroupGiftHistory) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *GroupGiftHistory) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *GroupGiftHistoryResponse) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *GroupGiftHistoryResponse) encodeFields(e *jx.Encoder) {
	{
		if s.GiftHistory.Set {
			e.FieldStart("gift_history")
			s.GiftHistory.Encode(e)
		}
	}
	{
		if s.NextPageValue.Set {
			e.FieldStart("next_page_value")
			s.NextPageValue.Encode(e)
		}
	}
}

var jsonFieldsNameOfGroupGiftHistoryResponse = [2]string{
	0: "gift_history",
	1: "next_page_value",
}

// Decode decodes GroupGiftHistoryResponse from json.
func (s *GroupGiftHistoryResponse) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode GroupGiftHistoryResponse to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "gift_history":
			if err := func() error {
				s.GiftHistory.Reset()
				if err := s.GiftHistory.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"gift_history\"")
			}
		case "next_page_value":
			if err := func() error {
				s.NextPageValue.Reset()
				if err := s.NextPageValue.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"next_page_value\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode GroupGiftHistoryResponse")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *GroupGiftHistoryResponse) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *GroupGiftHistoryResponse) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *GroupMuteUsersResponse) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *GroupMuteUsersResponse) encodeFields(e *jx.Encoder) {
	{
		if s.GroupUsers.Set {
			e.FieldStart("group_users")
			s.GroupUsers.Encode(e)
		}
	}
	{
		if s.NextCursor.Set {
			e.FieldStart("next_cursor")
			s.NextCursor.Encode(e)
		}
	}
	{
		if s.Total.Set {
			e.FieldStart("total")
			s.Total.Encode(e)
		}
	}
}

var jsonFieldsNameOfGroupMuteUsersResponse = [3]string{
	0: "group_users",
	1: "next_cursor",
	2: "total",
}

// Decode decodes GroupMuteUsersResponse from json.
func (s *GroupMuteUsersResponse) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode GroupMuteUsersResponse to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "group_users":
			if err := func() error {
				s.GroupUsers.Reset()
				if err := s.GroupUsers.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"group_users\"")
			}
		case "next_cursor":
			if err := func() error {
				s.NextCursor.Reset()
				if err := s.NextCursor.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"next_cursor\"")
			}
		case "total":
			if err := func() error {
				s.Total.Reset()
				if err := s.Total.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"total\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode GroupMuteUsersResponse")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *GroupMuteUsersResponse) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *GroupMuteUsersResponse) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *GroupNotificationSettingsResponse) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *GroupNotificationSettingsResponse) encodeFields(e *jx.Encoder) {
	{
		if s.Setting.Set {
			e.FieldStart("setting")
			s.Setting.Encode(e)
		}
	}
}

var jsonFieldsNameOfGroupNotificationSettingsResponse = [1]string{
	0: "setting",
}

// Decode decodes GroupNotificationSettingsResponse from json.
func (s *GroupNotificationSettingsResponse) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode GroupNotificationSettingsResponse to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "setting":
			if err := func() error {
				s.Setting.Reset()
				if err := s.Setting.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"setting\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode GroupNotificationSettingsResponse")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *GroupNotificationSettingsResponse) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *GroupNotificationSettingsResponse) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *GroupResponse) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *GroupResponse) encodeFields(e *jx.Encoder) {
	{
		if s.Group.Set {
			e.FieldStart("group")
			s.Group.Encode(e)
		}
	}
}

var jsonFieldsNameOfGroupResponse = [1]string{
	0: "group",
}

// Decode decodes GroupResponse from json.
func (s *GroupResponse) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode GroupResponse to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "group":
			if err := func() error {
				s.Group.Reset()
				if err := s.Group.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"group\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode GroupResponse")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *GroupResponse) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *GroupResponse) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *GroupThreadListResponse) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *GroupThreadListResponse) encodeFields(e *jx.Encoder) {
	{
		if s.NextPageValue.Set {
			e.FieldStart("next_page_value")
			s.NextPageValue.Encode(e)
		}
	}
	{
		if s.Threads.Set {
			e.FieldStart("threads")
			s.Threads.Encode(e)
		}
	}
}

var jsonFieldsNameOfGroupThreadListResponse = [2]string{
	0: "next_page_value",
	1: "threads",
}

// Decode decodes GroupThreadListResponse from json.
func (s *GroupThreadListResponse) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode GroupThreadListResponse to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "next_page_value":
			if err := func() error {
				s.NextPageValue.Reset()
				if err := s.NextPageValue.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"next_page_value\"")
			}
		case "threads":
			if err := func() error {
				s.Threads.Reset()
				if err := s.Threads.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"threads\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode GroupThreadListResponse")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *GroupThreadListResponse) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *GroupThreadListResponse) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *GroupUser) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *GroupUser) encodeFields(e *jx.Encoder) {
	{
		if s.IsModerator.Set {
			e.FieldStart("is_moderator")
			s.IsModerator.Encode(e)
		}
	}
	{
		if s.IsMute.Set {
			e.FieldStart("is_mute")
			s.IsMute.Encode(e)
		}
	}
	{
		if s.PendingDeputize.Set {
			e.FieldStart("pending_deputize")
			s.PendingDeputize.Encode(e)
		}
	}
	{
		if s.PendingTransfer.Set {
			e.FieldStart("pending_transfer")
			s.PendingTransfer.Encode(e)
		}
	}
	{
		if s.Title.Set {
			e.FieldStart("title")
			s.Title.Encode(e)
		}
	}
	{
		if s.User != nil {
			e.FieldStart("user")
			s.User.Encode(e)
		}
	}
}

var jsonFieldsNameOfGroupUser = [6]string{
	0: "is_moderator",
	1: "is_mute",
	2: "pending_deputize",
	3: "pending_transfer",
	4: "title",
	5: "user",
}

// Decode decodes GroupUser from json.
func (s *GroupUser) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode GroupUser to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "is_moderator":
			if err := func() error {
				s.IsModerator.Reset()
				if err := s.IsModerator.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"is_moderator\"")
			}
		case "is_mute":
			if err := func() error {
				s.IsMute.Reset()
				if err := s.IsMute.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"is_mute\"")
			}
		case "pending_deputize":
			if err := func() error {
				s.PendingDeputize.Reset()
				if err := s.PendingDeputize.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"pending_deputize\"")
			}
		case "pending_transfer":
			if err := func() error {
				s.PendingTransfer.Reset()
				if err := s.PendingTransfer.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"pending_transfer\"")
			}
		case "title":
			if err := func() error {
				s.Title.Reset()
				if err := s.Title.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"title\"")
			}
		case "user":
			if err := func() error {
				s.User = nil
				var elem RealmUser
				if err := elem.Decode(d); err != nil {
					return err
				}
				s.User = &elem
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"user\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode GroupUser")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *GroupUser) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *GroupUser) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *GroupUserResponse) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *GroupUserResponse) encodeFields(e *jx.Encoder) {
	{
		if s.GroupUser.Set {
			e.FieldStart("group_user")
			s.GroupUser.Encode(e)
		}
	}
}

var jsonFieldsNameOfGroupUserResponse = [1]string{
	0: "group_user",
}

// Decode decodes GroupUserResponse from json.
func (s *GroupUserResponse) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode GroupUserResponse to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "group_user":
			if err := func() error {
				s.GroupUser.Reset()
				if err := s.GroupUser.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"group_user\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode GroupUserResponse")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *GroupUserResponse) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *GroupUserResponse) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *GroupUsersResponse) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *GroupUsersResponse) encodeFields(e *jx.Encoder) {
	{
		if s.GroupUsers.Set {
			e.FieldStart("group_users")
			s.GroupUsers.Encode(e)
		}
	}
}

var jsonFieldsNameOfGroupUsersResponse = [1]string{
	0: "group_users",
}

// Decode decodes GroupUsersResponse from json.
func (s *GroupUsersResponse) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode GroupUsersResponse to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "group_users":
			if err := func() error {
				s.GroupUsers.Reset()
				if err := s.GroupUsers.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"group_users\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode GroupUsersResponse")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *GroupUsersResponse) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *GroupUsersResponse) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *GroupsRelatedResponse) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *GroupsRelatedResponse) encodeFields(e *jx.Encoder) {
	{
		if s.Groups.Set {
			e.FieldStart("groups")
			s.Groups.Encode(e)
		}
	}
	{
		if s.NextPageValue.Set {
			e.FieldStart("next_page_value")
			s.NextPageValue.Encode(e)
		}
	}
}

var jsonFieldsNameOfGroupsRelatedResponse = [2]string{
	0: "groups",
	1: "next_page_value",
}

// Decode decodes GroupsRelatedResponse from json.
func (s *GroupsRelatedResponse) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode GroupsRelatedResponse to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "groups":
			if err := func() error {
				s.Groups.Reset()
				if err := s.Groups.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"groups\"")
			}
		case "next_page_value":
			if err := func() error {
				s.NextPageValue.Reset()
				if err := s.NextPageValue.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"next_page_value\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode GroupsRelatedResponse")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *GroupsRelatedResponse) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *GroupsRelatedResponse) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *GroupsResponse) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *GroupsResponse) encodeFields(e *jx.Encoder) {
	{
		if s.Groups.Set {
			e.FieldStart("groups")
			s.Groups.Encode(e)
		}
	}
	{
		if s.PinnedGroups.Set {
			e.FieldStart("pinned_groups")
			s.PinnedGroups.Encode(e)
		}
	}
}

var jsonFieldsNameOfGroupsResponse = [2]string{
	0: "groups",
	1: "pinned_groups",
}

// Decode decodes GroupsResponse from json.
func (s *GroupsResponse) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode GroupsResponse to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "groups":
			if err := func() error {
				s.Groups.Reset()
				if err := s.Groups.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"groups\"")
			}
		case "pinned_groups":
			if err := func() error {
				s.PinnedGroups.Reset()
				if err := s.PinnedGroups.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"pinned_groups\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode GroupsResponse")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *GroupsResponse) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *GroupsResponse) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *HiddenResponse) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *HiddenResponse) encodeFields(e *jx.Encoder) {
	{
		if s.HiddenUsers.Set {
			e.FieldStart("hidden_users")
			s.HiddenUsers.Encode(e)
		}
	}
	{
		if s.Limit.Set {
			e.FieldStart("limit")
			s.Limit.Encode(e)
		}
	}
	{
		if s.NextPageValue.Set {
			e.FieldStart("next_page_value")
			s.NextPageValue.Encode(e)
		}
	}
	{
		if s.TotalCount.Set {
			e.FieldStart("total_count")
			s.TotalCount.Encode(e)
		}
	}
}

var jsonFieldsNameOfHiddenResponse = [4]string{
	0: "hidden_users",
	1: "limit",
	2: "next_page_value",
	3: "total_count",
}

// Decode decodes HiddenResponse from json.
func (s *HiddenResponse) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode HiddenResponse to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "hidden_users":
			if err := func() error {
				s.HiddenUsers.Reset()
				if err := s.HiddenUsers.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"hidden_users\"")
			}
		case "limit":
			if err := func() error {
				s.Limit.Reset()
				if err := s.Limit.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"limit\"")
			}
		case "next_page_value":
			if err := func() error {
				s.NextPageValue.Reset()
				if err := s.NextPageValue.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"next_page_value\"")
			}
		case "total_count":
			if err := func() error {
				s.TotalCount.Reset()
				if err := s.TotalCount.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"total_count\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode HiddenResponse")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *HiddenResponse) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *HiddenResponse) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *HimaUsersResponse) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *HimaUsersResponse) encodeFields(e *jx.Encoder) {
	{
		if s.HimaUsers.Set {
			e.FieldStart("hima_users")
			s.HimaUsers.Encode(e)
		}
	}
}

var jsonFieldsNameOfHimaUsersResponse = [1]string{
	0: "hima_users",
}

// Decode decodes HimaUsersResponse from json.
func (s *HimaUsersResponse) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode HimaUsersResponse to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "hima_users":
			if err := func() error {
				s.HimaUsers.Reset()
				if err := s.HimaUsers.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"hima_users\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode HimaUsersResponse")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *HimaUsersResponse) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *HimaUsersResponse) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *Interest) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *Interest) encodeFields(e *jx.Encoder) {
	{
		if s.Icon.Set {
			e.FieldStart("icon")
			s.Icon.Encode(e)
		}
	}
	{
		if s.ID.Set {
			e.FieldStart("id")
			s.ID.Encode(e)
		}
	}
	{
		if s.Name.Set {
			e.FieldStart("name")
			s.Name.Encode(e)
		}
	}
	{
		if s.Selected.Set {
			e.FieldStart("selected")
			s.Selected.Encode(e)
		}
	}
}

var jsonFieldsNameOfInterest = [4]string{
	0: "icon",
	1: "id",
	2: "name",
	3: "selected",
}

// Decode decodes Interest from json.
func (s *Interest) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Interest to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "icon":
			if err := func() error {
				s.Icon.Reset()
				if err := s.Icon.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"icon\"")
			}
		case "id":
			if err := func() error {
				s.ID.Reset()
				if err := s.ID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"id\"")
			}
		case "name":
			if err := func() error {
				s.Name.Reset()
				if err := s.Name.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"name\"")
			}
		case "selected":
			if err := func() error {
				s.Selected.Reset()
				if err := s.Selected.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"selected\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode Interest")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *Interest) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *Interest) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *LikePostsResponse) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *LikePostsResponse) encodeFields(e *jx.Encoder) {
	{
		if s.LikeIds.Set {
			e.FieldStart("like_ids")
			s.LikeIds.Encode(e)
		}
	}
}

var jsonFieldsNameOfLikePostsResponse = [1]string{
	0: "like_ids",
}

// Decode decodes LikePostsResponse from json.
func (s *LikePostsResponse) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode LikePostsResponse to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "like_ids":
			if err := func() error {
				s.LikeIds.Reset()
				if err := s.LikeIds.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"like_ids\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode LikePostsResponse")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *LikePostsResponse) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *LikePostsResponse) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *LoginEmailUserRequest) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *LoginEmailUserRequest) encodeFields(e *jx.Encoder) {
	{
		if s.APIKey.Set {
			e.FieldStart("api_key")
			s.APIKey.Encode(e)
		}
	}
	{
		if s.Email.Set {
			e.FieldStart("email")
			s.Email.Encode(e)
		}
	}
	{
		if s.Password.Set {
			e.FieldStart("password")
			s.Password.Encode(e)
		}
	}
	{
		if s.TwoFACode.Set {
			e.FieldStart("two_f_a_code")
			s.TwoFACode.Encode(e)
		}
	}
	{
		if s.UUID.Set {
			e.FieldStart("uuid")
			s.UUID.Encode(e)
		}
	}
}

var jsonFieldsNameOfLoginEmailUserRequest = [5]string{
	0: "api_key",
	1: "email",
	2: "password",
	3: "two_f_a_code",
	4: "uuid",
}

// Decode decodes LoginEmailUserRequest from json.
func (s *LoginEmailUserRequest) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode LoginEmailUserRequest to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "api_key":
			if err := func() error {
				s.APIKey.Reset()
				if err := s.APIKey.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"api_key\"")
			}
		case "email":
			if err := func() error {
				s.Email.Reset()
				if err := s.Email.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"email\"")
			}
		case "password":
			if err := func() error {
				s.Password.Reset()
				if err := s.Password.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"password\"")
			}
		case "two_f_a_code":
			if err := func() error {
				s.TwoFACode.Reset()
				if err := s.TwoFACode.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"two_f_a_code\"")
			}
		case "uuid":
			if err := func() error {
				s.UUID.Reset()
				if err := s.UUID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"uuid\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode LoginEmailUserRequest")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *LoginEmailUserRequest) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *LoginEmailUserRequest) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *LoginUpdateResponse) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *LoginUpdateResponse) encodeFields(e *jx.Encoder) {
	{
		if s.AccessToken.Set {
			e.FieldStart("access_token")
			s.AccessToken.Encode(e)
		}
	}
	{
		if s.ExpiresIn.Set {
			e.FieldStart("expires_in")
			s.ExpiresIn.Encode(e)
		}
	}
	{
		if s.RefreshToken.Set {
			e.FieldStart("refresh_token")
			s.RefreshToken.Encode(e)
		}
	}
	{
		if s.UserID.Set {
			e.FieldStart("user_id")
			s.UserID.Encode(e)
		}
	}
	{
		if s.Username.Set {
			e.FieldStart("username")
			s.Username.Encode(e)
		}
	}
}

var jsonFieldsNameOfLoginUpdateResponse = [5]string{
	0: "access_token",
	1: "expires_in",
	2: "refresh_token",
	3: "user_id",
	4: "username",
}

// Decode decodes LoginUpdateResponse from json.
func (s *LoginUpdateResponse) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode LoginUpdateResponse to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "access_token":
			if err := func() error {
				s.AccessToken.Reset()
				if err := s.AccessToken.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"access_token\"")
			}
		case "expires_in":
			if err := func() error {
				s.ExpiresIn.Reset()
				if err := s.ExpiresIn.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"expires_in\"")
			}
		case "refresh_token":
			if err := func() error {
				s.RefreshToken.Reset()
				if err := s.RefreshToken.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"refresh_token\"")
			}
		case "user_id":
			if err := func() error {
				s.UserID.Reset()
				if err := s.UserID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"user_id\"")
			}
		case "username":
			if err := func() error {
				s.Username.Reset()
				if err := s.Username.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"username\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode LoginUpdateResponse")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *LoginUpdateResponse) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *LoginUpdateResponse) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *LoginUserResponse) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *LoginUserResponse) encodeFields(e *jx.Encoder) {
	{
		if s.AccessToken.Set {
			e.FieldStart("access_token")
			s.AccessToken.Encode(e)
		}
	}
	{
		if s.ExpiresIn.Set {
			e.FieldStart("expires_in")
			s.ExpiresIn.Encode(e)
		}
	}
	{
		if s.IsNew.Set {
			e.FieldStart("is_new")
			s.IsNew.Encode(e)
		}
	}
	{
		if s.RefreshToken.Set {
			e.FieldStart("refresh_token")
			s.RefreshToken.Encode(e)
		}
	}
	{
		if s.SnsInfo.Set {
			e.FieldStart("sns_info")
			s.SnsInfo.Encode(e)
		}
	}
	{
		if s.UserID.Set {
			e.FieldStart("user_id")
			s.UserID.Encode(e)
		}
	}
	{
		if s.Username.Set {
			e.FieldStart("username")
			s.Username.Encode(e)
		}
	}
}

var jsonFieldsNameOfLoginUserResponse = [7]string{
	0: "access_token",
	1: "expires_in",
	2: "is_new",
	3: "refresh_token",
	4: "sns_info",
	5: "user_id",
	6: "username",
}

// Decode decodes LoginUserResponse from json.
func (s *LoginUserResponse) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode LoginUserResponse to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "access_token":
			if err := func() error {
				s.AccessToken.Reset()
				if err := s.AccessToken.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"access_token\"")
			}
		case "expires_in":
			if err := func() error {
				s.ExpiresIn.Reset()
				if err := s.ExpiresIn.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"expires_in\"")
			}
		case "is_new":
			if err := func() error {
				s.IsNew.Reset()
				if err := s.IsNew.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"is_new\"")
			}
		case "refresh_token":
			if err := func() error {
				s.RefreshToken.Reset()
				if err := s.RefreshToken.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"refresh_token\"")
			}
		case "sns_info":
			if err := func() error {
				s.SnsInfo.Reset()
				if err := s.SnsInfo.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"sns_info\"")
			}
		case "user_id":
			if err := func() error {
				s.UserID.Reset()
				if err := s.UserID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"user_id\"")
			}
		case "username":
			if err := func() error {
				s.Username.Reset()
				if err := s.Username.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"username\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode LoginUserResponse")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *LoginUserResponse) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *LoginUserResponse) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *MessageResponse) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *MessageResponse) encodeFields(e *jx.Encoder) {
	{
		if s.ConferenceCall.Set {
			e.FieldStart("conference_call")
			s.ConferenceCall.Encode(e)
		}
	}
	{
		if s.ID.Set {
			e.FieldStart("id")
			s.ID.Encode(e)
		}
	}
}

var jsonFieldsNameOfMessageResponse = [2]string{
	0: "conference_call",
	1: "id",
}

// Decode decodes MessageResponse from json.
func (s *MessageResponse) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode MessageResponse to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "conference_call":
			if err := func() error {
				s.ConferenceCall.Reset()
				if err := s.ConferenceCall.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"conference_call\"")
			}
		case "id":
			if err := func() error {
				s.ID.Reset()
				if err := s.ID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"id\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode MessageResponse")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *MessageResponse) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *MessageResponse) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *MessageTag) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *MessageTag) encodeFields(e *jx.Encoder) {
	{
		if s.Length.Set {
			e.FieldStart("length")
			s.Length.Encode(e)
		}
	}
	{
		if s.Offset.Set {
			e.FieldStart("offset")
			s.Offset.Encode(e)
		}
	}
	{
		if s.Type.Set {
			e.FieldStart("type")
			s.Type.Encode(e)
		}
	}
	{
		if s.UserID.Set {
			e.FieldStart("user_id")
			s.UserID.Encode(e)
		}
	}
}

var jsonFieldsNameOfMessageTag = [4]string{
	0: "length",
	1: "offset",
	2: "type",
	3: "user_id",
}

// Decode decodes MessageTag from json.
func (s *MessageTag) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode MessageTag to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "length":
			if err := func() error {
				s.Length.Reset()
				if err := s.Length.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"length\"")
			}
		case "offset":
			if err := func() error {
				s.Offset.Reset()
				if err := s.Offset.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"offset\"")
			}
		case "type":
			if err := func() error {
				s.Type.Reset()
				if err := s.Type.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"type\"")
			}
		case "user_id":
			if err := func() error {
				s.UserID.Reset()
				if err := s.UserID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"user_id\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode MessageTag")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *MessageTag) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *MessageTag) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes MessageType as json.
func (s MessageType) Encode(e *jx.Encoder) {
	e.Str(string(s))
}

// Decode decodes MessageType from json.
func (s *MessageType) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode MessageType to nil")
	}
	v, err := d.StrBytes()
	if err != nil {
		return err
	}
	// Try to use constant string.
	switch MessageType(v) {
	case MessageTypeText:
		*s = MessageTypeText
	case MessageTypeImage:
		*s = MessageTypeImage
	case MessageTypeEternalImage:
		*s = MessageTypeEternalImage
	case MessageTypeVideo:
		*s = MessageTypeVideo
	case MessageTypeEternalVideo:
		*s = MessageTypeEternalVideo
	case MessageTypeScreenshotWarning:
		*s = MessageTypeScreenshotWarning
	case MessageTypeGIF:
		*s = MessageTypeGIF
	case MessageTypeSticker:
		*s = MessageTypeSticker
	case MessageTypeIndividualCall:
		*s = MessageTypeIndividualCall
	case MessageTypeDeleted:
		*s = MessageTypeDeleted
	case MessageTypeUserJoined:
		*s = MessageTypeUserJoined
	case MessageTypeUserLeave:
		*s = MessageTypeUserLeave
	case MessageTypeOwnerKick:
		*s = MessageTypeOwnerKick
	default:
		*s = MessageType(v)
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s MessageType) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *MessageType) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *MessagesResponse) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *MessagesResponse) encodeFields(e *jx.Encoder) {
	{
		if s.Messages.Set {
			e.FieldStart("messages")
			s.Messages.Encode(e)
		}
	}
}

var jsonFieldsNameOfMessagesResponse = [1]string{
	0: "messages",
}

// Decode decodes MessagesResponse from json.
func (s *MessagesResponse) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode MessagesResponse to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "messages":
			if err := func() error {
				s.Messages.Reset()
				if err := s.Messages.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"messages\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode MessagesResponse")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *MessagesResponse) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *MessagesResponse) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *Metadata) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *Metadata) encodeFields(e *jx.Encoder) {
	{
		if s.Anonymous.Set {
			e.FieldStart("anonymous")
			s.Anonymous.Encode(e)
		}
	}
	{
		if s.Body.Set {
			e.FieldStart("body")
			s.Body.Encode(e)
		}
	}
	{
		if s.BulkInvitation.Set {
			e.FieldStart("bulk_invitation")
			s.BulkInvitation.Encode(e)
		}
	}
	{
		if s.ContentPreview.Set {
			e.FieldStart("content_preview")
			s.ContentPreview.Encode(e)
		}
	}
	{
		if s.GroupTopic.Set {
			e.FieldStart("group_topic")
			s.GroupTopic.Encode(e)
		}
	}
	{
		if s.Title.Set {
			e.FieldStart("title")
			s.Title.Encode(e)
		}
	}
	{
		if s.URL.Set {
			e.FieldStart("url")
			s.URL.Encode(e)
		}
	}
}

var jsonFieldsNameOfMetadata = [7]string{
	0: "anonymous",
	1: "body",
	2: "bulk_invitation",
	3: "content_preview",
	4: "group_topic",
	5: "title",
	6: "url",
}

// Decode decodes Metadata from json.
func (s *Metadata) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Metadata to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "anonymous":
			if err := func() error {
				s.Anonymous.Reset()
				if err := s.Anonymous.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"anonymous\"")
			}
		case "body":
			if err := func() error {
				s.Body.Reset()
				if err := s.Body.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"body\"")
			}
		case "bulk_invitation":
			if err := func() error {
				s.BulkInvitation.Reset()
				if err := s.BulkInvitation.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"bulk_invitation\"")
			}
		case "content_preview":
			if err := func() error {
				s.ContentPreview.Reset()
				if err := s.ContentPreview.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"content_preview\"")
			}
		case "group_topic":
			if err := func() error {
				s.GroupTopic.Reset()
				if err := s.GroupTopic.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"group_topic\"")
			}
		case "title":
			if err := func() error {
				s.Title.Reset()
				if err := s.Title.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"title\"")
			}
		case "url":
			if err := func() error {
				s.URL.Reset()
				if err := s.URL.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"url\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode Metadata")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *Metadata) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *Metadata) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *MuteKeyword) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *MuteKeyword) encodeFields(e *jx.Encoder) {
	{
		if s.Context.Set {
			e.FieldStart("context")
			s.Context.Encode(e)
		}
	}
	{
		if s.ID.Set {
			e.FieldStart("id")
			s.ID.Encode(e)
		}
	}
	{
		if s.Word.Set {
			e.FieldStart("word")
			s.Word.Encode(e)
		}
	}
}

var jsonFieldsNameOfMuteKeyword = [3]string{
	0: "context",
	1: "id",
	2: "word",
}

// Decode decodes MuteKeyword from json.
func (s *MuteKeyword) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode MuteKeyword to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "context":
			if err := func() error {
				s.Context.Reset()
				if err := s.Context.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"context\"")
			}
		case "id":
			if err := func() error {
				s.ID.Reset()
				if err := s.ID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"id\"")
			}
		case "word":
			if err := func() error {
				s.Word.Reset()
				if err := s.Word.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"word\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode MuteKeyword")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *MuteKeyword) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *MuteKeyword) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *MuteKeywordRequest) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *MuteKeywordRequest) encodeFields(e *jx.Encoder) {
	{
		if s.Context.Set {
			e.FieldStart("context")
			s.Context.Encode(e)
		}
	}
	{
		if s.Word.Set {
			e.FieldStart("word")
			s.Word.Encode(e)
		}
	}
}

var jsonFieldsNameOfMuteKeywordRequest = [2]string{
	0: "context",
	1: "word",
}

// Decode decodes MuteKeywordRequest from json.
func (s *MuteKeywordRequest) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode MuteKeywordRequest to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "context":
			if err := func() error {
				s.Context.Reset()
				if err := s.Context.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"context\"")
			}
		case "word":
			if err := func() error {
				s.Word.Reset()
				if err := s.Word.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"word\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode MuteKeywordRequest")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *MuteKeywordRequest) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *MuteKeywordRequest) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *MuteKeywordResponse) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *MuteKeywordResponse) encodeFields(e *jx.Encoder) {
	{
		if s.HiddenWords.Set {
			e.FieldStart("hidden_words")
			s.HiddenWords.Encode(e)
		}
	}
}

var jsonFieldsNameOfMuteKeywordResponse = [1]string{
	0: "hidden_words",
}

// Decode decodes MuteKeywordResponse from json.
func (s *MuteKeywordResponse) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode MuteKeywordResponse to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "hidden_words":
			if err := func() error {
				s.HiddenWords.Reset()
				if err := s.HiddenWords.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"hidden_words\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode MuteKeywordResponse")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *MuteKeywordResponse) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *MuteKeywordResponse) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *NotificationSettingResponse) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *NotificationSettingResponse) encodeFields(e *jx.Encoder) {
	{
		if s.Setting.Set {
			e.FieldStart("setting")
			s.Setting.Encode(e)
		}
	}
}

var jsonFieldsNameOfNotificationSettingResponse = [1]string{
	0: "setting",
}

// Decode decodes NotificationSettingResponse from json.
func (s *NotificationSettingResponse) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode NotificationSettingResponse to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "setting":
			if err := func() error {
				s.Setting.Reset()
				if err := s.Setting.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"setting\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode NotificationSettingResponse")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *NotificationSettingResponse) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *NotificationSettingResponse) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *OnlineStatus) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *OnlineStatus) encodeFields(e *jx.Encoder) {
	{
		if s.APIValue.Set {
			e.FieldStart("api_value")
			s.APIValue.Encode(e)
		}
	}
}

var jsonFieldsNameOfOnlineStatus = [1]string{
	0: "api_value",
}

// Decode decodes OnlineStatus from json.
func (s *OnlineStatus) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode OnlineStatus to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "api_value":
			if err := func() error {
				s.APIValue.Reset()
				if err := s.APIValue.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"api_value\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode OnlineStatus")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *OnlineStatus) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OnlineStatus) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes OnlineStatusEnum as json.
func (s OnlineStatusEnum) Encode(e *jx.Encoder) {
	e.Str(string(s))
}

// Decode decodes OnlineStatusEnum from json.
func (s *OnlineStatusEnum) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode OnlineStatusEnum to nil")
	}
	v, err := d.StrBytes()
	if err != nil {
		return err
	}
	// Try to use constant string.
	switch OnlineStatusEnum(v) {
	case OnlineStatusEnumOffline:
		*s = OnlineStatusEnumOffline
	case OnlineStatusEnumHidden:
		*s = OnlineStatusEnumHidden
	default:
		*s = OnlineStatusEnum(v)
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OnlineStatusEnum) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OnlineStatusEnum) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes Application as json.
func (o OptApplication) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	o.Value.Encode(e)
}

// Decode decodes Application from json.
func (o *OptApplication) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptApplication to nil")
	}
	o.Set = true
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptApplication) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptApplication) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes ChatInvitation as json.
func (o OptChatInvitation) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	o.Value.Encode(e)
}

// Decode decodes ChatInvitation from json.
func (o *OptChatInvitation) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptChatInvitation to nil")
	}
	o.Set = true
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptChatInvitation) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptChatInvitation) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes ChatRoomLastMessage as json.
func (o OptChatRoomLastMessage) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	o.Value.Encode(e)
}

// Decode decodes ChatRoomLastMessage from json.
func (o *OptChatRoomLastMessage) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptChatRoomLastMessage to nil")
	}
	o.Set = true
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptChatRoomLastMessage) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptChatRoomLastMessage) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes ConferenceCallBumpParams as json.
func (o OptConferenceCallBumpParams) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	o.Value.Encode(e)
}

// Decode decodes ConferenceCallBumpParams from json.
func (o *OptConferenceCallBumpParams) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptConferenceCallBumpParams to nil")
	}
	o.Set = true
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptConferenceCallBumpParams) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptConferenceCallBumpParams) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes Config as json.
func (o OptConfig) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	o.Value.Encode(e)
}

// Decode decodes Config from json.
func (o *OptConfig) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptConfig to nil")
	}
	o.Set = true
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptConfig) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptConfig) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes *CreateConferenceCallPostReqMessageTags as json.
func (o OptCreateConferenceCallPostReqMessageTags) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	o.Value.Encode(e)
}

// Decode decodes *CreateConferenceCallPostReqMessageTags from json.
func (o *OptCreateConferenceCallPostReqMessageTags) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptCreateConferenceCallPostReqMessageTags to nil")
	}
	o.Set = true
	o.Value = new(CreateConferenceCallPostReqMessageTags)
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptCreateConferenceCallPostReqMessageTags) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptCreateConferenceCallPostReqMessageTags) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes CreateGroupQuota as json.
func (o OptCreateGroupQuota) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	o.Value.Encode(e)
}

// Decode decodes CreateGroupQuota from json.
func (o *OptCreateGroupQuota) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptCreateGroupQuota to nil")
	}
	o.Set = true
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptCreateGroupQuota) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptCreateGroupQuota) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes *CreatePostReqMessageTags as json.
func (o OptCreatePostReqMessageTags) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	o.Value.Encode(e)
}

// Decode decodes *CreatePostReqMessageTags from json.
func (o *OptCreatePostReqMessageTags) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptCreatePostReqMessageTags to nil")
	}
	o.Set = true
	o.Value = new(CreatePostReqMessageTags)
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptCreatePostReqMessageTags) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptCreatePostReqMessageTags) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes *CreatePostReqSharedURL as json.
func (o OptCreatePostReqSharedURL) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	o.Value.Encode(e)
}

// Decode decodes *CreatePostReqSharedURL from json.
func (o *OptCreatePostReqSharedURL) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptCreatePostReqSharedURL to nil")
	}
	o.Set = true
	o.Value = new(CreatePostReqSharedURL)
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptCreatePostReqSharedURL) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptCreatePostReqSharedURL) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes *CreateSharePostReqMessageTags as json.
func (o OptCreateSharePostReqMessageTags) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	o.Value.Encode(e)
}

// Decode decodes *CreateSharePostReqMessageTags from json.
func (o *OptCreateSharePostReqMessageTags) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptCreateSharePostReqMessageTags to nil")
	}
	o.Set = true
	o.Value = new(CreateSharePostReqMessageTags)
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptCreateSharePostReqMessageTags) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptCreateSharePostReqMessageTags) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes *CreateThreadPostReqMessageTags as json.
func (o OptCreateThreadPostReqMessageTags) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	o.Value.Encode(e)
}

// Decode decodes *CreateThreadPostReqMessageTags from json.
func (o *OptCreateThreadPostReqMessageTags) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptCreateThreadPostReqMessageTags to nil")
	}
	o.Set = true
	o.Value = new(CreateThreadPostReqMessageTags)
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptCreateThreadPostReqMessageTags) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptCreateThreadPostReqMessageTags) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes *CreateThreadPostReqSharedURL as json.
func (o OptCreateThreadPostReqSharedURL) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	o.Value.Encode(e)
}

// Decode decodes *CreateThreadPostReqSharedURL from json.
func (o *OptCreateThreadPostReqSharedURL) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptCreateThreadPostReqSharedURL to nil")
	}
	o.Set = true
	o.Value = new(CreateThreadPostReqSharedURL)
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptCreateThreadPostReqSharedURL) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptCreateThreadPostReqSharedURL) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes Gender as json.
func (o OptGender) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	o.Value.Encode(e)
}

// Decode decodes Gender from json.
func (o *OptGender) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptGender to nil")
	}
	o.Set = true
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptGender) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptGender) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes GifImage as json.
func (o OptGifImage) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	o.Value.Encode(e)
}

// Decode decodes GifImage from json.
func (o *OptGifImage) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptGifImage to nil")
	}
	o.Set = true
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptGifImage) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptGifImage) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes Gift as json.
func (o OptGift) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	o.Value.Encode(e)
}

// Decode decodes Gift from json.
func (o *OptGift) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptGift to nil")
	}
	o.Set = true
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptGift) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptGift) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes GiftSlugItem as json.
func (o OptGiftSlugItem) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	o.Value.Encode(e)
}

// Decode decodes GiftSlugItem from json.
func (o *OptGiftSlugItem) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptGiftSlugItem to nil")
	}
	o.Set = true
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptGiftSlugItem) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptGiftSlugItem) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes GiftingAbility as json.
func (o OptGiftingAbility) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	o.Value.Encode(e)
}

// Decode decodes GiftingAbility from json.
func (o *OptGiftingAbility) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptGiftingAbility to nil")
	}
	o.Set = true
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptGiftingAbility) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptGiftingAbility) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes Group as json.
func (o OptGroup) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	o.Value.Encode(e)
}

// Decode decodes Group from json.
func (o *OptGroup) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptGroup to nil")
	}
	o.Set = true
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptGroup) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptGroup) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes GroupUser as json.
func (o OptGroupUser) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	o.Value.Encode(e)
}

// Decode decodes GroupUser from json.
func (o *OptGroupUser) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptGroupUser to nil")
	}
	o.Set = true
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptGroupUser) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptGroupUser) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes MessageType as json.
func (o OptMessageType) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	e.Str(string(o.Value))
}

// Decode decodes MessageType from json.
func (o *OptMessageType) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptMessageType to nil")
	}
	o.Set = true
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptMessageType) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptMessageType) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes Metadata as json.
func (o OptMetadata) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	o.Value.Encode(e)
}

// Decode decodes Metadata from json.
func (o *OptMetadata) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptMetadata to nil")
	}
	o.Set = true
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptMetadata) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptMetadata) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes MuteKeyword as json.
func (o OptMuteKeyword) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	o.Value.Encode(e)
}

// Decode decodes MuteKeyword from json.
func (o *OptMuteKeyword) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptMuteKeyword to nil")
	}
	o.Set = true
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptMuteKeyword) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptMuteKeyword) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes []Activity as json.
func (o OptNilActivityArray) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	if o.Null {
		e.Null()
		return
	}
	e.ArrStart()
	for _, elem := range o.Value {
		elem.Encode(e)
	}
	e.ArrEnd()
}

// Decode decodes []Activity from json.
func (o *OptNilActivityArray) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptNilActivityArray to nil")
	}
	if d.Next() == jx.Null {
		if err := d.Null(); err != nil {
			return err
		}

		var v []Activity
		o.Value = v
		o.Set = true
		o.Null = true
		return nil
	}
	o.Set = true
	o.Null = false
	o.Value = make([]Activity, 0)
	if err := d.Arr(func(d *jx.Decoder) error {
		var elem Activity
		if err := elem.Decode(d); err != nil {
			return err
		}
		o.Value = append(o.Value, elem)
		return nil
	}); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptNilActivityArray) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptNilActivityArray) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes []BanWord as json.
func (o OptNilBanWordArray) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	if o.Null {
		e.Null()
		return
	}
	e.ArrStart()
	for _, elem := range o.Value {
		elem.Encode(e)
	}
	e.ArrEnd()
}

// Decode decodes []BanWord from json.
func (o *OptNilBanWordArray) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptNilBanWordArray to nil")
	}
	if d.Next() == jx.Null {
		if err := d.Null(); err != nil {
			return err
		}

		var v []BanWord
		o.Value = v
		o.Set = true
		o.Null = true
		return nil
	}
	o.Set = true
	o.Null = false
	o.Value = make([]BanWord, 0)
	if err := d.Arr(func(d *jx.Decoder) error {
		var elem BanWord
		if err := elem.Decode(d); err != nil {
			return err
		}
		o.Value = append(o.Value, elem)
		return nil
	}); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptNilBanWordArray) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptNilBanWordArray) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes []Bgm as json.
func (o OptNilBgmArray) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	if o.Null {
		e.Null()
		return
	}
	e.ArrStart()
	for _, elem := range o.Value {
		elem.Encode(e)
	}
	e.ArrEnd()
}

// Decode decodes []Bgm from json.
func (o *OptNilBgmArray) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptNilBgmArray to nil")
	}
	if d.Next() == jx.Null {
		if err := d.Null(); err != nil {
			return err
		}

		var v []Bgm
		o.Value = v
		o.Set = true
		o.Null = true
		return nil
	}
	o.Set = true
	o.Null = false
	o.Value = make([]Bgm, 0)
	if err := d.Arr(func(d *jx.Decoder) error {
		var elem Bgm
		if err := elem.Decode(d); err != nil {
			return err
		}
		o.Value = append(o.Value, elem)
		return nil
	}); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptNilBgmArray) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptNilBgmArray) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes bool as json.
func (o OptNilBool) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	if o.Null {
		e.Null()
		return
	}
	e.Bool(bool(o.Value))
}

// Decode decodes bool from json.
func (o *OptNilBool) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptNilBool to nil")
	}
	if d.Next() == jx.Null {
		if err := d.Null(); err != nil {
			return err
		}

		var v bool
		o.Value = v
		o.Set = true
		o.Null = true
		return nil
	}
	o.Set = true
	o.Null = false
	v, err := d.Bool()
	if err != nil {
		return err
	}
	o.Value = bool(v)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptNilBool) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptNilBool) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes []Choice as json.
func (o OptNilChoiceArray) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	if o.Null {
		e.Null()
		return
	}
	e.ArrStart()
	for _, elem := range o.Value {
		elem.Encode(e)
	}
	e.ArrEnd()
}

// Decode decodes []Choice from json.
func (o *OptNilChoiceArray) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptNilChoiceArray to nil")
	}
	if d.Next() == jx.Null {
		if err := d.Null(); err != nil {
			return err
		}

		var v []Choice
		o.Value = v
		o.Set = true
		o.Null = true
		return nil
	}
	o.Set = true
	o.Null = false
	o.Value = make([]Choice, 0)
	if err := d.Arr(func(d *jx.Decoder) error {
		var elem Choice
		if err := elem.Decode(d); err != nil {
			return err
		}
		o.Value = append(o.Value, elem)
		return nil
	}); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptNilChoiceArray) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptNilChoiceArray) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes []ConferenceCallUserRole as json.
func (o OptNilConferenceCallUserRoleArray) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	if o.Null {
		e.Null()
		return
	}
	e.ArrStart()
	for _, elem := range o.Value {
		elem.Encode(e)
	}
	e.ArrEnd()
}

// Decode decodes []ConferenceCallUserRole from json.
func (o *OptNilConferenceCallUserRoleArray) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptNilConferenceCallUserRoleArray to nil")
	}
	if d.Next() == jx.Null {
		if err := d.Null(); err != nil {
			return err
		}

		var v []ConferenceCallUserRole
		o.Value = v
		o.Set = true
		o.Null = true
		return nil
	}
	o.Set = true
	o.Null = false
	o.Value = make([]ConferenceCallUserRole, 0)
	if err := d.Arr(func(d *jx.Decoder) error {
		var elem ConferenceCallUserRole
		if err := elem.Decode(d); err != nil {
			return err
		}
		o.Value = append(o.Value, elem)
		return nil
	}); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptNilConferenceCallUserRoleArray) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptNilConferenceCallUserRoleArray) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes float64 as json.
func (o OptNilFloat64) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	if o.Null {
		e.Null()
		return
	}
	e.Float64(float64(o.Value))
}

// Decode decodes float64 from json.
func (o *OptNilFloat64) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptNilFloat64 to nil")
	}
	if d.Next() == jx.Null {
		if err := d.Null(); err != nil {
			return err
		}

		var v float64
		o.Value = v
		o.Set = true
		o.Null = true
		return nil
	}
	o.Set = true
	o.Null = false
	v, err := d.Float64()
	if err != nil {
		return err
	}
	o.Value = float64(v)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptNilFloat64) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptNilFloat64) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes []FootprintDTO as json.
func (o OptNilFootprintDTOArray) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	if o.Null {
		e.Null()
		return
	}
	e.ArrStart()
	for _, elem := range o.Value {
		elem.Encode(e)
	}
	e.ArrEnd()
}

// Decode decodes []FootprintDTO from json.
func (o *OptNilFootprintDTOArray) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptNilFootprintDTOArray to nil")
	}
	if d.Next() == jx.Null {
		if err := d.Null(); err != nil {
			return err
		}

		var v []FootprintDTO
		o.Value = v
		o.Set = true
		o.Null = true
		return nil
	}
	o.Set = true
	o.Null = false
	o.Value = make([]FootprintDTO, 0)
	if err := d.Arr(func(d *jx.Decoder) error {
		var elem FootprintDTO
		if err := elem.Decode(d); err != nil {
			return err
		}
		o.Value = append(o.Value, elem)
		return nil
	}); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptNilFootprintDTOArray) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptNilFootprintDTOArray) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes []GiftCount as json.
func (o OptNilGiftCountArray) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	if o.Null {
		e.Null()
		return
	}
	e.ArrStart()
	for _, elem := range o.Value {
		elem.Encode(e)
	}
	e.ArrEnd()
}

// Decode decodes []GiftCount from json.
func (o *OptNilGiftCountArray) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptNilGiftCountArray to nil")
	}
	if d.Next() == jx.Null {
		if err := d.Null(); err != nil {
			return err
		}

		var v []GiftCount
		o.Value = v
		o.Set = true
		o.Null = true
		return nil
	}
	o.Set = true
	o.Null = false
	o.Value = make([]GiftCount, 0)
	if err := d.Arr(func(d *jx.Decoder) error {
		var elem GiftCount
		if err := elem.Decode(d); err != nil {
			return err
		}
		o.Value = append(o.Value, elem)
		return nil
	}); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptNilGiftCountArray) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptNilGiftCountArray) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes []GiftHistory as json.
func (o OptNilGiftHistoryArray) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	if o.Null {
		e.Null()
		return
	}
	e.ArrStart()
	for _, elem := range o.Value {
		elem.Encode(e)
	}
	e.ArrEnd()
}

// Decode decodes []GiftHistory from json.
func (o *OptNilGiftHistoryArray) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptNilGiftHistoryArray to nil")
	}
	if d.Next() == jx.Null {
		if err := d.Null(); err != nil {
			return err
		}

		var v []GiftHistory
		o.Value = v
		o.Set = true
		o.Null = true
		return nil
	}
	o.Set = true
	o.Null = false
	o.Value = make([]GiftHistory, 0)
	if err := d.Arr(func(d *jx.Decoder) error {
		var elem GiftHistory
		if err := elem.Decode(d); err != nil {
			return err
		}
		o.Value = append(o.Value, elem)
		return nil
	}); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptNilGiftHistoryArray) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptNilGiftHistoryArray) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes []Group as json.
func (o OptNilGroupArray) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	if o.Null {
		e.Null()
		return
	}
	e.ArrStart()
	for _, elem := range o.Value {
		elem.Encode(e)
	}
	e.ArrEnd()
}

// Decode decodes []Group from json.
func (o *OptNilGroupArray) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptNilGroupArray to nil")
	}
	if d.Next() == jx.Null {
		if err := d.Null(); err != nil {
			return err
		}

		var v []Group
		o.Value = v
		o.Set = true
		o.Null = true
		return nil
	}
	o.Set = true
	o.Null = false
	o.Value = make([]Group, 0)
	if err := d.Arr(func(d *jx.Decoder) error {
		var elem Group
		if err := elem.Decode(d); err != nil {
			return err
		}
		o.Value = append(o.Value, elem)
		return nil
	}); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptNilGroupArray) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptNilGroupArray) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes []GroupCategory as json.
func (o OptNilGroupCategoryArray) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	if o.Null {
		e.Null()
		return
	}
	e.ArrStart()
	for _, elem := range o.Value {
		elem.Encode(e)
	}
	e.ArrEnd()
}

// Decode decodes []GroupCategory from json.
func (o *OptNilGroupCategoryArray) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptNilGroupCategoryArray to nil")
	}
	if d.Next() == jx.Null {
		if err := d.Null(); err != nil {
			return err
		}

		var v []GroupCategory
		o.Value = v
		o.Set = true
		o.Null = true
		return nil
	}
	o.Set = true
	o.Null = false
	o.Value = make([]GroupCategory, 0)
	if err := d.Arr(func(d *jx.Decoder) error {
		var elem GroupCategory
		if err := elem.Decode(d); err != nil {
			return err
		}
		o.Value = append(o.Value, elem)
		return nil
	}); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptNilGroupCategoryArray) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptNilGroupCategoryArray) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes []GroupGiftHistory as json.
func (o OptNilGroupGiftHistoryArray) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	if o.Null {
		e.Null()
		return
	}
	e.ArrStart()
	for _, elem := range o.Value {
		elem.Encode(e)
	}
	e.ArrEnd()
}

// Decode decodes []GroupGiftHistory from json.
func (o *OptNilGroupGiftHistoryArray) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptNilGroupGiftHistoryArray to nil")
	}
	if d.Next() == jx.Null {
		if err := d.Null(); err != nil {
			return err
		}

		var v []GroupGiftHistory
		o.Value = v
		o.Set = true
		o.Null = true
		return nil
	}
	o.Set = true
	o.Null = false
	o.Value = make([]GroupGiftHistory, 0)
	if err := d.Arr(func(d *jx.Decoder) error {
		var elem GroupGiftHistory
		if err := elem.Decode(d); err != nil {
			return err
		}
		o.Value = append(o.Value, elem)
		return nil
	}); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptNilGroupGiftHistoryArray) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptNilGroupGiftHistoryArray) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes []GroupUser as json.
func (o OptNilGroupUserArray) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	if o.Null {
		e.Null()
		return
	}
	e.ArrStart()
	for _, elem := range o.Value {
		elem.Encode(e)
	}
	e.ArrEnd()
}

// Decode decodes []GroupUser from json.
func (o *OptNilGroupUserArray) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptNilGroupUserArray to nil")
	}
	if d.Next() == jx.Null {
		if err := d.Null(); err != nil {
			return err
		}

		var v []GroupUser
		o.Value = v
		o.Set = true
		o.Null = true
		return nil
	}
	o.Set = true
	o.Null = false
	o.Value = make([]GroupUser, 0)
	if err := d.Arr(func(d *jx.Decoder) error {
		var elem GroupUser
		if err := elem.Decode(d); err != nil {
			return err
		}
		o.Value = append(o.Value, elem)
		return nil
	}); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptNilGroupUserArray) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptNilGroupUserArray) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes int32 as json.
func (o OptNilInt32) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	if o.Null {
		e.Null()
		return
	}
	e.Int32(int32(o.Value))
}

// Decode decodes int32 from json.
func (o *OptNilInt32) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptNilInt32 to nil")
	}
	if d.Next() == jx.Null {
		if err := d.Null(); err != nil {
			return err
		}

		var v int32
		o.Value = v
		o.Set = true
		o.Null = true
		return nil
	}
	o.Set = true
	o.Null = false
	v, err := d.Int32()
	if err != nil {
		return err
	}
	o.Value = int32(v)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptNilInt32) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptNilInt32) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes int64 as json.
func (o OptNilInt64) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	if o.Null {
		e.Null()
		return
	}
	e.Int64(int64(o.Value))
}

// Decode decodes int64 from json.
func (o *OptNilInt64) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptNilInt64 to nil")
	}
	if d.Next() == jx.Null {
		if err := d.Null(); err != nil {
			return err
		}

		var v int64
		o.Value = v
		o.Set = true
		o.Null = true
		return nil
	}
	o.Set = true
	o.Null = false
	v, err := d.Int64()
	if err != nil {
		return err
	}
	o.Value = int64(v)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptNilInt64) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptNilInt64) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes []int64 as json.
func (o OptNilInt64Array) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	if o.Null {
		e.Null()
		return
	}
	e.ArrStart()
	for _, elem := range o.Value {
		e.Int64(elem)
	}
	e.ArrEnd()
}

// Decode decodes []int64 from json.
func (o *OptNilInt64Array) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptNilInt64Array to nil")
	}
	if d.Next() == jx.Null {
		if err := d.Null(); err != nil {
			return err
		}

		var v []int64
		o.Value = v
		o.Set = true
		o.Null = true
		return nil
	}
	o.Set = true
	o.Null = false
	o.Value = make([]int64, 0)
	if err := d.Arr(func(d *jx.Decoder) error {
		var elem int64
		v, err := d.Int64()
		elem = int64(v)
		if err != nil {
			return err
		}
		o.Value = append(o.Value, elem)
		return nil
	}); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptNilInt64Array) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptNilInt64Array) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes []Interest as json.
func (o OptNilInterestArray) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	if o.Null {
		e.Null()
		return
	}
	e.ArrStart()
	for _, elem := range o.Value {
		elem.Encode(e)
	}
	e.ArrEnd()
}

// Decode decodes []Interest from json.
func (o *OptNilInterestArray) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptNilInterestArray to nil")
	}
	if d.Next() == jx.Null {
		if err := d.Null(); err != nil {
			return err
		}

		var v []Interest
		o.Value = v
		o.Set = true
		o.Null = true
		return nil
	}
	o.Set = true
	o.Null = false
	o.Value = make([]Interest, 0)
	if err := d.Arr(func(d *jx.Decoder) error {
		var elem Interest
		if err := elem.Decode(d); err != nil {
			return err
		}
		o.Value = append(o.Value, elem)
		return nil
	}); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptNilInterestArray) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptNilInterestArray) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes []MessageTag as json.
func (o OptNilMessageTagArray) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	if o.Null {
		e.Null()
		return
	}
	e.ArrStart()
	for _, elem := range o.Value {
		elem.Encode(e)
	}
	e.ArrEnd()
}

// Decode decodes []MessageTag from json.
func (o *OptNilMessageTagArray) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptNilMessageTagArray to nil")
	}
	if d.Next() == jx.Null {
		if err := d.Null(); err != nil {
			return err
		}

		var v []MessageTag
		o.Value = v
		o.Set = true
		o.Null = true
		return nil
	}
	o.Set = true
	o.Null = false
	o.Value = make([]MessageTag, 0)
	if err := d.Arr(func(d *jx.Decoder) error {
		var elem MessageTag
		if err := elem.Decode(d); err != nil {
			return err
		}
		o.Value = append(o.Value, elem)
		return nil
	}); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptNilMessageTagArray) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptNilMessageTagArray) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes []MuteKeyword as json.
func (o OptNilMuteKeywordArray) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	if o.Null {
		e.Null()
		return
	}
	e.ArrStart()
	for _, elem := range o.Value {
		elem.Encode(e)
	}
	e.ArrEnd()
}

// Decode decodes []MuteKeyword from json.
func (o *OptNilMuteKeywordArray) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptNilMuteKeywordArray to nil")
	}
	if d.Next() == jx.Null {
		if err := d.Null(); err != nil {
			return err
		}

		var v []MuteKeyword
		o.Value = v
		o.Set = true
		o.Null = true
		return nil
	}
	o.Set = true
	o.Null = false
	o.Value = make([]MuteKeyword, 0)
	if err := d.Arr(func(d *jx.Decoder) error {
		var elem MuteKeyword
		if err := elem.Decode(d); err != nil {
			return err
		}
		o.Value = append(o.Value, elem)
		return nil
	}); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptNilMuteKeywordArray) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptNilMuteKeywordArray) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes []PopularWord as json.
func (o OptNilPopularWordArray) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	if o.Null {
		e.Null()
		return
	}
	e.ArrStart()
	for _, elem := range o.Value {
		elem.Encode(e)
	}
	e.ArrEnd()
}

// Decode decodes []PopularWord from json.
func (o *OptNilPopularWordArray) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptNilPopularWordArray to nil")
	}
	if d.Next() == jx.Null {
		if err := d.Null(); err != nil {
			return err
		}

		var v []PopularWord
		o.Value = v
		o.Set = true
		o.Null = true
		return nil
	}
	o.Set = true
	o.Null = false
	o.Value = make([]PopularWord, 0)
	if err := d.Arr(func(d *jx.Decoder) error {
		var elem PopularWord
		if err := elem.Decode(d); err != nil {
			return err
		}
		o.Value = append(o.Value, elem)
		return nil
	}); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptNilPopularWordArray) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptNilPopularWordArray) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes []Post as json.
func (o OptNilPostArray) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	if o.Null {
		e.Null()
		return
	}
	e.ArrStart()
	for _, elem := range o.Value {
		elem.Encode(e)
	}
	e.ArrEnd()
}

// Decode decodes []Post from json.
func (o *OptNilPostArray) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptNilPostArray to nil")
	}
	if d.Next() == jx.Null {
		if err := d.Null(); err != nil {
			return err
		}

		var v []Post
		o.Value = v
		o.Set = true
		o.Null = true
		return nil
	}
	o.Set = true
	o.Null = false
	o.Value = make([]Post, 0)
	if err := d.Arr(func(d *jx.Decoder) error {
		var elem Post
		if err := elem.Decode(d); err != nil {
			return err
		}
		o.Value = append(o.Value, elem)
		return nil
	}); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptNilPostArray) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptNilPostArray) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes []PostTag as json.
func (o OptNilPostTagArray) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	if o.Null {
		e.Null()
		return
	}
	e.ArrStart()
	for _, elem := range o.Value {
		elem.Encode(e)
	}
	e.ArrEnd()
}

// Decode decodes []PostTag from json.
func (o *OptNilPostTagArray) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptNilPostTagArray to nil")
	}
	if d.Next() == jx.Null {
		if err := d.Null(); err != nil {
			return err
		}

		var v []PostTag
		o.Value = v
		o.Set = true
		o.Null = true
		return nil
	}
	o.Set = true
	o.Null = false
	o.Value = make([]PostTag, 0)
	if err := d.Arr(func(d *jx.Decoder) error {
		var elem PostTag
		if err := elem.Decode(d); err != nil {
			return err
		}
		o.Value = append(o.Value, elem)
		return nil
	}); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptNilPostTagArray) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptNilPostTagArray) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes []PresignedUrl as json.
func (o OptNilPresignedUrlArray) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	if o.Null {
		e.Null()
		return
	}
	e.ArrStart()
	for _, elem := range o.Value {
		elem.Encode(e)
	}
	e.ArrEnd()
}

// Decode decodes []PresignedUrl from json.
func (o *OptNilPresignedUrlArray) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptNilPresignedUrlArray to nil")
	}
	if d.Next() == jx.Null {
		if err := d.Null(); err != nil {
			return err
		}

		var v []PresignedUrl
		o.Value = v
		o.Set = true
		o.Null = true
		return nil
	}
	o.Set = true
	o.Null = false
	o.Value = make([]PresignedUrl, 0)
	if err := d.Arr(func(d *jx.Decoder) error {
		var elem PresignedUrl
		if err := elem.Decode(d); err != nil {
			return err
		}
		o.Value = append(o.Value, elem)
		return nil
	}); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptNilPresignedUrlArray) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptNilPresignedUrlArray) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes []RealmChatRoom as json.
func (o OptNilRealmChatRoomArray) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	if o.Null {
		e.Null()
		return
	}
	e.ArrStart()
	for _, elem := range o.Value {
		elem.Encode(e)
	}
	e.ArrEnd()
}

// Decode decodes []RealmChatRoom from json.
func (o *OptNilRealmChatRoomArray) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptNilRealmChatRoomArray to nil")
	}
	if d.Next() == jx.Null {
		if err := d.Null(); err != nil {
			return err
		}

		var v []RealmChatRoom
		o.Value = v
		o.Set = true
		o.Null = true
		return nil
	}
	o.Set = true
	o.Null = false
	o.Value = make([]RealmChatRoom, 0)
	if err := d.Arr(func(d *jx.Decoder) error {
		var elem RealmChatRoom
		if err := elem.Decode(d); err != nil {
			return err
		}
		o.Value = append(o.Value, elem)
		return nil
	}); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptNilRealmChatRoomArray) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptNilRealmChatRoomArray) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes []RealmGame as json.
func (o OptNilRealmGameArray) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	if o.Null {
		e.Null()
		return
	}
	e.ArrStart()
	for _, elem := range o.Value {
		elem.Encode(e)
	}
	e.ArrEnd()
}

// Decode decodes []RealmGame from json.
func (o *OptNilRealmGameArray) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptNilRealmGameArray to nil")
	}
	if d.Next() == jx.Null {
		if err := d.Null(); err != nil {
			return err
		}

		var v []RealmGame
		o.Value = v
		o.Set = true
		o.Null = true
		return nil
	}
	o.Set = true
	o.Null = false
	o.Value = make([]RealmGame, 0)
	if err := d.Arr(func(d *jx.Decoder) error {
		var elem RealmGame
		if err := elem.Decode(d); err != nil {
			return err
		}
		o.Value = append(o.Value, elem)
		return nil
	}); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptNilRealmGameArray) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptNilRealmGameArray) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes []RealmGenre as json.
func (o OptNilRealmGenreArray) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	if o.Null {
		e.Null()
		return
	}
	e.ArrStart()
	for _, elem := range o.Value {
		elem.Encode(e)
	}
	e.ArrEnd()
}

// Decode decodes []RealmGenre from json.
func (o *OptNilRealmGenreArray) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptNilRealmGenreArray to nil")
	}
	if d.Next() == jx.Null {
		if err := d.Null(); err != nil {
			return err
		}

		var v []RealmGenre
		o.Value = v
		o.Set = true
		o.Null = true
		return nil
	}
	o.Set = true
	o.Null = false
	o.Value = make([]RealmGenre, 0)
	if err := d.Arr(func(d *jx.Decoder) error {
		var elem RealmGenre
		if err := elem.Decode(d); err != nil {
			return err
		}
		o.Value = append(o.Value, elem)
		return nil
	}); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptNilRealmGenreArray) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptNilRealmGenreArray) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes []RealmGift as json.
func (o OptNilRealmGiftArray) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	if o.Null {
		e.Null()
		return
	}
	e.ArrStart()
	for _, elem := range o.Value {
		elem.Encode(e)
	}
	e.ArrEnd()
}

// Decode decodes []RealmGift from json.
func (o *OptNilRealmGiftArray) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptNilRealmGiftArray to nil")
	}
	if d.Next() == jx.Null {
		if err := d.Null(); err != nil {
			return err
		}

		var v []RealmGift
		o.Value = v
		o.Set = true
		o.Null = true
		return nil
	}
	o.Set = true
	o.Null = false
	o.Value = make([]RealmGift, 0)
	if err := d.Arr(func(d *jx.Decoder) error {
		var elem RealmGift
		if err := elem.Decode(d); err != nil {
			return err
		}
		o.Value = append(o.Value, elem)
		return nil
	}); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptNilRealmGiftArray) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptNilRealmGiftArray) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes []RealmMessage as json.
func (o OptNilRealmMessageArray) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	if o.Null {
		e.Null()
		return
	}
	e.ArrStart()
	for _, elem := range o.Value {
		elem.Encode(e)
	}
	e.ArrEnd()
}

// Decode decodes []RealmMessage from json.
func (o *OptNilRealmMessageArray) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptNilRealmMessageArray to nil")
	}
	if d.Next() == jx.Null {
		if err := d.Null(); err != nil {
			return err
		}

		var v []RealmMessage
		o.Value = v
		o.Set = true
		o.Null = true
		return nil
	}
	o.Set = true
	o.Null = false
	o.Value = make([]RealmMessage, 0)
	if err := d.Arr(func(d *jx.Decoder) error {
		var elem RealmMessage
		if err := elem.Decode(d); err != nil {
			return err
		}
		o.Value = append(o.Value, elem)
		return nil
	}); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptNilRealmMessageArray) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptNilRealmMessageArray) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes []RealmUser as json.
func (o OptNilRealmUserArray) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	if o.Null {
		e.Null()
		return
	}
	e.ArrStart()
	for _, elem := range o.Value {
		elem.Encode(e)
	}
	e.ArrEnd()
}

// Decode decodes []RealmUser from json.
func (o *OptNilRealmUserArray) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptNilRealmUserArray to nil")
	}
	if d.Next() == jx.Null {
		if err := d.Null(); err != nil {
			return err
		}

		var v []RealmUser
		o.Value = v
		o.Set = true
		o.Null = true
		return nil
	}
	o.Set = true
	o.Null = false
	o.Value = make([]RealmUser, 0)
	if err := d.Arr(func(d *jx.Decoder) error {
		var elem RealmUser
		if err := elem.Decode(d); err != nil {
			return err
		}
		o.Value = append(o.Value, elem)
		return nil
	}); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptNilRealmUserArray) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptNilRealmUserArray) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes []ReceivedGift as json.
func (o OptNilReceivedGiftArray) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	if o.Null {
		e.Null()
		return
	}
	e.ArrStart()
	for _, elem := range o.Value {
		elem.Encode(e)
	}
	e.ArrEnd()
}

// Decode decodes []ReceivedGift from json.
func (o *OptNilReceivedGiftArray) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptNilReceivedGiftArray to nil")
	}
	if d.Next() == jx.Null {
		if err := d.Null(); err != nil {
			return err
		}

		var v []ReceivedGift
		o.Value = v
		o.Set = true
		o.Null = true
		return nil
	}
	o.Set = true
	o.Null = false
	o.Value = make([]ReceivedGift, 0)
	if err := d.Arr(func(d *jx.Decoder) error {
		var elem ReceivedGift
		if err := elem.Decode(d); err != nil {
			return err
		}
		o.Value = append(o.Value, elem)
		return nil
	}); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptNilReceivedGiftArray) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptNilReceivedGiftArray) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes []RefreshCounterRequest as json.
func (o OptNilRefreshCounterRequestArray) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	if o.Null {
		e.Null()
		return
	}
	e.ArrStart()
	for _, elem := range o.Value {
		elem.Encode(e)
	}
	e.ArrEnd()
}

// Decode decodes []RefreshCounterRequest from json.
func (o *OptNilRefreshCounterRequestArray) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptNilRefreshCounterRequestArray to nil")
	}
	if d.Next() == jx.Null {
		if err := d.Null(); err != nil {
			return err
		}

		var v []RefreshCounterRequest
		o.Value = v
		o.Set = true
		o.Null = true
		return nil
	}
	o.Set = true
	o.Null = false
	o.Value = make([]RefreshCounterRequest, 0)
	if err := d.Arr(func(d *jx.Decoder) error {
		var elem RefreshCounterRequest
		if err := elem.Decode(d); err != nil {
			return err
		}
		o.Value = append(o.Value, elem)
		return nil
	}); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptNilRefreshCounterRequestArray) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptNilRefreshCounterRequestArray) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes []Review as json.
func (o OptNilReviewArray) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	if o.Null {
		e.Null()
		return
	}
	e.ArrStart()
	for _, elem := range o.Value {
		elem.Encode(e)
	}
	e.ArrEnd()
}

// Decode decodes []Review from json.
func (o *OptNilReviewArray) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptNilReviewArray to nil")
	}
	if d.Next() == jx.Null {
		if err := d.Null(); err != nil {
			return err
		}

		var v []Review
		o.Value = v
		o.Set = true
		o.Null = true
		return nil
	}
	o.Set = true
	o.Null = false
	o.Value = make([]Review, 0)
	if err := d.Arr(func(d *jx.Decoder) error {
		var elem Review
		if err := elem.Decode(d); err != nil {
			return err
		}
		o.Value = append(o.Value, elem)
		return nil
	}); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptNilReviewArray) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptNilReviewArray) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes []Sticker as json.
func (o OptNilStickerArray) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	if o.Null {
		e.Null()
		return
	}
	e.ArrStart()
	for _, elem := range o.Value {
		elem.Encode(e)
	}
	e.ArrEnd()
}

// Decode decodes []Sticker from json.
func (o *OptNilStickerArray) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptNilStickerArray to nil")
	}
	if d.Next() == jx.Null {
		if err := d.Null(); err != nil {
			return err
		}

		var v []Sticker
		o.Value = v
		o.Set = true
		o.Null = true
		return nil
	}
	o.Set = true
	o.Null = false
	o.Value = make([]Sticker, 0)
	if err := d.Arr(func(d *jx.Decoder) error {
		var elem Sticker
		if err := elem.Decode(d); err != nil {
			return err
		}
		o.Value = append(o.Value, elem)
		return nil
	}); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptNilStickerArray) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptNilStickerArray) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes []StickerPack as json.
func (o OptNilStickerPackArray) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	if o.Null {
		e.Null()
		return
	}
	e.ArrStart()
	for _, elem := range o.Value {
		elem.Encode(e)
	}
	e.ArrEnd()
}

// Decode decodes []StickerPack from json.
func (o *OptNilStickerPackArray) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptNilStickerPackArray to nil")
	}
	if d.Next() == jx.Null {
		if err := d.Null(); err != nil {
			return err
		}

		var v []StickerPack
		o.Value = v
		o.Set = true
		o.Null = true
		return nil
	}
	o.Set = true
	o.Null = false
	o.Value = make([]StickerPack, 0)
	if err := d.Arr(func(d *jx.Decoder) error {
		var elem StickerPack
		if err := elem.Decode(d); err != nil {
			return err
		}
		o.Value = append(o.Value, elem)
		return nil
	}); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptNilStickerPackArray) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptNilStickerPackArray) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes string as json.
func (o OptNilString) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	if o.Null {
		e.Null()
		return
	}
	e.Str(string(o.Value))
}

// Decode decodes string from json.
func (o *OptNilString) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptNilString to nil")
	}
	if d.Next() == jx.Null {
		if err := d.Null(); err != nil {
			return err
		}

		var v string
		o.Value = v
		o.Set = true
		o.Null = true
		return nil
	}
	o.Set = true
	o.Null = false
	v, err := d.Str()
	if err != nil {
		return err
	}
	o.Value = string(v)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptNilString) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptNilString) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes []string as json.
func (o OptNilStringArray) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	if o.Null {
		e.Null()
		return
	}
	e.ArrStart()
	for _, elem := range o.Value {
		e.Str(elem)
	}
	e.ArrEnd()
}

// Decode decodes []string from json.
func (o *OptNilStringArray) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptNilStringArray to nil")
	}
	if d.Next() == jx.Null {
		if err := d.Null(); err != nil {
			return err
		}

		var v []string
		o.Value = v
		o.Set = true
		o.Null = true
		return nil
	}
	o.Set = true
	o.Null = false
	o.Value = make([]string, 0)
	if err := d.Arr(func(d *jx.Decoder) error {
		var elem string
		v, err := d.Str()
		elem = string(v)
		if err != nil {
			return err
		}
		o.Value = append(o.Value, elem)
		return nil
	}); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptNilStringArray) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptNilStringArray) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes []ThreadInfo as json.
func (o OptNilThreadInfoArray) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	if o.Null {
		e.Null()
		return
	}
	e.ArrStart()
	for _, elem := range o.Value {
		elem.Encode(e)
	}
	e.ArrEnd()
}

// Decode decodes []ThreadInfo from json.
func (o *OptNilThreadInfoArray) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptNilThreadInfoArray to nil")
	}
	if d.Next() == jx.Null {
		if err := d.Null(); err != nil {
			return err
		}

		var v []ThreadInfo
		o.Value = v
		o.Set = true
		o.Null = true
		return nil
	}
	o.Set = true
	o.Null = false
	o.Value = make([]ThreadInfo, 0)
	if err := d.Arr(func(d *jx.Decoder) error {
		var elem ThreadInfo
		if err := elem.Decode(d); err != nil {
			return err
		}
		o.Value = append(o.Value, elem)
		return nil
	}); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptNilThreadInfoArray) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptNilThreadInfoArray) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes []TransactionGiftReceived as json.
func (o OptNilTransactionGiftReceivedArray) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	if o.Null {
		e.Null()
		return
	}
	e.ArrStart()
	for _, elem := range o.Value {
		elem.Encode(e)
	}
	e.ArrEnd()
}

// Decode decodes []TransactionGiftReceived from json.
func (o *OptNilTransactionGiftReceivedArray) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptNilTransactionGiftReceivedArray to nil")
	}
	if d.Next() == jx.Null {
		if err := d.Null(); err != nil {
			return err
		}

		var v []TransactionGiftReceived
		o.Value = v
		o.Set = true
		o.Null = true
		return nil
	}
	o.Set = true
	o.Null = false
	o.Value = make([]TransactionGiftReceived, 0)
	if err := d.Arr(func(d *jx.Decoder) error {
		var elem TransactionGiftReceived
		if err := elem.Decode(d); err != nil {
			return err
		}
		o.Value = append(o.Value, elem)
		return nil
	}); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptNilTransactionGiftReceivedArray) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptNilTransactionGiftReceivedArray) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes []User as json.
func (o OptNilUserArray) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	if o.Null {
		e.Null()
		return
	}
	e.ArrStart()
	for _, elem := range o.Value {
		elem.Encode(e)
	}
	e.ArrEnd()
}

// Decode decodes []User from json.
func (o *OptNilUserArray) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptNilUserArray to nil")
	}
	if d.Next() == jx.Null {
		if err := d.Null(); err != nil {
			return err
		}

		var v []User
		o.Value = v
		o.Set = true
		o.Null = true
		return nil
	}
	o.Set = true
	o.Null = false
	o.Value = make([]User, 0)
	if err := d.Arr(func(d *jx.Decoder) error {
		var elem User
		if err := elem.Decode(d); err != nil {
			return err
		}
		o.Value = append(o.Value, elem)
		return nil
	}); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptNilUserArray) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptNilUserArray) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes []UserWrapper as json.
func (o OptNilUserWrapperArray) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	if o.Null {
		e.Null()
		return
	}
	e.ArrStart()
	for _, elem := range o.Value {
		elem.Encode(e)
	}
	e.ArrEnd()
}

// Decode decodes []UserWrapper from json.
func (o *OptNilUserWrapperArray) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptNilUserWrapperArray to nil")
	}
	if d.Next() == jx.Null {
		if err := d.Null(); err != nil {
			return err
		}

		var v []UserWrapper
		o.Value = v
		o.Set = true
		o.Null = true
		return nil
	}
	o.Set = true
	o.Null = false
	o.Value = make([]UserWrapper, 0)
	if err := d.Arr(func(d *jx.Decoder) error {
		var elem UserWrapper
		if err := elem.Decode(d); err != nil {
			return err
		}
		o.Value = append(o.Value, elem)
		return nil
	}); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptNilUserWrapperArray) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptNilUserWrapperArray) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes []Video as json.
func (o OptNilVideoArray) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	if o.Null {
		e.Null()
		return
	}
	e.ArrStart()
	for _, elem := range o.Value {
		elem.Encode(e)
	}
	e.ArrEnd()
}

// Decode decodes []Video from json.
func (o *OptNilVideoArray) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptNilVideoArray to nil")
	}
	if d.Next() == jx.Null {
		if err := d.Null(); err != nil {
			return err
		}

		var v []Video
		o.Value = v
		o.Set = true
		o.Null = true
		return nil
	}
	o.Set = true
	o.Null = false
	o.Value = make([]Video, 0)
	if err := d.Arr(func(d *jx.Decoder) error {
		var elem Video
		if err := elem.Decode(d); err != nil {
			return err
		}
		o.Value = append(o.Value, elem)
		return nil
	}); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptNilVideoArray) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptNilVideoArray) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes OnlineStatus as json.
func (o OptOnlineStatus) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	o.Value.Encode(e)
}

// Decode decodes OnlineStatus from json.
func (o *OptOnlineStatus) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptOnlineStatus to nil")
	}
	o.Set = true
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptOnlineStatus) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptOnlineStatus) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes OnlineStatusEnum as json.
func (o OptOnlineStatusEnum) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	e.Str(string(o.Value))
}

// Decode decodes OnlineStatusEnum from json.
func (o *OptOnlineStatusEnum) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptOnlineStatusEnum to nil")
	}
	o.Set = true
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptOnlineStatusEnum) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptOnlineStatusEnum) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes ParentMessage as json.
func (o OptParentMessage) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	o.Value.Encode(e)
}

// Decode decodes ParentMessage from json.
func (o *OptParentMessage) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptParentMessage to nil")
	}
	o.Set = true
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptParentMessage) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptParentMessage) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes Post as json.
func (o OptPost) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	o.Value.Encode(e)
}

// Decode decodes Post from json.
func (o *OptPost) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptPost to nil")
	}
	o.Set = true
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptPost) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptPost) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes PostType as json.
func (o OptPostType) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	e.Str(string(o.Value))
}

// Decode decodes PostType from json.
func (o *OptPostType) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptPostType to nil")
	}
	o.Set = true
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptPostType) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptPostType) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes RealmChatRoom as json.
func (o OptRealmChatRoom) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	o.Value.Encode(e)
}

// Decode decodes RealmChatRoom from json.
func (o *OptRealmChatRoom) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptRealmChatRoom to nil")
	}
	o.Set = true
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptRealmChatRoom) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptRealmChatRoom) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes RealmConferenceCall as json.
func (o OptRealmConferenceCall) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	o.Value.Encode(e)
}

// Decode decodes RealmConferenceCall from json.
func (o *OptRealmConferenceCall) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptRealmConferenceCall to nil")
	}
	o.Set = true
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptRealmConferenceCall) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptRealmConferenceCall) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes RealmGame as json.
func (o OptRealmGame) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	o.Value.Encode(e)
}

// Decode decodes RealmGame from json.
func (o *OptRealmGame) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptRealmGame to nil")
	}
	o.Set = true
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptRealmGame) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptRealmGame) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes RealmGenre as json.
func (o OptRealmGenre) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	o.Value.Encode(e)
}

// Decode decodes RealmGenre from json.
func (o *OptRealmGenre) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptRealmGenre to nil")
	}
	o.Set = true
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptRealmGenre) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptRealmGenre) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes RealmGift as json.
func (o OptRealmGift) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	o.Value.Encode(e)
}

// Decode decodes RealmGift from json.
func (o *OptRealmGift) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptRealmGift to nil")
	}
	o.Set = true
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptRealmGift) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptRealmGift) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes RealmGiftingAbility as json.
func (o OptRealmGiftingAbility) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	o.Value.Encode(e)
}

// Decode decodes RealmGiftingAbility from json.
func (o *OptRealmGiftingAbility) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptRealmGiftingAbility to nil")
	}
	o.Set = true
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptRealmGiftingAbility) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptRealmGiftingAbility) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes RealmPlatformDetails as json.
func (o OptRealmPlatformDetails) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	o.Value.Encode(e)
}

// Decode decodes RealmPlatformDetails from json.
func (o *OptRealmPlatformDetails) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptRealmPlatformDetails to nil")
	}
	o.Set = true
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptRealmPlatformDetails) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptRealmPlatformDetails) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes RealmUser as json.
func (o OptRealmUser) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	o.Value.Encode(e)
}

// Decode decodes RealmUser from json.
func (o *OptRealmUser) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptRealmUser to nil")
	}
	o.Set = true
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptRealmUser) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptRealmUser) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes *RepostReqMessageTags as json.
func (o OptRepostReqMessageTags) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	o.Value.Encode(e)
}

// Decode decodes *RepostReqMessageTags from json.
func (o *OptRepostReqMessageTags) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptRepostReqMessageTags to nil")
	}
	o.Set = true
	o.Value = new(RepostReqMessageTags)
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptRepostReqMessageTags) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptRepostReqMessageTags) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes *RepostReqSharedURL as json.
func (o OptRepostReqSharedURL) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	o.Value.Encode(e)
}

// Decode decodes *RepostReqSharedURL from json.
func (o *OptRepostReqSharedURL) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptRepostReqSharedURL to nil")
	}
	o.Set = true
	o.Value = new(RepostReqSharedURL)
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptRepostReqSharedURL) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptRepostReqSharedURL) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes ReviewRestriction as json.
func (o OptReviewRestriction) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	o.Value.Encode(e)
}

// Decode decodes ReviewRestriction from json.
func (o *OptReviewRestriction) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptReviewRestriction to nil")
	}
	o.Set = true
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptReviewRestriction) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptReviewRestriction) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes SearchCriteria as json.
func (o OptSearchCriteria) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	o.Value.Encode(e)
}

// Decode decodes SearchCriteria from json.
func (o *OptSearchCriteria) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptSearchCriteria to nil")
	}
	o.Set = true
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptSearchCriteria) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptSearchCriteria) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes Setting as json.
func (o OptSetting) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	o.Value.Encode(e)
}

// Decode decodes Setting from json.
func (o *OptSetting) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptSetting to nil")
	}
	o.Set = true
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptSetting) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptSetting) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes Settings as json.
func (o OptSettings) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	o.Value.Encode(e)
}

// Decode decodes Settings from json.
func (o *OptSettings) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptSettings to nil")
	}
	o.Set = true
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptSettings) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptSettings) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes Shareable as json.
func (o OptShareable) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	o.Value.Encode(e)
}

// Decode decodes Shareable from json.
func (o *OptShareable) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptShareable to nil")
	}
	o.Set = true
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptShareable) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptShareable) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes SharedUrl as json.
func (o OptSharedUrl) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	o.Value.Encode(e)
}

// Decode decodes SharedUrl from json.
func (o *OptSharedUrl) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptSharedUrl to nil")
	}
	o.Set = true
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptSharedUrl) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptSharedUrl) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes SignaturePayload as json.
func (o OptSignaturePayload) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	o.Value.Encode(e)
}

// Decode decodes SignaturePayload from json.
func (o *OptSignaturePayload) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptSignaturePayload to nil")
	}
	o.Set = true
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptSignaturePayload) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptSignaturePayload) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes SnsInfo as json.
func (o OptSnsInfo) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	o.Value.Encode(e)
}

// Decode decodes SnsInfo from json.
func (o *OptSnsInfo) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptSnsInfo to nil")
	}
	o.Set = true
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptSnsInfo) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptSnsInfo) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes Sticker as json.
func (o OptSticker) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	o.Value.Encode(e)
}

// Decode decodes Sticker from json.
func (o *OptSticker) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptSticker to nil")
	}
	o.Set = true
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptSticker) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptSticker) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes Survey as json.
func (o OptSurvey) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	o.Value.Encode(e)
}

// Decode decodes Survey from json.
func (o *OptSurvey) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptSurvey to nil")
	}
	o.Set = true
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptSurvey) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptSurvey) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes ThreadInfo as json.
func (o OptThreadInfo) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	o.Value.Encode(e)
}

// Decode decodes ThreadInfo from json.
func (o *OptThreadInfo) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptThreadInfo to nil")
	}
	o.Set = true
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptThreadInfo) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptThreadInfo) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes TimelineSettings as json.
func (o OptTimelineSettings) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	o.Value.Encode(e)
}

// Decode decodes TimelineSettings from json.
func (o *OptTimelineSettings) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptTimelineSettings to nil")
	}
	o.Set = true
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptTimelineSettings) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptTimelineSettings) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes Title as json.
func (o OptTitle) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	o.Value.Encode(e)
}

// Decode decodes Title from json.
func (o *OptTitle) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptTitle to nil")
	}
	o.Set = true
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptTitle) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptTitle) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes TransactionGiftReceivedItem as json.
func (o OptTransactionGiftReceivedItem) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	o.Value.Encode(e)
}

// Decode decodes TransactionGiftReceivedItem from json.
func (o *OptTransactionGiftReceivedItem) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptTransactionGiftReceivedItem to nil")
	}
	o.Set = true
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptTransactionGiftReceivedItem) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptTransactionGiftReceivedItem) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes TwoFaAuthRequiredDTO as json.
func (o OptTwoFaAuthRequiredDTO) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	o.Value.Encode(e)
}

// Decode decodes TwoFaAuthRequiredDTO from json.
func (o *OptTwoFaAuthRequiredDTO) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptTwoFaAuthRequiredDTO to nil")
	}
	o.Set = true
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptTwoFaAuthRequiredDTO) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptTwoFaAuthRequiredDTO) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes *UpdatePostReqMessageTags as json.
func (o OptUpdatePostReqMessageTags) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	o.Value.Encode(e)
}

// Decode decodes *UpdatePostReqMessageTags from json.
func (o *OptUpdatePostReqMessageTags) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptUpdatePostReqMessageTags to nil")
	}
	o.Set = true
	o.Value = new(UpdatePostReqMessageTags)
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptUpdatePostReqMessageTags) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptUpdatePostReqMessageTags) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes *UserConnectedBy as json.
func (o OptUserConnectedBy) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	o.Value.Encode(e)
}

// Decode decodes *UserConnectedBy from json.
func (o *OptUserConnectedBy) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptUserConnectedBy to nil")
	}
	o.Set = true
	o.Value = new(UserConnectedBy)
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptUserConnectedBy) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptUserConnectedBy) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes *UserCountry as json.
func (o OptUserCountry) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	o.Value.Encode(e)
}

// Decode decodes *UserCountry from json.
func (o *OptUserCountry) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptUserCountry to nil")
	}
	o.Set = true
	o.Value = new(UserCountry)
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptUserCountry) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptUserCountry) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes UserSetting as json.
func (o OptUserSetting) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	o.Value.Encode(e)
}

// Decode decodes UserSetting from json.
func (o *OptUserSetting) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptUserSetting to nil")
	}
	o.Set = true
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptUserSetting) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptUserSetting) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes UserUserDTO as json.
func (o OptUserUserDTO) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	o.Value.Encode(e)
}

// Decode decodes UserUserDTO from json.
func (o *OptUserUserDTO) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptUserUserDTO to nil")
	}
	o.Set = true
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptUserUserDTO) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptUserUserDTO) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *ParentMessage) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *ParentMessage) encodeFields(e *jx.Encoder) {
	{
		if s.Attachment.Set {
			e.FieldStart("attachment")
			s.Attachment.Encode(e)
		}
	}
	{
		if s.AttachmentThumbnail.Set {
			e.FieldStart("attachment_thumbnail")
			s.AttachmentThumbnail.Encode(e)
		}
	}
	{
		if s.CreatedAt.Set {
			e.FieldStart("created_at")
			s.CreatedAt.Encode(e)
		}
	}
	{
		if s.FontSize.Set {
			e.FieldStart("font_size")
			s.FontSize.Encode(e)
		}
	}
	{
		if s.GIF.Set {
			e.FieldStart("gif")
			s.GIF.Encode(e)
		}
	}
	{
		if s.ID.Set {
			e.FieldStart("id")
			s.ID.Encode(e)
		}
	}
	{
		if s.MessageType.Set {
			e.FieldStart("message_type")
			s.MessageType.Encode(e)
		}
	}
	{
		if s.RoomID.Set {
			e.FieldStart("room_id")
			s.RoomID.Encode(e)
		}
	}
	{
		if s.Sticker.Set {
			e.FieldStart("sticker")
			s.Sticker.Encode(e)
		}
	}
	{
		if s.Text.Set {
			e.FieldStart("text")
			s.Text.Encode(e)
		}
	}
	{
		if s.UserID.Set {
			e.FieldStart("user_id")
			s.UserID.Encode(e)
		}
	}
	{
		if s.VideoThumbnailURL.Set {
			e.FieldStart("video_thumbnail_url")
			s.VideoThumbnailURL.Encode(e)
		}
	}
	{
		if s.VideoURL.Set {
			e.FieldStart("video_url")
			s.VideoURL.Encode(e)
		}
	}
}

var jsonFieldsNameOfParentMessage = [13]string{
	0:  "attachment",
	1:  "attachment_thumbnail",
	2:  "created_at",
	3:  "font_size",
	4:  "gif",
	5:  "id",
	6:  "message_type",
	7:  "room_id",
	8:  "sticker",
	9:  "text",
	10: "user_id",
	11: "video_thumbnail_url",
	12: "video_url",
}

// Decode decodes ParentMessage from json.
func (s *ParentMessage) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode ParentMessage to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "attachment":
			if err := func() error {
				s.Attachment.Reset()
				if err := s.Attachment.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"attachment\"")
			}
		case "attachment_thumbnail":
			if err := func() error {
				s.AttachmentThumbnail.Reset()
				if err := s.AttachmentThumbnail.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"attachment_thumbnail\"")
			}
		case "created_at":
			if err := func() error {
				s.CreatedAt.Reset()
				if err := s.CreatedAt.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"created_at\"")
			}
		case "font_size":
			if err := func() error {
				s.FontSize.Reset()
				if err := s.FontSize.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"font_size\"")
			}
		case "gif":
			if err := func() error {
				s.GIF.Reset()
				if err := s.GIF.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"gif\"")
			}
		case "id":
			if err := func() error {
				s.ID.Reset()
				if err := s.ID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"id\"")
			}
		case "message_type":
			if err := func() error {
				s.MessageType.Reset()
				if err := s.MessageType.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"message_type\"")
			}
		case "room_id":
			if err := func() error {
				s.RoomID.Reset()
				if err := s.RoomID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"room_id\"")
			}
		case "sticker":
			if err := func() error {
				s.Sticker.Reset()
				if err := s.Sticker.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"sticker\"")
			}
		case "text":
			if err := func() error {
				s.Text.Reset()
				if err := s.Text.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"text\"")
			}
		case "user_id":
			if err := func() error {
				s.UserID.Reset()
				if err := s.UserID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"user_id\"")
			}
		case "video_thumbnail_url":
			if err := func() error {
				s.VideoThumbnailURL.Reset()
				if err := s.VideoThumbnailURL.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"video_thumbnail_url\"")
			}
		case "video_url":
			if err := func() error {
				s.VideoURL.Reset()
				if err := s.VideoURL.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"video_url\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode ParentMessage")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *ParentMessage) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *ParentMessage) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *PolicyAgreementsResponse) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *PolicyAgreementsResponse) encodeFields(e *jx.Encoder) {
	{
		if s.LatestDotmoneyAgreed.Set {
			e.FieldStart("latest_dotmoney_agreed")
			s.LatestDotmoneyAgreed.Encode(e)
		}
	}
	{
		if s.LatestEmpl2Agreed.Set {
			e.FieldStart("latest_empl2_agreed")
			s.LatestEmpl2Agreed.Encode(e)
		}
	}
	{
		if s.LatestPrivacyPolicyAgreed.Set {
			e.FieldStart("latest_privacy_policy_agreed")
			s.LatestPrivacyPolicyAgreed.Encode(e)
		}
	}
	{
		if s.LatestTermsOfUseAgreed.Set {
			e.FieldStart("latest_terms_of_use_agreed")
			s.LatestTermsOfUseAgreed.Encode(e)
		}
	}
}

var jsonFieldsNameOfPolicyAgreementsResponse = [4]string{
	0: "latest_dotmoney_agreed",
	1: "latest_empl2_agreed",
	2: "latest_privacy_policy_agreed",
	3: "latest_terms_of_use_agreed",
}

// Decode decodes PolicyAgreementsResponse from json.
func (s *PolicyAgreementsResponse) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode PolicyAgreementsResponse to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "latest_dotmoney_agreed":
			if err := func() error {
				s.LatestDotmoneyAgreed.Reset()
				if err := s.LatestDotmoneyAgreed.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"latest_dotmoney_agreed\"")
			}
		case "latest_empl2_agreed":
			if err := func() error {
				s.LatestEmpl2Agreed.Reset()
				if err := s.LatestEmpl2Agreed.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"latest_empl2_agreed\"")
			}
		case "latest_privacy_policy_agreed":
			if err := func() error {
				s.LatestPrivacyPolicyAgreed.Reset()
				if err := s.LatestPrivacyPolicyAgreed.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"latest_privacy_policy_agreed\"")
			}
		case "latest_terms_of_use_agreed":
			if err := func() error {
				s.LatestTermsOfUseAgreed.Reset()
				if err := s.LatestTermsOfUseAgreed.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"latest_terms_of_use_agreed\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode PolicyAgreementsResponse")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *PolicyAgreementsResponse) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *PolicyAgreementsResponse) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *PopularWord) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *PopularWord) encodeFields(e *jx.Encoder) {
	{
		if s.ID.Set {
			e.FieldStart("id")
			s.ID.Encode(e)
		}
	}
	{
		if s.Type.Set {
			e.FieldStart("type")
			s.Type.Encode(e)
		}
	}
	{
		if s.Word.Set {
			e.FieldStart("word")
			s.Word.Encode(e)
		}
	}
}

var jsonFieldsNameOfPopularWord = [3]string{
	0: "id",
	1: "type",
	2: "word",
}

// Decode decodes PopularWord from json.
func (s *PopularWord) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode PopularWord to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "id":
			if err := func() error {
				s.ID.Reset()
				if err := s.ID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"id\"")
			}
		case "type":
			if err := func() error {
				s.Type.Reset()
				if err := s.Type.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"type\"")
			}
		case "word":
			if err := func() error {
				s.Word.Reset()
				if err := s.Word.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"word\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode PopularWord")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *PopularWord) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *PopularWord) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *PopularWordsResponse) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *PopularWordsResponse) encodeFields(e *jx.Encoder) {
	{
		if s.PopularWords.Set {
			e.FieldStart("popular_words")
			s.PopularWords.Encode(e)
		}
	}
}

var jsonFieldsNameOfPopularWordsResponse = [1]string{
	0: "popular_words",
}

// Decode decodes PopularWordsResponse from json.
func (s *PopularWordsResponse) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode PopularWordsResponse to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "popular_words":
			if err := func() error {
				s.PopularWords.Reset()
				if err := s.PopularWords.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"popular_words\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode PopularWordsResponse")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *PopularWordsResponse) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *PopularWordsResponse) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *Post) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *Post) encodeFields(e *jx.Encoder) {
	{
		if s.Attachment.Set {
			e.FieldStart("attachment")
			s.Attachment.Encode(e)
		}
	}
	{
		if s.Attachment2.Set {
			e.FieldStart("attachment_2")
			s.Attachment2.Encode(e)
		}
	}
	{
		if s.Attachment2Thumbnail.Set {
			e.FieldStart("attachment_2_thumbnail")
			s.Attachment2Thumbnail.Encode(e)
		}
	}
	{
		if s.Attachment3.Set {
			e.FieldStart("attachment_3")
			s.Attachment3.Encode(e)
		}
	}
	{
		if s.Attachment3Thumbnail.Set {
			e.FieldStart("attachment_3_thumbnail")
			s.Attachment3Thumbnail.Encode(e)
		}
	}
	{
		if s.Attachment4.Set {
			e.FieldStart("attachment_4")
			s.Attachment4.Encode(e)
		}
	}
	{
		if s.Attachment4Thumbnail.Set {
			e.FieldStart("attachment_4_thumbnail")
			s.Attachment4Thumbnail.Encode(e)
		}
	}
	{
		if s.Attachment5.Set {
			e.FieldStart("attachment_5")
			s.Attachment5.Encode(e)
		}
	}
	{
		if s.Attachment5Thumbnail.Set {
			e.FieldStart("attachment_5_thumbnail")
			s.Attachment5Thumbnail.Encode(e)
		}
	}
	{
		if s.Attachment6.Set {
			e.FieldStart("attachment_6")
			s.Attachment6.Encode(e)
		}
	}
	{
		if s.Attachment6Thumbnail.Set {
			e.FieldStart("attachment_6_thumbnail")
			s.Attachment6Thumbnail.Encode(e)
		}
	}
	{
		if s.Attachment7.Set {
			e.FieldStart("attachment_7")
			s.Attachment7.Encode(e)
		}
	}
	{
		if s.Attachment7Thumbnail.Set {
			e.FieldStart("attachment_7_thumbnail")
			s.Attachment7Thumbnail.Encode(e)
		}
	}
	{
		if s.Attachment8.Set {
			e.FieldStart("attachment_8")
			s.Attachment8.Encode(e)
		}
	}
	{
		if s.Attachment8Thumbnail.Set {
			e.FieldStart("attachment_8_thumbnail")
			s.Attachment8Thumbnail.Encode(e)
		}
	}
	{
		if s.Attachment9.Set {
			e.FieldStart("attachment_9")
			s.Attachment9.Encode(e)
		}
	}
	{
		if s.Attachment9Thumbnail.Set {
			e.FieldStart("attachment_9_thumbnail")
			s.Attachment9Thumbnail.Encode(e)
		}
	}
	{
		if s.AttachmentThumbnail.Set {
			e.FieldStart("attachment_thumbnail")
			s.AttachmentThumbnail.Encode(e)
		}
	}
	{
		if s.Color.Set {
			e.FieldStart("color")
			s.Color.Encode(e)
		}
	}
	{
		if s.ConferenceCall.Set {
			e.FieldStart("conference_call")
			s.ConferenceCall.Encode(e)
		}
	}
	{
		if s.ConversationID.Set {
			e.FieldStart("conversation_id")
			s.ConversationID.Encode(e)
		}
	}
	{
		if s.CreatedAt.Set {
			e.FieldStart("created_at")
			s.CreatedAt.Encode(e)
		}
	}
	{
		if s.EditedAt.Set {
			e.FieldStart("edited_at")
			s.EditedAt.Encode(e)
		}
	}
	{
		if s.FontSize.Set {
			e.FieldStart("font_size")
			s.FontSize.Encode(e)
		}
	}
	{
		if s.GiftsCount.Set {
			e.FieldStart("gifts_count")
			s.GiftsCount.Encode(e)
		}
	}
	{
		if s.Group.Set {
			e.FieldStart("group")
			s.Group.Encode(e)
		}
	}
	{
		if s.GroupID.Set {
			e.FieldStart("group_id")
			s.GroupID.Encode(e)
		}
	}
	{
		if s.Highlighted.Set {
			e.FieldStart("highlighted")
			s.Highlighted.Encode(e)
		}
	}
	{
		if s.ID.Set {
			e.FieldStart("id")
			s.ID.Encode(e)
		}
	}
	{
		if s.InReplyTo.Set {
			e.FieldStart("in_reply_to")
			s.InReplyTo.Encode(e)
		}
	}
	{
		if s.InReplyToPost != nil {
			e.FieldStart("in_reply_to_post")
			s.InReplyToPost.Encode(e)
		}
	}
	{
		if s.InReplyToPostCount.Set {
			e.FieldStart("in_reply_to_post_count")
			s.InReplyToPostCount.Encode(e)
		}
	}
	{
		if s.IsFailToSend.Set {
			e.FieldStart("is_fail_to_send")
			s.IsFailToSend.Encode(e)
		}
	}
	{
		if s.IsMuted.Set {
			e.FieldStart("is_muted")
			s.IsMuted.Encode(e)
		}
	}
	{
		if s.Liked.Set {
			e.FieldStart("liked")
			s.Liked.Encode(e)
		}
	}
	{
		if s.Likers.Set {
			e.FieldStart("likers")
			s.Likers.Encode(e)
		}
	}
	{
		if s.LikersCount.Set {
			e.FieldStart("likers_count")
			s.LikersCount.Encode(e)
		}
	}
	{
		if s.LikesCount.Set {
			e.FieldStart("likes_count")
			s.LikesCount.Encode(e)
		}
	}
	{
		if s.Mentions.Set {
			e.FieldStart("mentions")
			s.Mentions.Encode(e)
		}
	}
	{
		if s.MessageTags.Set {
			e.FieldStart("message_tags")
			s.MessageTags.Encode(e)
		}
	}
	{
		if s.PostType.Set {
			e.FieldStart("post_type")
			s.PostType.Encode(e)
		}
	}
	{
		if s.ReportedCount.Set {
			e.FieldStart("reported_count")
			s.ReportedCount.Encode(e)
		}
	}
	{
		if s.Repostable.Set {
			e.FieldStart("repostable")
			s.Repostable.Encode(e)
		}
	}
	{
		if s.Reposted.Set {
			e.FieldStart("reposted")
			s.Reposted.Encode(e)
		}
	}
	{
		if s.RepostsCount.Set {
			e.FieldStart("reposts_count")
			s.RepostsCount.Encode(e)
		}
	}
	{
		if s.Shareable != nil {
			e.FieldStart("shareable")
			s.Shareable.Encode(e)
		}
	}
	{
		if s.SharedThread != nil {
			e.FieldStart("shared_thread")
			s.SharedThread.Encode(e)
		}
	}
	{
		if s.SharedURL.Set {
			e.FieldStart("shared_url")
			s.SharedURL.Encode(e)
		}
	}
	{
		if s.Survey.Set {
			e.FieldStart("survey")
			s.Survey.Encode(e)
		}
	}
	{
		if s.Tag.Set {
			e.FieldStart("tag")
			s.Tag.Encode(e)
		}
	}
	{
		if s.Text.Set {
			e.FieldStart("text")
			s.Text.Encode(e)
		}
	}
	{
		if s.Thread != nil {
			e.FieldStart("thread")
			s.Thread.Encode(e)
		}
	}
	{
		if s.ThreadID.Set {
			e.FieldStart("thread_id")
			s.ThreadID.Encode(e)
		}
	}
	{
		if s.TotalGiftsCount.Set {
			e.FieldStart("total_gifts_count")
			s.TotalGiftsCount.Encode(e)
		}
	}
	{
		if s.UpdatedAt.Set {
			e.FieldStart("updated_at")
			s.UpdatedAt.Encode(e)
		}
	}
	{
		if s.User.Set {
			e.FieldStart("user")
			s.User.Encode(e)
		}
	}
	{
		if s.Videos.Set {
			e.FieldStart("videos")
			s.Videos.Encode(e)
		}
	}
}

var jsonFieldsNameOfPost = [57]string{
	0:  "attachment",
	1:  "attachment_2",
	2:  "attachment_2_thumbnail",
	3:  "attachment_3",
	4:  "attachment_3_thumbnail",
	5:  "attachment_4",
	6:  "attachment_4_thumbnail",
	7:  "attachment_5",
	8:  "attachment_5_thumbnail",
	9:  "attachment_6",
	10: "attachment_6_thumbnail",
	11: "attachment_7",
	12: "attachment_7_thumbnail",
	13: "attachment_8",
	14: "attachment_8_thumbnail",
	15: "attachment_9",
	16: "attachment_9_thumbnail",
	17: "attachment_thumbnail",
	18: "color",
	19: "conference_call",
	20: "conversation_id",
	21: "created_at",
	22: "edited_at",
	23: "font_size",
	24: "gifts_count",
	25: "group",
	26: "group_id",
	27: "highlighted",
	28: "id",
	29: "in_reply_to",
	30: "in_reply_to_post",
	31: "in_reply_to_post_count",
	32: "is_fail_to_send",
	33: "is_muted",
	34: "liked",
	35: "likers",
	36: "likers_count",
	37: "likes_count",
	38: "mentions",
	39: "message_tags",
	40: "post_type",
	41: "reported_count",
	42: "repostable",
	43: "reposted",
	44: "reposts_count",
	45: "shareable",
	46: "shared_thread",
	47: "shared_url",
	48: "survey",
	49: "tag",
	50: "text",
	51: "thread",
	52: "thread_id",
	53: "total_gifts_count",
	54: "updated_at",
	55: "user",
	56: "videos",
}

// Decode decodes Post from json.
func (s *Post) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Post to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "attachment":
			if err := func() error {
				s.Attachment.Reset()
				if err := s.Attachment.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"attachment\"")
			}
		case "attachment_2":
			if err := func() error {
				s.Attachment2.Reset()
				if err := s.Attachment2.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"attachment_2\"")
			}
		case "attachment_2_thumbnail":
			if err := func() error {
				s.Attachment2Thumbnail.Reset()
				if err := s.Attachment2Thumbnail.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"attachment_2_thumbnail\"")
			}
		case "attachment_3":
			if err := func() error {
				s.Attachment3.Reset()
				if err := s.Attachment3.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"attachment_3\"")
			}
		case "attachment_3_thumbnail":
			if err := func() error {
				s.Attachment3Thumbnail.Reset()
				if err := s.Attachment3Thumbnail.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"attachment_3_thumbnail\"")
			}
		case "attachment_4":
			if err := func() error {
				s.Attachment4.Reset()
				if err := s.Attachment4.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"attachment_4\"")
			}
		case "attachment_4_thumbnail":
			if err := func() error {
				s.Attachment4Thumbnail.Reset()
				if err := s.Attachment4Thumbnail.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"attachment_4_thumbnail\"")
			}
		case "attachment_5":
			if err := func() error {
				s.Attachment5.Reset()
				if err := s.Attachment5.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"attachment_5\"")
			}
		case "attachment_5_thumbnail":
			if err := func() error {
				s.Attachment5Thumbnail.Reset()
				if err := s.Attachment5Thumbnail.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"attachment_5_thumbnail\"")
			}
		case "attachment_6":
			if err := func() error {
				s.Attachment6.Reset()
				if err := s.Attachment6.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"attachment_6\"")
			}
		case "attachment_6_thumbnail":
			if err := func() error {
				s.Attachment6Thumbnail.Reset()
				if err := s.Attachment6Thumbnail.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"attachment_6_thumbnail\"")
			}
		case "attachment_7":
			if err := func() error {
				s.Attachment7.Reset()
				if err := s.Attachment7.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"attachment_7\"")
			}
		case "attachment_7_thumbnail":
			if err := func() error {
				s.Attachment7Thumbnail.Reset()
				if err := s.Attachment7Thumbnail.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"attachment_7_thumbnail\"")
			}
		case "attachment_8":
			if err := func() error {
				s.Attachment8.Reset()
				if err := s.Attachment8.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"attachment_8\"")
			}
		case "attachment_8_thumbnail":
			if err := func() error {
				s.Attachment8Thumbnail.Reset()
				if err := s.Attachment8Thumbnail.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"attachment_8_thumbnail\"")
			}
		case "attachment_9":
			if err := func() error {
				s.Attachment9.Reset()
				if err := s.Attachment9.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"attachment_9\"")
			}
		case "attachment_9_thumbnail":
			if err := func() error {
				s.Attachment9Thumbnail.Reset()
				if err := s.Attachment9Thumbnail.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"attachment_9_thumbnail\"")
			}
		case "attachment_thumbnail":
			if err := func() error {
				s.AttachmentThumbnail.Reset()
				if err := s.AttachmentThumbnail.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"attachment_thumbnail\"")
			}
		case "color":
			if err := func() error {
				s.Color.Reset()
				if err := s.Color.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"color\"")
			}
		case "conference_call":
			if err := func() error {
				s.ConferenceCall.Reset()
				if err := s.ConferenceCall.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"conference_call\"")
			}
		case "conversation_id":
			if err := func() error {
				s.ConversationID.Reset()
				if err := s.ConversationID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"conversation_id\"")
			}
		case "created_at":
			if err := func() error {
				s.CreatedAt.Reset()
				if err := s.CreatedAt.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"created_at\"")
			}
		case "edited_at":
			if err := func() error {
				s.EditedAt.Reset()
				if err := s.EditedAt.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"edited_at\"")
			}
		case "font_size":
			if err := func() error {
				s.FontSize.Reset()
				if err := s.FontSize.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"font_size\"")
			}
		case "gifts_count":
			if err := func() error {
				s.GiftsCount.Reset()
				if err := s.GiftsCount.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"gifts_count\"")
			}
		case "group":
			if err := func() error {
				s.Group.Reset()
				if err := s.Group.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"group\"")
			}
		case "group_id":
			if err := func() error {
				s.GroupID.Reset()
				if err := s.GroupID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"group_id\"")
			}
		case "highlighted":
			if err := func() error {
				s.Highlighted.Reset()
				if err := s.Highlighted.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"highlighted\"")
			}
		case "id":
			if err := func() error {
				s.ID.Reset()
				if err := s.ID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"id\"")
			}
		case "in_reply_to":
			if err := func() error {
				s.InReplyTo.Reset()
				if err := s.InReplyTo.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"in_reply_to\"")
			}
		case "in_reply_to_post":
			if err := func() error {
				s.InReplyToPost = nil
				var elem Post
				if err := elem.Decode(d); err != nil {
					return err
				}
				s.InReplyToPost = &elem
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"in_reply_to_post\"")
			}
		case "in_reply_to_post_count":
			if err := func() error {
				s.InReplyToPostCount.Reset()
				if err := s.InReplyToPostCount.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"in_reply_to_post_count\"")
			}
		case "is_fail_to_send":
			if err := func() error {
				s.IsFailToSend.Reset()
				if err := s.IsFailToSend.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"is_fail_to_send\"")
			}
		case "is_muted":
			if err := func() error {
				s.IsMuted.Reset()
				if err := s.IsMuted.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"is_muted\"")
			}
		case "liked":
			if err := func() error {
				s.Liked.Reset()
				if err := s.Liked.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"liked\"")
			}
		case "likers":
			if err := func() error {
				s.Likers.Reset()
				if err := s.Likers.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"likers\"")
			}
		case "likers_count":
			if err := func() error {
				s.LikersCount.Reset()
				if err := s.LikersCount.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"likers_count\"")
			}
		case "likes_count":
			if err := func() error {
				s.LikesCount.Reset()
				if err := s.LikesCount.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"likes_count\"")
			}
		case "mentions":
			if err := func() error {
				s.Mentions.Reset()
				if err := s.Mentions.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"mentions\"")
			}
		case "message_tags":
			if err := func() error {
				s.MessageTags.Reset()
				if err := s.MessageTags.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"message_tags\"")
			}
		case "post_type":
			if err := func() error {
				s.PostType.Reset()
				if err := s.PostType.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"post_type\"")
			}
		case "reported_count":
			if err := func() error {
				s.ReportedCount.Reset()
				if err := s.ReportedCount.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"reported_count\"")
			}
		case "repostable":
			if err := func() error {
				s.Repostable.Reset()
				if err := s.Repostable.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"repostable\"")
			}
		case "reposted":
			if err := func() error {
				s.Reposted.Reset()
				if err := s.Reposted.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"reposted\"")
			}
		case "reposts_count":
			if err := func() error {
				s.RepostsCount.Reset()
				if err := s.RepostsCount.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"reposts_count\"")
			}
		case "shareable":
			if err := func() error {
				s.Shareable = nil
				var elem Shareable
				if err := elem.Decode(d); err != nil {
					return err
				}
				s.Shareable = &elem
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"shareable\"")
			}
		case "shared_thread":
			if err := func() error {
				s.SharedThread = nil
				var elem ThreadInfo
				if err := elem.Decode(d); err != nil {
					return err
				}
				s.SharedThread = &elem
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"shared_thread\"")
			}
		case "shared_url":
			if err := func() error {
				s.SharedURL.Reset()
				if err := s.SharedURL.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"shared_url\"")
			}
		case "survey":
			if err := func() error {
				s.Survey.Reset()
				if err := s.Survey.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"survey\"")
			}
		case "tag":
			if err := func() error {
				s.Tag.Reset()
				if err := s.Tag.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"tag\"")
			}
		case "text":
			if err := func() error {
				s.Text.Reset()
				if err := s.Text.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"text\"")
			}
		case "thread":
			if err := func() error {
				s.Thread = nil
				var elem ThreadInfo
				if err := elem.Decode(d); err != nil {
					return err
				}
				s.Thread = &elem
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"thread\"")
			}
		case "thread_id":
			if err := func() error {
				s.ThreadID.Reset()
				if err := s.ThreadID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"thread_id\"")
			}
		case "total_gifts_count":
			if err := func() error {
				s.TotalGiftsCount.Reset()
				if err := s.TotalGiftsCount.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"total_gifts_count\"")
			}
		case "updated_at":
			if err := func() error {
				s.UpdatedAt.Reset()
				if err := s.UpdatedAt.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"updated_at\"")
			}
		case "user":
			if err := func() error {
				s.User.Reset()
				if err := s.User.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"user\"")
			}
		case "videos":
			if err := func() error {
				s.Videos.Reset()
				if err := s.Videos.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"videos\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode Post")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *Post) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *Post) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *PostLikersResponse) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *PostLikersResponse) encodeFields(e *jx.Encoder) {
	{
		if s.LastID.Set {
			e.FieldStart("last_id")
			s.LastID.Encode(e)
		}
	}
	{
		if s.Users.Set {
			e.FieldStart("users")
			s.Users.Encode(e)
		}
	}
}

var jsonFieldsNameOfPostLikersResponse = [2]string{
	0: "last_id",
	1: "users",
}

// Decode decodes PostLikersResponse from json.
func (s *PostLikersResponse) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode PostLikersResponse to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "last_id":
			if err := func() error {
				s.LastID.Reset()
				if err := s.LastID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"last_id\"")
			}
		case "users":
			if err := func() error {
				s.Users.Reset()
				if err := s.Users.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"users\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode PostLikersResponse")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *PostLikersResponse) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *PostLikersResponse) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *PostResponse) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *PostResponse) encodeFields(e *jx.Encoder) {
	{
		if s.Post.Set {
			e.FieldStart("post")
			s.Post.Encode(e)
		}
	}
}

var jsonFieldsNameOfPostResponse = [1]string{
	0: "post",
}

// Decode decodes PostResponse from json.
func (s *PostResponse) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode PostResponse to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "post":
			if err := func() error {
				s.Post.Reset()
				if err := s.Post.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"post\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode PostResponse")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *PostResponse) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *PostResponse) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *PostTag) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *PostTag) encodeFields(e *jx.Encoder) {
	{
		if s.ID.Set {
			e.FieldStart("id")
			s.ID.Encode(e)
		}
	}
	{
		if s.PostHashtagsCount.Set {
			e.FieldStart("post_hashtags_count")
			s.PostHashtagsCount.Encode(e)
		}
	}
	{
		if s.Tag.Set {
			e.FieldStart("tag")
			s.Tag.Encode(e)
		}
	}
}

var jsonFieldsNameOfPostTag = [3]string{
	0: "id",
	1: "post_hashtags_count",
	2: "tag",
}

// Decode decodes PostTag from json.
func (s *PostTag) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode PostTag to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "id":
			if err := func() error {
				s.ID.Reset()
				if err := s.ID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"id\"")
			}
		case "post_hashtags_count":
			if err := func() error {
				s.PostHashtagsCount.Reset()
				if err := s.PostHashtagsCount.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"post_hashtags_count\"")
			}
		case "tag":
			if err := func() error {
				s.Tag.Reset()
				if err := s.Tag.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"tag\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode PostTag")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *PostTag) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *PostTag) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *PostTagsResponse) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *PostTagsResponse) encodeFields(e *jx.Encoder) {
	{
		if s.Tags.Set {
			e.FieldStart("tags")
			s.Tags.Encode(e)
		}
	}
}

var jsonFieldsNameOfPostTagsResponse = [1]string{
	0: "tags",
}

// Decode decodes PostTagsResponse from json.
func (s *PostTagsResponse) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode PostTagsResponse to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "tags":
			if err := func() error {
				s.Tags.Reset()
				if err := s.Tags.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"tags\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode PostTagsResponse")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *PostTagsResponse) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *PostTagsResponse) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes PostType as json.
func (s PostType) Encode(e *jx.Encoder) {
	e.Str(string(s))
}

// Decode decodes PostType from json.
func (s *PostType) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode PostType to nil")
	}
	v, err := d.StrBytes()
	if err != nil {
		return err
	}
	// Try to use constant string.
	switch PostType(v) {
	case PostTypeCall:
		*s = PostTypeCall
	case PostTypeImage:
		*s = PostTypeImage
	case PostTypeVideo:
		*s = PostTypeVideo
	case PostTypeSurvey:
		*s = PostTypeSurvey
	case PostTypeRepost:
		*s = PostTypeRepost
	case PostTypeThread:
		*s = PostTypeThread
	case PostTypeShareableGroup:
		*s = PostTypeShareableGroup
	case PostTypeShareableURL:
		*s = PostTypeShareableURL
	case PostTypeYoutube:
		*s = PostTypeYoutube
	case PostTypeShareableThread:
		*s = PostTypeShareableThread
	case PostTypeSharepal:
		*s = PostTypeSharepal
	case PostTypeUndefined:
		*s = PostTypeUndefined
	case PostTypeText:
		*s = PostTypeText
	default:
		*s = PostType(v)
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s PostType) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *PostType) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *PostsResponse) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *PostsResponse) encodeFields(e *jx.Encoder) {
	{
		if s.HasMoreHotPosts.Set {
			e.FieldStart("has_more_hot_posts")
			s.HasMoreHotPosts.Encode(e)
		}
	}
	{
		if s.NextPageValue.Set {
			e.FieldStart("next_page_value")
			s.NextPageValue.Encode(e)
		}
	}
	{
		if s.PinnedPosts.Set {
			e.FieldStart("pinned_posts")
			s.PinnedPosts.Encode(e)
		}
	}
	{
		if s.Posts.Set {
			e.FieldStart("posts")
			s.Posts.Encode(e)
		}
	}
}

var jsonFieldsNameOfPostsResponse = [4]string{
	0: "has_more_hot_posts",
	1: "next_page_value",
	2: "pinned_posts",
	3: "posts",
}

// Decode decodes PostsResponse from json.
func (s *PostsResponse) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode PostsResponse to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "has_more_hot_posts":
			if err := func() error {
				s.HasMoreHotPosts.Reset()
				if err := s.HasMoreHotPosts.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"has_more_hot_posts\"")
			}
		case "next_page_value":
			if err := func() error {
				s.NextPageValue.Reset()
				if err := s.NextPageValue.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"next_page_value\"")
			}
		case "pinned_posts":
			if err := func() error {
				s.PinnedPosts.Reset()
				if err := s.PinnedPosts.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"pinned_posts\"")
			}
		case "posts":
			if err := func() error {
				s.Posts.Reset()
				if err := s.Posts.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"posts\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode PostsResponse")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *PostsResponse) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *PostsResponse) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *PresignedUrl) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *PresignedUrl) encodeFields(e *jx.Encoder) {
	{
		if s.Filename.Set {
			e.FieldStart("filename")
			s.Filename.Encode(e)
		}
	}
	{
		if s.URL.Set {
			e.FieldStart("url")
			s.URL.Encode(e)
		}
	}
}

var jsonFieldsNameOfPresignedUrl = [2]string{
	0: "filename",
	1: "url",
}

// Decode decodes PresignedUrl from json.
func (s *PresignedUrl) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode PresignedUrl to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "filename":
			if err := func() error {
				s.Filename.Reset()
				if err := s.Filename.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"filename\"")
			}
		case "url":
			if err := func() error {
				s.URL.Reset()
				if err := s.URL.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"url\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode PresignedUrl")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *PresignedUrl) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *PresignedUrl) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *PresignedUrlResponse) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *PresignedUrlResponse) encodeFields(e *jx.Encoder) {
	{
		if s.PresignedURL.Set {
			e.FieldStart("presigned_url")
			s.PresignedURL.Encode(e)
		}
	}
}

var jsonFieldsNameOfPresignedUrlResponse = [1]string{
	0: "presigned_url",
}

// Decode decodes PresignedUrlResponse from json.
func (s *PresignedUrlResponse) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode PresignedUrlResponse to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "presigned_url":
			if err := func() error {
				s.PresignedURL.Reset()
				if err := s.PresignedURL.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"presigned_url\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode PresignedUrlResponse")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *PresignedUrlResponse) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *PresignedUrlResponse) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *PresignedUrlsResponse) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *PresignedUrlsResponse) encodeFields(e *jx.Encoder) {
	{
		if s.PresignedUrls.Set {
			e.FieldStart("presigned_urls")
			s.PresignedUrls.Encode(e)
		}
	}
}

var jsonFieldsNameOfPresignedUrlsResponse = [1]string{
	0: "presigned_urls",
}

// Decode decodes PresignedUrlsResponse from json.
func (s *PresignedUrlsResponse) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode PresignedUrlsResponse to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "presigned_urls":
			if err := func() error {
				s.PresignedUrls.Reset()
				if err := s.PresignedUrls.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"presigned_urls\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode PresignedUrlsResponse")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *PresignedUrlsResponse) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *PresignedUrlsResponse) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *ReadAttachmentRequest) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *ReadAttachmentRequest) encodeFields(e *jx.Encoder) {
	{
		if s.AttachmentMsgIds.Set {
			e.FieldStart("attachment_msg_ids")
			s.AttachmentMsgIds.Encode(e)
		}
	}
}

var jsonFieldsNameOfReadAttachmentRequest = [1]string{
	0: "attachment_msg_ids",
}

// Decode decodes ReadAttachmentRequest from json.
func (s *ReadAttachmentRequest) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode ReadAttachmentRequest to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "attachment_msg_ids":
			if err := func() error {
				s.AttachmentMsgIds.Reset()
				if err := s.AttachmentMsgIds.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"attachment_msg_ids\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode ReadAttachmentRequest")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *ReadAttachmentRequest) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *ReadAttachmentRequest) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *RealmChatRoom) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *RealmChatRoom) encodeFields(e *jx.Encoder) {
	{
		if s.Background.Set {
			e.FieldStart("background")
			s.Background.Encode(e)
		}
	}
	{
		if s.ID.Set {
			e.FieldStart("id")
			s.ID.Encode(e)
		}
	}
	{
		if s.IsGroup.Set {
			e.FieldStart("is_group")
			s.IsGroup.Encode(e)
		}
	}
	{
		if s.IsRequest.Set {
			e.FieldStart("is_request")
			s.IsRequest.Encode(e)
		}
	}
	{
		if s.LastMessage.Set {
			e.FieldStart("last_message")
			s.LastMessage.Encode(e)
		}
	}
	{
		if s.Members.Set {
			e.FieldStart("members")
			s.Members.Encode(e)
		}
	}
	{
		if s.Name.Set {
			e.FieldStart("name")
			s.Name.Encode(e)
		}
	}
	{
		if s.Owner.Set {
			e.FieldStart("owner")
			s.Owner.Encode(e)
		}
	}
	{
		if s.UnreadCount.Set {
			e.FieldStart("unread_count")
			s.UnreadCount.Encode(e)
		}
	}
	{
		if s.UpdatedAt.Set {
			e.FieldStart("updated_at")
			s.UpdatedAt.Encode(e)
		}
	}
	{
		if s.UserSetting.Set {
			e.FieldStart("user_setting")
			s.UserSetting.Encode(e)
		}
	}
}

var jsonFieldsNameOfRealmChatRoom = [11]string{
	0:  "background",
	1:  "id",
	2:  "is_group",
	3:  "is_request",
	4:  "last_message",
	5:  "members",
	6:  "name",
	7:  "owner",
	8:  "unread_count",
	9:  "updated_at",
	10: "user_setting",
}

// Decode decodes RealmChatRoom from json.
func (s *RealmChatRoom) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode RealmChatRoom to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "background":
			if err := func() error {
				s.Background.Reset()
				if err := s.Background.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"background\"")
			}
		case "id":
			if err := func() error {
				s.ID.Reset()
				if err := s.ID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"id\"")
			}
		case "is_group":
			if err := func() error {
				s.IsGroup.Reset()
				if err := s.IsGroup.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"is_group\"")
			}
		case "is_request":
			if err := func() error {
				s.IsRequest.Reset()
				if err := s.IsRequest.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"is_request\"")
			}
		case "last_message":
			if err := func() error {
				s.LastMessage.Reset()
				if err := s.LastMessage.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"last_message\"")
			}
		case "members":
			if err := func() error {
				s.Members.Reset()
				if err := s.Members.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"members\"")
			}
		case "name":
			if err := func() error {
				s.Name.Reset()
				if err := s.Name.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"name\"")
			}
		case "owner":
			if err := func() error {
				s.Owner.Reset()
				if err := s.Owner.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"owner\"")
			}
		case "unread_count":
			if err := func() error {
				s.UnreadCount.Reset()
				if err := s.UnreadCount.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"unread_count\"")
			}
		case "updated_at":
			if err := func() error {
				s.UpdatedAt.Reset()
				if err := s.UpdatedAt.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"updated_at\"")
			}
		case "user_setting":
			if err := func() error {
				s.UserSetting.Reset()
				if err := s.UserSetting.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"user_setting\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode RealmChatRoom")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *RealmChatRoom) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *RealmChatRoom) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *RealmConferenceCall) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *RealmConferenceCall) encodeFields(e *jx.Encoder) {
	{
		if s.Active.Set {
			e.FieldStart("active")
			s.Active.Encode(e)
		}
	}
	{
		if s.AgoraChannel.Set {
			e.FieldStart("agora_channel")
			s.AgoraChannel.Encode(e)
		}
	}
	{
		if s.AgoraToken.Set {
			e.FieldStart("agora_token")
			s.AgoraToken.Encode(e)
		}
	}
	{
		if s.AnonymousCallUsersCount.Set {
			e.FieldStart("anonymous_call_users_count")
			s.AnonymousCallUsersCount.Encode(e)
		}
	}
	{
		if s.BumpParams.Set {
			e.FieldStart("bump_params")
			s.BumpParams.Encode(e)
		}
	}
	{
		if s.CallType.Set {
			e.FieldStart("call_type")
			s.CallType.Encode(e)
		}
	}
	{
		if s.ConferenceCallUserRoles.Set {
			e.FieldStart("conference_call_user_roles")
			s.ConferenceCallUserRoles.Encode(e)
		}
	}
	{
		if s.ConferenceCallUsers.Set {
			e.FieldStart("conference_call_users")
			s.ConferenceCallUsers.Encode(e)
		}
	}
	{
		if s.ConferenceCallUsersCount.Set {
			e.FieldStart("conference_call_users_count")
			s.ConferenceCallUsersCount.Encode(e)
		}
	}
	{
		if s.DurationSeconds.Set {
			e.FieldStart("duration_seconds")
			s.DurationSeconds.Encode(e)
		}
	}
	{
		if s.Game.Set {
			e.FieldStart("game")
			s.Game.Encode(e)
		}
	}
	{
		if s.Genre.Set {
			e.FieldStart("genre")
			s.Genre.Encode(e)
		}
	}
	{
		if s.GroupID.Set {
			e.FieldStart("group_id")
			s.GroupID.Encode(e)
		}
	}
	{
		if s.ID.Set {
			e.FieldStart("id")
			s.ID.Encode(e)
		}
	}
	{
		if s.JoinableBy.Set {
			e.FieldStart("joinable_by")
			s.JoinableBy.Encode(e)
		}
	}
	{
		if s.MaxParticipants.Set {
			e.FieldStart("max_participants")
			s.MaxParticipants.Encode(e)
		}
	}
	{
		if s.PostID.Set {
			e.FieldStart("post_id")
			s.PostID.Encode(e)
		}
	}
	{
		if s.Server.Set {
			e.FieldStart("server")
			s.Server.Encode(e)
		}
	}
}

var jsonFieldsNameOfRealmConferenceCall = [18]string{
	0:  "active",
	1:  "agora_channel",
	2:  "agora_token",
	3:  "anonymous_call_users_count",
	4:  "bump_params",
	5:  "call_type",
	6:  "conference_call_user_roles",
	7:  "conference_call_users",
	8:  "conference_call_users_count",
	9:  "duration_seconds",
	10: "game",
	11: "genre",
	12: "group_id",
	13: "id",
	14: "joinable_by",
	15: "max_participants",
	16: "post_id",
	17: "server",
}

// Decode decodes RealmConferenceCall from json.
func (s *RealmConferenceCall) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode RealmConferenceCall to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "active":
			if err := func() error {
				s.Active.Reset()
				if err := s.Active.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"active\"")
			}
		case "agora_channel":
			if err := func() error {
				s.AgoraChannel.Reset()
				if err := s.AgoraChannel.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"agora_channel\"")
			}
		case "agora_token":
			if err := func() error {
				s.AgoraToken.Reset()
				if err := s.AgoraToken.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"agora_token\"")
			}
		case "anonymous_call_users_count":
			if err := func() error {
				s.AnonymousCallUsersCount.Reset()
				if err := s.AnonymousCallUsersCount.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"anonymous_call_users_count\"")
			}
		case "bump_params":
			if err := func() error {
				s.BumpParams.Reset()
				if err := s.BumpParams.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"bump_params\"")
			}
		case "call_type":
			if err := func() error {
				s.CallType.Reset()
				if err := s.CallType.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"call_type\"")
			}
		case "conference_call_user_roles":
			if err := func() error {
				s.ConferenceCallUserRoles.Reset()
				if err := s.ConferenceCallUserRoles.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"conference_call_user_roles\"")
			}
		case "conference_call_users":
			if err := func() error {
				s.ConferenceCallUsers.Reset()
				if err := s.ConferenceCallUsers.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"conference_call_users\"")
			}
		case "conference_call_users_count":
			if err := func() error {
				s.ConferenceCallUsersCount.Reset()
				if err := s.ConferenceCallUsersCount.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"conference_call_users_count\"")
			}
		case "duration_seconds":
			if err := func() error {
				s.DurationSeconds.Reset()
				if err := s.DurationSeconds.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"duration_seconds\"")
			}
		case "game":
			if err := func() error {
				s.Game.Reset()
				if err := s.Game.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"game\"")
			}
		case "genre":
			if err := func() error {
				s.Genre.Reset()
				if err := s.Genre.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"genre\"")
			}
		case "group_id":
			if err := func() error {
				s.GroupID.Reset()
				if err := s.GroupID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"group_id\"")
			}
		case "id":
			if err := func() error {
				s.ID.Reset()
				if err := s.ID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"id\"")
			}
		case "joinable_by":
			if err := func() error {
				s.JoinableBy.Reset()
				if err := s.JoinableBy.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"joinable_by\"")
			}
		case "max_participants":
			if err := func() error {
				s.MaxParticipants.Reset()
				if err := s.MaxParticipants.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"max_participants\"")
			}
		case "post_id":
			if err := func() error {
				s.PostID.Reset()
				if err := s.PostID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"post_id\"")
			}
		case "server":
			if err := func() error {
				s.Server.Reset()
				if err := s.Server.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"server\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode RealmConferenceCall")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *RealmConferenceCall) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *RealmConferenceCall) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *RealmGame) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *RealmGame) encodeFields(e *jx.Encoder) {
	{
		if s.IconURL.Set {
			e.FieldStart("icon_url")
			s.IconURL.Encode(e)
		}
	}
	{
		if s.ID.Set {
			e.FieldStart("id")
			s.ID.Encode(e)
		}
	}
	{
		if s.PlatformDetails.Set {
			e.FieldStart("platform_details")
			s.PlatformDetails.Encode(e)
		}
	}
	{
		if s.Title.Set {
			e.FieldStart("title")
			s.Title.Encode(e)
		}
	}
}

var jsonFieldsNameOfRealmGame = [4]string{
	0: "icon_url",
	1: "id",
	2: "platform_details",
	3: "title",
}

// Decode decodes RealmGame from json.
func (s *RealmGame) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode RealmGame to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "icon_url":
			if err := func() error {
				s.IconURL.Reset()
				if err := s.IconURL.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"icon_url\"")
			}
		case "id":
			if err := func() error {
				s.ID.Reset()
				if err := s.ID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"id\"")
			}
		case "platform_details":
			if err := func() error {
				s.PlatformDetails.Reset()
				if err := s.PlatformDetails.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"platform_details\"")
			}
		case "title":
			if err := func() error {
				s.Title.Reset()
				if err := s.Title.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"title\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode RealmGame")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *RealmGame) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *RealmGame) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *RealmGenre) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *RealmGenre) encodeFields(e *jx.Encoder) {
	{
		if s.IconURL.Set {
			e.FieldStart("icon_url")
			s.IconURL.Encode(e)
		}
	}
	{
		if s.ID.Set {
			e.FieldStart("id")
			s.ID.Encode(e)
		}
	}
	{
		if s.Title.Set {
			e.FieldStart("title")
			s.Title.Encode(e)
		}
	}
}

var jsonFieldsNameOfRealmGenre = [3]string{
	0: "icon_url",
	1: "id",
	2: "title",
}

// Decode decodes RealmGenre from json.
func (s *RealmGenre) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode RealmGenre to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "icon_url":
			if err := func() error {
				s.IconURL.Reset()
				if err := s.IconURL.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"icon_url\"")
			}
		case "id":
			if err := func() error {
				s.ID.Reset()
				if err := s.ID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"id\"")
			}
		case "title":
			if err := func() error {
				s.Title.Reset()
				if err := s.Title.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"title\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode RealmGenre")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *RealmGenre) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *RealmGenre) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *RealmGift) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *RealmGift) encodeFields(e *jx.Encoder) {
	{
		if s.Currency.Set {
			e.FieldStart("currency")
			s.Currency.Encode(e)
		}
	}
	{
		if s.Icon.Set {
			e.FieldStart("icon")
			s.Icon.Encode(e)
		}
	}
	{
		if s.IconThumbnail.Set {
			e.FieldStart("icon_thumbnail")
			s.IconThumbnail.Encode(e)
		}
	}
	{
		if s.ID.Set {
			e.FieldStart("id")
			s.ID.Encode(e)
		}
	}
	{
		if s.Price.Set {
			e.FieldStart("price")
			s.Price.Encode(e)
		}
	}
	{
		if s.Slug.Set {
			e.FieldStart("slug")
			s.Slug.Encode(e)
		}
	}
	{
		if s.Title.Set {
			e.FieldStart("title")
			s.Title.Encode(e)
		}
	}
}

var jsonFieldsNameOfRealmGift = [7]string{
	0: "currency",
	1: "icon",
	2: "icon_thumbnail",
	3: "id",
	4: "price",
	5: "slug",
	6: "title",
}

// Decode decodes RealmGift from json.
func (s *RealmGift) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode RealmGift to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "currency":
			if err := func() error {
				s.Currency.Reset()
				if err := s.Currency.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"currency\"")
			}
		case "icon":
			if err := func() error {
				s.Icon.Reset()
				if err := s.Icon.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"icon\"")
			}
		case "icon_thumbnail":
			if err := func() error {
				s.IconThumbnail.Reset()
				if err := s.IconThumbnail.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"icon_thumbnail\"")
			}
		case "id":
			if err := func() error {
				s.ID.Reset()
				if err := s.ID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"id\"")
			}
		case "price":
			if err := func() error {
				s.Price.Reset()
				if err := s.Price.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"price\"")
			}
		case "slug":
			if err := func() error {
				s.Slug.Reset()
				if err := s.Slug.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"slug\"")
			}
		case "title":
			if err := func() error {
				s.Title.Reset()
				if err := s.Title.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"title\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode RealmGift")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *RealmGift) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *RealmGift) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *RealmGiftingAbility) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *RealmGiftingAbility) encodeFields(e *jx.Encoder) {
	{
		if s.CanReceive.Set {
			e.FieldStart("can_receive")
			s.CanReceive.Encode(e)
		}
	}
	{
		if s.CanSend.Set {
			e.FieldStart("can_send")
			s.CanSend.Encode(e)
		}
	}
	{
		if s.Enabled.Set {
			e.FieldStart("enabled")
			s.Enabled.Encode(e)
		}
	}
	{
		if s.UserID.Set {
			e.FieldStart("user_id")
			s.UserID.Encode(e)
		}
	}
}

var jsonFieldsNameOfRealmGiftingAbility = [4]string{
	0: "can_receive",
	1: "can_send",
	2: "enabled",
	3: "user_id",
}

// Decode decodes RealmGiftingAbility from json.
func (s *RealmGiftingAbility) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode RealmGiftingAbility to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "can_receive":
			if err := func() error {
				s.CanReceive.Reset()
				if err := s.CanReceive.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"can_receive\"")
			}
		case "can_send":
			if err := func() error {
				s.CanSend.Reset()
				if err := s.CanSend.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"can_send\"")
			}
		case "enabled":
			if err := func() error {
				s.Enabled.Reset()
				if err := s.Enabled.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"enabled\"")
			}
		case "user_id":
			if err := func() error {
				s.UserID.Reset()
				if err := s.UserID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"user_id\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode RealmGiftingAbility")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *RealmGiftingAbility) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *RealmGiftingAbility) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *RealmMessage) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *RealmMessage) encodeFields(e *jx.Encoder) {
	{
		if s.Attachment.Set {
			e.FieldStart("attachment")
			s.Attachment.Encode(e)
		}
	}
	{
		if s.AttachmentAndroid.Set {
			e.FieldStart("attachment_android")
			s.AttachmentAndroid.Encode(e)
		}
	}
	{
		if s.AttachmentReadCount.Set {
			e.FieldStart("attachment_read_count")
			s.AttachmentReadCount.Encode(e)
		}
	}
	{
		if s.AttachmentThumbnail.Set {
			e.FieldStart("attachment_thumbnail")
			s.AttachmentThumbnail.Encode(e)
		}
	}
	{
		if s.ConferenceCall.Set {
			e.FieldStart("conference_call")
			s.ConferenceCall.Encode(e)
		}
	}
	{
		if s.CreatedAt.Set {
			e.FieldStart("created_at")
			s.CreatedAt.Encode(e)
		}
	}
	{
		if s.FontSize.Set {
			e.FieldStart("font_size")
			s.FontSize.Encode(e)
		}
	}
	{
		if s.GIF.Set {
			e.FieldStart("gif")
			s.GIF.Encode(e)
		}
	}
	{
		if s.ID.Set {
			e.FieldStart("id")
			s.ID.Encode(e)
		}
	}
	{
		if s.Invitation.Set {
			e.FieldStart("invitation")
			s.Invitation.Encode(e)
		}
	}
	{
		if s.IsError.Set {
			e.FieldStart("is_error")
			s.IsError.Encode(e)
		}
	}
	{
		if s.IsSent.Set {
			e.FieldStart("is_sent")
			s.IsSent.Encode(e)
		}
	}
	{
		if s.MessageType.Set {
			e.FieldStart("message_type")
			s.MessageType.Encode(e)
		}
	}
	{
		if s.Parent.Set {
			e.FieldStart("parent")
			s.Parent.Encode(e)
		}
	}
	{
		if s.RefreshRetryCount.Set {
			e.FieldStart("refresh_retry_count")
			s.RefreshRetryCount.Encode(e)
		}
	}
	{
		if s.RoomID.Set {
			e.FieldStart("room_id")
			s.RoomID.Encode(e)
		}
	}
	{
		if s.Sticker.Set {
			e.FieldStart("sticker")
			s.Sticker.Encode(e)
		}
	}
	{
		if s.Text.Set {
			e.FieldStart("text")
			s.Text.Encode(e)
		}
	}
	{
		if s.UserID.Set {
			e.FieldStart("user_id")
			s.UserID.Encode(e)
		}
	}
	{
		if s.VideoProcessed.Set {
			e.FieldStart("video_processed")
			s.VideoProcessed.Encode(e)
		}
	}
	{
		if s.VideoThumbnailURL.Set {
			e.FieldStart("video_thumbnail_url")
			s.VideoThumbnailURL.Encode(e)
		}
	}
	{
		if s.VideoURL.Set {
			e.FieldStart("video_url")
			s.VideoURL.Encode(e)
		}
	}
}

var jsonFieldsNameOfRealmMessage = [22]string{
	0:  "attachment",
	1:  "attachment_android",
	2:  "attachment_read_count",
	3:  "attachment_thumbnail",
	4:  "conference_call",
	5:  "created_at",
	6:  "font_size",
	7:  "gif",
	8:  "id",
	9:  "invitation",
	10: "is_error",
	11: "is_sent",
	12: "message_type",
	13: "parent",
	14: "refresh_retry_count",
	15: "room_id",
	16: "sticker",
	17: "text",
	18: "user_id",
	19: "video_processed",
	20: "video_thumbnail_url",
	21: "video_url",
}

// Decode decodes RealmMessage from json.
func (s *RealmMessage) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode RealmMessage to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "attachment":
			if err := func() error {
				s.Attachment.Reset()
				if err := s.Attachment.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"attachment\"")
			}
		case "attachment_android":
			if err := func() error {
				s.AttachmentAndroid.Reset()
				if err := s.AttachmentAndroid.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"attachment_android\"")
			}
		case "attachment_read_count":
			if err := func() error {
				s.AttachmentReadCount.Reset()
				if err := s.AttachmentReadCount.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"attachment_read_count\"")
			}
		case "attachment_thumbnail":
			if err := func() error {
				s.AttachmentThumbnail.Reset()
				if err := s.AttachmentThumbnail.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"attachment_thumbnail\"")
			}
		case "conference_call":
			if err := func() error {
				s.ConferenceCall.Reset()
				if err := s.ConferenceCall.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"conference_call\"")
			}
		case "created_at":
			if err := func() error {
				s.CreatedAt.Reset()
				if err := s.CreatedAt.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"created_at\"")
			}
		case "font_size":
			if err := func() error {
				s.FontSize.Reset()
				if err := s.FontSize.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"font_size\"")
			}
		case "gif":
			if err := func() error {
				s.GIF.Reset()
				if err := s.GIF.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"gif\"")
			}
		case "id":
			if err := func() error {
				s.ID.Reset()
				if err := s.ID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"id\"")
			}
		case "invitation":
			if err := func() error {
				s.Invitation.Reset()
				if err := s.Invitation.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"invitation\"")
			}
		case "is_error":
			if err := func() error {
				s.IsError.Reset()
				if err := s.IsError.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"is_error\"")
			}
		case "is_sent":
			if err := func() error {
				s.IsSent.Reset()
				if err := s.IsSent.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"is_sent\"")
			}
		case "message_type":
			if err := func() error {
				s.MessageType.Reset()
				if err := s.MessageType.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"message_type\"")
			}
		case "parent":
			if err := func() error {
				s.Parent.Reset()
				if err := s.Parent.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"parent\"")
			}
		case "refresh_retry_count":
			if err := func() error {
				s.RefreshRetryCount.Reset()
				if err := s.RefreshRetryCount.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"refresh_retry_count\"")
			}
		case "room_id":
			if err := func() error {
				s.RoomID.Reset()
				if err := s.RoomID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"room_id\"")
			}
		case "sticker":
			if err := func() error {
				s.Sticker.Reset()
				if err := s.Sticker.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"sticker\"")
			}
		case "text":
			if err := func() error {
				s.Text.Reset()
				if err := s.Text.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"text\"")
			}
		case "user_id":
			if err := func() error {
				s.UserID.Reset()
				if err := s.UserID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"user_id\"")
			}
		case "video_processed":
			if err := func() error {
				s.VideoProcessed.Reset()
				if err := s.VideoProcessed.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"video_processed\"")
			}
		case "video_thumbnail_url":
			if err := func() error {
				s.VideoThumbnailURL.Reset()
				if err := s.VideoThumbnailURL.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"video_thumbnail_url\"")
			}
		case "video_url":
			if err := func() error {
				s.VideoURL.Reset()
				if err := s.VideoURL.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"video_url\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode RealmMessage")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *RealmMessage) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *RealmMessage) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *RealmPlatformDetails) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *RealmPlatformDetails) encodeFields(e *jx.Encoder) {
	{
		if s.AffiliateURL.Set {
			e.FieldStart("affiliate_url")
			s.AffiliateURL.Encode(e)
		}
	}
	{
		if s.PackageID.Set {
			e.FieldStart("package_id")
			s.PackageID.Encode(e)
		}
	}
}

var jsonFieldsNameOfRealmPlatformDetails = [2]string{
	0: "affiliate_url",
	1: "package_id",
}

// Decode decodes RealmPlatformDetails from json.
func (s *RealmPlatformDetails) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode RealmPlatformDetails to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "affiliate_url":
			if err := func() error {
				s.AffiliateURL.Reset()
				if err := s.AffiliateURL.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"affiliate_url\"")
			}
		case "package_id":
			if err := func() error {
				s.PackageID.Reset()
				if err := s.PackageID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"package_id\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode RealmPlatformDetails")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *RealmPlatformDetails) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *RealmPlatformDetails) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *RealmUser) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *RealmUser) encodeFields(e *jx.Encoder) {
	{
		if s.AgeVerified.Set {
			e.FieldStart("age_verified")
			s.AgeVerified.Encode(e)
		}
	}
	{
		if s.AppleConnected.Set {
			e.FieldStart("apple_connected")
			s.AppleConnected.Encode(e)
		}
	}
	{
		if s.Biography.Set {
			e.FieldStart("biography")
			s.Biography.Encode(e)
		}
	}
	{
		if s.BirthDate.Set {
			e.FieldStart("birth_date")
			s.BirthDate.Encode(e)
		}
	}
	{
		if s.BlockingLimit.Set {
			e.FieldStart("blocking_limit")
			s.BlockingLimit.Encode(e)
		}
	}
	{
		if s.ChatRequest.Set {
			e.FieldStart("chat_request")
			s.ChatRequest.Encode(e)
		}
	}
	{
		if s.ConnectedBy.Set {
			e.FieldStart("connected_by")
			s.ConnectedBy.Encode(e)
		}
	}
	{
		if s.ContactPhones.Set {
			e.FieldStart("contact_phones")
			s.ContactPhones.Encode(e)
		}
	}
	{
		if s.CountryCode.Set {
			e.FieldStart("country_code")
			s.CountryCode.Encode(e)
		}
	}
	{
		if s.CoverImage.Set {
			e.FieldStart("cover_image")
			s.CoverImage.Encode(e)
		}
	}
	{
		if s.CoverImageThumbnail.Set {
			e.FieldStart("cover_image_thumbnail")
			s.CoverImageThumbnail.Encode(e)
		}
	}
	{
		if s.CreatedAtSeconds.Set {
			e.FieldStart("created_at_seconds")
			s.CreatedAtSeconds.Encode(e)
		}
	}
	{
		if s.DangerousUser.Set {
			e.FieldStart("dangerous_user")
			s.DangerousUser.Encode(e)
		}
	}
	{
		if s.EmailConfirmed.Set {
			e.FieldStart("email_confirmed")
			s.EmailConfirmed.Encode(e)
		}
	}
	{
		if s.FacebookConnected.Set {
			e.FieldStart("facebook_connected")
			s.FacebookConnected.Encode(e)
		}
	}
	{
		if s.FollowPending.Set {
			e.FieldStart("follow_pending")
			s.FollowPending.Encode(e)
		}
	}
	{
		if s.FollowedBy.Set {
			e.FieldStart("followed_by")
			s.FollowedBy.Encode(e)
		}
	}
	{
		if s.FollowedByPending.Set {
			e.FieldStart("followed_by_pending")
			s.FollowedByPending.Encode(e)
		}
	}
	{
		if s.FollowersCount.Set {
			e.FieldStart("followers_count")
			s.FollowersCount.Encode(e)
		}
	}
	{
		if s.Following.Set {
			e.FieldStart("following")
			s.Following.Encode(e)
		}
	}
	{
		if s.FollowingsCount.Set {
			e.FieldStart("followings_count")
			s.FollowingsCount.Encode(e)
		}
	}
	{
		if s.Frame.Set {
			e.FieldStart("frame")
			s.Frame.Encode(e)
		}
	}
	{
		if s.FrameThumbnail.Set {
			e.FieldStart("frame_thumbnail")
			s.FrameThumbnail.Encode(e)
		}
	}
	{
		if s.FromDifferentGenerationAndTrusted.Set {
			e.FieldStart("from_different_generation_and_trusted")
			s.FromDifferentGenerationAndTrusted.Encode(e)
		}
	}
	{
		if s.Gender.Set {
			e.FieldStart("gender")
			s.Gender.Encode(e)
		}
	}
	{
		if s.Generation.Set {
			e.FieldStart("generation")
			s.Generation.Encode(e)
		}
	}
	{
		if s.GiftingAbility.Set {
			e.FieldStart("gifting_ability")
			s.GiftingAbility.Encode(e)
		}
	}
	{
		if s.GoogleConnected.Set {
			e.FieldStart("google_connected")
			s.GoogleConnected.Encode(e)
		}
	}
	{
		if s.GroupPhoneOn.Set {
			e.FieldStart("group_phone_on")
			s.GroupPhoneOn.Encode(e)
		}
	}
	{
		if s.GroupUser.Set {
			e.FieldStart("group_user")
			s.GroupUser.Encode(e)
		}
	}
	{
		if s.GroupVideoOn.Set {
			e.FieldStart("group_video_on")
			s.GroupVideoOn.Encode(e)
		}
	}
	{
		if s.GroupsUsersCount.Set {
			e.FieldStart("groups_users_count")
			s.GroupsUsersCount.Encode(e)
		}
	}
	{
		if s.Hidden.Set {
			e.FieldStart("hidden")
			s.Hidden.Encode(e)
		}
	}
	{
		if s.HideVip.Set {
			e.FieldStart("hide_vip")
			s.HideVip.Encode(e)
		}
	}
	{
		if s.ID.Set {
			e.FieldStart("id")
			s.ID.Encode(e)
		}
	}
	{
		if s.InterestsCount.Set {
			e.FieldStart("interests_count")
			s.InterestsCount.Encode(e)
		}
	}
	{
		if s.InterestsSelected.Set {
			e.FieldStart("interests_selected")
			s.InterestsSelected.Encode(e)
		}
	}
	{
		if s.IsPrivate.Set {
			e.FieldStart("is_private")
			s.IsPrivate.Encode(e)
		}
	}
	{
		if s.LastLoggedInAtSeconds.Set {
			e.FieldStart("last_logged_in_at_seconds")
			s.LastLoggedInAtSeconds.Encode(e)
		}
	}
	{
		if s.LineConnected.Set {
			e.FieldStart("line_connected")
			s.LineConnected.Encode(e)
		}
	}
	{
		if s.LoginStreakCount.Set {
			e.FieldStart("login_streak_count")
			s.LoginStreakCount.Encode(e)
		}
	}
	{
		if s.MaskedEmail.Set {
			e.FieldStart("masked_email")
			s.MaskedEmail.Encode(e)
		}
	}
	{
		if s.MediaCount.Set {
			e.FieldStart("media_count")
			s.MediaCount.Encode(e)
		}
	}
	{
		if s.MobileNumber.Set {
			e.FieldStart("mobile_number")
			s.MobileNumber.Encode(e)
		}
	}
	{
		if s.NewUser.Set {
			e.FieldStart("new_user")
			s.NewUser.Encode(e)
		}
	}
	{
		if s.Nickname.Set {
			e.FieldStart("nickname")
			s.Nickname.Encode(e)
		}
	}
	{
		if s.OnlineStatus.Set {
			e.FieldStart("online_status")
			s.OnlineStatus.Encode(e)
		}
	}
	{
		if s.PhoneOn.Set {
			e.FieldStart("phone_on")
			s.PhoneOn.Encode(e)
		}
	}
	{
		if s.PostsCount.Set {
			e.FieldStart("posts_count")
			s.PostsCount.Encode(e)
		}
	}
	{
		if s.Prefecture.Set {
			e.FieldStart("prefecture")
			s.Prefecture.Encode(e)
		}
	}
	{
		if s.ProfileIcon.Set {
			e.FieldStart("profile_icon")
			s.ProfileIcon.Encode(e)
		}
	}
	{
		if s.ProfileIconThumbnail.Set {
			e.FieldStart("profile_icon_thumbnail")
			s.ProfileIconThumbnail.Encode(e)
		}
	}
	{
		if s.PushNotification.Set {
			e.FieldStart("push_notification")
			s.PushNotification.Encode(e)
		}
	}
	{
		if s.RecentlyKenta.Set {
			e.FieldStart("recently_kenta")
			s.RecentlyKenta.Encode(e)
		}
	}
	{
		if s.RestrictedReviewBy.Set {
			e.FieldStart("restricted_review_by")
			s.RestrictedReviewBy.Encode(e)
		}
	}
	{
		if s.ReviewsCount.Set {
			e.FieldStart("reviews_count")
			s.ReviewsCount.Encode(e)
		}
	}
	{
		if s.Title.Set {
			e.FieldStart("title")
			s.Title.Encode(e)
		}
	}
	{
		if s.TotalGiftsReceivedCount.Set {
			e.FieldStart("total_gifts_received_count")
			s.TotalGiftsReceivedCount.Encode(e)
		}
	}
	{
		if s.TwitterID.Set {
			e.FieldStart("twitter_id")
			s.TwitterID.Encode(e)
		}
	}
	{
		if s.TwoFaEnabled.Set {
			e.FieldStart("two_fa_enabled")
			s.TwoFaEnabled.Encode(e)
		}
	}
	{
		if s.UpdatedTimeMillis.Set {
			e.FieldStart("updated_time_millis")
			s.UpdatedTimeMillis.Encode(e)
		}
	}
	{
		if s.Username.Set {
			e.FieldStart("username")
			s.Username.Encode(e)
		}
	}
	{
		if s.UUID.Set {
			e.FieldStart("uuid")
			s.UUID.Encode(e)
		}
	}
	{
		if s.VideoOn.Set {
			e.FieldStart("video_on")
			s.VideoOn.Encode(e)
		}
	}
	{
		if s.Vip.Set {
			e.FieldStart("vip")
			s.Vip.Encode(e)
		}
	}
	{
		if s.VipUntilSeconds.Set {
			e.FieldStart("vip_until_seconds")
			s.VipUntilSeconds.Encode(e)
		}
	}
	{
		if s.WorldIDConnected.Set {
			e.FieldStart("world_id_connected")
			s.WorldIDConnected.Encode(e)
		}
	}
}

var jsonFieldsNameOfRealmUser = [67]string{
	0:  "age_verified",
	1:  "apple_connected",
	2:  "biography",
	3:  "birth_date",
	4:  "blocking_limit",
	5:  "chat_request",
	6:  "connected_by",
	7:  "contact_phones",
	8:  "country_code",
	9:  "cover_image",
	10: "cover_image_thumbnail",
	11: "created_at_seconds",
	12: "dangerous_user",
	13: "email_confirmed",
	14: "facebook_connected",
	15: "follow_pending",
	16: "followed_by",
	17: "followed_by_pending",
	18: "followers_count",
	19: "following",
	20: "followings_count",
	21: "frame",
	22: "frame_thumbnail",
	23: "from_different_generation_and_trusted",
	24: "gender",
	25: "generation",
	26: "gifting_ability",
	27: "google_connected",
	28: "group_phone_on",
	29: "group_user",
	30: "group_video_on",
	31: "groups_users_count",
	32: "hidden",
	33: "hide_vip",
	34: "id",
	35: "interests_count",
	36: "interests_selected",
	37: "is_private",
	38: "last_logged_in_at_seconds",
	39: "line_connected",
	40: "login_streak_count",
	41: "masked_email",
	42: "media_count",
	43: "mobile_number",
	44: "new_user",
	45: "nickname",
	46: "online_status",
	47: "phone_on",
	48: "posts_count",
	49: "prefecture",
	50: "profile_icon",
	51: "profile_icon_thumbnail",
	52: "push_notification",
	53: "recently_kenta",
	54: "restricted_review_by",
	55: "reviews_count",
	56: "title",
	57: "total_gifts_received_count",
	58: "twitter_id",
	59: "two_fa_enabled",
	60: "updated_time_millis",
	61: "username",
	62: "uuid",
	63: "video_on",
	64: "vip",
	65: "vip_until_seconds",
	66: "world_id_connected",
}

// Decode decodes RealmUser from json.
func (s *RealmUser) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode RealmUser to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "age_verified":
			if err := func() error {
				s.AgeVerified.Reset()
				if err := s.AgeVerified.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"age_verified\"")
			}
		case "apple_connected":
			if err := func() error {
				s.AppleConnected.Reset()
				if err := s.AppleConnected.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"apple_connected\"")
			}
		case "biography":
			if err := func() error {
				s.Biography.Reset()
				if err := s.Biography.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"biography\"")
			}
		case "birth_date":
			if err := func() error {
				s.BirthDate.Reset()
				if err := s.BirthDate.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"birth_date\"")
			}
		case "blocking_limit":
			if err := func() error {
				s.BlockingLimit.Reset()
				if err := s.BlockingLimit.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"blocking_limit\"")
			}
		case "chat_request":
			if err := func() error {
				s.ChatRequest.Reset()
				if err := s.ChatRequest.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"chat_request\"")
			}
		case "connected_by":
			if err := func() error {
				s.ConnectedBy.Reset()
				if err := s.ConnectedBy.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"connected_by\"")
			}
		case "contact_phones":
			if err := func() error {
				s.ContactPhones.Reset()
				if err := s.ContactPhones.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"contact_phones\"")
			}
		case "country_code":
			if err := func() error {
				s.CountryCode.Reset()
				if err := s.CountryCode.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"country_code\"")
			}
		case "cover_image":
			if err := func() error {
				s.CoverImage.Reset()
				if err := s.CoverImage.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"cover_image\"")
			}
		case "cover_image_thumbnail":
			if err := func() error {
				s.CoverImageThumbnail.Reset()
				if err := s.CoverImageThumbnail.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"cover_image_thumbnail\"")
			}
		case "created_at_seconds":
			if err := func() error {
				s.CreatedAtSeconds.Reset()
				if err := s.CreatedAtSeconds.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"created_at_seconds\"")
			}
		case "dangerous_user":
			if err := func() error {
				s.DangerousUser.Reset()
				if err := s.DangerousUser.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"dangerous_user\"")
			}
		case "email_confirmed":
			if err := func() error {
				s.EmailConfirmed.Reset()
				if err := s.EmailConfirmed.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"email_confirmed\"")
			}
		case "facebook_connected":
			if err := func() error {
				s.FacebookConnected.Reset()
				if err := s.FacebookConnected.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"facebook_connected\"")
			}
		case "follow_pending":
			if err := func() error {
				s.FollowPending.Reset()
				if err := s.FollowPending.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"follow_pending\"")
			}
		case "followed_by":
			if err := func() error {
				s.FollowedBy.Reset()
				if err := s.FollowedBy.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"followed_by\"")
			}
		case "followed_by_pending":
			if err := func() error {
				s.FollowedByPending.Reset()
				if err := s.FollowedByPending.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"followed_by_pending\"")
			}
		case "followers_count":
			if err := func() error {
				s.FollowersCount.Reset()
				if err := s.FollowersCount.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"followers_count\"")
			}
		case "following":
			if err := func() error {
				s.Following.Reset()
				if err := s.Following.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"following\"")
			}
		case "followings_count":
			if err := func() error {
				s.FollowingsCount.Reset()
				if err := s.FollowingsCount.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"followings_count\"")
			}
		case "frame":
			if err := func() error {
				s.Frame.Reset()
				if err := s.Frame.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"frame\"")
			}
		case "frame_thumbnail":
			if err := func() error {
				s.FrameThumbnail.Reset()
				if err := s.FrameThumbnail.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"frame_thumbnail\"")
			}
		case "from_different_generation_and_trusted":
			if err := func() error {
				s.FromDifferentGenerationAndTrusted.Reset()
				if err := s.FromDifferentGenerationAndTrusted.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"from_different_generation_and_trusted\"")
			}
		case "gender":
			if err := func() error {
				s.Gender.Reset()
				if err := s.Gender.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"gender\"")
			}
		case "generation":
			if err := func() error {
				s.Generation.Reset()
				if err := s.Generation.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"generation\"")
			}
		case "gifting_ability":
			if err := func() error {
				s.GiftingAbility.Reset()
				if err := s.GiftingAbility.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"gifting_ability\"")
			}
		case "google_connected":
			if err := func() error {
				s.GoogleConnected.Reset()
				if err := s.GoogleConnected.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"google_connected\"")
			}
		case "group_phone_on":
			if err := func() error {
				s.GroupPhoneOn.Reset()
				if err := s.GroupPhoneOn.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"group_phone_on\"")
			}
		case "group_user":
			if err := func() error {
				s.GroupUser.Reset()
				if err := s.GroupUser.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"group_user\"")
			}
		case "group_video_on":
			if err := func() error {
				s.GroupVideoOn.Reset()
				if err := s.GroupVideoOn.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"group_video_on\"")
			}
		case "groups_users_count":
			if err := func() error {
				s.GroupsUsersCount.Reset()
				if err := s.GroupsUsersCount.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"groups_users_count\"")
			}
		case "hidden":
			if err := func() error {
				s.Hidden.Reset()
				if err := s.Hidden.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"hidden\"")
			}
		case "hide_vip":
			if err := func() error {
				s.HideVip.Reset()
				if err := s.HideVip.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"hide_vip\"")
			}
		case "id":
			if err := func() error {
				s.ID.Reset()
				if err := s.ID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"id\"")
			}
		case "interests_count":
			if err := func() error {
				s.InterestsCount.Reset()
				if err := s.InterestsCount.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"interests_count\"")
			}
		case "interests_selected":
			if err := func() error {
				s.InterestsSelected.Reset()
				if err := s.InterestsSelected.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"interests_selected\"")
			}
		case "is_private":
			if err := func() error {
				s.IsPrivate.Reset()
				if err := s.IsPrivate.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"is_private\"")
			}
		case "last_logged_in_at_seconds":
			if err := func() error {
				s.LastLoggedInAtSeconds.Reset()
				if err := s.LastLoggedInAtSeconds.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"last_logged_in_at_seconds\"")
			}
		case "line_connected":
			if err := func() error {
				s.LineConnected.Reset()
				if err := s.LineConnected.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"line_connected\"")
			}
		case "login_streak_count":
			if err := func() error {
				s.LoginStreakCount.Reset()
				if err := s.LoginStreakCount.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"login_streak_count\"")
			}
		case "masked_email":
			if err := func() error {
				s.MaskedEmail.Reset()
				if err := s.MaskedEmail.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"masked_email\"")
			}
		case "media_count":
			if err := func() error {
				s.MediaCount.Reset()
				if err := s.MediaCount.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"media_count\"")
			}
		case "mobile_number":
			if err := func() error {
				s.MobileNumber.Reset()
				if err := s.MobileNumber.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"mobile_number\"")
			}
		case "new_user":
			if err := func() error {
				s.NewUser.Reset()
				if err := s.NewUser.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"new_user\"")
			}
		case "nickname":
			if err := func() error {
				s.Nickname.Reset()
				if err := s.Nickname.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"nickname\"")
			}
		case "online_status":
			if err := func() error {
				s.OnlineStatus.Reset()
				if err := s.OnlineStatus.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"online_status\"")
			}
		case "phone_on":
			if err := func() error {
				s.PhoneOn.Reset()
				if err := s.PhoneOn.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"phone_on\"")
			}
		case "posts_count":
			if err := func() error {
				s.PostsCount.Reset()
				if err := s.PostsCount.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"posts_count\"")
			}
		case "prefecture":
			if err := func() error {
				s.Prefecture.Reset()
				if err := s.Prefecture.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"prefecture\"")
			}
		case "profile_icon":
			if err := func() error {
				s.ProfileIcon.Reset()
				if err := s.ProfileIcon.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"profile_icon\"")
			}
		case "profile_icon_thumbnail":
			if err := func() error {
				s.ProfileIconThumbnail.Reset()
				if err := s.ProfileIconThumbnail.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"profile_icon_thumbnail\"")
			}
		case "push_notification":
			if err := func() error {
				s.PushNotification.Reset()
				if err := s.PushNotification.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"push_notification\"")
			}
		case "recently_kenta":
			if err := func() error {
				s.RecentlyKenta.Reset()
				if err := s.RecentlyKenta.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"recently_kenta\"")
			}
		case "restricted_review_by":
			if err := func() error {
				s.RestrictedReviewBy.Reset()
				if err := s.RestrictedReviewBy.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"restricted_review_by\"")
			}
		case "reviews_count":
			if err := func() error {
				s.ReviewsCount.Reset()
				if err := s.ReviewsCount.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"reviews_count\"")
			}
		case "title":
			if err := func() error {
				s.Title.Reset()
				if err := s.Title.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"title\"")
			}
		case "total_gifts_received_count":
			if err := func() error {
				s.TotalGiftsReceivedCount.Reset()
				if err := s.TotalGiftsReceivedCount.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"total_gifts_received_count\"")
			}
		case "twitter_id":
			if err := func() error {
				s.TwitterID.Reset()
				if err := s.TwitterID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"twitter_id\"")
			}
		case "two_fa_enabled":
			if err := func() error {
				s.TwoFaEnabled.Reset()
				if err := s.TwoFaEnabled.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"two_fa_enabled\"")
			}
		case "updated_time_millis":
			if err := func() error {
				s.UpdatedTimeMillis.Reset()
				if err := s.UpdatedTimeMillis.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"updated_time_millis\"")
			}
		case "username":
			if err := func() error {
				s.Username.Reset()
				if err := s.Username.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"username\"")
			}
		case "uuid":
			if err := func() error {
				s.UUID.Reset()
				if err := s.UUID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"uuid\"")
			}
		case "video_on":
			if err := func() error {
				s.VideoOn.Reset()
				if err := s.VideoOn.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"video_on\"")
			}
		case "vip":
			if err := func() error {
				s.Vip.Reset()
				if err := s.Vip.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"vip\"")
			}
		case "vip_until_seconds":
			if err := func() error {
				s.VipUntilSeconds.Reset()
				if err := s.VipUntilSeconds.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"vip_until_seconds\"")
			}
		case "world_id_connected":
			if err := func() error {
				s.WorldIDConnected.Reset()
				if err := s.WorldIDConnected.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"world_id_connected\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode RealmUser")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *RealmUser) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *RealmUser) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *ReceivedGift) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *ReceivedGift) encodeFields(e *jx.Encoder) {
	{
		if s.Gift.Set {
			e.FieldStart("gift")
			s.Gift.Encode(e)
		}
	}
	{
		if s.ReceivedCount.Set {
			e.FieldStart("received_count")
			s.ReceivedCount.Encode(e)
		}
	}
	{
		if s.Senders.Set {
			e.FieldStart("senders")
			s.Senders.Encode(e)
		}
	}
	{
		if s.TotalSendersCount.Set {
			e.FieldStart("total_senders_count")
			s.TotalSendersCount.Encode(e)
		}
	}
}

var jsonFieldsNameOfReceivedGift = [4]string{
	0: "gift",
	1: "received_count",
	2: "senders",
	3: "total_senders_count",
}

// Decode decodes ReceivedGift from json.
func (s *ReceivedGift) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode ReceivedGift to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "gift":
			if err := func() error {
				s.Gift.Reset()
				if err := s.Gift.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"gift\"")
			}
		case "received_count":
			if err := func() error {
				s.ReceivedCount.Reset()
				if err := s.ReceivedCount.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"received_count\"")
			}
		case "senders":
			if err := func() error {
				s.Senders.Reset()
				if err := s.Senders.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"senders\"")
			}
		case "total_senders_count":
			if err := func() error {
				s.TotalSendersCount.Reset()
				if err := s.TotalSendersCount.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"total_senders_count\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode ReceivedGift")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *ReceivedGift) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *ReceivedGift) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *RefreshCounterRequest) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *RefreshCounterRequest) encodeFields(e *jx.Encoder) {
	{
		if s.Counter.Set {
			e.FieldStart("counter")
			s.Counter.Encode(e)
		}
	}
	{
		if s.LastRequestedAt.Set {
			e.FieldStart("last_requested_at")
			s.LastRequestedAt.Encode(e)
		}
	}
	{
		if s.Status.Set {
			e.FieldStart("status")
			s.Status.Encode(e)
		}
	}
}

var jsonFieldsNameOfRefreshCounterRequest = [3]string{
	0: "counter",
	1: "last_requested_at",
	2: "status",
}

// Decode decodes RefreshCounterRequest from json.
func (s *RefreshCounterRequest) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode RefreshCounterRequest to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "counter":
			if err := func() error {
				s.Counter.Reset()
				if err := s.Counter.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"counter\"")
			}
		case "last_requested_at":
			if err := func() error {
				s.LastRequestedAt.Reset()
				if err := s.LastRequestedAt.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"last_requested_at\"")
			}
		case "status":
			if err := func() error {
				s.Status.Reset()
				if err := s.Status.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"status\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode RefreshCounterRequest")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *RefreshCounterRequest) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *RefreshCounterRequest) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *RefreshCounterRequestsResponse) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *RefreshCounterRequestsResponse) encodeFields(e *jx.Encoder) {
	{
		if s.ResetCounterRequests.Set {
			e.FieldStart("reset_counter_requests")
			s.ResetCounterRequests.Encode(e)
		}
	}
}

var jsonFieldsNameOfRefreshCounterRequestsResponse = [1]string{
	0: "reset_counter_requests",
}

// Decode decodes RefreshCounterRequestsResponse from json.
func (s *RefreshCounterRequestsResponse) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode RefreshCounterRequestsResponse to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "reset_counter_requests":
			if err := func() error {
				s.ResetCounterRequests.Reset()
				if err := s.ResetCounterRequests.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"reset_counter_requests\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode RefreshCounterRequestsResponse")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *RefreshCounterRequestsResponse) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *RefreshCounterRequestsResponse) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *RepostReqMessageTags) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *RepostReqMessageTags) encodeFields(e *jx.Encoder) {
}

var jsonFieldsNameOfRepostReqMessageTags = [0]string{}

// Decode decodes RepostReqMessageTags from json.
func (s *RepostReqMessageTags) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode RepostReqMessageTags to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		default:
			return d.Skip()
		}
	}); err != nil {
		return errors.Wrap(err, "decode RepostReqMessageTags")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *RepostReqMessageTags) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *RepostReqMessageTags) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *RepostReqSharedURL) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *RepostReqSharedURL) encodeFields(e *jx.Encoder) {
}

var jsonFieldsNameOfRepostReqSharedURL = [0]string{}

// Decode decodes RepostReqSharedURL from json.
func (s *RepostReqSharedURL) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode RepostReqSharedURL to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		default:
			return d.Skip()
		}
	}); err != nil {
		return errors.Wrap(err, "decode RepostReqSharedURL")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *RepostReqSharedURL) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *RepostReqSharedURL) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *Review) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *Review) encodeFields(e *jx.Encoder) {
	{
		if s.Comment.Set {
			e.FieldStart("comment")
			s.Comment.Encode(e)
		}
	}
	{
		if s.CreatedAt.Set {
			e.FieldStart("created_at")
			s.CreatedAt.Encode(e)
		}
	}
	{
		if s.ID.Set {
			e.FieldStart("id")
			s.ID.Encode(e)
		}
	}
	{
		if s.MutualReview.Set {
			e.FieldStart("mutual_review")
			s.MutualReview.Encode(e)
		}
	}
	{
		if s.ReportedCount.Set {
			e.FieldStart("reported_count")
			s.ReportedCount.Encode(e)
		}
	}
	{
		if s.Reviewer.Set {
			e.FieldStart("reviewer")
			s.Reviewer.Encode(e)
		}
	}
	{
		if s.User.Set {
			e.FieldStart("user")
			s.User.Encode(e)
		}
	}
}

var jsonFieldsNameOfReview = [7]string{
	0: "comment",
	1: "created_at",
	2: "id",
	3: "mutual_review",
	4: "reported_count",
	5: "reviewer",
	6: "user",
}

// Decode decodes Review from json.
func (s *Review) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Review to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "comment":
			if err := func() error {
				s.Comment.Reset()
				if err := s.Comment.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"comment\"")
			}
		case "created_at":
			if err := func() error {
				s.CreatedAt.Reset()
				if err := s.CreatedAt.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"created_at\"")
			}
		case "id":
			if err := func() error {
				s.ID.Reset()
				if err := s.ID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"id\"")
			}
		case "mutual_review":
			if err := func() error {
				s.MutualReview.Reset()
				if err := s.MutualReview.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"mutual_review\"")
			}
		case "reported_count":
			if err := func() error {
				s.ReportedCount.Reset()
				if err := s.ReportedCount.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"reported_count\"")
			}
		case "reviewer":
			if err := func() error {
				s.Reviewer.Reset()
				if err := s.Reviewer.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"reviewer\"")
			}
		case "user":
			if err := func() error {
				s.User.Reset()
				if err := s.User.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"user\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode Review")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *Review) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *Review) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *ReviewRestriction) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *ReviewRestriction) encodeFields(e *jx.Encoder) {
	{
		if s.APIValue.Set {
			e.FieldStart("api_value")
			s.APIValue.Encode(e)
		}
	}
}

var jsonFieldsNameOfReviewRestriction = [1]string{
	0: "api_value",
}

// Decode decodes ReviewRestriction from json.
func (s *ReviewRestriction) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode ReviewRestriction to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "api_value":
			if err := func() error {
				s.APIValue.Reset()
				if err := s.APIValue.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"api_value\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode ReviewRestriction")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *ReviewRestriction) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *ReviewRestriction) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *ReviewsResponse) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *ReviewsResponse) encodeFields(e *jx.Encoder) {
	{
		if s.PinnedReviews.Set {
			e.FieldStart("pinned_reviews")
			s.PinnedReviews.Encode(e)
		}
	}
	{
		if s.Reviews.Set {
			e.FieldStart("reviews")
			s.Reviews.Encode(e)
		}
	}
}

var jsonFieldsNameOfReviewsResponse = [2]string{
	0: "pinned_reviews",
	1: "reviews",
}

// Decode decodes ReviewsResponse from json.
func (s *ReviewsResponse) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode ReviewsResponse to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "pinned_reviews":
			if err := func() error {
				s.PinnedReviews.Reset()
				if err := s.PinnedReviews.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"pinned_reviews\"")
			}
		case "reviews":
			if err := func() error {
				s.Reviews.Reset()
				if err := s.Reviews.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"reviews\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode ReviewsResponse")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *ReviewsResponse) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *ReviewsResponse) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *RtmTokenResponse) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *RtmTokenResponse) encodeFields(e *jx.Encoder) {
	{
		if s.Token.Set {
			e.FieldStart("token")
			s.Token.Encode(e)
		}
	}
}

var jsonFieldsNameOfRtmTokenResponse = [1]string{
	0: "token",
}

// Decode decodes RtmTokenResponse from json.
func (s *RtmTokenResponse) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode RtmTokenResponse to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "token":
			if err := func() error {
				s.Token.Reset()
				if err := s.Token.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"token\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode RtmTokenResponse")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *RtmTokenResponse) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *RtmTokenResponse) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *SearchCriteria) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *SearchCriteria) encodeFields(e *jx.Encoder) {
	{
		if s.Biography.Set {
			e.FieldStart("biography")
			s.Biography.Encode(e)
		}
	}
	{
		if s.Gender.Set {
			e.FieldStart("gender")
			s.Gender.Encode(e)
		}
	}
	{
		if s.Nickname.Set {
			e.FieldStart("nickname")
			s.Nickname.Encode(e)
		}
	}
	{
		if s.Prefecture.Set {
			e.FieldStart("prefecture")
			s.Prefecture.Encode(e)
		}
	}
	{
		if s.Username.Set {
			e.FieldStart("username")
			s.Username.Encode(e)
		}
	}
}

var jsonFieldsNameOfSearchCriteria = [5]string{
	0: "biography",
	1: "gender",
	2: "nickname",
	3: "prefecture",
	4: "username",
}

// Decode decodes SearchCriteria from json.
func (s *SearchCriteria) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode SearchCriteria to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "biography":
			if err := func() error {
				s.Biography.Reset()
				if err := s.Biography.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"biography\"")
			}
		case "gender":
			if err := func() error {
				s.Gender.Reset()
				if err := s.Gender.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"gender\"")
			}
		case "nickname":
			if err := func() error {
				s.Nickname.Reset()
				if err := s.Nickname.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"nickname\"")
			}
		case "prefecture":
			if err := func() error {
				s.Prefecture.Reset()
				if err := s.Prefecture.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"prefecture\"")
			}
		case "username":
			if err := func() error {
				s.Username.Reset()
				if err := s.Username.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"username\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode SearchCriteria")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *SearchCriteria) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *SearchCriteria) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *SearchUsersRequest) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *SearchUsersRequest) encodeFields(e *jx.Encoder) {
	{
		if s.User.Set {
			e.FieldStart("user")
			s.User.Encode(e)
		}
	}
}

var jsonFieldsNameOfSearchUsersRequest = [1]string{
	0: "user",
}

// Decode decodes SearchUsersRequest from json.
func (s *SearchUsersRequest) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode SearchUsersRequest to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "user":
			if err := func() error {
				s.User.Reset()
				if err := s.User.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"user\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode SearchUsersRequest")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *SearchUsersRequest) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *SearchUsersRequest) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *Setting) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *Setting) encodeFields(e *jx.Encoder) {
	{
		if s.NotificationGroupJoin.Set {
			e.FieldStart("notification_group_join")
			s.NotificationGroupJoin.Encode(e)
		}
	}
	{
		if s.NotificationGroupMessageTagAll.Set {
			e.FieldStart("notification_group_message_tag_all")
			s.NotificationGroupMessageTagAll.Encode(e)
		}
	}
	{
		if s.NotificationGroupPost.Set {
			e.FieldStart("notification_group_post")
			s.NotificationGroupPost.Encode(e)
		}
	}
	{
		if s.NotificationGroupRequest.Set {
			e.FieldStart("notification_group_request")
			s.NotificationGroupRequest.Encode(e)
		}
	}
}

var jsonFieldsNameOfSetting = [4]string{
	0: "notification_group_join",
	1: "notification_group_message_tag_all",
	2: "notification_group_post",
	3: "notification_group_request",
}

// Decode decodes Setting from json.
func (s *Setting) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Setting to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "notification_group_join":
			if err := func() error {
				s.NotificationGroupJoin.Reset()
				if err := s.NotificationGroupJoin.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"notification_group_join\"")
			}
		case "notification_group_message_tag_all":
			if err := func() error {
				s.NotificationGroupMessageTagAll.Reset()
				if err := s.NotificationGroupMessageTagAll.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"notification_group_message_tag_all\"")
			}
		case "notification_group_post":
			if err := func() error {
				s.NotificationGroupPost.Reset()
				if err := s.NotificationGroupPost.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"notification_group_post\"")
			}
		case "notification_group_request":
			if err := func() error {
				s.NotificationGroupRequest.Reset()
				if err := s.NotificationGroupRequest.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"notification_group_request\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode Setting")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *Setting) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *Setting) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *Settings) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *Settings) encodeFields(e *jx.Encoder) {
	{
		if s.AgeRestrictedOnReview.Set {
			e.FieldStart("age_restricted_on_review")
			s.AgeRestrictedOnReview.Encode(e)
		}
	}
	{
		if s.AllowReposts.Set {
			e.FieldStart("allow_reposts")
			s.AllowReposts.Encode(e)
		}
	}
	{
		if s.CautionUserChat.Set {
			e.FieldStart("caution_user_chat")
			s.CautionUserChat.Encode(e)
		}
	}
	{
		if s.FollowingOnlyCallInvite.Set {
			e.FieldStart("following_only_call_invite")
			s.FollowingOnlyCallInvite.Encode(e)
		}
	}
	{
		if s.FollowingOnlyGroupInvite.Set {
			e.FieldStart("following_only_group_invite")
			s.FollowingOnlyGroupInvite.Encode(e)
		}
	}
	{
		if s.FollowingRestrictedOnReview.Set {
			e.FieldStart("following_restricted_on_review")
			s.FollowingRestrictedOnReview.Encode(e)
		}
	}
	{
		if s.HideActiveCall.Set {
			e.FieldStart("hide_active_call")
			s.HideActiveCall.Encode(e)
		}
	}
	{
		if s.HideGiftsReceived.Set {
			e.FieldStart("hide_gifts_received")
			s.HideGiftsReceived.Encode(e)
		}
	}
	{
		if s.HideHotPost.Set {
			e.FieldStart("hide_hot_post")
			s.HideHotPost.Encode(e)
		}
	}
	{
		if s.HideOnInvitable.Set {
			e.FieldStart("hide_on_invitable")
			s.HideOnInvitable.Encode(e)
		}
	}
	{
		if s.HideOnlineStatus.Set {
			e.FieldStart("hide_online_status")
			s.HideOnlineStatus.Encode(e)
		}
	}
	{
		if s.HideReplyFollowingTimeline.Set {
			e.FieldStart("hide_reply_following_timeline")
			s.HideReplyFollowingTimeline.Encode(e)
		}
	}
	{
		if s.HideReplyPublicTimeline.Set {
			e.FieldStart("hide_reply_public_timeline")
			s.HideReplyPublicTimeline.Encode(e)
		}
	}
	{
		if s.HideVip.Set {
			e.FieldStart("hide_vip")
			s.HideVip.Encode(e)
		}
	}
	{
		if s.InvisibleOnUserSearch.Set {
			e.FieldStart("invisible_on_user_search")
			s.InvisibleOnUserSearch.Encode(e)
		}
	}
	{
		if s.NoReplyGroupTimeline.Set {
			e.FieldStart("no_reply_group_timeline")
			s.NoReplyGroupTimeline.Encode(e)
		}
	}
	{
		if s.NotificationBirthdayToFollowers.Set {
			e.FieldStart("notification_birthday_to_followers")
			s.NotificationBirthdayToFollowers.Encode(e)
		}
	}
	{
		if s.NotificationBulkCallInvite.Set {
			e.FieldStart("notification_bulk_call_invite")
			s.NotificationBulkCallInvite.Encode(e)
		}
	}
	{
		if s.NotificationCallInvite.Set {
			e.FieldStart("notification_call_invite")
			s.NotificationCallInvite.Encode(e)
		}
	}
	{
		if s.NotificationChat.Set {
			e.FieldStart("notification_chat")
			s.NotificationChat.Encode(e)
		}
	}
	{
		if s.NotificationChatDelete.Set {
			e.FieldStart("notification_chat_delete")
			s.NotificationChatDelete.Encode(e)
		}
	}
	{
		if s.NotificationContactFriend.Set {
			e.FieldStart("notification_contact_friend")
			s.NotificationContactFriend.Encode(e)
		}
	}
	{
		if s.NotificationDailyQuest.Set {
			e.FieldStart("notification_daily_quest")
			s.NotificationDailyQuest.Encode(e)
		}
	}
	{
		if s.NotificationDailySummary.Set {
			e.FieldStart("notification_daily_summary")
			s.NotificationDailySummary.Encode(e)
		}
	}
	{
		if s.NotificationFollow.Set {
			e.FieldStart("notification_follow")
			s.NotificationFollow.Encode(e)
		}
	}
	{
		if s.NotificationFollowAccept.Set {
			e.FieldStart("notification_follow_accept")
			s.NotificationFollowAccept.Encode(e)
		}
	}
	{
		if s.NotificationFollowRequest.Set {
			e.FieldStart("notification_follow_request")
			s.NotificationFollowRequest.Encode(e)
		}
	}
	{
		if s.NotificationFollowerConferenceCall.Set {
			e.FieldStart("notification_follower_conference_call")
			s.NotificationFollowerConferenceCall.Encode(e)
		}
	}
	{
		if s.NotificationFollowerCreateGroup.Set {
			e.FieldStart("notification_follower_create_group")
			s.NotificationFollowerCreateGroup.Encode(e)
		}
	}
	{
		if s.NotificationFollowingBirthdateOn.Set {
			e.FieldStart("notification_following_birthdate_on")
			s.NotificationFollowingBirthdateOn.Encode(e)
		}
	}
	{
		if s.NotificationFollowingPostAfterBreak.Set {
			e.FieldStart("notification_following_post_after_break")
			s.NotificationFollowingPostAfterBreak.Encode(e)
		}
	}
	{
		if s.NotificationFollowingsInCall.Set {
			e.FieldStart("notification_followings_in_call")
			s.NotificationFollowingsInCall.Encode(e)
		}
	}
	{
		if s.NotificationFootprint.Set {
			e.FieldStart("notification_footprint")
			s.NotificationFootprint.Encode(e)
		}
	}
	{
		if s.NotificationGiftReceive.Set {
			e.FieldStart("notification_gift_receive")
			s.NotificationGiftReceive.Encode(e)
		}
	}
	{
		if s.NotificationGroupAccept.Set {
			e.FieldStart("notification_group_accept")
			s.NotificationGroupAccept.Encode(e)
		}
	}
	{
		if s.NotificationGroupConferenceCall.Set {
			e.FieldStart("notification_group_conference_call")
			s.NotificationGroupConferenceCall.Encode(e)
		}
	}
	{
		if s.NotificationGroupInvite.Set {
			e.FieldStart("notification_group_invite")
			s.NotificationGroupInvite.Encode(e)
		}
	}
	{
		if s.NotificationGroupJoin.Set {
			e.FieldStart("notification_group_join")
			s.NotificationGroupJoin.Encode(e)
		}
	}
	{
		if s.NotificationGroupMessageTagAll.Set {
			e.FieldStart("notification_group_message_tag_all")
			s.NotificationGroupMessageTagAll.Encode(e)
		}
	}
	{
		if s.NotificationGroupModerator.Set {
			e.FieldStart("notification_group_moderator")
			s.NotificationGroupModerator.Encode(e)
		}
	}
	{
		if s.NotificationGroupPost.Set {
			e.FieldStart("notification_group_post")
			s.NotificationGroupPost.Encode(e)
		}
	}
	{
		if s.NotificationGroupRequest.Set {
			e.FieldStart("notification_group_request")
			s.NotificationGroupRequest.Encode(e)
		}
	}
	{
		if s.NotificationHimaNow.Set {
			e.FieldStart("notification_hima_now")
			s.NotificationHimaNow.Encode(e)
		}
	}
	{
		if s.NotificationInviteeCreateAccount.Set {
			e.FieldStart("notification_invitee_create_account")
			s.NotificationInviteeCreateAccount.Encode(e)
		}
	}
	{
		if s.NotificationLatestNews.Set {
			e.FieldStart("notification_latest_news")
			s.NotificationLatestNews.Encode(e)
		}
	}
	{
		if s.NotificationLike.Set {
			e.FieldStart("notification_like")
			s.NotificationLike.Encode(e)
		}
	}
	{
		if s.NotificationMessageTag.Set {
			e.FieldStart("notification_message_tag")
			s.NotificationMessageTag.Encode(e)
		}
	}
	{
		if s.NotificationPopularPost.Set {
			e.FieldStart("notification_popular_post")
			s.NotificationPopularPost.Encode(e)
		}
	}
	{
		if s.NotificationProfileScreenshot.Set {
			e.FieldStart("notification_profile_screenshot")
			s.NotificationProfileScreenshot.Encode(e)
		}
	}
	{
		if s.NotificationReply.Set {
			e.FieldStart("notification_reply")
			s.NotificationReply.Encode(e)
		}
	}
	{
		if s.NotificationRepost.Set {
			e.FieldStart("notification_repost")
			s.NotificationRepost.Encode(e)
		}
	}
	{
		if s.NotificationReview.Set {
			e.FieldStart("notification_review")
			s.NotificationReview.Encode(e)
		}
	}
	{
		if s.NotificationSecurityWarning.Set {
			e.FieldStart("notification_security_warning")
			s.NotificationSecurityWarning.Encode(e)
		}
	}
	{
		if s.NotificationTwitterFriend.Set {
			e.FieldStart("notification_twitter_friend")
			s.NotificationTwitterFriend.Encode(e)
		}
	}
	{
		if s.PrivacyMode.Set {
			e.FieldStart("privacy_mode")
			s.PrivacyMode.Encode(e)
		}
	}
	{
		if s.PrivatePost.Set {
			e.FieldStart("private_post")
			s.PrivatePost.Encode(e)
		}
	}
	{
		if s.PrivateUserTimeline.Set {
			e.FieldStart("private_user_timeline")
			s.PrivateUserTimeline.Encode(e)
		}
	}
	{
		if s.ShowTotalGiftsReceivedCount.Set {
			e.FieldStart("show_total_gifts_received_count")
			s.ShowTotalGiftsReceivedCount.Encode(e)
		}
	}
	{
		if s.VipInvisibleFootprintMode.Set {
			e.FieldStart("vip_invisible_footprint_mode")
			s.VipInvisibleFootprintMode.Encode(e)
		}
	}
	{
		if s.VisibleOnSnsFriendRecommendation.Set {
			e.FieldStart("visible_on_sns_friend_recommendation")
			s.VisibleOnSnsFriendRecommendation.Encode(e)
		}
	}
}

var jsonFieldsNameOfSettings = [60]string{
	0:  "age_restricted_on_review",
	1:  "allow_reposts",
	2:  "caution_user_chat",
	3:  "following_only_call_invite",
	4:  "following_only_group_invite",
	5:  "following_restricted_on_review",
	6:  "hide_active_call",
	7:  "hide_gifts_received",
	8:  "hide_hot_post",
	9:  "hide_on_invitable",
	10: "hide_online_status",
	11: "hide_reply_following_timeline",
	12: "hide_reply_public_timeline",
	13: "hide_vip",
	14: "invisible_on_user_search",
	15: "no_reply_group_timeline",
	16: "notification_birthday_to_followers",
	17: "notification_bulk_call_invite",
	18: "notification_call_invite",
	19: "notification_chat",
	20: "notification_chat_delete",
	21: "notification_contact_friend",
	22: "notification_daily_quest",
	23: "notification_daily_summary",
	24: "notification_follow",
	25: "notification_follow_accept",
	26: "notification_follow_request",
	27: "notification_follower_conference_call",
	28: "notification_follower_create_group",
	29: "notification_following_birthdate_on",
	30: "notification_following_post_after_break",
	31: "notification_followings_in_call",
	32: "notification_footprint",
	33: "notification_gift_receive",
	34: "notification_group_accept",
	35: "notification_group_conference_call",
	36: "notification_group_invite",
	37: "notification_group_join",
	38: "notification_group_message_tag_all",
	39: "notification_group_moderator",
	40: "notification_group_post",
	41: "notification_group_request",
	42: "notification_hima_now",
	43: "notification_invitee_create_account",
	44: "notification_latest_news",
	45: "notification_like",
	46: "notification_message_tag",
	47: "notification_popular_post",
	48: "notification_profile_screenshot",
	49: "notification_reply",
	50: "notification_repost",
	51: "notification_review",
	52: "notification_security_warning",
	53: "notification_twitter_friend",
	54: "privacy_mode",
	55: "private_post",
	56: "private_user_timeline",
	57: "show_total_gifts_received_count",
	58: "vip_invisible_footprint_mode",
	59: "visible_on_sns_friend_recommendation",
}

// Decode decodes Settings from json.
func (s *Settings) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Settings to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "age_restricted_on_review":
			if err := func() error {
				s.AgeRestrictedOnReview.Reset()
				if err := s.AgeRestrictedOnReview.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"age_restricted_on_review\"")
			}
		case "allow_reposts":
			if err := func() error {
				s.AllowReposts.Reset()
				if err := s.AllowReposts.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"allow_reposts\"")
			}
		case "caution_user_chat":
			if err := func() error {
				s.CautionUserChat.Reset()
				if err := s.CautionUserChat.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"caution_user_chat\"")
			}
		case "following_only_call_invite":
			if err := func() error {
				s.FollowingOnlyCallInvite.Reset()
				if err := s.FollowingOnlyCallInvite.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"following_only_call_invite\"")
			}
		case "following_only_group_invite":
			if err := func() error {
				s.FollowingOnlyGroupInvite.Reset()
				if err := s.FollowingOnlyGroupInvite.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"following_only_group_invite\"")
			}
		case "following_restricted_on_review":
			if err := func() error {
				s.FollowingRestrictedOnReview.Reset()
				if err := s.FollowingRestrictedOnReview.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"following_restricted_on_review\"")
			}
		case "hide_active_call":
			if err := func() error {
				s.HideActiveCall.Reset()
				if err := s.HideActiveCall.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"hide_active_call\"")
			}
		case "hide_gifts_received":
			if err := func() error {
				s.HideGiftsReceived.Reset()
				if err := s.HideGiftsReceived.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"hide_gifts_received\"")
			}
		case "hide_hot_post":
			if err := func() error {
				s.HideHotPost.Reset()
				if err := s.HideHotPost.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"hide_hot_post\"")
			}
		case "hide_on_invitable":
			if err := func() error {
				s.HideOnInvitable.Reset()
				if err := s.HideOnInvitable.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"hide_on_invitable\"")
			}
		case "hide_online_status":
			if err := func() error {
				s.HideOnlineStatus.Reset()
				if err := s.HideOnlineStatus.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"hide_online_status\"")
			}
		case "hide_reply_following_timeline":
			if err := func() error {
				s.HideReplyFollowingTimeline.Reset()
				if err := s.HideReplyFollowingTimeline.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"hide_reply_following_timeline\"")
			}
		case "hide_reply_public_timeline":
			if err := func() error {
				s.HideReplyPublicTimeline.Reset()
				if err := s.HideReplyPublicTimeline.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"hide_reply_public_timeline\"")
			}
		case "hide_vip":
			if err := func() error {
				s.HideVip.Reset()
				if err := s.HideVip.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"hide_vip\"")
			}
		case "invisible_on_user_search":
			if err := func() error {
				s.InvisibleOnUserSearch.Reset()
				if err := s.InvisibleOnUserSearch.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"invisible_on_user_search\"")
			}
		case "no_reply_group_timeline":
			if err := func() error {
				s.NoReplyGroupTimeline.Reset()
				if err := s.NoReplyGroupTimeline.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"no_reply_group_timeline\"")
			}
		case "notification_birthday_to_followers":
			if err := func() error {
				s.NotificationBirthdayToFollowers.Reset()
				if err := s.NotificationBirthdayToFollowers.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"notification_birthday_to_followers\"")
			}
		case "notification_bulk_call_invite":
			if err := func() error {
				s.NotificationBulkCallInvite.Reset()
				if err := s.NotificationBulkCallInvite.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"notification_bulk_call_invite\"")
			}
		case "notification_call_invite":
			if err := func() error {
				s.NotificationCallInvite.Reset()
				if err := s.NotificationCallInvite.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"notification_call_invite\"")
			}
		case "notification_chat":
			if err := func() error {
				s.NotificationChat.Reset()
				if err := s.NotificationChat.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"notification_chat\"")
			}
		case "notification_chat_delete":
			if err := func() error {
				s.NotificationChatDelete.Reset()
				if err := s.NotificationChatDelete.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"notification_chat_delete\"")
			}
		case "notification_contact_friend":
			if err := func() error {
				s.NotificationContactFriend.Reset()
				if err := s.NotificationContactFriend.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"notification_contact_friend\"")
			}
		case "notification_daily_quest":
			if err := func() error {
				s.NotificationDailyQuest.Reset()
				if err := s.NotificationDailyQuest.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"notification_daily_quest\"")
			}
		case "notification_daily_summary":
			if err := func() error {
				s.NotificationDailySummary.Reset()
				if err := s.NotificationDailySummary.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"notification_daily_summary\"")
			}
		case "notification_follow":
			if err := func() error {
				s.NotificationFollow.Reset()
				if err := s.NotificationFollow.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"notification_follow\"")
			}
		case "notification_follow_accept":
			if err := func() error {
				s.NotificationFollowAccept.Reset()
				if err := s.NotificationFollowAccept.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"notification_follow_accept\"")
			}
		case "notification_follow_request":
			if err := func() error {
				s.NotificationFollowRequest.Reset()
				if err := s.NotificationFollowRequest.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"notification_follow_request\"")
			}
		case "notification_follower_conference_call":
			if err := func() error {
				s.NotificationFollowerConferenceCall.Reset()
				if err := s.NotificationFollowerConferenceCall.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"notification_follower_conference_call\"")
			}
		case "notification_follower_create_group":
			if err := func() error {
				s.NotificationFollowerCreateGroup.Reset()
				if err := s.NotificationFollowerCreateGroup.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"notification_follower_create_group\"")
			}
		case "notification_following_birthdate_on":
			if err := func() error {
				s.NotificationFollowingBirthdateOn.Reset()
				if err := s.NotificationFollowingBirthdateOn.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"notification_following_birthdate_on\"")
			}
		case "notification_following_post_after_break":
			if err := func() error {
				s.NotificationFollowingPostAfterBreak.Reset()
				if err := s.NotificationFollowingPostAfterBreak.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"notification_following_post_after_break\"")
			}
		case "notification_followings_in_call":
			if err := func() error {
				s.NotificationFollowingsInCall.Reset()
				if err := s.NotificationFollowingsInCall.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"notification_followings_in_call\"")
			}
		case "notification_footprint":
			if err := func() error {
				s.NotificationFootprint.Reset()
				if err := s.NotificationFootprint.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"notification_footprint\"")
			}
		case "notification_gift_receive":
			if err := func() error {
				s.NotificationGiftReceive.Reset()
				if err := s.NotificationGiftReceive.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"notification_gift_receive\"")
			}
		case "notification_group_accept":
			if err := func() error {
				s.NotificationGroupAccept.Reset()
				if err := s.NotificationGroupAccept.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"notification_group_accept\"")
			}
		case "notification_group_conference_call":
			if err := func() error {
				s.NotificationGroupConferenceCall.Reset()
				if err := s.NotificationGroupConferenceCall.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"notification_group_conference_call\"")
			}
		case "notification_group_invite":
			if err := func() error {
				s.NotificationGroupInvite.Reset()
				if err := s.NotificationGroupInvite.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"notification_group_invite\"")
			}
		case "notification_group_join":
			if err := func() error {
				s.NotificationGroupJoin.Reset()
				if err := s.NotificationGroupJoin.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"notification_group_join\"")
			}
		case "notification_group_message_tag_all":
			if err := func() error {
				s.NotificationGroupMessageTagAll.Reset()
				if err := s.NotificationGroupMessageTagAll.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"notification_group_message_tag_all\"")
			}
		case "notification_group_moderator":
			if err := func() error {
				s.NotificationGroupModerator.Reset()
				if err := s.NotificationGroupModerator.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"notification_group_moderator\"")
			}
		case "notification_group_post":
			if err := func() error {
				s.NotificationGroupPost.Reset()
				if err := s.NotificationGroupPost.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"notification_group_post\"")
			}
		case "notification_group_request":
			if err := func() error {
				s.NotificationGroupRequest.Reset()
				if err := s.NotificationGroupRequest.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"notification_group_request\"")
			}
		case "notification_hima_now":
			if err := func() error {
				s.NotificationHimaNow.Reset()
				if err := s.NotificationHimaNow.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"notification_hima_now\"")
			}
		case "notification_invitee_create_account":
			if err := func() error {
				s.NotificationInviteeCreateAccount.Reset()
				if err := s.NotificationInviteeCreateAccount.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"notification_invitee_create_account\"")
			}
		case "notification_latest_news":
			if err := func() error {
				s.NotificationLatestNews.Reset()
				if err := s.NotificationLatestNews.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"notification_latest_news\"")
			}
		case "notification_like":
			if err := func() error {
				s.NotificationLike.Reset()
				if err := s.NotificationLike.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"notification_like\"")
			}
		case "notification_message_tag":
			if err := func() error {
				s.NotificationMessageTag.Reset()
				if err := s.NotificationMessageTag.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"notification_message_tag\"")
			}
		case "notification_popular_post":
			if err := func() error {
				s.NotificationPopularPost.Reset()
				if err := s.NotificationPopularPost.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"notification_popular_post\"")
			}
		case "notification_profile_screenshot":
			if err := func() error {
				s.NotificationProfileScreenshot.Reset()
				if err := s.NotificationProfileScreenshot.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"notification_profile_screenshot\"")
			}
		case "notification_reply":
			if err := func() error {
				s.NotificationReply.Reset()
				if err := s.NotificationReply.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"notification_reply\"")
			}
		case "notification_repost":
			if err := func() error {
				s.NotificationRepost.Reset()
				if err := s.NotificationRepost.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"notification_repost\"")
			}
		case "notification_review":
			if err := func() error {
				s.NotificationReview.Reset()
				if err := s.NotificationReview.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"notification_review\"")
			}
		case "notification_security_warning":
			if err := func() error {
				s.NotificationSecurityWarning.Reset()
				if err := s.NotificationSecurityWarning.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"notification_security_warning\"")
			}
		case "notification_twitter_friend":
			if err := func() error {
				s.NotificationTwitterFriend.Reset()
				if err := s.NotificationTwitterFriend.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"notification_twitter_friend\"")
			}
		case "privacy_mode":
			if err := func() error {
				s.PrivacyMode.Reset()
				if err := s.PrivacyMode.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"privacy_mode\"")
			}
		case "private_post":
			if err := func() error {
				s.PrivatePost.Reset()
				if err := s.PrivatePost.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"private_post\"")
			}
		case "private_user_timeline":
			if err := func() error {
				s.PrivateUserTimeline.Reset()
				if err := s.PrivateUserTimeline.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"private_user_timeline\"")
			}
		case "show_total_gifts_received_count":
			if err := func() error {
				s.ShowTotalGiftsReceivedCount.Reset()
				if err := s.ShowTotalGiftsReceivedCount.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"show_total_gifts_received_count\"")
			}
		case "vip_invisible_footprint_mode":
			if err := func() error {
				s.VipInvisibleFootprintMode.Reset()
				if err := s.VipInvisibleFootprintMode.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"vip_invisible_footprint_mode\"")
			}
		case "visible_on_sns_friend_recommendation":
			if err := func() error {
				s.VisibleOnSnsFriendRecommendation.Reset()
				if err := s.VisibleOnSnsFriendRecommendation.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"visible_on_sns_friend_recommendation\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode Settings")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *Settings) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *Settings) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *Shareable) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *Shareable) encodeFields(e *jx.Encoder) {
	{
		if s.Group.Set {
			e.FieldStart("group")
			s.Group.Encode(e)
		}
	}
	{
		if s.Post.Set {
			e.FieldStart("post")
			s.Post.Encode(e)
		}
	}
	{
		if s.Thread.Set {
			e.FieldStart("thread")
			s.Thread.Encode(e)
		}
	}
}

var jsonFieldsNameOfShareable = [3]string{
	0: "group",
	1: "post",
	2: "thread",
}

// Decode decodes Shareable from json.
func (s *Shareable) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Shareable to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "group":
			if err := func() error {
				s.Group.Reset()
				if err := s.Group.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"group\"")
			}
		case "post":
			if err := func() error {
				s.Post.Reset()
				if err := s.Post.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"post\"")
			}
		case "thread":
			if err := func() error {
				s.Thread.Reset()
				if err := s.Thread.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"thread\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode Shareable")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *Shareable) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *Shareable) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *SharedUrl) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *SharedUrl) encodeFields(e *jx.Encoder) {
	{
		if s.Description.Set {
			e.FieldStart("description")
			s.Description.Encode(e)
		}
	}
	{
		if s.ImageURL.Set {
			e.FieldStart("image_url")
			s.ImageURL.Encode(e)
		}
	}
	{
		if s.Title.Set {
			e.FieldStart("title")
			s.Title.Encode(e)
		}
	}
	{
		if s.URL.Set {
			e.FieldStart("url")
			s.URL.Encode(e)
		}
	}
}

var jsonFieldsNameOfSharedUrl = [4]string{
	0: "description",
	1: "image_url",
	2: "title",
	3: "url",
}

// Decode decodes SharedUrl from json.
func (s *SharedUrl) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode SharedUrl to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "description":
			if err := func() error {
				s.Description.Reset()
				if err := s.Description.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"description\"")
			}
		case "image_url":
			if err := func() error {
				s.ImageURL.Reset()
				if err := s.ImageURL.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"image_url\"")
			}
		case "title":
			if err := func() error {
				s.Title.Reset()
				if err := s.Title.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"title\"")
			}
		case "url":
			if err := func() error {
				s.URL.Reset()
				if err := s.URL.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"url\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode SharedUrl")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *SharedUrl) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *SharedUrl) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *SignaturePayload) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *SignaturePayload) encodeFields(e *jx.Encoder) {
	{
		if s.Action.Set {
			e.FieldStart("action")
			s.Action.Encode(e)
		}
	}
	{
		if s.CallID.Set {
			e.FieldStart("call_id")
			s.CallID.Encode(e)
		}
	}
	{
		if s.ReceiverUUID.Set {
			e.FieldStart("receiver_uuid")
			s.ReceiverUUID.Encode(e)
		}
	}
	{
		if s.SenderUUID.Set {
			e.FieldStart("sender_uuid")
			s.SenderUUID.Encode(e)
		}
	}
	{
		if s.Signature.Set {
			e.FieldStart("signature")
			s.Signature.Encode(e)
		}
	}
	{
		if s.Timestamp.Set {
			e.FieldStart("timestamp")
			s.Timestamp.Encode(e)
		}
	}
}

var jsonFieldsNameOfSignaturePayload = [6]string{
	0: "action",
	1: "call_id",
	2: "receiver_uuid",
	3: "sender_uuid",
	4: "signature",
	5: "timestamp",
}

// Decode decodes SignaturePayload from json.
func (s *SignaturePayload) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode SignaturePayload to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "action":
			if err := func() error {
				s.Action.Reset()
				if err := s.Action.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"action\"")
			}
		case "call_id":
			if err := func() error {
				s.CallID.Reset()
				if err := s.CallID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"call_id\"")
			}
		case "receiver_uuid":
			if err := func() error {
				s.ReceiverUUID.Reset()
				if err := s.ReceiverUUID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"receiver_uuid\"")
			}
		case "sender_uuid":
			if err := func() error {
				s.SenderUUID.Reset()
				if err := s.SenderUUID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"sender_uuid\"")
			}
		case "signature":
			if err := func() error {
				s.Signature.Reset()
				if err := s.Signature.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"signature\"")
			}
		case "timestamp":
			if err := func() error {
				s.Timestamp.Reset()
				if err := s.Timestamp.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"timestamp\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode SignaturePayload")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *SignaturePayload) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *SignaturePayload) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *SnsInfo) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *SnsInfo) encodeFields(e *jx.Encoder) {
	{
		if s.Biography.Set {
			e.FieldStart("biography")
			s.Biography.Encode(e)
		}
	}
	{
		if s.Gender.Set {
			e.FieldStart("gender")
			s.Gender.Encode(e)
		}
	}
	{
		if s.Nickname.Set {
			e.FieldStart("nickname")
			s.Nickname.Encode(e)
		}
	}
	{
		if s.ProfileImage.Set {
			e.FieldStart("profile_image")
			s.ProfileImage.Encode(e)
		}
	}
	{
		if s.Type.Set {
			e.FieldStart("type")
			s.Type.Encode(e)
		}
	}
	{
		if s.UID.Set {
			e.FieldStart("uid")
			s.UID.Encode(e)
		}
	}
}

var jsonFieldsNameOfSnsInfo = [6]string{
	0: "biography",
	1: "gender",
	2: "nickname",
	3: "profile_image",
	4: "type",
	5: "uid",
}

// Decode decodes SnsInfo from json.
func (s *SnsInfo) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode SnsInfo to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "biography":
			if err := func() error {
				s.Biography.Reset()
				if err := s.Biography.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"biography\"")
			}
		case "gender":
			if err := func() error {
				s.Gender.Reset()
				if err := s.Gender.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"gender\"")
			}
		case "nickname":
			if err := func() error {
				s.Nickname.Reset()
				if err := s.Nickname.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"nickname\"")
			}
		case "profile_image":
			if err := func() error {
				s.ProfileImage.Reset()
				if err := s.ProfileImage.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"profile_image\"")
			}
		case "type":
			if err := func() error {
				s.Type.Reset()
				if err := s.Type.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"type\"")
			}
		case "uid":
			if err := func() error {
				s.UID.Reset()
				if err := s.UID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"uid\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode SnsInfo")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *SnsInfo) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *SnsInfo) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *Sticker) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *Sticker) encodeFields(e *jx.Encoder) {
	{
		if s.Extension.Set {
			e.FieldStart("extension")
			s.Extension.Encode(e)
		}
	}
	{
		if s.Height.Set {
			e.FieldStart("height")
			s.Height.Encode(e)
		}
	}
	{
		if s.ID.Set {
			e.FieldStart("id")
			s.ID.Encode(e)
		}
	}
	{
		if s.StickerPackID.Set {
			e.FieldStart("sticker_pack_id")
			s.StickerPackID.Encode(e)
		}
	}
	{
		if s.URL.Set {
			e.FieldStart("url")
			s.URL.Encode(e)
		}
	}
	{
		if s.Width.Set {
			e.FieldStart("width")
			s.Width.Encode(e)
		}
	}
}

var jsonFieldsNameOfSticker = [6]string{
	0: "extension",
	1: "height",
	2: "id",
	3: "sticker_pack_id",
	4: "url",
	5: "width",
}

// Decode decodes Sticker from json.
func (s *Sticker) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Sticker to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "extension":
			if err := func() error {
				s.Extension.Reset()
				if err := s.Extension.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"extension\"")
			}
		case "height":
			if err := func() error {
				s.Height.Reset()
				if err := s.Height.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"height\"")
			}
		case "id":
			if err := func() error {
				s.ID.Reset()
				if err := s.ID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"id\"")
			}
		case "sticker_pack_id":
			if err := func() error {
				s.StickerPackID.Reset()
				if err := s.StickerPackID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"sticker_pack_id\"")
			}
		case "url":
			if err := func() error {
				s.URL.Reset()
				if err := s.URL.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"url\"")
			}
		case "width":
			if err := func() error {
				s.Width.Reset()
				if err := s.Width.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"width\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode Sticker")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *Sticker) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *Sticker) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *StickerPack) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *StickerPack) encodeFields(e *jx.Encoder) {
	{
		if s.Cover.Set {
			e.FieldStart("cover")
			s.Cover.Encode(e)
		}
	}
	{
		if s.Description.Set {
			e.FieldStart("description")
			s.Description.Encode(e)
		}
	}
	{
		if s.ID.Set {
			e.FieldStart("id")
			s.ID.Encode(e)
		}
	}
	{
		if s.Name.Set {
			e.FieldStart("name")
			s.Name.Encode(e)
		}
	}
	{
		if s.Order.Set {
			e.FieldStart("order")
			s.Order.Encode(e)
		}
	}
	{
		if s.Stickers.Set {
			e.FieldStart("stickers")
			s.Stickers.Encode(e)
		}
	}
}

var jsonFieldsNameOfStickerPack = [6]string{
	0: "cover",
	1: "description",
	2: "id",
	3: "name",
	4: "order",
	5: "stickers",
}

// Decode decodes StickerPack from json.
func (s *StickerPack) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode StickerPack to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "cover":
			if err := func() error {
				s.Cover.Reset()
				if err := s.Cover.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"cover\"")
			}
		case "description":
			if err := func() error {
				s.Description.Reset()
				if err := s.Description.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"description\"")
			}
		case "id":
			if err := func() error {
				s.ID.Reset()
				if err := s.ID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"id\"")
			}
		case "name":
			if err := func() error {
				s.Name.Reset()
				if err := s.Name.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"name\"")
			}
		case "order":
			if err := func() error {
				s.Order.Reset()
				if err := s.Order.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"order\"")
			}
		case "stickers":
			if err := func() error {
				s.Stickers.Reset()
				if err := s.Stickers.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"stickers\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode StickerPack")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *StickerPack) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *StickerPack) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *StickerPacksResponse) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *StickerPacksResponse) encodeFields(e *jx.Encoder) {
	{
		if s.StickerPacks.Set {
			e.FieldStart("sticker_packs")
			s.StickerPacks.Encode(e)
		}
	}
}

var jsonFieldsNameOfStickerPacksResponse = [1]string{
	0: "sticker_packs",
}

// Decode decodes StickerPacksResponse from json.
func (s *StickerPacksResponse) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode StickerPacksResponse to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "sticker_packs":
			if err := func() error {
				s.StickerPacks.Reset()
				if err := s.StickerPacks.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"sticker_packs\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode StickerPacksResponse")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *StickerPacksResponse) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *StickerPacksResponse) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *Survey) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *Survey) encodeFields(e *jx.Encoder) {
	{
		if s.Choices.Set {
			e.FieldStart("choices")
			s.Choices.Encode(e)
		}
	}
	{
		if s.ID.Set {
			e.FieldStart("id")
			s.ID.Encode(e)
		}
	}
	{
		if s.Voted.Set {
			e.FieldStart("voted")
			s.Voted.Encode(e)
		}
	}
	{
		if s.VotesCount.Set {
			e.FieldStart("votes_count")
			s.VotesCount.Encode(e)
		}
	}
}

var jsonFieldsNameOfSurvey = [4]string{
	0: "choices",
	1: "id",
	2: "voted",
	3: "votes_count",
}

// Decode decodes Survey from json.
func (s *Survey) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Survey to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "choices":
			if err := func() error {
				s.Choices.Reset()
				if err := s.Choices.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"choices\"")
			}
		case "id":
			if err := func() error {
				s.ID.Reset()
				if err := s.ID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"id\"")
			}
		case "voted":
			if err := func() error {
				s.Voted.Reset()
				if err := s.Voted.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"voted\"")
			}
		case "votes_count":
			if err := func() error {
				s.VotesCount.Reset()
				if err := s.VotesCount.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"votes_count\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode Survey")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *Survey) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *Survey) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *ThreadInfo) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *ThreadInfo) encodeFields(e *jx.Encoder) {
	{
		if s.CreatedAt.Set {
			e.FieldStart("created_at")
			s.CreatedAt.Encode(e)
		}
	}
	{
		if s.ID.Set {
			e.FieldStart("id")
			s.ID.Encode(e)
		}
	}
	{
		if s.IsJoined.Set {
			e.FieldStart("is_joined")
			s.IsJoined.Encode(e)
		}
	}
	{
		if s.LastPost.Set {
			e.FieldStart("last_post")
			s.LastPost.Encode(e)
		}
	}
	{
		if s.NewUpdates.Set {
			e.FieldStart("new_updates")
			s.NewUpdates.Encode(e)
		}
	}
	{
		if s.Owner.Set {
			e.FieldStart("owner")
			s.Owner.Encode(e)
		}
	}
	{
		if s.PostsCount.Set {
			e.FieldStart("posts_count")
			s.PostsCount.Encode(e)
		}
	}
	{
		if s.ThreadIcon.Set {
			e.FieldStart("thread_icon")
			s.ThreadIcon.Encode(e)
		}
	}
	{
		if s.Title.Set {
			e.FieldStart("title")
			s.Title.Encode(e)
		}
	}
	{
		if s.UnreadCount.Set {
			e.FieldStart("unread_count")
			s.UnreadCount.Encode(e)
		}
	}
	{
		if s.UpdatedAt.Set {
			e.FieldStart("updated_at")
			s.UpdatedAt.Encode(e)
		}
	}
}

var jsonFieldsNameOfThreadInfo = [11]string{
	0:  "created_at",
	1:  "id",
	2:  "is_joined",
	3:  "last_post",
	4:  "new_updates",
	5:  "owner",
	6:  "posts_count",
	7:  "thread_icon",
	8:  "title",
	9:  "unread_count",
	10: "updated_at",
}

// Decode decodes ThreadInfo from json.
func (s *ThreadInfo) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode ThreadInfo to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "created_at":
			if err := func() error {
				s.CreatedAt.Reset()
				if err := s.CreatedAt.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"created_at\"")
			}
		case "id":
			if err := func() error {
				s.ID.Reset()
				if err := s.ID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"id\"")
			}
		case "is_joined":
			if err := func() error {
				s.IsJoined.Reset()
				if err := s.IsJoined.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"is_joined\"")
			}
		case "last_post":
			if err := func() error {
				s.LastPost.Reset()
				if err := s.LastPost.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"last_post\"")
			}
		case "new_updates":
			if err := func() error {
				s.NewUpdates.Reset()
				if err := s.NewUpdates.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"new_updates\"")
			}
		case "owner":
			if err := func() error {
				s.Owner.Reset()
				if err := s.Owner.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"owner\"")
			}
		case "posts_count":
			if err := func() error {
				s.PostsCount.Reset()
				if err := s.PostsCount.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"posts_count\"")
			}
		case "thread_icon":
			if err := func() error {
				s.ThreadIcon.Reset()
				if err := s.ThreadIcon.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"thread_icon\"")
			}
		case "title":
			if err := func() error {
				s.Title.Reset()
				if err := s.Title.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"title\"")
			}
		case "unread_count":
			if err := func() error {
				s.UnreadCount.Reset()
				if err := s.UnreadCount.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"unread_count\"")
			}
		case "updated_at":
			if err := func() error {
				s.UpdatedAt.Reset()
				if err := s.UpdatedAt.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"updated_at\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode ThreadInfo")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *ThreadInfo) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *ThreadInfo) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *TimelineSettings) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *TimelineSettings) encodeFields(e *jx.Encoder) {
	{
		if s.FavesFilter.Set {
			e.FieldStart("faves_filter")
			s.FavesFilter.Encode(e)
		}
	}
	{
		if s.HideHotPost.Set {
			e.FieldStart("hide_hot_post")
			s.HideHotPost.Encode(e)
		}
	}
	{
		if s.HideReplyFollowingTimeline.Set {
			e.FieldStart("hide_reply_following_timeline")
			s.HideReplyFollowingTimeline.Encode(e)
		}
	}
	{
		if s.HideReplyPublicTimeline.Set {
			e.FieldStart("hide_reply_public_timeline")
			s.HideReplyPublicTimeline.Encode(e)
		}
	}
}

var jsonFieldsNameOfTimelineSettings = [4]string{
	0: "faves_filter",
	1: "hide_hot_post",
	2: "hide_reply_following_timeline",
	3: "hide_reply_public_timeline",
}

// Decode decodes TimelineSettings from json.
func (s *TimelineSettings) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode TimelineSettings to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "faves_filter":
			if err := func() error {
				s.FavesFilter.Reset()
				if err := s.FavesFilter.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"faves_filter\"")
			}
		case "hide_hot_post":
			if err := func() error {
				s.HideHotPost.Reset()
				if err := s.HideHotPost.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"hide_hot_post\"")
			}
		case "hide_reply_following_timeline":
			if err := func() error {
				s.HideReplyFollowingTimeline.Reset()
				if err := s.HideReplyFollowingTimeline.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"hide_reply_following_timeline\"")
			}
		case "hide_reply_public_timeline":
			if err := func() error {
				s.HideReplyPublicTimeline.Reset()
				if err := s.HideReplyPublicTimeline.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"hide_reply_public_timeline\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode TimelineSettings")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *TimelineSettings) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *TimelineSettings) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *Title) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *Title) encodeFields(e *jx.Encoder) {
	{
		if s.APIValue.Set {
			e.FieldStart("api_value")
			s.APIValue.Encode(e)
		}
	}
}

var jsonFieldsNameOfTitle = [1]string{
	0: "api_value",
}

// Decode decodes Title from json.
func (s *Title) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Title to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "api_value":
			if err := func() error {
				s.APIValue.Reset()
				if err := s.APIValue.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"api_value\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode Title")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *Title) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *Title) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *TokenResponse) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *TokenResponse) encodeFields(e *jx.Encoder) {
	{
		if s.AccessToken.Set {
			e.FieldStart("access_token")
			s.AccessToken.Encode(e)
		}
	}
	{
		if s.CreatedAt.Set {
			e.FieldStart("created_at")
			s.CreatedAt.Encode(e)
		}
	}
	{
		if s.ExpiresIn.Set {
			e.FieldStart("expires_in")
			s.ExpiresIn.Encode(e)
		}
	}
	{
		if s.ID.Set {
			e.FieldStart("id")
			s.ID.Encode(e)
		}
	}
	{
		if s.RefreshToken.Set {
			e.FieldStart("refresh_token")
			s.RefreshToken.Encode(e)
		}
	}
}

var jsonFieldsNameOfTokenResponse = [5]string{
	0: "access_token",
	1: "created_at",
	2: "expires_in",
	3: "id",
	4: "refresh_token",
}

// Decode decodes TokenResponse from json.
func (s *TokenResponse) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode TokenResponse to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "access_token":
			if err := func() error {
				s.AccessToken.Reset()
				if err := s.AccessToken.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"access_token\"")
			}
		case "created_at":
			if err := func() error {
				s.CreatedAt.Reset()
				if err := s.CreatedAt.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"created_at\"")
			}
		case "expires_in":
			if err := func() error {
				s.ExpiresIn.Reset()
				if err := s.ExpiresIn.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"expires_in\"")
			}
		case "id":
			if err := func() error {
				s.ID.Reset()
				if err := s.ID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"id\"")
			}
		case "refresh_token":
			if err := func() error {
				s.RefreshToken.Reset()
				if err := s.RefreshToken.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"refresh_token\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode TokenResponse")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *TokenResponse) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *TokenResponse) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *TotalChatRequestResponse) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *TotalChatRequestResponse) encodeFields(e *jx.Encoder) {
	{
		if s.Total.Set {
			e.FieldStart("total")
			s.Total.Encode(e)
		}
	}
}

var jsonFieldsNameOfTotalChatRequestResponse = [1]string{
	0: "total",
}

// Decode decodes TotalChatRequestResponse from json.
func (s *TotalChatRequestResponse) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode TotalChatRequestResponse to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "total":
			if err := func() error {
				s.Total.Reset()
				if err := s.Total.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"total\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode TotalChatRequestResponse")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *TotalChatRequestResponse) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *TotalChatRequestResponse) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *TransactionGiftReceived) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *TransactionGiftReceived) encodeFields(e *jx.Encoder) {
	{
		if s.Icon.Set {
			e.FieldStart("icon")
			s.Icon.Encode(e)
		}
	}
	{
		if s.Item.Set {
			e.FieldStart("item")
			s.Item.Encode(e)
		}
	}
	{
		if s.ItemID.Set {
			e.FieldStart("item_id")
			s.ItemID.Encode(e)
		}
	}
	{
		if s.Quantity.Set {
			e.FieldStart("quantity")
			s.Quantity.Encode(e)
		}
	}
}

var jsonFieldsNameOfTransactionGiftReceived = [4]string{
	0: "icon",
	1: "item",
	2: "item_id",
	3: "quantity",
}

// Decode decodes TransactionGiftReceived from json.
func (s *TransactionGiftReceived) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode TransactionGiftReceived to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "icon":
			if err := func() error {
				s.Icon.Reset()
				if err := s.Icon.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"icon\"")
			}
		case "item":
			if err := func() error {
				s.Item.Reset()
				if err := s.Item.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"item\"")
			}
		case "item_id":
			if err := func() error {
				s.ItemID.Reset()
				if err := s.ItemID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"item_id\"")
			}
		case "quantity":
			if err := func() error {
				s.Quantity.Reset()
				if err := s.Quantity.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"quantity\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode TransactionGiftReceived")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *TransactionGiftReceived) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *TransactionGiftReceived) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *TransactionGiftReceivedItem) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *TransactionGiftReceivedItem) encodeFields(e *jx.Encoder) {
	{
		if s.Currency.Set {
			e.FieldStart("currency")
			s.Currency.Encode(e)
		}
	}
	{
		if s.ID.Set {
			e.FieldStart("id")
			s.ID.Encode(e)
		}
	}
	{
		if s.Price.Set {
			e.FieldStart("price")
			s.Price.Encode(e)
		}
	}
	{
		if s.Title.Set {
			e.FieldStart("title")
			s.Title.Encode(e)
		}
	}
}

var jsonFieldsNameOfTransactionGiftReceivedItem = [4]string{
	0: "currency",
	1: "id",
	2: "price",
	3: "title",
}

// Decode decodes TransactionGiftReceivedItem from json.
func (s *TransactionGiftReceivedItem) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode TransactionGiftReceivedItem to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "currency":
			if err := func() error {
				s.Currency.Reset()
				if err := s.Currency.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"currency\"")
			}
		case "id":
			if err := func() error {
				s.ID.Reset()
				if err := s.ID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"id\"")
			}
		case "price":
			if err := func() error {
				s.Price.Reset()
				if err := s.Price.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"price\"")
			}
		case "title":
			if err := func() error {
				s.Title.Reset()
				if err := s.Title.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"title\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode TransactionGiftReceivedItem")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *TransactionGiftReceivedItem) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *TransactionGiftReceivedItem) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *TwoFAStatusResponse) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *TwoFAStatusResponse) encodeFields(e *jx.Encoder) {
	{
		if s.TwoFaAuthRequired.Set {
			e.FieldStart("two_fa_auth_required")
			s.TwoFaAuthRequired.Encode(e)
		}
	}
	{
		if s.TwoFaEnabled.Set {
			e.FieldStart("two_fa_enabled")
			s.TwoFaEnabled.Encode(e)
		}
	}
}

var jsonFieldsNameOfTwoFAStatusResponse = [2]string{
	0: "two_fa_auth_required",
	1: "two_fa_enabled",
}

// Decode decodes TwoFAStatusResponse from json.
func (s *TwoFAStatusResponse) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode TwoFAStatusResponse to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "two_fa_auth_required":
			if err := func() error {
				s.TwoFaAuthRequired.Reset()
				if err := s.TwoFaAuthRequired.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"two_fa_auth_required\"")
			}
		case "two_fa_enabled":
			if err := func() error {
				s.TwoFaEnabled.Reset()
				if err := s.TwoFaEnabled.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"two_fa_enabled\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode TwoFAStatusResponse")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *TwoFAStatusResponse) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *TwoFAStatusResponse) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *TwoFaAuthRequiredDTO) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *TwoFaAuthRequiredDTO) encodeFields(e *jx.Encoder) {
	{
		if s.Login.Set {
			e.FieldStart("login")
			s.Login.Encode(e)
		}
	}
	{
		if s.Wallet.Set {
			e.FieldStart("wallet")
			s.Wallet.Encode(e)
		}
	}
}

var jsonFieldsNameOfTwoFaAuthRequiredDTO = [2]string{
	0: "login",
	1: "wallet",
}

// Decode decodes TwoFaAuthRequiredDTO from json.
func (s *TwoFaAuthRequiredDTO) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode TwoFaAuthRequiredDTO to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "login":
			if err := func() error {
				s.Login.Reset()
				if err := s.Login.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"login\"")
			}
		case "wallet":
			if err := func() error {
				s.Wallet.Reset()
				if err := s.Wallet.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"wallet\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode TwoFaAuthRequiredDTO")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *TwoFaAuthRequiredDTO) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *TwoFaAuthRequiredDTO) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *TwoStepAuthEnabledResponse) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *TwoStepAuthEnabledResponse) encodeFields(e *jx.Encoder) {
	{
		if s.RecoveryCodes.Set {
			e.FieldStart("recovery_codes")
			s.RecoveryCodes.Encode(e)
		}
	}
}

var jsonFieldsNameOfTwoStepAuthEnabledResponse = [1]string{
	0: "recovery_codes",
}

// Decode decodes TwoStepAuthEnabledResponse from json.
func (s *TwoStepAuthEnabledResponse) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode TwoStepAuthEnabledResponse to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "recovery_codes":
			if err := func() error {
				s.RecoveryCodes.Reset()
				if err := s.RecoveryCodes.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"recovery_codes\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode TwoStepAuthEnabledResponse")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *TwoStepAuthEnabledResponse) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *TwoStepAuthEnabledResponse) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *TwoStepAuthRequestInfoResponse) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *TwoStepAuthRequestInfoResponse) encodeFields(e *jx.Encoder) {
	{
		if s.QrCode.Set {
			e.FieldStart("qr_code")
			s.QrCode.Encode(e)
		}
	}
	{
		if s.Secret.Set {
			e.FieldStart("secret")
			s.Secret.Encode(e)
		}
	}
}

var jsonFieldsNameOfTwoStepAuthRequestInfoResponse = [2]string{
	0: "qr_code",
	1: "secret",
}

// Decode decodes TwoStepAuthRequestInfoResponse from json.
func (s *TwoStepAuthRequestInfoResponse) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode TwoStepAuthRequestInfoResponse to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "qr_code":
			if err := func() error {
				s.QrCode.Reset()
				if err := s.QrCode.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"qr_code\"")
			}
		case "secret":
			if err := func() error {
				s.Secret.Reset()
				if err := s.Secret.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"secret\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode TwoStepAuthRequestInfoResponse")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *TwoStepAuthRequestInfoResponse) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *TwoStepAuthRequestInfoResponse) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *UnreadStatusResponse) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *UnreadStatusResponse) encodeFields(e *jx.Encoder) {
	{
		if s.IsUnread.Set {
			e.FieldStart("is_unread")
			s.IsUnread.Encode(e)
		}
	}
}

var jsonFieldsNameOfUnreadStatusResponse = [1]string{
	0: "is_unread",
}

// Decode decodes UnreadStatusResponse from json.
func (s *UnreadStatusResponse) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode UnreadStatusResponse to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "is_unread":
			if err := func() error {
				s.IsUnread.Reset()
				if err := s.IsUnread.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"is_unread\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode UnreadStatusResponse")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *UnreadStatusResponse) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *UnreadStatusResponse) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *UpdatePostReqMessageTags) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *UpdatePostReqMessageTags) encodeFields(e *jx.Encoder) {
}

var jsonFieldsNameOfUpdatePostReqMessageTags = [0]string{}

// Decode decodes UpdatePostReqMessageTags from json.
func (s *UpdatePostReqMessageTags) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode UpdatePostReqMessageTags to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		default:
			return d.Skip()
		}
	}); err != nil {
		return errors.Wrap(err, "decode UpdatePostReqMessageTags")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *UpdatePostReqMessageTags) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *UpdatePostReqMessageTags) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *User) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *User) encodeFields(e *jx.Encoder) {
	{
		if s.Biography.Set {
			e.FieldStart("biography")
			s.Biography.Encode(e)
		}
	}
	{
		if s.BirthDate.Set {
			e.FieldStart("birth_date")
			s.BirthDate.Encode(e)
		}
	}
	{
		if s.BlockingLimit.Set {
			e.FieldStart("blocking_limit")
			s.BlockingLimit.Encode(e)
		}
	}
	{
		if s.ConnectedBy.Set {
			e.FieldStart("connected_by")
			s.ConnectedBy.Encode(e)
		}
	}
	{
		if s.ContactPhones.Set {
			e.FieldStart("contact_phones")
			s.ContactPhones.Encode(e)
		}
	}
	{
		if s.Country.Set {
			e.FieldStart("country")
			s.Country.Encode(e)
		}
	}
	{
		if s.CoverImage.Set {
			e.FieldStart("cover_image")
			s.CoverImage.Encode(e)
		}
	}
	{
		if s.CoverImageThumbnail.Set {
			e.FieldStart("cover_image_thumbnail")
			s.CoverImageThumbnail.Encode(e)
		}
	}
	{
		if s.CreatedAtMillis.Set {
			e.FieldStart("created_at_millis")
			s.CreatedAtMillis.Encode(e)
		}
	}
	{
		if s.FollowersCount.Set {
			e.FieldStart("followers_count")
			s.FollowersCount.Encode(e)
		}
	}
	{
		if s.FollowingsCount.Set {
			e.FieldStart("followings_count")
			s.FollowingsCount.Encode(e)
		}
	}
	{
		if s.Gender.Set {
			e.FieldStart("gender")
			s.Gender.Encode(e)
		}
	}
	{
		if s.GiftReceivedCounter.Set {
			e.FieldStart("gift_received_counter")
			s.GiftReceivedCounter.Encode(e)
		}
	}
	{
		if s.GiftingAbility.Set {
			e.FieldStart("gifting_ability")
			s.GiftingAbility.Encode(e)
		}
	}
	{
		if s.GroupsUsersCount.Set {
			e.FieldStart("groups_users_count")
			s.GroupsUsersCount.Encode(e)
		}
	}
	{
		if s.ID.Set {
			e.FieldStart("id")
			s.ID.Encode(e)
		}
	}
	{
		if s.InterestsCount.Set {
			e.FieldStart("interests_count")
			s.InterestsCount.Encode(e)
		}
	}
	{
		if s.IsAgeVerified.Set {
			e.FieldStart("is_age_verified")
			s.IsAgeVerified.Encode(e)
		}
	}
	{
		if s.IsAppleConnected.Set {
			e.FieldStart("is_apple_connected")
			s.IsAppleConnected.Encode(e)
		}
	}
	{
		if s.IsChatRequestOn.Set {
			e.FieldStart("is_chat_request_on")
			s.IsChatRequestOn.Encode(e)
		}
	}
	{
		if s.IsDangerous.Set {
			e.FieldStart("is_dangerous")
			s.IsDangerous.Encode(e)
		}
	}
	{
		if s.IsEmailConfirmed.Set {
			e.FieldStart("is_email_confirmed")
			s.IsEmailConfirmed.Encode(e)
		}
	}
	{
		if s.IsFacebookConnected.Set {
			e.FieldStart("is_facebook_connected")
			s.IsFacebookConnected.Encode(e)
		}
	}
	{
		if s.IsFollowPending.Set {
			e.FieldStart("is_follow_pending")
			s.IsFollowPending.Encode(e)
		}
	}
	{
		if s.IsFollowedBy.Set {
			e.FieldStart("is_followed_by")
			s.IsFollowedBy.Encode(e)
		}
	}
	{
		if s.IsFollowedByPending.Set {
			e.FieldStart("is_followed_by_pending")
			s.IsFollowedByPending.Encode(e)
		}
	}
	{
		if s.IsFollowing.Set {
			e.FieldStart("is_following")
			s.IsFollowing.Encode(e)
		}
	}
	{
		if s.IsFromDifferentGenerationAndTrusted.Set {
			e.FieldStart("is_from_different_generation_and_trusted")
			s.IsFromDifferentGenerationAndTrusted.Encode(e)
		}
	}
	{
		if s.IsGoogleConnected.Set {
			e.FieldStart("is_google_connected")
			s.IsGoogleConnected.Encode(e)
		}
	}
	{
		if s.IsGroupPhoneOn.Set {
			e.FieldStart("is_group_phone_on")
			s.IsGroupPhoneOn.Encode(e)
		}
	}
	{
		if s.IsGroupVideoOn.Set {
			e.FieldStart("is_group_video_on")
			s.IsGroupVideoOn.Encode(e)
		}
	}
	{
		if s.IsHidden.Set {
			e.FieldStart("is_hidden")
			s.IsHidden.Encode(e)
		}
	}
	{
		if s.IsHideVip.Set {
			e.FieldStart("is_hide_vip")
			s.IsHideVip.Encode(e)
		}
	}
	{
		if s.IsInterestsSelected.Set {
			e.FieldStart("is_interests_selected")
			s.IsInterestsSelected.Encode(e)
		}
	}
	{
		if s.IsLineConnected.Set {
			e.FieldStart("is_line_connected")
			s.IsLineConnected.Encode(e)
		}
	}
	{
		if s.IsMuted.Set {
			e.FieldStart("is_muted")
			s.IsMuted.Encode(e)
		}
	}
	{
		if s.IsNew.Set {
			e.FieldStart("is_new")
			s.IsNew.Encode(e)
		}
	}
	{
		if s.IsPhoneOn.Set {
			e.FieldStart("is_phone_on")
			s.IsPhoneOn.Encode(e)
		}
	}
	{
		if s.IsPrivate.Set {
			e.FieldStart("is_private")
			s.IsPrivate.Encode(e)
		}
	}
	{
		if s.IsRecentlyPenalized.Set {
			e.FieldStart("is_recently_penalized")
			s.IsRecentlyPenalized.Encode(e)
		}
	}
	{
		if s.IsRegistered.Set {
			e.FieldStart("is_registered")
			s.IsRegistered.Encode(e)
		}
	}
	{
		if s.IsTwoFaEnabled.Set {
			e.FieldStart("is_two_fa_enabled")
			s.IsTwoFaEnabled.Encode(e)
		}
	}
	{
		if s.IsVideoOn.Set {
			e.FieldStart("is_video_on")
			s.IsVideoOn.Encode(e)
		}
	}
	{
		if s.IsVip.Set {
			e.FieldStart("is_vip")
			s.IsVip.Encode(e)
		}
	}
	{
		if s.IsWorldIDConnected.Set {
			e.FieldStart("is_world_id_connected")
			s.IsWorldIDConnected.Encode(e)
		}
	}
	{
		if s.LastLoggedInAtMillis.Set {
			e.FieldStart("last_logged_in_at_millis")
			s.LastLoggedInAtMillis.Encode(e)
		}
	}
	{
		if s.LoginStreakCount.Set {
			e.FieldStart("login_streak_count")
			s.LoginStreakCount.Encode(e)
		}
	}
	{
		if s.MaskedEmail.Set {
			e.FieldStart("masked_email")
			s.MaskedEmail.Encode(e)
		}
	}
	{
		if s.MatchingID.Set {
			e.FieldStart("matching_id")
			s.MatchingID.Encode(e)
		}
	}
	{
		if s.MediaCount.Set {
			e.FieldStart("media_count")
			s.MediaCount.Encode(e)
		}
	}
	{
		if s.MobileNumber.Set {
			e.FieldStart("mobile_number")
			s.MobileNumber.Encode(e)
		}
	}
	{
		if s.Nickname.Set {
			e.FieldStart("nickname")
			s.Nickname.Encode(e)
		}
	}
	{
		if s.OnlineStatus.Set {
			e.FieldStart("online_status")
			s.OnlineStatus.Encode(e)
		}
	}
	{
		if s.PostsCount.Set {
			e.FieldStart("posts_count")
			s.PostsCount.Encode(e)
		}
	}
	{
		if s.Prefecture.Set {
			e.FieldStart("prefecture")
			s.Prefecture.Encode(e)
		}
	}
	{
		if s.ProfileIcon.Set {
			e.FieldStart("profile_icon")
			s.ProfileIcon.Encode(e)
		}
	}
	{
		if s.ProfileIconFrame.Set {
			e.FieldStart("profile_icon_frame")
			s.ProfileIconFrame.Encode(e)
		}
	}
	{
		if s.ProfileIconFrameThumbnail.Set {
			e.FieldStart("profile_icon_frame_thumbnail")
			s.ProfileIconFrameThumbnail.Encode(e)
		}
	}
	{
		if s.ProfileIconThumbnail.Set {
			e.FieldStart("profile_icon_thumbnail")
			s.ProfileIconThumbnail.Encode(e)
		}
	}
	{
		if s.ReviewRestriction.Set {
			e.FieldStart("review_restriction")
			s.ReviewRestriction.Encode(e)
		}
	}
	{
		if s.ReviewsCount.Set {
			e.FieldStart("reviews_count")
			s.ReviewsCount.Encode(e)
		}
	}
	{
		if s.Title.Set {
			e.FieldStart("title")
			s.Title.Encode(e)
		}
	}
	{
		if s.TwitterID.Set {
			e.FieldStart("twitter_id")
			s.TwitterID.Encode(e)
		}
	}
	{
		if s.Username.Set {
			e.FieldStart("username")
			s.Username.Encode(e)
		}
	}
	{
		if s.UUID.Set {
			e.FieldStart("uuid")
			s.UUID.Encode(e)
		}
	}
	{
		if s.VipUntilMillis.Set {
			e.FieldStart("vip_until_millis")
			s.VipUntilMillis.Encode(e)
		}
	}
}

var jsonFieldsNameOfUser = [66]string{
	0:  "biography",
	1:  "birth_date",
	2:  "blocking_limit",
	3:  "connected_by",
	4:  "contact_phones",
	5:  "country",
	6:  "cover_image",
	7:  "cover_image_thumbnail",
	8:  "created_at_millis",
	9:  "followers_count",
	10: "followings_count",
	11: "gender",
	12: "gift_received_counter",
	13: "gifting_ability",
	14: "groups_users_count",
	15: "id",
	16: "interests_count",
	17: "is_age_verified",
	18: "is_apple_connected",
	19: "is_chat_request_on",
	20: "is_dangerous",
	21: "is_email_confirmed",
	22: "is_facebook_connected",
	23: "is_follow_pending",
	24: "is_followed_by",
	25: "is_followed_by_pending",
	26: "is_following",
	27: "is_from_different_generation_and_trusted",
	28: "is_google_connected",
	29: "is_group_phone_on",
	30: "is_group_video_on",
	31: "is_hidden",
	32: "is_hide_vip",
	33: "is_interests_selected",
	34: "is_line_connected",
	35: "is_muted",
	36: "is_new",
	37: "is_phone_on",
	38: "is_private",
	39: "is_recently_penalized",
	40: "is_registered",
	41: "is_two_fa_enabled",
	42: "is_video_on",
	43: "is_vip",
	44: "is_world_id_connected",
	45: "last_logged_in_at_millis",
	46: "login_streak_count",
	47: "masked_email",
	48: "matching_id",
	49: "media_count",
	50: "mobile_number",
	51: "nickname",
	52: "online_status",
	53: "posts_count",
	54: "prefecture",
	55: "profile_icon",
	56: "profile_icon_frame",
	57: "profile_icon_frame_thumbnail",
	58: "profile_icon_thumbnail",
	59: "review_restriction",
	60: "reviews_count",
	61: "title",
	62: "twitter_id",
	63: "username",
	64: "uuid",
	65: "vip_until_millis",
}

// Decode decodes User from json.
func (s *User) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode User to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "biography":
			if err := func() error {
				s.Biography.Reset()
				if err := s.Biography.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"biography\"")
			}
		case "birth_date":
			if err := func() error {
				s.BirthDate.Reset()
				if err := s.BirthDate.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"birth_date\"")
			}
		case "blocking_limit":
			if err := func() error {
				s.BlockingLimit.Reset()
				if err := s.BlockingLimit.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"blocking_limit\"")
			}
		case "connected_by":
			if err := func() error {
				s.ConnectedBy.Reset()
				if err := s.ConnectedBy.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"connected_by\"")
			}
		case "contact_phones":
			if err := func() error {
				s.ContactPhones.Reset()
				if err := s.ContactPhones.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"contact_phones\"")
			}
		case "country":
			if err := func() error {
				s.Country.Reset()
				if err := s.Country.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"country\"")
			}
		case "cover_image":
			if err := func() error {
				s.CoverImage.Reset()
				if err := s.CoverImage.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"cover_image\"")
			}
		case "cover_image_thumbnail":
			if err := func() error {
				s.CoverImageThumbnail.Reset()
				if err := s.CoverImageThumbnail.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"cover_image_thumbnail\"")
			}
		case "created_at_millis":
			if err := func() error {
				s.CreatedAtMillis.Reset()
				if err := s.CreatedAtMillis.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"created_at_millis\"")
			}
		case "followers_count":
			if err := func() error {
				s.FollowersCount.Reset()
				if err := s.FollowersCount.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"followers_count\"")
			}
		case "followings_count":
			if err := func() error {
				s.FollowingsCount.Reset()
				if err := s.FollowingsCount.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"followings_count\"")
			}
		case "gender":
			if err := func() error {
				s.Gender.Reset()
				if err := s.Gender.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"gender\"")
			}
		case "gift_received_counter":
			if err := func() error {
				s.GiftReceivedCounter.Reset()
				if err := s.GiftReceivedCounter.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"gift_received_counter\"")
			}
		case "gifting_ability":
			if err := func() error {
				s.GiftingAbility.Reset()
				if err := s.GiftingAbility.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"gifting_ability\"")
			}
		case "groups_users_count":
			if err := func() error {
				s.GroupsUsersCount.Reset()
				if err := s.GroupsUsersCount.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"groups_users_count\"")
			}
		case "id":
			if err := func() error {
				s.ID.Reset()
				if err := s.ID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"id\"")
			}
		case "interests_count":
			if err := func() error {
				s.InterestsCount.Reset()
				if err := s.InterestsCount.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"interests_count\"")
			}
		case "is_age_verified":
			if err := func() error {
				s.IsAgeVerified.Reset()
				if err := s.IsAgeVerified.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"is_age_verified\"")
			}
		case "is_apple_connected":
			if err := func() error {
				s.IsAppleConnected.Reset()
				if err := s.IsAppleConnected.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"is_apple_connected\"")
			}
		case "is_chat_request_on":
			if err := func() error {
				s.IsChatRequestOn.Reset()
				if err := s.IsChatRequestOn.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"is_chat_request_on\"")
			}
		case "is_dangerous":
			if err := func() error {
				s.IsDangerous.Reset()
				if err := s.IsDangerous.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"is_dangerous\"")
			}
		case "is_email_confirmed":
			if err := func() error {
				s.IsEmailConfirmed.Reset()
				if err := s.IsEmailConfirmed.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"is_email_confirmed\"")
			}
		case "is_facebook_connected":
			if err := func() error {
				s.IsFacebookConnected.Reset()
				if err := s.IsFacebookConnected.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"is_facebook_connected\"")
			}
		case "is_follow_pending":
			if err := func() error {
				s.IsFollowPending.Reset()
				if err := s.IsFollowPending.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"is_follow_pending\"")
			}
		case "is_followed_by":
			if err := func() error {
				s.IsFollowedBy.Reset()
				if err := s.IsFollowedBy.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"is_followed_by\"")
			}
		case "is_followed_by_pending":
			if err := func() error {
				s.IsFollowedByPending.Reset()
				if err := s.IsFollowedByPending.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"is_followed_by_pending\"")
			}
		case "is_following":
			if err := func() error {
				s.IsFollowing.Reset()
				if err := s.IsFollowing.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"is_following\"")
			}
		case "is_from_different_generation_and_trusted":
			if err := func() error {
				s.IsFromDifferentGenerationAndTrusted.Reset()
				if err := s.IsFromDifferentGenerationAndTrusted.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"is_from_different_generation_and_trusted\"")
			}
		case "is_google_connected":
			if err := func() error {
				s.IsGoogleConnected.Reset()
				if err := s.IsGoogleConnected.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"is_google_connected\"")
			}
		case "is_group_phone_on":
			if err := func() error {
				s.IsGroupPhoneOn.Reset()
				if err := s.IsGroupPhoneOn.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"is_group_phone_on\"")
			}
		case "is_group_video_on":
			if err := func() error {
				s.IsGroupVideoOn.Reset()
				if err := s.IsGroupVideoOn.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"is_group_video_on\"")
			}
		case "is_hidden":
			if err := func() error {
				s.IsHidden.Reset()
				if err := s.IsHidden.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"is_hidden\"")
			}
		case "is_hide_vip":
			if err := func() error {
				s.IsHideVip.Reset()
				if err := s.IsHideVip.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"is_hide_vip\"")
			}
		case "is_interests_selected":
			if err := func() error {
				s.IsInterestsSelected.Reset()
				if err := s.IsInterestsSelected.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"is_interests_selected\"")
			}
		case "is_line_connected":
			if err := func() error {
				s.IsLineConnected.Reset()
				if err := s.IsLineConnected.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"is_line_connected\"")
			}
		case "is_muted":
			if err := func() error {
				s.IsMuted.Reset()
				if err := s.IsMuted.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"is_muted\"")
			}
		case "is_new":
			if err := func() error {
				s.IsNew.Reset()
				if err := s.IsNew.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"is_new\"")
			}
		case "is_phone_on":
			if err := func() error {
				s.IsPhoneOn.Reset()
				if err := s.IsPhoneOn.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"is_phone_on\"")
			}
		case "is_private":
			if err := func() error {
				s.IsPrivate.Reset()
				if err := s.IsPrivate.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"is_private\"")
			}
		case "is_recently_penalized":
			if err := func() error {
				s.IsRecentlyPenalized.Reset()
				if err := s.IsRecentlyPenalized.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"is_recently_penalized\"")
			}
		case "is_registered":
			if err := func() error {
				s.IsRegistered.Reset()
				if err := s.IsRegistered.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"is_registered\"")
			}
		case "is_two_fa_enabled":
			if err := func() error {
				s.IsTwoFaEnabled.Reset()
				if err := s.IsTwoFaEnabled.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"is_two_fa_enabled\"")
			}
		case "is_video_on":
			if err := func() error {
				s.IsVideoOn.Reset()
				if err := s.IsVideoOn.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"is_video_on\"")
			}
		case "is_vip":
			if err := func() error {
				s.IsVip.Reset()
				if err := s.IsVip.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"is_vip\"")
			}
		case "is_world_id_connected":
			if err := func() error {
				s.IsWorldIDConnected.Reset()
				if err := s.IsWorldIDConnected.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"is_world_id_connected\"")
			}
		case "last_logged_in_at_millis":
			if err := func() error {
				s.LastLoggedInAtMillis.Reset()
				if err := s.LastLoggedInAtMillis.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"last_logged_in_at_millis\"")
			}
		case "login_streak_count":
			if err := func() error {
				s.LoginStreakCount.Reset()
				if err := s.LoginStreakCount.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"login_streak_count\"")
			}
		case "masked_email":
			if err := func() error {
				s.MaskedEmail.Reset()
				if err := s.MaskedEmail.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"masked_email\"")
			}
		case "matching_id":
			if err := func() error {
				s.MatchingID.Reset()
				if err := s.MatchingID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"matching_id\"")
			}
		case "media_count":
			if err := func() error {
				s.MediaCount.Reset()
				if err := s.MediaCount.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"media_count\"")
			}
		case "mobile_number":
			if err := func() error {
				s.MobileNumber.Reset()
				if err := s.MobileNumber.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"mobile_number\"")
			}
		case "nickname":
			if err := func() error {
				s.Nickname.Reset()
				if err := s.Nickname.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"nickname\"")
			}
		case "online_status":
			if err := func() error {
				s.OnlineStatus.Reset()
				if err := s.OnlineStatus.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"online_status\"")
			}
		case "posts_count":
			if err := func() error {
				s.PostsCount.Reset()
				if err := s.PostsCount.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"posts_count\"")
			}
		case "prefecture":
			if err := func() error {
				s.Prefecture.Reset()
				if err := s.Prefecture.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"prefecture\"")
			}
		case "profile_icon":
			if err := func() error {
				s.ProfileIcon.Reset()
				if err := s.ProfileIcon.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"profile_icon\"")
			}
		case "profile_icon_frame":
			if err := func() error {
				s.ProfileIconFrame.Reset()
				if err := s.ProfileIconFrame.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"profile_icon_frame\"")
			}
		case "profile_icon_frame_thumbnail":
			if err := func() error {
				s.ProfileIconFrameThumbnail.Reset()
				if err := s.ProfileIconFrameThumbnail.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"profile_icon_frame_thumbnail\"")
			}
		case "profile_icon_thumbnail":
			if err := func() error {
				s.ProfileIconThumbnail.Reset()
				if err := s.ProfileIconThumbnail.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"profile_icon_thumbnail\"")
			}
		case "review_restriction":
			if err := func() error {
				s.ReviewRestriction.Reset()
				if err := s.ReviewRestriction.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"review_restriction\"")
			}
		case "reviews_count":
			if err := func() error {
				s.ReviewsCount.Reset()
				if err := s.ReviewsCount.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"reviews_count\"")
			}
		case "title":
			if err := func() error {
				s.Title.Reset()
				if err := s.Title.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"title\"")
			}
		case "twitter_id":
			if err := func() error {
				s.TwitterID.Reset()
				if err := s.TwitterID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"twitter_id\"")
			}
		case "username":
			if err := func() error {
				s.Username.Reset()
				if err := s.Username.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"username\"")
			}
		case "uuid":
			if err := func() error {
				s.UUID.Reset()
				if err := s.UUID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"uuid\"")
			}
		case "vip_until_millis":
			if err := func() error {
				s.VipUntilMillis.Reset()
				if err := s.VipUntilMillis.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"vip_until_millis\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode User")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *User) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *User) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *UserConnectedBy) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *UserConnectedBy) encodeFields(e *jx.Encoder) {
}

var jsonFieldsNameOfUserConnectedBy = [0]string{}

// Decode decodes UserConnectedBy from json.
func (s *UserConnectedBy) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode UserConnectedBy to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		default:
			return d.Skip()
		}
	}); err != nil {
		return errors.Wrap(err, "decode UserConnectedBy")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *UserConnectedBy) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *UserConnectedBy) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *UserCountry) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *UserCountry) encodeFields(e *jx.Encoder) {
}

var jsonFieldsNameOfUserCountry = [0]string{}

// Decode decodes UserCountry from json.
func (s *UserCountry) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode UserCountry to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		default:
			return d.Skip()
		}
	}); err != nil {
		return errors.Wrap(err, "decode UserCountry")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *UserCountry) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *UserCountry) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *UserCustomDefinitionsResponse) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *UserCustomDefinitionsResponse) encodeFields(e *jx.Encoder) {
	{
		if s.Age.Set {
			e.FieldStart("age")
			s.Age.Encode(e)
		}
	}
	{
		if s.CreatedAt.Set {
			e.FieldStart("created_at")
			s.CreatedAt.Encode(e)
		}
	}
	{
		if s.FollowersCount.Set {
			e.FieldStart("followers_count")
			s.FollowersCount.Encode(e)
		}
	}
	{
		if s.FollowingsCount.Set {
			e.FieldStart("followings_count")
			s.FollowingsCount.Encode(e)
		}
	}
	{
		if s.LastLoggedinAt.Set {
			e.FieldStart("last_loggedin_at")
			s.LastLoggedinAt.Encode(e)
		}
	}
	{
		if s.ReportedCount.Set {
			e.FieldStart("reported_count")
			s.ReportedCount.Encode(e)
		}
	}
	{
		if s.Status.Set {
			e.FieldStart("status")
			s.Status.Encode(e)
		}
	}
}

var jsonFieldsNameOfUserCustomDefinitionsResponse = [7]string{
	0: "age",
	1: "created_at",
	2: "followers_count",
	3: "followings_count",
	4: "last_loggedin_at",
	5: "reported_count",
	6: "status",
}

// Decode decodes UserCustomDefinitionsResponse from json.
func (s *UserCustomDefinitionsResponse) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode UserCustomDefinitionsResponse to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "age":
			if err := func() error {
				s.Age.Reset()
				if err := s.Age.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"age\"")
			}
		case "created_at":
			if err := func() error {
				s.CreatedAt.Reset()
				if err := s.CreatedAt.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"created_at\"")
			}
		case "followers_count":
			if err := func() error {
				s.FollowersCount.Reset()
				if err := s.FollowersCount.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"followers_count\"")
			}
		case "followings_count":
			if err := func() error {
				s.FollowingsCount.Reset()
				if err := s.FollowingsCount.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"followings_count\"")
			}
		case "last_loggedin_at":
			if err := func() error {
				s.LastLoggedinAt.Reset()
				if err := s.LastLoggedinAt.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"last_loggedin_at\"")
			}
		case "reported_count":
			if err := func() error {
				s.ReportedCount.Reset()
				if err := s.ReportedCount.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"reported_count\"")
			}
		case "status":
			if err := func() error {
				s.Status.Reset()
				if err := s.Status.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"status\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode UserCustomDefinitionsResponse")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *UserCustomDefinitionsResponse) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *UserCustomDefinitionsResponse) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *UserInterestsResponse) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *UserInterestsResponse) encodeFields(e *jx.Encoder) {
	{
		if s.Interests.Set {
			e.FieldStart("interests")
			s.Interests.Encode(e)
		}
	}
}

var jsonFieldsNameOfUserInterestsResponse = [1]string{
	0: "interests",
}

// Decode decodes UserInterestsResponse from json.
func (s *UserInterestsResponse) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode UserInterestsResponse to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "interests":
			if err := func() error {
				s.Interests.Reset()
				if err := s.Interests.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"interests\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode UserInterestsResponse")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *UserInterestsResponse) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *UserInterestsResponse) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *UserResponse) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *UserResponse) encodeFields(e *jx.Encoder) {
	{
		if s.AppleConnected.Set {
			e.FieldStart("apple_connected")
			s.AppleConnected.Encode(e)
		}
	}
	{
		if s.BirthDate.Set {
			e.FieldStart("birth_date")
			s.BirthDate.Encode(e)
		}
	}
	{
		if s.BlockingLimit.Set {
			e.FieldStart("blocking_limit")
			s.BlockingLimit.Encode(e)
		}
	}
	{
		if s.EmailConfirmed.Set {
			e.FieldStart("email_confirmed")
			s.EmailConfirmed.Encode(e)
		}
	}
	{
		if s.FacebookConnected.Set {
			e.FieldStart("facebook_connected")
			s.FacebookConnected.Encode(e)
		}
	}
	{
		if s.GiftingAbility.Set {
			e.FieldStart("gifting_ability")
			s.GiftingAbility.Encode(e)
		}
	}
	{
		if s.GoogleConnected.Set {
			e.FieldStart("google_connected")
			s.GoogleConnected.Encode(e)
		}
	}
	{
		if s.GroupPhoneOn.Set {
			e.FieldStart("group_phone_on")
			s.GroupPhoneOn.Encode(e)
		}
	}
	{
		if s.GroupVideoOn.Set {
			e.FieldStart("group_video_on")
			s.GroupVideoOn.Encode(e)
		}
	}
	{
		if s.InterestsCount.Set {
			e.FieldStart("interests_count")
			s.InterestsCount.Encode(e)
		}
	}
	{
		if s.LineConnected.Set {
			e.FieldStart("line_connected")
			s.LineConnected.Encode(e)
		}
	}
	{
		if s.MaskedEmail.Set {
			e.FieldStart("masked_email")
			s.MaskedEmail.Encode(e)
		}
	}
	{
		if s.PhoneOn.Set {
			e.FieldStart("phone_on")
			s.PhoneOn.Encode(e)
		}
	}
	{
		if s.PushNotification.Set {
			e.FieldStart("push_notification")
			s.PushNotification.Encode(e)
		}
	}
	{
		if s.TwitterID.Set {
			e.FieldStart("twitter_id")
			s.TwitterID.Encode(e)
		}
	}
	{
		if s.User.Set {
			e.FieldStart("user")
			s.User.Encode(e)
		}
	}
	{
		if s.UUID.Set {
			e.FieldStart("uuid")
			s.UUID.Encode(e)
		}
	}
	{
		if s.VideoOn.Set {
			e.FieldStart("video_on")
			s.VideoOn.Encode(e)
		}
	}
	{
		if s.VipUntilSeconds.Set {
			e.FieldStart("vip_until_seconds")
			s.VipUntilSeconds.Encode(e)
		}
	}
	{
		if s.WorldIDConnected.Set {
			e.FieldStart("world_id_connected")
			s.WorldIDConnected.Encode(e)
		}
	}
}

var jsonFieldsNameOfUserResponse = [20]string{
	0:  "apple_connected",
	1:  "birth_date",
	2:  "blocking_limit",
	3:  "email_confirmed",
	4:  "facebook_connected",
	5:  "gifting_ability",
	6:  "google_connected",
	7:  "group_phone_on",
	8:  "group_video_on",
	9:  "interests_count",
	10: "line_connected",
	11: "masked_email",
	12: "phone_on",
	13: "push_notification",
	14: "twitter_id",
	15: "user",
	16: "uuid",
	17: "video_on",
	18: "vip_until_seconds",
	19: "world_id_connected",
}

// Decode decodes UserResponse from json.
func (s *UserResponse) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode UserResponse to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "apple_connected":
			if err := func() error {
				s.AppleConnected.Reset()
				if err := s.AppleConnected.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"apple_connected\"")
			}
		case "birth_date":
			if err := func() error {
				s.BirthDate.Reset()
				if err := s.BirthDate.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"birth_date\"")
			}
		case "blocking_limit":
			if err := func() error {
				s.BlockingLimit.Reset()
				if err := s.BlockingLimit.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"blocking_limit\"")
			}
		case "email_confirmed":
			if err := func() error {
				s.EmailConfirmed.Reset()
				if err := s.EmailConfirmed.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"email_confirmed\"")
			}
		case "facebook_connected":
			if err := func() error {
				s.FacebookConnected.Reset()
				if err := s.FacebookConnected.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"facebook_connected\"")
			}
		case "gifting_ability":
			if err := func() error {
				s.GiftingAbility.Reset()
				if err := s.GiftingAbility.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"gifting_ability\"")
			}
		case "google_connected":
			if err := func() error {
				s.GoogleConnected.Reset()
				if err := s.GoogleConnected.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"google_connected\"")
			}
		case "group_phone_on":
			if err := func() error {
				s.GroupPhoneOn.Reset()
				if err := s.GroupPhoneOn.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"group_phone_on\"")
			}
		case "group_video_on":
			if err := func() error {
				s.GroupVideoOn.Reset()
				if err := s.GroupVideoOn.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"group_video_on\"")
			}
		case "interests_count":
			if err := func() error {
				s.InterestsCount.Reset()
				if err := s.InterestsCount.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"interests_count\"")
			}
		case "line_connected":
			if err := func() error {
				s.LineConnected.Reset()
				if err := s.LineConnected.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"line_connected\"")
			}
		case "masked_email":
			if err := func() error {
				s.MaskedEmail.Reset()
				if err := s.MaskedEmail.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"masked_email\"")
			}
		case "phone_on":
			if err := func() error {
				s.PhoneOn.Reset()
				if err := s.PhoneOn.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"phone_on\"")
			}
		case "push_notification":
			if err := func() error {
				s.PushNotification.Reset()
				if err := s.PushNotification.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"push_notification\"")
			}
		case "twitter_id":
			if err := func() error {
				s.TwitterID.Reset()
				if err := s.TwitterID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"twitter_id\"")
			}
		case "user":
			if err := func() error {
				s.User.Reset()
				if err := s.User.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"user\"")
			}
		case "uuid":
			if err := func() error {
				s.UUID.Reset()
				if err := s.UUID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"uuid\"")
			}
		case "video_on":
			if err := func() error {
				s.VideoOn.Reset()
				if err := s.VideoOn.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"video_on\"")
			}
		case "vip_until_seconds":
			if err := func() error {
				s.VipUntilSeconds.Reset()
				if err := s.VipUntilSeconds.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"vip_until_seconds\"")
			}
		case "world_id_connected":
			if err := func() error {
				s.WorldIDConnected.Reset()
				if err := s.WorldIDConnected.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"world_id_connected\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode UserResponse")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *UserResponse) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *UserResponse) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *UserSetting) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *UserSetting) encodeFields(e *jx.Encoder) {
	{
		if s.NotificationChat.Set {
			e.FieldStart("notification_chat")
			s.NotificationChat.Encode(e)
		}
	}
}

var jsonFieldsNameOfUserSetting = [1]string{
	0: "notification_chat",
}

// Decode decodes UserSetting from json.
func (s *UserSetting) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode UserSetting to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "notification_chat":
			if err := func() error {
				s.NotificationChat.Reset()
				if err := s.NotificationChat.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"notification_chat\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode UserSetting")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *UserSetting) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *UserSetting) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *UserTimestampResponse) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *UserTimestampResponse) encodeFields(e *jx.Encoder) {
	{
		if s.Country.Set {
			e.FieldStart("country")
			s.Country.Encode(e)
		}
	}
	{
		if s.IPAddress.Set {
			e.FieldStart("ip_address")
			s.IPAddress.Encode(e)
		}
	}
	{
		if s.Time.Set {
			e.FieldStart("time")
			s.Time.Encode(e)
		}
	}
}

var jsonFieldsNameOfUserTimestampResponse = [3]string{
	0: "country",
	1: "ip_address",
	2: "time",
}

// Decode decodes UserTimestampResponse from json.
func (s *UserTimestampResponse) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode UserTimestampResponse to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "country":
			if err := func() error {
				s.Country.Reset()
				if err := s.Country.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"country\"")
			}
		case "ip_address":
			if err := func() error {
				s.IPAddress.Reset()
				if err := s.IPAddress.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"ip_address\"")
			}
		case "time":
			if err := func() error {
				s.Time.Reset()
				if err := s.Time.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"time\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode UserTimestampResponse")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *UserTimestampResponse) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *UserTimestampResponse) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *UserUserDTO) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *UserUserDTO) encodeFields(e *jx.Encoder) {
	{
		if s.AgeVerified.Set {
			e.FieldStart("age_verified")
			s.AgeVerified.Encode(e)
		}
	}
	{
		if s.Biography.Set {
			e.FieldStart("biography")
			s.Biography.Encode(e)
		}
	}
	{
		if s.CoverImage.Set {
			e.FieldStart("cover_image")
			s.CoverImage.Encode(e)
		}
	}
	{
		if s.CoverImageThumbnail.Set {
			e.FieldStart("cover_image_thumbnail")
			s.CoverImageThumbnail.Encode(e)
		}
	}
	{
		if s.CreatedAt.Set {
			e.FieldStart("created_at")
			s.CreatedAt.Encode(e)
		}
	}
	{
		if s.DangerousUser.Set {
			e.FieldStart("dangerous_user")
			s.DangerousUser.Encode(e)
		}
	}
	{
		if s.FollowPending.Set {
			e.FieldStart("follow_pending")
			s.FollowPending.Encode(e)
		}
	}
	{
		if s.FollowedBy.Set {
			e.FieldStart("followed_by")
			s.FollowedBy.Encode(e)
		}
	}
	{
		if s.FollowersCount.Set {
			e.FieldStart("followers_count")
			s.FollowersCount.Encode(e)
		}
	}
	{
		if s.Following.Set {
			e.FieldStart("following")
			s.Following.Encode(e)
		}
	}
	{
		if s.FollowingsCount.Set {
			e.FieldStart("followings_count")
			s.FollowingsCount.Encode(e)
		}
	}
	{
		if s.Gender.Set {
			e.FieldStart("gender")
			s.Gender.Encode(e)
		}
	}
	{
		if s.GroupsUsersCount.Set {
			e.FieldStart("groups_users_count")
			s.GroupsUsersCount.Encode(e)
		}
	}
	{
		if s.Hidden.Set {
			e.FieldStart("hidden")
			s.Hidden.Encode(e)
		}
	}
	{
		if s.HideVip.Set {
			e.FieldStart("hide_vip")
			s.HideVip.Encode(e)
		}
	}
	{
		if s.ID.Set {
			e.FieldStart("id")
			s.ID.Encode(e)
		}
	}
	{
		if s.IsPrivate.Set {
			e.FieldStart("is_private")
			s.IsPrivate.Encode(e)
		}
	}
	{
		if s.NewUser.Set {
			e.FieldStart("new_user")
			s.NewUser.Encode(e)
		}
	}
	{
		if s.Nickname.Set {
			e.FieldStart("nickname")
			s.Nickname.Encode(e)
		}
	}
	{
		if s.PostsCount.Set {
			e.FieldStart("posts_count")
			s.PostsCount.Encode(e)
		}
	}
	{
		if s.Prefecture.Set {
			e.FieldStart("prefecture")
			s.Prefecture.Encode(e)
		}
	}
	{
		if s.ProfileIcon.Set {
			e.FieldStart("profile_icon")
			s.ProfileIcon.Encode(e)
		}
	}
	{
		if s.ProfileIconThumbnail.Set {
			e.FieldStart("profile_icon_thumbnail")
			s.ProfileIconThumbnail.Encode(e)
		}
	}
	{
		if s.RecentlyKenta.Set {
			e.FieldStart("recently_kenta")
			s.RecentlyKenta.Encode(e)
		}
	}
	{
		if s.ReviewsCount.Set {
			e.FieldStart("reviews_count")
			s.ReviewsCount.Encode(e)
		}
	}
	{
		if s.Title.Set {
			e.FieldStart("title")
			s.Title.Encode(e)
		}
	}
	{
		if s.Vip.Set {
			e.FieldStart("vip")
			s.Vip.Encode(e)
		}
	}
}

var jsonFieldsNameOfUserUserDTO = [27]string{
	0:  "age_verified",
	1:  "biography",
	2:  "cover_image",
	3:  "cover_image_thumbnail",
	4:  "created_at",
	5:  "dangerous_user",
	6:  "follow_pending",
	7:  "followed_by",
	8:  "followers_count",
	9:  "following",
	10: "followings_count",
	11: "gender",
	12: "groups_users_count",
	13: "hidden",
	14: "hide_vip",
	15: "id",
	16: "is_private",
	17: "new_user",
	18: "nickname",
	19: "posts_count",
	20: "prefecture",
	21: "profile_icon",
	22: "profile_icon_thumbnail",
	23: "recently_kenta",
	24: "reviews_count",
	25: "title",
	26: "vip",
}

// Decode decodes UserUserDTO from json.
func (s *UserUserDTO) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode UserUserDTO to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "age_verified":
			if err := func() error {
				s.AgeVerified.Reset()
				if err := s.AgeVerified.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"age_verified\"")
			}
		case "biography":
			if err := func() error {
				s.Biography.Reset()
				if err := s.Biography.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"biography\"")
			}
		case "cover_image":
			if err := func() error {
				s.CoverImage.Reset()
				if err := s.CoverImage.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"cover_image\"")
			}
		case "cover_image_thumbnail":
			if err := func() error {
				s.CoverImageThumbnail.Reset()
				if err := s.CoverImageThumbnail.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"cover_image_thumbnail\"")
			}
		case "created_at":
			if err := func() error {
				s.CreatedAt.Reset()
				if err := s.CreatedAt.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"created_at\"")
			}
		case "dangerous_user":
			if err := func() error {
				s.DangerousUser.Reset()
				if err := s.DangerousUser.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"dangerous_user\"")
			}
		case "follow_pending":
			if err := func() error {
				s.FollowPending.Reset()
				if err := s.FollowPending.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"follow_pending\"")
			}
		case "followed_by":
			if err := func() error {
				s.FollowedBy.Reset()
				if err := s.FollowedBy.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"followed_by\"")
			}
		case "followers_count":
			if err := func() error {
				s.FollowersCount.Reset()
				if err := s.FollowersCount.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"followers_count\"")
			}
		case "following":
			if err := func() error {
				s.Following.Reset()
				if err := s.Following.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"following\"")
			}
		case "followings_count":
			if err := func() error {
				s.FollowingsCount.Reset()
				if err := s.FollowingsCount.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"followings_count\"")
			}
		case "gender":
			if err := func() error {
				s.Gender.Reset()
				if err := s.Gender.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"gender\"")
			}
		case "groups_users_count":
			if err := func() error {
				s.GroupsUsersCount.Reset()
				if err := s.GroupsUsersCount.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"groups_users_count\"")
			}
		case "hidden":
			if err := func() error {
				s.Hidden.Reset()
				if err := s.Hidden.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"hidden\"")
			}
		case "hide_vip":
			if err := func() error {
				s.HideVip.Reset()
				if err := s.HideVip.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"hide_vip\"")
			}
		case "id":
			if err := func() error {
				s.ID.Reset()
				if err := s.ID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"id\"")
			}
		case "is_private":
			if err := func() error {
				s.IsPrivate.Reset()
				if err := s.IsPrivate.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"is_private\"")
			}
		case "new_user":
			if err := func() error {
				s.NewUser.Reset()
				if err := s.NewUser.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"new_user\"")
			}
		case "nickname":
			if err := func() error {
				s.Nickname.Reset()
				if err := s.Nickname.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"nickname\"")
			}
		case "posts_count":
			if err := func() error {
				s.PostsCount.Reset()
				if err := s.PostsCount.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"posts_count\"")
			}
		case "prefecture":
			if err := func() error {
				s.Prefecture.Reset()
				if err := s.Prefecture.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"prefecture\"")
			}
		case "profile_icon":
			if err := func() error {
				s.ProfileIcon.Reset()
				if err := s.ProfileIcon.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"profile_icon\"")
			}
		case "profile_icon_thumbnail":
			if err := func() error {
				s.ProfileIconThumbnail.Reset()
				if err := s.ProfileIconThumbnail.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"profile_icon_thumbnail\"")
			}
		case "recently_kenta":
			if err := func() error {
				s.RecentlyKenta.Reset()
				if err := s.RecentlyKenta.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"recently_kenta\"")
			}
		case "reviews_count":
			if err := func() error {
				s.ReviewsCount.Reset()
				if err := s.ReviewsCount.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"reviews_count\"")
			}
		case "title":
			if err := func() error {
				s.Title.Reset()
				if err := s.Title.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"title\"")
			}
		case "vip":
			if err := func() error {
				s.Vip.Reset()
				if err := s.Vip.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"vip\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode UserUserDTO")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *UserUserDTO) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *UserUserDTO) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *UserWrapper) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *UserWrapper) encodeFields(e *jx.Encoder) {
	{
		if s.ID.Set {
			e.FieldStart("id")
			s.ID.Encode(e)
		}
	}
	{
		if s.User.Set {
			e.FieldStart("user")
			s.User.Encode(e)
		}
	}
}

var jsonFieldsNameOfUserWrapper = [2]string{
	0: "id",
	1: "user",
}

// Decode decodes UserWrapper from json.
func (s *UserWrapper) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode UserWrapper to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "id":
			if err := func() error {
				s.ID.Reset()
				if err := s.ID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"id\"")
			}
		case "user":
			if err := func() error {
				s.User.Reset()
				if err := s.User.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"user\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode UserWrapper")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *UserWrapper) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *UserWrapper) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *UsersByTimestampResponse) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *UsersByTimestampResponse) encodeFields(e *jx.Encoder) {
	{
		if s.LastTimestamp.Set {
			e.FieldStart("last_timestamp")
			s.LastTimestamp.Encode(e)
		}
	}
	{
		if s.Users.Set {
			e.FieldStart("users")
			s.Users.Encode(e)
		}
	}
}

var jsonFieldsNameOfUsersByTimestampResponse = [2]string{
	0: "last_timestamp",
	1: "users",
}

// Decode decodes UsersByTimestampResponse from json.
func (s *UsersByTimestampResponse) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode UsersByTimestampResponse to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "last_timestamp":
			if err := func() error {
				s.LastTimestamp.Reset()
				if err := s.LastTimestamp.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"last_timestamp\"")
			}
		case "users":
			if err := func() error {
				s.Users.Reset()
				if err := s.Users.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"users\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode UsersByTimestampResponse")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *UsersByTimestampResponse) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *UsersByTimestampResponse) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *UsersResponse) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *UsersResponse) encodeFields(e *jx.Encoder) {
	{
		if s.NextPageValue.Set {
			e.FieldStart("next_page_value")
			s.NextPageValue.Encode(e)
		}
	}
	{
		if s.Users.Set {
			e.FieldStart("users")
			s.Users.Encode(e)
		}
	}
}

var jsonFieldsNameOfUsersResponse = [2]string{
	0: "next_page_value",
	1: "users",
}

// Decode decodes UsersResponse from json.
func (s *UsersResponse) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode UsersResponse to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "next_page_value":
			if err := func() error {
				s.NextPageValue.Reset()
				if err := s.NextPageValue.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"next_page_value\"")
			}
		case "users":
			if err := func() error {
				s.Users.Reset()
				if err := s.Users.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"users\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode UsersResponse")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *UsersResponse) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *UsersResponse) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *ValidationPostResponse) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *ValidationPostResponse) encodeFields(e *jx.Encoder) {
	{
		if s.IsAllowToPost.Set {
			e.FieldStart("is_allow_to_post")
			s.IsAllowToPost.Encode(e)
		}
	}
}

var jsonFieldsNameOfValidationPostResponse = [1]string{
	0: "is_allow_to_post",
}

// Decode decodes ValidationPostResponse from json.
func (s *ValidationPostResponse) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode ValidationPostResponse to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "is_allow_to_post":
			if err := func() error {
				s.IsAllowToPost.Reset()
				if err := s.IsAllowToPost.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"is_allow_to_post\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode ValidationPostResponse")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *ValidationPostResponse) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *ValidationPostResponse) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *Video) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *Video) encodeFields(e *jx.Encoder) {
	{
		if s.Bitrate.Set {
			e.FieldStart("bitrate")
			s.Bitrate.Encode(e)
		}
	}
	{
		if s.Completed.Set {
			e.FieldStart("completed")
			s.Completed.Encode(e)
		}
	}
	{
		if s.Height.Set {
			e.FieldStart("height")
			s.Height.Encode(e)
		}
	}
	{
		if s.ID.Set {
			e.FieldStart("id")
			s.ID.Encode(e)
		}
	}
	{
		if s.ThumbnailBigURL.Set {
			e.FieldStart("thumbnail_big_url")
			s.ThumbnailBigURL.Encode(e)
		}
	}
	{
		if s.ThumbnailURL.Set {
			e.FieldStart("thumbnail_url")
			s.ThumbnailURL.Encode(e)
		}
	}
	{
		if s.VideoURL.Set {
			e.FieldStart("video_url")
			s.VideoURL.Encode(e)
		}
	}
	{
		if s.ViewsCount.Set {
			e.FieldStart("views_count")
			s.ViewsCount.Encode(e)
		}
	}
	{
		if s.Width.Set {
			e.FieldStart("width")
			s.Width.Encode(e)
		}
	}
}

var jsonFieldsNameOfVideo = [9]string{
	0: "bitrate",
	1: "completed",
	2: "height",
	3: "id",
	4: "thumbnail_big_url",
	5: "thumbnail_url",
	6: "video_url",
	7: "views_count",
	8: "width",
}

// Decode decodes Video from json.
func (s *Video) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Video to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "bitrate":
			if err := func() error {
				s.Bitrate.Reset()
				if err := s.Bitrate.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"bitrate\"")
			}
		case "completed":
			if err := func() error {
				s.Completed.Reset()
				if err := s.Completed.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"completed\"")
			}
		case "height":
			if err := func() error {
				s.Height.Reset()
				if err := s.Height.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"height\"")
			}
		case "id":
			if err := func() error {
				s.ID.Reset()
				if err := s.ID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"id\"")
			}
		case "thumbnail_big_url":
			if err := func() error {
				s.ThumbnailBigURL.Reset()
				if err := s.ThumbnailBigURL.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"thumbnail_big_url\"")
			}
		case "thumbnail_url":
			if err := func() error {
				s.ThumbnailURL.Reset()
				if err := s.ThumbnailURL.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"thumbnail_url\"")
			}
		case "video_url":
			if err := func() error {
				s.VideoURL.Reset()
				if err := s.VideoURL.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"video_url\"")
			}
		case "views_count":
			if err := func() error {
				s.ViewsCount.Reset()
				if err := s.ViewsCount.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"views_count\"")
			}
		case "width":
			if err := func() error {
				s.Width.Reset()
				if err := s.Width.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"width\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode Video")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *Video) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *Video) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *VoteSurveyResponse) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *VoteSurveyResponse) encodeFields(e *jx.Encoder) {
	{
		if s.Survey.Set {
			e.FieldStart("survey")
			s.Survey.Encode(e)
		}
	}
}

var jsonFieldsNameOfVoteSurveyResponse = [1]string{
	0: "survey",
}

// Decode decodes VoteSurveyResponse from json.
func (s *VoteSurveyResponse) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode VoteSurveyResponse to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "survey":
			if err := func() error {
				s.Survey.Reset()
				if err := s.Survey.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"survey\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode VoteSurveyResponse")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *VoteSurveyResponse) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *VoteSurveyResponse) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *Walkthrough) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *Walkthrough) encodeFields(e *jx.Encoder) {
	{
		if s.Title.Set {
			e.FieldStart("title")
			s.Title.Encode(e)
		}
	}
	{
		if s.URL.Set {
			e.FieldStart("url")
			s.URL.Encode(e)
		}
	}
}

var jsonFieldsNameOfWalkthrough = [2]string{
	0: "title",
	1: "url",
}

// Decode decodes Walkthrough from json.
func (s *Walkthrough) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Walkthrough to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "title":
			if err := func() error {
				s.Title.Reset()
				if err := s.Title.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"title\"")
			}
		case "url":
			if err := func() error {
				s.URL.Reset()
				if err := s.URL.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"url\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode Walkthrough")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *Walkthrough) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *Walkthrough) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *WebSocketTokenResponse) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *WebSocketTokenResponse) encodeFields(e *jx.Encoder) {
	{
		if s.Token.Set {
			e.FieldStart("token")
			s.Token.Encode(e)
		}
	}
}

var jsonFieldsNameOfWebSocketTokenResponse = [1]string{
	0: "token",
}

// Decode decodes WebSocketTokenResponse from json.
func (s *WebSocketTokenResponse) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode WebSocketTokenResponse to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "token":
			if err := func() error {
				s.Token.Reset()
				if err := s.Token.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"token\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode WebSocketTokenResponse")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *WebSocketTokenResponse) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *WebSocketTokenResponse) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}
