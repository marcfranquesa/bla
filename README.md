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

The shortened URL will be in the response body, and a deletion token will be provided in the 'X-Token' response header. You can view both using:

```sh
curl -X POST https://bla.cat -d 'https://github.com/marcfranquesa/bla' -i
```

To delete a shortened URL, use the DELETE method with the URL's ID and the deletion token:

```sh
curl -X DELETE https://bla.cat/abc123 -H "Authorization: token"
```
