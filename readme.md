### Generate docs

In order to generate openapi in `/openapi`, we'll have to do this:
- OpenApi spec is `openapi/apispec.json` (OpenAPI 3).
- Install redoc with NPM: `npm i -g redoc-cli`.
- Generate the openapi to: `redoc-cli bundle -o openapi/index.html openapi/apispec.json`.
- Generated openapi is `openapi/index.html`.

Tools:
* For vs code users, install OpenApi plugin: https://marketplace.visualstudio.com/items?itemName=42Crunch.vscode-openapi
* Swagger Online Editor: https://editor.swagger.io/