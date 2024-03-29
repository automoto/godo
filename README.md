## GODO

Simple command line client for generating markdown TODO lists.

You can have the markdown TODO list generated as a file in the directory of your choosing or have it printed to your terminal for easy copy and pasting.

### Building and Installation
Clone this repository and run the following command inside the repo:

`go build -o godo output_util.go main.go`

This will create a binary locally you can run commands against already, like so:

`./godo --td "get paid, get fed, get sleep"`

Copy the binary to your local bin to access the CLI anywhere.

`sudo cp godo /usr/local/bin`

Now you can just do plain `godo` commands:

`godo --td "be smart, be fast, be precise"`


### Configure directory for your markdown TODO list files

**Note:** If no directory for outputting your markdown files is set or found, the markdown will be printed out to the terminal.

Setting the directory for your markdown files to save to is very handy for pointing `godo` to a Dropbox or Google Drive folder for easy access.

You have two options:

Pass in the output directory via the `-p` command line option, like so:

`./godo -td="buy food, get gas" -p "/home/your_path_goes_here""`

Or do it by setting the `TODO_PATH` environment variable:

`export TODO_PATH="/home/your_path_goes_here"`

**Note:** Any directories passed in via the `-p` option take precedence over the value set in the `TODO_PATH` environment variable.

### Copy TODO list to your clip board

In ubuntu we can copy our generated TODO list to the clipboard using `xclip`, like so:

`./godo -td="get paid, get fed, get sleep" | xclip -sel clip`

### Tests

Make sure you can do a `dep ensure` to pull in all the packages needed for running the tests

Run the tests:

`go test -v`