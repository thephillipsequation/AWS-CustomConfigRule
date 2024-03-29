## Quick Demo of AWS Custom Config Rules using GO

This project will build and deploy:
* a pipeline for deploying custom AWS Config rules
* a sample lambda / config rule written in golang

I've used this as a boilerplate for many projects over the years, feel free to steal the pattern and make it your own :)

## Requirements
You must have an environment variable set for `GIT_TOKEN` that is equal to an OAUTH token capable of cloning this (or your own hosted) repo in order for the pipeline to work. 

For more info on setting up an OAUTH token please visit the following [page](https://help.github.com/articles/creating-a-personal-access-token-for-the-command-line/).

NOTE: if you are using `OSX` you should `brew install watch` for this script to work properly


## Directions to Launch the User Account Pipeline
To launch the pipeline stack
` sh pipeline.sh`
To update the pipeline stack
`sh pipeline.sh update`
To delete the pipeline stack
`sh pipeline.sh delete`
