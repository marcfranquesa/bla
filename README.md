# Bla 🗣

URL shortener, sample usage:

```sh
curl -X POST https://bla.cat -d 'https://github.com/marcfranquesa/bla'
```

Also accepts JSON:

```sh
curl -X POST https://bla.cat \
  -H "Content-Type: application/json" \
  -d '{"url":"https://github.com/marcfranquesa/bla"}'
```

