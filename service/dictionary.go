package service

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/jiten-mobile/service/graph/model"
)

type DictionaryStore struct {
	client *firestore.Client
}

func NewDictionaryService(client *firestore.Client) *DictionaryStore {
	return &DictionaryStore{
		client: client,
	}
}

func (db *DictionaryStore) GetWord(ctx context.Context, word string) (*model.Word, error) {

	doc, err := db.client.Collection("dictionary").Doc(word).Get(ctx)
	if err != nil {
		return nil, err
	}
	var wordData model.Word
	if err := doc.DataTo(&wordData); err != nil {
		return nil, err
	}
	wordData.ID = wordData.Word
	return &wordData, nil
}

// func (db *DictionaryStore) GetTranslation(ctx context.Context, word string, targetLang string) (*model.Translation, error) {
// 	doc, err := db.client.Collection("dictionary").Doc(word).Get(ctx)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var wordData model.Word
// 	if err := doc.DataTo(&wordData); err != nil {
// 		return nil, err
// 	}

// definitions := wordData.Definitions

// doc, err := db.client.Collection("translations").Doc(word).Collection(targetLang).Doc(word).Get(ctx)
// if err != nil {
// 	return nil, err
// }

// var translationDoc map[string]model.Translation
// if err := doc.DataTo(translationDoc); err != nil {
// 	return nil, err
// }

// return nil, nil
// }

func (db *DictionaryStore) WriteWord(ctx context.Context, wordData *model.Word) error {
	wordData.ID = wordData.Word
	docRef := db.client.Collection("dictionary").Doc(wordData.Word)
	_, err := docRef.Set(ctx, wordData)
	return err
}

func (db *DictionaryStore) WriteTranslation(ctx context.Context, word string, targetLang string, translationData *model.Translation) error {
	docRef := db.client.Collection("translations").Doc(word).Collection(targetLang).Doc(word)
	_, err := docRef.Set(ctx, translationData)
	return err
}
