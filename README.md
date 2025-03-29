# Langding pages
Programming languages that include the syntax on their landing page.

## Syntax
```toml
[go]
name = "Go"
url = "go.dev"
date = 2025-03-29
syntax = true
noscroll = false
interactive = true

[toml]
name = "TOML"
url = "toml.io"
date = 2025-03-29
syntax = true
noscroll = true
interactive = false
```

Languages are defined [languages.toml](languages.toml), which is formatted in
[TOML](toml.io).

Each table in the file defines information for a particular language. The header
must unambiguously identify the langauge. All lowercase letters is preferred.

Each table has the following keys:

| Key         | Type       | Required | Description                                                                            |
|:------------|:-----------|:---------|:---------------------------------------------------------------------------------------|
| name        | string     | yes      | The name of the language in its canonical form.                                        |
| url         | string     | yes      | A URL to the language's landing page.                                                  |
| date        | local date | yes      | The date when the landing page was last checked.                                       |
| syntax      | boolean    | yes      | Whether the landing page displays the syntax of the language.                          |
| noscroll    | boolean    | no       | If `syntax=true`, whether the user does not need to scroll to see the syntax.          |
| interactive | boolean    | no       | Whether the landing page includes a widget for interacting with the language directly. |

- For a landing page to qualify as `syntax=true`, the intention should be to
  showcase the syntax itself, rather than being a part of the design of the
  page.
- `url` is required. The landing page should be for the language specifically.
  If a language does not have an authoritative website, then it is not to be
  included.
- The most simple form of a URL that will sucessfully navigate to the langauge's
  landing page is preferred. For example, `toml.io` versus `https://toml.io/`.
- `noscroll` is based on a reasonable standard screen size of 1920x1080 pixels
  with a Device Pixel Ratio of 1.
- An interactive widget is usually prefilled with an example, so
  `interactive=true` usually implies `syntax=true`.
