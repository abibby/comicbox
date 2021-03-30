// Code generated by github.com/Yamashou/gqlgenc, DO NOT EDIT.

package anilist

import (
	"context"
	"net/http"

	"github.com/Yamashou/gqlgenc/client"
)

type Client struct {
	Client *client.Client
}

func NewClient(cli *http.Client, baseURL string, options ...client.HTTPRequestOption) *Client {
	return &Client{Client: client.NewClient(cli, baseURL, options...)}
}

type Query struct {
	Page                *Page                "json:\"Page\" graphql:\"Page\""
	Media               *Media               "json:\"Media\" graphql:\"Media\""
	MediaTrend          *MediaTrend          "json:\"MediaTrend\" graphql:\"MediaTrend\""
	AiringSchedule      *AiringSchedule      "json:\"AiringSchedule\" graphql:\"AiringSchedule\""
	Character           *Character           "json:\"Character\" graphql:\"Character\""
	Staff               *Staff               "json:\"Staff\" graphql:\"Staff\""
	MediaList           *MediaList           "json:\"MediaList\" graphql:\"MediaList\""
	MediaListCollection *MediaListCollection "json:\"MediaListCollection\" graphql:\"MediaListCollection\""
	GenreCollection     []*string            "json:\"GenreCollection\" graphql:\"GenreCollection\""
	MediaTagCollection  []*MediaTag          "json:\"MediaTagCollection\" graphql:\"MediaTagCollection\""
	User                *User                "json:\"User\" graphql:\"User\""
	Viewer              *User                "json:\"Viewer\" graphql:\"Viewer\""
	Notification        NotificationUnion    "json:\"Notification\" graphql:\"Notification\""
	Studio              *Studio              "json:\"Studio\" graphql:\"Studio\""
	Review              *Review              "json:\"Review\" graphql:\"Review\""
	Activity            ActivityUnion        "json:\"Activity\" graphql:\"Activity\""
	ActivityReply       *ActivityReply       "json:\"ActivityReply\" graphql:\"ActivityReply\""
	Following           *User                "json:\"Following\" graphql:\"Following\""
	Follower            *User                "json:\"Follower\" graphql:\"Follower\""
	Thread              *Thread              "json:\"Thread\" graphql:\"Thread\""
	ThreadComment       []*ThreadComment     "json:\"ThreadComment\" graphql:\"ThreadComment\""
	Recommendation      *Recommendation      "json:\"Recommendation\" graphql:\"Recommendation\""
	Like                *User                "json:\"Like\" graphql:\"Like\""
	Markdown            *ParsedMarkdown      "json:\"Markdown\" graphql:\"Markdown\""
	AniChartUser        *AniChartUser        "json:\"AniChartUser\" graphql:\"AniChartUser\""
	SiteStatistics      *SiteStatistics      "json:\"SiteStatistics\" graphql:\"SiteStatistics\""
}

type Mutation struct {
	UpdateUser                 *User            "json:\"UpdateUser\" graphql:\"UpdateUser\""
	SaveMediaListEntry         *MediaList       "json:\"SaveMediaListEntry\" graphql:\"SaveMediaListEntry\""
	UpdateMediaListEntries     []*MediaList     "json:\"UpdateMediaListEntries\" graphql:\"UpdateMediaListEntries\""
	DeleteMediaListEntry       *Deleted         "json:\"DeleteMediaListEntry\" graphql:\"DeleteMediaListEntry\""
	DeleteCustomList           *Deleted         "json:\"DeleteCustomList\" graphql:\"DeleteCustomList\""
	SaveTextActivity           *TextActivity    "json:\"SaveTextActivity\" graphql:\"SaveTextActivity\""
	SaveMessageActivity        *MessageActivity "json:\"SaveMessageActivity\" graphql:\"SaveMessageActivity\""
	SaveListActivity           *ListActivity    "json:\"SaveListActivity\" graphql:\"SaveListActivity\""
	DeleteActivity             *Deleted         "json:\"DeleteActivity\" graphql:\"DeleteActivity\""
	ToggleActivitySubscription ActivityUnion    "json:\"ToggleActivitySubscription\" graphql:\"ToggleActivitySubscription\""
	SaveActivityReply          *ActivityReply   "json:\"SaveActivityReply\" graphql:\"SaveActivityReply\""
	DeleteActivityReply        *Deleted         "json:\"DeleteActivityReply\" graphql:\"DeleteActivityReply\""
	ToggleLike                 []*User          "json:\"ToggleLike\" graphql:\"ToggleLike\""
	ToggleLikeV2               LikeableUnion    "json:\"ToggleLikeV2\" graphql:\"ToggleLikeV2\""
	ToggleFollow               *User            "json:\"ToggleFollow\" graphql:\"ToggleFollow\""
	ToggleFavourite            *Favourites      "json:\"ToggleFavourite\" graphql:\"ToggleFavourite\""
	UpdateFavouriteOrder       *Favourites      "json:\"UpdateFavouriteOrder\" graphql:\"UpdateFavouriteOrder\""
	SaveReview                 *Review          "json:\"SaveReview\" graphql:\"SaveReview\""
	DeleteReview               *Deleted         "json:\"DeleteReview\" graphql:\"DeleteReview\""
	RateReview                 *Review          "json:\"RateReview\" graphql:\"RateReview\""
	SaveRecommendation         *Recommendation  "json:\"SaveRecommendation\" graphql:\"SaveRecommendation\""
	SaveThread                 *Thread          "json:\"SaveThread\" graphql:\"SaveThread\""
	DeleteThread               *Deleted         "json:\"DeleteThread\" graphql:\"DeleteThread\""
	ToggleThreadSubscription   *Thread          "json:\"ToggleThreadSubscription\" graphql:\"ToggleThreadSubscription\""
	SaveThreadComment          *ThreadComment   "json:\"SaveThreadComment\" graphql:\"SaveThreadComment\""
	DeleteThreadComment        *Deleted         "json:\"DeleteThreadComment\" graphql:\"DeleteThreadComment\""
	UpdateAniChartSettings     *string          "json:\"UpdateAniChartSettings\" graphql:\"UpdateAniChartSettings\""
	UpdateAniChartHighlights   *string          "json:\"UpdateAniChartHighlights\" graphql:\"UpdateAniChartHighlights\""
}
type SaveMediaListEntryPayload struct {
	SaveMediaListEntry *struct {
		MediaID         int  "json:\"mediaId\" graphql:\"mediaId\""
		Progress        *int "json:\"progress\" graphql:\"progress\""
		ProgressVolumes *int "json:\"progressVolumes\" graphql:\"progressVolumes\""
		Media           *struct {
			ID    int "json:\"id\" graphql:\"id\""
			Title *struct {
				UserPreferred *string "json:\"userPreferred\" graphql:\"userPreferred\""
			} "json:\"title\" graphql:\"title\""
		} "json:\"media\" graphql:\"media\""
	} "json:\"SaveMediaListEntry\" graphql:\"SaveMediaListEntry\""
}

const SaveMediaListEntryQuery = `mutation SaveMediaListEntry ($mediaId: Int!, $progress: Int, $progressVolumes: Int) {
	SaveMediaListEntry(mediaId: $mediaId, progress: $progress, progressVolumes: $progressVolumes) {
		mediaId
		progress
		progressVolumes
		media {
			id
			title {
				userPreferred
			}
		}
	}
}
`

func (c *Client) SaveMediaListEntry(ctx context.Context, mediaID int, progress *int, progressVolumes *int, httpRequestOptions ...client.HTTPRequestOption) (*SaveMediaListEntryPayload, error) {
	vars := map[string]interface{}{
		"mediaId":         mediaID,
		"progress":        progress,
		"progressVolumes": progressVolumes,
	}

	var res SaveMediaListEntryPayload
	if err := c.Client.Post(ctx, SaveMediaListEntryQuery, &res, vars, httpRequestOptions...); err != nil {
		return nil, err
	}

	return &res, nil
}