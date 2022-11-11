# quizlog-go

Backend server of `Quizlog`.

Project for graduation in `Design Projects(Capstone Design)` in `CAU`.

## About Quizlog

`Quizlog` is Quiz platform based on user's posts on `Tistory`. Users can make their own Quiz for self-study.

Try in [Quizlog Web App](http://quizlog-react.s3-website.ap-northeast-2.amazonaws.com/)

## Project Stack

### Frontend

- `javascript`
- `react.js`

### Backend

- `golang`
- `fiber`
- `ent`
- `facebookgo/inject`

## Setup

`config.yaml` should be placed in project root directory.
Follow below form:
```yaml
auth:
  secret_key: "SECRET_KEY_FOR_JWT_ENCRYPT"
tistory:
  client_id: "TISTORY_CLIENT_ID"
  client_secret: "TISTORY_SECRET_ID"
  redirect_uri: "REDIRECT_URI_FOR_OAUTH"
```

### Auth
`Quizlog` use `jwt` token for authentication, and use `HS256` algorithm for enryption and decryption. So `secret key` for `HS256` is needed in `config.yaml`

### Tistory Client

Needs below options for `Tistory Blog API`.:
- `client_id` : Tistory App ID.
- `client_secret` : Tistory App Secret Key.
- `redirect_uri` : redirection URI after Authentication.

Register your app in [Tistory](https://www.tistory.com/), and set in `config.yaml`.