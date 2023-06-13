package xml

import "testing"

type Foo struct {
	Bar     string
	Comment string
}

// Custom XML marshaler for Foo
func (i Foo) MarshalXML(e *Encoder, start StartElement) error {
	attrs := []Attr{
		{
			Name:  Name{Local: "bar"},
			Value: i.Bar,
		},
		{
			Name:  Name{Local: "comment"},
			Value: i.Comment,
		},
	}

	// Create a self-closing tag for Item
	empty := EmptyElement{
		Name: Name{
			Space: "",
			Local: "foo",
		},
		Attr: attrs,
	}

	// can not use Encode or EncodeElement here, because they will not emit self-closing tag
	err := e.EncodeToken(empty)
	if err != nil {
		return err
	}

	// Flush must be called since we are not using Encode or EncodeElement
	if err := e.Flush(); err != nil {
		return err
	}

	return nil
}

func TestSelfClodingTagFoo(t *testing.T) {
	expectedXML := `<foo bar="hello" comment="world"/>`

	foo := Foo{
		Bar:     "hello",
		Comment: "world",
	}

	marshaledXML, err := MarshalIndent(foo, "", "  ")
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	if string(marshaledXML) != expectedXML {
		t.Errorf("Expected marshaled XML:\n%s\n\nGot:\n%s", expectedXML, marshaledXML)
	}
}
