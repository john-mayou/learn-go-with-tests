package dictionary

type Dictionary map[string]string

func Search(dict Dictionary, word string) string {
	return dict[word]
}
