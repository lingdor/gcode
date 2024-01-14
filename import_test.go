package gcode

import (
	"fmt"
	"testing"
)

func TestAddImport(t *testing.T) {

	code := `package aa`
	bs, err := AddImport([]byte(code), "arr", "github.com/lingdor/magicarray")
	if err != nil {
		panic(err)
	}
	expect := `package aa
import(
    arr "github.com/lingdor/magicarray"
)
`
	if string(bs) != expect {
		t.Errorf("val:%q,expect:%q", string(bs), expect)
	}

	code = string(bs)
	bs, err = AddImport([]byte(code), "arr", "github.com/lingdor/magicarray")
	if err != nil {
		panic(err)
	}
	if string(bs) != expect {
		t.Errorf("\nval:%q\nexpect:%q", string(bs), expect)
	}
	code = string(bs)
	bs, err = AddImport([]byte(code), "", "github.com/lingdor/magicarray/gsql")
	if err != nil {
		panic(err)
	}
	expect = `package aa
import(
    arr "github.com/lingdor/magicarray"
    "github.com/lingdor/magicarray/gsql"
)
`
	if string(bs) != expect {
		t.Errorf("\nval:%q\nexpect:%q", string(bs), expect)
	}

	fmt.Printf(string(bs))

}
