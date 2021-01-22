# Concurrency and CLI

1. Create new repository called `go-cnc-cli-training` and initialize go modules. Also copy paste `.github` directory from root of this project to add CI to the project.

2. You need to develop simple message processor that can send messages via phone, email or both of those delivery systems. This CLI tool wil take path to the `messages.json` file and send all the messages from that file to the receivers. An example of calling such service would look like this (use `flag` package to get arguments):
    ```go
    go run main.go -f ./testdata/messages.json
    ```
    Please refer to `messages.json` file for example dataset.

    Couple of requirements before you started:
    
    * `phone` and `email` workers should run as a separate go routine, meaning there will be only two workers `phoneWorker` and `emailWorker`.
    * if you encounter the message with `all` value in `by` property (look at `messages.json` for reference) you should send that message to both `email` and `phone` worker.
    * assume that you don't know the amount of messages, use unbuffered channels and wait groups to finish this task.

3. Here's some pseudo code to help you grasp the task in case you need help:

    ```pseudo
    This program will spin up two workers
    one for email and one for phone messages.

    function emailWorker(messages, waitGroup)
      do your processing here don't forget to use waitGroup
    end

    function phoneWorker(messages, waitGroup)
      do your processing here don't forget to use waitGroup
    end

    function main()
      first receive -f flag from the CLI
      
      then read all the messages from the file and parse JSON

      create your wait group and channels

      spin up your workers (don't forget to increment
      waitGroup before spinning up every worker)

      send all you messages (don't forget to check how the message 
      needs to be delivered before sending it in the channel)

      close your channels and set waitGroup
      to wait till all workers is finished
    end

    ```