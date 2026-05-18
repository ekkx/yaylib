// Code generated; DO NOT EDIT.

package oas

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/ogen-go/ogen/uri"
)

var (
	rn269AllowedHeaders = map[string]string{
		"POST": "Content-Type",
	}
	rn77AllowedHeaders = map[string]string{
		"POST": "Content-Type",
	}
	rn226AllowedHeaders = map[string]string{
		"POST": "Content-Type",
	}
	rn233AllowedHeaders = map[string]string{
		"POST": "Content-Type",
	}
	rn235AllowedHeaders = map[string]string{
		"POST": "Content-Type",
	}
	rn22AllowedHeaders = map[string]string{
		"PUT": "Content-Type",
	}
	rn339AllowedHeaders = map[string]string{
		"PUT": "Content-Type",
	}
	rn1AllowedHeaders = map[string]string{
		"POST": "Content-Type",
	}
	rn54AllowedHeaders = map[string]string{
		"POST": "Content-Type",
	}
	rn35AllowedHeaders = map[string]string{
		"POST": "Content-Type",
	}
	rn282AllowedHeaders = map[string]string{
		"POST": "Content-Type",
	}
	rn286AllowedHeaders = map[string]string{
		"POST": "Content-Type",
	}
	rn304AllowedHeaders = map[string]string{
		"POST": "Content-Type",
	}
	rn227AllowedHeaders = map[string]string{
		"POST": "Content-Type",
	}
	rn317AllowedHeaders = map[string]string{
		"POST": "Content-Type",
	}
	rn362AllowedHeaders = map[string]string{
		"PUT": "Content-Type",
	}
	rn218AllowedHeaders = map[string]string{
		"POST": "Content-Type",
	}
	rn219AllowedHeaders = map[string]string{
		"POST": "Content-Type",
	}
	rn39AllowedHeaders = map[string]string{
		"POST": "Content-Type",
	}
	rn273AllowedHeaders = map[string]string{
		"POST": "Content-Type",
	}
	rn279AllowedHeaders = map[string]string{
		"POST": "Content-Type",
	}
	rn280AllowedHeaders = map[string]string{
		"POST": "Content-Type",
	}
	rn177AllowedHeaders = map[string]string{
		"POST": "Content-Type",
	}
	rn349AllowedHeaders = map[string]string{
		"POST": "Content-Type",
	}
	rn44AllowedHeaders = map[string]string{
		"POST": "Content-Type",
	}
	rn10AllowedHeaders = map[string]string{
		"PUT": "Content-Type",
	}
	rn46AllowedHeaders = map[string]string{
		"POST": "Content-Type,X-Jwt",
	}
	rn82AllowedHeaders = map[string]string{
		"POST": "Content-Type",
	}
	rn26AllowedHeaders = map[string]string{
		"PUT": "Content-Type",
	}
	rn28AllowedHeaders = map[string]string{
		"PUT": "Content-Type",
	}
	rn108AllowedHeaders = map[string]string{
		"POST": "Content-Type",
	}
	rn207AllowedHeaders = map[string]string{
		"PUT": "Content-Type",
	}
	rn345AllowedHeaders = map[string]string{
		"PUT": "Content-Type",
	}
	rn215AllowedHeaders = map[string]string{
		"GET": "X-Jwt",
	}
	rn259AllowedHeaders = map[string]string{
		"POST": "Content-Type",
	}
	rn185AllowedHeaders = map[string]string{
		"POST": "Content-Type",
	}
	rn309AllowedHeaders = map[string]string{
		"PUT": "Content-Type",
	}
	rn41AllowedHeaders = map[string]string{
		"POST": "Content-Type",
	}
	rn67AllowedHeaders = map[string]string{
		"PUT": "Content-Type",
	}
	rn71AllowedHeaders = map[string]string{
		"PUT": "Content-Type",
	}
	rn221AllowedHeaders = map[string]string{
		"POST": "Content-Type",
	}
	rn319AllowedHeaders = map[string]string{
		"POST": "Content-Type",
	}
	rn223AllowedHeaders = map[string]string{
		"POST": "Content-Type",
	}
	rn229AllowedHeaders = map[string]string{
		"POST": "Content-Type",
	}
	rn343AllowedHeaders = map[string]string{
		"POST": "Content-Type",
	}
	rn138AllowedHeaders = map[string]string{
		"POST": "Content-Type",
	}
	rn277AllowedHeaders = map[string]string{
		"PUT": "Content-Type",
	}
	rn237AllowedHeaders = map[string]string{
		"POST": "Content-Type",
	}
	rn61AllowedHeaders = map[string]string{
		"POST": "Content-Type",
	}
	rn36AllowedHeaders = map[string]string{
		"POST": "Content-Type",
	}
	rn43AllowedHeaders = map[string]string{
		"POST": "Content-Type",
	}
	rn161AllowedHeaders = map[string]string{
		"GET": "Cache-Control",
	}
	rn358AllowedHeaders = map[string]string{
		"POST": "Content-Type",
	}
	rn93AllowedHeaders = map[string]string{
		"POST": "Content-Type",
	}
	rn69AllowedHeaders = map[string]string{
		"POST": "Content-Type",
	}
	rn211AllowedHeaders = map[string]string{
		"POST": "Content-Type",
	}
	rn305AllowedHeaders = map[string]string{
		"POST": "Content-Type",
	}
	rn232AllowedHeaders = map[string]string{
		"POST": "Content-Type",
	}
	rn33AllowedHeaders = map[string]string{
		"POST": "Content-Type",
	}
	rn340AllowedHeaders = map[string]string{
		"POST": "Content-Type",
	}
	rn316AllowedHeaders = map[string]string{
		"POST": "Content-Type",
	}
	rn297AllowedHeaders = map[string]string{
		"POST": "Content-Type",
	}
	rn38AllowedHeaders = map[string]string{
		"POST": "Content-Type",
	}
	rn66AllowedHeaders = map[string]string{
		"POST": "Content-Type",
	}
	rn298AllowedHeaders = map[string]string{
		"POST": "Content-Type",
	}
	rn322AllowedHeaders = map[string]string{
		"POST": "Content-Type",
	}
	rn344AllowedHeaders = map[string]string{
		"POST": "Content-Type",
	}
	rn40AllowedHeaders = map[string]string{
		"POST": "Content-Type,X-Jwt",
	}
	rn303AllowedHeaders = map[string]string{
		"POST": "Content-Type,X-Jwt",
	}
	rn261AllowedHeaders = map[string]string{
		"PUT": "Content-Type",
	}
	rn264AllowedHeaders = map[string]string{
		"POST": "Content-Type",
	}
	rn300AllowedHeaders = map[string]string{
		"POST": "Content-Type",
	}
	rn68AllowedHeaders = map[string]string{
		"POST": "Content-Type",
	}
	rn347AllowedHeaders = map[string]string{
		"POST": "Content-Type",
	}
	rn257AllowedHeaders = map[string]string{
		"POST": "Content-Type",
	}
	rn302AllowedHeaders = map[string]string{
		"POST": "Content-Type",
	}
)

func (s *Server) cutPrefix(path string) (string, bool) {
	prefix := s.cfg.Prefix
	if prefix == "" {
		return path, true
	}
	if !strings.HasPrefix(path, prefix) {
		// Prefix doesn't match.
		return "", false
	}
	// Cut prefix from the path.
	return strings.TrimPrefix(path, prefix), true
}

// ServeHTTP serves http request as defined by OpenAPI v3 specification,
// calling handler that matches the path or returning not found error.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	elem := r.URL.Path
	elemIsEscaped := false
	if rawPath := r.URL.RawPath; rawPath != "" {
		if normalized, ok := uri.NormalizeEscapedPath(rawPath); ok {
			elem = normalized
			elemIsEscaped = strings.ContainsRune(elem, '%')
		}
	}

	elem, ok := s.cutPrefix(elem)
	if !ok || len(elem) == 0 {
		s.notFound(w, r)
		return
	}
	args := [2]string{}

	// Static code generated router with unwrapped path search.
	switch {
	default:
		if len(elem) == 0 {
			break
		}
		switch elem[0] {
		case '/': // Prefix: "/"

			if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
				elem = elem[l:]
			} else {
				break
			}

			if len(elem) == 0 {
				break
			}
			switch elem[0] {
			case 'a': // Prefix: "api/"
				origElem := elem
				if l := len("api/"); len(elem) >= l && elem[0:l] == "api/" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'a': // Prefix: "apps/"

					if l := len("apps/"); len(elem) >= l && elem[0:l] == "apps/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "app"
					// Leaf parameter, slashes are prohibited
					idx := strings.IndexByte(elem, '/')
					if idx >= 0 {
						break
					}
					args[0] = elem
					elem = ""

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "GET":
							s.handleGetAppConfigRequest([1]string{
								args[0],
							}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, notAllowedParams{
								allowedMethods: "GET",
								allowedHeaders: nil,
								acceptPost:     "",
								acceptPatch:    "",
							})
						}

						return
					}

				case 'u': // Prefix: "user_activities"

					if l := len("user_activities"); len(elem) >= l && elem[0:l] == "user_activities" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "GET":
							s.handleGetUserActivitiesV1Request([0]string{}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, notAllowedParams{
								allowedMethods: "GET",
								allowedHeaders: nil,
								acceptPost:     "",
								acceptPatch:    "",
							})
						}

						return
					}

				case 'v': // Prefix: "v"

					if l := len("v"); len(elem) >= l && elem[0:l] == "v" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						break
					}
					switch elem[0] {
					case '1': // Prefix: "1/oauth/token"

						if l := len("1/oauth/token"); len(elem) >= l && elem[0:l] == "1/oauth/token" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch r.Method {
							case "POST":
								s.handleOauthTokenRequest([0]string{}, elemIsEscaped, w, r)
							default:
								s.notAllowed(w, r, notAllowedParams{
									allowedMethods: "POST",
									allowedHeaders: rn269AllowedHeaders,
									acceptPost:     "application/x-www-form-urlencoded",
									acceptPatch:    "",
								})
							}

							return
						}

					case '2': // Prefix: "2/user_activities"

						if l := len("2/user_activities"); len(elem) >= l && elem[0:l] == "2/user_activities" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch r.Method {
							case "GET":
								s.handleGetUserActivitiesRequest([0]string{}, elemIsEscaped, w, r)
							default:
								s.notAllowed(w, r, notAllowedParams{
									allowedMethods: "GET",
									allowedHeaders: nil,
									acceptPost:     "",
									acceptPatch:    "",
								})
							}

							return
						}

					}

				}

				elem = origElem
			case 'v': // Prefix: "v"
				origElem := elem
				if l := len("v"); len(elem) >= l && elem[0:l] == "v" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case '1': // Prefix: "1/"

					if l := len("1/"); len(elem) >= l && elem[0:l] == "1/" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						break
					}
					switch elem[0] {
					case 'b': // Prefix: "buckets/presigned_urls"

						if l := len("buckets/presigned_urls"); len(elem) >= l && elem[0:l] == "buckets/presigned_urls" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch r.Method {
							case "GET":
								s.handleGetBucketPresignedUrlsRequest([0]string{}, elemIsEscaped, w, r)
							default:
								s.notAllowed(w, r, notAllowedParams{
									allowedMethods: "GET",
									allowedHeaders: nil,
									acceptPost:     "",
									acceptPatch:    "",
								})
							}

							return
						}

					case 'c': // Prefix: "c"

						if l := len("c"); len(elem) >= l && elem[0:l] == "c" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							break
						}
						switch elem[0] {
						case 'a': // Prefix: "alls/"

							if l := len("alls/"); len(elem) >= l && elem[0:l] == "alls/" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								break
							}
							switch elem[0] {
							case 'a': // Prefix: "action_signature/generate"
								origElem := elem
								if l := len("action_signature/generate"); len(elem) >= l && elem[0:l] == "action_signature/generate" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch r.Method {
									case "POST":
										s.handleGenerateCallActionSignatureRequest([0]string{}, elemIsEscaped, w, r)
									default:
										s.notAllowed(w, r, notAllowedParams{
											allowedMethods: "POST",
											allowedHeaders: rn77AllowedHeaders,
											acceptPost:     "multipart/form-data",
											acceptPatch:    "",
										})
									}

									return
								}

								elem = origElem
							case 'b': // Prefix: "bgm"
								origElem := elem
								if l := len("bgm"); len(elem) >= l && elem[0:l] == "bgm" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch r.Method {
									case "GET":
										s.handleGetCallBgmsRequest([0]string{}, elemIsEscaped, w, r)
									default:
										s.notAllowed(w, r, notAllowedParams{
											allowedMethods: "GET",
											allowedHeaders: nil,
											acceptPost:     "",
											acceptPatch:    "",
										})
									}

									return
								}

								elem = origElem
							case 'c': // Prefix: "conference_calls/"
								origElem := elem
								if l := len("conference_calls/"); len(elem) >= l && elem[0:l] == "conference_calls/" {
									elem = elem[l:]
								} else {
									break
								}

								// Param: "call_id"
								// Match until "/"
								idx := strings.IndexByte(elem, '/')
								if idx < 0 {
									idx = len(elem)
								}
								args[0] = elem[:idx]
								elem = elem[idx:]

								if len(elem) == 0 {
									break
								}
								switch elem[0] {
								case '/': // Prefix: "/invite"

									if l := len("/invite"); len(elem) >= l && elem[0:l] == "/invite" {
										elem = elem[l:]
									} else {
										break
									}

									if len(elem) == 0 {
										// Leaf node.
										switch r.Method {
										case "POST":
											s.handleInviteToConferenceCallRequest([1]string{
												args[0],
											}, elemIsEscaped, w, r)
										default:
											s.notAllowed(w, r, notAllowedParams{
												allowedMethods: "POST",
												allowedHeaders: rn226AllowedHeaders,
												acceptPost:     "multipart/form-data",
												acceptPatch:    "",
											})
										}

										return
									}

								}

								elem = origElem
							case 'l': // Prefix: "leave_"
								origElem := elem
								if l := len("leave_"); len(elem) >= l && elem[0:l] == "leave_" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									break
								}
								switch elem[0] {
								case 'a': // Prefix: "agora_channel"

									if l := len("agora_channel"); len(elem) >= l && elem[0:l] == "agora_channel" {
										elem = elem[l:]
									} else {
										break
									}

									if len(elem) == 0 {
										// Leaf node.
										switch r.Method {
										case "POST":
											s.handleLeaveAgoraChannelRequest([0]string{}, elemIsEscaped, w, r)
										default:
											s.notAllowed(w, r, notAllowedParams{
												allowedMethods: "POST",
												allowedHeaders: rn233AllowedHeaders,
												acceptPost:     "multipart/form-data",
												acceptPatch:    "",
											})
										}

										return
									}

								case 'c': // Prefix: "conference_call"

									if l := len("conference_call"); len(elem) >= l && elem[0:l] == "conference_call" {
										elem = elem[l:]
									} else {
										break
									}

									if len(elem) == 0 {
										// Leaf node.
										switch r.Method {
										case "POST":
											s.handleLeaveConferenceCallRequest([0]string{}, elemIsEscaped, w, r)
										default:
											s.notAllowed(w, r, notAllowedParams{
												allowedMethods: "POST",
												allowedHeaders: rn235AllowedHeaders,
												acceptPost:     "multipart/form-data",
												acceptPatch:    "",
											})
										}

										return
									}

								}

								elem = origElem
							case 'p': // Prefix: "phone_status/"
								origElem := elem
								if l := len("phone_status/"); len(elem) >= l && elem[0:l] == "phone_status/" {
									elem = elem[l:]
								} else {
									break
								}

								// Param: "opponent_id"
								// Leaf parameter, slashes are prohibited
								idx := strings.IndexByte(elem, '/')
								if idx >= 0 {
									break
								}
								args[0] = elem
								elem = ""

								if len(elem) == 0 {
									// Leaf node.
									switch r.Method {
									case "GET":
										s.handleGetPhoneStatusRequest([1]string{
											args[0],
										}, elemIsEscaped, w, r)
									default:
										s.notAllowed(w, r, notAllowedParams{
											allowedMethods: "GET",
											allowedHeaders: nil,
											acceptPost:     "",
											acceptPatch:    "",
										})
									}

									return
								}

								elem = origElem
							}
							// Param: "call_id"
							// Match until "/"
							idx := strings.IndexByte(elem, '/')
							if idx < 0 {
								idx = len(elem)
							}
							args[0] = elem[:idx]
							elem = elem[idx:]

							if len(elem) == 0 {
								switch r.Method {
								case "PUT":
									s.handleUpdateCallRequest([1]string{
										args[0],
									}, elemIsEscaped, w, r)
								default:
									s.notAllowed(w, r, notAllowedParams{
										allowedMethods: "PUT",
										allowedHeaders: rn22AllowedHeaders,
										acceptPost:     "",
										acceptPatch:    "",
									})
								}

								return
							}
							switch elem[0] {
							case '/': // Prefix: "/"

								if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									break
								}
								switch elem[0] {
								case 'b': // Prefix: "bu"

									if l := len("bu"); len(elem) >= l && elem[0:l] == "bu" {
										elem = elem[l:]
									} else {
										break
									}

									if len(elem) == 0 {
										break
									}
									switch elem[0] {
									case 'l': // Prefix: "lk_invite"

										if l := len("lk_invite"); len(elem) >= l && elem[0:l] == "lk_invite" {
											elem = elem[l:]
										} else {
											break
										}

										if len(elem) == 0 {
											// Leaf node.
											switch r.Method {
											case "POST":
												s.handleBulkInviteToCallRequest([1]string{
													args[0],
												}, elemIsEscaped, w, r)
											default:
												s.notAllowed(w, r, notAllowedParams{
													allowedMethods: "POST",
													allowedHeaders: nil,
													acceptPost:     "",
													acceptPatch:    "",
												})
											}

											return
										}

									case 'm': // Prefix: "mp"

										if l := len("mp"); len(elem) >= l && elem[0:l] == "mp" {
											elem = elem[l:]
										} else {
											break
										}

										if len(elem) == 0 {
											// Leaf node.
											switch r.Method {
											case "POST":
												s.handleBumpCallRequest([1]string{
													args[0],
												}, elemIsEscaped, w, r)
											default:
												s.notAllowed(w, r, notAllowedParams{
													allowedMethods: "POST",
													allowedHeaders: nil,
													acceptPost:     "",
													acceptPatch:    "",
												})
											}

											return
										}

									}

								case 'g': // Prefix: "gift_transactions"

									if l := len("gift_transactions"); len(elem) >= l && elem[0:l] == "gift_transactions" {
										elem = elem[l:]
									} else {
										break
									}

									if len(elem) == 0 {
										// Leaf node.
										switch r.Method {
										case "GET":
											s.handleGetCallGiftHistoryRequest([1]string{
												args[0],
											}, elemIsEscaped, w, r)
										default:
											s.notAllowed(w, r, notAllowedParams{
												allowedMethods: "GET",
												allowedHeaders: nil,
												acceptPost:     "",
												acceptPatch:    "",
											})
										}

										return
									}

								case 'u': // Prefix: "users/"

									if l := len("users/"); len(elem) >= l && elem[0:l] == "users/" {
										elem = elem[l:]
									} else {
										break
									}

									if len(elem) == 0 {
										break
									}
									switch elem[0] {
									case 'i': // Prefix: "invitable"
										origElem := elem
										if l := len("invitable"); len(elem) >= l && elem[0:l] == "invitable" {
											elem = elem[l:]
										} else {
											break
										}

										if len(elem) == 0 {
											// Leaf node.
											switch r.Method {
											case "GET":
												s.handleGetInvitableCallUsersRequest([1]string{
													args[0],
												}, elemIsEscaped, w, r)
											default:
												s.notAllowed(w, r, notAllowedParams{
													allowedMethods: "GET",
													allowedHeaders: nil,
													acceptPost:     "",
													acceptPatch:    "",
												})
											}

											return
										}

										elem = origElem
									}
									// Param: "user_id"
									// Leaf parameter, slashes are prohibited
									idx := strings.IndexByte(elem, '/')
									if idx >= 0 {
										break
									}
									args[1] = elem
									elem = ""

									if len(elem) == 0 {
										// Leaf node.
										switch r.Method {
										case "PUT":
											s.handleUpdateCallUserRequest([2]string{
												args[0],
												args[1],
											}, elemIsEscaped, w, r)
										default:
											s.notAllowed(w, r, notAllowedParams{
												allowedMethods: "PUT",
												allowedHeaders: rn339AllowedHeaders,
												acceptPost:     "",
												acceptPatch:    "",
											})
										}

										return
									}

								}

							}

						case 'h': // Prefix: "hat_rooms/"

							if l := len("hat_rooms/"); len(elem) >= l && elem[0:l] == "hat_rooms/" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								break
							}
							switch elem[0] {
							case 'a': // Prefix: "accept_chat_request"
								origElem := elem
								if l := len("accept_chat_request"); len(elem) >= l && elem[0:l] == "accept_chat_request" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch r.Method {
									case "POST":
										s.handleAcceptChatRequestRequest([0]string{}, elemIsEscaped, w, r)
									default:
										s.notAllowed(w, r, notAllowedParams{
											allowedMethods: "POST",
											allowedHeaders: rn1AllowedHeaders,
											acceptPost:     "multipart/form-data",
											acceptPatch:    "",
										})
									}

									return
								}

								elem = origElem
							case 'm': // Prefix: "ma"
								origElem := elem
								if l := len("ma"); len(elem) >= l && elem[0:l] == "ma" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									break
								}
								switch elem[0] {
								case 'i': // Prefix: "in_list"

									if l := len("in_list"); len(elem) >= l && elem[0:l] == "in_list" {
										elem = elem[l:]
									} else {
										break
									}

									if len(elem) == 0 {
										// Leaf node.
										switch r.Method {
										case "GET":
											s.handleGetMainChatRoomsRequest([0]string{}, elemIsEscaped, w, r)
										default:
											s.notAllowed(w, r, notAllowedParams{
												allowedMethods: "GET",
												allowedHeaders: nil,
												acceptPost:     "",
												acceptPatch:    "",
											})
										}

										return
									}

								case 's': // Prefix: "ss_destroy"

									if l := len("ss_destroy"); len(elem) >= l && elem[0:l] == "ss_destroy" {
										elem = elem[l:]
									} else {
										break
									}

									if len(elem) == 0 {
										// Leaf node.
										switch r.Method {
										case "POST":
											s.handleDeleteChatRoomsRequest([0]string{}, elemIsEscaped, w, r)
										default:
											s.notAllowed(w, r, notAllowedParams{
												allowedMethods: "POST",
												allowedHeaders: rn54AllowedHeaders,
												acceptPost:     "multipart/form-data",
												acceptPatch:    "",
											})
										}

										return
									}

								}

								elem = origElem
							case 'n': // Prefix: "new"
								origElem := elem
								if l := len("new"); len(elem) >= l && elem[0:l] == "new" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch r.Method {
									case "POST":
										s.handleCreateChatRoomV1Request([0]string{}, elemIsEscaped, w, r)
									default:
										s.notAllowed(w, r, notAllowedParams{
											allowedMethods: "POST",
											allowedHeaders: rn35AllowedHeaders,
											acceptPost:     "multipart/form-data",
											acceptPatch:    "",
										})
									}

									return
								}

								elem = origElem
							case 'r': // Prefix: "request_list"
								origElem := elem
								if l := len("request_list"); len(elem) >= l && elem[0:l] == "request_list" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch r.Method {
									case "GET":
										s.handleGetChatRequestsRequest([0]string{}, elemIsEscaped, w, r)
									default:
										s.notAllowed(w, r, notAllowedParams{
											allowedMethods: "GET",
											allowedHeaders: nil,
											acceptPost:     "",
											acceptPatch:    "",
										})
									}

									return
								}

								elem = origElem
							case 't': // Prefix: "total_chat_request"
								origElem := elem
								if l := len("total_chat_request"); len(elem) >= l && elem[0:l] == "total_chat_request" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch r.Method {
									case "GET":
										s.handleGetChatRequestCountRequest([0]string{}, elemIsEscaped, w, r)
									default:
										s.notAllowed(w, r, notAllowedParams{
											allowedMethods: "GET",
											allowedHeaders: nil,
											acceptPost:     "",
											acceptPatch:    "",
										})
									}

									return
								}

								elem = origElem
							case 'u': // Prefix: "unread_status"
								origElem := elem
								if l := len("unread_status"); len(elem) >= l && elem[0:l] == "unread_status" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch r.Method {
									case "GET":
										s.handleGetChatUnreadStatusRequest([0]string{}, elemIsEscaped, w, r)
									default:
										s.notAllowed(w, r, notAllowedParams{
											allowedMethods: "GET",
											allowedHeaders: nil,
											acceptPost:     "",
											acceptPatch:    "",
										})
									}

									return
								}

								elem = origElem
							}
							// Param: "room_id"
							// Match until "/"
							idx := strings.IndexByte(elem, '/')
							if idx < 0 {
								idx = len(elem)
							}
							args[0] = elem[:idx]
							elem = elem[idx:]

							if len(elem) == 0 {
								break
							}
							switch elem[0] {
							case '/': // Prefix: "/"

								if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									break
								}
								switch elem[0] {
								case 'a': // Prefix: "attachments_read"

									if l := len("attachments_read"); len(elem) >= l && elem[0:l] == "attachments_read" {
										elem = elem[l:]
									} else {
										break
									}

									if len(elem) == 0 {
										// Leaf node.
										switch r.Method {
										case "POST":
											s.handleReadChatAttachmentsRequest([1]string{
												args[0],
											}, elemIsEscaped, w, r)
										default:
											s.notAllowed(w, r, notAllowedParams{
												allowedMethods: "POST",
												allowedHeaders: rn282AllowedHeaders,
												acceptPost:     "application/json",
												acceptPatch:    "",
											})
										}

										return
									}

								case 'm': // Prefix: "messages/"

									if l := len("messages/"); len(elem) >= l && elem[0:l] == "messages/" {
										elem = elem[l:]
									} else {
										break
									}

									// Param: "message_id"
									// Match until "/"
									idx := strings.IndexByte(elem, '/')
									if idx < 0 {
										idx = len(elem)
									}
									args[1] = elem[:idx]
									elem = elem[idx:]

									if len(elem) == 0 {
										break
									}
									switch elem[0] {
									case '/': // Prefix: "/delete"

										if l := len("/delete"); len(elem) >= l && elem[0:l] == "/delete" {
											elem = elem[l:]
										} else {
											break
										}

										if len(elem) == 0 {
											// Leaf node.
											switch r.Method {
											case "DELETE":
												s.handleDeleteChatMessageRequest([2]string{
													args[0],
													args[1],
												}, elemIsEscaped, w, r)
											default:
												s.notAllowed(w, r, notAllowedParams{
													allowedMethods: "DELETE",
													allowedHeaders: nil,
													acceptPost:     "",
													acceptPatch:    "",
												})
											}

											return
										}

									}

								case 'p': // Prefix: "pinned"

									if l := len("pinned"); len(elem) >= l && elem[0:l] == "pinned" {
										elem = elem[l:]
									} else {
										break
									}

									if len(elem) == 0 {
										// Leaf node.
										switch r.Method {
										case "DELETE":
											s.handleUnpinChatRoomRequest([1]string{
												args[0],
											}, elemIsEscaped, w, r)
										case "POST":
											s.handlePinChatRoomRequest([1]string{
												args[0],
											}, elemIsEscaped, w, r)
										default:
											s.notAllowed(w, r, notAllowedParams{
												allowedMethods: "DELETE,POST",
												allowedHeaders: nil,
												acceptPost:     "",
												acceptPatch:    "",
											})
										}

										return
									}

								case 'v': // Prefix: "videos_read"

									if l := len("videos_read"); len(elem) >= l && elem[0:l] == "videos_read" {
										elem = elem[l:]
									} else {
										break
									}

									if len(elem) == 0 {
										// Leaf node.
										switch r.Method {
										case "POST":
											s.handleReadChatVideosRequest([1]string{
												args[0],
											}, elemIsEscaped, w, r)
										default:
											s.notAllowed(w, r, notAllowedParams{
												allowedMethods: "POST",
												allowedHeaders: rn286AllowedHeaders,
												acceptPost:     "multipart/form-data",
												acceptPatch:    "",
											})
										}

										return
									}

								}

							}

						}

					case 'e': // Prefix: "email_verification_urls"

						if l := len("email_verification_urls"); len(elem) >= l && elem[0:l] == "email_verification_urls" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch r.Method {
							case "POST":
								s.handleRequestEmailVerificationRequest([0]string{}, elemIsEscaped, w, r)
							default:
								s.notAllowed(w, r, notAllowedParams{
									allowedMethods: "POST",
									allowedHeaders: rn304AllowedHeaders,
									acceptPost:     "multipart/form-data",
									acceptPatch:    "",
								})
							}

							return
						}

					case 'g': // Prefix: "g"

						if l := len("g"); len(elem) >= l && elem[0:l] == "g" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							break
						}
						switch elem[0] {
						case 'a': // Prefix: "ames/apps"

							if l := len("ames/apps"); len(elem) >= l && elem[0:l] == "ames/apps" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								switch r.Method {
								case "GET":
									s.handleListGameAppsRequest([0]string{}, elemIsEscaped, w, r)
								default:
									s.notAllowed(w, r, notAllowedParams{
										allowedMethods: "GET",
										allowedHeaders: nil,
										acceptPost:     "",
										acceptPatch:    "",
									})
								}

								return
							}
							switch elem[0] {
							case '/': // Prefix: "/"

								if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
									elem = elem[l:]
								} else {
									break
								}

								// Param: "app_id"
								// Match until "/"
								idx := strings.IndexByte(elem, '/')
								if idx < 0 {
									idx = len(elem)
								}
								args[0] = elem[:idx]
								elem = elem[idx:]

								if len(elem) == 0 {
									break
								}
								switch elem[0] {
								case '/': // Prefix: "/walkthroughs"

									if l := len("/walkthroughs"); len(elem) >= l && elem[0:l] == "/walkthroughs" {
										elem = elem[l:]
									} else {
										break
									}

									if len(elem) == 0 {
										// Leaf node.
										switch r.Method {
										case "GET":
											s.handleListGameWalkthroughsRequest([1]string{
												args[0],
											}, elemIsEscaped, w, r)
										default:
											s.notAllowed(w, r, notAllowedParams{
												allowedMethods: "GET",
												allowedHeaders: nil,
												acceptPost:     "",
												acceptPatch:    "",
											})
										}

										return
									}

								}

							}

						case 'e': // Prefix: "enres"

							if l := len("enres"); len(elem) >= l && elem[0:l] == "enres" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch r.Method {
								case "GET":
									s.handleListGenresRequest([0]string{}, elemIsEscaped, w, r)
								default:
									s.notAllowed(w, r, notAllowedParams{
										allowedMethods: "GET",
										allowedHeaders: nil,
										acceptPost:     "",
										acceptPatch:    "",
									})
								}

								return
							}

						case 'r': // Prefix: "roup"

							if l := len("roup"); len(elem) >= l && elem[0:l] == "roup" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								break
							}
							switch elem[0] {
							case '_': // Prefix: "_mute/"

								if l := len("_mute/"); len(elem) >= l && elem[0:l] == "_mute/" {
									elem = elem[l:]
								} else {
									break
								}

								// Param: "id"
								// Match until "/"
								idx := strings.IndexByte(elem, '/')
								if idx < 0 {
									idx = len(elem)
								}
								args[0] = elem[:idx]
								elem = elem[idx:]

								if len(elem) == 0 {
									break
								}
								switch elem[0] {
								case '/': // Prefix: "/"

									if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
										elem = elem[l:]
									} else {
										break
									}

									if len(elem) == 0 {
										break
									}
									switch elem[0] {
									case 'm': // Prefix: "mute"

										if l := len("mute"); len(elem) >= l && elem[0:l] == "mute" {
											elem = elem[l:]
										} else {
											break
										}

										if len(elem) == 0 {
											break
										}
										switch elem[0] {
										case '/': // Prefix: "/"

											if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
												elem = elem[l:]
											} else {
												break
											}

											// Param: "user_id"
											// Leaf parameter, slashes are prohibited
											idx := strings.IndexByte(elem, '/')
											if idx >= 0 {
												break
											}
											args[1] = elem
											elem = ""

											if len(elem) == 0 {
												// Leaf node.
												switch r.Method {
												case "POST":
													s.handleMuteGroupUserRequest([2]string{
														args[0],
														args[1],
													}, elemIsEscaped, w, r)
												default:
													s.notAllowed(w, r, notAllowedParams{
														allowedMethods: "POST",
														allowedHeaders: nil,
														acceptPost:     "",
														acceptPatch:    "",
													})
												}

												return
											}

										case 'd': // Prefix: "d_users"

											if l := len("d_users"); len(elem) >= l && elem[0:l] == "d_users" {
												elem = elem[l:]
											} else {
												break
											}

											if len(elem) == 0 {
												// Leaf node.
												switch r.Method {
												case "GET":
													s.handleListMutedGroupUsersRequest([1]string{
														args[0],
													}, elemIsEscaped, w, r)
												default:
													s.notAllowed(w, r, notAllowedParams{
														allowedMethods: "GET",
														allowedHeaders: nil,
														acceptPost:     "",
														acceptPatch:    "",
													})
												}

												return
											}

										}

									case 'u': // Prefix: "unmute/"

										if l := len("unmute/"); len(elem) >= l && elem[0:l] == "unmute/" {
											elem = elem[l:]
										} else {
											break
										}

										// Param: "user_id"
										// Leaf parameter, slashes are prohibited
										idx := strings.IndexByte(elem, '/')
										if idx >= 0 {
											break
										}
										args[1] = elem
										elem = ""

										if len(elem) == 0 {
											// Leaf node.
											switch r.Method {
											case "DELETE":
												s.handleUnmuteGroupUserRequest([2]string{
													args[0],
													args[1],
												}, elemIsEscaped, w, r)
											default:
												s.notAllowed(w, r, notAllowedParams{
													allowedMethods: "DELETE",
													allowedHeaders: nil,
													acceptPost:     "",
													acceptPatch:    "",
												})
											}

											return
										}

									}

								}

							case 's': // Prefix: "s/"

								if l := len("s/"); len(elem) >= l && elem[0:l] == "s/" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									break
								}
								switch elem[0] {
								case 'c': // Prefix: "c"
									origElem := elem
									if l := len("c"); len(elem) >= l && elem[0:l] == "c" {
										elem = elem[l:]
									} else {
										break
									}

									if len(elem) == 0 {
										break
									}
									switch elem[0] {
									case 'a': // Prefix: "ategories"

										if l := len("ategories"); len(elem) >= l && elem[0:l] == "ategories" {
											elem = elem[l:]
										} else {
											break
										}

										if len(elem) == 0 {
											// Leaf node.
											switch r.Method {
											case "GET":
												s.handleListGroupCategoriesRequest([0]string{}, elemIsEscaped, w, r)
											default:
												s.notAllowed(w, r, notAllowedParams{
													allowedMethods: "GET",
													allowedHeaders: nil,
													acceptPost:     "",
													acceptPatch:    "",
												})
											}

											return
										}

									case 'r': // Prefix: "reated_quota"

										if l := len("reated_quota"); len(elem) >= l && elem[0:l] == "reated_quota" {
											elem = elem[l:]
										} else {
											break
										}

										if len(elem) == 0 {
											// Leaf node.
											switch r.Method {
											case "GET":
												s.handleGetGroupCreateQuotaRequest([0]string{}, elemIsEscaped, w, r)
											default:
												s.notAllowed(w, r, notAllowedParams{
													allowedMethods: "GET",
													allowedHeaders: nil,
													acceptPost:     "",
													acceptPatch:    "",
												})
											}

											return
										}

									}

									elem = origElem
								case 'j': // Prefix: "joined_statuses"
									origElem := elem
									if l := len("joined_statuses"); len(elem) >= l && elem[0:l] == "joined_statuses" {
										elem = elem[l:]
									} else {
										break
									}

									if len(elem) == 0 {
										// Leaf node.
										switch r.Method {
										case "GET":
											s.handleGetJoinedGroupStatusesRequest([0]string{}, elemIsEscaped, w, r)
										default:
											s.notAllowed(w, r, notAllowedParams{
												allowedMethods: "GET",
												allowedHeaders: nil,
												acceptPost:     "",
												acceptPatch:    "",
											})
										}

										return
									}

									elem = origElem
								case 'u': // Prefix: "u"
									origElem := elem
									if l := len("u"); len(elem) >= l && elem[0:l] == "u" {
										elem = elem[l:]
									} else {
										break
									}

									if len(elem) == 0 {
										break
									}
									switch elem[0] {
									case 'n': // Prefix: "nread_status"

										if l := len("nread_status"); len(elem) >= l && elem[0:l] == "nread_status" {
											elem = elem[l:]
										} else {
											break
										}

										if len(elem) == 0 {
											// Leaf node.
											switch r.Method {
											case "GET":
												s.handleGetGroupUnreadStatusRequest([0]string{}, elemIsEscaped, w, r)
											default:
												s.notAllowed(w, r, notAllowedParams{
													allowedMethods: "GET",
													allowedHeaders: nil,
													acceptPost:     "",
													acceptPatch:    "",
												})
											}

											return
										}

									case 's': // Prefix: "ser_group_list"

										if l := len("ser_group_list"); len(elem) >= l && elem[0:l] == "ser_group_list" {
											elem = elem[l:]
										} else {
											break
										}

										if len(elem) == 0 {
											// Leaf node.
											switch r.Method {
											case "GET":
												s.handleGetUserGroupListRequest([0]string{}, elemIsEscaped, w, r)
											default:
												s.notAllowed(w, r, notAllowedParams{
													allowedMethods: "GET",
													allowedHeaders: nil,
													acceptPost:     "",
													acceptPatch:    "",
												})
											}

											return
										}

									}

									elem = origElem
								}
								// Param: "id"
								// Match until "/"
								idx := strings.IndexByte(elem, '/')
								if idx < 0 {
									idx = len(elem)
								}
								args[0] = elem[:idx]
								elem = elem[idx:]

								if len(elem) == 0 {
									switch r.Method {
									case "GET":
										s.handleGetGroupRequest([1]string{
											args[0],
										}, elemIsEscaped, w, r)
									default:
										s.notAllowed(w, r, notAllowedParams{
											allowedMethods: "GET",
											allowedHeaders: nil,
											acceptPost:     "",
											acceptPatch:    "",
										})
									}

									return
								}
								switch elem[0] {
								case '/': // Prefix: "/"

									if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
										elem = elem[l:]
									} else {
										break
									}

									if len(elem) == 0 {
										break
									}
									switch elem[0] {
									case 'a': // Prefix: "accept/"

										if l := len("accept/"); len(elem) >= l && elem[0:l] == "accept/" {
											elem = elem[l:]
										} else {
											break
										}

										// Param: "userId"
										// Leaf parameter, slashes are prohibited
										idx := strings.IndexByte(elem, '/')
										if idx >= 0 {
											break
										}
										args[1] = elem
										elem = ""

										if len(elem) == 0 {
											// Leaf node.
											switch r.Method {
											case "POST":
												s.handleAcceptGroupJoinRequestRequest([2]string{
													args[0],
													args[1],
												}, elemIsEscaped, w, r)
											default:
												s.notAllowed(w, r, notAllowedParams{
													allowedMethods: "POST",
													allowedHeaders: nil,
													acceptPost:     "",
													acceptPatch:    "",
												})
											}

											return
										}

									case 'b': // Prefix: "ban"

										if l := len("ban"); len(elem) >= l && elem[0:l] == "ban" {
											elem = elem[l:]
										} else {
											break
										}

										if len(elem) == 0 {
											break
										}
										switch elem[0] {
										case '/': // Prefix: "/"

											if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
												elem = elem[l:]
											} else {
												break
											}

											// Param: "userId"
											// Leaf parameter, slashes are prohibited
											idx := strings.IndexByte(elem, '/')
											if idx >= 0 {
												break
											}
											args[1] = elem
											elem = ""

											if len(elem) == 0 {
												// Leaf node.
												switch r.Method {
												case "POST":
													s.handleBanGroupUserRequest([2]string{
														args[0],
														args[1],
													}, elemIsEscaped, w, r)
												default:
													s.notAllowed(w, r, notAllowedParams{
														allowedMethods: "POST",
														allowedHeaders: nil,
														acceptPost:     "",
														acceptPatch:    "",
													})
												}

												return
											}

										case '_': // Prefix: "_list"

											if l := len("_list"); len(elem) >= l && elem[0:l] == "_list" {
												elem = elem[l:]
											} else {
												break
											}

											if len(elem) == 0 {
												// Leaf node.
												switch r.Method {
												case "GET":
													s.handleGetGroupBanListRequest([1]string{
														args[0],
													}, elemIsEscaped, w, r)
												default:
													s.notAllowed(w, r, notAllowedParams{
														allowedMethods: "GET",
														allowedHeaders: nil,
														acceptPost:     "",
														acceptPatch:    "",
													})
												}

												return
											}

										}

									case 'd': // Prefix: "de"

										if l := len("de"); len(elem) >= l && elem[0:l] == "de" {
											elem = elem[l:]
										} else {
											break
										}

										if len(elem) == 0 {
											break
										}
										switch elem[0] {
										case 'c': // Prefix: "cline/"

											if l := len("cline/"); len(elem) >= l && elem[0:l] == "cline/" {
												elem = elem[l:]
											} else {
												break
											}

											// Param: "userId"
											// Leaf parameter, slashes are prohibited
											idx := strings.IndexByte(elem, '/')
											if idx >= 0 {
												break
											}
											args[1] = elem
											elem = ""

											if len(elem) == 0 {
												// Leaf node.
												switch r.Method {
												case "POST":
													s.handleDeclineGroupJoinRequestRequest([2]string{
														args[0],
														args[1],
													}, elemIsEscaped, w, r)
												default:
													s.notAllowed(w, r, notAllowedParams{
														allowedMethods: "POST",
														allowedHeaders: nil,
														acceptPost:     "",
														acceptPatch:    "",
													})
												}

												return
											}

										case 'p': // Prefix: "putize"

											if l := len("putize"); len(elem) >= l && elem[0:l] == "putize" {
												elem = elem[l:]
											} else {
												break
											}

											if len(elem) == 0 {
												switch r.Method {
												case "DELETE":
													s.handleRemoveGroupDeputiesRequest([1]string{
														args[0],
													}, elemIsEscaped, w, r)
												case "PUT":
													s.handleDeputizeGroupUsersRequest([1]string{
														args[0],
													}, elemIsEscaped, w, r)
												default:
													s.notAllowed(w, r, notAllowedParams{
														allowedMethods: "DELETE,PUT",
														allowedHeaders: nil,
														acceptPost:     "",
														acceptPatch:    "",
													})
												}

												return
											}
											switch elem[0] {
											case '/': // Prefix: "/"

												if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
													elem = elem[l:]
												} else {
													break
												}

												// Param: "user_id"
												// Match until "/"
												idx := strings.IndexByte(elem, '/')
												if idx < 0 {
													idx = len(elem)
												}
												args[1] = elem[:idx]
												elem = elem[idx:]

												if len(elem) == 0 {
													break
												}
												switch elem[0] {
												case '/': // Prefix: "/withdraw"

													if l := len("/withdraw"); len(elem) >= l && elem[0:l] == "/withdraw" {
														elem = elem[l:]
													} else {
														break
													}

													if len(elem) == 0 {
														// Leaf node.
														switch r.Method {
														case "PUT":
															s.handleWithdrawGroupDeputyRequest([2]string{
																args[0],
																args[1],
															}, elemIsEscaped, w, r)
														default:
															s.notAllowed(w, r, notAllowedParams{
																allowedMethods: "PUT",
																allowedHeaders: nil,
																acceptPost:     "",
																acceptPatch:    "",
															})
														}

														return
													}

												}

											}

										}

									case 'f': // Prefix: "fire/"

										if l := len("fire/"); len(elem) >= l && elem[0:l] == "fire/" {
											elem = elem[l:]
										} else {
											break
										}

										// Param: "user_id"
										// Leaf parameter, slashes are prohibited
										idx := strings.IndexByte(elem, '/')
										if idx >= 0 {
											break
										}
										args[1] = elem
										elem = ""

										if len(elem) == 0 {
											// Leaf node.
											switch r.Method {
											case "POST":
												s.handleFireGroupUserRequest([2]string{
													args[0],
													args[1],
												}, elemIsEscaped, w, r)
											default:
												s.notAllowed(w, r, notAllowedParams{
													allowedMethods: "POST",
													allowedHeaders: nil,
													acceptPost:     "",
													acceptPatch:    "",
												})
											}

											return
										}

									case 'g': // Prefix: "gift_"

										if l := len("gift_"); len(elem) >= l && elem[0:l] == "gift_" {
											elem = elem[l:]
										} else {
											break
										}

										if len(elem) == 0 {
											break
										}
										switch elem[0] {
										case 'h': // Prefix: "history"

											if l := len("history"); len(elem) >= l && elem[0:l] == "history" {
												elem = elem[l:]
											} else {
												break
											}

											if len(elem) == 0 {
												// Leaf node.
												switch r.Method {
												case "GET":
													s.handleGetGroupGiftHistoryRequest([1]string{
														args[0],
													}, elemIsEscaped, w, r)
												default:
													s.notAllowed(w, r, notAllowedParams{
														allowedMethods: "GET",
														allowedHeaders: nil,
														acceptPost:     "",
														acceptPatch:    "",
													})
												}

												return
											}

										case 't': // Prefix: "transactions"

											if l := len("transactions"); len(elem) >= l && elem[0:l] == "transactions" {
												elem = elem[l:]
											} else {
												break
											}

											if len(elem) == 0 {
												// Leaf node.
												switch r.Method {
												case "GET":
													s.handleGetGroupGiftTransactionsRequest([1]string{
														args[0],
													}, elemIsEscaped, w, r)
												default:
													s.notAllowed(w, r, notAllowedParams{
														allowedMethods: "GET",
														allowedHeaders: nil,
														acceptPost:     "",
														acceptPatch:    "",
													})
												}

												return
											}

										}

									case 'h': // Prefix: "highlights"

										if l := len("highlights"); len(elem) >= l && elem[0:l] == "highlights" {
											elem = elem[l:]
										} else {
											break
										}

										if len(elem) == 0 {
											switch r.Method {
											case "GET":
												s.handleGetGroupHighlightsRequest([1]string{
													args[0],
												}, elemIsEscaped, w, r)
											default:
												s.notAllowed(w, r, notAllowedParams{
													allowedMethods: "GET",
													allowedHeaders: nil,
													acceptPost:     "",
													acceptPatch:    "",
												})
											}

											return
										}
										switch elem[0] {
										case '/': // Prefix: "/"

											if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
												elem = elem[l:]
											} else {
												break
											}

											// Param: "post_id"
											// Leaf parameter, slashes are prohibited
											idx := strings.IndexByte(elem, '/')
											if idx >= 0 {
												break
											}
											args[1] = elem
											elem = ""

											if len(elem) == 0 {
												// Leaf node.
												switch r.Method {
												case "DELETE":
													s.handleUnpinGroupHighlightPostRequest([2]string{
														args[0],
														args[1],
													}, elemIsEscaped, w, r)
												case "PUT":
													s.handlePinGroupHighlightPostRequest([2]string{
														args[0],
														args[1],
													}, elemIsEscaped, w, r)
												default:
													s.notAllowed(w, r, notAllowedParams{
														allowedMethods: "DELETE,PUT",
														allowedHeaders: nil,
														acceptPost:     "",
														acceptPatch:    "",
													})
												}

												return
											}

										}

									case 'i': // Prefix: "invite"

										if l := len("invite"); len(elem) >= l && elem[0:l] == "invite" {
											elem = elem[l:]
										} else {
											break
										}

										if len(elem) == 0 {
											// Leaf node.
											switch r.Method {
											case "POST":
												s.handleInviteToGroupRequest([1]string{
													args[0],
												}, elemIsEscaped, w, r)
											default:
												s.notAllowed(w, r, notAllowedParams{
													allowedMethods: "POST",
													allowedHeaders: rn227AllowedHeaders,
													acceptPost:     "multipart/form-data",
													acceptPatch:    "",
												})
											}

											return
										}

									case 'j': // Prefix: "join"

										if l := len("join"); len(elem) >= l && elem[0:l] == "join" {
											elem = elem[l:]
										} else {
											break
										}

										if len(elem) == 0 {
											// Leaf node.
											switch r.Method {
											case "POST":
												s.handleJoinGroupRequest([1]string{
													args[0],
												}, elemIsEscaped, w, r)
											default:
												s.notAllowed(w, r, notAllowedParams{
													allowedMethods: "POST",
													allowedHeaders: nil,
													acceptPost:     "",
													acceptPatch:    "",
												})
											}

											return
										}

									case 'l': // Prefix: "leave"

										if l := len("leave"); len(elem) >= l && elem[0:l] == "leave" {
											elem = elem[l:]
										} else {
											break
										}

										if len(elem) == 0 {
											// Leaf node.
											switch r.Method {
											case "DELETE":
												s.handleLeaveGroupRequest([1]string{
													args[0],
												}, elemIsEscaped, w, r)
											default:
												s.notAllowed(w, r, notAllowedParams{
													allowedMethods: "DELETE",
													allowedHeaders: nil,
													acceptPost:     "",
													acceptPatch:    "",
												})
											}

											return
										}

									case 'm': // Prefix: "members/"

										if l := len("members/"); len(elem) >= l && elem[0:l] == "members/" {
											elem = elem[l:]
										} else {
											break
										}

										// Param: "userId"
										// Leaf parameter, slashes are prohibited
										idx := strings.IndexByte(elem, '/')
										if idx >= 0 {
											break
										}
										args[1] = elem
										elem = ""

										if len(elem) == 0 {
											// Leaf node.
											switch r.Method {
											case "GET":
												s.handleGetGroupMemberRequest([2]string{
													args[0],
													args[1],
												}, elemIsEscaped, w, r)
											default:
												s.notAllowed(w, r, notAllowedParams{
													allowedMethods: "GET",
													allowedHeaders: nil,
													acceptPost:     "",
													acceptPatch:    "",
												})
											}

											return
										}

									case 'r': // Prefix: "re"

										if l := len("re"); len(elem) >= l && elem[0:l] == "re" {
											elem = elem[l:]
										} else {
											break
										}

										if len(elem) == 0 {
											break
										}
										switch elem[0] {
										case 'c': // Prefix: "ceived_gifts/"

											if l := len("ceived_gifts/"); len(elem) >= l && elem[0:l] == "ceived_gifts/" {
												elem = elem[l:]
											} else {
												break
											}

											// Param: "gift_id"
											// Match until "/"
											idx := strings.IndexByte(elem, '/')
											if idx < 0 {
												idx = len(elem)
											}
											args[1] = elem[:idx]
											elem = elem[idx:]

											if len(elem) == 0 {
												break
											}
											switch elem[0] {
											case '/': // Prefix: "/senders"

												if l := len("/senders"); len(elem) >= l && elem[0:l] == "/senders" {
													elem = elem[l:]
												} else {
													break
												}

												if len(elem) == 0 {
													// Leaf node.
													switch r.Method {
													case "GET":
														s.handleGetGroupReceivedGiftSendersRequest([2]string{
															args[0],
															args[1],
														}, elemIsEscaped, w, r)
													default:
														s.notAllowed(w, r, notAllowedParams{
															allowedMethods: "GET",
															allowedHeaders: nil,
															acceptPost:     "",
															acceptPatch:    "",
														})
													}

													return
												}

											}

										case 'l': // Prefix: "lat"

											if l := len("lat"); len(elem) >= l && elem[0:l] == "lat" {
												elem = elem[l:]
											} else {
												break
											}

											if len(elem) == 0 {
												break
											}
											switch elem[0] {
											case 'a': // Prefix: "able"

												if l := len("able"); len(elem) >= l && elem[0:l] == "able" {
													elem = elem[l:]
												} else {
													break
												}

												if len(elem) == 0 {
													// Leaf node.
													switch r.Method {
													case "GET":
														s.handleGetRelatableGroupsRequest([1]string{
															args[0],
														}, elemIsEscaped, w, r)
													default:
														s.notAllowed(w, r, notAllowedParams{
															allowedMethods: "GET",
															allowedHeaders: nil,
															acceptPost:     "",
															acceptPatch:    "",
														})
													}

													return
												}

											case 'e': // Prefix: "ed"

												if l := len("ed"); len(elem) >= l && elem[0:l] == "ed" {
													elem = elem[l:]
												} else {
													break
												}

												if len(elem) == 0 {
													// Leaf node.
													switch r.Method {
													case "DELETE":
														s.handleRemoveRelatedGroupsRequest([1]string{
															args[0],
														}, elemIsEscaped, w, r)
													case "GET":
														s.handleGetRelatedGroupsRequest([1]string{
															args[0],
														}, elemIsEscaped, w, r)
													case "PUT":
														s.handleUpdateRelatedGroupsRequest([1]string{
															args[0],
														}, elemIsEscaped, w, r)
													default:
														s.notAllowed(w, r, notAllowedParams{
															allowedMethods: "DELETE,GET,PUT",
															allowedHeaders: nil,
															acceptPost:     "",
															acceptPatch:    "",
														})
													}

													return
												}

											}

										case 'q': // Prefix: "quest_walkthrough"

											if l := len("quest_walkthrough"); len(elem) >= l && elem[0:l] == "quest_walkthrough" {
												elem = elem[l:]
											} else {
												break
											}

											if len(elem) == 0 {
												// Leaf node.
												switch r.Method {
												case "POST":
													s.handleRequestGroupWalkthroughRequest([1]string{
														args[0],
													}, elemIsEscaped, w, r)
												default:
													s.notAllowed(w, r, notAllowedParams{
														allowedMethods: "POST",
														allowedHeaders: nil,
														acceptPost:     "",
														acceptPatch:    "",
													})
												}

												return
											}

										}

									case 's': // Prefix: "set_title"

										if l := len("set_title"); len(elem) >= l && elem[0:l] == "set_title" {
											elem = elem[l:]
										} else {
											break
										}

										if len(elem) == 0 {
											// Leaf node.
											switch r.Method {
											case "POST":
												s.handleSetGroupTitleRequest([1]string{
													args[0],
												}, elemIsEscaped, w, r)
											default:
												s.notAllowed(w, r, notAllowedParams{
													allowedMethods: "POST",
													allowedHeaders: rn317AllowedHeaders,
													acceptPost:     "multipart/form-data",
													acceptPatch:    "",
												})
											}

											return
										}

									case 't': // Prefix: "t"

										if l := len("t"); len(elem) >= l && elem[0:l] == "t" {
											elem = elem[l:]
										} else {
											break
										}

										if len(elem) == 0 {
											break
										}
										switch elem[0] {
										case 'a': // Prefix: "ake_over"

											if l := len("ake_over"); len(elem) >= l && elem[0:l] == "ake_over" {
												elem = elem[l:]
											} else {
												break
											}

											if len(elem) == 0 {
												// Leaf node.
												switch r.Method {
												case "POST":
													s.handleTakeOverGroupRequest([1]string{
														args[0],
													}, elemIsEscaped, w, r)
												default:
													s.notAllowed(w, r, notAllowedParams{
														allowedMethods: "POST",
														allowedHeaders: nil,
														acceptPost:     "",
														acceptPatch:    "",
													})
												}

												return
											}

										case 'r': // Prefix: "ransfer"

											if l := len("ransfer"); len(elem) >= l && elem[0:l] == "ransfer" {
												elem = elem[l:]
											} else {
												break
											}

											if len(elem) == 0 {
												switch r.Method {
												case "DELETE":
													s.handleCancelGroupTransferRequest([1]string{
														args[0],
													}, elemIsEscaped, w, r)
												case "PUT":
													s.handleAcceptGroupTransferRequest([1]string{
														args[0],
													}, elemIsEscaped, w, r)
												default:
													s.notAllowed(w, r, notAllowedParams{
														allowedMethods: "DELETE,PUT",
														allowedHeaders: nil,
														acceptPost:     "",
														acceptPatch:    "",
													})
												}

												return
											}
											switch elem[0] {
											case '/': // Prefix: "/withdraw"

												if l := len("/withdraw"); len(elem) >= l && elem[0:l] == "/withdraw" {
													elem = elem[l:]
												} else {
													break
												}

												if len(elem) == 0 {
													// Leaf node.
													switch r.Method {
													case "PUT":
														s.handleWithdrawGroupTransferRequest([1]string{
															args[0],
														}, elemIsEscaped, w, r)
													default:
														s.notAllowed(w, r, notAllowedParams{
															allowedMethods: "PUT",
															allowedHeaders: rn362AllowedHeaders,
															acceptPost:     "",
															acceptPatch:    "",
														})
													}

													return
												}

											}

										}

									case 'u': // Prefix: "u"

										if l := len("u"); len(elem) >= l && elem[0:l] == "u" {
											elem = elem[l:]
										} else {
											break
										}

										if len(elem) == 0 {
											break
										}
										switch elem[0] {
										case 'n': // Prefix: "nban/"

											if l := len("nban/"); len(elem) >= l && elem[0:l] == "nban/" {
												elem = elem[l:]
											} else {
												break
											}

											// Param: "userId"
											// Leaf parameter, slashes are prohibited
											idx := strings.IndexByte(elem, '/')
											if idx >= 0 {
												break
											}
											args[1] = elem
											elem = ""

											if len(elem) == 0 {
												// Leaf node.
												switch r.Method {
												case "POST":
													s.handleUnbanGroupUserRequest([2]string{
														args[0],
														args[1],
													}, elemIsEscaped, w, r)
												default:
													s.notAllowed(w, r, notAllowedParams{
														allowedMethods: "POST",
														allowedHeaders: nil,
														acceptPost:     "",
														acceptPatch:    "",
													})
												}

												return
											}

										case 's': // Prefix: "sers/invitable"

											if l := len("sers/invitable"); len(elem) >= l && elem[0:l] == "sers/invitable" {
												elem = elem[l:]
											} else {
												break
											}

											if len(elem) == 0 {
												// Leaf node.
												switch r.Method {
												case "GET":
													s.handleGetInvitableGroupUsersRequest([1]string{
														args[0],
													}, elemIsEscaped, w, r)
												default:
													s.notAllowed(w, r, notAllowedParams{
														allowedMethods: "GET",
														allowedHeaders: nil,
														acceptPost:     "",
														acceptPatch:    "",
													})
												}

												return
											}

										}

									case 'v': // Prefix: "visit"

										if l := len("visit"); len(elem) >= l && elem[0:l] == "visit" {
											elem = elem[l:]
										} else {
											break
										}

										if len(elem) == 0 {
											// Leaf node.
											switch r.Method {
											case "POST":
												s.handleVisitGroupRequest([1]string{
													args[0],
												}, elemIsEscaped, w, r)
											default:
												s.notAllowed(w, r, notAllowedParams{
													allowedMethods: "POST",
													allowedHeaders: nil,
													acceptPost:     "",
													acceptPatch:    "",
												})
											}

											return
										}

									}

								}

							}

						}

					case 'h': // Prefix: "hidden/"

						if l := len("hidden/"); len(elem) >= l && elem[0:l] == "hidden/" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							break
						}
						switch elem[0] {
						case 'c': // Prefix: "chats"

							if l := len("chats"); len(elem) >= l && elem[0:l] == "chats" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch r.Method {
								case "DELETE":
									s.handleUnhideChatsRequest([0]string{}, elemIsEscaped, w, r)
								case "GET":
									s.handleListHiddenChatsRequest([0]string{}, elemIsEscaped, w, r)
								case "POST":
									s.handleHideChatsRequest([0]string{}, elemIsEscaped, w, r)
								default:
									s.notAllowed(w, r, notAllowedParams{
										allowedMethods: "DELETE,GET,POST",
										allowedHeaders: rn218AllowedHeaders,
										acceptPost:     "multipart/form-data",
										acceptPatch:    "",
									})
								}

								return
							}

						case 'u': // Prefix: "users"

							if l := len("users"); len(elem) >= l && elem[0:l] == "users" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch r.Method {
								case "DELETE":
									s.handleUnhideUsersRequest([0]string{}, elemIsEscaped, w, r)
								case "GET":
									s.handleListHiddenUsersRequest([0]string{}, elemIsEscaped, w, r)
								case "POST":
									s.handleHideUsersRequest([0]string{}, elemIsEscaped, w, r)
								default:
									s.notAllowed(w, r, notAllowedParams{
										allowedMethods: "DELETE,GET,POST",
										allowedHeaders: rn219AllowedHeaders,
										acceptPost:     "multipart/form-data",
										acceptPatch:    "",
									})
								}

								return
							}

						case 'w': // Prefix: "words"

							if l := len("words"); len(elem) >= l && elem[0:l] == "words" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch r.Method {
								case "DELETE":
									s.handleDeleteMuteKeywordRequest([0]string{}, elemIsEscaped, w, r)
								case "GET":
									s.handleListMuteKeywordsRequest([0]string{}, elemIsEscaped, w, r)
								case "POST":
									s.handleCreateMuteKeywordRequest([0]string{}, elemIsEscaped, w, r)
								default:
									s.notAllowed(w, r, notAllowedParams{
										allowedMethods: "DELETE,GET,POST",
										allowedHeaders: rn39AllowedHeaders,
										acceptPost:     "application/json",
										acceptPatch:    "",
									})
								}

								return
							}

						}

					case 'p': // Prefix: "p"

						if l := len("p"); len(elem) >= l && elem[0:l] == "p" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							break
						}
						switch elem[0] {
						case 'i': // Prefix: "inned/"

							if l := len("inned/"); len(elem) >= l && elem[0:l] == "inned/" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								break
							}
							switch elem[0] {
							case 'g': // Prefix: "groups"

								if l := len("groups"); len(elem) >= l && elem[0:l] == "groups" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									switch r.Method {
									case "POST":
										s.handlePinGroupRequest([0]string{}, elemIsEscaped, w, r)
									default:
										s.notAllowed(w, r, notAllowedParams{
											allowedMethods: "POST",
											allowedHeaders: rn273AllowedHeaders,
											acceptPost:     "multipart/form-data",
											acceptPatch:    "",
										})
									}

									return
								}
								switch elem[0] {
								case '/': // Prefix: "/"

									if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
										elem = elem[l:]
									} else {
										break
									}

									// Param: "id"
									// Leaf parameter, slashes are prohibited
									idx := strings.IndexByte(elem, '/')
									if idx >= 0 {
										break
									}
									args[0] = elem
									elem = ""

									if len(elem) == 0 {
										// Leaf node.
										switch r.Method {
										case "DELETE":
											s.handleUnpinGroupRequest([1]string{
												args[0],
											}, elemIsEscaped, w, r)
										default:
											s.notAllowed(w, r, notAllowedParams{
												allowedMethods: "DELETE",
												allowedHeaders: nil,
												acceptPost:     "",
												acceptPatch:    "",
											})
										}

										return
									}

								}

							case 'p': // Prefix: "posts"

								if l := len("posts"); len(elem) >= l && elem[0:l] == "posts" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch r.Method {
									case "POST":
										s.handlePinPostRequest([0]string{}, elemIsEscaped, w, r)
									default:
										s.notAllowed(w, r, notAllowedParams{
											allowedMethods: "POST",
											allowedHeaders: rn279AllowedHeaders,
											acceptPost:     "multipart/form-data",
											acceptPatch:    "",
										})
									}

									return
								}

							case 'r': // Prefix: "reviews"

								if l := len("reviews"); len(elem) >= l && elem[0:l] == "reviews" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									switch r.Method {
									case "POST":
										s.handlePinReviewRequest([0]string{}, elemIsEscaped, w, r)
									default:
										s.notAllowed(w, r, notAllowedParams{
											allowedMethods: "POST",
											allowedHeaders: rn280AllowedHeaders,
											acceptPost:     "multipart/form-data",
											acceptPatch:    "",
										})
									}

									return
								}
								switch elem[0] {
								case '/': // Prefix: "/"

									if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
										elem = elem[l:]
									} else {
										break
									}

									// Param: "id"
									// Leaf parameter, slashes are prohibited
									idx := strings.IndexByte(elem, '/')
									if idx >= 0 {
										break
									}
									args[0] = elem
									elem = ""

									if len(elem) == 0 {
										// Leaf node.
										switch r.Method {
										case "DELETE":
											s.handleUnpinReviewRequest([1]string{
												args[0],
											}, elemIsEscaped, w, r)
										default:
											s.notAllowed(w, r, notAllowedParams{
												allowedMethods: "DELETE",
												allowedHeaders: nil,
												acceptPost:     "",
												acceptPatch:    "",
											})
										}

										return
									}

								}

							}

						case 'o': // Prefix: "osts/"

							if l := len("osts/"); len(elem) >= l && elem[0:l] == "osts/" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								break
							}
							switch elem[0] {
							case 'a': // Prefix: "active_call"
								origElem := elem
								if l := len("active_call"); len(elem) >= l && elem[0:l] == "active_call" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch r.Method {
									case "GET":
										s.handleGetActiveCallPostRequest([0]string{}, elemIsEscaped, w, r)
									default:
										s.notAllowed(w, r, notAllowedParams{
											allowedMethods: "GET",
											allowedHeaders: nil,
											acceptPost:     "",
											acceptPatch:    "",
										})
									}

									return
								}

								elem = origElem
							case 'd': // Prefix: "delete_all_post"
								origElem := elem
								if l := len("delete_all_post"); len(elem) >= l && elem[0:l] == "delete_all_post" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch r.Method {
									case "POST":
										s.handleDeleteAllPostsRequest([0]string{}, elemIsEscaped, w, r)
									default:
										s.notAllowed(w, r, notAllowedParams{
											allowedMethods: "POST",
											allowedHeaders: nil,
											acceptPost:     "",
											acceptPatch:    "",
										})
									}

									return
								}

								elem = origElem
							case 'r': // Prefix: "recommended_tag"
								origElem := elem
								if l := len("recommended_tag"); len(elem) >= l && elem[0:l] == "recommended_tag" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch r.Method {
									case "POST":
										s.handleGetRecommendedPostTagsRequest([0]string{}, elemIsEscaped, w, r)
									default:
										s.notAllowed(w, r, notAllowedParams{
											allowedMethods: "POST",
											allowedHeaders: rn177AllowedHeaders,
											acceptPost:     "multipart/form-data",
											acceptPatch:    "",
										})
									}

									return
								}

								elem = origElem
							case 'v': // Prefix: "v"
								origElem := elem
								if l := len("v"); len(elem) >= l && elem[0:l] == "v" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									break
								}
								switch elem[0] {
								case 'a': // Prefix: "alidate"

									if l := len("alidate"); len(elem) >= l && elem[0:l] == "alidate" {
										elem = elem[l:]
									} else {
										break
									}

									if len(elem) == 0 {
										// Leaf node.
										switch r.Method {
										case "POST":
											s.handleValidatePostRequest([0]string{}, elemIsEscaped, w, r)
										default:
											s.notAllowed(w, r, notAllowedParams{
												allowedMethods: "POST",
												allowedHeaders: rn349AllowedHeaders,
												acceptPost:     "multipart/form-data",
												acceptPatch:    "",
											})
										}

										return
									}

								case 'i': // Prefix: "ideos/"

									if l := len("ideos/"); len(elem) >= l && elem[0:l] == "ideos/" {
										elem = elem[l:]
									} else {
										break
									}

									// Param: "videoId"
									// Match until "/"
									idx := strings.IndexByte(elem, '/')
									if idx < 0 {
										idx = len(elem)
									}
									args[0] = elem[:idx]
									elem = elem[idx:]

									if len(elem) == 0 {
										break
									}
									switch elem[0] {
									case '/': // Prefix: "/view"

										if l := len("/view"); len(elem) >= l && elem[0:l] == "/view" {
											elem = elem[l:]
										} else {
											break
										}

										if len(elem) == 0 {
											// Leaf node.
											switch r.Method {
											case "POST":
												s.handleViewPostVideoRequest([1]string{
													args[0],
												}, elemIsEscaped, w, r)
											default:
												s.notAllowed(w, r, notAllowedParams{
													allowedMethods: "POST",
													allowedHeaders: nil,
													acceptPost:     "",
													acceptPatch:    "",
												})
											}

											return
										}

									}

								}

								elem = origElem
							}
							// Param: "post_id"
							// Match until "/"
							idx := strings.IndexByte(elem, '/')
							if idx < 0 {
								idx = len(elem)
							}
							args[0] = elem[:idx]
							elem = elem[idx:]

							if len(elem) == 0 {
								break
							}
							switch elem[0] {
							case '/': // Prefix: "/"

								if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									break
								}
								switch elem[0] {
								case 'g': // Prefix: "gift_transactions"

									if l := len("gift_transactions"); len(elem) >= l && elem[0:l] == "gift_transactions" {
										elem = elem[l:]
									} else {
										break
									}

									if len(elem) == 0 {
										// Leaf node.
										switch r.Method {
										case "GET":
											s.handleGetPostGiftTransactionsRequest([1]string{
												args[0],
											}, elemIsEscaped, w, r)
										default:
											s.notAllowed(w, r, notAllowedParams{
												allowedMethods: "GET",
												allowedHeaders: nil,
												acceptPost:     "",
												acceptPatch:    "",
											})
										}

										return
									}

								case 'l': // Prefix: "likers"

									if l := len("likers"); len(elem) >= l && elem[0:l] == "likers" {
										elem = elem[l:]
									} else {
										break
									}

									if len(elem) == 0 {
										// Leaf node.
										switch r.Method {
										case "GET":
											s.handleGetPostLikersRequest([1]string{
												args[0],
											}, elemIsEscaped, w, r)
										default:
											s.notAllowed(w, r, notAllowedParams{
												allowedMethods: "GET",
												allowedHeaders: nil,
												acceptPost:     "",
												acceptPatch:    "",
											})
										}

										return
									}

								case 'u': // Prefix: "unlike"

									if l := len("unlike"); len(elem) >= l && elem[0:l] == "unlike" {
										elem = elem[l:]
									} else {
										break
									}

									if len(elem) == 0 {
										// Leaf node.
										switch r.Method {
										case "POST":
											s.handleUnlikePostRequest([1]string{
												args[0],
											}, elemIsEscaped, w, r)
										default:
											s.notAllowed(w, r, notAllowedParams{
												allowedMethods: "POST",
												allowedHeaders: nil,
												acceptPost:     "",
												acceptPatch:    "",
											})
										}

										return
									}

								}

							}

						}

					case 'r': // Prefix: "received_gifts"

						if l := len("received_gifts"); len(elem) >= l && elem[0:l] == "received_gifts" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							switch r.Method {
							case "GET":
								s.handleListReceivedGiftsV1Request([0]string{}, elemIsEscaped, w, r)
							default:
								s.notAllowed(w, r, notAllowedParams{
									allowedMethods: "GET",
									allowedHeaders: nil,
									acceptPost:     "",
									acceptPatch:    "",
								})
							}

							return
						}
						switch elem[0] {
						case '/': // Prefix: "/"

							if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
								elem = elem[l:]
							} else {
								break
							}

							// Param: "gift_id"
							// Match until "/"
							idx := strings.IndexByte(elem, '/')
							if idx < 0 {
								idx = len(elem)
							}
							args[0] = elem[:idx]
							elem = elem[idx:]

							if len(elem) == 0 {
								switch r.Method {
								case "GET":
									s.handleGetReceivedGiftTransactionRequest([1]string{
										args[0],
									}, elemIsEscaped, w, r)
								default:
									s.notAllowed(w, r, notAllowedParams{
										allowedMethods: "GET",
										allowedHeaders: nil,
										acceptPost:     "",
										acceptPatch:    "",
									})
								}

								return
							}
							switch elem[0] {
							case '/': // Prefix: "/senders"

								if l := len("/senders"); len(elem) >= l && elem[0:l] == "/senders" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch r.Method {
									case "GET":
										s.handleGetReceivedGiftSendersRequest([1]string{
											args[0],
										}, elemIsEscaped, w, r)
									default:
										s.notAllowed(w, r, notAllowedParams{
											allowedMethods: "GET",
											allowedHeaders: nil,
											acceptPost:     "",
											acceptPatch:    "",
										})
									}

									return
								}

							}

						}

					case 't': // Prefix: "threads"

						if l := len("threads"); len(elem) >= l && elem[0:l] == "threads" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							switch r.Method {
							case "GET":
								s.handleListThreadsRequest([0]string{}, elemIsEscaped, w, r)
							case "POST":
								s.handleCreateThreadRequest([0]string{}, elemIsEscaped, w, r)
							default:
								s.notAllowed(w, r, notAllowedParams{
									allowedMethods: "GET,POST",
									allowedHeaders: rn44AllowedHeaders,
									acceptPost:     "application/json",
									acceptPatch:    "",
								})
							}

							return
						}
						switch elem[0] {
						case '/': // Prefix: "/"

							if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								break
							}
							switch elem[0] {
							case 'j': // Prefix: "joined_statuses"
								origElem := elem
								if l := len("joined_statuses"); len(elem) >= l && elem[0:l] == "joined_statuses" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch r.Method {
									case "GET":
										s.handleGetJoinedThreadStatusesRequest([0]string{}, elemIsEscaped, w, r)
									default:
										s.notAllowed(w, r, notAllowedParams{
											allowedMethods: "GET",
											allowedHeaders: nil,
											acceptPost:     "",
											acceptPatch:    "",
										})
									}

									return
								}

								elem = origElem
							}
							// Param: "thread_id"
							// Match until "/"
							idx := strings.IndexByte(elem, '/')
							if idx < 0 {
								idx = len(elem)
							}
							args[0] = elem[:idx]
							elem = elem[idx:]

							if len(elem) == 0 {
								switch r.Method {
								case "DELETE":
									s.handleDeleteThreadRequest([1]string{
										args[0],
									}, elemIsEscaped, w, r)
								case "GET":
									s.handleGetThreadRequest([1]string{
										args[0],
									}, elemIsEscaped, w, r)
								case "PUT":
									s.handleUpdateThreadRequest([1]string{
										args[0],
									}, elemIsEscaped, w, r)
								default:
									s.notAllowed(w, r, notAllowedParams{
										allowedMethods: "DELETE,GET,PUT",
										allowedHeaders: rn10AllowedHeaders,
										acceptPost:     "",
										acceptPatch:    "",
									})
								}

								return
							}
							switch elem[0] {
							case '/': // Prefix: "/"

								if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									break
								}
								switch elem[0] {
								case 'm': // Prefix: "members/"

									if l := len("members/"); len(elem) >= l && elem[0:l] == "members/" {
										elem = elem[l:]
									} else {
										break
									}

									// Param: "id"
									// Leaf parameter, slashes are prohibited
									idx := strings.IndexByte(elem, '/')
									if idx >= 0 {
										break
									}
									args[1] = elem
									elem = ""

									if len(elem) == 0 {
										// Leaf node.
										switch r.Method {
										case "DELETE":
											s.handleRemoveThreadMemberRequest([2]string{
												args[0],
												args[1],
											}, elemIsEscaped, w, r)
										case "POST":
											s.handleAddThreadMemberRequest([2]string{
												args[0],
												args[1],
											}, elemIsEscaped, w, r)
										default:
											s.notAllowed(w, r, notAllowedParams{
												allowedMethods: "DELETE,POST",
												allowedHeaders: nil,
												acceptPost:     "",
												acceptPatch:    "",
											})
										}

										return
									}

								case 'p': // Prefix: "posts"

									if l := len("posts"); len(elem) >= l && elem[0:l] == "posts" {
										elem = elem[l:]
									} else {
										break
									}

									if len(elem) == 0 {
										// Leaf node.
										switch r.Method {
										case "GET":
											s.handleGetThreadPostsRequest([1]string{
												args[0],
											}, elemIsEscaped, w, r)
										case "POST":
											s.handleCreateThreadPostRequest([1]string{
												args[0],
											}, elemIsEscaped, w, r)
										default:
											s.notAllowed(w, r, notAllowedParams{
												allowedMethods: "GET,POST",
												allowedHeaders: rn46AllowedHeaders,
												acceptPost:     "multipart/form-data",
												acceptPatch:    "",
											})
										}

										return
									}

								}

							}

						}

					case 'u': // Prefix: "users/"

						if l := len("users/"); len(elem) >= l && elem[0:l] == "users/" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							break
						}
						switch elem[0] {
						case 'a': // Prefix: "a"
							origElem := elem
							if l := len("a"); len(elem) >= l && elem[0:l] == "a" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								break
							}
							switch elem[0] {
							case 'c': // Prefix: "ctive_followings"

								if l := len("ctive_followings"); len(elem) >= l && elem[0:l] == "ctive_followings" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch r.Method {
									case "GET":
										s.handleGetActiveFollowingsRequest([0]string{}, elemIsEscaped, w, r)
									default:
										s.notAllowed(w, r, notAllowedParams{
											allowedMethods: "GET",
											allowedHeaders: nil,
											acceptPost:     "",
											acceptPatch:    "",
										})
									}

									return
								}

							case 'd': // Prefix: "dditonal_notification_setting"

								if l := len("dditonal_notification_setting"); len(elem) >= l && elem[0:l] == "dditonal_notification_setting" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch r.Method {
									case "GET":
										s.handleGetAdditionalNotificationSettingRequest([0]string{}, elemIsEscaped, w, r)
									case "POST":
										s.handleUpdateAdditionalNotificationSettingRequest([0]string{}, elemIsEscaped, w, r)
									default:
										s.notAllowed(w, r, notAllowedParams{
											allowedMethods: "GET,POST",
											allowedHeaders: rn82AllowedHeaders,
											acceptPost:     "multipart/form-data",
											acceptPatch:    "",
										})
									}

									return
								}

							case 'l': // Prefix: "live"

								if l := len("live"); len(elem) >= l && elem[0:l] == "live" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch r.Method {
									case "POST":
										s.handlePingAliveRequest([0]string{}, elemIsEscaped, w, r)
									default:
										s.notAllowed(w, r, notAllowedParams{
											allowedMethods: "POST",
											allowedHeaders: nil,
											acceptPost:     "",
											acceptPatch:    "",
										})
									}

									return
								}

							}

							elem = origElem
						case 'b': // Prefix: "block_ids"
							origElem := elem
							if l := len("block_ids"); len(elem) >= l && elem[0:l] == "block_ids" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch r.Method {
								case "GET":
									s.handleGetBlockedUserIdsRequest([0]string{}, elemIsEscaped, w, r)
								default:
									s.notAllowed(w, r, notAllowedParams{
										allowedMethods: "GET",
										allowedHeaders: nil,
										acceptPost:     "",
										acceptPatch:    "",
									})
								}

								return
							}

							elem = origElem
						case 'c': // Prefix: "c"
							origElem := elem
							if l := len("c"); len(elem) >= l && elem[0:l] == "c" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								break
							}
							switch elem[0] {
							case 'h': // Prefix: "hange_"

								if l := len("hange_"); len(elem) >= l && elem[0:l] == "hange_" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									break
								}
								switch elem[0] {
								case 'e': // Prefix: "email"

									if l := len("email"); len(elem) >= l && elem[0:l] == "email" {
										elem = elem[l:]
									} else {
										break
									}

									if len(elem) == 0 {
										// Leaf node.
										switch r.Method {
										case "PUT":
											s.handleChangeEmailRequest([0]string{}, elemIsEscaped, w, r)
										default:
											s.notAllowed(w, r, notAllowedParams{
												allowedMethods: "PUT",
												allowedHeaders: rn26AllowedHeaders,
												acceptPost:     "",
												acceptPatch:    "",
											})
										}

										return
									}

								case 'p': // Prefix: "password"

									if l := len("password"); len(elem) >= l && elem[0:l] == "password" {
										elem = elem[l:]
									} else {
										break
									}

									if len(elem) == 0 {
										// Leaf node.
										switch r.Method {
										case "PUT":
											s.handleChangePasswordRequest([0]string{}, elemIsEscaped, w, r)
										default:
											s.notAllowed(w, r, notAllowedParams{
												allowedMethods: "PUT",
												allowedHeaders: rn28AllowedHeaders,
												acceptPost:     "",
												acceptPatch:    "",
											})
										}

										return
									}

								}

							case 'u': // Prefix: "ustom_definitions"

								if l := len("ustom_definitions"); len(elem) >= l && elem[0:l] == "ustom_definitions" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch r.Method {
									case "GET":
										s.handleGetUserCustomDefinitionsRequest([0]string{}, elemIsEscaped, w, r)
									default:
										s.notAllowed(w, r, notAllowedParams{
											allowedMethods: "GET",
											allowedHeaders: nil,
											acceptPost:     "",
											acceptPatch:    "",
										})
									}

									return
								}

							}

							elem = origElem
						case 'd': // Prefix: "default_settings"
							origElem := elem
							if l := len("default_settings"); len(elem) >= l && elem[0:l] == "default_settings" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch r.Method {
								case "GET":
									s.handleGetDefaultSettingsRequest([0]string{}, elemIsEscaped, w, r)
								default:
									s.notAllowed(w, r, notAllowedParams{
										allowedMethods: "GET",
										allowedHeaders: nil,
										acceptPost:     "",
										acceptPatch:    "",
									})
								}

								return
							}

							elem = origElem
						case 'f': // Prefix: "following"
							origElem := elem
							if l := len("following"); len(elem) >= l && elem[0:l] == "following" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								break
							}
							switch elem[0] {
							case '_': // Prefix: "_born_today"

								if l := len("_born_today"); len(elem) >= l && elem[0:l] == "_born_today" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch r.Method {
									case "GET":
										s.handleGetFollowingsBornTodayRequest([0]string{}, elemIsEscaped, w, r)
									default:
										s.notAllowed(w, r, notAllowedParams{
											allowedMethods: "GET",
											allowedHeaders: nil,
											acceptPost:     "",
											acceptPatch:    "",
										})
									}

									return
								}

							case 's': // Prefix: "s/chatable"

								if l := len("s/chatable"); len(elem) >= l && elem[0:l] == "s/chatable" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch r.Method {
									case "POST":
										s.handleGetChatableFollowingsRequest([0]string{}, elemIsEscaped, w, r)
									default:
										s.notAllowed(w, r, notAllowedParams{
											allowedMethods: "POST",
											allowedHeaders: rn108AllowedHeaders,
											acceptPost:     "application/json",
											acceptPatch:    "",
										})
									}

									return
								}

							}

							elem = origElem
						case 'h': // Prefix: "hima"
							origElem := elem
							if l := len("hima"); len(elem) >= l && elem[0:l] == "hima" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch r.Method {
								case "POST":
									s.handleSetHimaRequest([0]string{}, elemIsEscaped, w, r)
								default:
									s.notAllowed(w, r, notAllowedParams{
										allowedMethods: "POST",
										allowedHeaders: nil,
										acceptPost:     "",
										acceptPatch:    "",
									})
								}

								return
							}

							elem = origElem
						case 'i': // Prefix: "interests"
							origElem := elem
							if l := len("interests"); len(elem) >= l && elem[0:l] == "interests" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch r.Method {
								case "GET":
									s.handleGetUserInterestsRequest([0]string{}, elemIsEscaped, w, r)
								case "PUT":
									s.handleUpdateUserInterestsRequest([0]string{}, elemIsEscaped, w, r)
								default:
									s.notAllowed(w, r, notAllowedParams{
										allowedMethods: "GET,PUT",
										allowedHeaders: rn207AllowedHeaders,
										acceptPost:     "",
										acceptPatch:    "",
									})
								}

								return
							}

							elem = origElem
						case 'l': // Prefix: "l"
							origElem := elem
							if l := len("l"); len(elem) >= l && elem[0:l] == "l" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								break
							}
							switch elem[0] {
							case 'a': // Prefix: "anguage"

								if l := len("anguage"); len(elem) >= l && elem[0:l] == "anguage" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch r.Method {
									case "PUT":
										s.handleUpdateLanguageRequest([0]string{}, elemIsEscaped, w, r)
									default:
										s.notAllowed(w, r, notAllowedParams{
											allowedMethods: "PUT",
											allowedHeaders: rn345AllowedHeaders,
											acceptPost:     "",
											acceptPatch:    "",
										})
									}

									return
								}

							case 'i': // Prefix: "ist_id"

								if l := len("ist_id"); len(elem) >= l && elem[0:l] == "ist_id" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch r.Method {
									case "GET":
										s.handleGetUsersByIdsRequest([0]string{}, elemIsEscaped, w, r)
									default:
										s.notAllowed(w, r, notAllowedParams{
											allowedMethods: "GET",
											allowedHeaders: rn215AllowedHeaders,
											acceptPost:     "",
											acceptPatch:    "",
										})
									}

									return
								}

							case 'o': // Prefix: "ogout"

								if l := len("ogout"); len(elem) >= l && elem[0:l] == "ogout" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch r.Method {
									case "POST":
										s.handleLogoutRequest([0]string{}, elemIsEscaped, w, r)
									default:
										s.notAllowed(w, r, notAllowedParams{
											allowedMethods: "POST",
											allowedHeaders: rn259AllowedHeaders,
											acceptPost:     "multipart/form-data",
											acceptPatch:    "",
										})
									}

									return
								}

							}

							elem = origElem
						case 'p': // Prefix: "p"
							origElem := elem
							if l := len("p"); len(elem) >= l && elem[0:l] == "p" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								break
							}
							switch elem[0] {
							case 'o': // Prefix: "olicy_agreements"

								if l := len("olicy_agreements"); len(elem) >= l && elem[0:l] == "olicy_agreements" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									switch r.Method {
									case "GET":
										s.handleGetPolicyAgreementsRequest([0]string{}, elemIsEscaped, w, r)
									default:
										s.notAllowed(w, r, notAllowedParams{
											allowedMethods: "GET",
											allowedHeaders: nil,
											acceptPost:     "",
											acceptPatch:    "",
										})
									}

									return
								}
								switch elem[0] {
								case '/': // Prefix: "/"

									if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
										elem = elem[l:]
									} else {
										break
									}

									// Param: "type"
									// Leaf parameter, slashes are prohibited
									idx := strings.IndexByte(elem, '/')
									if idx >= 0 {
										break
									}
									args[0] = elem
									elem = ""

									if len(elem) == 0 {
										// Leaf node.
										switch r.Method {
										case "POST":
											s.handleAgreePolicyRequest([1]string{
												args[0],
											}, elemIsEscaped, w, r)
										default:
											s.notAllowed(w, r, notAllowedParams{
												allowedMethods: "POST",
												allowedHeaders: nil,
												acceptPost:     "",
												acceptPatch:    "",
											})
										}

										return
									}

								}

							case 'r': // Prefix: "resigned_url"

								if l := len("resigned_url"); len(elem) >= l && elem[0:l] == "resigned_url" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch r.Method {
									case "GET":
										s.handleGetUserPresignedUrlRequest([0]string{}, elemIsEscaped, w, r)
									default:
										s.notAllowed(w, r, notAllowedParams{
											allowedMethods: "GET",
											allowedHeaders: nil,
											acceptPost:     "",
											acceptPatch:    "",
										})
									}

									return
								}

							}

							elem = origElem
						case 'q': // Prefix: "qr_codes/"
							origElem := elem
							if l := len("qr_codes/"); len(elem) >= l && elem[0:l] == "qr_codes/" {
								elem = elem[l:]
							} else {
								break
							}

							// Param: "qr"
							// Leaf parameter, slashes are prohibited
							idx := strings.IndexByte(elem, '/')
							if idx >= 0 {
								break
							}
							args[0] = elem
							elem = ""

							if len(elem) == 0 {
								// Leaf node.
								switch r.Method {
								case "GET":
									s.handleGetUserByQrRequest([1]string{
										args[0],
									}, elemIsEscaped, w, r)
								default:
									s.notAllowed(w, r, notAllowedParams{
										allowedMethods: "GET",
										allowedHeaders: nil,
										acceptPost:     "",
										acceptPatch:    "",
									})
								}

								return
							}

							elem = origElem
						case 'r': // Prefix: "re"
							origElem := elem
							if l := len("re"); len(elem) >= l && elem[0:l] == "re" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								break
							}
							switch elem[0] {
							case 's': // Prefix: "set_"

								if l := len("set_"); len(elem) >= l && elem[0:l] == "set_" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									break
								}
								switch elem[0] {
								case 'c': // Prefix: "counters"

									if l := len("counters"); len(elem) >= l && elem[0:l] == "counters" {
										elem = elem[l:]
									} else {
										break
									}

									if len(elem) == 0 {
										// Leaf node.
										switch r.Method {
										case "GET":
											s.handleGetResetCountersRequest([0]string{}, elemIsEscaped, w, r)
										case "POST":
											s.handleResetCountersRequest([0]string{}, elemIsEscaped, w, r)
										default:
											s.notAllowed(w, r, notAllowedParams{
												allowedMethods: "GET,POST",
												allowedHeaders: rn185AllowedHeaders,
												acceptPost:     "multipart/form-data",
												acceptPatch:    "",
											})
										}

										return
									}

								case 'p': // Prefix: "password"

									if l := len("password"); len(elem) >= l && elem[0:l] == "password" {
										elem = elem[l:]
									} else {
										break
									}

									if len(elem) == 0 {
										// Leaf node.
										switch r.Method {
										case "PUT":
											s.handleResetPasswordRequest([0]string{}, elemIsEscaped, w, r)
										default:
											s.notAllowed(w, r, notAllowedParams{
												allowedMethods: "PUT",
												allowedHeaders: rn309AllowedHeaders,
												acceptPost:     "",
												acceptPatch:    "",
											})
										}

										return
									}

								}

							case 'v': // Prefix: "views"

								if l := len("views"); len(elem) >= l && elem[0:l] == "views" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									switch r.Method {
									case "DELETE":
										s.handleDeleteMyReviewsRequest([0]string{}, elemIsEscaped, w, r)
									case "POST":
										s.handleCreateReviewRequest([0]string{}, elemIsEscaped, w, r)
									default:
										s.notAllowed(w, r, notAllowedParams{
											allowedMethods: "DELETE,POST",
											allowedHeaders: rn41AllowedHeaders,
											acceptPost:     "multipart/form-data",
											acceptPatch:    "",
										})
									}

									return
								}
								switch elem[0] {
								case '/': // Prefix: "/mine"

									if l := len("/mine"); len(elem) >= l && elem[0:l] == "/mine" {
										elem = elem[l:]
									} else {
										break
									}

									if len(elem) == 0 {
										// Leaf node.
										switch r.Method {
										case "GET":
											s.handleGetMyReviewsRequest([0]string{}, elemIsEscaped, w, r)
										default:
											s.notAllowed(w, r, notAllowedParams{
												allowedMethods: "GET",
												allowedHeaders: nil,
												acceptPost:     "",
												acceptPatch:    "",
											})
										}

										return
									}

								}

							}

							elem = origElem
						case 's': // Prefix: "se"
							origElem := elem
							if l := len("se"); len(elem) >= l && elem[0:l] == "se" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								break
							}
							switch elem[0] {
							case 'a': // Prefix: "arch"

								if l := len("arch"); len(elem) >= l && elem[0:l] == "arch" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch r.Method {
									case "GET":
										s.handleSearchUsersRequest([0]string{}, elemIsEscaped, w, r)
									default:
										s.notAllowed(w, r, notAllowedParams{
											allowedMethods: "GET",
											allowedHeaders: nil,
											acceptPost:     "",
											acceptPatch:    "",
										})
									}

									return
								}

							case 'c': // Prefix: "cret"

								if l := len("cret"); len(elem) >= l && elem[0:l] == "cret" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									switch r.Method {
									case "GET":
										s.handleGetTwoFactorAuthRequestInfoRequest([0]string{}, elemIsEscaped, w, r)
									default:
										s.notAllowed(w, r, notAllowedParams{
											allowedMethods: "GET",
											allowedHeaders: nil,
											acceptPost:     "",
											acceptPatch:    "",
										})
									}

									return
								}
								switch elem[0] {
								case '/': // Prefix: "/"

									if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
										elem = elem[l:]
									} else {
										break
									}

									if len(elem) == 0 {
										break
									}
									switch elem[0] {
									case 'd': // Prefix: "disable"

										if l := len("disable"); len(elem) >= l && elem[0:l] == "disable" {
											elem = elem[l:]
										} else {
											break
										}

										if len(elem) == 0 {
											// Leaf node.
											switch r.Method {
											case "PUT":
												s.handleDisableTwoFactorAuthRequest([0]string{}, elemIsEscaped, w, r)
											default:
												s.notAllowed(w, r, notAllowedParams{
													allowedMethods: "PUT",
													allowedHeaders: rn67AllowedHeaders,
													acceptPost:     "",
													acceptPatch:    "",
												})
											}

											return
										}

									case 'e': // Prefix: "enable"

										if l := len("enable"); len(elem) >= l && elem[0:l] == "enable" {
											elem = elem[l:]
										} else {
											break
										}

										if len(elem) == 0 {
											// Leaf node.
											switch r.Method {
											case "PUT":
												s.handleEnableTwoFactorAuthRequest([0]string{}, elemIsEscaped, w, r)
											default:
												s.notAllowed(w, r, notAllowedParams{
													allowedMethods: "PUT",
													allowedHeaders: rn71AllowedHeaders,
													acceptPost:     "",
													acceptPatch:    "",
												})
											}

											return
										}

									case 's': // Prefix: "status"

										if l := len("status"); len(elem) >= l && elem[0:l] == "status" {
											elem = elem[l:]
										} else {
											break
										}

										if len(elem) == 0 {
											// Leaf node.
											switch r.Method {
											case "GET":
												s.handleGetTwoFactorAuthStatusRequest([0]string{}, elemIsEscaped, w, r)
											default:
												s.notAllowed(w, r, notAllowedParams{
													allowedMethods: "GET",
													allowedHeaders: nil,
													acceptPost:     "",
													acceptPatch:    "",
												})
											}

											return
										}

									}

								}

							}

							elem = origElem
						case 'w': // Prefix: "ws_token"
							origElem := elem
							if l := len("ws_token"); len(elem) >= l && elem[0:l] == "ws_token" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch r.Method {
								case "GET":
									s.handleGetWebSocketTokenRequest([0]string{}, elemIsEscaped, w, r)
								default:
									s.notAllowed(w, r, notAllowedParams{
										allowedMethods: "GET",
										allowedHeaders: nil,
										acceptPost:     "",
										acceptPatch:    "",
									})
								}

								return
							}

							elem = origElem
						}
						// Param: "id"
						// Match until "/"
						idx := strings.IndexByte(elem, '/')
						if idx < 0 {
							idx = len(elem)
						}
						args[0] = elem[:idx]
						elem = elem[idx:]

						if len(elem) == 0 {
							break
						}
						switch elem[0] {
						case '/': // Prefix: "/"

							if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								break
							}
							switch elem[0] {
							case 'b': // Prefix: "b"

								if l := len("b"); len(elem) >= l && elem[0:l] == "b" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									break
								}
								switch elem[0] {
								case 'l': // Prefix: "lock"

									if l := len("lock"); len(elem) >= l && elem[0:l] == "lock" {
										elem = elem[l:]
									} else {
										break
									}

									if len(elem) == 0 {
										// Leaf node.
										switch r.Method {
										case "POST":
											s.handleBlockUserRequest([1]string{
												args[0],
											}, elemIsEscaped, w, r)
										default:
											s.notAllowed(w, r, notAllowedParams{
												allowedMethods: "POST",
												allowedHeaders: nil,
												acceptPost:     "",
												acceptPatch:    "",
											})
										}

										return
									}

								case 'o': // Prefix: "ookmarks"

									if l := len("ookmarks"); len(elem) >= l && elem[0:l] == "ookmarks" {
										elem = elem[l:]
									} else {
										break
									}

									if len(elem) == 0 {
										switch r.Method {
										case "GET":
											s.handleGetBookmarkedPostsRequest([1]string{
												args[0],
											}, elemIsEscaped, w, r)
										default:
											s.notAllowed(w, r, notAllowedParams{
												allowedMethods: "GET",
												allowedHeaders: nil,
												acceptPost:     "",
												acceptPatch:    "",
											})
										}

										return
									}
									switch elem[0] {
									case '/': // Prefix: "/"

										if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
											elem = elem[l:]
										} else {
											break
										}

										// Param: "id"
										// Leaf parameter, slashes are prohibited
										idx := strings.IndexByte(elem, '/')
										if idx >= 0 {
											break
										}
										args[1] = elem
										elem = ""

										if len(elem) == 0 {
											// Leaf node.
											switch r.Method {
											case "DELETE":
												s.handleDeleteBookmarkRequest([2]string{
													args[0],
													args[1],
												}, elemIsEscaped, w, r)
											case "PUT":
												s.handleCreateBookmarkRequest([2]string{
													args[0],
													args[1],
												}, elemIsEscaped, w, r)
											default:
												s.notAllowed(w, r, notAllowedParams{
													allowedMethods: "DELETE,PUT",
													allowedHeaders: nil,
													acceptPost:     "",
													acceptPatch:    "",
												})
											}

											return
										}

									}

								}

							case 'f': // Prefix: "follow_recommended"

								if l := len("follow_recommended"); len(elem) >= l && elem[0:l] == "follow_recommended" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch r.Method {
									case "GET":
										s.handleGetRecommendedFollowUsersRequest([1]string{
											args[0],
										}, elemIsEscaped, w, r)
									default:
										s.notAllowed(w, r, notAllowedParams{
											allowedMethods: "GET",
											allowedHeaders: nil,
											acceptPost:     "",
											acceptPatch:    "",
										})
									}

									return
								}

							case 'g': // Prefix: "gift_transactions"

								if l := len("gift_transactions"); len(elem) >= l && elem[0:l] == "gift_transactions" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch r.Method {
									case "GET":
										s.handleGetUserGiftTransactionsRequest([1]string{
											args[0],
										}, elemIsEscaped, w, r)
									default:
										s.notAllowed(w, r, notAllowedParams{
											allowedMethods: "GET",
											allowedHeaders: nil,
											acceptPost:     "",
											acceptPatch:    "",
										})
									}

									return
								}

							}

						}

					}

				case '2': // Prefix: "2/"

					if l := len("2/"); len(elem) >= l && elem[0:l] == "2/" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						break
					}
					switch elem[0] {
					case 'c': // Prefix: "c"

						if l := len("c"); len(elem) >= l && elem[0:l] == "c" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							break
						}
						switch elem[0] {
						case 'a': // Prefix: "alls/"

							if l := len("alls/"); len(elem) >= l && elem[0:l] == "alls/" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								break
							}
							switch elem[0] {
							case 'a': // Prefix: "action_signature/validate"

								if l := len("action_signature/validate"); len(elem) >= l && elem[0:l] == "action_signature/validate" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch r.Method {
									case "GET":
										s.handleValidateCallActionSignatureRequest([0]string{}, elemIsEscaped, w, r)
									default:
										s.notAllowed(w, r, notAllowedParams{
											allowedMethods: "GET",
											allowedHeaders: nil,
											acceptPost:     "",
											acceptPatch:    "",
										})
									}

									return
								}

							case 'c': // Prefix: "conferences/"

								if l := len("conferences/"); len(elem) >= l && elem[0:l] == "conferences/" {
									elem = elem[l:]
								} else {
									break
								}

								// Param: "call_id"
								// Leaf parameter, slashes are prohibited
								idx := strings.IndexByte(elem, '/')
								if idx >= 0 {
									break
								}
								args[0] = elem
								elem = ""

								if len(elem) == 0 {
									// Leaf node.
									switch r.Method {
									case "GET":
										s.handleGetConferenceCallRequest([1]string{
											args[0],
										}, elemIsEscaped, w, r)
									default:
										s.notAllowed(w, r, notAllowedParams{
											allowedMethods: "GET",
											allowedHeaders: nil,
											acceptPost:     "",
											acceptPatch:    "",
										})
									}

									return
								}

							case 'i': // Prefix: "invite"

								if l := len("invite"); len(elem) >= l && elem[0:l] == "invite" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch r.Method {
									case "POST":
										s.handleInviteToCallRequest([0]string{}, elemIsEscaped, w, r)
									default:
										s.notAllowed(w, r, notAllowedParams{
											allowedMethods: "POST",
											allowedHeaders: rn221AllowedHeaders,
											acceptPost:     "multipart/form-data",
											acceptPatch:    "",
										})
									}

									return
								}

							case 's': // Prefix: "start_conference_call"

								if l := len("start_conference_call"); len(elem) >= l && elem[0:l] == "start_conference_call" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch r.Method {
									case "POST":
										s.handleStartConferenceCallRequest([0]string{}, elemIsEscaped, w, r)
									default:
										s.notAllowed(w, r, notAllowedParams{
											allowedMethods: "POST",
											allowedHeaders: rn319AllowedHeaders,
											acceptPost:     "multipart/form-data",
											acceptPatch:    "",
										})
									}

									return
								}

							}

						case 'h': // Prefix: "hat_rooms/"

							if l := len("hat_rooms/"); len(elem) >= l && elem[0:l] == "hat_rooms/" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								break
							}
							switch elem[0] {
							case 'u': // Prefix: "update"
								origElem := elem
								if l := len("update"); len(elem) >= l && elem[0:l] == "update" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch r.Method {
									case "GET":
										s.handleGetUpdatedChatRoomsRequest([0]string{}, elemIsEscaped, w, r)
									default:
										s.notAllowed(w, r, notAllowedParams{
											allowedMethods: "GET",
											allowedHeaders: nil,
											acceptPost:     "",
											acceptPatch:    "",
										})
									}

									return
								}

								elem = origElem
							}
							// Param: "id"
							// Match until "/"
							idx := strings.IndexByte(elem, '/')
							if idx < 0 {
								idx = len(elem)
							}
							args[0] = elem[:idx]
							elem = elem[idx:]

							if len(elem) == 0 {
								switch r.Method {
								case "GET":
									s.handleGetChatRoomRequest([1]string{
										args[0],
									}, elemIsEscaped, w, r)
								default:
									s.notAllowed(w, r, notAllowedParams{
										allowedMethods: "GET",
										allowedHeaders: nil,
										acceptPost:     "",
										acceptPatch:    "",
									})
								}

								return
							}
							switch elem[0] {
							case '/': // Prefix: "/"

								if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									break
								}
								switch elem[0] {
								case 'b': // Prefix: "background"

									if l := len("background"); len(elem) >= l && elem[0:l] == "background" {
										elem = elem[l:]
									} else {
										break
									}

									if len(elem) == 0 {
										// Leaf node.
										switch r.Method {
										case "DELETE":
											s.handleRemoveChatRoomBackgroundRequest([1]string{
												args[0],
											}, elemIsEscaped, w, r)
										default:
											s.notAllowed(w, r, notAllowedParams{
												allowedMethods: "DELETE",
												allowedHeaders: nil,
												acceptPost:     "",
												acceptPatch:    "",
											})
										}

										return
									}

								case 'i': // Prefix: "invite"

									if l := len("invite"); len(elem) >= l && elem[0:l] == "invite" {
										elem = elem[l:]
									} else {
										break
									}

									if len(elem) == 0 {
										// Leaf node.
										switch r.Method {
										case "POST":
											s.handleInviteToChatRoomRequest([1]string{
												args[0],
											}, elemIsEscaped, w, r)
										default:
											s.notAllowed(w, r, notAllowedParams{
												allowedMethods: "POST",
												allowedHeaders: rn223AllowedHeaders,
												acceptPost:     "multipart/form-data",
												acceptPatch:    "",
											})
										}

										return
									}

								case 'k': // Prefix: "kick"

									if l := len("kick"); len(elem) >= l && elem[0:l] == "kick" {
										elem = elem[l:]
									} else {
										break
									}

									if len(elem) == 0 {
										// Leaf node.
										switch r.Method {
										case "POST":
											s.handleKickFromChatRoomRequest([1]string{
												args[0],
											}, elemIsEscaped, w, r)
										default:
											s.notAllowed(w, r, notAllowedParams{
												allowedMethods: "POST",
												allowedHeaders: rn229AllowedHeaders,
												acceptPost:     "multipart/form-data",
												acceptPatch:    "",
											})
										}

										return
									}

								case 'm': // Prefix: "messages"

									if l := len("messages"); len(elem) >= l && elem[0:l] == "messages" {
										elem = elem[l:]
									} else {
										break
									}

									if len(elem) == 0 {
										switch r.Method {
										case "GET":
											s.handleGetChatMessagesRequest([1]string{
												args[0],
											}, elemIsEscaped, w, r)
										default:
											s.notAllowed(w, r, notAllowedParams{
												allowedMethods: "GET",
												allowedHeaders: nil,
												acceptPost:     "",
												acceptPatch:    "",
											})
										}

										return
									}
									switch elem[0] {
									case '/': // Prefix: "/"

										if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
											elem = elem[l:]
										} else {
											break
										}

										// Param: "message_id"
										// Match until "/"
										idx := strings.IndexByte(elem, '/')
										if idx < 0 {
											idx = len(elem)
										}
										args[1] = elem[:idx]
										elem = elem[idx:]

										if len(elem) == 0 {
											break
										}
										switch elem[0] {
										case '/': // Prefix: "/read"

											if l := len("/read"); len(elem) >= l && elem[0:l] == "/read" {
												elem = elem[l:]
											} else {
												break
											}

											if len(elem) == 0 {
												// Leaf node.
												switch r.Method {
												case "POST":
													s.handleReadChatMessageRequest([2]string{
														args[0],
														args[1],
													}, elemIsEscaped, w, r)
												default:
													s.notAllowed(w, r, notAllowedParams{
														allowedMethods: "POST",
														allowedHeaders: nil,
														acceptPost:     "",
														acceptPatch:    "",
													})
												}

												return
											}

										}

									}

								}

							}

						case 'o': // Prefix: "onversations/"

							if l := len("onversations/"); len(elem) >= l && elem[0:l] == "onversations/" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								break
							}
							switch elem[0] {
							case 'r': // Prefix: "root_posts"
								origElem := elem
								if l := len("root_posts"); len(elem) >= l && elem[0:l] == "root_posts" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch r.Method {
									case "GET":
										s.handleGetRootPostsRequest([0]string{}, elemIsEscaped, w, r)
									default:
										s.notAllowed(w, r, notAllowedParams{
											allowedMethods: "GET",
											allowedHeaders: nil,
											acceptPost:     "",
											acceptPatch:    "",
										})
									}

									return
								}

								elem = origElem
							}
							// Param: "id"
							// Leaf parameter, slashes are prohibited
							idx := strings.IndexByte(elem, '/')
							if idx >= 0 {
								break
							}
							args[0] = elem
							elem = ""

							if len(elem) == 0 {
								// Leaf node.
								switch r.Method {
								case "GET":
									s.handleGetConversationRequest([1]string{
										args[0],
									}, elemIsEscaped, w, r)
								default:
									s.notAllowed(w, r, notAllowedParams{
										allowedMethods: "GET",
										allowedHeaders: nil,
										acceptPost:     "",
										acceptPatch:    "",
									})
								}

								return
							}

						}

					case 'g': // Prefix: "g"

						if l := len("g"); len(elem) >= l && elem[0:l] == "g" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							break
						}
						switch elem[0] {
						case 'i': // Prefix: "ifts"

							if l := len("ifts"); len(elem) >= l && elem[0:l] == "ifts" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch r.Method {
								case "GET":
									s.handleListGiftsRequest([0]string{}, elemIsEscaped, w, r)
								default:
									s.notAllowed(w, r, notAllowedParams{
										allowedMethods: "GET",
										allowedHeaders: nil,
										acceptPost:     "",
										acceptPatch:    "",
									})
								}

								return
							}

						case 'r': // Prefix: "roups"

							if l := len("roups"); len(elem) >= l && elem[0:l] == "roups" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								switch r.Method {
								case "GET":
									s.handleListGroupsRequest([0]string{}, elemIsEscaped, w, r)
								default:
									s.notAllowed(w, r, notAllowedParams{
										allowedMethods: "GET",
										allowedHeaders: nil,
										acceptPost:     "",
										acceptPatch:    "",
									})
								}

								return
							}
							switch elem[0] {
							case '/': // Prefix: "/"

								if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									break
								}
								switch elem[0] {
								case 'm': // Prefix: "mine"
									origElem := elem
									if l := len("mine"); len(elem) >= l && elem[0:l] == "mine" {
										elem = elem[l:]
									} else {
										break
									}

									if len(elem) == 0 {
										// Leaf node.
										switch r.Method {
										case "GET":
											s.handleListMyGroupsRequest([0]string{}, elemIsEscaped, w, r)
										default:
											s.notAllowed(w, r, notAllowedParams{
												allowedMethods: "GET",
												allowedHeaders: nil,
												acceptPost:     "",
												acceptPatch:    "",
											})
										}

										return
									}

									elem = origElem
								}
								// Param: "id"
								// Match until "/"
								idx := strings.IndexByte(elem, '/')
								if idx < 0 {
									idx = len(elem)
								}
								args[0] = elem[:idx]
								elem = elem[idx:]

								if len(elem) == 0 {
									break
								}
								switch elem[0] {
								case '/': // Prefix: "/"

									if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
										elem = elem[l:]
									} else {
										break
									}

									if len(elem) == 0 {
										break
									}
									switch elem[0] {
									case 'm': // Prefix: "members"

										if l := len("members"); len(elem) >= l && elem[0:l] == "members" {
											elem = elem[l:]
										} else {
											break
										}

										if len(elem) == 0 {
											// Leaf node.
											switch r.Method {
											case "GET":
												s.handleGetGroupMembersRequest([1]string{
													args[0],
												}, elemIsEscaped, w, r)
											default:
												s.notAllowed(w, r, notAllowedParams{
													allowedMethods: "GET",
													allowedHeaders: nil,
													acceptPost:     "",
													acceptPatch:    "",
												})
											}

											return
										}

									case 'p': // Prefix: "posts/search"

										if l := len("posts/search"); len(elem) >= l && elem[0:l] == "posts/search" {
											elem = elem[l:]
										} else {
											break
										}

										if len(elem) == 0 {
											// Leaf node.
											switch r.Method {
											case "GET":
												s.handleSearchGroupPostsRequest([1]string{
													args[0],
												}, elemIsEscaped, w, r)
											default:
												s.notAllowed(w, r, notAllowedParams{
													allowedMethods: "GET",
													allowedHeaders: nil,
													acceptPost:     "",
													acceptPatch:    "",
												})
											}

											return
										}

									}

								}

							}

						}

					case 'n': // Prefix: "notification_settings/"

						if l := len("notification_settings/"); len(elem) >= l && elem[0:l] == "notification_settings/" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							break
						}
						switch elem[0] {
						case 'c': // Prefix: "chat_rooms/"

							if l := len("chat_rooms/"); len(elem) >= l && elem[0:l] == "chat_rooms/" {
								elem = elem[l:]
							} else {
								break
							}

							// Param: "id"
							// Leaf parameter, slashes are prohibited
							idx := strings.IndexByte(elem, '/')
							if idx >= 0 {
								break
							}
							args[0] = elem
							elem = ""

							if len(elem) == 0 {
								// Leaf node.
								switch r.Method {
								case "POST":
									s.handleUpdateChatRoomNotificationSettingsRequest([1]string{
										args[0],
									}, elemIsEscaped, w, r)
								default:
									s.notAllowed(w, r, notAllowedParams{
										allowedMethods: "POST",
										allowedHeaders: rn343AllowedHeaders,
										acceptPost:     "multipart/form-data",
										acceptPatch:    "",
									})
								}

								return
							}

						case 'g': // Prefix: "groups/"

							if l := len("groups/"); len(elem) >= l && elem[0:l] == "groups/" {
								elem = elem[l:]
							} else {
								break
							}

							// Param: "id"
							// Leaf parameter, slashes are prohibited
							idx := strings.IndexByte(elem, '/')
							if idx >= 0 {
								break
							}
							args[0] = elem
							elem = ""

							if len(elem) == 0 {
								// Leaf node.
								switch r.Method {
								case "GET":
									s.handleGetGroupNotificationSettingsRequest([1]string{
										args[0],
									}, elemIsEscaped, w, r)
								case "POST":
									s.handleUpdateGroupNotificationSettingsRequest([1]string{
										args[0],
									}, elemIsEscaped, w, r)
								default:
									s.notAllowed(w, r, notAllowedParams{
										allowedMethods: "GET,POST",
										allowedHeaders: rn138AllowedHeaders,
										acceptPost:     "multipart/form-data",
										acceptPatch:    "",
									})
								}

								return
							}

						}

					case 'p': // Prefix: "posts/"

						if l := len("posts/"); len(elem) >= l && elem[0:l] == "posts/" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							break
						}
						switch elem[0] {
						case 'c': // Prefix: "call_"
							origElem := elem
							if l := len("call_"); len(elem) >= l && elem[0:l] == "call_" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								break
							}
							switch elem[0] {
							case 'f': // Prefix: "followers_timeline"

								if l := len("followers_timeline"); len(elem) >= l && elem[0:l] == "followers_timeline" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch r.Method {
									case "GET":
										s.handleGetCallFollowersTimelineRequest([0]string{}, elemIsEscaped, w, r)
									default:
										s.notAllowed(w, r, notAllowedParams{
											allowedMethods: "GET",
											allowedHeaders: nil,
											acceptPost:     "",
											acceptPatch:    "",
										})
									}

									return
								}

							case 't': // Prefix: "timeline"

								if l := len("timeline"); len(elem) >= l && elem[0:l] == "timeline" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch r.Method {
									case "GET":
										s.handleGetCallTimelineRequest([0]string{}, elemIsEscaped, w, r)
									default:
										s.notAllowed(w, r, notAllowedParams{
											allowedMethods: "GET",
											allowedHeaders: nil,
											acceptPost:     "",
											acceptPatch:    "",
										})
									}

									return
								}

							}

							elem = origElem
						case 'f': // Prefix: "following_timeline"
							origElem := elem
							if l := len("following_timeline"); len(elem) >= l && elem[0:l] == "following_timeline" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch r.Method {
								case "GET":
									s.handleGetFollowingTimelineRequest([0]string{}, elemIsEscaped, w, r)
								default:
									s.notAllowed(w, r, notAllowedParams{
										allowedMethods: "GET",
										allowedHeaders: nil,
										acceptPost:     "",
										acceptPatch:    "",
									})
								}

								return
							}

							elem = origElem
						case 'g': // Prefix: "group_"
							origElem := elem
							if l := len("group_"); len(elem) >= l && elem[0:l] == "group_" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								break
							}
							switch elem[0] {
							case 'p': // Prefix: "pinned_post"

								if l := len("pinned_post"); len(elem) >= l && elem[0:l] == "pinned_post" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch r.Method {
									case "DELETE":
										s.handleUnpinGroupPostRequest([0]string{}, elemIsEscaped, w, r)
									case "PUT":
										s.handlePinGroupPostRequest([0]string{}, elemIsEscaped, w, r)
									default:
										s.notAllowed(w, r, notAllowedParams{
											allowedMethods: "DELETE,PUT",
											allowedHeaders: rn277AllowedHeaders,
											acceptPost:     "",
											acceptPatch:    "",
										})
									}

									return
								}

							case 't': // Prefix: "timeline"

								if l := len("timeline"); len(elem) >= l && elem[0:l] == "timeline" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch r.Method {
									case "GET":
										s.handleGetGroupTimelineRequest([0]string{}, elemIsEscaped, w, r)
									default:
										s.notAllowed(w, r, notAllowedParams{
											allowedMethods: "GET",
											allowedHeaders: nil,
											acceptPost:     "",
											acceptPatch:    "",
										})
									}

									return
								}

							}

							elem = origElem
						case 'l': // Prefix: "like"
							origElem := elem
							if l := len("like"); len(elem) >= l && elem[0:l] == "like" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch r.Method {
								case "POST":
									s.handleLikePostsRequest([0]string{}, elemIsEscaped, w, r)
								default:
									s.notAllowed(w, r, notAllowedParams{
										allowedMethods: "POST",
										allowedHeaders: rn237AllowedHeaders,
										acceptPost:     "multipart/form-data",
										acceptPatch:    "",
									})
								}

								return
							}

							elem = origElem
						case 'm': // Prefix: "m"
							origElem := elem
							if l := len("m"); len(elem) >= l && elem[0:l] == "m" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								break
							}
							switch elem[0] {
							case 'a': // Prefix: "ass_destroy"

								if l := len("ass_destroy"); len(elem) >= l && elem[0:l] == "ass_destroy" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch r.Method {
									case "POST":
										s.handleDeletePostsRequest([0]string{}, elemIsEscaped, w, r)
									default:
										s.notAllowed(w, r, notAllowedParams{
											allowedMethods: "POST",
											allowedHeaders: rn61AllowedHeaders,
											acceptPost:     "multipart/form-data",
											acceptPatch:    "",
										})
									}

									return
								}

							case 'i': // Prefix: "ine"

								if l := len("ine"); len(elem) >= l && elem[0:l] == "ine" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch r.Method {
									case "GET":
										s.handleGetMyPostsRequest([0]string{}, elemIsEscaped, w, r)
									default:
										s.notAllowed(w, r, notAllowedParams{
											allowedMethods: "GET",
											allowedHeaders: nil,
											acceptPost:     "",
											acceptPatch:    "",
										})
									}

									return
								}

							case 'u': // Prefix: "ultiple"

								if l := len("ultiple"); len(elem) >= l && elem[0:l] == "ultiple" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch r.Method {
									case "GET":
										s.handleGetPostsByIdsRequest([0]string{}, elemIsEscaped, w, r)
									default:
										s.notAllowed(w, r, notAllowedParams{
											allowedMethods: "GET",
											allowedHeaders: nil,
											acceptPost:     "",
											acceptPatch:    "",
										})
									}

									return
								}

							}

							elem = origElem
						case 'n': // Prefix: "new_"
							origElem := elem
							if l := len("new_"); len(elem) >= l && elem[0:l] == "new_" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								break
							}
							switch elem[0] {
							case 'c': // Prefix: "conference_call"

								if l := len("conference_call"); len(elem) >= l && elem[0:l] == "conference_call" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch r.Method {
									case "POST":
										s.handleCreateConferenceCallPostRequest([0]string{}, elemIsEscaped, w, r)
									default:
										s.notAllowed(w, r, notAllowedParams{
											allowedMethods: "POST",
											allowedHeaders: rn36AllowedHeaders,
											acceptPost:     "multipart/form-data",
											acceptPatch:    "",
										})
									}

									return
								}

							case 's': // Prefix: "share_post"

								if l := len("share_post"); len(elem) >= l && elem[0:l] == "share_post" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch r.Method {
									case "POST":
										s.handleCreateSharePostRequest([0]string{}, elemIsEscaped, w, r)
									default:
										s.notAllowed(w, r, notAllowedParams{
											allowedMethods: "POST",
											allowedHeaders: rn43AllowedHeaders,
											acceptPost:     "multipart/form-data",
											acceptPatch:    "",
										})
									}

									return
								}

							}

							elem = origElem
						case 'r': // Prefix: "rec"
							origElem := elem
							if l := len("rec"); len(elem) >= l && elem[0:l] == "rec" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								break
							}
							switch elem[0] {
							case 'e': // Prefix: "ent_engagement"

								if l := len("ent_engagement"); len(elem) >= l && elem[0:l] == "ent_engagement" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch r.Method {
									case "GET":
										s.handleGetRecentEngagementPostsRequest([0]string{}, elemIsEscaped, w, r)
									default:
										s.notAllowed(w, r, notAllowedParams{
											allowedMethods: "GET",
											allowedHeaders: nil,
											acceptPost:     "",
											acceptPatch:    "",
										})
									}

									return
								}

							case 'o': // Prefix: "ommended_timeline"

								if l := len("ommended_timeline"); len(elem) >= l && elem[0:l] == "ommended_timeline" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch r.Method {
									case "GET":
										s.handleGetRecommendedTimelineRequest([0]string{}, elemIsEscaped, w, r)
									default:
										s.notAllowed(w, r, notAllowedParams{
											allowedMethods: "GET",
											allowedHeaders: nil,
											acceptPost:     "",
											acceptPatch:    "",
										})
									}

									return
								}

							}

							elem = origElem
						case 's': // Prefix: "search"
							origElem := elem
							if l := len("search"); len(elem) >= l && elem[0:l] == "search" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch r.Method {
								case "GET":
									s.handleSearchPostsRequest([0]string{}, elemIsEscaped, w, r)
								default:
									s.notAllowed(w, r, notAllowedParams{
										allowedMethods: "GET",
										allowedHeaders: nil,
										acceptPost:     "",
										acceptPatch:    "",
									})
								}

								return
							}

							elem = origElem
						case 't': // Prefix: "tags/"
							origElem := elem
							if l := len("tags/"); len(elem) >= l && elem[0:l] == "tags/" {
								elem = elem[l:]
							} else {
								break
							}

							// Param: "tag"
							// Leaf parameter, slashes are prohibited
							idx := strings.IndexByte(elem, '/')
							if idx >= 0 {
								break
							}
							args[0] = elem
							elem = ""

							if len(elem) == 0 {
								// Leaf node.
								switch r.Method {
								case "GET":
									s.handleGetPostsByTagRequest([1]string{
										args[0],
									}, elemIsEscaped, w, r)
								default:
									s.notAllowed(w, r, notAllowedParams{
										allowedMethods: "GET",
										allowedHeaders: nil,
										acceptPost:     "",
										acceptPatch:    "",
									})
								}

								return
							}

							elem = origElem
						case 'u': // Prefix: "u"
							origElem := elem
							if l := len("u"); len(elem) >= l && elem[0:l] == "u" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								break
							}
							switch elem[0] {
							case 'r': // Prefix: "rl_metadata"

								if l := len("rl_metadata"); len(elem) >= l && elem[0:l] == "rl_metadata" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch r.Method {
									case "GET":
										s.handleGetPostUrlMetadataRequest([0]string{}, elemIsEscaped, w, r)
									default:
										s.notAllowed(w, r, notAllowedParams{
											allowedMethods: "GET",
											allowedHeaders: nil,
											acceptPost:     "",
											acceptPatch:    "",
										})
									}

									return
								}

							case 's': // Prefix: "ser_timeline"

								if l := len("ser_timeline"); len(elem) >= l && elem[0:l] == "ser_timeline" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch r.Method {
									case "GET":
										s.handleGetUserTimelineRequest([0]string{}, elemIsEscaped, w, r)
									default:
										s.notAllowed(w, r, notAllowedParams{
											allowedMethods: "GET",
											allowedHeaders: nil,
											acceptPost:     "",
											acceptPatch:    "",
										})
									}

									return
								}

							}

							elem = origElem
						}
						// Param: "id"
						// Match until one of "/t"
						idx := strings.IndexAny(elem, "/t")
						if idx < 0 {
							idx = len(elem)
						}
						args[0] = elem[:idx]
						elem = elem[idx:]

						if len(elem) == 0 {
							switch r.Method {
							case "GET":
								s.handleGetPostRequest([1]string{
									args[0],
								}, elemIsEscaped, w, r)
							default:
								s.notAllowed(w, r, notAllowedParams{
									allowedMethods: "GET",
									allowedHeaders: rn161AllowedHeaders,
									acceptPost:     "",
									acceptPatch:    "",
								})
							}

							return
						}
						switch elem[0] {
						case '/': // Prefix: "/reposts"

							if l := len("/reposts"); len(elem) >= l && elem[0:l] == "/reposts" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch r.Method {
								case "GET":
									s.handleGetPostRepostsRequest([1]string{
										args[0],
									}, elemIsEscaped, w, r)
								default:
									s.notAllowed(w, r, notAllowedParams{
										allowedMethods: "GET",
										allowedHeaders: nil,
										acceptPost:     "",
										acceptPatch:    "",
									})
								}

								return
							}

						case 't': // Prefix: "timeline"

							if l := len("timeline"); len(elem) >= l && elem[0:l] == "timeline" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch r.Method {
								case "GET":
									s.handleGetTimelineRequest([1]string{
										args[0],
									}, elemIsEscaped, w, r)
								default:
									s.notAllowed(w, r, notAllowedParams{
										allowedMethods: "GET",
										allowedHeaders: nil,
										acceptPost:     "",
										acceptPatch:    "",
									})
								}

								return
							}

						}

					case 'r': // Prefix: "received_gifts"

						if l := len("received_gifts"); len(elem) >= l && elem[0:l] == "received_gifts" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch r.Method {
							case "GET":
								s.handleListReceivedGiftsRequest([0]string{}, elemIsEscaped, w, r)
							default:
								s.notAllowed(w, r, notAllowedParams{
									allowedMethods: "GET",
									allowedHeaders: nil,
									acceptPost:     "",
									acceptPatch:    "",
								})
							}

							return
						}

					case 's': // Prefix: "s"

						if l := len("s"); len(elem) >= l && elem[0:l] == "s" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							break
						}
						switch elem[0] {
						case 't': // Prefix: "ticker_packs"

							if l := len("ticker_packs"); len(elem) >= l && elem[0:l] == "ticker_packs" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch r.Method {
								case "GET":
									s.handleListStickerPacksRequest([0]string{}, elemIsEscaped, w, r)
								default:
									s.notAllowed(w, r, notAllowedParams{
										allowedMethods: "GET",
										allowedHeaders: nil,
										acceptPost:     "",
										acceptPatch:    "",
									})
								}

								return
							}

						case 'u': // Prefix: "urveys/"

							if l := len("urveys/"); len(elem) >= l && elem[0:l] == "urveys/" {
								elem = elem[l:]
							} else {
								break
							}

							// Param: "id"
							// Match until "/"
							idx := strings.IndexByte(elem, '/')
							if idx < 0 {
								idx = len(elem)
							}
							args[0] = elem[:idx]
							elem = elem[idx:]

							if len(elem) == 0 {
								break
							}
							switch elem[0] {
							case '/': // Prefix: "/vote"

								if l := len("/vote"); len(elem) >= l && elem[0:l] == "/vote" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch r.Method {
									case "POST":
										s.handleVoteSurveyRequest([1]string{
											args[0],
										}, elemIsEscaped, w, r)
									default:
										s.notAllowed(w, r, notAllowedParams{
											allowedMethods: "POST",
											allowedHeaders: rn358AllowedHeaders,
											acceptPost:     "multipart/form-data",
											acceptPatch:    "",
										})
									}

									return
								}

							}

						}

					case 'u': // Prefix: "users/"

						if l := len("users/"); len(elem) >= l && elem[0:l] == "users/" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							break
						}
						switch elem[0] {
						case 'b': // Prefix: "blocked"
							origElem := elem
							if l := len("blocked"); len(elem) >= l && elem[0:l] == "blocked" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch r.Method {
								case "POST":
									s.handleGetBlockedUsersRequest([0]string{}, elemIsEscaped, w, r)
								default:
									s.notAllowed(w, r, notAllowedParams{
										allowedMethods: "POST",
										allowedHeaders: rn93AllowedHeaders,
										acceptPost:     "application/json",
										acceptPatch:    "",
									})
								}

								return
							}

							elem = origElem
						case 'e': // Prefix: "edit"
							origElem := elem
							if l := len("edit"); len(elem) >= l && elem[0:l] == "edit" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch r.Method {
								case "POST":
									s.handleEditUserV2Request([0]string{}, elemIsEscaped, w, r)
								default:
									s.notAllowed(w, r, notAllowedParams{
										allowedMethods: "POST",
										allowedHeaders: rn69AllowedHeaders,
										acceptPost:     "multipart/form-data",
										acceptPatch:    "",
									})
								}

								return
							}

							elem = origElem
						case 'f': // Prefix: "f"
							origElem := elem
							if l := len("f"); len(elem) >= l && elem[0:l] == "f" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								break
							}
							switch elem[0] {
							case 'o': // Prefix: "ollow"

								if l := len("ollow"); len(elem) >= l && elem[0:l] == "ollow" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									switch r.Method {
									case "POST":
										s.handleFollowUsersRequest([0]string{}, elemIsEscaped, w, r)
									default:
										s.notAllowed(w, r, notAllowedParams{
											allowedMethods: "POST",
											allowedHeaders: nil,
											acceptPost:     "",
											acceptPatch:    "",
										})
									}

									return
								}
								switch elem[0] {
								case '_': // Prefix: "_requests"

									if l := len("_requests"); len(elem) >= l && elem[0:l] == "_requests" {
										elem = elem[l:]
									} else {
										break
									}

									if len(elem) == 0 {
										switch r.Method {
										case "GET":
											s.handleGetFollowRequestsRequest([0]string{}, elemIsEscaped, w, r)
										default:
											s.notAllowed(w, r, notAllowedParams{
												allowedMethods: "GET",
												allowedHeaders: nil,
												acceptPost:     "",
												acceptPatch:    "",
											})
										}

										return
									}
									switch elem[0] {
									case '_': // Prefix: "_count"

										if l := len("_count"); len(elem) >= l && elem[0:l] == "_count" {
											elem = elem[l:]
										} else {
											break
										}

										if len(elem) == 0 {
											// Leaf node.
											switch r.Method {
											case "GET":
												s.handleGetFollowRequestCountRequest([0]string{}, elemIsEscaped, w, r)
											default:
												s.notAllowed(w, r, notAllowedParams{
													allowedMethods: "GET",
													allowedHeaders: nil,
													acceptPost:     "",
													acceptPatch:    "",
												})
											}

											return
										}

									}

								}

							case 'r': // Prefix: "resh/"

								if l := len("resh/"); len(elem) >= l && elem[0:l] == "resh/" {
									elem = elem[l:]
								} else {
									break
								}

								// Param: "id"
								// Leaf parameter, slashes are prohibited
								idx := strings.IndexByte(elem, '/')
								if idx >= 0 {
									break
								}
								args[0] = elem
								elem = ""

								if len(elem) == 0 {
									// Leaf node.
									switch r.Method {
									case "GET":
										s.handleGetFreshUserRequest([1]string{
											args[0],
										}, elemIsEscaped, w, r)
									default:
										s.notAllowed(w, r, notAllowedParams{
											allowedMethods: "GET",
											allowedHeaders: nil,
											acceptPost:     "",
											acceptPatch:    "",
										})
									}

									return
								}

							}

							elem = origElem
						case 'h': // Prefix: "hima_users"
							origElem := elem
							if l := len("hima_users"); len(elem) >= l && elem[0:l] == "hima_users" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch r.Method {
								case "GET":
									s.handleGetHimaUsersRequest([0]string{}, elemIsEscaped, w, r)
								default:
									s.notAllowed(w, r, notAllowedParams{
										allowedMethods: "GET",
										allowedHeaders: nil,
										acceptPost:     "",
										acceptPatch:    "",
									})
								}

								return
							}

							elem = origElem
						case 'i': // Prefix: "info/"
							origElem := elem
							if l := len("info/"); len(elem) >= l && elem[0:l] == "info/" {
								elem = elem[l:]
							} else {
								break
							}

							// Param: "id"
							// Leaf parameter, slashes are prohibited
							idx := strings.IndexByte(elem, '/')
							if idx >= 0 {
								break
							}
							args[0] = elem
							elem = ""

							if len(elem) == 0 {
								// Leaf node.
								switch r.Method {
								case "GET":
									s.handleGetUserInfoRequest([1]string{
										args[0],
									}, elemIsEscaped, w, r)
								default:
									s.notAllowed(w, r, notAllowedParams{
										allowedMethods: "GET",
										allowedHeaders: nil,
										acceptPost:     "",
										acceptPatch:    "",
									})
								}

								return
							}

							elem = origElem
						case 'r': // Prefix: "re"
							origElem := elem
							if l := len("re"); len(elem) >= l && elem[0:l] == "re" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								break
							}
							switch elem[0] {
							case 'm': // Prefix: "move_"

								if l := len("move_"); len(elem) >= l && elem[0:l] == "move_" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									break
								}
								switch elem[0] {
								case 'c': // Prefix: "cover_image"

									if l := len("cover_image"); len(elem) >= l && elem[0:l] == "cover_image" {
										elem = elem[l:]
									} else {
										break
									}

									if len(elem) == 0 {
										// Leaf node.
										switch r.Method {
										case "POST":
											s.handleRemoveCoverImageRequest([0]string{}, elemIsEscaped, w, r)
										default:
											s.notAllowed(w, r, notAllowedParams{
												allowedMethods: "POST",
												allowedHeaders: nil,
												acceptPost:     "",
												acceptPatch:    "",
											})
										}

										return
									}

								case 'p': // Prefix: "profile_photo"

									if l := len("profile_photo"); len(elem) >= l && elem[0:l] == "profile_photo" {
										elem = elem[l:]
									} else {
										break
									}

									if len(elem) == 0 {
										// Leaf node.
										switch r.Method {
										case "POST":
											s.handleRemoveProfilePhotoRequest([0]string{}, elemIsEscaped, w, r)
										default:
											s.notAllowed(w, r, notAllowedParams{
												allowedMethods: "POST",
												allowedHeaders: nil,
												acceptPost:     "",
												acceptPatch:    "",
											})
										}

										return
									}

								}

							case 's': // Prefix: "send_confirm_email"

								if l := len("send_confirm_email"); len(elem) >= l && elem[0:l] == "send_confirm_email" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch r.Method {
									case "POST":
										s.handleResendConfirmEmailRequest([0]string{}, elemIsEscaped, w, r)
									default:
										s.notAllowed(w, r, notAllowedParams{
											allowedMethods: "POST",
											allowedHeaders: nil,
											acceptPost:     "",
											acceptPatch:    "",
										})
									}

									return
								}

							case 'v': // Prefix: "views/"

								if l := len("views/"); len(elem) >= l && elem[0:l] == "views/" {
									elem = elem[l:]
								} else {
									break
								}

								// Param: "id"
								// Leaf parameter, slashes are prohibited
								idx := strings.IndexByte(elem, '/')
								if idx >= 0 {
									break
								}
								args[0] = elem
								elem = ""

								if len(elem) == 0 {
									// Leaf node.
									switch r.Method {
									case "GET":
										s.handleGetUserReviewsRequest([1]string{
											args[0],
										}, elemIsEscaped, w, r)
									case "POST":
										s.handleReplyToReviewRequest([1]string{
											args[0],
										}, elemIsEscaped, w, r)
									default:
										s.notAllowed(w, r, notAllowedParams{
											allowedMethods: "GET,POST",
											allowedHeaders: rn211AllowedHeaders,
											acceptPost:     "multipart/form-data",
											acceptPatch:    "",
										})
									}

									return
								}

							}

							elem = origElem
						case 't': // Prefix: "timestamp"
							origElem := elem
							if l := len("timestamp"); len(elem) >= l && elem[0:l] == "timestamp" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch r.Method {
								case "GET":
									s.handleGetUserTimestampRequest([0]string{}, elemIsEscaped, w, r)
								default:
									s.notAllowed(w, r, notAllowedParams{
										allowedMethods: "GET",
										allowedHeaders: nil,
										acceptPost:     "",
										acceptPatch:    "",
									})
								}

								return
							}

							elem = origElem
						}
						// Param: "user_id"
						// Match until "/"
						idx := strings.IndexByte(elem, '/')
						if idx < 0 {
							idx = len(elem)
						}
						args[0] = elem[:idx]
						elem = elem[idx:]

						if len(elem) == 0 {
							switch r.Method {
							case "GET":
								s.handleGetUserRequest([1]string{
									args[0],
								}, elemIsEscaped, w, r)
							default:
								s.notAllowed(w, r, notAllowedParams{
									allowedMethods: "GET",
									allowedHeaders: nil,
									acceptPost:     "",
									acceptPatch:    "",
								})
							}

							return
						}
						switch elem[0] {
						case '/': // Prefix: "/"

							if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								break
							}
							switch elem[0] {
							case 'f': // Prefix: "fo"

								if l := len("fo"); len(elem) >= l && elem[0:l] == "fo" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									break
								}
								switch elem[0] {
								case 'l': // Prefix: "llow"

									if l := len("llow"); len(elem) >= l && elem[0:l] == "llow" {
										elem = elem[l:]
									} else {
										break
									}

									if len(elem) == 0 {
										switch r.Method {
										case "POST":
											s.handleFollowUserRequest([1]string{
												args[0],
											}, elemIsEscaped, w, r)
										default:
											s.notAllowed(w, r, notAllowedParams{
												allowedMethods: "POST",
												allowedHeaders: nil,
												acceptPost:     "",
												acceptPatch:    "",
											})
										}

										return
									}
									switch elem[0] {
									case '_': // Prefix: "_request"

										if l := len("_request"); len(elem) >= l && elem[0:l] == "_request" {
											elem = elem[l:]
										} else {
											break
										}

										if len(elem) == 0 {
											// Leaf node.
											switch r.Method {
											case "POST":
												s.handleRequestFollowRequest([1]string{
													args[0],
												}, elemIsEscaped, w, r)
											default:
												s.notAllowed(w, r, notAllowedParams{
													allowedMethods: "POST",
													allowedHeaders: rn305AllowedHeaders,
													acceptPost:     "multipart/form-data",
													acceptPatch:    "",
												})
											}

											return
										}

									}

								case 'o': // Prefix: "otprints/"

									if l := len("otprints/"); len(elem) >= l && elem[0:l] == "otprints/" {
										elem = elem[l:]
									} else {
										break
									}

									// Param: "footprint_id"
									// Leaf parameter, slashes are prohibited
									idx := strings.IndexByte(elem, '/')
									if idx >= 0 {
										break
									}
									args[1] = elem
									elem = ""

									if len(elem) == 0 {
										// Leaf node.
										switch r.Method {
										case "DELETE":
											s.handleDeleteFootprintRequest([2]string{
												args[0],
												args[1],
											}, elemIsEscaped, w, r)
										default:
											s.notAllowed(w, r, notAllowedParams{
												allowedMethods: "DELETE",
												allowedHeaders: nil,
												acceptPost:     "",
												acceptPatch:    "",
											})
										}

										return
									}

								}

							case 'u': // Prefix: "un"

								if l := len("un"); len(elem) >= l && elem[0:l] == "un" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									break
								}
								switch elem[0] {
								case 'b': // Prefix: "block"

									if l := len("block"); len(elem) >= l && elem[0:l] == "block" {
										elem = elem[l:]
									} else {
										break
									}

									if len(elem) == 0 {
										// Leaf node.
										switch r.Method {
										case "POST":
											s.handleUnblockUserRequest([1]string{
												args[0],
											}, elemIsEscaped, w, r)
										default:
											s.notAllowed(w, r, notAllowedParams{
												allowedMethods: "POST",
												allowedHeaders: nil,
												acceptPost:     "",
												acceptPatch:    "",
											})
										}

										return
									}

								case 'f': // Prefix: "follow"

									if l := len("follow"); len(elem) >= l && elem[0:l] == "follow" {
										elem = elem[l:]
									} else {
										break
									}

									if len(elem) == 0 {
										// Leaf node.
										switch r.Method {
										case "POST":
											s.handleUnfollowUserRequest([1]string{
												args[0],
											}, elemIsEscaped, w, r)
										default:
											s.notAllowed(w, r, notAllowedParams{
												allowedMethods: "POST",
												allowedHeaders: nil,
												acceptPost:     "",
												acceptPatch:    "",
											})
										}

										return
									}

								}

							}

						}

					}

				case '3': // Prefix: "3/"

					if l := len("3/"); len(elem) >= l && elem[0:l] == "3/" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						break
					}
					switch elem[0] {
					case 'c': // Prefix: "c"

						if l := len("c"); len(elem) >= l && elem[0:l] == "c" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							break
						}
						switch elem[0] {
						case 'a': // Prefix: "alls/"

							if l := len("alls/"); len(elem) >= l && elem[0:l] == "alls/" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								break
							}
							switch elem[0] {
							case 'c': // Prefix: "conference_calls/"
								origElem := elem
								if l := len("conference_calls/"); len(elem) >= l && elem[0:l] == "conference_calls/" {
									elem = elem[l:]
								} else {
									break
								}

								// Param: "call_id"
								// Match until "/"
								idx := strings.IndexByte(elem, '/')
								if idx < 0 {
									idx = len(elem)
								}
								args[0] = elem[:idx]
								elem = elem[idx:]

								if len(elem) == 0 {
									break
								}
								switch elem[0] {
								case '/': // Prefix: "/kick"

									if l := len("/kick"); len(elem) >= l && elem[0:l] == "/kick" {
										elem = elem[l:]
									} else {
										break
									}

									if len(elem) == 0 {
										// Leaf node.
										switch r.Method {
										case "POST":
											s.handleKickFromConferenceCallRequest([1]string{
												args[0],
											}, elemIsEscaped, w, r)
										default:
											s.notAllowed(w, r, notAllowedParams{
												allowedMethods: "POST",
												allowedHeaders: rn232AllowedHeaders,
												acceptPost:     "multipart/form-data",
												acceptPatch:    "",
											})
										}

										return
									}

								}

								elem = origElem
							}
							// Param: "call_id"
							// Match until "/"
							idx := strings.IndexByte(elem, '/')
							if idx < 0 {
								idx = len(elem)
							}
							args[0] = elem[:idx]
							elem = elem[idx:]

							if len(elem) == 0 {
								break
							}
							switch elem[0] {
							case '/': // Prefix: "/agora_rtm_token"

								if l := len("/agora_rtm_token"); len(elem) >= l && elem[0:l] == "/agora_rtm_token" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch r.Method {
									case "GET":
										s.handleGetAgoraRtmTokenRequest([1]string{
											args[0],
										}, elemIsEscaped, w, r)
									default:
										s.notAllowed(w, r, notAllowedParams{
											allowedMethods: "GET",
											allowedHeaders: nil,
											acceptPost:     "",
											acceptPatch:    "",
										})
									}

									return
								}

							}

						case 'h': // Prefix: "hat_rooms/"

							if l := len("hat_rooms/"); len(elem) >= l && elem[0:l] == "hat_rooms/" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								break
							}
							switch elem[0] {
							case 'n': // Prefix: "new"
								origElem := elem
								if l := len("new"); len(elem) >= l && elem[0:l] == "new" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch r.Method {
									case "POST":
										s.handleCreateChatRoomRequest([0]string{}, elemIsEscaped, w, r)
									default:
										s.notAllowed(w, r, notAllowedParams{
											allowedMethods: "POST",
											allowedHeaders: rn33AllowedHeaders,
											acceptPost:     "multipart/form-data",
											acceptPatch:    "",
										})
									}

									return
								}

								elem = origElem
							}
							// Param: "chat_room_id"
							// Match until "/"
							idx := strings.IndexByte(elem, '/')
							if idx < 0 {
								idx = len(elem)
							}
							args[0] = elem[:idx]
							elem = elem[idx:]

							if len(elem) == 0 {
								break
							}
							switch elem[0] {
							case '/': // Prefix: "/"

								if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									break
								}
								switch elem[0] {
								case 'e': // Prefix: "edit"

									if l := len("edit"); len(elem) >= l && elem[0:l] == "edit" {
										elem = elem[l:]
									} else {
										break
									}

									if len(elem) == 0 {
										// Leaf node.
										switch r.Method {
										case "POST":
											s.handleUpdateChatRoomRequest([1]string{
												args[0],
											}, elemIsEscaped, w, r)
										default:
											s.notAllowed(w, r, notAllowedParams{
												allowedMethods: "POST",
												allowedHeaders: rn340AllowedHeaders,
												acceptPost:     "multipart/form-data",
												acceptPatch:    "",
											})
										}

										return
									}

								case 'm': // Prefix: "messages/new"

									if l := len("messages/new"); len(elem) >= l && elem[0:l] == "messages/new" {
										elem = elem[l:]
									} else {
										break
									}

									if len(elem) == 0 {
										// Leaf node.
										switch r.Method {
										case "POST":
											s.handleSendChatMessageRequest([1]string{
												args[0],
											}, elemIsEscaped, w, r)
										default:
											s.notAllowed(w, r, notAllowedParams{
												allowedMethods: "POST",
												allowedHeaders: rn316AllowedHeaders,
												acceptPost:     "multipart/form-data",
												acceptPatch:    "",
											})
										}

										return
									}

								case 'r': // Prefix: "report"

									if l := len("report"); len(elem) >= l && elem[0:l] == "report" {
										elem = elem[l:]
									} else {
										break
									}

									if len(elem) == 0 {
										// Leaf node.
										switch r.Method {
										case "POST":
											s.handleReportChatRoomRequest([1]string{
												args[0],
											}, elemIsEscaped, w, r)
										default:
											s.notAllowed(w, r, notAllowedParams{
												allowedMethods: "POST",
												allowedHeaders: rn297AllowedHeaders,
												acceptPost:     "multipart/form-data",
												acceptPatch:    "",
											})
										}

										return
									}

								}

							}

						}

					case 'g': // Prefix: "groups/"

						if l := len("groups/"); len(elem) >= l && elem[0:l] == "groups/" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							break
						}
						switch elem[0] {
						case 'n': // Prefix: "new"
							origElem := elem
							if l := len("new"); len(elem) >= l && elem[0:l] == "new" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch r.Method {
								case "POST":
									s.handleCreateGroupRequest([0]string{}, elemIsEscaped, w, r)
								default:
									s.notAllowed(w, r, notAllowedParams{
										allowedMethods: "POST",
										allowedHeaders: rn38AllowedHeaders,
										acceptPost:     "multipart/form-data",
										acceptPatch:    "",
									})
								}

								return
							}

							elem = origElem
						}
						// Param: "group_id"
						// Match until "/"
						idx := strings.IndexByte(elem, '/')
						if idx < 0 {
							idx = len(elem)
						}
						args[0] = elem[:idx]
						elem = elem[idx:]

						if len(elem) == 0 {
							break
						}
						switch elem[0] {
						case '/': // Prefix: "/"

							if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								break
							}
							switch elem[0] {
							case 'c': // Prefix: "cover"

								if l := len("cover"); len(elem) >= l && elem[0:l] == "cover" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch r.Method {
									case "DELETE":
										s.handleRemoveGroupCoverRequest([1]string{
											args[0],
										}, elemIsEscaped, w, r)
									default:
										s.notAllowed(w, r, notAllowedParams{
											allowedMethods: "DELETE",
											allowedHeaders: nil,
											acceptPost:     "",
											acceptPatch:    "",
										})
									}

									return
								}

							case 'd': // Prefix: "deputize/mass"

								if l := len("deputize/mass"); len(elem) >= l && elem[0:l] == "deputize/mass" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch r.Method {
									case "POST":
										s.handleDeputizeGroupUsersMassRequest([1]string{
											args[0],
										}, elemIsEscaped, w, r)
									default:
										s.notAllowed(w, r, notAllowedParams{
											allowedMethods: "POST",
											allowedHeaders: rn66AllowedHeaders,
											acceptPost:     "multipart/form-data",
											acceptPatch:    "",
										})
									}

									return
								}

							case 'i': // Prefix: "icon"

								if l := len("icon"); len(elem) >= l && elem[0:l] == "icon" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch r.Method {
									case "DELETE":
										s.handleRemoveGroupIconRequest([1]string{
											args[0],
										}, elemIsEscaped, w, r)
									default:
										s.notAllowed(w, r, notAllowedParams{
											allowedMethods: "DELETE",
											allowedHeaders: nil,
											acceptPost:     "",
											acceptPatch:    "",
										})
									}

									return
								}

							case 'r': // Prefix: "report"

								if l := len("report"); len(elem) >= l && elem[0:l] == "report" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch r.Method {
									case "POST":
										s.handleReportGroupRequest([1]string{
											args[0],
										}, elemIsEscaped, w, r)
									default:
										s.notAllowed(w, r, notAllowedParams{
											allowedMethods: "POST",
											allowedHeaders: rn298AllowedHeaders,
											acceptPost:     "multipart/form-data",
											acceptPatch:    "",
										})
									}

									return
								}

							case 't': // Prefix: "transfer"

								if l := len("transfer"); len(elem) >= l && elem[0:l] == "transfer" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch r.Method {
									case "POST":
										s.handleTransferGroupRequest([1]string{
											args[0],
										}, elemIsEscaped, w, r)
									default:
										s.notAllowed(w, r, notAllowedParams{
											allowedMethods: "POST",
											allowedHeaders: rn322AllowedHeaders,
											acceptPost:     "multipart/form-data",
											acceptPatch:    "",
										})
									}

									return
								}

							case 'u': // Prefix: "update"

								if l := len("update"); len(elem) >= l && elem[0:l] == "update" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch r.Method {
									case "POST":
										s.handleUpdateGroupRequest([1]string{
											args[0],
										}, elemIsEscaped, w, r)
									default:
										s.notAllowed(w, r, notAllowedParams{
											allowedMethods: "POST",
											allowedHeaders: rn344AllowedHeaders,
											acceptPost:     "multipart/form-data",
											acceptPatch:    "",
										})
									}

									return
								}

							}

						}

					case 'p': // Prefix: "posts/"

						if l := len("posts/"); len(elem) >= l && elem[0:l] == "posts/" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							break
						}
						switch elem[0] {
						case 'n': // Prefix: "new"
							origElem := elem
							if l := len("new"); len(elem) >= l && elem[0:l] == "new" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch r.Method {
								case "POST":
									s.handleCreatePostRequest([0]string{}, elemIsEscaped, w, r)
								default:
									s.notAllowed(w, r, notAllowedParams{
										allowedMethods: "POST",
										allowedHeaders: rn40AllowedHeaders,
										acceptPost:     "multipart/form-data",
										acceptPatch:    "",
									})
								}

								return
							}

							elem = origElem
						case 'r': // Prefix: "repost"
							origElem := elem
							if l := len("repost"); len(elem) >= l && elem[0:l] == "repost" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch r.Method {
								case "POST":
									s.handleRepostRequest([0]string{}, elemIsEscaped, w, r)
								default:
									s.notAllowed(w, r, notAllowedParams{
										allowedMethods: "POST",
										allowedHeaders: rn303AllowedHeaders,
										acceptPost:     "multipart/form-data",
										acceptPatch:    "",
									})
								}

								return
							}

							elem = origElem
						}
						// Param: "id"
						// Match until "/"
						idx := strings.IndexByte(elem, '/')
						if idx < 0 {
							idx = len(elem)
						}
						args[0] = elem[:idx]
						elem = elem[idx:]

						if len(elem) == 0 {
							switch r.Method {
							case "PUT":
								s.handleUpdatePostRequest([1]string{
									args[0],
								}, elemIsEscaped, w, r)
							default:
								s.notAllowed(w, r, notAllowedParams{
									allowedMethods: "PUT",
									allowedHeaders: rn261AllowedHeaders,
									acceptPost:     "",
									acceptPatch:    "",
								})
							}

							return
						}
						switch elem[0] {
						case '/': // Prefix: "/"

							if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								break
							}
							switch elem[0] {
							case 'm': // Prefix: "move_to_thread"

								if l := len("move_to_thread"); len(elem) >= l && elem[0:l] == "move_to_thread" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									switch r.Method {
									case "POST":
										s.handleMovePostToThreadRequest([1]string{
											args[0],
										}, elemIsEscaped, w, r)
									default:
										s.notAllowed(w, r, notAllowedParams{
											allowedMethods: "POST",
											allowedHeaders: rn264AllowedHeaders,
											acceptPost:     "multipart/form-data",
											acceptPatch:    "",
										})
									}

									return
								}
								switch elem[0] {
								case '/': // Prefix: "/"

									if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
										elem = elem[l:]
									} else {
										break
									}

									// Param: "thread_id"
									// Leaf parameter, slashes are prohibited
									idx := strings.IndexByte(elem, '/')
									if idx >= 0 {
										break
									}
									args[1] = elem
									elem = ""

									if len(elem) == 0 {
										// Leaf node.
										switch r.Method {
										case "PUT":
											s.handleMovePostToSpecificThreadRequest([2]string{
												args[0],
												args[1],
											}, elemIsEscaped, w, r)
										default:
											s.notAllowed(w, r, notAllowedParams{
												allowedMethods: "PUT",
												allowedHeaders: nil,
												acceptPost:     "",
												acceptPatch:    "",
											})
										}

										return
									}

								}

							case 'r': // Prefix: "report"

								if l := len("report"); len(elem) >= l && elem[0:l] == "report" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch r.Method {
									case "POST":
										s.handleReportPostRequest([1]string{
											args[0],
										}, elemIsEscaped, w, r)
									default:
										s.notAllowed(w, r, notAllowedParams{
											allowedMethods: "POST",
											allowedHeaders: rn300AllowedHeaders,
											acceptPost:     "multipart/form-data",
											acceptPatch:    "",
										})
									}

									return
								}

							}

						}

					case 'u': // Prefix: "users/"

						if l := len("users/"); len(elem) >= l && elem[0:l] == "users/" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							break
						}
						switch elem[0] {
						case 'e': // Prefix: "edit"
							origElem := elem
							if l := len("edit"); len(elem) >= l && elem[0:l] == "edit" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch r.Method {
								case "POST":
									s.handleEditUserRequest([0]string{}, elemIsEscaped, w, r)
								default:
									s.notAllowed(w, r, notAllowedParams{
										allowedMethods: "POST",
										allowedHeaders: rn68AllowedHeaders,
										acceptPost:     "multipart/form-data",
										acceptPatch:    "",
									})
								}

								return
							}

							elem = origElem
						case 'f': // Prefix: "footprints"
							origElem := elem
							if l := len("footprints"); len(elem) >= l && elem[0:l] == "footprints" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch r.Method {
								case "GET":
									s.handleGetFootprintsRequest([0]string{}, elemIsEscaped, w, r)
								default:
									s.notAllowed(w, r, notAllowedParams{
										allowedMethods: "GET",
										allowedHeaders: nil,
										acceptPost:     "",
										acceptPatch:    "",
									})
								}

								return
							}

							elem = origElem
						case 'l': // Prefix: "login_"
							origElem := elem
							if l := len("login_"); len(elem) >= l && elem[0:l] == "login_" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								break
							}
							switch elem[0] {
							case 'u': // Prefix: "update"

								if l := len("update"); len(elem) >= l && elem[0:l] == "update" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch r.Method {
									case "POST":
										s.handleUpdateLoginRequest([0]string{}, elemIsEscaped, w, r)
									default:
										s.notAllowed(w, r, notAllowedParams{
											allowedMethods: "POST",
											allowedHeaders: rn347AllowedHeaders,
											acceptPost:     "multipart/form-data",
											acceptPatch:    "",
										})
									}

									return
								}

							case 'w': // Prefix: "with_email"

								if l := len("with_email"); len(elem) >= l && elem[0:l] == "with_email" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch r.Method {
									case "POST":
										s.handleLoginWithEmailRequest([0]string{}, elemIsEscaped, w, r)
									default:
										s.notAllowed(w, r, notAllowedParams{
											allowedMethods: "POST",
											allowedHeaders: rn257AllowedHeaders,
											acceptPost:     "application/json",
											acceptPatch:    "",
										})
									}

									return
								}

							}

							elem = origElem
						}
						// Param: "id"
						// Match until "/"
						idx := strings.IndexByte(elem, '/')
						if idx < 0 {
							idx = len(elem)
						}
						args[0] = elem[:idx]
						elem = elem[idx:]

						if len(elem) == 0 {
							break
						}
						switch elem[0] {
						case '/': // Prefix: "/"

							if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								break
							}
							switch elem[0] {
							case 'f': // Prefix: "follow"

								if l := len("follow"); len(elem) >= l && elem[0:l] == "follow" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									break
								}
								switch elem[0] {
								case 'e': // Prefix: "ers"

									if l := len("ers"); len(elem) >= l && elem[0:l] == "ers" {
										elem = elem[l:]
									} else {
										break
									}

									if len(elem) == 0 {
										// Leaf node.
										switch r.Method {
										case "GET":
											s.handleGetUserFollowersRequest([1]string{
												args[0],
											}, elemIsEscaped, w, r)
										default:
											s.notAllowed(w, r, notAllowedParams{
												allowedMethods: "GET",
												allowedHeaders: nil,
												acceptPost:     "",
												acceptPatch:    "",
											})
										}

										return
									}

								case 'i': // Prefix: "ings"

									if l := len("ings"); len(elem) >= l && elem[0:l] == "ings" {
										elem = elem[l:]
									} else {
										break
									}

									if len(elem) == 0 {
										// Leaf node.
										switch r.Method {
										case "GET":
											s.handleGetUserFollowingsRequest([1]string{
												args[0],
											}, elemIsEscaped, w, r)
										default:
											s.notAllowed(w, r, notAllowedParams{
												allowedMethods: "GET",
												allowedHeaders: nil,
												acceptPost:     "",
												acceptPatch:    "",
											})
										}

										return
									}

								}

							case 'r': // Prefix: "report"

								if l := len("report"); len(elem) >= l && elem[0:l] == "report" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch r.Method {
									case "POST":
										s.handleReportUserRequest([1]string{
											args[0],
										}, elemIsEscaped, w, r)
									default:
										s.notAllowed(w, r, notAllowedParams{
											allowedMethods: "POST",
											allowedHeaders: rn302AllowedHeaders,
											acceptPost:     "multipart/form-data",
											acceptPatch:    "",
										})
									}

									return
								}

							}

						}

					}

				}

				elem = origElem
			}
			// Param: "countryApiValue"
			// Match until "/"
			idx := strings.IndexByte(elem, '/')
			if idx < 0 {
				idx = len(elem)
			}
			args[0] = elem[:idx]
			elem = elem[idx:]

			if len(elem) == 0 {
				break
			}
			switch elem[0] {
			case '/': // Prefix: "/api/"

				if l := len("/api/"); len(elem) >= l && elem[0:l] == "/api/" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'a': // Prefix: "apps/"

					if l := len("apps/"); len(elem) >= l && elem[0:l] == "apps/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "app"
					// Match until "/"
					idx := strings.IndexByte(elem, '/')
					if idx < 0 {
						idx = len(elem)
					}
					args[1] = elem[:idx]
					elem = elem[idx:]

					if len(elem) == 0 {
						break
					}
					switch elem[0] {
					case '/': // Prefix: "/popular_words"

						if l := len("/popular_words"); len(elem) >= l && elem[0:l] == "/popular_words" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch r.Method {
							case "GET":
								s.handleGetPopularWordsRequest([2]string{
									args[0],
									args[1],
								}, elemIsEscaped, w, r)
							default:
								s.notAllowed(w, r, notAllowedParams{
									allowedMethods: "GET",
									allowedHeaders: nil,
									acceptPost:     "",
									acceptPatch:    "",
								})
							}

							return
						}

					}

				case 'v': // Prefix: "v2/banned_words"

					if l := len("v2/banned_words"); len(elem) >= l && elem[0:l] == "v2/banned_words" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "GET":
							s.handleGetBannedWordsRequest([1]string{
								args[0],
							}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, notAllowedParams{
								allowedMethods: "GET",
								allowedHeaders: nil,
								acceptPost:     "",
								acceptPatch:    "",
							})
						}

						return
					}

				}

			}

		}
	}
	s.notFound(w, r)
}

// Route is route object.
type Route struct {
	name           string
	summary        string
	operationID    string
	operationGroup string
	pathPattern    string
	count          int
	args           [2]string
}

// Name returns ogen operation name.
//
// It is guaranteed to be unique and not empty.
func (r Route) Name() string {
	return r.name
}

// Summary returns OpenAPI summary.
func (r Route) Summary() string {
	return r.summary
}

// OperationID returns OpenAPI operationId.
func (r Route) OperationID() string {
	return r.operationID
}

// OperationGroup returns the x-ogen-operation-group value.
func (r Route) OperationGroup() string {
	return r.operationGroup
}

// PathPattern returns OpenAPI path.
func (r Route) PathPattern() string {
	return r.pathPattern
}

// Args returns parsed arguments.
func (r Route) Args() []string {
	return r.args[:r.count]
}

// FindRoute finds Route for given method and path.
//
// Note: this method does not unescape path or handle reserved characters in path properly. Use FindPath instead.
func (s *Server) FindRoute(method, path string) (Route, bool) {
	return s.FindPath(method, &url.URL{Path: path})
}

// FindPath finds Route for given method and URL.
func (s *Server) FindPath(method string, u *url.URL) (r Route, _ bool) {
	var (
		elem = u.Path
		args = r.args
	)
	if rawPath := u.RawPath; rawPath != "" {
		if normalized, ok := uri.NormalizeEscapedPath(rawPath); ok {
			elem = normalized
		}
		defer func() {
			for i, arg := range r.args[:r.count] {
				if unescaped, err := url.PathUnescape(arg); err == nil {
					r.args[i] = unescaped
				}
			}
		}()
	}

	elem, ok := s.cutPrefix(elem)
	if !ok {
		return r, false
	}

	// Static code generated router with unwrapped path search.
	switch {
	default:
		if len(elem) == 0 {
			break
		}
		switch elem[0] {
		case '/': // Prefix: "/"

			if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
				elem = elem[l:]
			} else {
				break
			}

			if len(elem) == 0 {
				break
			}
			switch elem[0] {
			case 'a': // Prefix: "api/"
				origElem := elem
				if l := len("api/"); len(elem) >= l && elem[0:l] == "api/" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'a': // Prefix: "apps/"

					if l := len("apps/"); len(elem) >= l && elem[0:l] == "apps/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "app"
					// Leaf parameter, slashes are prohibited
					idx := strings.IndexByte(elem, '/')
					if idx >= 0 {
						break
					}
					args[0] = elem
					elem = ""

					if len(elem) == 0 {
						// Leaf node.
						switch method {
						case "GET":
							r.name = GetAppConfigOperation
							r.summary = ""
							r.operationID = "getAppConfig"
							r.operationGroup = ""
							r.pathPattern = "/api/apps/{app}"
							r.args = args
							r.count = 1
							return r, true
						default:
							return
						}
					}

				case 'u': // Prefix: "user_activities"

					if l := len("user_activities"); len(elem) >= l && elem[0:l] == "user_activities" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch method {
						case "GET":
							r.name = GetUserActivitiesV1Operation
							r.summary = ""
							r.operationID = "getUserActivitiesV1"
							r.operationGroup = ""
							r.pathPattern = "/api/user_activities"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}

				case 'v': // Prefix: "v"

					if l := len("v"); len(elem) >= l && elem[0:l] == "v" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						break
					}
					switch elem[0] {
					case '1': // Prefix: "1/oauth/token"

						if l := len("1/oauth/token"); len(elem) >= l && elem[0:l] == "1/oauth/token" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch method {
							case "POST":
								r.name = OauthTokenOperation
								r.summary = ""
								r.operationID = "oauthToken"
								r.operationGroup = ""
								r.pathPattern = "/api/v1/oauth/token"
								r.args = args
								r.count = 0
								return r, true
							default:
								return
							}
						}

					case '2': // Prefix: "2/user_activities"

						if l := len("2/user_activities"); len(elem) >= l && elem[0:l] == "2/user_activities" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch method {
							case "GET":
								r.name = GetUserActivitiesOperation
								r.summary = ""
								r.operationID = "getUserActivities"
								r.operationGroup = ""
								r.pathPattern = "/api/v2/user_activities"
								r.args = args
								r.count = 0
								return r, true
							default:
								return
							}
						}

					}

				}

				elem = origElem
			case 'v': // Prefix: "v"
				origElem := elem
				if l := len("v"); len(elem) >= l && elem[0:l] == "v" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case '1': // Prefix: "1/"

					if l := len("1/"); len(elem) >= l && elem[0:l] == "1/" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						break
					}
					switch elem[0] {
					case 'b': // Prefix: "buckets/presigned_urls"

						if l := len("buckets/presigned_urls"); len(elem) >= l && elem[0:l] == "buckets/presigned_urls" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch method {
							case "GET":
								r.name = GetBucketPresignedUrlsOperation
								r.summary = ""
								r.operationID = "getBucketPresignedUrls"
								r.operationGroup = ""
								r.pathPattern = "/v1/buckets/presigned_urls"
								r.args = args
								r.count = 0
								return r, true
							default:
								return
							}
						}

					case 'c': // Prefix: "c"

						if l := len("c"); len(elem) >= l && elem[0:l] == "c" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							break
						}
						switch elem[0] {
						case 'a': // Prefix: "alls/"

							if l := len("alls/"); len(elem) >= l && elem[0:l] == "alls/" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								break
							}
							switch elem[0] {
							case 'a': // Prefix: "action_signature/generate"
								origElem := elem
								if l := len("action_signature/generate"); len(elem) >= l && elem[0:l] == "action_signature/generate" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch method {
									case "POST":
										r.name = GenerateCallActionSignatureOperation
										r.summary = ""
										r.operationID = "generateCallActionSignature"
										r.operationGroup = ""
										r.pathPattern = "/v1/calls/action_signature/generate"
										r.args = args
										r.count = 0
										return r, true
									default:
										return
									}
								}

								elem = origElem
							case 'b': // Prefix: "bgm"
								origElem := elem
								if l := len("bgm"); len(elem) >= l && elem[0:l] == "bgm" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch method {
									case "GET":
										r.name = GetCallBgmsOperation
										r.summary = ""
										r.operationID = "getCallBgms"
										r.operationGroup = ""
										r.pathPattern = "/v1/calls/bgm"
										r.args = args
										r.count = 0
										return r, true
									default:
										return
									}
								}

								elem = origElem
							case 'c': // Prefix: "conference_calls/"
								origElem := elem
								if l := len("conference_calls/"); len(elem) >= l && elem[0:l] == "conference_calls/" {
									elem = elem[l:]
								} else {
									break
								}

								// Param: "call_id"
								// Match until "/"
								idx := strings.IndexByte(elem, '/')
								if idx < 0 {
									idx = len(elem)
								}
								args[0] = elem[:idx]
								elem = elem[idx:]

								if len(elem) == 0 {
									break
								}
								switch elem[0] {
								case '/': // Prefix: "/invite"

									if l := len("/invite"); len(elem) >= l && elem[0:l] == "/invite" {
										elem = elem[l:]
									} else {
										break
									}

									if len(elem) == 0 {
										// Leaf node.
										switch method {
										case "POST":
											r.name = InviteToConferenceCallOperation
											r.summary = ""
											r.operationID = "inviteToConferenceCall"
											r.operationGroup = ""
											r.pathPattern = "/v1/calls/conference_calls/{call_id}/invite"
											r.args = args
											r.count = 1
											return r, true
										default:
											return
										}
									}

								}

								elem = origElem
							case 'l': // Prefix: "leave_"
								origElem := elem
								if l := len("leave_"); len(elem) >= l && elem[0:l] == "leave_" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									break
								}
								switch elem[0] {
								case 'a': // Prefix: "agora_channel"

									if l := len("agora_channel"); len(elem) >= l && elem[0:l] == "agora_channel" {
										elem = elem[l:]
									} else {
										break
									}

									if len(elem) == 0 {
										// Leaf node.
										switch method {
										case "POST":
											r.name = LeaveAgoraChannelOperation
											r.summary = ""
											r.operationID = "leaveAgoraChannel"
											r.operationGroup = ""
											r.pathPattern = "/v1/calls/leave_agora_channel"
											r.args = args
											r.count = 0
											return r, true
										default:
											return
										}
									}

								case 'c': // Prefix: "conference_call"

									if l := len("conference_call"); len(elem) >= l && elem[0:l] == "conference_call" {
										elem = elem[l:]
									} else {
										break
									}

									if len(elem) == 0 {
										// Leaf node.
										switch method {
										case "POST":
											r.name = LeaveConferenceCallOperation
											r.summary = ""
											r.operationID = "leaveConferenceCall"
											r.operationGroup = ""
											r.pathPattern = "/v1/calls/leave_conference_call"
											r.args = args
											r.count = 0
											return r, true
										default:
											return
										}
									}

								}

								elem = origElem
							case 'p': // Prefix: "phone_status/"
								origElem := elem
								if l := len("phone_status/"); len(elem) >= l && elem[0:l] == "phone_status/" {
									elem = elem[l:]
								} else {
									break
								}

								// Param: "opponent_id"
								// Leaf parameter, slashes are prohibited
								idx := strings.IndexByte(elem, '/')
								if idx >= 0 {
									break
								}
								args[0] = elem
								elem = ""

								if len(elem) == 0 {
									// Leaf node.
									switch method {
									case "GET":
										r.name = GetPhoneStatusOperation
										r.summary = ""
										r.operationID = "getPhoneStatus"
										r.operationGroup = ""
										r.pathPattern = "/v1/calls/phone_status/{opponent_id}"
										r.args = args
										r.count = 1
										return r, true
									default:
										return
									}
								}

								elem = origElem
							}
							// Param: "call_id"
							// Match until "/"
							idx := strings.IndexByte(elem, '/')
							if idx < 0 {
								idx = len(elem)
							}
							args[0] = elem[:idx]
							elem = elem[idx:]

							if len(elem) == 0 {
								switch method {
								case "PUT":
									r.name = UpdateCallOperation
									r.summary = ""
									r.operationID = "updateCall"
									r.operationGroup = ""
									r.pathPattern = "/v1/calls/{call_id}"
									r.args = args
									r.count = 1
									return r, true
								default:
									return
								}
							}
							switch elem[0] {
							case '/': // Prefix: "/"

								if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									break
								}
								switch elem[0] {
								case 'b': // Prefix: "bu"

									if l := len("bu"); len(elem) >= l && elem[0:l] == "bu" {
										elem = elem[l:]
									} else {
										break
									}

									if len(elem) == 0 {
										break
									}
									switch elem[0] {
									case 'l': // Prefix: "lk_invite"

										if l := len("lk_invite"); len(elem) >= l && elem[0:l] == "lk_invite" {
											elem = elem[l:]
										} else {
											break
										}

										if len(elem) == 0 {
											// Leaf node.
											switch method {
											case "POST":
												r.name = BulkInviteToCallOperation
												r.summary = ""
												r.operationID = "bulkInviteToCall"
												r.operationGroup = ""
												r.pathPattern = "/v1/calls/{call_id}/bulk_invite"
												r.args = args
												r.count = 1
												return r, true
											default:
												return
											}
										}

									case 'm': // Prefix: "mp"

										if l := len("mp"); len(elem) >= l && elem[0:l] == "mp" {
											elem = elem[l:]
										} else {
											break
										}

										if len(elem) == 0 {
											// Leaf node.
											switch method {
											case "POST":
												r.name = BumpCallOperation
												r.summary = ""
												r.operationID = "bumpCall"
												r.operationGroup = ""
												r.pathPattern = "/v1/calls/{call_id}/bump"
												r.args = args
												r.count = 1
												return r, true
											default:
												return
											}
										}

									}

								case 'g': // Prefix: "gift_transactions"

									if l := len("gift_transactions"); len(elem) >= l && elem[0:l] == "gift_transactions" {
										elem = elem[l:]
									} else {
										break
									}

									if len(elem) == 0 {
										// Leaf node.
										switch method {
										case "GET":
											r.name = GetCallGiftHistoryOperation
											r.summary = ""
											r.operationID = "getCallGiftHistory"
											r.operationGroup = ""
											r.pathPattern = "/v1/calls/{call_id}/gift_transactions"
											r.args = args
											r.count = 1
											return r, true
										default:
											return
										}
									}

								case 'u': // Prefix: "users/"

									if l := len("users/"); len(elem) >= l && elem[0:l] == "users/" {
										elem = elem[l:]
									} else {
										break
									}

									if len(elem) == 0 {
										break
									}
									switch elem[0] {
									case 'i': // Prefix: "invitable"
										origElem := elem
										if l := len("invitable"); len(elem) >= l && elem[0:l] == "invitable" {
											elem = elem[l:]
										} else {
											break
										}

										if len(elem) == 0 {
											// Leaf node.
											switch method {
											case "GET":
												r.name = GetInvitableCallUsersOperation
												r.summary = ""
												r.operationID = "getInvitableCallUsers"
												r.operationGroup = ""
												r.pathPattern = "/v1/calls/{call_id}/users/invitable"
												r.args = args
												r.count = 1
												return r, true
											default:
												return
											}
										}

										elem = origElem
									}
									// Param: "user_id"
									// Leaf parameter, slashes are prohibited
									idx := strings.IndexByte(elem, '/')
									if idx >= 0 {
										break
									}
									args[1] = elem
									elem = ""

									if len(elem) == 0 {
										// Leaf node.
										switch method {
										case "PUT":
											r.name = UpdateCallUserOperation
											r.summary = ""
											r.operationID = "updateCallUser"
											r.operationGroup = ""
											r.pathPattern = "/v1/calls/{call_id}/users/{user_id}"
											r.args = args
											r.count = 2
											return r, true
										default:
											return
										}
									}

								}

							}

						case 'h': // Prefix: "hat_rooms/"

							if l := len("hat_rooms/"); len(elem) >= l && elem[0:l] == "hat_rooms/" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								break
							}
							switch elem[0] {
							case 'a': // Prefix: "accept_chat_request"
								origElem := elem
								if l := len("accept_chat_request"); len(elem) >= l && elem[0:l] == "accept_chat_request" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch method {
									case "POST":
										r.name = AcceptChatRequestOperation
										r.summary = ""
										r.operationID = "acceptChatRequest"
										r.operationGroup = ""
										r.pathPattern = "/v1/chat_rooms/accept_chat_request"
										r.args = args
										r.count = 0
										return r, true
									default:
										return
									}
								}

								elem = origElem
							case 'm': // Prefix: "ma"
								origElem := elem
								if l := len("ma"); len(elem) >= l && elem[0:l] == "ma" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									break
								}
								switch elem[0] {
								case 'i': // Prefix: "in_list"

									if l := len("in_list"); len(elem) >= l && elem[0:l] == "in_list" {
										elem = elem[l:]
									} else {
										break
									}

									if len(elem) == 0 {
										// Leaf node.
										switch method {
										case "GET":
											r.name = GetMainChatRoomsOperation
											r.summary = ""
											r.operationID = "getMainChatRooms"
											r.operationGroup = ""
											r.pathPattern = "/v1/chat_rooms/main_list"
											r.args = args
											r.count = 0
											return r, true
										default:
											return
										}
									}

								case 's': // Prefix: "ss_destroy"

									if l := len("ss_destroy"); len(elem) >= l && elem[0:l] == "ss_destroy" {
										elem = elem[l:]
									} else {
										break
									}

									if len(elem) == 0 {
										// Leaf node.
										switch method {
										case "POST":
											r.name = DeleteChatRoomsOperation
											r.summary = ""
											r.operationID = "deleteChatRooms"
											r.operationGroup = ""
											r.pathPattern = "/v1/chat_rooms/mass_destroy"
											r.args = args
											r.count = 0
											return r, true
										default:
											return
										}
									}

								}

								elem = origElem
							case 'n': // Prefix: "new"
								origElem := elem
								if l := len("new"); len(elem) >= l && elem[0:l] == "new" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch method {
									case "POST":
										r.name = CreateChatRoomV1Operation
										r.summary = ""
										r.operationID = "createChatRoomV1"
										r.operationGroup = ""
										r.pathPattern = "/v1/chat_rooms/new"
										r.args = args
										r.count = 0
										return r, true
									default:
										return
									}
								}

								elem = origElem
							case 'r': // Prefix: "request_list"
								origElem := elem
								if l := len("request_list"); len(elem) >= l && elem[0:l] == "request_list" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch method {
									case "GET":
										r.name = GetChatRequestsOperation
										r.summary = ""
										r.operationID = "getChatRequests"
										r.operationGroup = ""
										r.pathPattern = "/v1/chat_rooms/request_list"
										r.args = args
										r.count = 0
										return r, true
									default:
										return
									}
								}

								elem = origElem
							case 't': // Prefix: "total_chat_request"
								origElem := elem
								if l := len("total_chat_request"); len(elem) >= l && elem[0:l] == "total_chat_request" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch method {
									case "GET":
										r.name = GetChatRequestCountOperation
										r.summary = ""
										r.operationID = "getChatRequestCount"
										r.operationGroup = ""
										r.pathPattern = "/v1/chat_rooms/total_chat_request"
										r.args = args
										r.count = 0
										return r, true
									default:
										return
									}
								}

								elem = origElem
							case 'u': // Prefix: "unread_status"
								origElem := elem
								if l := len("unread_status"); len(elem) >= l && elem[0:l] == "unread_status" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch method {
									case "GET":
										r.name = GetChatUnreadStatusOperation
										r.summary = ""
										r.operationID = "getChatUnreadStatus"
										r.operationGroup = ""
										r.pathPattern = "/v1/chat_rooms/unread_status"
										r.args = args
										r.count = 0
										return r, true
									default:
										return
									}
								}

								elem = origElem
							}
							// Param: "room_id"
							// Match until "/"
							idx := strings.IndexByte(elem, '/')
							if idx < 0 {
								idx = len(elem)
							}
							args[0] = elem[:idx]
							elem = elem[idx:]

							if len(elem) == 0 {
								break
							}
							switch elem[0] {
							case '/': // Prefix: "/"

								if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									break
								}
								switch elem[0] {
								case 'a': // Prefix: "attachments_read"

									if l := len("attachments_read"); len(elem) >= l && elem[0:l] == "attachments_read" {
										elem = elem[l:]
									} else {
										break
									}

									if len(elem) == 0 {
										// Leaf node.
										switch method {
										case "POST":
											r.name = ReadChatAttachmentsOperation
											r.summary = ""
											r.operationID = "readChatAttachments"
											r.operationGroup = ""
											r.pathPattern = "/v1/chat_rooms/{id}/attachments_read"
											r.args = args
											r.count = 1
											return r, true
										default:
											return
										}
									}

								case 'm': // Prefix: "messages/"

									if l := len("messages/"); len(elem) >= l && elem[0:l] == "messages/" {
										elem = elem[l:]
									} else {
										break
									}

									// Param: "message_id"
									// Match until "/"
									idx := strings.IndexByte(elem, '/')
									if idx < 0 {
										idx = len(elem)
									}
									args[1] = elem[:idx]
									elem = elem[idx:]

									if len(elem) == 0 {
										break
									}
									switch elem[0] {
									case '/': // Prefix: "/delete"

										if l := len("/delete"); len(elem) >= l && elem[0:l] == "/delete" {
											elem = elem[l:]
										} else {
											break
										}

										if len(elem) == 0 {
											// Leaf node.
											switch method {
											case "DELETE":
												r.name = DeleteChatMessageOperation
												r.summary = ""
												r.operationID = "deleteChatMessage"
												r.operationGroup = ""
												r.pathPattern = "/v1/chat_rooms/{room_id}/messages/{message_id}/delete"
												r.args = args
												r.count = 2
												return r, true
											default:
												return
											}
										}

									}

								case 'p': // Prefix: "pinned"

									if l := len("pinned"); len(elem) >= l && elem[0:l] == "pinned" {
										elem = elem[l:]
									} else {
										break
									}

									if len(elem) == 0 {
										// Leaf node.
										switch method {
										case "DELETE":
											r.name = UnpinChatRoomOperation
											r.summary = ""
											r.operationID = "unpinChatRoom"
											r.operationGroup = ""
											r.pathPattern = "/v1/chat_rooms/{id}/pinned"
											r.args = args
											r.count = 1
											return r, true
										case "POST":
											r.name = PinChatRoomOperation
											r.summary = ""
											r.operationID = "pinChatRoom"
											r.operationGroup = ""
											r.pathPattern = "/v1/chat_rooms/{id}/pinned"
											r.args = args
											r.count = 1
											return r, true
										default:
											return
										}
									}

								case 'v': // Prefix: "videos_read"

									if l := len("videos_read"); len(elem) >= l && elem[0:l] == "videos_read" {
										elem = elem[l:]
									} else {
										break
									}

									if len(elem) == 0 {
										// Leaf node.
										switch method {
										case "POST":
											r.name = ReadChatVideosOperation
											r.summary = ""
											r.operationID = "readChatVideos"
											r.operationGroup = ""
											r.pathPattern = "/v1/chat_rooms/{id}/videos_read"
											r.args = args
											r.count = 1
											return r, true
										default:
											return
										}
									}

								}

							}

						}

					case 'e': // Prefix: "email_verification_urls"

						if l := len("email_verification_urls"); len(elem) >= l && elem[0:l] == "email_verification_urls" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch method {
							case "POST":
								r.name = RequestEmailVerificationOperation
								r.summary = ""
								r.operationID = "requestEmailVerification"
								r.operationGroup = ""
								r.pathPattern = "/v1/email_verification_urls"
								r.args = args
								r.count = 0
								return r, true
							default:
								return
							}
						}

					case 'g': // Prefix: "g"

						if l := len("g"); len(elem) >= l && elem[0:l] == "g" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							break
						}
						switch elem[0] {
						case 'a': // Prefix: "ames/apps"

							if l := len("ames/apps"); len(elem) >= l && elem[0:l] == "ames/apps" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								switch method {
								case "GET":
									r.name = ListGameAppsOperation
									r.summary = ""
									r.operationID = "listGameApps"
									r.operationGroup = ""
									r.pathPattern = "/v1/games/apps"
									r.args = args
									r.count = 0
									return r, true
								default:
									return
								}
							}
							switch elem[0] {
							case '/': // Prefix: "/"

								if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
									elem = elem[l:]
								} else {
									break
								}

								// Param: "app_id"
								// Match until "/"
								idx := strings.IndexByte(elem, '/')
								if idx < 0 {
									idx = len(elem)
								}
								args[0] = elem[:idx]
								elem = elem[idx:]

								if len(elem) == 0 {
									break
								}
								switch elem[0] {
								case '/': // Prefix: "/walkthroughs"

									if l := len("/walkthroughs"); len(elem) >= l && elem[0:l] == "/walkthroughs" {
										elem = elem[l:]
									} else {
										break
									}

									if len(elem) == 0 {
										// Leaf node.
										switch method {
										case "GET":
											r.name = ListGameWalkthroughsOperation
											r.summary = ""
											r.operationID = "listGameWalkthroughs"
											r.operationGroup = ""
											r.pathPattern = "/v1/games/apps/{app_id}/walkthroughs"
											r.args = args
											r.count = 1
											return r, true
										default:
											return
										}
									}

								}

							}

						case 'e': // Prefix: "enres"

							if l := len("enres"); len(elem) >= l && elem[0:l] == "enres" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch method {
								case "GET":
									r.name = ListGenresOperation
									r.summary = ""
									r.operationID = "listGenres"
									r.operationGroup = ""
									r.pathPattern = "/v1/genres"
									r.args = args
									r.count = 0
									return r, true
								default:
									return
								}
							}

						case 'r': // Prefix: "roup"

							if l := len("roup"); len(elem) >= l && elem[0:l] == "roup" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								break
							}
							switch elem[0] {
							case '_': // Prefix: "_mute/"

								if l := len("_mute/"); len(elem) >= l && elem[0:l] == "_mute/" {
									elem = elem[l:]
								} else {
									break
								}

								// Param: "id"
								// Match until "/"
								idx := strings.IndexByte(elem, '/')
								if idx < 0 {
									idx = len(elem)
								}
								args[0] = elem[:idx]
								elem = elem[idx:]

								if len(elem) == 0 {
									break
								}
								switch elem[0] {
								case '/': // Prefix: "/"

									if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
										elem = elem[l:]
									} else {
										break
									}

									if len(elem) == 0 {
										break
									}
									switch elem[0] {
									case 'm': // Prefix: "mute"

										if l := len("mute"); len(elem) >= l && elem[0:l] == "mute" {
											elem = elem[l:]
										} else {
											break
										}

										if len(elem) == 0 {
											break
										}
										switch elem[0] {
										case '/': // Prefix: "/"

											if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
												elem = elem[l:]
											} else {
												break
											}

											// Param: "user_id"
											// Leaf parameter, slashes are prohibited
											idx := strings.IndexByte(elem, '/')
											if idx >= 0 {
												break
											}
											args[1] = elem
											elem = ""

											if len(elem) == 0 {
												// Leaf node.
												switch method {
												case "POST":
													r.name = MuteGroupUserOperation
													r.summary = ""
													r.operationID = "muteGroupUser"
													r.operationGroup = ""
													r.pathPattern = "/v1/group_mute/{id}/mute/{user_id}"
													r.args = args
													r.count = 2
													return r, true
												default:
													return
												}
											}

										case 'd': // Prefix: "d_users"

											if l := len("d_users"); len(elem) >= l && elem[0:l] == "d_users" {
												elem = elem[l:]
											} else {
												break
											}

											if len(elem) == 0 {
												// Leaf node.
												switch method {
												case "GET":
													r.name = ListMutedGroupUsersOperation
													r.summary = ""
													r.operationID = "listMutedGroupUsers"
													r.operationGroup = ""
													r.pathPattern = "/v1/group_mute/{id}/muted_users"
													r.args = args
													r.count = 1
													return r, true
												default:
													return
												}
											}

										}

									case 'u': // Prefix: "unmute/"

										if l := len("unmute/"); len(elem) >= l && elem[0:l] == "unmute/" {
											elem = elem[l:]
										} else {
											break
										}

										// Param: "user_id"
										// Leaf parameter, slashes are prohibited
										idx := strings.IndexByte(elem, '/')
										if idx >= 0 {
											break
										}
										args[1] = elem
										elem = ""

										if len(elem) == 0 {
											// Leaf node.
											switch method {
											case "DELETE":
												r.name = UnmuteGroupUserOperation
												r.summary = ""
												r.operationID = "unmuteGroupUser"
												r.operationGroup = ""
												r.pathPattern = "/v1/group_mute/{id}/unmute/{user_id}"
												r.args = args
												r.count = 2
												return r, true
											default:
												return
											}
										}

									}

								}

							case 's': // Prefix: "s/"

								if l := len("s/"); len(elem) >= l && elem[0:l] == "s/" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									break
								}
								switch elem[0] {
								case 'c': // Prefix: "c"
									origElem := elem
									if l := len("c"); len(elem) >= l && elem[0:l] == "c" {
										elem = elem[l:]
									} else {
										break
									}

									if len(elem) == 0 {
										break
									}
									switch elem[0] {
									case 'a': // Prefix: "ategories"

										if l := len("ategories"); len(elem) >= l && elem[0:l] == "ategories" {
											elem = elem[l:]
										} else {
											break
										}

										if len(elem) == 0 {
											// Leaf node.
											switch method {
											case "GET":
												r.name = ListGroupCategoriesOperation
												r.summary = ""
												r.operationID = "listGroupCategories"
												r.operationGroup = ""
												r.pathPattern = "/v1/groups/categories"
												r.args = args
												r.count = 0
												return r, true
											default:
												return
											}
										}

									case 'r': // Prefix: "reated_quota"

										if l := len("reated_quota"); len(elem) >= l && elem[0:l] == "reated_quota" {
											elem = elem[l:]
										} else {
											break
										}

										if len(elem) == 0 {
											// Leaf node.
											switch method {
											case "GET":
												r.name = GetGroupCreateQuotaOperation
												r.summary = ""
												r.operationID = "getGroupCreateQuota"
												r.operationGroup = ""
												r.pathPattern = "/v1/groups/created_quota"
												r.args = args
												r.count = 0
												return r, true
											default:
												return
											}
										}

									}

									elem = origElem
								case 'j': // Prefix: "joined_statuses"
									origElem := elem
									if l := len("joined_statuses"); len(elem) >= l && elem[0:l] == "joined_statuses" {
										elem = elem[l:]
									} else {
										break
									}

									if len(elem) == 0 {
										// Leaf node.
										switch method {
										case "GET":
											r.name = GetJoinedGroupStatusesOperation
											r.summary = ""
											r.operationID = "getJoinedGroupStatuses"
											r.operationGroup = ""
											r.pathPattern = "/v1/groups/joined_statuses"
											r.args = args
											r.count = 0
											return r, true
										default:
											return
										}
									}

									elem = origElem
								case 'u': // Prefix: "u"
									origElem := elem
									if l := len("u"); len(elem) >= l && elem[0:l] == "u" {
										elem = elem[l:]
									} else {
										break
									}

									if len(elem) == 0 {
										break
									}
									switch elem[0] {
									case 'n': // Prefix: "nread_status"

										if l := len("nread_status"); len(elem) >= l && elem[0:l] == "nread_status" {
											elem = elem[l:]
										} else {
											break
										}

										if len(elem) == 0 {
											// Leaf node.
											switch method {
											case "GET":
												r.name = GetGroupUnreadStatusOperation
												r.summary = ""
												r.operationID = "getGroupUnreadStatus"
												r.operationGroup = ""
												r.pathPattern = "/v1/groups/unread_status"
												r.args = args
												r.count = 0
												return r, true
											default:
												return
											}
										}

									case 's': // Prefix: "ser_group_list"

										if l := len("ser_group_list"); len(elem) >= l && elem[0:l] == "ser_group_list" {
											elem = elem[l:]
										} else {
											break
										}

										if len(elem) == 0 {
											// Leaf node.
											switch method {
											case "GET":
												r.name = GetUserGroupListOperation
												r.summary = ""
												r.operationID = "getUserGroupList"
												r.operationGroup = ""
												r.pathPattern = "/v1/groups/user_group_list"
												r.args = args
												r.count = 0
												return r, true
											default:
												return
											}
										}

									}

									elem = origElem
								}
								// Param: "id"
								// Match until "/"
								idx := strings.IndexByte(elem, '/')
								if idx < 0 {
									idx = len(elem)
								}
								args[0] = elem[:idx]
								elem = elem[idx:]

								if len(elem) == 0 {
									switch method {
									case "GET":
										r.name = GetGroupOperation
										r.summary = ""
										r.operationID = "getGroup"
										r.operationGroup = ""
										r.pathPattern = "/v1/groups/{id}"
										r.args = args
										r.count = 1
										return r, true
									default:
										return
									}
								}
								switch elem[0] {
								case '/': // Prefix: "/"

									if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
										elem = elem[l:]
									} else {
										break
									}

									if len(elem) == 0 {
										break
									}
									switch elem[0] {
									case 'a': // Prefix: "accept/"

										if l := len("accept/"); len(elem) >= l && elem[0:l] == "accept/" {
											elem = elem[l:]
										} else {
											break
										}

										// Param: "userId"
										// Leaf parameter, slashes are prohibited
										idx := strings.IndexByte(elem, '/')
										if idx >= 0 {
											break
										}
										args[1] = elem
										elem = ""

										if len(elem) == 0 {
											// Leaf node.
											switch method {
											case "POST":
												r.name = AcceptGroupJoinRequestOperation
												r.summary = ""
												r.operationID = "acceptGroupJoinRequest"
												r.operationGroup = ""
												r.pathPattern = "/v1/groups/{id}/accept/{userId}"
												r.args = args
												r.count = 2
												return r, true
											default:
												return
											}
										}

									case 'b': // Prefix: "ban"

										if l := len("ban"); len(elem) >= l && elem[0:l] == "ban" {
											elem = elem[l:]
										} else {
											break
										}

										if len(elem) == 0 {
											break
										}
										switch elem[0] {
										case '/': // Prefix: "/"

											if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
												elem = elem[l:]
											} else {
												break
											}

											// Param: "userId"
											// Leaf parameter, slashes are prohibited
											idx := strings.IndexByte(elem, '/')
											if idx >= 0 {
												break
											}
											args[1] = elem
											elem = ""

											if len(elem) == 0 {
												// Leaf node.
												switch method {
												case "POST":
													r.name = BanGroupUserOperation
													r.summary = ""
													r.operationID = "banGroupUser"
													r.operationGroup = ""
													r.pathPattern = "/v1/groups/{id}/ban/{userId}"
													r.args = args
													r.count = 2
													return r, true
												default:
													return
												}
											}

										case '_': // Prefix: "_list"

											if l := len("_list"); len(elem) >= l && elem[0:l] == "_list" {
												elem = elem[l:]
											} else {
												break
											}

											if len(elem) == 0 {
												// Leaf node.
												switch method {
												case "GET":
													r.name = GetGroupBanListOperation
													r.summary = ""
													r.operationID = "getGroupBanList"
													r.operationGroup = ""
													r.pathPattern = "/v1/groups/{id}/ban_list"
													r.args = args
													r.count = 1
													return r, true
												default:
													return
												}
											}

										}

									case 'd': // Prefix: "de"

										if l := len("de"); len(elem) >= l && elem[0:l] == "de" {
											elem = elem[l:]
										} else {
											break
										}

										if len(elem) == 0 {
											break
										}
										switch elem[0] {
										case 'c': // Prefix: "cline/"

											if l := len("cline/"); len(elem) >= l && elem[0:l] == "cline/" {
												elem = elem[l:]
											} else {
												break
											}

											// Param: "userId"
											// Leaf parameter, slashes are prohibited
											idx := strings.IndexByte(elem, '/')
											if idx >= 0 {
												break
											}
											args[1] = elem
											elem = ""

											if len(elem) == 0 {
												// Leaf node.
												switch method {
												case "POST":
													r.name = DeclineGroupJoinRequestOperation
													r.summary = ""
													r.operationID = "declineGroupJoinRequest"
													r.operationGroup = ""
													r.pathPattern = "/v1/groups/{id}/decline/{userId}"
													r.args = args
													r.count = 2
													return r, true
												default:
													return
												}
											}

										case 'p': // Prefix: "putize"

											if l := len("putize"); len(elem) >= l && elem[0:l] == "putize" {
												elem = elem[l:]
											} else {
												break
											}

											if len(elem) == 0 {
												switch method {
												case "DELETE":
													r.name = RemoveGroupDeputiesOperation
													r.summary = ""
													r.operationID = "removeGroupDeputies"
													r.operationGroup = ""
													r.pathPattern = "/v1/groups/{id}/deputize"
													r.args = args
													r.count = 1
													return r, true
												case "PUT":
													r.name = DeputizeGroupUsersOperation
													r.summary = ""
													r.operationID = "deputizeGroupUsers"
													r.operationGroup = ""
													r.pathPattern = "/v1/groups/{id}/deputize"
													r.args = args
													r.count = 1
													return r, true
												default:
													return
												}
											}
											switch elem[0] {
											case '/': // Prefix: "/"

												if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
													elem = elem[l:]
												} else {
													break
												}

												// Param: "user_id"
												// Match until "/"
												idx := strings.IndexByte(elem, '/')
												if idx < 0 {
													idx = len(elem)
												}
												args[1] = elem[:idx]
												elem = elem[idx:]

												if len(elem) == 0 {
													break
												}
												switch elem[0] {
												case '/': // Prefix: "/withdraw"

													if l := len("/withdraw"); len(elem) >= l && elem[0:l] == "/withdraw" {
														elem = elem[l:]
													} else {
														break
													}

													if len(elem) == 0 {
														// Leaf node.
														switch method {
														case "PUT":
															r.name = WithdrawGroupDeputyOperation
															r.summary = ""
															r.operationID = "withdrawGroupDeputy"
															r.operationGroup = ""
															r.pathPattern = "/v1/groups/{group_id}/deputize/{user_id}/withdraw"
															r.args = args
															r.count = 2
															return r, true
														default:
															return
														}
													}

												}

											}

										}

									case 'f': // Prefix: "fire/"

										if l := len("fire/"); len(elem) >= l && elem[0:l] == "fire/" {
											elem = elem[l:]
										} else {
											break
										}

										// Param: "user_id"
										// Leaf parameter, slashes are prohibited
										idx := strings.IndexByte(elem, '/')
										if idx >= 0 {
											break
										}
										args[1] = elem
										elem = ""

										if len(elem) == 0 {
											// Leaf node.
											switch method {
											case "POST":
												r.name = FireGroupUserOperation
												r.summary = ""
												r.operationID = "fireGroupUser"
												r.operationGroup = ""
												r.pathPattern = "/v1/groups/{group_id}/fire/{user_id}"
												r.args = args
												r.count = 2
												return r, true
											default:
												return
											}
										}

									case 'g': // Prefix: "gift_"

										if l := len("gift_"); len(elem) >= l && elem[0:l] == "gift_" {
											elem = elem[l:]
										} else {
											break
										}

										if len(elem) == 0 {
											break
										}
										switch elem[0] {
										case 'h': // Prefix: "history"

											if l := len("history"); len(elem) >= l && elem[0:l] == "history" {
												elem = elem[l:]
											} else {
												break
											}

											if len(elem) == 0 {
												// Leaf node.
												switch method {
												case "GET":
													r.name = GetGroupGiftHistoryOperation
													r.summary = ""
													r.operationID = "getGroupGiftHistory"
													r.operationGroup = ""
													r.pathPattern = "/v1/groups/{group_id}/gift_history"
													r.args = args
													r.count = 1
													return r, true
												default:
													return
												}
											}

										case 't': // Prefix: "transactions"

											if l := len("transactions"); len(elem) >= l && elem[0:l] == "transactions" {
												elem = elem[l:]
											} else {
												break
											}

											if len(elem) == 0 {
												// Leaf node.
												switch method {
												case "GET":
													r.name = GetGroupGiftTransactionsOperation
													r.summary = ""
													r.operationID = "getGroupGiftTransactions"
													r.operationGroup = ""
													r.pathPattern = "/v1/groups/{group_id}/gift_transactions"
													r.args = args
													r.count = 1
													return r, true
												default:
													return
												}
											}

										}

									case 'h': // Prefix: "highlights"

										if l := len("highlights"); len(elem) >= l && elem[0:l] == "highlights" {
											elem = elem[l:]
										} else {
											break
										}

										if len(elem) == 0 {
											switch method {
											case "GET":
												r.name = GetGroupHighlightsOperation
												r.summary = ""
												r.operationID = "getGroupHighlights"
												r.operationGroup = ""
												r.pathPattern = "/v1/groups/{group_id}/highlights"
												r.args = args
												r.count = 1
												return r, true
											default:
												return
											}
										}
										switch elem[0] {
										case '/': // Prefix: "/"

											if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
												elem = elem[l:]
											} else {
												break
											}

											// Param: "post_id"
											// Leaf parameter, slashes are prohibited
											idx := strings.IndexByte(elem, '/')
											if idx >= 0 {
												break
											}
											args[1] = elem
											elem = ""

											if len(elem) == 0 {
												// Leaf node.
												switch method {
												case "DELETE":
													r.name = UnpinGroupHighlightPostOperation
													r.summary = ""
													r.operationID = "unpinGroupHighlightPost"
													r.operationGroup = ""
													r.pathPattern = "/v1/groups/{group_id}/highlights/{post_id}"
													r.args = args
													r.count = 2
													return r, true
												case "PUT":
													r.name = PinGroupHighlightPostOperation
													r.summary = ""
													r.operationID = "pinGroupHighlightPost"
													r.operationGroup = ""
													r.pathPattern = "/v1/groups/{group_id}/highlights/{post_id}"
													r.args = args
													r.count = 2
													return r, true
												default:
													return
												}
											}

										}

									case 'i': // Prefix: "invite"

										if l := len("invite"); len(elem) >= l && elem[0:l] == "invite" {
											elem = elem[l:]
										} else {
											break
										}

										if len(elem) == 0 {
											// Leaf node.
											switch method {
											case "POST":
												r.name = InviteToGroupOperation
												r.summary = ""
												r.operationID = "inviteToGroup"
												r.operationGroup = ""
												r.pathPattern = "/v1/groups/{id}/invite"
												r.args = args
												r.count = 1
												return r, true
											default:
												return
											}
										}

									case 'j': // Prefix: "join"

										if l := len("join"); len(elem) >= l && elem[0:l] == "join" {
											elem = elem[l:]
										} else {
											break
										}

										if len(elem) == 0 {
											// Leaf node.
											switch method {
											case "POST":
												r.name = JoinGroupOperation
												r.summary = ""
												r.operationID = "joinGroup"
												r.operationGroup = ""
												r.pathPattern = "/v1/groups/{id}/join"
												r.args = args
												r.count = 1
												return r, true
											default:
												return
											}
										}

									case 'l': // Prefix: "leave"

										if l := len("leave"); len(elem) >= l && elem[0:l] == "leave" {
											elem = elem[l:]
										} else {
											break
										}

										if len(elem) == 0 {
											// Leaf node.
											switch method {
											case "DELETE":
												r.name = LeaveGroupOperation
												r.summary = ""
												r.operationID = "leaveGroup"
												r.operationGroup = ""
												r.pathPattern = "/v1/groups/{id}/leave"
												r.args = args
												r.count = 1
												return r, true
											default:
												return
											}
										}

									case 'm': // Prefix: "members/"

										if l := len("members/"); len(elem) >= l && elem[0:l] == "members/" {
											elem = elem[l:]
										} else {
											break
										}

										// Param: "userId"
										// Leaf parameter, slashes are prohibited
										idx := strings.IndexByte(elem, '/')
										if idx >= 0 {
											break
										}
										args[1] = elem
										elem = ""

										if len(elem) == 0 {
											// Leaf node.
											switch method {
											case "GET":
												r.name = GetGroupMemberOperation
												r.summary = ""
												r.operationID = "getGroupMember"
												r.operationGroup = ""
												r.pathPattern = "/v1/groups/{id}/members/{userId}"
												r.args = args
												r.count = 2
												return r, true
											default:
												return
											}
										}

									case 'r': // Prefix: "re"

										if l := len("re"); len(elem) >= l && elem[0:l] == "re" {
											elem = elem[l:]
										} else {
											break
										}

										if len(elem) == 0 {
											break
										}
										switch elem[0] {
										case 'c': // Prefix: "ceived_gifts/"

											if l := len("ceived_gifts/"); len(elem) >= l && elem[0:l] == "ceived_gifts/" {
												elem = elem[l:]
											} else {
												break
											}

											// Param: "gift_id"
											// Match until "/"
											idx := strings.IndexByte(elem, '/')
											if idx < 0 {
												idx = len(elem)
											}
											args[1] = elem[:idx]
											elem = elem[idx:]

											if len(elem) == 0 {
												break
											}
											switch elem[0] {
											case '/': // Prefix: "/senders"

												if l := len("/senders"); len(elem) >= l && elem[0:l] == "/senders" {
													elem = elem[l:]
												} else {
													break
												}

												if len(elem) == 0 {
													// Leaf node.
													switch method {
													case "GET":
														r.name = GetGroupReceivedGiftSendersOperation
														r.summary = ""
														r.operationID = "getGroupReceivedGiftSenders"
														r.operationGroup = ""
														r.pathPattern = "/v1/groups/{group_id}/received_gifts/{gift_id}/senders"
														r.args = args
														r.count = 2
														return r, true
													default:
														return
													}
												}

											}

										case 'l': // Prefix: "lat"

											if l := len("lat"); len(elem) >= l && elem[0:l] == "lat" {
												elem = elem[l:]
											} else {
												break
											}

											if len(elem) == 0 {
												break
											}
											switch elem[0] {
											case 'a': // Prefix: "able"

												if l := len("able"); len(elem) >= l && elem[0:l] == "able" {
													elem = elem[l:]
												} else {
													break
												}

												if len(elem) == 0 {
													// Leaf node.
													switch method {
													case "GET":
														r.name = GetRelatableGroupsOperation
														r.summary = ""
														r.operationID = "getRelatableGroups"
														r.operationGroup = ""
														r.pathPattern = "/v1/groups/{id}/relatable"
														r.args = args
														r.count = 1
														return r, true
													default:
														return
													}
												}

											case 'e': // Prefix: "ed"

												if l := len("ed"); len(elem) >= l && elem[0:l] == "ed" {
													elem = elem[l:]
												} else {
													break
												}

												if len(elem) == 0 {
													// Leaf node.
													switch method {
													case "DELETE":
														r.name = RemoveRelatedGroupsOperation
														r.summary = ""
														r.operationID = "removeRelatedGroups"
														r.operationGroup = ""
														r.pathPattern = "/v1/groups/{id}/related"
														r.args = args
														r.count = 1
														return r, true
													case "GET":
														r.name = GetRelatedGroupsOperation
														r.summary = ""
														r.operationID = "getRelatedGroups"
														r.operationGroup = ""
														r.pathPattern = "/v1/groups/{id}/related"
														r.args = args
														r.count = 1
														return r, true
													case "PUT":
														r.name = UpdateRelatedGroupsOperation
														r.summary = ""
														r.operationID = "updateRelatedGroups"
														r.operationGroup = ""
														r.pathPattern = "/v1/groups/{id}/related"
														r.args = args
														r.count = 1
														return r, true
													default:
														return
													}
												}

											}

										case 'q': // Prefix: "quest_walkthrough"

											if l := len("quest_walkthrough"); len(elem) >= l && elem[0:l] == "quest_walkthrough" {
												elem = elem[l:]
											} else {
												break
											}

											if len(elem) == 0 {
												// Leaf node.
												switch method {
												case "POST":
													r.name = RequestGroupWalkthroughOperation
													r.summary = ""
													r.operationID = "requestGroupWalkthrough"
													r.operationGroup = ""
													r.pathPattern = "/v1/groups/{id}/request_walkthrough"
													r.args = args
													r.count = 1
													return r, true
												default:
													return
												}
											}

										}

									case 's': // Prefix: "set_title"

										if l := len("set_title"); len(elem) >= l && elem[0:l] == "set_title" {
											elem = elem[l:]
										} else {
											break
										}

										if len(elem) == 0 {
											// Leaf node.
											switch method {
											case "POST":
												r.name = SetGroupTitleOperation
												r.summary = ""
												r.operationID = "setGroupTitle"
												r.operationGroup = ""
												r.pathPattern = "/v1/groups/{id}/set_title"
												r.args = args
												r.count = 1
												return r, true
											default:
												return
											}
										}

									case 't': // Prefix: "t"

										if l := len("t"); len(elem) >= l && elem[0:l] == "t" {
											elem = elem[l:]
										} else {
											break
										}

										if len(elem) == 0 {
											break
										}
										switch elem[0] {
										case 'a': // Prefix: "ake_over"

											if l := len("ake_over"); len(elem) >= l && elem[0:l] == "ake_over" {
												elem = elem[l:]
											} else {
												break
											}

											if len(elem) == 0 {
												// Leaf node.
												switch method {
												case "POST":
													r.name = TakeOverGroupOperation
													r.summary = ""
													r.operationID = "takeOverGroup"
													r.operationGroup = ""
													r.pathPattern = "/v1/groups/{id}/take_over"
													r.args = args
													r.count = 1
													return r, true
												default:
													return
												}
											}

										case 'r': // Prefix: "ransfer"

											if l := len("ransfer"); len(elem) >= l && elem[0:l] == "ransfer" {
												elem = elem[l:]
											} else {
												break
											}

											if len(elem) == 0 {
												switch method {
												case "DELETE":
													r.name = CancelGroupTransferOperation
													r.summary = ""
													r.operationID = "cancelGroupTransfer"
													r.operationGroup = ""
													r.pathPattern = "/v1/groups/{id}/transfer"
													r.args = args
													r.count = 1
													return r, true
												case "PUT":
													r.name = AcceptGroupTransferOperation
													r.summary = ""
													r.operationID = "acceptGroupTransfer"
													r.operationGroup = ""
													r.pathPattern = "/v1/groups/{id}/transfer"
													r.args = args
													r.count = 1
													return r, true
												default:
													return
												}
											}
											switch elem[0] {
											case '/': // Prefix: "/withdraw"

												if l := len("/withdraw"); len(elem) >= l && elem[0:l] == "/withdraw" {
													elem = elem[l:]
												} else {
													break
												}

												if len(elem) == 0 {
													// Leaf node.
													switch method {
													case "PUT":
														r.name = WithdrawGroupTransferOperation
														r.summary = ""
														r.operationID = "withdrawGroupTransfer"
														r.operationGroup = ""
														r.pathPattern = "/v1/groups/{id}/transfer/withdraw"
														r.args = args
														r.count = 1
														return r, true
													default:
														return
													}
												}

											}

										}

									case 'u': // Prefix: "u"

										if l := len("u"); len(elem) >= l && elem[0:l] == "u" {
											elem = elem[l:]
										} else {
											break
										}

										if len(elem) == 0 {
											break
										}
										switch elem[0] {
										case 'n': // Prefix: "nban/"

											if l := len("nban/"); len(elem) >= l && elem[0:l] == "nban/" {
												elem = elem[l:]
											} else {
												break
											}

											// Param: "userId"
											// Leaf parameter, slashes are prohibited
											idx := strings.IndexByte(elem, '/')
											if idx >= 0 {
												break
											}
											args[1] = elem
											elem = ""

											if len(elem) == 0 {
												// Leaf node.
												switch method {
												case "POST":
													r.name = UnbanGroupUserOperation
													r.summary = ""
													r.operationID = "unbanGroupUser"
													r.operationGroup = ""
													r.pathPattern = "/v1/groups/{id}/unban/{userId}"
													r.args = args
													r.count = 2
													return r, true
												default:
													return
												}
											}

										case 's': // Prefix: "sers/invitable"

											if l := len("sers/invitable"); len(elem) >= l && elem[0:l] == "sers/invitable" {
												elem = elem[l:]
											} else {
												break
											}

											if len(elem) == 0 {
												// Leaf node.
												switch method {
												case "GET":
													r.name = GetInvitableGroupUsersOperation
													r.summary = ""
													r.operationID = "getInvitableGroupUsers"
													r.operationGroup = ""
													r.pathPattern = "/v1/groups/{group_id}/users/invitable"
													r.args = args
													r.count = 1
													return r, true
												default:
													return
												}
											}

										}

									case 'v': // Prefix: "visit"

										if l := len("visit"); len(elem) >= l && elem[0:l] == "visit" {
											elem = elem[l:]
										} else {
											break
										}

										if len(elem) == 0 {
											// Leaf node.
											switch method {
											case "POST":
												r.name = VisitGroupOperation
												r.summary = ""
												r.operationID = "visitGroup"
												r.operationGroup = ""
												r.pathPattern = "/v1/groups/{id}/visit"
												r.args = args
												r.count = 1
												return r, true
											default:
												return
											}
										}

									}

								}

							}

						}

					case 'h': // Prefix: "hidden/"

						if l := len("hidden/"); len(elem) >= l && elem[0:l] == "hidden/" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							break
						}
						switch elem[0] {
						case 'c': // Prefix: "chats"

							if l := len("chats"); len(elem) >= l && elem[0:l] == "chats" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch method {
								case "DELETE":
									r.name = UnhideChatsOperation
									r.summary = ""
									r.operationID = "unhideChats"
									r.operationGroup = ""
									r.pathPattern = "/v1/hidden/chats"
									r.args = args
									r.count = 0
									return r, true
								case "GET":
									r.name = ListHiddenChatsOperation
									r.summary = ""
									r.operationID = "listHiddenChats"
									r.operationGroup = ""
									r.pathPattern = "/v1/hidden/chats"
									r.args = args
									r.count = 0
									return r, true
								case "POST":
									r.name = HideChatsOperation
									r.summary = ""
									r.operationID = "hideChats"
									r.operationGroup = ""
									r.pathPattern = "/v1/hidden/chats"
									r.args = args
									r.count = 0
									return r, true
								default:
									return
								}
							}

						case 'u': // Prefix: "users"

							if l := len("users"); len(elem) >= l && elem[0:l] == "users" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch method {
								case "DELETE":
									r.name = UnhideUsersOperation
									r.summary = ""
									r.operationID = "unhideUsers"
									r.operationGroup = ""
									r.pathPattern = "/v1/hidden/users"
									r.args = args
									r.count = 0
									return r, true
								case "GET":
									r.name = ListHiddenUsersOperation
									r.summary = ""
									r.operationID = "listHiddenUsers"
									r.operationGroup = ""
									r.pathPattern = "/v1/hidden/users"
									r.args = args
									r.count = 0
									return r, true
								case "POST":
									r.name = HideUsersOperation
									r.summary = ""
									r.operationID = "hideUsers"
									r.operationGroup = ""
									r.pathPattern = "/v1/hidden/users"
									r.args = args
									r.count = 0
									return r, true
								default:
									return
								}
							}

						case 'w': // Prefix: "words"

							if l := len("words"); len(elem) >= l && elem[0:l] == "words" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch method {
								case "DELETE":
									r.name = DeleteMuteKeywordOperation
									r.summary = ""
									r.operationID = "deleteMuteKeyword"
									r.operationGroup = ""
									r.pathPattern = "/v1/hidden/words"
									r.args = args
									r.count = 0
									return r, true
								case "GET":
									r.name = ListMuteKeywordsOperation
									r.summary = ""
									r.operationID = "listMuteKeywords"
									r.operationGroup = ""
									r.pathPattern = "/v1/hidden/words"
									r.args = args
									r.count = 0
									return r, true
								case "POST":
									r.name = CreateMuteKeywordOperation
									r.summary = ""
									r.operationID = "createMuteKeyword"
									r.operationGroup = ""
									r.pathPattern = "/v1/hidden/words"
									r.args = args
									r.count = 0
									return r, true
								default:
									return
								}
							}

						}

					case 'p': // Prefix: "p"

						if l := len("p"); len(elem) >= l && elem[0:l] == "p" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							break
						}
						switch elem[0] {
						case 'i': // Prefix: "inned/"

							if l := len("inned/"); len(elem) >= l && elem[0:l] == "inned/" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								break
							}
							switch elem[0] {
							case 'g': // Prefix: "groups"

								if l := len("groups"); len(elem) >= l && elem[0:l] == "groups" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									switch method {
									case "POST":
										r.name = PinGroupOperation
										r.summary = ""
										r.operationID = "pinGroup"
										r.operationGroup = ""
										r.pathPattern = "/v1/pinned/groups"
										r.args = args
										r.count = 0
										return r, true
									default:
										return
									}
								}
								switch elem[0] {
								case '/': // Prefix: "/"

									if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
										elem = elem[l:]
									} else {
										break
									}

									// Param: "id"
									// Leaf parameter, slashes are prohibited
									idx := strings.IndexByte(elem, '/')
									if idx >= 0 {
										break
									}
									args[0] = elem
									elem = ""

									if len(elem) == 0 {
										// Leaf node.
										switch method {
										case "DELETE":
											r.name = UnpinGroupOperation
											r.summary = ""
											r.operationID = "unpinGroup"
											r.operationGroup = ""
											r.pathPattern = "/v1/pinned/groups/{id}"
											r.args = args
											r.count = 1
											return r, true
										default:
											return
										}
									}

								}

							case 'p': // Prefix: "posts"

								if l := len("posts"); len(elem) >= l && elem[0:l] == "posts" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch method {
									case "POST":
										r.name = PinPostOperation
										r.summary = ""
										r.operationID = "pinPost"
										r.operationGroup = ""
										r.pathPattern = "/v1/pinned/posts"
										r.args = args
										r.count = 0
										return r, true
									default:
										return
									}
								}

							case 'r': // Prefix: "reviews"

								if l := len("reviews"); len(elem) >= l && elem[0:l] == "reviews" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									switch method {
									case "POST":
										r.name = PinReviewOperation
										r.summary = ""
										r.operationID = "pinReview"
										r.operationGroup = ""
										r.pathPattern = "/v1/pinned/reviews"
										r.args = args
										r.count = 0
										return r, true
									default:
										return
									}
								}
								switch elem[0] {
								case '/': // Prefix: "/"

									if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
										elem = elem[l:]
									} else {
										break
									}

									// Param: "id"
									// Leaf parameter, slashes are prohibited
									idx := strings.IndexByte(elem, '/')
									if idx >= 0 {
										break
									}
									args[0] = elem
									elem = ""

									if len(elem) == 0 {
										// Leaf node.
										switch method {
										case "DELETE":
											r.name = UnpinReviewOperation
											r.summary = ""
											r.operationID = "unpinReview"
											r.operationGroup = ""
											r.pathPattern = "/v1/pinned/reviews/{id}"
											r.args = args
											r.count = 1
											return r, true
										default:
											return
										}
									}

								}

							}

						case 'o': // Prefix: "osts/"

							if l := len("osts/"); len(elem) >= l && elem[0:l] == "osts/" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								break
							}
							switch elem[0] {
							case 'a': // Prefix: "active_call"
								origElem := elem
								if l := len("active_call"); len(elem) >= l && elem[0:l] == "active_call" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch method {
									case "GET":
										r.name = GetActiveCallPostOperation
										r.summary = ""
										r.operationID = "getActiveCallPost"
										r.operationGroup = ""
										r.pathPattern = "/v1/posts/active_call"
										r.args = args
										r.count = 0
										return r, true
									default:
										return
									}
								}

								elem = origElem
							case 'd': // Prefix: "delete_all_post"
								origElem := elem
								if l := len("delete_all_post"); len(elem) >= l && elem[0:l] == "delete_all_post" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch method {
									case "POST":
										r.name = DeleteAllPostsOperation
										r.summary = ""
										r.operationID = "deleteAllPosts"
										r.operationGroup = ""
										r.pathPattern = "/v1/posts/delete_all_post"
										r.args = args
										r.count = 0
										return r, true
									default:
										return
									}
								}

								elem = origElem
							case 'r': // Prefix: "recommended_tag"
								origElem := elem
								if l := len("recommended_tag"); len(elem) >= l && elem[0:l] == "recommended_tag" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch method {
									case "POST":
										r.name = GetRecommendedPostTagsOperation
										r.summary = ""
										r.operationID = "getRecommendedPostTags"
										r.operationGroup = ""
										r.pathPattern = "/v1/posts/recommended_tag"
										r.args = args
										r.count = 0
										return r, true
									default:
										return
									}
								}

								elem = origElem
							case 'v': // Prefix: "v"
								origElem := elem
								if l := len("v"); len(elem) >= l && elem[0:l] == "v" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									break
								}
								switch elem[0] {
								case 'a': // Prefix: "alidate"

									if l := len("alidate"); len(elem) >= l && elem[0:l] == "alidate" {
										elem = elem[l:]
									} else {
										break
									}

									if len(elem) == 0 {
										// Leaf node.
										switch method {
										case "POST":
											r.name = ValidatePostOperation
											r.summary = ""
											r.operationID = "validatePost"
											r.operationGroup = ""
											r.pathPattern = "/v1/posts/validate"
											r.args = args
											r.count = 0
											return r, true
										default:
											return
										}
									}

								case 'i': // Prefix: "ideos/"

									if l := len("ideos/"); len(elem) >= l && elem[0:l] == "ideos/" {
										elem = elem[l:]
									} else {
										break
									}

									// Param: "videoId"
									// Match until "/"
									idx := strings.IndexByte(elem, '/')
									if idx < 0 {
										idx = len(elem)
									}
									args[0] = elem[:idx]
									elem = elem[idx:]

									if len(elem) == 0 {
										break
									}
									switch elem[0] {
									case '/': // Prefix: "/view"

										if l := len("/view"); len(elem) >= l && elem[0:l] == "/view" {
											elem = elem[l:]
										} else {
											break
										}

										if len(elem) == 0 {
											// Leaf node.
											switch method {
											case "POST":
												r.name = ViewPostVideoOperation
												r.summary = ""
												r.operationID = "viewPostVideo"
												r.operationGroup = ""
												r.pathPattern = "/v1/posts/videos/{videoId}/view"
												r.args = args
												r.count = 1
												return r, true
											default:
												return
											}
										}

									}

								}

								elem = origElem
							}
							// Param: "post_id"
							// Match until "/"
							idx := strings.IndexByte(elem, '/')
							if idx < 0 {
								idx = len(elem)
							}
							args[0] = elem[:idx]
							elem = elem[idx:]

							if len(elem) == 0 {
								break
							}
							switch elem[0] {
							case '/': // Prefix: "/"

								if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									break
								}
								switch elem[0] {
								case 'g': // Prefix: "gift_transactions"

									if l := len("gift_transactions"); len(elem) >= l && elem[0:l] == "gift_transactions" {
										elem = elem[l:]
									} else {
										break
									}

									if len(elem) == 0 {
										// Leaf node.
										switch method {
										case "GET":
											r.name = GetPostGiftTransactionsOperation
											r.summary = ""
											r.operationID = "getPostGiftTransactions"
											r.operationGroup = ""
											r.pathPattern = "/v1/posts/{post_id}/gift_transactions"
											r.args = args
											r.count = 1
											return r, true
										default:
											return
										}
									}

								case 'l': // Prefix: "likers"

									if l := len("likers"); len(elem) >= l && elem[0:l] == "likers" {
										elem = elem[l:]
									} else {
										break
									}

									if len(elem) == 0 {
										// Leaf node.
										switch method {
										case "GET":
											r.name = GetPostLikersOperation
											r.summary = ""
											r.operationID = "getPostLikers"
											r.operationGroup = ""
											r.pathPattern = "/v1/posts/{id}/likers"
											r.args = args
											r.count = 1
											return r, true
										default:
											return
										}
									}

								case 'u': // Prefix: "unlike"

									if l := len("unlike"); len(elem) >= l && elem[0:l] == "unlike" {
										elem = elem[l:]
									} else {
										break
									}

									if len(elem) == 0 {
										// Leaf node.
										switch method {
										case "POST":
											r.name = UnlikePostOperation
											r.summary = ""
											r.operationID = "unlikePost"
											r.operationGroup = ""
											r.pathPattern = "/v1/posts/{id}/unlike"
											r.args = args
											r.count = 1
											return r, true
										default:
											return
										}
									}

								}

							}

						}

					case 'r': // Prefix: "received_gifts"

						if l := len("received_gifts"); len(elem) >= l && elem[0:l] == "received_gifts" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							switch method {
							case "GET":
								r.name = ListReceivedGiftsV1Operation
								r.summary = ""
								r.operationID = "listReceivedGiftsV1"
								r.operationGroup = ""
								r.pathPattern = "/v1/received_gifts"
								r.args = args
								r.count = 0
								return r, true
							default:
								return
							}
						}
						switch elem[0] {
						case '/': // Prefix: "/"

							if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
								elem = elem[l:]
							} else {
								break
							}

							// Param: "gift_id"
							// Match until "/"
							idx := strings.IndexByte(elem, '/')
							if idx < 0 {
								idx = len(elem)
							}
							args[0] = elem[:idx]
							elem = elem[idx:]

							if len(elem) == 0 {
								switch method {
								case "GET":
									r.name = GetReceivedGiftTransactionOperation
									r.summary = ""
									r.operationID = "getReceivedGiftTransaction"
									r.operationGroup = ""
									r.pathPattern = "/v1/received_gifts/{gift_transaction_uuid}"
									r.args = args
									r.count = 1
									return r, true
								default:
									return
								}
							}
							switch elem[0] {
							case '/': // Prefix: "/senders"

								if l := len("/senders"); len(elem) >= l && elem[0:l] == "/senders" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch method {
									case "GET":
										r.name = GetReceivedGiftSendersOperation
										r.summary = ""
										r.operationID = "getReceivedGiftSenders"
										r.operationGroup = ""
										r.pathPattern = "/v1/received_gifts/{gift_id}/senders"
										r.args = args
										r.count = 1
										return r, true
									default:
										return
									}
								}

							}

						}

					case 't': // Prefix: "threads"

						if l := len("threads"); len(elem) >= l && elem[0:l] == "threads" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							switch method {
							case "GET":
								r.name = ListThreadsOperation
								r.summary = ""
								r.operationID = "listThreads"
								r.operationGroup = ""
								r.pathPattern = "/v1/threads"
								r.args = args
								r.count = 0
								return r, true
							case "POST":
								r.name = CreateThreadOperation
								r.summary = ""
								r.operationID = "createThread"
								r.operationGroup = ""
								r.pathPattern = "/v1/threads"
								r.args = args
								r.count = 0
								return r, true
							default:
								return
							}
						}
						switch elem[0] {
						case '/': // Prefix: "/"

							if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								break
							}
							switch elem[0] {
							case 'j': // Prefix: "joined_statuses"
								origElem := elem
								if l := len("joined_statuses"); len(elem) >= l && elem[0:l] == "joined_statuses" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch method {
									case "GET":
										r.name = GetJoinedThreadStatusesOperation
										r.summary = ""
										r.operationID = "getJoinedThreadStatuses"
										r.operationGroup = ""
										r.pathPattern = "/v1/threads/joined_statuses"
										r.args = args
										r.count = 0
										return r, true
									default:
										return
									}
								}

								elem = origElem
							}
							// Param: "thread_id"
							// Match until "/"
							idx := strings.IndexByte(elem, '/')
							if idx < 0 {
								idx = len(elem)
							}
							args[0] = elem[:idx]
							elem = elem[idx:]

							if len(elem) == 0 {
								switch method {
								case "DELETE":
									r.name = DeleteThreadOperation
									r.summary = ""
									r.operationID = "deleteThread"
									r.operationGroup = ""
									r.pathPattern = "/v1/threads/{id}"
									r.args = args
									r.count = 1
									return r, true
								case "GET":
									r.name = GetThreadOperation
									r.summary = ""
									r.operationID = "getThread"
									r.operationGroup = ""
									r.pathPattern = "/v1/threads/{id}"
									r.args = args
									r.count = 1
									return r, true
								case "PUT":
									r.name = UpdateThreadOperation
									r.summary = ""
									r.operationID = "updateThread"
									r.operationGroup = ""
									r.pathPattern = "/v1/threads/{id}"
									r.args = args
									r.count = 1
									return r, true
								default:
									return
								}
							}
							switch elem[0] {
							case '/': // Prefix: "/"

								if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									break
								}
								switch elem[0] {
								case 'm': // Prefix: "members/"

									if l := len("members/"); len(elem) >= l && elem[0:l] == "members/" {
										elem = elem[l:]
									} else {
										break
									}

									// Param: "id"
									// Leaf parameter, slashes are prohibited
									idx := strings.IndexByte(elem, '/')
									if idx >= 0 {
										break
									}
									args[1] = elem
									elem = ""

									if len(elem) == 0 {
										// Leaf node.
										switch method {
										case "DELETE":
											r.name = RemoveThreadMemberOperation
											r.summary = ""
											r.operationID = "removeThreadMember"
											r.operationGroup = ""
											r.pathPattern = "/v1/threads/{thread_id}/members/{id}"
											r.args = args
											r.count = 2
											return r, true
										case "POST":
											r.name = AddThreadMemberOperation
											r.summary = ""
											r.operationID = "addThreadMember"
											r.operationGroup = ""
											r.pathPattern = "/v1/threads/{thread_id}/members/{id}"
											r.args = args
											r.count = 2
											return r, true
										default:
											return
										}
									}

								case 'p': // Prefix: "posts"

									if l := len("posts"); len(elem) >= l && elem[0:l] == "posts" {
										elem = elem[l:]
									} else {
										break
									}

									if len(elem) == 0 {
										// Leaf node.
										switch method {
										case "GET":
											r.name = GetThreadPostsOperation
											r.summary = ""
											r.operationID = "getThreadPosts"
											r.operationGroup = ""
											r.pathPattern = "/v1/threads/{id}/posts"
											r.args = args
											r.count = 1
											return r, true
										case "POST":
											r.name = CreateThreadPostOperation
											r.summary = ""
											r.operationID = "createThreadPost"
											r.operationGroup = ""
											r.pathPattern = "/v1/threads/{id}/posts"
											r.args = args
											r.count = 1
											return r, true
										default:
											return
										}
									}

								}

							}

						}

					case 'u': // Prefix: "users/"

						if l := len("users/"); len(elem) >= l && elem[0:l] == "users/" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							break
						}
						switch elem[0] {
						case 'a': // Prefix: "a"
							origElem := elem
							if l := len("a"); len(elem) >= l && elem[0:l] == "a" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								break
							}
							switch elem[0] {
							case 'c': // Prefix: "ctive_followings"

								if l := len("ctive_followings"); len(elem) >= l && elem[0:l] == "ctive_followings" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch method {
									case "GET":
										r.name = GetActiveFollowingsOperation
										r.summary = ""
										r.operationID = "getActiveFollowings"
										r.operationGroup = ""
										r.pathPattern = "/v1/users/active_followings"
										r.args = args
										r.count = 0
										return r, true
									default:
										return
									}
								}

							case 'd': // Prefix: "dditonal_notification_setting"

								if l := len("dditonal_notification_setting"); len(elem) >= l && elem[0:l] == "dditonal_notification_setting" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch method {
									case "GET":
										r.name = GetAdditionalNotificationSettingOperation
										r.summary = ""
										r.operationID = "getAdditionalNotificationSetting"
										r.operationGroup = ""
										r.pathPattern = "/v1/users/additonal_notification_setting"
										r.args = args
										r.count = 0
										return r, true
									case "POST":
										r.name = UpdateAdditionalNotificationSettingOperation
										r.summary = ""
										r.operationID = "updateAdditionalNotificationSetting"
										r.operationGroup = ""
										r.pathPattern = "/v1/users/additonal_notification_setting"
										r.args = args
										r.count = 0
										return r, true
									default:
										return
									}
								}

							case 'l': // Prefix: "live"

								if l := len("live"); len(elem) >= l && elem[0:l] == "live" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch method {
									case "POST":
										r.name = PingAliveOperation
										r.summary = ""
										r.operationID = "pingAlive"
										r.operationGroup = ""
										r.pathPattern = "/v1/users/alive"
										r.args = args
										r.count = 0
										return r, true
									default:
										return
									}
								}

							}

							elem = origElem
						case 'b': // Prefix: "block_ids"
							origElem := elem
							if l := len("block_ids"); len(elem) >= l && elem[0:l] == "block_ids" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch method {
								case "GET":
									r.name = GetBlockedUserIdsOperation
									r.summary = ""
									r.operationID = "getBlockedUserIds"
									r.operationGroup = ""
									r.pathPattern = "/v1/users/block_ids"
									r.args = args
									r.count = 0
									return r, true
								default:
									return
								}
							}

							elem = origElem
						case 'c': // Prefix: "c"
							origElem := elem
							if l := len("c"); len(elem) >= l && elem[0:l] == "c" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								break
							}
							switch elem[0] {
							case 'h': // Prefix: "hange_"

								if l := len("hange_"); len(elem) >= l && elem[0:l] == "hange_" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									break
								}
								switch elem[0] {
								case 'e': // Prefix: "email"

									if l := len("email"); len(elem) >= l && elem[0:l] == "email" {
										elem = elem[l:]
									} else {
										break
									}

									if len(elem) == 0 {
										// Leaf node.
										switch method {
										case "PUT":
											r.name = ChangeEmailOperation
											r.summary = ""
											r.operationID = "changeEmail"
											r.operationGroup = ""
											r.pathPattern = "/v1/users/change_email"
											r.args = args
											r.count = 0
											return r, true
										default:
											return
										}
									}

								case 'p': // Prefix: "password"

									if l := len("password"); len(elem) >= l && elem[0:l] == "password" {
										elem = elem[l:]
									} else {
										break
									}

									if len(elem) == 0 {
										// Leaf node.
										switch method {
										case "PUT":
											r.name = ChangePasswordOperation
											r.summary = ""
											r.operationID = "changePassword"
											r.operationGroup = ""
											r.pathPattern = "/v1/users/change_password"
											r.args = args
											r.count = 0
											return r, true
										default:
											return
										}
									}

								}

							case 'u': // Prefix: "ustom_definitions"

								if l := len("ustom_definitions"); len(elem) >= l && elem[0:l] == "ustom_definitions" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch method {
									case "GET":
										r.name = GetUserCustomDefinitionsOperation
										r.summary = ""
										r.operationID = "getUserCustomDefinitions"
										r.operationGroup = ""
										r.pathPattern = "/v1/users/custom_definitions"
										r.args = args
										r.count = 0
										return r, true
									default:
										return
									}
								}

							}

							elem = origElem
						case 'd': // Prefix: "default_settings"
							origElem := elem
							if l := len("default_settings"); len(elem) >= l && elem[0:l] == "default_settings" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch method {
								case "GET":
									r.name = GetDefaultSettingsOperation
									r.summary = ""
									r.operationID = "getDefaultSettings"
									r.operationGroup = ""
									r.pathPattern = "/v1/users/default_settings"
									r.args = args
									r.count = 0
									return r, true
								default:
									return
								}
							}

							elem = origElem
						case 'f': // Prefix: "following"
							origElem := elem
							if l := len("following"); len(elem) >= l && elem[0:l] == "following" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								break
							}
							switch elem[0] {
							case '_': // Prefix: "_born_today"

								if l := len("_born_today"); len(elem) >= l && elem[0:l] == "_born_today" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch method {
									case "GET":
										r.name = GetFollowingsBornTodayOperation
										r.summary = ""
										r.operationID = "getFollowingsBornToday"
										r.operationGroup = ""
										r.pathPattern = "/v1/users/following_born_today"
										r.args = args
										r.count = 0
										return r, true
									default:
										return
									}
								}

							case 's': // Prefix: "s/chatable"

								if l := len("s/chatable"); len(elem) >= l && elem[0:l] == "s/chatable" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch method {
									case "POST":
										r.name = GetChatableFollowingsOperation
										r.summary = ""
										r.operationID = "getChatableFollowings"
										r.operationGroup = ""
										r.pathPattern = "/v1/users/followings/chatable"
										r.args = args
										r.count = 0
										return r, true
									default:
										return
									}
								}

							}

							elem = origElem
						case 'h': // Prefix: "hima"
							origElem := elem
							if l := len("hima"); len(elem) >= l && elem[0:l] == "hima" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch method {
								case "POST":
									r.name = SetHimaOperation
									r.summary = ""
									r.operationID = "setHima"
									r.operationGroup = ""
									r.pathPattern = "/v1/users/hima"
									r.args = args
									r.count = 0
									return r, true
								default:
									return
								}
							}

							elem = origElem
						case 'i': // Prefix: "interests"
							origElem := elem
							if l := len("interests"); len(elem) >= l && elem[0:l] == "interests" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch method {
								case "GET":
									r.name = GetUserInterestsOperation
									r.summary = ""
									r.operationID = "getUserInterests"
									r.operationGroup = ""
									r.pathPattern = "/v1/users/interests"
									r.args = args
									r.count = 0
									return r, true
								case "PUT":
									r.name = UpdateUserInterestsOperation
									r.summary = ""
									r.operationID = "updateUserInterests"
									r.operationGroup = ""
									r.pathPattern = "/v1/users/interests"
									r.args = args
									r.count = 0
									return r, true
								default:
									return
								}
							}

							elem = origElem
						case 'l': // Prefix: "l"
							origElem := elem
							if l := len("l"); len(elem) >= l && elem[0:l] == "l" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								break
							}
							switch elem[0] {
							case 'a': // Prefix: "anguage"

								if l := len("anguage"); len(elem) >= l && elem[0:l] == "anguage" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch method {
									case "PUT":
										r.name = UpdateLanguageOperation
										r.summary = ""
										r.operationID = "updateLanguage"
										r.operationGroup = ""
										r.pathPattern = "/v1/users/language"
										r.args = args
										r.count = 0
										return r, true
									default:
										return
									}
								}

							case 'i': // Prefix: "ist_id"

								if l := len("ist_id"); len(elem) >= l && elem[0:l] == "ist_id" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch method {
									case "GET":
										r.name = GetUsersByIdsOperation
										r.summary = ""
										r.operationID = "getUsersByIds"
										r.operationGroup = ""
										r.pathPattern = "/v1/users/list_id"
										r.args = args
										r.count = 0
										return r, true
									default:
										return
									}
								}

							case 'o': // Prefix: "ogout"

								if l := len("ogout"); len(elem) >= l && elem[0:l] == "ogout" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch method {
									case "POST":
										r.name = LogoutOperation
										r.summary = ""
										r.operationID = "logout"
										r.operationGroup = ""
										r.pathPattern = "/v1/users/logout"
										r.args = args
										r.count = 0
										return r, true
									default:
										return
									}
								}

							}

							elem = origElem
						case 'p': // Prefix: "p"
							origElem := elem
							if l := len("p"); len(elem) >= l && elem[0:l] == "p" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								break
							}
							switch elem[0] {
							case 'o': // Prefix: "olicy_agreements"

								if l := len("olicy_agreements"); len(elem) >= l && elem[0:l] == "olicy_agreements" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									switch method {
									case "GET":
										r.name = GetPolicyAgreementsOperation
										r.summary = ""
										r.operationID = "getPolicyAgreements"
										r.operationGroup = ""
										r.pathPattern = "/v1/users/policy_agreements"
										r.args = args
										r.count = 0
										return r, true
									default:
										return
									}
								}
								switch elem[0] {
								case '/': // Prefix: "/"

									if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
										elem = elem[l:]
									} else {
										break
									}

									// Param: "type"
									// Leaf parameter, slashes are prohibited
									idx := strings.IndexByte(elem, '/')
									if idx >= 0 {
										break
									}
									args[0] = elem
									elem = ""

									if len(elem) == 0 {
										// Leaf node.
										switch method {
										case "POST":
											r.name = AgreePolicyOperation
											r.summary = ""
											r.operationID = "agreePolicy"
											r.operationGroup = ""
											r.pathPattern = "/v1/users/policy_agreements/{type}"
											r.args = args
											r.count = 1
											return r, true
										default:
											return
										}
									}

								}

							case 'r': // Prefix: "resigned_url"

								if l := len("resigned_url"); len(elem) >= l && elem[0:l] == "resigned_url" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch method {
									case "GET":
										r.name = GetUserPresignedUrlOperation
										r.summary = ""
										r.operationID = "getUserPresignedUrl"
										r.operationGroup = ""
										r.pathPattern = "/v1/users/presigned_url"
										r.args = args
										r.count = 0
										return r, true
									default:
										return
									}
								}

							}

							elem = origElem
						case 'q': // Prefix: "qr_codes/"
							origElem := elem
							if l := len("qr_codes/"); len(elem) >= l && elem[0:l] == "qr_codes/" {
								elem = elem[l:]
							} else {
								break
							}

							// Param: "qr"
							// Leaf parameter, slashes are prohibited
							idx := strings.IndexByte(elem, '/')
							if idx >= 0 {
								break
							}
							args[0] = elem
							elem = ""

							if len(elem) == 0 {
								// Leaf node.
								switch method {
								case "GET":
									r.name = GetUserByQrOperation
									r.summary = ""
									r.operationID = "getUserByQr"
									r.operationGroup = ""
									r.pathPattern = "/v1/users/qr_codes/{qr}"
									r.args = args
									r.count = 1
									return r, true
								default:
									return
								}
							}

							elem = origElem
						case 'r': // Prefix: "re"
							origElem := elem
							if l := len("re"); len(elem) >= l && elem[0:l] == "re" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								break
							}
							switch elem[0] {
							case 's': // Prefix: "set_"

								if l := len("set_"); len(elem) >= l && elem[0:l] == "set_" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									break
								}
								switch elem[0] {
								case 'c': // Prefix: "counters"

									if l := len("counters"); len(elem) >= l && elem[0:l] == "counters" {
										elem = elem[l:]
									} else {
										break
									}

									if len(elem) == 0 {
										// Leaf node.
										switch method {
										case "GET":
											r.name = GetResetCountersOperation
											r.summary = ""
											r.operationID = "getResetCounters"
											r.operationGroup = ""
											r.pathPattern = "/v1/users/reset_counters"
											r.args = args
											r.count = 0
											return r, true
										case "POST":
											r.name = ResetCountersOperation
											r.summary = ""
											r.operationID = "resetCounters"
											r.operationGroup = ""
											r.pathPattern = "/v1/users/reset_counters"
											r.args = args
											r.count = 0
											return r, true
										default:
											return
										}
									}

								case 'p': // Prefix: "password"

									if l := len("password"); len(elem) >= l && elem[0:l] == "password" {
										elem = elem[l:]
									} else {
										break
									}

									if len(elem) == 0 {
										// Leaf node.
										switch method {
										case "PUT":
											r.name = ResetPasswordOperation
											r.summary = ""
											r.operationID = "resetPassword"
											r.operationGroup = ""
											r.pathPattern = "/v1/users/reset_password"
											r.args = args
											r.count = 0
											return r, true
										default:
											return
										}
									}

								}

							case 'v': // Prefix: "views"

								if l := len("views"); len(elem) >= l && elem[0:l] == "views" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									switch method {
									case "DELETE":
										r.name = DeleteMyReviewsOperation
										r.summary = ""
										r.operationID = "deleteMyReviews"
										r.operationGroup = ""
										r.pathPattern = "/v1/users/reviews"
										r.args = args
										r.count = 0
										return r, true
									case "POST":
										r.name = CreateReviewOperation
										r.summary = ""
										r.operationID = "createReview"
										r.operationGroup = ""
										r.pathPattern = "/v1/users/reviews"
										r.args = args
										r.count = 0
										return r, true
									default:
										return
									}
								}
								switch elem[0] {
								case '/': // Prefix: "/mine"

									if l := len("/mine"); len(elem) >= l && elem[0:l] == "/mine" {
										elem = elem[l:]
									} else {
										break
									}

									if len(elem) == 0 {
										// Leaf node.
										switch method {
										case "GET":
											r.name = GetMyReviewsOperation
											r.summary = ""
											r.operationID = "getMyReviews"
											r.operationGroup = ""
											r.pathPattern = "/v1/users/reviews/mine"
											r.args = args
											r.count = 0
											return r, true
										default:
											return
										}
									}

								}

							}

							elem = origElem
						case 's': // Prefix: "se"
							origElem := elem
							if l := len("se"); len(elem) >= l && elem[0:l] == "se" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								break
							}
							switch elem[0] {
							case 'a': // Prefix: "arch"

								if l := len("arch"); len(elem) >= l && elem[0:l] == "arch" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch method {
									case "GET":
										r.name = SearchUsersOperation
										r.summary = ""
										r.operationID = "searchUsers"
										r.operationGroup = ""
										r.pathPattern = "/v1/users/search"
										r.args = args
										r.count = 0
										return r, true
									default:
										return
									}
								}

							case 'c': // Prefix: "cret"

								if l := len("cret"); len(elem) >= l && elem[0:l] == "cret" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									switch method {
									case "GET":
										r.name = GetTwoFactorAuthRequestInfoOperation
										r.summary = ""
										r.operationID = "getTwoFactorAuthRequestInfo"
										r.operationGroup = ""
										r.pathPattern = "/v1/users/secret"
										r.args = args
										r.count = 0
										return r, true
									default:
										return
									}
								}
								switch elem[0] {
								case '/': // Prefix: "/"

									if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
										elem = elem[l:]
									} else {
										break
									}

									if len(elem) == 0 {
										break
									}
									switch elem[0] {
									case 'd': // Prefix: "disable"

										if l := len("disable"); len(elem) >= l && elem[0:l] == "disable" {
											elem = elem[l:]
										} else {
											break
										}

										if len(elem) == 0 {
											// Leaf node.
											switch method {
											case "PUT":
												r.name = DisableTwoFactorAuthOperation
												r.summary = ""
												r.operationID = "disableTwoFactorAuth"
												r.operationGroup = ""
												r.pathPattern = "/v1/users/secret/disable"
												r.args = args
												r.count = 0
												return r, true
											default:
												return
											}
										}

									case 'e': // Prefix: "enable"

										if l := len("enable"); len(elem) >= l && elem[0:l] == "enable" {
											elem = elem[l:]
										} else {
											break
										}

										if len(elem) == 0 {
											// Leaf node.
											switch method {
											case "PUT":
												r.name = EnableTwoFactorAuthOperation
												r.summary = ""
												r.operationID = "enableTwoFactorAuth"
												r.operationGroup = ""
												r.pathPattern = "/v1/users/secret/enable"
												r.args = args
												r.count = 0
												return r, true
											default:
												return
											}
										}

									case 's': // Prefix: "status"

										if l := len("status"); len(elem) >= l && elem[0:l] == "status" {
											elem = elem[l:]
										} else {
											break
										}

										if len(elem) == 0 {
											// Leaf node.
											switch method {
											case "GET":
												r.name = GetTwoFactorAuthStatusOperation
												r.summary = ""
												r.operationID = "getTwoFactorAuthStatus"
												r.operationGroup = ""
												r.pathPattern = "/v1/users/secret/status"
												r.args = args
												r.count = 0
												return r, true
											default:
												return
											}
										}

									}

								}

							}

							elem = origElem
						case 'w': // Prefix: "ws_token"
							origElem := elem
							if l := len("ws_token"); len(elem) >= l && elem[0:l] == "ws_token" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch method {
								case "GET":
									r.name = GetWebSocketTokenOperation
									r.summary = ""
									r.operationID = "getWebSocketToken"
									r.operationGroup = ""
									r.pathPattern = "/v1/users/ws_token"
									r.args = args
									r.count = 0
									return r, true
								default:
									return
								}
							}

							elem = origElem
						}
						// Param: "id"
						// Match until "/"
						idx := strings.IndexByte(elem, '/')
						if idx < 0 {
							idx = len(elem)
						}
						args[0] = elem[:idx]
						elem = elem[idx:]

						if len(elem) == 0 {
							break
						}
						switch elem[0] {
						case '/': // Prefix: "/"

							if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								break
							}
							switch elem[0] {
							case 'b': // Prefix: "b"

								if l := len("b"); len(elem) >= l && elem[0:l] == "b" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									break
								}
								switch elem[0] {
								case 'l': // Prefix: "lock"

									if l := len("lock"); len(elem) >= l && elem[0:l] == "lock" {
										elem = elem[l:]
									} else {
										break
									}

									if len(elem) == 0 {
										// Leaf node.
										switch method {
										case "POST":
											r.name = BlockUserOperation
											r.summary = ""
											r.operationID = "blockUser"
											r.operationGroup = ""
											r.pathPattern = "/v1/users/{id}/block"
											r.args = args
											r.count = 1
											return r, true
										default:
											return
										}
									}

								case 'o': // Prefix: "ookmarks"

									if l := len("ookmarks"); len(elem) >= l && elem[0:l] == "ookmarks" {
										elem = elem[l:]
									} else {
										break
									}

									if len(elem) == 0 {
										switch method {
										case "GET":
											r.name = GetBookmarkedPostsOperation
											r.summary = ""
											r.operationID = "getBookmarkedPosts"
											r.operationGroup = ""
											r.pathPattern = "/v1/users/{user_id}/bookmarks"
											r.args = args
											r.count = 1
											return r, true
										default:
											return
										}
									}
									switch elem[0] {
									case '/': // Prefix: "/"

										if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
											elem = elem[l:]
										} else {
											break
										}

										// Param: "id"
										// Leaf parameter, slashes are prohibited
										idx := strings.IndexByte(elem, '/')
										if idx >= 0 {
											break
										}
										args[1] = elem
										elem = ""

										if len(elem) == 0 {
											// Leaf node.
											switch method {
											case "DELETE":
												r.name = DeleteBookmarkOperation
												r.summary = ""
												r.operationID = "deleteBookmark"
												r.operationGroup = ""
												r.pathPattern = "/v1/users/{user_id}/bookmarks/{id}"
												r.args = args
												r.count = 2
												return r, true
											case "PUT":
												r.name = CreateBookmarkOperation
												r.summary = ""
												r.operationID = "createBookmark"
												r.operationGroup = ""
												r.pathPattern = "/v1/users/{user_id}/bookmarks/{id}"
												r.args = args
												r.count = 2
												return r, true
											default:
												return
											}
										}

									}

								}

							case 'f': // Prefix: "follow_recommended"

								if l := len("follow_recommended"); len(elem) >= l && elem[0:l] == "follow_recommended" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch method {
									case "GET":
										r.name = GetRecommendedFollowUsersOperation
										r.summary = ""
										r.operationID = "getRecommendedFollowUsers"
										r.operationGroup = ""
										r.pathPattern = "/v1/users/{id}/follow_recommended"
										r.args = args
										r.count = 1
										return r, true
									default:
										return
									}
								}

							case 'g': // Prefix: "gift_transactions"

								if l := len("gift_transactions"); len(elem) >= l && elem[0:l] == "gift_transactions" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch method {
									case "GET":
										r.name = GetUserGiftTransactionsOperation
										r.summary = ""
										r.operationID = "getUserGiftTransactions"
										r.operationGroup = ""
										r.pathPattern = "/v1/users/{user_id}/gift_transactions"
										r.args = args
										r.count = 1
										return r, true
									default:
										return
									}
								}

							}

						}

					}

				case '2': // Prefix: "2/"

					if l := len("2/"); len(elem) >= l && elem[0:l] == "2/" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						break
					}
					switch elem[0] {
					case 'c': // Prefix: "c"

						if l := len("c"); len(elem) >= l && elem[0:l] == "c" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							break
						}
						switch elem[0] {
						case 'a': // Prefix: "alls/"

							if l := len("alls/"); len(elem) >= l && elem[0:l] == "alls/" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								break
							}
							switch elem[0] {
							case 'a': // Prefix: "action_signature/validate"

								if l := len("action_signature/validate"); len(elem) >= l && elem[0:l] == "action_signature/validate" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch method {
									case "GET":
										r.name = ValidateCallActionSignatureOperation
										r.summary = ""
										r.operationID = "validateCallActionSignature"
										r.operationGroup = ""
										r.pathPattern = "/v2/calls/action_signature/validate"
										r.args = args
										r.count = 0
										return r, true
									default:
										return
									}
								}

							case 'c': // Prefix: "conferences/"

								if l := len("conferences/"); len(elem) >= l && elem[0:l] == "conferences/" {
									elem = elem[l:]
								} else {
									break
								}

								// Param: "call_id"
								// Leaf parameter, slashes are prohibited
								idx := strings.IndexByte(elem, '/')
								if idx >= 0 {
									break
								}
								args[0] = elem
								elem = ""

								if len(elem) == 0 {
									// Leaf node.
									switch method {
									case "GET":
										r.name = GetConferenceCallOperation
										r.summary = ""
										r.operationID = "getConferenceCall"
										r.operationGroup = ""
										r.pathPattern = "/v2/calls/conferences/{call_id}"
										r.args = args
										r.count = 1
										return r, true
									default:
										return
									}
								}

							case 'i': // Prefix: "invite"

								if l := len("invite"); len(elem) >= l && elem[0:l] == "invite" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch method {
									case "POST":
										r.name = InviteToCallOperation
										r.summary = ""
										r.operationID = "inviteToCall"
										r.operationGroup = ""
										r.pathPattern = "/v2/calls/invite"
										r.args = args
										r.count = 0
										return r, true
									default:
										return
									}
								}

							case 's': // Prefix: "start_conference_call"

								if l := len("start_conference_call"); len(elem) >= l && elem[0:l] == "start_conference_call" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch method {
									case "POST":
										r.name = StartConferenceCallOperation
										r.summary = ""
										r.operationID = "startConferenceCall"
										r.operationGroup = ""
										r.pathPattern = "/v2/calls/start_conference_call"
										r.args = args
										r.count = 0
										return r, true
									default:
										return
									}
								}

							}

						case 'h': // Prefix: "hat_rooms/"

							if l := len("hat_rooms/"); len(elem) >= l && elem[0:l] == "hat_rooms/" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								break
							}
							switch elem[0] {
							case 'u': // Prefix: "update"
								origElem := elem
								if l := len("update"); len(elem) >= l && elem[0:l] == "update" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch method {
									case "GET":
										r.name = GetUpdatedChatRoomsOperation
										r.summary = ""
										r.operationID = "getUpdatedChatRooms"
										r.operationGroup = ""
										r.pathPattern = "/v2/chat_rooms/update"
										r.args = args
										r.count = 0
										return r, true
									default:
										return
									}
								}

								elem = origElem
							}
							// Param: "id"
							// Match until "/"
							idx := strings.IndexByte(elem, '/')
							if idx < 0 {
								idx = len(elem)
							}
							args[0] = elem[:idx]
							elem = elem[idx:]

							if len(elem) == 0 {
								switch method {
								case "GET":
									r.name = GetChatRoomOperation
									r.summary = ""
									r.operationID = "getChatRoom"
									r.operationGroup = ""
									r.pathPattern = "/v2/chat_rooms/{id}"
									r.args = args
									r.count = 1
									return r, true
								default:
									return
								}
							}
							switch elem[0] {
							case '/': // Prefix: "/"

								if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									break
								}
								switch elem[0] {
								case 'b': // Prefix: "background"

									if l := len("background"); len(elem) >= l && elem[0:l] == "background" {
										elem = elem[l:]
									} else {
										break
									}

									if len(elem) == 0 {
										// Leaf node.
										switch method {
										case "DELETE":
											r.name = RemoveChatRoomBackgroundOperation
											r.summary = ""
											r.operationID = "removeChatRoomBackground"
											r.operationGroup = ""
											r.pathPattern = "/v2/chat_rooms/{id}/background"
											r.args = args
											r.count = 1
											return r, true
										default:
											return
										}
									}

								case 'i': // Prefix: "invite"

									if l := len("invite"); len(elem) >= l && elem[0:l] == "invite" {
										elem = elem[l:]
									} else {
										break
									}

									if len(elem) == 0 {
										// Leaf node.
										switch method {
										case "POST":
											r.name = InviteToChatRoomOperation
											r.summary = ""
											r.operationID = "inviteToChatRoom"
											r.operationGroup = ""
											r.pathPattern = "/v2/chat_rooms/{id}/invite"
											r.args = args
											r.count = 1
											return r, true
										default:
											return
										}
									}

								case 'k': // Prefix: "kick"

									if l := len("kick"); len(elem) >= l && elem[0:l] == "kick" {
										elem = elem[l:]
									} else {
										break
									}

									if len(elem) == 0 {
										// Leaf node.
										switch method {
										case "POST":
											r.name = KickFromChatRoomOperation
											r.summary = ""
											r.operationID = "kickFromChatRoom"
											r.operationGroup = ""
											r.pathPattern = "/v2/chat_rooms/{id}/kick"
											r.args = args
											r.count = 1
											return r, true
										default:
											return
										}
									}

								case 'm': // Prefix: "messages"

									if l := len("messages"); len(elem) >= l && elem[0:l] == "messages" {
										elem = elem[l:]
									} else {
										break
									}

									if len(elem) == 0 {
										switch method {
										case "GET":
											r.name = GetChatMessagesOperation
											r.summary = ""
											r.operationID = "getChatMessages"
											r.operationGroup = ""
											r.pathPattern = "/v2/chat_rooms/{id}/messages"
											r.args = args
											r.count = 1
											return r, true
										default:
											return
										}
									}
									switch elem[0] {
									case '/': // Prefix: "/"

										if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
											elem = elem[l:]
										} else {
											break
										}

										// Param: "message_id"
										// Match until "/"
										idx := strings.IndexByte(elem, '/')
										if idx < 0 {
											idx = len(elem)
										}
										args[1] = elem[:idx]
										elem = elem[idx:]

										if len(elem) == 0 {
											break
										}
										switch elem[0] {
										case '/': // Prefix: "/read"

											if l := len("/read"); len(elem) >= l && elem[0:l] == "/read" {
												elem = elem[l:]
											} else {
												break
											}

											if len(elem) == 0 {
												// Leaf node.
												switch method {
												case "POST":
													r.name = ReadChatMessageOperation
													r.summary = ""
													r.operationID = "readChatMessage"
													r.operationGroup = ""
													r.pathPattern = "/v2/chat_rooms/{id}/messages/{message_id}/read"
													r.args = args
													r.count = 2
													return r, true
												default:
													return
												}
											}

										}

									}

								}

							}

						case 'o': // Prefix: "onversations/"

							if l := len("onversations/"); len(elem) >= l && elem[0:l] == "onversations/" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								break
							}
							switch elem[0] {
							case 'r': // Prefix: "root_posts"
								origElem := elem
								if l := len("root_posts"); len(elem) >= l && elem[0:l] == "root_posts" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch method {
									case "GET":
										r.name = GetRootPostsOperation
										r.summary = ""
										r.operationID = "getRootPosts"
										r.operationGroup = ""
										r.pathPattern = "/v2/conversations/root_posts"
										r.args = args
										r.count = 0
										return r, true
									default:
										return
									}
								}

								elem = origElem
							}
							// Param: "id"
							// Leaf parameter, slashes are prohibited
							idx := strings.IndexByte(elem, '/')
							if idx >= 0 {
								break
							}
							args[0] = elem
							elem = ""

							if len(elem) == 0 {
								// Leaf node.
								switch method {
								case "GET":
									r.name = GetConversationOperation
									r.summary = ""
									r.operationID = "getConversation"
									r.operationGroup = ""
									r.pathPattern = "/v2/conversations/{id}"
									r.args = args
									r.count = 1
									return r, true
								default:
									return
								}
							}

						}

					case 'g': // Prefix: "g"

						if l := len("g"); len(elem) >= l && elem[0:l] == "g" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							break
						}
						switch elem[0] {
						case 'i': // Prefix: "ifts"

							if l := len("ifts"); len(elem) >= l && elem[0:l] == "ifts" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch method {
								case "GET":
									r.name = ListGiftsOperation
									r.summary = ""
									r.operationID = "listGifts"
									r.operationGroup = ""
									r.pathPattern = "/v2/gifts"
									r.args = args
									r.count = 0
									return r, true
								default:
									return
								}
							}

						case 'r': // Prefix: "roups"

							if l := len("roups"); len(elem) >= l && elem[0:l] == "roups" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								switch method {
								case "GET":
									r.name = ListGroupsOperation
									r.summary = ""
									r.operationID = "listGroups"
									r.operationGroup = ""
									r.pathPattern = "/v2/groups"
									r.args = args
									r.count = 0
									return r, true
								default:
									return
								}
							}
							switch elem[0] {
							case '/': // Prefix: "/"

								if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									break
								}
								switch elem[0] {
								case 'm': // Prefix: "mine"
									origElem := elem
									if l := len("mine"); len(elem) >= l && elem[0:l] == "mine" {
										elem = elem[l:]
									} else {
										break
									}

									if len(elem) == 0 {
										// Leaf node.
										switch method {
										case "GET":
											r.name = ListMyGroupsOperation
											r.summary = ""
											r.operationID = "listMyGroups"
											r.operationGroup = ""
											r.pathPattern = "/v2/groups/mine"
											r.args = args
											r.count = 0
											return r, true
										default:
											return
										}
									}

									elem = origElem
								}
								// Param: "id"
								// Match until "/"
								idx := strings.IndexByte(elem, '/')
								if idx < 0 {
									idx = len(elem)
								}
								args[0] = elem[:idx]
								elem = elem[idx:]

								if len(elem) == 0 {
									break
								}
								switch elem[0] {
								case '/': // Prefix: "/"

									if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
										elem = elem[l:]
									} else {
										break
									}

									if len(elem) == 0 {
										break
									}
									switch elem[0] {
									case 'm': // Prefix: "members"

										if l := len("members"); len(elem) >= l && elem[0:l] == "members" {
											elem = elem[l:]
										} else {
											break
										}

										if len(elem) == 0 {
											// Leaf node.
											switch method {
											case "GET":
												r.name = GetGroupMembersOperation
												r.summary = ""
												r.operationID = "getGroupMembers"
												r.operationGroup = ""
												r.pathPattern = "/v2/groups/{id}/members"
												r.args = args
												r.count = 1
												return r, true
											default:
												return
											}
										}

									case 'p': // Prefix: "posts/search"

										if l := len("posts/search"); len(elem) >= l && elem[0:l] == "posts/search" {
											elem = elem[l:]
										} else {
											break
										}

										if len(elem) == 0 {
											// Leaf node.
											switch method {
											case "GET":
												r.name = SearchGroupPostsOperation
												r.summary = ""
												r.operationID = "searchGroupPosts"
												r.operationGroup = ""
												r.pathPattern = "/v2/groups/{id}/posts/search"
												r.args = args
												r.count = 1
												return r, true
											default:
												return
											}
										}

									}

								}

							}

						}

					case 'n': // Prefix: "notification_settings/"

						if l := len("notification_settings/"); len(elem) >= l && elem[0:l] == "notification_settings/" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							break
						}
						switch elem[0] {
						case 'c': // Prefix: "chat_rooms/"

							if l := len("chat_rooms/"); len(elem) >= l && elem[0:l] == "chat_rooms/" {
								elem = elem[l:]
							} else {
								break
							}

							// Param: "id"
							// Leaf parameter, slashes are prohibited
							idx := strings.IndexByte(elem, '/')
							if idx >= 0 {
								break
							}
							args[0] = elem
							elem = ""

							if len(elem) == 0 {
								// Leaf node.
								switch method {
								case "POST":
									r.name = UpdateChatRoomNotificationSettingsOperation
									r.summary = ""
									r.operationID = "updateChatRoomNotificationSettings"
									r.operationGroup = ""
									r.pathPattern = "/v2/notification_settings/chat_rooms/{id}"
									r.args = args
									r.count = 1
									return r, true
								default:
									return
								}
							}

						case 'g': // Prefix: "groups/"

							if l := len("groups/"); len(elem) >= l && elem[0:l] == "groups/" {
								elem = elem[l:]
							} else {
								break
							}

							// Param: "id"
							// Leaf parameter, slashes are prohibited
							idx := strings.IndexByte(elem, '/')
							if idx >= 0 {
								break
							}
							args[0] = elem
							elem = ""

							if len(elem) == 0 {
								// Leaf node.
								switch method {
								case "GET":
									r.name = GetGroupNotificationSettingsOperation
									r.summary = ""
									r.operationID = "getGroupNotificationSettings"
									r.operationGroup = ""
									r.pathPattern = "/v2/notification_settings/groups/{id}"
									r.args = args
									r.count = 1
									return r, true
								case "POST":
									r.name = UpdateGroupNotificationSettingsOperation
									r.summary = ""
									r.operationID = "updateGroupNotificationSettings"
									r.operationGroup = ""
									r.pathPattern = "/v2/notification_settings/groups/{id}"
									r.args = args
									r.count = 1
									return r, true
								default:
									return
								}
							}

						}

					case 'p': // Prefix: "posts/"

						if l := len("posts/"); len(elem) >= l && elem[0:l] == "posts/" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							break
						}
						switch elem[0] {
						case 'c': // Prefix: "call_"
							origElem := elem
							if l := len("call_"); len(elem) >= l && elem[0:l] == "call_" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								break
							}
							switch elem[0] {
							case 'f': // Prefix: "followers_timeline"

								if l := len("followers_timeline"); len(elem) >= l && elem[0:l] == "followers_timeline" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch method {
									case "GET":
										r.name = GetCallFollowersTimelineOperation
										r.summary = ""
										r.operationID = "getCallFollowersTimeline"
										r.operationGroup = ""
										r.pathPattern = "/v2/posts/call_followers_timeline"
										r.args = args
										r.count = 0
										return r, true
									default:
										return
									}
								}

							case 't': // Prefix: "timeline"

								if l := len("timeline"); len(elem) >= l && elem[0:l] == "timeline" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch method {
									case "GET":
										r.name = GetCallTimelineOperation
										r.summary = ""
										r.operationID = "getCallTimeline"
										r.operationGroup = ""
										r.pathPattern = "/v2/posts/call_timeline"
										r.args = args
										r.count = 0
										return r, true
									default:
										return
									}
								}

							}

							elem = origElem
						case 'f': // Prefix: "following_timeline"
							origElem := elem
							if l := len("following_timeline"); len(elem) >= l && elem[0:l] == "following_timeline" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch method {
								case "GET":
									r.name = GetFollowingTimelineOperation
									r.summary = ""
									r.operationID = "getFollowingTimeline"
									r.operationGroup = ""
									r.pathPattern = "/v2/posts/following_timeline"
									r.args = args
									r.count = 0
									return r, true
								default:
									return
								}
							}

							elem = origElem
						case 'g': // Prefix: "group_"
							origElem := elem
							if l := len("group_"); len(elem) >= l && elem[0:l] == "group_" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								break
							}
							switch elem[0] {
							case 'p': // Prefix: "pinned_post"

								if l := len("pinned_post"); len(elem) >= l && elem[0:l] == "pinned_post" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch method {
									case "DELETE":
										r.name = UnpinGroupPostOperation
										r.summary = ""
										r.operationID = "unpinGroupPost"
										r.operationGroup = ""
										r.pathPattern = "/v2/posts/group_pinned_post"
										r.args = args
										r.count = 0
										return r, true
									case "PUT":
										r.name = PinGroupPostOperation
										r.summary = ""
										r.operationID = "pinGroupPost"
										r.operationGroup = ""
										r.pathPattern = "/v2/posts/group_pinned_post"
										r.args = args
										r.count = 0
										return r, true
									default:
										return
									}
								}

							case 't': // Prefix: "timeline"

								if l := len("timeline"); len(elem) >= l && elem[0:l] == "timeline" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch method {
									case "GET":
										r.name = GetGroupTimelineOperation
										r.summary = ""
										r.operationID = "getGroupTimeline"
										r.operationGroup = ""
										r.pathPattern = "/v2/posts/group_timeline"
										r.args = args
										r.count = 0
										return r, true
									default:
										return
									}
								}

							}

							elem = origElem
						case 'l': // Prefix: "like"
							origElem := elem
							if l := len("like"); len(elem) >= l && elem[0:l] == "like" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch method {
								case "POST":
									r.name = LikePostsOperation
									r.summary = ""
									r.operationID = "likePosts"
									r.operationGroup = ""
									r.pathPattern = "/v2/posts/like"
									r.args = args
									r.count = 0
									return r, true
								default:
									return
								}
							}

							elem = origElem
						case 'm': // Prefix: "m"
							origElem := elem
							if l := len("m"); len(elem) >= l && elem[0:l] == "m" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								break
							}
							switch elem[0] {
							case 'a': // Prefix: "ass_destroy"

								if l := len("ass_destroy"); len(elem) >= l && elem[0:l] == "ass_destroy" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch method {
									case "POST":
										r.name = DeletePostsOperation
										r.summary = ""
										r.operationID = "deletePosts"
										r.operationGroup = ""
										r.pathPattern = "/v2/posts/mass_destroy"
										r.args = args
										r.count = 0
										return r, true
									default:
										return
									}
								}

							case 'i': // Prefix: "ine"

								if l := len("ine"); len(elem) >= l && elem[0:l] == "ine" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch method {
									case "GET":
										r.name = GetMyPostsOperation
										r.summary = ""
										r.operationID = "getMyPosts"
										r.operationGroup = ""
										r.pathPattern = "/v2/posts/mine"
										r.args = args
										r.count = 0
										return r, true
									default:
										return
									}
								}

							case 'u': // Prefix: "ultiple"

								if l := len("ultiple"); len(elem) >= l && elem[0:l] == "ultiple" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch method {
									case "GET":
										r.name = GetPostsByIdsOperation
										r.summary = ""
										r.operationID = "getPostsByIds"
										r.operationGroup = ""
										r.pathPattern = "/v2/posts/multiple"
										r.args = args
										r.count = 0
										return r, true
									default:
										return
									}
								}

							}

							elem = origElem
						case 'n': // Prefix: "new_"
							origElem := elem
							if l := len("new_"); len(elem) >= l && elem[0:l] == "new_" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								break
							}
							switch elem[0] {
							case 'c': // Prefix: "conference_call"

								if l := len("conference_call"); len(elem) >= l && elem[0:l] == "conference_call" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch method {
									case "POST":
										r.name = CreateConferenceCallPostOperation
										r.summary = ""
										r.operationID = "createConferenceCallPost"
										r.operationGroup = ""
										r.pathPattern = "/v2/posts/new_conference_call"
										r.args = args
										r.count = 0
										return r, true
									default:
										return
									}
								}

							case 's': // Prefix: "share_post"

								if l := len("share_post"); len(elem) >= l && elem[0:l] == "share_post" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch method {
									case "POST":
										r.name = CreateSharePostOperation
										r.summary = ""
										r.operationID = "createSharePost"
										r.operationGroup = ""
										r.pathPattern = "/v2/posts/new_share_post"
										r.args = args
										r.count = 0
										return r, true
									default:
										return
									}
								}

							}

							elem = origElem
						case 'r': // Prefix: "rec"
							origElem := elem
							if l := len("rec"); len(elem) >= l && elem[0:l] == "rec" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								break
							}
							switch elem[0] {
							case 'e': // Prefix: "ent_engagement"

								if l := len("ent_engagement"); len(elem) >= l && elem[0:l] == "ent_engagement" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch method {
									case "GET":
										r.name = GetRecentEngagementPostsOperation
										r.summary = ""
										r.operationID = "getRecentEngagementPosts"
										r.operationGroup = ""
										r.pathPattern = "/v2/posts/recent_engagement"
										r.args = args
										r.count = 0
										return r, true
									default:
										return
									}
								}

							case 'o': // Prefix: "ommended_timeline"

								if l := len("ommended_timeline"); len(elem) >= l && elem[0:l] == "ommended_timeline" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch method {
									case "GET":
										r.name = GetRecommendedTimelineOperation
										r.summary = ""
										r.operationID = "getRecommendedTimeline"
										r.operationGroup = ""
										r.pathPattern = "/v2/posts/recommended_timeline"
										r.args = args
										r.count = 0
										return r, true
									default:
										return
									}
								}

							}

							elem = origElem
						case 's': // Prefix: "search"
							origElem := elem
							if l := len("search"); len(elem) >= l && elem[0:l] == "search" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch method {
								case "GET":
									r.name = SearchPostsOperation
									r.summary = ""
									r.operationID = "searchPosts"
									r.operationGroup = ""
									r.pathPattern = "/v2/posts/search"
									r.args = args
									r.count = 0
									return r, true
								default:
									return
								}
							}

							elem = origElem
						case 't': // Prefix: "tags/"
							origElem := elem
							if l := len("tags/"); len(elem) >= l && elem[0:l] == "tags/" {
								elem = elem[l:]
							} else {
								break
							}

							// Param: "tag"
							// Leaf parameter, slashes are prohibited
							idx := strings.IndexByte(elem, '/')
							if idx >= 0 {
								break
							}
							args[0] = elem
							elem = ""

							if len(elem) == 0 {
								// Leaf node.
								switch method {
								case "GET":
									r.name = GetPostsByTagOperation
									r.summary = ""
									r.operationID = "getPostsByTag"
									r.operationGroup = ""
									r.pathPattern = "/v2/posts/tags/{tag}"
									r.args = args
									r.count = 1
									return r, true
								default:
									return
								}
							}

							elem = origElem
						case 'u': // Prefix: "u"
							origElem := elem
							if l := len("u"); len(elem) >= l && elem[0:l] == "u" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								break
							}
							switch elem[0] {
							case 'r': // Prefix: "rl_metadata"

								if l := len("rl_metadata"); len(elem) >= l && elem[0:l] == "rl_metadata" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch method {
									case "GET":
										r.name = GetPostUrlMetadataOperation
										r.summary = ""
										r.operationID = "getPostUrlMetadata"
										r.operationGroup = ""
										r.pathPattern = "/v2/posts/url_metadata"
										r.args = args
										r.count = 0
										return r, true
									default:
										return
									}
								}

							case 's': // Prefix: "ser_timeline"

								if l := len("ser_timeline"); len(elem) >= l && elem[0:l] == "ser_timeline" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch method {
									case "GET":
										r.name = GetUserTimelineOperation
										r.summary = ""
										r.operationID = "getUserTimeline"
										r.operationGroup = ""
										r.pathPattern = "/v2/posts/user_timeline"
										r.args = args
										r.count = 0
										return r, true
									default:
										return
									}
								}

							}

							elem = origElem
						}
						// Param: "id"
						// Match until one of "/t"
						idx := strings.IndexAny(elem, "/t")
						if idx < 0 {
							idx = len(elem)
						}
						args[0] = elem[:idx]
						elem = elem[idx:]

						if len(elem) == 0 {
							switch method {
							case "GET":
								r.name = GetPostOperation
								r.summary = ""
								r.operationID = "getPost"
								r.operationGroup = ""
								r.pathPattern = "/v2/posts/{id}"
								r.args = args
								r.count = 1
								return r, true
							default:
								return
							}
						}
						switch elem[0] {
						case '/': // Prefix: "/reposts"

							if l := len("/reposts"); len(elem) >= l && elem[0:l] == "/reposts" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch method {
								case "GET":
									r.name = GetPostRepostsOperation
									r.summary = ""
									r.operationID = "getPostReposts"
									r.operationGroup = ""
									r.pathPattern = "/v2/posts/{id}/reposts"
									r.args = args
									r.count = 1
									return r, true
								default:
									return
								}
							}

						case 't': // Prefix: "timeline"

							if l := len("timeline"); len(elem) >= l && elem[0:l] == "timeline" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch method {
								case "GET":
									r.name = GetTimelineOperation
									r.summary = ""
									r.operationID = "getTimeline"
									r.operationGroup = ""
									r.pathPattern = "/v2/posts/{noreply_mode}timeline"
									r.args = args
									r.count = 1
									return r, true
								default:
									return
								}
							}

						}

					case 'r': // Prefix: "received_gifts"

						if l := len("received_gifts"); len(elem) >= l && elem[0:l] == "received_gifts" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch method {
							case "GET":
								r.name = ListReceivedGiftsOperation
								r.summary = ""
								r.operationID = "listReceivedGifts"
								r.operationGroup = ""
								r.pathPattern = "/v2/received_gifts"
								r.args = args
								r.count = 0
								return r, true
							default:
								return
							}
						}

					case 's': // Prefix: "s"

						if l := len("s"); len(elem) >= l && elem[0:l] == "s" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							break
						}
						switch elem[0] {
						case 't': // Prefix: "ticker_packs"

							if l := len("ticker_packs"); len(elem) >= l && elem[0:l] == "ticker_packs" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch method {
								case "GET":
									r.name = ListStickerPacksOperation
									r.summary = ""
									r.operationID = "listStickerPacks"
									r.operationGroup = ""
									r.pathPattern = "/v2/sticker_packs"
									r.args = args
									r.count = 0
									return r, true
								default:
									return
								}
							}

						case 'u': // Prefix: "urveys/"

							if l := len("urveys/"); len(elem) >= l && elem[0:l] == "urveys/" {
								elem = elem[l:]
							} else {
								break
							}

							// Param: "id"
							// Match until "/"
							idx := strings.IndexByte(elem, '/')
							if idx < 0 {
								idx = len(elem)
							}
							args[0] = elem[:idx]
							elem = elem[idx:]

							if len(elem) == 0 {
								break
							}
							switch elem[0] {
							case '/': // Prefix: "/vote"

								if l := len("/vote"); len(elem) >= l && elem[0:l] == "/vote" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch method {
									case "POST":
										r.name = VoteSurveyOperation
										r.summary = ""
										r.operationID = "voteSurvey"
										r.operationGroup = ""
										r.pathPattern = "/v2/surveys/{id}/vote"
										r.args = args
										r.count = 1
										return r, true
									default:
										return
									}
								}

							}

						}

					case 'u': // Prefix: "users/"

						if l := len("users/"); len(elem) >= l && elem[0:l] == "users/" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							break
						}
						switch elem[0] {
						case 'b': // Prefix: "blocked"
							origElem := elem
							if l := len("blocked"); len(elem) >= l && elem[0:l] == "blocked" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch method {
								case "POST":
									r.name = GetBlockedUsersOperation
									r.summary = ""
									r.operationID = "getBlockedUsers"
									r.operationGroup = ""
									r.pathPattern = "/v2/users/blocked"
									r.args = args
									r.count = 0
									return r, true
								default:
									return
								}
							}

							elem = origElem
						case 'e': // Prefix: "edit"
							origElem := elem
							if l := len("edit"); len(elem) >= l && elem[0:l] == "edit" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch method {
								case "POST":
									r.name = EditUserV2Operation
									r.summary = ""
									r.operationID = "editUserV2"
									r.operationGroup = ""
									r.pathPattern = "/v2/users/edit"
									r.args = args
									r.count = 0
									return r, true
								default:
									return
								}
							}

							elem = origElem
						case 'f': // Prefix: "f"
							origElem := elem
							if l := len("f"); len(elem) >= l && elem[0:l] == "f" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								break
							}
							switch elem[0] {
							case 'o': // Prefix: "ollow"

								if l := len("ollow"); len(elem) >= l && elem[0:l] == "ollow" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									switch method {
									case "POST":
										r.name = FollowUsersOperation
										r.summary = ""
										r.operationID = "followUsers"
										r.operationGroup = ""
										r.pathPattern = "/v2/users/follow"
										r.args = args
										r.count = 0
										return r, true
									default:
										return
									}
								}
								switch elem[0] {
								case '_': // Prefix: "_requests"

									if l := len("_requests"); len(elem) >= l && elem[0:l] == "_requests" {
										elem = elem[l:]
									} else {
										break
									}

									if len(elem) == 0 {
										switch method {
										case "GET":
											r.name = GetFollowRequestsOperation
											r.summary = ""
											r.operationID = "getFollowRequests"
											r.operationGroup = ""
											r.pathPattern = "/v2/users/follow_requests"
											r.args = args
											r.count = 0
											return r, true
										default:
											return
										}
									}
									switch elem[0] {
									case '_': // Prefix: "_count"

										if l := len("_count"); len(elem) >= l && elem[0:l] == "_count" {
											elem = elem[l:]
										} else {
											break
										}

										if len(elem) == 0 {
											// Leaf node.
											switch method {
											case "GET":
												r.name = GetFollowRequestCountOperation
												r.summary = ""
												r.operationID = "getFollowRequestCount"
												r.operationGroup = ""
												r.pathPattern = "/v2/users/follow_requests_count"
												r.args = args
												r.count = 0
												return r, true
											default:
												return
											}
										}

									}

								}

							case 'r': // Prefix: "resh/"

								if l := len("resh/"); len(elem) >= l && elem[0:l] == "resh/" {
									elem = elem[l:]
								} else {
									break
								}

								// Param: "id"
								// Leaf parameter, slashes are prohibited
								idx := strings.IndexByte(elem, '/')
								if idx >= 0 {
									break
								}
								args[0] = elem
								elem = ""

								if len(elem) == 0 {
									// Leaf node.
									switch method {
									case "GET":
										r.name = GetFreshUserOperation
										r.summary = ""
										r.operationID = "getFreshUser"
										r.operationGroup = ""
										r.pathPattern = "/v2/users/fresh/{id}"
										r.args = args
										r.count = 1
										return r, true
									default:
										return
									}
								}

							}

							elem = origElem
						case 'h': // Prefix: "hima_users"
							origElem := elem
							if l := len("hima_users"); len(elem) >= l && elem[0:l] == "hima_users" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch method {
								case "GET":
									r.name = GetHimaUsersOperation
									r.summary = ""
									r.operationID = "getHimaUsers"
									r.operationGroup = ""
									r.pathPattern = "/v2/users/hima_users"
									r.args = args
									r.count = 0
									return r, true
								default:
									return
								}
							}

							elem = origElem
						case 'i': // Prefix: "info/"
							origElem := elem
							if l := len("info/"); len(elem) >= l && elem[0:l] == "info/" {
								elem = elem[l:]
							} else {
								break
							}

							// Param: "id"
							// Leaf parameter, slashes are prohibited
							idx := strings.IndexByte(elem, '/')
							if idx >= 0 {
								break
							}
							args[0] = elem
							elem = ""

							if len(elem) == 0 {
								// Leaf node.
								switch method {
								case "GET":
									r.name = GetUserInfoOperation
									r.summary = ""
									r.operationID = "getUserInfo"
									r.operationGroup = ""
									r.pathPattern = "/v2/users/info/{id}"
									r.args = args
									r.count = 1
									return r, true
								default:
									return
								}
							}

							elem = origElem
						case 'r': // Prefix: "re"
							origElem := elem
							if l := len("re"); len(elem) >= l && elem[0:l] == "re" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								break
							}
							switch elem[0] {
							case 'm': // Prefix: "move_"

								if l := len("move_"); len(elem) >= l && elem[0:l] == "move_" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									break
								}
								switch elem[0] {
								case 'c': // Prefix: "cover_image"

									if l := len("cover_image"); len(elem) >= l && elem[0:l] == "cover_image" {
										elem = elem[l:]
									} else {
										break
									}

									if len(elem) == 0 {
										// Leaf node.
										switch method {
										case "POST":
											r.name = RemoveCoverImageOperation
											r.summary = ""
											r.operationID = "removeCoverImage"
											r.operationGroup = ""
											r.pathPattern = "/v2/users/remove_cover_image"
											r.args = args
											r.count = 0
											return r, true
										default:
											return
										}
									}

								case 'p': // Prefix: "profile_photo"

									if l := len("profile_photo"); len(elem) >= l && elem[0:l] == "profile_photo" {
										elem = elem[l:]
									} else {
										break
									}

									if len(elem) == 0 {
										// Leaf node.
										switch method {
										case "POST":
											r.name = RemoveProfilePhotoOperation
											r.summary = ""
											r.operationID = "removeProfilePhoto"
											r.operationGroup = ""
											r.pathPattern = "/v2/users/remove_profile_photo"
											r.args = args
											r.count = 0
											return r, true
										default:
											return
										}
									}

								}

							case 's': // Prefix: "send_confirm_email"

								if l := len("send_confirm_email"); len(elem) >= l && elem[0:l] == "send_confirm_email" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch method {
									case "POST":
										r.name = ResendConfirmEmailOperation
										r.summary = ""
										r.operationID = "resendConfirmEmail"
										r.operationGroup = ""
										r.pathPattern = "/v2/users/resend_confirm_email"
										r.args = args
										r.count = 0
										return r, true
									default:
										return
									}
								}

							case 'v': // Prefix: "views/"

								if l := len("views/"); len(elem) >= l && elem[0:l] == "views/" {
									elem = elem[l:]
								} else {
									break
								}

								// Param: "id"
								// Leaf parameter, slashes are prohibited
								idx := strings.IndexByte(elem, '/')
								if idx >= 0 {
									break
								}
								args[0] = elem
								elem = ""

								if len(elem) == 0 {
									// Leaf node.
									switch method {
									case "GET":
										r.name = GetUserReviewsOperation
										r.summary = ""
										r.operationID = "getUserReviews"
										r.operationGroup = ""
										r.pathPattern = "/v2/users/reviews/{id}"
										r.args = args
										r.count = 1
										return r, true
									case "POST":
										r.name = ReplyToReviewOperation
										r.summary = ""
										r.operationID = "replyToReview"
										r.operationGroup = ""
										r.pathPattern = "/v2/users/reviews/{id}"
										r.args = args
										r.count = 1
										return r, true
									default:
										return
									}
								}

							}

							elem = origElem
						case 't': // Prefix: "timestamp"
							origElem := elem
							if l := len("timestamp"); len(elem) >= l && elem[0:l] == "timestamp" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch method {
								case "GET":
									r.name = GetUserTimestampOperation
									r.summary = ""
									r.operationID = "getUserTimestamp"
									r.operationGroup = ""
									r.pathPattern = "/v2/users/timestamp"
									r.args = args
									r.count = 0
									return r, true
								default:
									return
								}
							}

							elem = origElem
						}
						// Param: "user_id"
						// Match until "/"
						idx := strings.IndexByte(elem, '/')
						if idx < 0 {
							idx = len(elem)
						}
						args[0] = elem[:idx]
						elem = elem[idx:]

						if len(elem) == 0 {
							switch method {
							case "GET":
								r.name = GetUserOperation
								r.summary = ""
								r.operationID = "getUser"
								r.operationGroup = ""
								r.pathPattern = "/v2/users/{id}"
								r.args = args
								r.count = 1
								return r, true
							default:
								return
							}
						}
						switch elem[0] {
						case '/': // Prefix: "/"

							if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								break
							}
							switch elem[0] {
							case 'f': // Prefix: "fo"

								if l := len("fo"); len(elem) >= l && elem[0:l] == "fo" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									break
								}
								switch elem[0] {
								case 'l': // Prefix: "llow"

									if l := len("llow"); len(elem) >= l && elem[0:l] == "llow" {
										elem = elem[l:]
									} else {
										break
									}

									if len(elem) == 0 {
										switch method {
										case "POST":
											r.name = FollowUserOperation
											r.summary = ""
											r.operationID = "followUser"
											r.operationGroup = ""
											r.pathPattern = "/v2/users/{id}/follow"
											r.args = args
											r.count = 1
											return r, true
										default:
											return
										}
									}
									switch elem[0] {
									case '_': // Prefix: "_request"

										if l := len("_request"); len(elem) >= l && elem[0:l] == "_request" {
											elem = elem[l:]
										} else {
											break
										}

										if len(elem) == 0 {
											// Leaf node.
											switch method {
											case "POST":
												r.name = RequestFollowOperation
												r.summary = ""
												r.operationID = "requestFollow"
												r.operationGroup = ""
												r.pathPattern = "/v2/users/{target_id}/follow_request"
												r.args = args
												r.count = 1
												return r, true
											default:
												return
											}
										}

									}

								case 'o': // Prefix: "otprints/"

									if l := len("otprints/"); len(elem) >= l && elem[0:l] == "otprints/" {
										elem = elem[l:]
									} else {
										break
									}

									// Param: "footprint_id"
									// Leaf parameter, slashes are prohibited
									idx := strings.IndexByte(elem, '/')
									if idx >= 0 {
										break
									}
									args[1] = elem
									elem = ""

									if len(elem) == 0 {
										// Leaf node.
										switch method {
										case "DELETE":
											r.name = DeleteFootprintOperation
											r.summary = ""
											r.operationID = "deleteFootprint"
											r.operationGroup = ""
											r.pathPattern = "/v2/users/{user_id}/footprints/{footprint_id}"
											r.args = args
											r.count = 2
											return r, true
										default:
											return
										}
									}

								}

							case 'u': // Prefix: "un"

								if l := len("un"); len(elem) >= l && elem[0:l] == "un" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									break
								}
								switch elem[0] {
								case 'b': // Prefix: "block"

									if l := len("block"); len(elem) >= l && elem[0:l] == "block" {
										elem = elem[l:]
									} else {
										break
									}

									if len(elem) == 0 {
										// Leaf node.
										switch method {
										case "POST":
											r.name = UnblockUserOperation
											r.summary = ""
											r.operationID = "unblockUser"
											r.operationGroup = ""
											r.pathPattern = "/v2/users/{id}/unblock"
											r.args = args
											r.count = 1
											return r, true
										default:
											return
										}
									}

								case 'f': // Prefix: "follow"

									if l := len("follow"); len(elem) >= l && elem[0:l] == "follow" {
										elem = elem[l:]
									} else {
										break
									}

									if len(elem) == 0 {
										// Leaf node.
										switch method {
										case "POST":
											r.name = UnfollowUserOperation
											r.summary = ""
											r.operationID = "unfollowUser"
											r.operationGroup = ""
											r.pathPattern = "/v2/users/{id}/unfollow"
											r.args = args
											r.count = 1
											return r, true
										default:
											return
										}
									}

								}

							}

						}

					}

				case '3': // Prefix: "3/"

					if l := len("3/"); len(elem) >= l && elem[0:l] == "3/" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						break
					}
					switch elem[0] {
					case 'c': // Prefix: "c"

						if l := len("c"); len(elem) >= l && elem[0:l] == "c" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							break
						}
						switch elem[0] {
						case 'a': // Prefix: "alls/"

							if l := len("alls/"); len(elem) >= l && elem[0:l] == "alls/" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								break
							}
							switch elem[0] {
							case 'c': // Prefix: "conference_calls/"
								origElem := elem
								if l := len("conference_calls/"); len(elem) >= l && elem[0:l] == "conference_calls/" {
									elem = elem[l:]
								} else {
									break
								}

								// Param: "call_id"
								// Match until "/"
								idx := strings.IndexByte(elem, '/')
								if idx < 0 {
									idx = len(elem)
								}
								args[0] = elem[:idx]
								elem = elem[idx:]

								if len(elem) == 0 {
									break
								}
								switch elem[0] {
								case '/': // Prefix: "/kick"

									if l := len("/kick"); len(elem) >= l && elem[0:l] == "/kick" {
										elem = elem[l:]
									} else {
										break
									}

									if len(elem) == 0 {
										// Leaf node.
										switch method {
										case "POST":
											r.name = KickFromConferenceCallOperation
											r.summary = ""
											r.operationID = "kickFromConferenceCall"
											r.operationGroup = ""
											r.pathPattern = "/v3/calls/conference_calls/{call_id}/kick"
											r.args = args
											r.count = 1
											return r, true
										default:
											return
										}
									}

								}

								elem = origElem
							}
							// Param: "call_id"
							// Match until "/"
							idx := strings.IndexByte(elem, '/')
							if idx < 0 {
								idx = len(elem)
							}
							args[0] = elem[:idx]
							elem = elem[idx:]

							if len(elem) == 0 {
								break
							}
							switch elem[0] {
							case '/': // Prefix: "/agora_rtm_token"

								if l := len("/agora_rtm_token"); len(elem) >= l && elem[0:l] == "/agora_rtm_token" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch method {
									case "GET":
										r.name = GetAgoraRtmTokenOperation
										r.summary = ""
										r.operationID = "getAgoraRtmToken"
										r.operationGroup = ""
										r.pathPattern = "/v3/calls/{call_id}/agora_rtm_token"
										r.args = args
										r.count = 1
										return r, true
									default:
										return
									}
								}

							}

						case 'h': // Prefix: "hat_rooms/"

							if l := len("hat_rooms/"); len(elem) >= l && elem[0:l] == "hat_rooms/" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								break
							}
							switch elem[0] {
							case 'n': // Prefix: "new"
								origElem := elem
								if l := len("new"); len(elem) >= l && elem[0:l] == "new" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch method {
									case "POST":
										r.name = CreateChatRoomOperation
										r.summary = ""
										r.operationID = "createChatRoom"
										r.operationGroup = ""
										r.pathPattern = "/v3/chat_rooms/new"
										r.args = args
										r.count = 0
										return r, true
									default:
										return
									}
								}

								elem = origElem
							}
							// Param: "chat_room_id"
							// Match until "/"
							idx := strings.IndexByte(elem, '/')
							if idx < 0 {
								idx = len(elem)
							}
							args[0] = elem[:idx]
							elem = elem[idx:]

							if len(elem) == 0 {
								break
							}
							switch elem[0] {
							case '/': // Prefix: "/"

								if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									break
								}
								switch elem[0] {
								case 'e': // Prefix: "edit"

									if l := len("edit"); len(elem) >= l && elem[0:l] == "edit" {
										elem = elem[l:]
									} else {
										break
									}

									if len(elem) == 0 {
										// Leaf node.
										switch method {
										case "POST":
											r.name = UpdateChatRoomOperation
											r.summary = ""
											r.operationID = "updateChatRoom"
											r.operationGroup = ""
											r.pathPattern = "/v3/chat_rooms/{id}/edit"
											r.args = args
											r.count = 1
											return r, true
										default:
											return
										}
									}

								case 'm': // Prefix: "messages/new"

									if l := len("messages/new"); len(elem) >= l && elem[0:l] == "messages/new" {
										elem = elem[l:]
									} else {
										break
									}

									if len(elem) == 0 {
										// Leaf node.
										switch method {
										case "POST":
											r.name = SendChatMessageOperation
											r.summary = ""
											r.operationID = "sendChatMessage"
											r.operationGroup = ""
											r.pathPattern = "/v3/chat_rooms/{id}/messages/new"
											r.args = args
											r.count = 1
											return r, true
										default:
											return
										}
									}

								case 'r': // Prefix: "report"

									if l := len("report"); len(elem) >= l && elem[0:l] == "report" {
										elem = elem[l:]
									} else {
										break
									}

									if len(elem) == 0 {
										// Leaf node.
										switch method {
										case "POST":
											r.name = ReportChatRoomOperation
											r.summary = ""
											r.operationID = "reportChatRoom"
											r.operationGroup = ""
											r.pathPattern = "/v3/chat_rooms/{chat_room_id}/report"
											r.args = args
											r.count = 1
											return r, true
										default:
											return
										}
									}

								}

							}

						}

					case 'g': // Prefix: "groups/"

						if l := len("groups/"); len(elem) >= l && elem[0:l] == "groups/" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							break
						}
						switch elem[0] {
						case 'n': // Prefix: "new"
							origElem := elem
							if l := len("new"); len(elem) >= l && elem[0:l] == "new" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch method {
								case "POST":
									r.name = CreateGroupOperation
									r.summary = ""
									r.operationID = "createGroup"
									r.operationGroup = ""
									r.pathPattern = "/v3/groups/new"
									r.args = args
									r.count = 0
									return r, true
								default:
									return
								}
							}

							elem = origElem
						}
						// Param: "group_id"
						// Match until "/"
						idx := strings.IndexByte(elem, '/')
						if idx < 0 {
							idx = len(elem)
						}
						args[0] = elem[:idx]
						elem = elem[idx:]

						if len(elem) == 0 {
							break
						}
						switch elem[0] {
						case '/': // Prefix: "/"

							if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								break
							}
							switch elem[0] {
							case 'c': // Prefix: "cover"

								if l := len("cover"); len(elem) >= l && elem[0:l] == "cover" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch method {
									case "DELETE":
										r.name = RemoveGroupCoverOperation
										r.summary = ""
										r.operationID = "removeGroupCover"
										r.operationGroup = ""
										r.pathPattern = "/v3/groups/{id}/cover"
										r.args = args
										r.count = 1
										return r, true
									default:
										return
									}
								}

							case 'd': // Prefix: "deputize/mass"

								if l := len("deputize/mass"); len(elem) >= l && elem[0:l] == "deputize/mass" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch method {
									case "POST":
										r.name = DeputizeGroupUsersMassOperation
										r.summary = ""
										r.operationID = "deputizeGroupUsersMass"
										r.operationGroup = ""
										r.pathPattern = "/v3/groups/{group_id}/deputize/mass"
										r.args = args
										r.count = 1
										return r, true
									default:
										return
									}
								}

							case 'i': // Prefix: "icon"

								if l := len("icon"); len(elem) >= l && elem[0:l] == "icon" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch method {
									case "DELETE":
										r.name = RemoveGroupIconOperation
										r.summary = ""
										r.operationID = "removeGroupIcon"
										r.operationGroup = ""
										r.pathPattern = "/v3/groups/{id}/icon"
										r.args = args
										r.count = 1
										return r, true
									default:
										return
									}
								}

							case 'r': // Prefix: "report"

								if l := len("report"); len(elem) >= l && elem[0:l] == "report" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch method {
									case "POST":
										r.name = ReportGroupOperation
										r.summary = ""
										r.operationID = "reportGroup"
										r.operationGroup = ""
										r.pathPattern = "/v3/groups/{group_id}/report"
										r.args = args
										r.count = 1
										return r, true
									default:
										return
									}
								}

							case 't': // Prefix: "transfer"

								if l := len("transfer"); len(elem) >= l && elem[0:l] == "transfer" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch method {
									case "POST":
										r.name = TransferGroupOperation
										r.summary = ""
										r.operationID = "transferGroup"
										r.operationGroup = ""
										r.pathPattern = "/v3/groups/{id}/transfer"
										r.args = args
										r.count = 1
										return r, true
									default:
										return
									}
								}

							case 'u': // Prefix: "update"

								if l := len("update"); len(elem) >= l && elem[0:l] == "update" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch method {
									case "POST":
										r.name = UpdateGroupOperation
										r.summary = ""
										r.operationID = "updateGroup"
										r.operationGroup = ""
										r.pathPattern = "/v3/groups/{id}/update"
										r.args = args
										r.count = 1
										return r, true
									default:
										return
									}
								}

							}

						}

					case 'p': // Prefix: "posts/"

						if l := len("posts/"); len(elem) >= l && elem[0:l] == "posts/" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							break
						}
						switch elem[0] {
						case 'n': // Prefix: "new"
							origElem := elem
							if l := len("new"); len(elem) >= l && elem[0:l] == "new" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch method {
								case "POST":
									r.name = CreatePostOperation
									r.summary = ""
									r.operationID = "createPost"
									r.operationGroup = ""
									r.pathPattern = "/v3/posts/new"
									r.args = args
									r.count = 0
									return r, true
								default:
									return
								}
							}

							elem = origElem
						case 'r': // Prefix: "repost"
							origElem := elem
							if l := len("repost"); len(elem) >= l && elem[0:l] == "repost" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch method {
								case "POST":
									r.name = RepostOperation
									r.summary = ""
									r.operationID = "repost"
									r.operationGroup = ""
									r.pathPattern = "/v3/posts/repost"
									r.args = args
									r.count = 0
									return r, true
								default:
									return
								}
							}

							elem = origElem
						}
						// Param: "id"
						// Match until "/"
						idx := strings.IndexByte(elem, '/')
						if idx < 0 {
							idx = len(elem)
						}
						args[0] = elem[:idx]
						elem = elem[idx:]

						if len(elem) == 0 {
							switch method {
							case "PUT":
								r.name = UpdatePostOperation
								r.summary = ""
								r.operationID = "updatePost"
								r.operationGroup = ""
								r.pathPattern = "/v3/posts/{id}"
								r.args = args
								r.count = 1
								return r, true
							default:
								return
							}
						}
						switch elem[0] {
						case '/': // Prefix: "/"

							if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								break
							}
							switch elem[0] {
							case 'm': // Prefix: "move_to_thread"

								if l := len("move_to_thread"); len(elem) >= l && elem[0:l] == "move_to_thread" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									switch method {
									case "POST":
										r.name = MovePostToThreadOperation
										r.summary = ""
										r.operationID = "movePostToThread"
										r.operationGroup = ""
										r.pathPattern = "/v3/posts/{id}/move_to_thread"
										r.args = args
										r.count = 1
										return r, true
									default:
										return
									}
								}
								switch elem[0] {
								case '/': // Prefix: "/"

									if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
										elem = elem[l:]
									} else {
										break
									}

									// Param: "thread_id"
									// Leaf parameter, slashes are prohibited
									idx := strings.IndexByte(elem, '/')
									if idx >= 0 {
										break
									}
									args[1] = elem
									elem = ""

									if len(elem) == 0 {
										// Leaf node.
										switch method {
										case "PUT":
											r.name = MovePostToSpecificThreadOperation
											r.summary = ""
											r.operationID = "movePostToSpecificThread"
											r.operationGroup = ""
											r.pathPattern = "/v3/posts/{id}/move_to_thread/{thread_id}"
											r.args = args
											r.count = 2
											return r, true
										default:
											return
										}
									}

								}

							case 'r': // Prefix: "report"

								if l := len("report"); len(elem) >= l && elem[0:l] == "report" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch method {
									case "POST":
										r.name = ReportPostOperation
										r.summary = ""
										r.operationID = "reportPost"
										r.operationGroup = ""
										r.pathPattern = "/v3/posts/{post_id}/report"
										r.args = args
										r.count = 1
										return r, true
									default:
										return
									}
								}

							}

						}

					case 'u': // Prefix: "users/"

						if l := len("users/"); len(elem) >= l && elem[0:l] == "users/" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							break
						}
						switch elem[0] {
						case 'e': // Prefix: "edit"
							origElem := elem
							if l := len("edit"); len(elem) >= l && elem[0:l] == "edit" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch method {
								case "POST":
									r.name = EditUserOperation
									r.summary = ""
									r.operationID = "editUser"
									r.operationGroup = ""
									r.pathPattern = "/v3/users/edit"
									r.args = args
									r.count = 0
									return r, true
								default:
									return
								}
							}

							elem = origElem
						case 'f': // Prefix: "footprints"
							origElem := elem
							if l := len("footprints"); len(elem) >= l && elem[0:l] == "footprints" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch method {
								case "GET":
									r.name = GetFootprintsOperation
									r.summary = ""
									r.operationID = "getFootprints"
									r.operationGroup = ""
									r.pathPattern = "/v3/users/footprints"
									r.args = args
									r.count = 0
									return r, true
								default:
									return
								}
							}

							elem = origElem
						case 'l': // Prefix: "login_"
							origElem := elem
							if l := len("login_"); len(elem) >= l && elem[0:l] == "login_" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								break
							}
							switch elem[0] {
							case 'u': // Prefix: "update"

								if l := len("update"); len(elem) >= l && elem[0:l] == "update" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch method {
									case "POST":
										r.name = UpdateLoginOperation
										r.summary = ""
										r.operationID = "updateLogin"
										r.operationGroup = ""
										r.pathPattern = "/v3/users/login_update"
										r.args = args
										r.count = 0
										return r, true
									default:
										return
									}
								}

							case 'w': // Prefix: "with_email"

								if l := len("with_email"); len(elem) >= l && elem[0:l] == "with_email" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch method {
									case "POST":
										r.name = LoginWithEmailOperation
										r.summary = ""
										r.operationID = "loginWithEmail"
										r.operationGroup = ""
										r.pathPattern = "/v3/users/login_with_email"
										r.args = args
										r.count = 0
										return r, true
									default:
										return
									}
								}

							}

							elem = origElem
						}
						// Param: "id"
						// Match until "/"
						idx := strings.IndexByte(elem, '/')
						if idx < 0 {
							idx = len(elem)
						}
						args[0] = elem[:idx]
						elem = elem[idx:]

						if len(elem) == 0 {
							break
						}
						switch elem[0] {
						case '/': // Prefix: "/"

							if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								break
							}
							switch elem[0] {
							case 'f': // Prefix: "follow"

								if l := len("follow"); len(elem) >= l && elem[0:l] == "follow" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									break
								}
								switch elem[0] {
								case 'e': // Prefix: "ers"

									if l := len("ers"); len(elem) >= l && elem[0:l] == "ers" {
										elem = elem[l:]
									} else {
										break
									}

									if len(elem) == 0 {
										// Leaf node.
										switch method {
										case "GET":
											r.name = GetUserFollowersOperation
											r.summary = ""
											r.operationID = "getUserFollowers"
											r.operationGroup = ""
											r.pathPattern = "/v3/users/{id}/followers"
											r.args = args
											r.count = 1
											return r, true
										default:
											return
										}
									}

								case 'i': // Prefix: "ings"

									if l := len("ings"); len(elem) >= l && elem[0:l] == "ings" {
										elem = elem[l:]
									} else {
										break
									}

									if len(elem) == 0 {
										// Leaf node.
										switch method {
										case "GET":
											r.name = GetUserFollowingsOperation
											r.summary = ""
											r.operationID = "getUserFollowings"
											r.operationGroup = ""
											r.pathPattern = "/v3/users/{id}/followings"
											r.args = args
											r.count = 1
											return r, true
										default:
											return
										}
									}

								}

							case 'r': // Prefix: "report"

								if l := len("report"); len(elem) >= l && elem[0:l] == "report" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch method {
									case "POST":
										r.name = ReportUserOperation
										r.summary = ""
										r.operationID = "reportUser"
										r.operationGroup = ""
										r.pathPattern = "/v3/users/{user_id}/report"
										r.args = args
										r.count = 1
										return r, true
									default:
										return
									}
								}

							}

						}

					}

				}

				elem = origElem
			}
			// Param: "countryApiValue"
			// Match until "/"
			idx := strings.IndexByte(elem, '/')
			if idx < 0 {
				idx = len(elem)
			}
			args[0] = elem[:idx]
			elem = elem[idx:]

			if len(elem) == 0 {
				break
			}
			switch elem[0] {
			case '/': // Prefix: "/api/"

				if l := len("/api/"); len(elem) >= l && elem[0:l] == "/api/" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'a': // Prefix: "apps/"

					if l := len("apps/"); len(elem) >= l && elem[0:l] == "apps/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "app"
					// Match until "/"
					idx := strings.IndexByte(elem, '/')
					if idx < 0 {
						idx = len(elem)
					}
					args[1] = elem[:idx]
					elem = elem[idx:]

					if len(elem) == 0 {
						break
					}
					switch elem[0] {
					case '/': // Prefix: "/popular_words"

						if l := len("/popular_words"); len(elem) >= l && elem[0:l] == "/popular_words" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch method {
							case "GET":
								r.name = GetPopularWordsOperation
								r.summary = ""
								r.operationID = "getPopularWords"
								r.operationGroup = ""
								r.pathPattern = "/{countryApiValue}/api/apps/{app}/popular_words"
								r.args = args
								r.count = 2
								return r, true
							default:
								return
							}
						}

					}

				case 'v': // Prefix: "v2/banned_words"

					if l := len("v2/banned_words"); len(elem) >= l && elem[0:l] == "v2/banned_words" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch method {
						case "GET":
							r.name = GetBannedWordsOperation
							r.summary = ""
							r.operationID = "getBannedWords"
							r.operationGroup = ""
							r.pathPattern = "/{countryApiValue}/api/v2/banned_words"
							r.args = args
							r.count = 1
							return r, true
						default:
							return
						}
					}

				}

			}

		}
	}
	return r, false
}
