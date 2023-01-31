//  Source: lemmy/crates/db_schema/src/lib.rs
// Code generated by go.arsenm.dev/go-lemmy/cmd/gen (struct generator). DO NOT EDIT.

package types

type SortType string

const (
	SortTypeActive       SortType = "Active"
	SortTypeHot          SortType = "Hot"
	SortTypeNew          SortType = "New"
	SortTypeOld          SortType = "Old"
	SortTypeTopDay       SortType = "TopDay"
	SortTypeTopWeek      SortType = "TopWeek"
	SortTypeTopMonth     SortType = "TopMonth"
	SortTypeTopYear      SortType = "TopYear"
	SortTypeTopAll       SortType = "TopAll"
	SortTypeMostComments SortType = "MostComments"
	SortTypeNewComments  SortType = "NewComments"
)

type CommentSortType string

const (
	CommentSortTypeHot CommentSortType = "Hot"
	CommentSortTypeTop CommentSortType = "Top"
	CommentSortTypeNew CommentSortType = "New"
	CommentSortTypeOld CommentSortType = "Old"
)

type ListingType string

const (
	ListingTypeAll        ListingType = "All"
	ListingTypeLocal      ListingType = "Local"
	ListingTypeSubscribed ListingType = "Subscribed"
)

type SearchType string

const (
	SearchTypeAll         SearchType = "All"
	SearchTypeComments    SearchType = "Comments"
	SearchTypePosts       SearchType = "Posts"
	SearchTypeCommunities SearchType = "Communities"
	SearchTypeUsers       SearchType = "Users"
	SearchTypeUrl         SearchType = "Url"
)

type SubscribedType string

const (
	SubscribedTypeSubscribed    SubscribedType = "Subscribed"
	SubscribedTypeNotSubscribed SubscribedType = "NotSubscribed"
	SubscribedTypePending       SubscribedType = "Pending"
)

type ModlogActionType string

const (
	ModlogActionTypeAll                  ModlogActionType = "All"
	ModlogActionTypeModRemovePost        ModlogActionType = "ModRemovePost"
	ModlogActionTypeModLockPost          ModlogActionType = "ModLockPost"
	ModlogActionTypeModFeaturePost       ModlogActionType = "ModFeaturePost"
	ModlogActionTypeModRemoveComment     ModlogActionType = "ModRemoveComment"
	ModlogActionTypeModRemoveCommunity   ModlogActionType = "ModRemoveCommunity"
	ModlogActionTypeModBanFromCommunity  ModlogActionType = "ModBanFromCommunity"
	ModlogActionTypeModAddCommunity      ModlogActionType = "ModAddCommunity"
	ModlogActionTypeModTransferCommunity ModlogActionType = "ModTransferCommunity"
	ModlogActionTypeModAdd               ModlogActionType = "ModAdd"
	ModlogActionTypeModBan               ModlogActionType = "ModBan"
	ModlogActionTypeModHideCommunity     ModlogActionType = "ModHideCommunity"
	ModlogActionTypeAdminPurgePerson     ModlogActionType = "AdminPurgePerson"
	ModlogActionTypeAdminPurgeCommunity  ModlogActionType = "AdminPurgeCommunity"
	ModlogActionTypeAdminPurgePost       ModlogActionType = "AdminPurgePost"
	ModlogActionTypeAdminPurgeComment    ModlogActionType = "AdminPurgeComment"
)

type PostFeatureType string

const (
	PostFeatureTypeLocal     PostFeatureType = "Local"
	PostFeatureTypeCommunity PostFeatureType = "Community"
)
