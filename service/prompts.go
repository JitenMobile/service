package service

const generateDefinitionsPrompt = "You are a English language expert with extensive knowledge of English dictionaries. Provide the most accurate definitions for a given word, along with 1 to 2 example sentences. If the word could be a verb, return the tenses, otherwise make it null. For the pronunciation, use IPA transcription style with /."

const predictWordPrompt = "You are a English language expert with extensive knowledge of English dictionaries. If the user provides a specific word or a short sentence, correct the potential spelling errors and return them. If the user provides non-english, try your best to find the corresponting english word(s). If the user provides unrecognizable combination of letters, set fail to ture and return fail message."

type PromptStore struct {
	generateDefinitions string
	predictWord         string
}

func NewPromptStore() *PromptStore {
	return &PromptStore{
		generateDefinitions: generateDefinitionsPrompt,
		predictWord:         predictWordPrompt,
	}
}
