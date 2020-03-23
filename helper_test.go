package log

import "testing"

func TestJSON(t *testing.T) {
	x := &struct {
		A string `json:"a"`
		B int    `json:"b"`
		C string `json:"-"`
	}{
		A: "a_val",
		B: 100,
		C: "c",
	}
	have := JSON(x)
	want := `{"a":"a_val","b":100}`
	if have != want {
		t.Errorf("have:%s, want:%s", have, want)
		return
	}
}

func TestXML(t *testing.T) {
	x := &struct {
		XMLName struct{} `xml:"msg"`
		A       string   `xml:"a"`
		B       int      `xml:"b"`
		C       string   `xml:"-"`
	}{
		A: "a_val",
		B: 100,
		C: "c",
	}
	have := XML(x)
	want := `<msg><a>a_val</a><b>100</b></msg>`
	if have != want {
		t.Errorf("have:%s, want:%s", have, want)
		return
	}
}
