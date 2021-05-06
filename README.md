# proxy
Use your own domain with file.coffee

# How will it work?
The proxy will make requests to the file.coffee server and act as a seperate webserver serving the file.coffee content. It's dockerized and will in the future have caching and automatic SSL certificates.

# Todo
- [ ] Making it automatically generate Let's Encrypt certificates.
- [x] ~~Add a caching system, like Varnish.~~ Made a simple cache system that saves the file after the first request for as long as needed.
- [x] ~~The proxy is currently a bit slower than using file.coffee directly, we're still looking into methods to speed it~~ Using Caching, the 2nd request will be a lot faster.
up.

# Requirements
* A server with Docker and Docker compose

# Setup
### 🐳 Deploy it to port 80 with NGINX (recommended)
Copy `.env.example` to `.env` and fill in the file.

To just run the instance, simply do `docker-compose up` in the root of this project after cloning it. If you want to 
deploy it you can simply run `docker-compose up -d` and it will deploy in the background.

### 🛠 Manual setup (not recommended)
1. Go to the `src` folder
2. `go mod download` to collect all the dependencies
3. `go run main.go`
4. It's running

# ShareX
It's really easy to use the proxy instance with ShareX, simply edit the `sharex-config.sxcu` file in the root of the
 project.

The only thing you have to change is replace `https://your-domain.com/` with your own domain and you're ready to use it.

*It uses the normal file.coffee domain for uploading to ensure everything gets fully uploaded. It's **not** recommended
 that you change this since there's not feature in the proxy that handles file uploading.*

# Privacy, Safety & Security
Depending on the settings used in te proxy, it might store data for instead of just proxy. This is to make the proxy 
faster but also makes it that the files _could_ be accessed if there was direct access to your proxy. We highly recommend
only using the proxy for yourself as a while-label type way of using our platform.

When using Docker Compose, it creates an internal network between the containers with the NGINX server being the only 
of public exposure. It's highly recommended to keep it this way and not deploying it manually for security reasons.

# License
This repository is using the MIT license.
