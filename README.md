# Perl

I think I named this folder `perl` when there were like two lines of code which
could simulate perls [`warn`](https://perldoc.perl.org/functions/warn.html)
feature. I never bothered to re-name it and I guess this wasn't meant to ever
see the light outside my folder. But here we are...

## Warn

Warn is used to simulate the way perl does warnings but in golang. This is nice
to debug where your code is doing, setting markers which will tell you which
line you got to etcetera.

```go
func debugMe(maybe1, maybe2, maybe3 bool) {
    if maybe1 {
        warn.Here()
    }

    if maybe2 {
        warn.Here("In maybe2!")
    }

    if maybe3 {
        warn.Here(fmt.Sprintf("Uhh, I'm having an error: %s", "UNKNOWN"))
    }
}
```

## Grep

Grep is small utils that will find things in lists. It has features like
[`grep`](https://perldoc.perl.org/functions/grep.html) and
[`List::Util::first`](https://perldoc.perl.org/List/Util.html#first).

```go
func lists() {
    newList := grep.List("string-literal", []string{"foo", "string-literal", "bar"})
    // newList == []string{"string-literal"}

    newList := grep.List("/re(gexp)?/", []string{"re", "redo", "do"})
    // newList == []string{"re", "redo"}

    newList := grep.First("item", []string{"not item", "item", "item"})
    // newList == "item"

    isFound := grep.Find([]string{"a", "b", "c"}, func(s string) bool {
        return s == "b"
    })
    // isFound == true
}
```

## Dumper

Dumper simulates [`Data::Dumper`](https://metacpan.org/pod/Data::Dumper) by
stringifying data structures. It's just a wrapper around
[go-spew](https://github.com/davecgh/go-spew) but hard coded to never call
`String()` or `Error()` if those are implemented.

```go
type X struct {
    A int
    B int
    C int
}

func (x *X) String() string {
    return strconv.Itoa(x.A + x.B + x.C)
}

func dump() {
    x := X{2, 4, 6}

    spew.Dump(x)
    // (main.X) 12

    data.Dumper(x)
    // (main.X) {
    //  A: (int) 2,
    //  B: (int) 4,
    //  C: (int) 6
    // }
}
```
