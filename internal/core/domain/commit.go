package domain

type Commit struct {
	CommitType    CommitType
	CommitMessage string
	CommitEmoji   CommitEmoji
}

func (c *Commit) String() string {
	return c.CommitType.Id + ": " + c.CommitEmoji.Code + " " + c.CommitMessage
}
