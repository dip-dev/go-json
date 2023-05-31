# json/opt

json/opt is a package for determining the presence or absence of fields in JSON.

> **Note**
> The following types are not supported because they are not normally handled in JSON.
> - `rune`
> - `complex64`, `complex128`
> - `uintptr`

## sample
### Unmarshal
haskey is false for fields that do not exist in json

```golang
func unmarshalSample() {
	o := struct {
		Item1 opt.String `json:"item1"`
		Item2 opt.String `json:"item2"`
		Item3 opt.String `json:"item3"`
	}{}

	j := []byte(`{"item1": "abc", "item3": null}`)
	_ = json.Unmarshal(j, &o)
	fmt.Println(o.Item1.Value())  // "abc"
	fmt.Println(o.Item1.HasKey()) // true
	fmt.Println(o.Item2.Value())  // nil
	fmt.Println(o.Item2.HasKey()) // false
	fmt.Println(o.Item3.Value())  // nil
	fmt.Println(o.Item3.HasKey()) // true
}

```
### Marshal
> **Note**
> The fields (`v`, `hasKey`) that opt's own type has are not output. 
> Also, if you want to express that the key does not exist, use the json tag omitempty.


```golang

func marshalSample() {
    toPtr := func(s string) *string {
		return &s
	}
    o := struct {
		Item1 *opt.String `json:"item1,omitempty"`
		Item2 *opt.String `json:"item2,omitempty"`
		Item3 *opt.String `json:"item3,omitempty"`
	}{
        Item1: &opt.String{v: toPtr("sample"), hasKey: true},
		Item2: &opt.String{}, // hasKey is false, but the field is output
        Item3: nil,
    }
    got, _ := json.Marshal(o)
    fmt.Println(string(got)) // {"item1":"sample", "item2":null}
}
```