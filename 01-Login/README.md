# Auth0 + Go Web App Sample + Google App Engine

## Running the App

To run the app, make sure you have **go** and **goget** installed.

Download SDK for your platform from here: `https://developers.google.com/appengine/downloads#Google_App_Engine_SDK_for_Go`

Update the app.yaml and provide your Auth0 credentials.

```yaml
application: auth0-golang-web-app
version: 1
runtime: go
api_version: go1

handlers:
- url: /public
  static_dir: public

- url: /.*
  script: _go_app

env_variables:
  AUTH0_CLIENT_ID: "{CLIENT_ID}"
  AUTH0_DOMAIN: "{DOMAIN}"
  AUTH0_CLIENT_SECRET: "{CLIENT_SECRET}"
  AUTH0_CALLBACK_URL: "http://localhost:8080/callback"
```

Once you've set your Auth0 credentials in the `.env` file, run `go get .` to install the Go dependencies.

Run `dev_appserver.py app.yaml` to start the app and navigate to [http://localhost:8080/](http://localhost:8080/)
