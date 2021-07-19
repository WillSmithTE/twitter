package main

type TwitterSearchResponse struct {
	Statuses       []TwitterStatus `json:"statuses"`
	SearchMetadata TwitterMetadata `json:"search_metadata"`
}

type TwitterStatus struct {
	CreatedAt        string `json:"created_at"`
	ID               int64  `json:"id"`
	IDStr            string `json:"id_str"`
	FullText         string `json:"full_text"`
	Truncated        bool   `json:"truncated"`
	DisplayTextRange []int  `json:"display_text_range"`
	Entities         struct {
		Hashtags []struct {
			Text    string `json:"text"`
			Indices []int  `json:"indices"`
		} `json:"hashtags"`
		Symbols      []interface{} `json:"symbols"`
		UserMentions []struct {
			ScreenName string `json:"screen_name"`
			Name       string `json:"name"`
			ID         int64  `json:"id"`
			IDStr      string `json:"id_str"`
			Indices    []int  `json:"indices"`
		} `json:"user_mentions"`
		Urls []interface{} `json:"urls"`
	} `json:"entities"`
	Metadata struct {
		IsoLanguageCode string `json:"iso_language_code"`
		ResultType      string `json:"result_type"`
	} `json:"metadata"`
	Source               string `json:"source"`
	InReplyToStatusID    int64  `json:"in_reply_to_status_id"`
	InReplyToStatusIDStr string `json:"in_reply_to_status_id_str"`
	InReplyToUserID      int64  `json:"in_reply_to_user_id"`
	InReplyToUserIDStr   string `json:"in_reply_to_user_id_str"`
	InReplyToScreenName  string `json:"in_reply_to_screen_name"`
	User                 struct {
		ID          int64       `json:"id"`
		IDStr       string      `json:"id_str"`
		Name        string      `json:"name"`
		ScreenName  string      `json:"screen_name"`
		Location    string      `json:"location"`
		Description string      `json:"description"`
		URL         interface{} `json:"url"`
		Entities    struct {
			Description struct {
				Urls []interface{} `json:"urls"`
			} `json:"description"`
		} `json:"entities"`
		Protected                      bool          `json:"protected"`
		FollowersCount                 int           `json:"followers_count"`
		FriendsCount                   int           `json:"friends_count"`
		ListedCount                    int           `json:"listed_count"`
		CreatedAt                      string        `json:"created_at"`
		FavouritesCount                int           `json:"favourites_count"`
		UtcOffset                      interface{}   `json:"utc_offset"`
		TimeZone                       interface{}   `json:"time_zone"`
		GeoEnabled                     bool          `json:"geo_enabled"`
		Verified                       bool          `json:"verified"`
		StatusesCount                  int           `json:"statuses_count"`
		Lang                           interface{}   `json:"lang"`
		ContributorsEnabled            bool          `json:"contributors_enabled"`
		IsTranslator                   bool          `json:"is_translator"`
		IsTranslationEnabled           bool          `json:"is_translation_enabled"`
		ProfileBackgroundColor         string        `json:"profile_background_color"`
		ProfileBackgroundImageURL      interface{}   `json:"profile_background_image_url"`
		ProfileBackgroundImageURLHTTPS interface{}   `json:"profile_background_image_url_https"`
		ProfileBackgroundTile          bool          `json:"profile_background_tile"`
		ProfileImageURL                string        `json:"profile_image_url"`
		ProfileImageURLHTTPS           string        `json:"profile_image_url_https"`
		ProfileLinkColor               string        `json:"profile_link_color"`
		ProfileSidebarBorderColor      string        `json:"profile_sidebar_border_color"`
		ProfileSidebarFillColor        string        `json:"profile_sidebar_fill_color"`
		ProfileTextColor               string        `json:"profile_text_color"`
		ProfileUseBackgroundImage      bool          `json:"profile_use_background_image"`
		HasExtendedProfile             bool          `json:"has_extended_profile"`
		DefaultProfile                 bool          `json:"default_profile"`
		DefaultProfileImage            bool          `json:"default_profile_image"`
		Following                      interface{}   `json:"following"`
		FollowRequestSent              interface{}   `json:"follow_request_sent"`
		Notifications                  interface{}   `json:"notifications"`
		TranslatorType                 string        `json:"translator_type"`
		WithheldInCountries            []interface{} `json:"withheld_in_countries"`
	} `json:"user"`
	Geo               interface{} `json:"geo"`
	Coordinates       interface{} `json:"coordinates"`
	Place             interface{} `json:"place"`
	Contributors      interface{} `json:"contributors"`
	IsQuoteStatus     bool        `json:"is_quote_status"`
	RetweetCount      int         `json:"retweet_count"`
	FavoriteCount     int         `json:"favorite_count"`
	Favorited         bool        `json:"favorited"`
	Retweeted         bool        `json:"retweeted"`
	Lang              string      `json:"lang"`
	QuotedStatusID    int64       `json:"quoted_status_id,omitempty"`
	QuotedStatusIDStr string      `json:"quoted_status_id_str,omitempty"`
	QuotedStatus      struct {
		CreatedAt        string `json:"created_at"`
		ID               int64  `json:"id"`
		IDStr            string `json:"id_str"`
		FullText         string `json:"full_text"`
		Truncated        bool   `json:"truncated"`
		DisplayTextRange []int  `json:"display_text_range"`
		Entities         struct {
			Hashtags     []interface{} `json:"hashtags"`
			Symbols      []interface{} `json:"symbols"`
			UserMentions []interface{} `json:"user_mentions"`
			Urls         []interface{} `json:"urls"`
		} `json:"entities"`
		Metadata struct {
			IsoLanguageCode string `json:"iso_language_code"`
			ResultType      string `json:"result_type"`
		} `json:"metadata"`
		Source               string `json:"source"`
		InReplyToStatusID    int64  `json:"in_reply_to_status_id"`
		InReplyToStatusIDStr string `json:"in_reply_to_status_id_str"`
		InReplyToUserID      int    `json:"in_reply_to_user_id"`
		InReplyToUserIDStr   string `json:"in_reply_to_user_id_str"`
		InReplyToScreenName  string `json:"in_reply_to_screen_name"`
		User                 struct {
			ID          int    `json:"id"`
			IDStr       string `json:"id_str"`
			Name        string `json:"name"`
			ScreenName  string `json:"screen_name"`
			Location    string `json:"location"`
			Description string `json:"description"`
			URL         string `json:"url"`
			Entities    struct {
				URL struct {
					Urls []struct {
						URL         string `json:"url"`
						ExpandedURL string `json:"expanded_url"`
						DisplayURL  string `json:"display_url"`
						Indices     []int  `json:"indices"`
					} `json:"urls"`
				} `json:"url"`
				Description struct {
					Urls []interface{} `json:"urls"`
				} `json:"description"`
			} `json:"entities"`
			Protected                      bool          `json:"protected"`
			FollowersCount                 int           `json:"followers_count"`
			FriendsCount                   int           `json:"friends_count"`
			ListedCount                    int           `json:"listed_count"`
			CreatedAt                      string        `json:"created_at"`
			FavouritesCount                int           `json:"favourites_count"`
			UtcOffset                      interface{}   `json:"utc_offset"`
			TimeZone                       interface{}   `json:"time_zone"`
			GeoEnabled                     bool          `json:"geo_enabled"`
			Verified                       bool          `json:"verified"`
			StatusesCount                  int           `json:"statuses_count"`
			Lang                           interface{}   `json:"lang"`
			ContributorsEnabled            bool          `json:"contributors_enabled"`
			IsTranslator                   bool          `json:"is_translator"`
			IsTranslationEnabled           bool          `json:"is_translation_enabled"`
			ProfileBackgroundColor         string        `json:"profile_background_color"`
			ProfileBackgroundImageURL      string        `json:"profile_background_image_url"`
			ProfileBackgroundImageURLHTTPS string        `json:"profile_background_image_url_https"`
			ProfileBackgroundTile          bool          `json:"profile_background_tile"`
			ProfileImageURL                string        `json:"profile_image_url"`
			ProfileImageURLHTTPS           string        `json:"profile_image_url_https"`
			ProfileBannerURL               string        `json:"profile_banner_url"`
			ProfileLinkColor               string        `json:"profile_link_color"`
			ProfileSidebarBorderColor      string        `json:"profile_sidebar_border_color"`
			ProfileSidebarFillColor        string        `json:"profile_sidebar_fill_color"`
			ProfileTextColor               string        `json:"profile_text_color"`
			ProfileUseBackgroundImage      bool          `json:"profile_use_background_image"`
			HasExtendedProfile             bool          `json:"has_extended_profile"`
			DefaultProfile                 bool          `json:"default_profile"`
			DefaultProfileImage            bool          `json:"default_profile_image"`
			Following                      interface{}   `json:"following"`
			FollowRequestSent              interface{}   `json:"follow_request_sent"`
			Notifications                  interface{}   `json:"notifications"`
			TranslatorType                 string        `json:"translator_type"`
			WithheldInCountries            []interface{} `json:"withheld_in_countries"`
		} `json:"user"`
		Geo           interface{} `json:"geo"`
		Coordinates   interface{} `json:"coordinates"`
		Place         interface{} `json:"place"`
		Contributors  interface{} `json:"contributors"`
		IsQuoteStatus bool        `json:"is_quote_status"`
		RetweetCount  int         `json:"retweet_count"`
		FavoriteCount int         `json:"favorite_count"`
		Favorited     bool        `json:"favorited"`
		Retweeted     bool        `json:"retweeted"`
		Lang          string      `json:"lang"`
	} `json:"quoted_status,omitempty"`
	PossiblySensitive bool `json:"possibly_sensitive,omitempty"`
	ExtendedEntities  struct {
		Media []struct {
			ID            int64  `json:"id"`
			IDStr         string `json:"id_str"`
			Indices       []int  `json:"indices"`
			MediaURL      string `json:"media_url"`
			MediaURLHTTPS string `json:"media_url_https"`
			URL           string `json:"url"`
			DisplayURL    string `json:"display_url"`
			ExpandedURL   string `json:"expanded_url"`
			Type          string `json:"type"`
			Sizes         struct {
				Large struct {
					W      int    `json:"w"`
					H      int    `json:"h"`
					Resize string `json:"resize"`
				} `json:"large"`
				Thumb struct {
					W      int    `json:"w"`
					H      int    `json:"h"`
					Resize string `json:"resize"`
				} `json:"thumb"`
				Small struct {
					W      int    `json:"w"`
					H      int    `json:"h"`
					Resize string `json:"resize"`
				} `json:"small"`
				Medium struct {
					W      int    `json:"w"`
					H      int    `json:"h"`
					Resize string `json:"resize"`
				} `json:"medium"`
			} `json:"sizes"`
		} `json:"media"`
	} `json:"extended_entities,omitempty"`
}

type TwitterMetadata struct {
	CompletedIn float64 `json:"completed_in"`
	MaxID       int64   `json:"max_id"`
	MaxIDStr    string  `json:"max_id_str"`
	NextResults string  `json:"next_results"`
	Query       string  `json:"query"`
	RefreshURL  string  `json:"refresh_url"`
	Count       int     `json:"count"`
	SinceID     int     `json:"since_id"`
	SinceIDStr  string  `json:"since_id_str"`
}
