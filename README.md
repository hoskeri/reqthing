## What

ReqThing - Thing that makes requests.

## Why

Like curl, with some extras:

- persistent connections to servers.
- cache responses, transparently
- watch url for changes.

## How

```bash
reqthing \
  -state "/var/cache/reqthing.0"
  -url "https://www.example.com/download.mp4"

```
