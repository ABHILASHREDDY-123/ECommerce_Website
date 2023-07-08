package handlers

import (
	"testing"

	"github.com/ABHILASHREDDY-123/ECommerce_Website.git/models"
)

func TestSendEmail(T *testing.T){
	email := &models.Email{
		Subject:"This is a subject",
		Body : "This is body",
		SenderEmail: "gajulapallyabhilash@gmail.com",
		ReceiverEmail: "20je0362@cse.iitism.ac.in",
		Password: "soyehlgkdtsicbzw",
	}
	err := SendEmail(email)
	if err != nil {
		T.Logf("FAILED")
	} else {
		T.Logf("Test PASSED")
	}
}