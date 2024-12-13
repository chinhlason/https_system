# How to gen ssl cert and key if you are using windows

1. run an alpine container by command `docker run -it --name genssl-container alpine:latest /bin/sh`
2. download openssl by command `apk add --no-cache openssl`
3. make a directory to store the cert and key by command `mkdir /ssl`, cd to that directory by command `cd /ssl`
4. generate a private key by command `openssl genrsa -out private.key 2048`
5. generate a csr by command `openssl req -new -key private.key -out cert.csr` (in Common Name you need to put your domain name, if you run app locally you can put "localhost"), after that you can use this csr to gen a public cert or private cert
6. in this case, I will generate a self-signed cert by command `openssl x509 -req -days 365 -in cert.csr -signkey private.key -out public.crt`, note that this cert will not be trusted by browser, I suggest you to buy a cert from trusted CA
7. finally, you need copy private key and certificate from container to your host device by command `docker cp genssl-container:/ssl/ ./server` and `docker cp genssl-container:/ssl/ ./client` (don't push this ssl folder to public repo)

# Run this project
1. After you have ssl folder, you can change directory to 'client' and 'server', easily run this project by command `docker-compose up -d --build`
2. The backend server will run on port 3000, and frontend server will run on port 4002, you can access the frontend server by open browser and go to `https://localhost:4002`
3. If you want to see more details about the config, you can check the `Dockerfile` in each folder
