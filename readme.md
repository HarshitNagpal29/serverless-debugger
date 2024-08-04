# Serverless Debugger

Serverless Debugger is a web application designed to facilitate the debugging of serverless functions deployed on AWS Lambda and Google Cloud Platform (GCP). It offers features to list functions, invoke them, update their code, fetch logs, and add breakpoints for debugging purposes.

## Table of Contents

- [Serverless Debugger](#serverless-debugger)
  - [Table of Contents](#table-of-contents)
  - [Features](#features)
  - [Prerequisites](#prerequisites)
  - [Project Structure](#project-structure)
  - [Setup Instructions](#setup-instructions)
    - [Backend Setup](#backend-setup)
    - [Frontend Setup](#frontend-setup)
  - [Running the Application](#running-the-application)
  - [Usage](#usage)
    - [AWS Functions](#aws-functions)
    - [GCP Functions](#gcp-functions)
    - [Function Logs](#function-logs)
    - [Debugger](#debugger)
  - [API Endpoints](#api-endpoints)
  - [Contributing](#contributing)

## Features

- **List Functions**: Retrieve and display a list of serverless functions from AWS Lambda and GCP.
- **Invoke Functions**: Invoke serverless functions and display the output.
- **Update Functions**: Update the code of serverless functions.
- **Fetch Logs**: Retrieve and display logs for serverless functions within a specified time range.
- **Add Breakpoints**: Add breakpoints to serverless functions for debugging purposes.

## Prerequisites

Before setting up the project, ensure you have the following installed on your machine:

- Go (for backend development)
- Node.js and npm (for frontend development)
- Docker (optional, for containerization)
- An IDE or text editor of your choice (e.g., VS Code)

## Project Structure

The project is organized into backend and frontend directories:

```
serverless-debugger/
├── backend/
│   ├── handlers/
│   │   ├── debuggerhandlers/
│   │   │   └── debuggerHandlers.go
│   │   ├── handlers.go
│   │   └── log_handlers/
│   │       └── log_Handlers.go
│   ├── pkg/
│   │   ├── aws/
│   │   │   └── awsClient.go
│   │   ├── gcp/
│   │   │   └── gcpClient.go
│   ├── main.go
│   └── go.mod
├── frontend/
│   ├── public/
│   │   ├── index.html
│   │   └── ...
│   ├── src/
│   │   ├── components/
│   │   │   ├── AWSFunctions.jsx
│   │   │   ├── GCPFunctions.jsx
│   │   │   ├── FunctionLogs.jsx
│   │   │   └── Debugger.jsx
│   │   ├── api.js
│   │   ├── App.jsx
│   │   ├── index.jsx
│   │   └── ...
│   ├── package.json
│   └── ...
├── .gitignore
└── README.md
```

- **backend/**: Contains Go code for backend server, including handlers for HTTP requests (`handlers/`), log handlers (`log_handlers/`), and service implementations (`pkg/`).
- **frontend/**: Contains React code for frontend UI, including components (`src/components/`), API functions (`src/api.js`), and main application files (`src/App.jsx`, `src/index.jsx`).

## Setup Instructions

Follow these steps to set up and run the Serverless Debugger application locally.

### Backend Setup

1. **Clone the repository:**

   ```bash
   git clone https://github.com/HarshitNagpal29/serverless-debugger.git
   cd serverless-debugger

2. **Install dependencies:**

    Ensure you have Go installed. Then, navigate to the backend/ directory and install dependencies using:

   ```bash
   go mod tidy

3. **Set up environment variables:**

   ```bash
   PORT=8080
   AWS_REGION=your_aws_region
   AWS_ACCESS_KEY=your_aws_access_key
   AWS_SECRET_ACCESS_KEY=your_aws_secret_access_key
   GCP_Project_ID=your_gcp_ProjectId
   GCP_Credentials_File=your_Credentials_File

4. **Run the backend server:**

   ```bash
   go run main.go
   


## Usage

### AWS Functions

1. **View AWS Functions:**
   * Navigate to the AWS Functions section.
   * A list of AWS Lambda functions will be displayed.

2. **Invoke AWS Function:**
   * Click the "Invoke" button next to the desired function.
   * The function will be invoked, and the response will be logged in the console.

3. **Update AWS Function:**
   * Click the "Update" button next to the desired function.
   * Update the function's code through the provided input fields.
   * The function will be updated, and the response will be logged in the console.

### GCP Functions

1. **View GCP Functions:**
   * Navigate to the GCP Functions section.
   * A list of GCP Cloud Functions will be displayed.

2. **Invoke GCP Function:**
   * Click the "Invoke" button next to the desired function.
   * The function will be invoked, and the response will be logged in the console.

3. **Update GCP Function:**
   * Click the "Update" button next to the desired function.
   * Update the function's code through the provided input fields.
   * The function will be updated, and the response will be logged in the console.

### Function Logs

1. **Fetch Logs:**
   * Navigate to the Function Logs section.
   * Select the service (AWS or GCP), and enter the function name, start time, and end time.
   * Click the "Fetch Logs" button.
   * The logs will be displayed in a list format.

### Debugger

1. **Add Breakpoint:**
   * Navigate to the Debugger section.
   * Select the service (AWS or GCP), and enter the function name, file name, and line number.
   * Click the "Add Breakpoint" button.
   * The breakpoint will be added, and the response will be logged in the console.

## API Endpoints

The backend server exposes several API endpoints to interact with AWS and GCP functions:

### AWS Lambda

* `GET /aws/functions`: List AWS Lambda functions.
* `POST /aws/invoke/:functionName`: Invoke an AWS Lambda function.
* `POST /aws/update/:functionName`: Update an AWS Lambda function.
* `GET /aws/logs/:functionName?startTime=...&endTime=...`: Fetch AWS Lambda logs.
* `POST /aws/debugger/addBreakPoint/:functionName?fileName=...&lineNumber=...`: Add breakpoint for AWS Lambda function.

### GCP Cloud Functions

* `GET /gcp/functions`: List GCP Cloud Functions.
* `POST /gcp/invoke/:functionName`: Invoke a GCP Cloud Function.
* `POST /gcp/update/:functionName`: Update a GCP Cloud Function.
* `GET /gcp/logs/:functionName?startTime=...&endTime=...`: Fetch GCP Cloud Function logs.
* `POST /gcp/debugger/addBreakPoint/:functionName?fileName=...&lineNumber=...`: Add breakpoint for GCP Cloud Function.