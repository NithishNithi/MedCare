FROM ubuntu:latest

COPY . .

RUN apt-get update && apt-get install -y ca-certificates

EXPOSE 5000
EXPOSE 5005

CMD [ "./medcare" ]


# Docker Image run Command
# docker run -d --name medcare-container -p 5000:5000 --network medcare-network medcare 