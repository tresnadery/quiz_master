# Manual Installation

## Requirement

* golang installed

* mysql database

# Install the application
This script must be run on the root of the folder

``` ./bin/setup.sh ```

# How to run the command line

``` ./bin/quiz_master [command] [arg] [flag] ```

# List Command

List Question

``` ./bin/quiz_master list_question```

Detail Question

``` ./bin/quiz_master question <number> ```

Create Question

``` ./bin/quiz_master create_question <number> <question> <answer>```

Answer Question

``` ./bin/quiz_master answer_question <number> <answer>```

Delete Question

``` ./bin/quiz_master delete_question <number>```