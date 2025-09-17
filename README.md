# GLUI

Gitlab (terminal)ui to navigate gitlab

// to build
## Features

*basics*

list the pipelines
list the merge request
list the issues
list the jobs
list the log from the job with auto refresh

*use case*
navigate from merge request to pipeline to job to log
create a new pipeline with custom inputs on a branch to test things
create a merge request from the console using the gitlab template with the merge request title standard configured in this repo
easy navigate to the favorite repos
set a monitor on the pipeline of a merge request and get a notification if success or fail
cli integration for non-interactive ai e.g. Amazon q usage, clear --help to get directly what you want
download/see artifacts from pipelines for easy debugging

*user/developer interface*
same ui as k9s
easy navigation using the keyboard
checks for the height and the width automatically
uses gitlab api and uses the key from the .env file
? need some help where to store the keys

## Architecture

uses a core, with an interactive terminal ui but also the commands can be executed on the commandline
probably using go

## Planning

0. let's first make a design to cover all features, identify missing features, keep a TODO.md
1. make a simple set up using go, make the basic navigation work, see k9s documentation
2. make the core menu supporting all features, but with mock functionality
3. make the connection with real gitlab, support only cloud for now, on prem, let's do that later
4. make small improvements, stick to the plan
5. on every change create a small log file

inspiration:
- ~/git/hub/k9s
- ~/git/hub/glab-tui
