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
