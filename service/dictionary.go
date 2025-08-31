package service

import (
	"context"
	"fmt"
	"strings"

	"cloud.google.com/go/firestore"
	"github.com/jiten-mobile/service/graph/model"
)

type DictionaryStore struct {
	client *firestore.Client
}

const (
	DictionaryCollection  = "dictionary"
	TranslationCollection = "translations"
	NotFoundPrefix        = "[Not Found]"
)

func NewDictionaryService(client *firestore.Client) *DictionaryStore {
	return &DictionaryStore{
		client: client,
	}
}

func (db *DictionaryStore) GetWord(ctx context.Context, word string) (*model.Word, error) {

	docName := strings.ToLower(word)
	doc, err := db.client.Collection(DictionaryCollection).Doc(docName).Get(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s, %v", NotFoundPrefix, err) // this helps resolve identify not found condition
	}
	var wordData model.Word
	if err := doc.DataTo(&wordData); err != nil {
		return nil, err
	}
	wordData.ID = wordData.Word
	return &wordData, nil
}

// Get translation data from DB
func (db *DictionaryStore) GetTranslation(ctx context.Context, word string, targetLang string) (*model.Translation, error) {
	docName := strings.ToLower(word)
	doc, err := db.client.Collection(TranslationCollection).Doc(docName).Collection(targetLang).Doc(docName).Get(ctx)
	if err != nil {
		return nil, err
	}

	var translationData model.Translation
	if err := doc.DataTo(&translationData); err != nil {
		return nil, err
	}

	return &translationData, nil
}

func (db *DictionaryStore) WriteWord(ctx context.Context, wordData *model.Word) error {
	wordData.ID = wordData.Word
	docName := strings.ToLower(wordData.Word)
	docRef := db.client.Collection(DictionaryCollection).Doc(docName)
	_, err := docRef.Set(ctx, wordData)
	return err
}

func (db *DictionaryStore) WriteTranslation(ctx context.Context, word string, targetLang string, translationData *model.Translation) error {
	docName := strings.ToLower(word)
	docRef := db.client.Collection(TranslationCollection).Doc(docName).Collection(targetLang).Doc(docName)
	_, err := docRef.Set(ctx, translationData)
	return err
}
