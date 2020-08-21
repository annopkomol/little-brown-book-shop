# Little Brown Book Shop

## Database Migration Guide (MySQL)

 1. Install migration cli 
[https://github.com/golang-migrate/migrate/tree/master/cmd/migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)

 2. Go to script directory using `cd ./db/script/` and create database config file<br> 
 `touch db_conn_url.txt`.
 
 3. Edit `db_conn_url.txt`, see example in `db_conn_url.example.txt` file.
 
 4. Run migration script using command `sh run_migration.sh` and follow the instruction.
 
## Deployment

#### API
1. Set Environment Variables on your environment, see example at `.env.example` inside `./api/` path

2. For Production Environment, use `Dockerfile` instead of `Dockerfile.dev` to build Docker image (minimize final image size and
 cut off unnecessary features such as hot-reloading
 
3. Upload your image to your prefer Container Registry and deploy from that image.

#### Front-end
1. `touch env.production.local` in `./web/` and set `VUE_APP_API_URL={{your_api_base_url}}` 
2. Run `yarn build` inside `./web/` path, the published directory will be store in `./web/dist/`
3. Serve `./dist` directory by any HTTP Server of your choice.
4. Configure your HTTP server to fallback to index.html for any requests that do not match a static file.<br> 
[Click Me](https://router.vuejs.org/guide/essentials/history-mode.html) for more info.


## Local Development

1. `touch .env` at `./api/` path, see example in `.env.example`
2. `touch env.development.local` at `./web/` and set `VUE_APP_API_URL=localhost:{{your_api_port}}`
3. run `docker-compose build`
4. run `docker-compose up`