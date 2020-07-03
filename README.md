# proxy [![Build Status](https://drone.lngzl.nl/api/badges/filecoffee/proxy/status.svg)](https://drone.lngzl.nl/filecoffee/proxy)
Use your own domain with file.coffee

# How will it work?
The proxy will make requests to the file.coffee server and act as a seperate webserver serving the file.coffee content. It's dockerized and will in the future have caching and automatic SSL certificates.

# Todo
- [ ] Add a caching system, like Varnish.
- [ ] Making it automatically generate Let's Encrypt certificates.

# Requirements
* A server with Docker and Docker compose

# Setup
### 🐳 Deploy it to port 80 with NGINX
To just run the instance, simply do `docker-compose up` in the root of this project after cloning it. If you want to deploy it you can simply run `docker-compose up -d` and it will deploy in the background.

### 🛠 Manual setup (not all features)
1. Go to the `src` folder
2. `go mod download` to collect all the dependencies
3. `go run main.go`
4. It's running

# ShareX
It's really easy to use the proxy instance with ShareX, simply edit the `sharex-config.sxcu` file in the root of the project.

The only thing you have to change is replace `https://your-domain.com/` with your own domain and you're ready to use it.

*It uses the normal file.coffee domain for uploading to ensure everything gets fully uploaded. It's **not** recommended that you change this since there's not feature in the proxy that handles file uploading.*