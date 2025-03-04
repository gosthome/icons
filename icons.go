package icons

import (
	"encoding"
	"errors"
	"fmt"
	"regexp"
	"sync"

	"fyne.io/fyne/v2"
)

//go:generate bash -c "cd _gen && go run ."

var iconRE = regexp.MustCompile(`^$|^([\w\-]+):([\w\-]+)$`)

type Icon struct {
	Collection string
	Icon       string
}

func Parse(s string) (*Icon, error) {
	ret := &Icon{}
	err := ret.UnmarshalString(s)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// String implements fmt.Stringer.
func (i *Icon) String() string {
	return i.Collection + ":" + i.Icon
}

// MarshalText implements encoding.TextMarshaler.
func (i *Icon) MarshalText() (text []byte, err error) {
	return []byte(i.String()), nil
}

func (i *Icon) UnmarshalString(text string) error {
	data := iconRE.FindStringSubmatch(text)
	if data == nil {
		return errors.New("icon does not adhere to format collection:icon-name, ex. mdi:chip")
	}
	i.Collection = data[1]
	i.Icon = data[2]
	return nil
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (i *Icon) UnmarshalText(text []byte) error {
	return i.UnmarshalString(string(text))
}

func (i *Icon) GetResource() fyne.Resource {
	collectionsMux.RLock()
	defer collectionsMux.RUnlock()
	coll, ok := collections[i.Collection]
	if !ok {
		return nil
	}
	icon, ok := coll[i.Icon]
	if !ok {
		return nil
	}
	return icon
}

var _ encoding.TextUnmarshaler = (*Icon)(nil)
var _ encoding.TextMarshaler = (*Icon)(nil)
var _ fmt.Stringer = (*Icon)(nil)

type Collection map[string]*fyne.StaticResource

type Collections map[string]Collection

func RegisteredCollection(name string, c Collection) Collection {
	collectionsMux.Lock()
	defer collectionsMux.Unlock()
	collections[name] = c
	return c
}

var (
	collections    Collections = make(Collections)
	collectionsMux sync.RWMutex
)
