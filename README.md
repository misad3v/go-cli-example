# go-cli-example

## Build
In order to make this lab works, run the following commands:
```python
cd helpers/docker
docker-compose -f docker-compose.yaml up
```

After ran the `docker-compose` you can perform:
1. get into the container terminal: `docker exec -it example-app bash` 
2. init the `go.mod`: `go mod init github.com/misad3v/go-cli-example.git`

## Cobra
To add a Command Line Interface to a project, do the following:
1. Make sure you add the cobra lib: `go install github.com/spf13/cobra/cobra@latest` this was installed when we built the container 
2. Then, initialize the cobra: `cobra init` in the root folder.
3. This will create a `cmd` folder and a `root.go` file inside it. 
4. To add new commands to an existed CLI: `cobra init <command-name>` like: `cobra init bmake`
<br>
[click here to get more examples](https://github.com/spf13/cobra)