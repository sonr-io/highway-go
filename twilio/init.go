package twilio

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"
	"time"

	db "github.com/sonr-io/highway-go/database"
	"github.com/sonr-io/highway-go/models"
	twilio "github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SendSMS(phone string, db *db.MongoClient) *models.Verify {

	client := twilio.NewRestClientWithParams(twilio.RestClientParams{
		Username: os.Getenv("TWILIO_SID"),
		Password: os.Getenv("TWILIO_AUTH_ID"),
	})

	params := &openapi.CreateMessageParams{}
	params.SetTo("+1" + phone)
	params.SetFrom(os.Getenv("TWILIO_PHONE"))

	// generate verification code and message
	verificationCode := verificationCodeGenerator()
	message := "Hello from TagMatch! Your verification code is " + strconv.FormatInt(verificationCode, 10)

	params.SetBody(message)

	_, err := client.ApiV2010.CreateMessage(params)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("SMS sent successfully!")
	}

	// This is dumb
	zero := 0

	ver := &models.Verify{
		VerifyID:         primitive.NewObjectID().Hex(),
		VerificationCode: strconv.FormatInt(verificationCode, 10),
		PhoneNum:         phone,
		CreatedAt:        time.Now().String(),
		Verified:         false,
		Attempts:         &zero,
		VerifyOpenDate:   time.Now().Format(time.RFC3339),
	}

	//db.CreateVerification(ver)

	return ver
}

func verificationCodeGenerator() int64 {
	max := big.NewInt(899999)
	n, err := rand.Int(rand.Reader, max)
	if err != nil {
		log.Print(err)
		log.Print("\nRand function failed in twilio package\n")
		return 123456
	}

	n.Add(n, big.NewInt(100000))

	return n.Int64()
}
