# Langding pages
Programming languages that include the syntax on their landing page.

## Syntax
```toml
[go]
name = "Go"
url = "go.dev"
date = 2025-03-29
syntax = true
scroll = true
interactive = true

[toml]
name = "TOML"
url = "toml.io"
date = 2025-03-29
syntax = true
scroll = false
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
| scroll      | boolean    | no       | If `syntax=true`, whether the user must scroll to see the syntax.                      |
| interactive | boolean    | no       | Whether the landing page includes a widget for interacting with the language directly. |

- `url` is required; if a language does not have an authoritative website, then
  it is not to be included.
- The most simple form of a URL that will sucessfully navigate to the langauge's
  landing page is preferred. For example, `toml.io` versus `https://toml.io/`.
- `scroll` is based on a reasonable standard screen size of 1920x1080 pixels
  with a Device Pixel Ratio of 1.
- `interactive=true` usually implies `syntax=true`.
