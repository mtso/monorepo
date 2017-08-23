# gateway

tackdb.com gateway server.

## Notes

Should keep user session by cookie, and cache JWTs from auth server.
Then attach JWT to the header when passing requests to Cloud API.

If the API request is sent with BasicAuth, try to authenticate with BasicAuth
