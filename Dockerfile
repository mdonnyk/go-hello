FROM alpine:latest 
WORKDIR /app
EXPOSE 8080 
COPY . /app
CMD /app/hello 
