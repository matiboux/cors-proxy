# CORS Proxy

Simple CORS proxy server to run locally that passthrough all requests to the target server to bypass CORS restrictions of your browser.


## Usage

Start the CORS Proxy server locally at `http://localhost` (port 80) using Docker:

```bash
docker run -p 80:8080 matiboux/cors-proxy
```

Replace `-p 80:8080` with `-p <port>:8080` to use a different port and access the server at `http://localhost:<port>`.

You can pull and run the CORS Proxy Docker image from either:
- The [Docker Hub](https://hub.docker.com/r/matiboux/cors-proxy): `docker pull matiboux/cors-proxy`
- The [GitHub Container Registry](https://github.com/matiboux/cors-proxy/pkgs/container/cors-proxy): `docker pull ghcr.io/matiboux/cors-proxy`


## Development

### CORS Proxy

Use this command to run the program locally for development:

```sh
docker compose watch
# or: docker compose up
```

Using `watch`, you'll benefit from file changes watching for sync & rebuild.

Use [DockerC](https://github.com/matiboux/dockerc) for shortened commands: `dockerc - @w`.

### Documentation

Use this command to run the documentation site locally for development:

```sh
docker compose -f docker-compose-docs.yml -f docker-compose-docs.override.yml watch
# or: dockerc docs.override @w
```


## License

Copyright (c) 2024-2025 [Matiboux](https://github.com/matiboux) ([matiboux.me](https://matiboux.me))

Licensed under the [MIT License](https://opensource.org/license/MIT). You can see a copy in the [LICENSE](LICENSE) file.
