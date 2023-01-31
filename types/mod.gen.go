//  Source: lemmy/crates/api_common/src/websocket/mod.rs
// Code generated by go.arsenm.dev/go-lemmy/cmd/gen (struct generator). DO NOT EDIT.

package types

type UserOperation string

const (
	UserOperationLogin                                 UserOperation = "Login"
	UserOperationGetCaptcha                            UserOperation = "GetCaptcha"
	UserOperationSaveComment                           UserOperation = "SaveComment"
	UserOperationCreateCommentLike                     UserOperation = "CreateCommentLike"
	UserOperationCreateCommentReport                   UserOperation = "CreateCommentReport"
	UserOperationResolveCommentReport                  UserOperation = "ResolveCommentReport"
	UserOperationListCommentReports                    UserOperation = "ListCommentReports"
	UserOperationCreatePostLike                        UserOperation = "CreatePostLike"
	UserOperationLockPost                              UserOperation = "LockPost"
	UserOperationFeaturePost                           UserOperation = "FeaturePost"
	UserOperationMarkPostAsRead                        UserOperation = "MarkPostAsRead"
	UserOperationSavePost                              UserOperation = "SavePost"
	UserOperationCreatePostReport                      UserOperation = "CreatePostReport"
	UserOperationResolvePostReport                     UserOperation = "ResolvePostReport"
	UserOperationListPostReports                       UserOperation = "ListPostReports"
	UserOperationGetReportCount                        UserOperation = "GetReportCount"
	UserOperationGetUnreadCount                        UserOperation = "GetUnreadCount"
	UserOperationVerifyEmail                           UserOperation = "VerifyEmail"
	UserOperationFollowCommunity                       UserOperation = "FollowCommunity"
	UserOperationGetReplies                            UserOperation = "GetReplies"
	UserOperationGetPersonMentions                     UserOperation = "GetPersonMentions"
	UserOperationMarkPersonMentionAsRead               UserOperation = "MarkPersonMentionAsRead"
	UserOperationMarkCommentReplyAsRead                UserOperation = "MarkCommentReplyAsRead"
	UserOperationGetModlog                             UserOperation = "GetModlog"
	UserOperationBanFromCommunity                      UserOperation = "BanFromCommunity"
	UserOperationAddModToCommunity                     UserOperation = "AddModToCommunity"
	UserOperationAddAdmin                              UserOperation = "AddAdmin"
	UserOperationGetUnreadRegistrationApplicationCount UserOperation = "GetUnreadRegistrationApplicationCount"
	UserOperationListRegistrationApplications          UserOperation = "ListRegistrationApplications"
	UserOperationApproveRegistrationApplication        UserOperation = "ApproveRegistrationApplication"
	UserOperationBanPerson                             UserOperation = "BanPerson"
	UserOperationGetBannedPersons                      UserOperation = "GetBannedPersons"
	UserOperationMarkAllAsRead                         UserOperation = "MarkAllAsRead"
	UserOperationSaveUserSettings                      UserOperation = "SaveUserSettings"
	UserOperationTransferCommunity                     UserOperation = "TransferCommunity"
	UserOperationLeaveAdmin                            UserOperation = "LeaveAdmin"
	UserOperationPasswordReset                         UserOperation = "PasswordReset"
	UserOperationPasswordChange                        UserOperation = "PasswordChange"
	UserOperationMarkPrivateMessageAsRead              UserOperation = "MarkPrivateMessageAsRead"
	UserOperationCreatePrivateMessageReport            UserOperation = "CreatePrivateMessageReport"
	UserOperationResolvePrivateMessageReport           UserOperation = "ResolvePrivateMessageReport"
	UserOperationListPrivateMessageReports             UserOperation = "ListPrivateMessageReports"
	UserOperationUserJoin                              UserOperation = "UserJoin"
	UserOperationPostJoin                              UserOperation = "PostJoin"
	UserOperationCommunityJoin                         UserOperation = "CommunityJoin"
	UserOperationModJoin                               UserOperation = "ModJoin"
	UserOperationChangePassword                        UserOperation = "ChangePassword"
	UserOperationGetSiteMetadata                       UserOperation = "GetSiteMetadata"
	UserOperationBlockCommunity                        UserOperation = "BlockCommunity"
	UserOperationBlockPerson                           UserOperation = "BlockPerson"
	UserOperationPurgePerson                           UserOperation = "PurgePerson"
	UserOperationPurgeCommunity                        UserOperation = "PurgeCommunity"
	UserOperationPurgePost                             UserOperation = "PurgePost"
	UserOperationPurgeComment                          UserOperation = "PurgeComment"
)

type UserOperationCRUD string

const (
	UserOperationCRUDCreateSite           UserOperationCRUD = "CreateSite"
	UserOperationCRUDGetSite              UserOperationCRUD = "GetSite"
	UserOperationCRUDEditSite             UserOperationCRUD = "EditSite"
	UserOperationCRUDCreateCommunity      UserOperationCRUD = "CreateCommunity"
	UserOperationCRUDListCommunities      UserOperationCRUD = "ListCommunities"
	UserOperationCRUDEditCommunity        UserOperationCRUD = "EditCommunity"
	UserOperationCRUDDeleteCommunity      UserOperationCRUD = "DeleteCommunity"
	UserOperationCRUDRemoveCommunity      UserOperationCRUD = "RemoveCommunity"
	UserOperationCRUDCreatePost           UserOperationCRUD = "CreatePost"
	UserOperationCRUDGetPost              UserOperationCRUD = "GetPost"
	UserOperationCRUDEditPost             UserOperationCRUD = "EditPost"
	UserOperationCRUDDeletePost           UserOperationCRUD = "DeletePost"
	UserOperationCRUDRemovePost           UserOperationCRUD = "RemovePost"
	UserOperationCRUDCreateComment        UserOperationCRUD = "CreateComment"
	UserOperationCRUDGetComment           UserOperationCRUD = "GetComment"
	UserOperationCRUDEditComment          UserOperationCRUD = "EditComment"
	UserOperationCRUDDeleteComment        UserOperationCRUD = "DeleteComment"
	UserOperationCRUDRemoveComment        UserOperationCRUD = "RemoveComment"
	UserOperationCRUDRegister             UserOperationCRUD = "Register"
	UserOperationCRUDDeleteAccount        UserOperationCRUD = "DeleteAccount"
	UserOperationCRUDCreatePrivateMessage UserOperationCRUD = "CreatePrivateMessage"
	UserOperationCRUDGetPrivateMessages   UserOperationCRUD = "GetPrivateMessages"
	UserOperationCRUDEditPrivateMessage   UserOperationCRUD = "EditPrivateMessage"
	UserOperationCRUDDeletePrivateMessage UserOperationCRUD = "DeletePrivateMessage"
)

type UserOperationApub string

const (
	UserOperationApubGetPosts         UserOperationApub = "GetPosts"
	UserOperationApubGetCommunity     UserOperationApub = "GetCommunity"
	UserOperationApubGetComments      UserOperationApub = "GetComments"
	UserOperationApubGetPersonDetails UserOperationApub = "GetPersonDetails"
	UserOperationApubSearch           UserOperationApub = "Search"
	UserOperationApubResolveObject    UserOperationApub = "ResolveObject"
)
