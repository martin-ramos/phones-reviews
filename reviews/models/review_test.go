package models

import "testing"

func NewReview(stars int, comment string) *CreateReviewCMD {
	return &CreateReviewCMD{
		Stars:   stars,
		Comment: comment,
	}
}

func Test_createReviewValidateWithCorrectParams(t *testing.T) {
	r := NewReview(4, "The iphone x looks good")

	err := r.validate()

	if err != nil {
		t.Error("The validation did not pass")
		t.Fail()
	}
}

func Test_createReviewValidateShouldFail(t *testing.T) {
	r := NewReview(8, "The iphone x looks good")

	err := r.validate()

	if err == nil {
		t.Error("Should stars between 1-8 or chars of comment less than 400")
		t.Fail()
	}
}
