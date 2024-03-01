# Go CRUD API

A simple Go URL Shortner API using [Gin](https://github.com/gin-gonic/gin) framework for managing resources.

## Table of Contents

- [Introduction](#introduction)
- [Prerequisites](#prerequisites)
- [Getting Started](#getting-started)
  - [Installation](#installation)
  - [Configuration](#configuration)
- [Usage](#usage)
  - [Endpoints](#endpoints)
  - [Examples](#examples)
- [Contributing](#contributing)
- [License](#license)

## Introduction

This API provides a basic URL Shortner interface for managing URLs. It is built using the Go programming language and the Gin web framework.

## Prerequisites

Before you begin, ensure you have met the following requirements:

- Go installed (version 1.21)
- Git installed
- [Gin](https://github.com/gin-gonic/gin) installed (used for routing)
- or Simply enter `go mod tidy` in console, to install all dependencies.

## Getting Started

### Installation

1. Clone the repository:

   ```shell
   git clone https://github.com/sameergaikwad222/urlshortner.git
   cd urlshortner
   ```

2. Install All Dependencies

   ```shell
   go mod tidy
   ```

3. Set Configuration File. (path can be found in .gitignore file)
   Set Config File name as given in .gitignore and set all values accordingly

4. Build & Run the API

   ```shell
   go run main.go
   ```

## Contributing

Contributions are welcome! Please fork the repository and create a pull request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.
