# goff-id

Server to generate unique IDs, implemented in Go ([golang](https://golang.org)).

## APIs

- `GET /sf64`: generate a Twitter Snowflake 64-bit id as an integer.
- `GET /sf64hex`: generate a Twitter Snowflake 64-bit id as a hex (base 16) string.
- `GET /sf64ascii`: generate a Twitter Snowflake 64-bit id as an ascii (base 36) string.
- `GET /sf128`: generate a Twitter Snowflake 128-bit id as an integer.
- `GET /sf128hex`: generate a Twitter Snowflake 128-bit id as a hex (base 16) string.
- `GET /sf128ascii`: generate a Twitter Snowflake 128-bit id as an ascii (base 36) string.

(output is JSON-encoded).

## License

The MIT License (MIT). See [LICENSE.txt](LICENSE.txt).

## History

Current version: `1.0.0`

### 2019-01-04 - v1.0.0

First release:

- Generate Twitter's Snowflake IDs:
  - 64-bit & 128-bit IDs.
  - integer, hex (base 16) and ascii (base 36) formats.
