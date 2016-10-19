# hello-world
## Just a tiny reference golang app that uses Alpine as it's base OS image

### Compile to binary:

sh build.sh

### Build Docker image:

docker build -t go-hello-world .

### Run Docker image:

docker run  -p 8080:8080 -i -t go-hello-world

### (optional) Apply ecr tag to it

docker tag go-hello-world:latest 563743801656.dkr.ecr.us-east-1.amazonaws.com/turner-ca-sandbox:latest

### or a versioned tag

docker tag go-hello-world:latest 563743801656.dkr.ecr.us-east-1.amazonaws.com/turner-ca-sandbox:v2


## push it to aws registry


aws ecr get-login --region us-east-1
### run docker login with that result above
docker push 563743801656.dkr.ecr.us-east-1.amazonaws.com/turner-ca-sandbox:v2


# docker run  -e "MODEL_CONTROLLER=pal" -e "ENVIRONMENT=prod" -p 8000:5000 -i -t 23eaf9d7951d  


MODEL_CONTROLLER
