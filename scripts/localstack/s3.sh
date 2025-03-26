#!/bin/bash

# Create S3 bucket for inventory
awslocal s3api create-bucket --bucket inventory
