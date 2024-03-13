package models

import (
	"golang_app/golangApp/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestCreateWeddingData(t *testing.T) {
	var err error
	var userResult string
	var objID primitive.ObjectID
	var result string
	user := models.User{
		Username:  "testuser1",
		Email:     "test1@example.com",
		Firstname: "Test",
		Lastname:  "User",
		Authtoken: "testtoken",
	}
	userResult, err = models.CreateUser(user)
	objID, err = primitive.ObjectIDFromHex(userResult)

	weddingData := models.WeddingData{
		UserID: objID,
		Bride: models.BrideGroom{
			Firstname:  "Devia",
			Lastname:   "Nur Fadillah",
			ProfileURL: "",
			SocialMedia: []models.SocialMedia{
				{
					Username: "devia_insta",
					Link:     "https://instagram.com/devia_insta",
					Platform: "instagram",
				},
			},
			Parent: models.Parent{
				Father: struct {
					Firstname string `json:"firstname"`
					Lastname  string `json:"lastname"`
				}{
					Firstname: "Masdi",
					Lastname:  "Aja",
				},
				Mother: struct {
					Firstname string `json:"firstname"`
					Lastname  string `json:"lastname"`
				}{
					Firstname: "Evi",
					Lastname:  "Aja",
				},
			},
		},
		Groom: models.BrideGroom{
			Firstname:  "Soeltan",
			Lastname:   "Zaki Rizaldy",
			ProfileURL: "",
			SocialMedia: []models.SocialMedia{
				{
					Username: "soeltan_insta",
					Link:     "https://instagram.com/soeltan_insta",
					Platform: "instagram",
				},
			},
			Parent: models.Parent{
				Father: struct {
					Firstname string `json:"firstname"`
					Lastname  string `json:"lastname"`
				}{
					Firstname: "Masdi",
					Lastname:  "",
				},
				Mother: struct {
					Firstname string `json:"firstname"`
					Lastname  string `json:"lastname"`
				}{
					Firstname: "Evi",
					Lastname:  "",
				},
			},
		},
		Wedding: models.Wedding{
			Date: "15-01-2025",
			Akad: models.AkadResepsi{
				Date:    "15-01-2025",
				Start:   "09:00",
				End:     "selesai",
				Address: "Plataran Menteng, Jalan HOS. Cokroaminoto, RT.6/RW.4, Gondangdia, Kota Jakarta Pusat, Daerah Khusus Ibukota Jakarta, Indonesia",
				MapURL:  "https://www.google.com/maps/place/Plataran+Menteng/@-6.1919015,106.8285812,17z/data=!3m1!4b1!4m6!3m5!1s0x2e69f423590651f7:0x983424b56075bd8!8m2!3d-6.1919015!4d106.8285812!16s%2Fg%2F11c5sw4rnx?entry=ttu",
			},
			Resepsi: models.AkadResepsi{
				Date:    "15-01-2025",
				Start:   "12:00",
				End:     "selesai",
				Address: "Plataran Menteng, Jalan HOS. Cokroaminoto, RT.6/RW.4, Gondangdia, Kota Jakarta Pusat, Daerah Khusus Ibukota Jakarta, Indonesia",
				MapURL:  "https://www.google.com/maps/place/Plataran+Menteng/@-6.1919015,106.8285812,17z/data=!3m1!4b1!4m6!3m5!1s0x2e69f423590651f7:0x983424b56075bd8!8m2!3d-6.1919015!4d106.8285812!16s%2Fg%2F11c5sw4rnx?entry=ttu",
			},
		},
	}

	result, err = models.CreateWeddingData(weddingData)
	assert.Nil(t, err, "expected no error")
	assert.NotNil(t, result, "expected insert result not to be nil")
}

func TestGetWeddingDataByUserId(t *testing.T) {
	email := "test1@example.com"
	userId := models.GetUserByEmail(email)
	result := models.GetWeddingDataByUserId(userId.Id.Hex())
	assert.NotNil(t, result, "expected Get result not to be nil")

}
