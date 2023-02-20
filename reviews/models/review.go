package models

import (
	"errors"
	"time"
)

const maxLengthComment = 400

// Review
type Review struct {
	Id      int64
	Stars   int       // 1 - 5
	Comment string    // max 400 characters
	Date    time.Time // created At
}

// CreateReviewCMD
type CreateReviewCMD struct {
	Stars   int    `json:"stars"`
	Comment string `json:"comment"`
}

func (cmd *CreateReviewCMD) validate() error {
	if cmd.Stars < 1 || cmd.Stars > 5 {
		return errors.New("stars must be between 1 - 5")
	}

	if len(cmd.Comment) > maxLengthComment {
		return errors.New("comment must be less than 400 chars")
	}
	return nil
}
