# notification-service

Pub/Sub based notification service API for sending email and mobile push notifications. Developed using Go / Gin, AWS SES, SNS, S3.

<br/>
<br/>

## Directory structure

The directory structure is as follows:

- **conf/**  
  - Configuration files for the service.
  
- **email/**  
  - Handlers and configurations for email notifications using AWS SES.

- **models/**  
  - Data models for the notification service.

- **notification/**  
  - Core logic for sending notifications.

- **routes/**  
  - API route definitions for the notification service.

- **utils/**  
  - Utility functions used across the service.

- `main.go`: Entry point of the Go-based notification service.
- `go.mod` & `go.sum`: Dependency management files.

<br/>
<br/>

## Overview

### Design

The service is used mainly with a monitoring and logging service. The monitoring service can be found <a href="https://www.sitemonitoring.io/">here</a>. Similar services can be found <a href="https://whimsical.com/web-microservices-6uqvwWZtcBFsNJB2hepGy1">here</a> and below:

#### Similar services

<img width="834" alt="image" src="https://github.com/user-attachments/assets/b54088e7-870c-46dd-9cf6-2e5ec27d9d5c">
