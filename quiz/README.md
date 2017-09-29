### Problem1: Bash Shell Script sample:

Implement a script which has the following usage format:

```bash
./bash_example [--help|-h]
./bash_example --create-file=<filename> [--no-prompt] [--verbose]
```
  This script should create a file with the name provided in the args with <filename>. This file should contain all 50 states with one state per line. This should be done in following scenarios:
  * File <filename> does not exist or --no-prompt option is used
  * In case the file exists, this script should prompt the user:
  ```bash
    “File exists. Overwrite (y/n) ?”
  ```
    The script should process user input and either overwrite file or not.
    If the user input is not valid, i.e. neither ‘y’ nor ’n’, the script should just prompt the user again with the same prompt until a valid input is provided.

    If the --verbose option is set, the following messages should be printed in appropriate portions of code.

    ```bash
File removed
File created
File already exists
```
**Return Codes:**
 * In case of bad command line args, the help showing usage should be printed and returned with error (non zero code).
 * In all successful case return code should be 0

### Problem2: Python/Go sample:

Implement the following as described below in either go or python:

```bash
uniquify [--help|-h]
uniquify --file=<filename> --output=<output-filename> [—verbose]
```

This script/program should input a file with the name provided in the args called <filename>. This input file can contain millions of **lines** with duplicates.

The script/program should output a file with name provided in the args called <output-filename> with all duplicate **lines** removed.

A line is delimited by \n or \r or both

### Problem3: Go sample:

You will need to implement a program in go as described below:

```bash
gosample [-help|-h]
gosample -urls=<comma-seperated-one-or-more-urls>
```

This binary should be able to read multiple urls and generate a word (a word only contains letters A-Za-z0-9) frequency table. Output can be stored in files with names url1.txt, url2.txt … in the following format:
```bash
  url: <url>
    word1: count
    word2: count
```

Example:

See samples/problem3/url1.txt

**[Extra Credit]**

Handle URLs in parallel and demonstrate reduced time to process all URLs. Note: You only need to implement serial or parallel but not both.

## Submitting your responses for the tests

We will accept submissions as a github repo. For this, you would need the following:

Have an account on github.com
Have a repository called aporeto which is public
Have directories called
```bash
samples/problem1
samples/problem2
samples/problem3
```
Use README.md files in this directory to provide insights into your thought process if any.

## Here are things you would like to pay attention to:

  * Code samples should be written not just for functionality but also readability
  * Pay attention to the algorithms implemented and the order (Big-O notation) that your implementation provides. These are not just functionality and readability but also performance.
  * Test your code for all required scenarios.

## Compiling/Dependency resolution

**go code**

  * We will be running "go get ./..." and "go build"

**python code**

  * We will be using python 2.7. If you want me to install specific dependencies, please provide a shell script called download.sh and requirements file called requirements.txt in the following format

  ```bash
  #!/bin/bash

  pip install -r requirements.txt
  ```
