# shepherd
oidc, auth2.0 repository for external package comparison

# application

An authentication authorization application was created for each oidc external package.
Both use https://pkg.go.dev/golang.org/x/oauth2 for the auth2.0 part.

- coreos
    - OIDC building application using https://github.com/coreos/go-oidc package.
- zitadel
    - OIDC building application using https://github.com/zitadel/oidc package.

Both projects are also created with reference to the server project configuration in the [basic-package](https://go.dev/doc/modules/layout#basic-package).

# 

```
docker compose build

docker compose up
```

|application|url|
|---|---|
|coreos|http://localhost:8080|
|zitadel|http://localhost:8081|

