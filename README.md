#build
docker build --no-cache -t example-api -f Dockerfile .
docker run -d  -p 8080:8080 example-api