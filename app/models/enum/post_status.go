package enum

//PostStatus is the status of a given post
type PostStatus int

var (
	//PostOpen is the default status
	PostOpen PostStatus
	//PostInDevelopment is used when the post has been accepted and work is in progress
	PostInDevelopment PostStatus = 1
	//PostAlpha is used when the feature is in alpha testing
	PostAlpha PostStatus = 2
	//PostBeta is used when the feature is in beta testing
	PostBeta PostStatus = 3
	//PostCompleted is used when the post has been accepted and already implemented
	PostCompleted PostStatus = 4
	//PostDeclined is used when organizers decide to decline an post
	PostDeclined PostStatus = 5
	//PostPlanned is used when organizers have accepted an post and it's on the roadmap
	PostPlanned PostStatus = 6
	//PostDuplicate is used when the post has already been posted before
	PostDuplicate PostStatus = 7
	//PostDeleted is used when the post is completely removed from the site and should never be shown again
	PostDeleted PostStatus = 8
)
var postStatusIDs = map[PostStatus]string{
	PostOpen:          "open",
	PostInDevelopment: "in-development",
	PostAlpha:         "alpha",
	PostBeta:          "beta",
	PostCompleted:     "completed",
	PostDeclined:      "declined",
	PostPlanned:       "planned",
	PostDuplicate:     "duplicate",
	PostDeleted:       "deleted",
}

var postStatusNames = map[string]PostStatus{
	"open":           PostOpen,
	"in-development": PostInDevelopment,
	"started":        PostInDevelopment, // Legacy support for old 'started' status
	"alpha":          PostAlpha,
	"beta":           PostBeta,
	"completed":      PostCompleted,
	"declined":       PostDeclined,
	"planned":        PostPlanned,
	"duplicate":      PostDuplicate,
	"deleted":        PostDeleted,
}

// MarshalText returns the Text version of the post status
func (status PostStatus) MarshalText() ([]byte, error) {
	return []byte(postStatusIDs[status]), nil
}

// UnmarshalText parse string into a post status
func (status *PostStatus) UnmarshalText(text []byte) error {
	*status = postStatusNames[string(text)]
	return nil
}

// Name returns the name of a post status
func (status PostStatus) Name() string {
	name, ok := postStatusIDs[status]
	if ok {
		return name
	}
	return "unknown"
}
