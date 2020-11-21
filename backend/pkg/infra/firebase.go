package infra

import (
	"context"
	"link-note/backend/pkg/model"
	"log"
	"strings"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"

	firebase "firebase.google.com/go"
	"google.golang.org/api/iterator"
)

type FirestoreAuth struct {
	Type                        string `json:"type"`
	Project_id                  string `json:"project_id"`
	Private_key_id              string `json:"private_key_id"`
	Private_key                 string `json:"private_key"`
	Client_email                string `json:"client_email"`
	Client_id                   string `json:"client_id"`
	Auth_uri                    string `json:"auth_uri"`
	Token_uri                   string `json:"token_uri"`
	Auth_provider_x509_cert_url string `json:"auth_provider_x509_cert_url"`
	Client_x509_cert_url        string `json:"client_x509_cert_url"`
}

type FireBaseClient struct {
	FireBase  *firebase.App
	FireStore *firestore.Client
	Ctx       context.Context
}

type FireBaseHandler interface {
	InsertData(data model.Content)
	DeleteData(uid, id string) error
	GetData(uid string) []model.Content
	AuthJWT(jwt string) error
}

type FireBase struct {
	FireBaseHandler
}

func Init_firebase() FireBaseHandler {

	ctx := context.Background()
	sa := option.WithCredentialsFile("./firestore.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		return nil
	}
	client, err := app.Firestore(ctx)

	if err != nil {
		return nil
	}

	return &FireBaseClient{
		FireBase:  app,
		FireStore: client,
		Ctx:       ctx,
	}
}

func (fb *FireBaseClient) InsertData(data model.Content) {

	_, updateError := fb.FireStore.Collection(data.Uid).Doc(data.Content_id).Set(fb.Ctx, map[string]interface{}{
		"content_id": data.Content_id,
		"comment":    data.Comment,
		"url":        data.Url,
		"date":       data.Date,
	}, firestore.MergeAll)
	if updateError != nil {
		log.Printf("An error has occurred: %s", updateError)
	}
}

func (fb *FireBaseClient) DeleteData(uid, id string) error {

	_, err := fb.FireStore.Collection(uid).Doc(id).Delete(fb.Ctx)

	if err != nil {
		return err
	}

	return nil

}

func (fb *FireBaseClient) GetData(uid string) []model.Content {

	var res_data []model.Content
	iter := fb.FireStore.Collection(uid).Documents(fb.Ctx)

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}
		data := doc.Data()
		var content model.Content
		content.Comment = data["comment"].(string)
		content.Url = data["url"].(string)
		content.Content_id = data["content_id"].(string)
		content.Date = int(data["date"].(int64))

		res_data = append(res_data, content)
	}

	return res_data

}

func (fb *FireBaseClient) AuthJWT(jwt string) error {

	auth, err := fb.FireBase.Auth(fb.Ctx)
	if err != nil {

		return err

	}

	idToken := strings.Replace(jwt, "Bearer ", "", 1)

	_, err = auth.VerifyIDToken(fb.Ctx, idToken)
	if err != nil {

		return err
	}
	return nil
}
