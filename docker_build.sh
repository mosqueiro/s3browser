docker build --load --platform linux/amd64 -t s3browser .
docker tag s3browser mosqueiro/s3browser
docker push mosqueiro/s3browser
