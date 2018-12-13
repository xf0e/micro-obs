package item

import (
	"strconv"
	"testing"

	"github.com/obitech/micro-obs/util"
)

var (
	sampleItems = []struct {
		name string
		desc string
		qty  int
	}{
		{"test", "test", 0},
		{"orange", "a juicy fruit", 100},
		{"😍", "lovely smily", 999},
		{"     ", "﷽", 249093419},
	}
)

func TestItem(t *testing.T) {
	var items []*Item

	t.Run("Create new item", func(t *testing.T) {
		for _, tt := range sampleItems {
			i, err := NewItem(tt.name, tt.desc, tt.qty)
			if err != nil {
				t.Errorf("failed to create new item: %#v", err)
			}
			items = append(items, i)
		}
	})

	t.Run("Confirm hash conversions", func(t *testing.T) {
		for _, i := range items {
			s, err := util.HashIDToString(i.ID)
			if err != nil {
				t.Errorf("unable to decode %#v to string: %#v", i.ID, err)
			}
			if s != i.Name {
				t.Errorf("HashIDToString(%#v), expected: %#v, got: %#v", i.ID, i.Name, s)
			}
		}
	})

	t.Run("Redis marshalling", func(t *testing.T) {
		prsKeys := []string{"name", "desc", "qty"}
		for _, i := range items {
			key, fv := i.marshalRedis()
			if key != i.ID {
				t.Errorf("marshaling unsuccesful, expected key = %#v, got key = %#v", i.ID, key)
			}
			for _, k := range prsKeys {
				if _, prs := fv[k]; !prs {
					t.Errorf("key %s not present in marshalled map.", k)
				}
			}
			if fv["name"] != i.Name {
				t.Errorf("marshaling unsuccesful, expected name = %#v, got name = %#v", i.Name, fv["name"])
			}
			if fv["desc"] != i.Desc {
				t.Errorf("marshaling unsuccesful, expected desc = %#v, got desc = %#v", i.Name, fv["desc"])
			}
			if fv["qty"] != strconv.Itoa(i.Qty) {
				t.Errorf("marshaling unsuccesful, expected qty = %#v, got qty = %#v", i.Qty, fv["qty"])
			}
		}
	})
}
