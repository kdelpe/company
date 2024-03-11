# Project Name

The Company Management System is a web application designed to streamline employee and branch management for a company. It allows users to manage employee data, including their personal information, salaries, and branch assignments. Additionally, it provides features to manage branch information such as branch managers, start dates, and suppliers associated with each branch. The application is built using Go programming language and utilizes the Gin web framework for the backend. Data is stored in a MySQL database, and the application provides RESTful APIs for interacting with the data.

## Key Features:

- **Employee management**: Add, update, and retrieve employee data including personal details and salaries.
- **Branch management**: Manage branch information such as branch name, managers, and associated suppliers.
- **Branch-supplier association**: Track suppliers associated with each branch and the type of supplies they provide.
- **Robust data handling**: Handles nullable values effectively, ensuring accurate retrieval and representation of data.
This project aims to provide a comprehensive solution for managing company resources efficiently. Contributions and feedback are welcome!

## Table of Contents
- [Installation](#installation)
- [Usage](#usage-)
- [Tests](#test-)
- [Author](#author)

## Installation

1. Clone the repository
```bash
$ git clone https://github.com/kdelpe/company.git
$ cd project-folder
$ go build # or any other package manager command
```

2. Create `.env` file and update the following variables:
```text
DB_USERNAME="your_DB_username"
DB_PASSWORD="your_DB_password"
```

## Usage 

```bash
$ go run .
```

## Test 

## Author
[Kervens Delpe](https://github.com/kdelpe)