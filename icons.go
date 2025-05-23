package icons

import (
	"encoding"
	"errors"
	"fmt"
	"regexp"
	"slices"
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
func (i Icon) String() string {
	return i.Collection + ":" + i.Icon
}

// MarshalText implements encoding.TextMarshaler.
func (i Icon) MarshalText() (text []byte, err error) {
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

type AllIcons []Icon

func FromCollectionKeys(keys map[string][]string) AllIcons {
	return slices.Collect(func(yield func(Icon) bool) {
		for col, icons := range keys {
			for _, icon := range icons {
				if !yield(Icon{Collection: col, Icon: icon}) {
					return
				}
			}
		}
	})
}

type IconResource interface {
	Name() string
	Content() []byte
}

type Collections[T any, PT interface {
	*T
	IconResource
}] interface {
	Lookup(collection, icon string) PT
}

func GetResource[T any, PT interface {
	*T
	IconResource
}](coll Collections[T, PT], i *Icon) PT {
	if i == nil {
		return nil
	}
	return coll.Lookup(i.Collection, i.Icon)
}

var _ encoding.TextUnmarshaler = (*Icon)(nil)
var _ encoding.TextMarshaler = (*Icon)(nil)
var _ fmt.Stringer = (*Icon)(nil)
