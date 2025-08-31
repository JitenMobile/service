package service

const generateDefinitionsPrompt = "You are a English language expert with extensive knowledge of English dictionaries, and your knowledge is up-to-date. Provide the most accurate definitions for a given word, and also the nowadays meaning if it exists, along with at least 1 and at most 3 example sentences, put into the examples field. For the part of speech, return type only, no need to do further explaination. If the word could be a verb, return the tenses, otherwise make it null. For the pronunciation, use IPA transcription style with /. Synonyms, antonyms, and related terms should all be single word."

const predictWordPrompt = "You are a English language expert with extensive knowledge of English dictionaries. If the user provides a specific word or a short sentence, correct the potential spelling errors and return them. If the user provides non-english, try your best to find the corresponting english word(s). If the user provides unrecognizable combination of letters, set fail to ture and return fail message."

const generationJsonDescription = "Generate word definitions where each definition object contains its own examples array. Do not put examples at the top level - they must be inside each definition object along with partOfSpeech and meaning."

const translationJsonDescription = "Translate the given object fields into the target language. Each field in the object should be translated accurately while preserving its original meaning and context. Ensure proper grammar and syntax in the target language."

const translationPrefixPrompt = "You are a professional translator who can precisely translate the each fields of the given object from English to "

type PromptStore struct {
	generateDefinitionsPrompt  string
	predictWordPrompt          string
	translationPrefixPrompt    string
	generationJsonDescription  string
	translationJsonDescription string
}

func (p *PromptStore) GetTranslationPrompt(targetLang string) string {
	return translationPrefixPrompt + targetLang
}

func NewPromptStore() *PromptStore {
	return &PromptStore{
		generateDefinitionsPrompt:  generateDefinitionsPrompt,
		predictWordPrompt:          predictWordPrompt,
		translationPrefixPrompt:    translationPrefixPrompt,
		generationJsonDescription:  generationJsonDescription,
		translationJsonDescription: translationJsonDescription,
	}
}
