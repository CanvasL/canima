# canima

## API
|Purpose|URI|Method|SC|
|-|-|-|-|
|register|/user|POST|201, 400, 500|
|login|/user/:username|POST|200, 400, 500|
|getUserInfo|/user/:username|GET|200, 400, 401, 403, 500|
|logout|/user/:username|DELETE|204, 400, 401, 403, 500|