# full-stack-crud
GoLang backend which serves a Vue SPA and provides a simple REST API to interact with a Postgres database. The whole bundle is containerized with Docker.

* The Go backend utilizes Gin to serve the API and static files, and GORM to handle database transactions. 
* The Vue SPA is built into a single `index.html` file, which contains the necessary JS and CSS. This is good for very simple applications, but it's a bad design choice for bigger pages, as it'll take longer to load.

## Building and running
Start the application by running `docker compose up --build` inside the root directory. The server listens on `localhost:8080` by default.
