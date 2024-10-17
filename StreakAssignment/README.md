# Find Pairs API

## Overview

The Find Pairs API is a simple backend service that takes an array of integers and a target value, returning all pairs of indices where the sum of the elements at those indices equals the target value. This API is built using Go and follows good coding practices to ensure maintainability, performance, and clarity.

## Features

- Accepts a JSON request containing an array of integers and a target sum.
- Returns all pairs of indices whose corresponding elements sum to the target.
- Handles edge cases, including:
  - Empty arrays.
  - Arrays with duplicate numbers.
  - Cases where no pairs match the target.
- Validates input to ensure correctness.

## API Endpoint

### `POST /find-pairs`

#### Request Body

The request should be a JSON object with the following structure:

```json
{
  "numbers": [1, 2, 3, 4, 5],
  "target": 6
}


Steps to Run
Clone the repository:
git clone <repository-url>

cd find-pairs-api
Install dependencies (if any):
go mod tidy

Run the server:
go run main.go

The server will start, and you can access the API at http://localhost:8080/find-pairs.