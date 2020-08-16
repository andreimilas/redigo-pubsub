# redigo Pub/Sub
## Author
Andrei Milas

## Description
Pub/Sub example written using https://github.com/gomodule/redigo

## Running the project
1. Run ```docker-compose up``` for starting a Docker container with Redis listening on port 6379
2. Start a subscriber by running ```go run subscribe/subscribe.go <<channel_name>>```; if ```<<channel_name>>``` is not specified, the subscriber will listen to the default channel -- ```sample```
3. Publish messages by running ```go run publish/publish.go <<channel_name>> <<message>>``` in a different terminal; if ```<<channel_name>>``` or ```<<message>>``` are not specified, the default values will be used -- ```sample``` for channel name and ```hello``` for message.
4. You should see the published messages logged in the subscriber terminal.
