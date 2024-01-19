# go-cli
Inspired by Powerful Command-Line Applications in Go 

## TODO:
### 02-CLI
• Update the custom usage function to include additional instructions on
how to provide new tasks to the tool.
• Include test cases for the remaining options, such as -complete.
• Update the tests to use the TODO_FILENAME environment variable instead of hard-coding the test file name so that it doesn’t cause conflicts with an existing file.
• Update the getTask() function allowing it to handle multiline input from
STDIN. Each line should be a new task in the list.