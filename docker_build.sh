# docker build --load -t s3browser .
# docker run -p 8080:3000 s3browser
docker build --load --platform linux/amd64 -t s3browser .
docker tag s3browser mosqueiro/s3browser
docker push mosqueiro/s3browser
