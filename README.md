# go-mstdn

Package **mstdn** provides tools for working with the **Mastodon API**, which is becoming the (defacto) stand client-server API for the **Fediverse**.
Note that the **Mastodon API** is not (yet) officially part of **ActivityPub**.

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/sourcecode.social/reiver/go-mstdn

[![GoDoc](https://godoc.org/sourcecode.social/reiver/go-mstdn?status.svg)](https://godoc.org/sourcecode.social/reiver/go-mstdn)

## Import

To import package **mstdn** use `import` code like the follownig:
```
import "sourcecode.social/reiver/go-mstdn"
```

To import the Mastodon API **entities** use `import` code like the following:
```
import "sourcecode.social/reiver/go-mstdn/ent"
```

To import the Mastodon API **entities** for **administrators** use `import` code like the following:
```
import "sourcecode.social/reiver/go-mstdn/ent/admn"
```

## Installation

To install package **mstdn** do the following:
```
GOPROXY=direct go get https://sourcecode.social/reiver/go-mstdn
```

## Author

Package **mstdn** was written by [Charles Iliya Krempeaux](http://changelog.ca)
