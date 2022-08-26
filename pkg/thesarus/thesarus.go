package thesarus

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type BigHuge struct {
	ApiKey string
}

func New(apikey string) *BigHuge {
	return &BigHuge{
		ApiKey: apikey,
	}
}

type synonym struct {
	Verb *verb
}

type verb struct {
	Syn []string
	Rel []string
}

type Thesarus interface {
	Synonym(term string) ([]string, error)
}

func (b *BigHuge) Synonyms(term string) ([]string, error) {
	url := fmt.Sprintf("https://words.bighugelabs.com/api/2/%s/%s/json", b.ApiKey, term)
	r, err := http.Get(url)

	if err != nil {
		return nil, fmt.Errorf("thesarus: failed to look for synonym of %s. %s", term, err.Error())
	}

	defer r.Body.Close()
	var data synonym

	if err = json.NewDecoder(r.Body).Decode(&data); err != nil {

		return nil, err
	}

	var syns []string

	if data.Verb != nil {
		syns = append(syns, data.Verb.Syn...)
	}

	return syns, nil
}
