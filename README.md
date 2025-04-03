# Langding pages
Programming languages that include the syntax on their landing page.

## Syntax
```go
{
	Name:        "Go",
	URL:         "go.dev",
	Date:        "2025-03-29",
	Syntax:      true,
	NoScroll:    false,
	Interactive: true,
},
{
	Name:        "TOML",
	URL:         "toml.io",
	Date:        "2025-03-29",
	Syntax:      true,
	NoScroll:    true,
	Interactive: false,
},
```

Languages are defined in [languages.go](languages.go) as a [Go](https://go.dev)
slice.

Each entry in the slice defines information for a particular language.

Each entry has the following keys:

| Key         | Type   | Description                                                                            |
|:------------|:-------|:---------------------------------------------------------------------------------------|
| Name        | string | The name of the language in its canonical form.                                        |
| URL         | string | A URL to the language's landing page.                                                  |
| Date        | string | The date (formatted "2006-01-02") when the landing page was last checked.              |
| Syntax      | bool   | Whether the landing page displays the syntax of the language.                          |
| NoScroll    | bool   | If `Syntax=true`, whether the user does not need to scroll to see the syntax.          |
| Interactive | bool   | Whether the landing page includes a widget for interacting with the language directly. |

- For a landing page to qualify as `Syntax=true`, the intention should be to
  showcase the syntax itself, rather than being a part of the design of the
  page.
- `URL` is required. The landing page should be for the language specifically.
  If a language does not have an authoritative website, then it is not to be
  included.
- The most simple form of a URL that will sucessfully navigate to the langauge's
  landing page is preferred. For example, `toml.io` versus `https://toml.io/`.
- `NoScroll` is based on a reasonable standard screen size of 1920x1080 pixels
  with a Device Pixel Ratio of 1.
- An interactive widget is usually prefilled with an example, so
  `Interactive=true` usually implies `Syntax=true`.
