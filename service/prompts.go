package service

const generateDefinitionsPrompt = "You are a English language expert with extensive knowledge of English dictionaries. Provide the most accurate definitions for a given word, along with 2 example sentences, put into the examples field. If the word could be a verb, return the tenses, otherwise make it null. For the pronunciation, use IPA transcription style with /."

const predictWordPrompt = "You are a English language expert with extensive knowledge of English dictionaries. If the user provides a specific word or a short sentence, correct the potential spelling errors and return them. If the user provides non-english, try your best to find the corresponting english word(s). If the user provides unrecognizable combination of letters, set fail to ture and return fail message."

const jsonDescription = "Generate word definitions where each definition object contains its own examples array. Do not put examples at the top level - they must be inside each definition object along with partOfSpeech and meaning."

type PromptStore struct {
}

func (p *PromptStore) GenerateDefinitionsPrompt() string {
	return generateDefinitionsPrompt
}

func (p *PromptStore) PredictWordPrompt() string {
	return predictWordPrompt
}

func (p *PromptStore) JsonDescription() string {
	return jsonDescription
}

func NewPromptStore() *PromptStore {
	return &PromptStore{}
}
