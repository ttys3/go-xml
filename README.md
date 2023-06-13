# go-xml

golang xml package which add marshal self-closing tag support

the code applied from https://go-review.googlesource.com/c/go/+/469495

## Oops

I found this https://twitter.com/ZeCoffee/status/766349635359211520

![José Coelho @ZeCoffee golang encoding/xml works fine until you need a self-closing tag...](https://user-images.githubusercontent.com/41882455/245492933-32b95e6d-e409-4e33-9f9b-1ec61258c617.png)

## usage

Custom MarshalXML ref [https://pkg.go.dev/encoding/xml#Marshal](https://pkg.go.dev/encoding/xml#example-package-CustomMarshalXML)

self-closing tag example:

```go
import "github.com/ttys3/go-xml"

// no `xml` struct tag is needed or can be used here
// since we handle all this in `MarshalXML`
type Foo struct {
	Bar     string
	Comment string
}

// Custom XML marshaler for Foo
func (i Foo) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	attrs := []xml.Attr{
		{
			Name:  xml.Name{Local: "bar"},
			Value: i.Bar,
		},
		{
			Name:  xml.Name{Local: "comment"},
			Value: i.Comment,
		},
	}

	// Create a self-closing tag for Item
	empty := xml.EmptyElement{
		Name: xml.Name{
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

	marshaledXML, err := xml.MarshalIndent(foo, "", "  ")
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	if string(marshaledXML) != expectedXML {
		t.Errorf("Expected marshaled XML:\n%s\n\nGot:\n%s", expectedXML, marshaledXML)
	}
}
```


## related issues

see https://github.com/golang/go/issues/21399

https://go-review.googlesource.com/c/go/+/469495

https://go-review.googlesource.com/c/go/+/59830

https://github.com/nemith/netconf/pull/27/files
