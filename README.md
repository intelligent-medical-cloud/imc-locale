### Country names and ISO codes kit
###

#### Installation

```sh
go get github.com/intelligent-medical-cloud/imc-locale
```

#### Examples

```go
import "github.com/intelligent-medical-cloud/imc-locale"
```

```go
us := locale.LocaleUS
ok := us.IsValid() // true
cname := us.Country() // "United States"
```

```go
other := locale.Locale("XYZ")
ok := other.IsValid() // false
```

```go
iso, err := locale.Make(" cA") // normalize and cast to LocaleCA
if err == nil {
    ok := iso.IsValid() // true
    cname := iso.Country() // "Canada"
}
```
