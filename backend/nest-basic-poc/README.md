# NestJS Basic POC

This project is a basic proof-of-concept application built with NestJS.

## Prerequisites

* Node.js (version >= 16)
* npm or yarn

## Installation

1.  Clone the repository.
2.  Navigate to the `backend/nest-basic-poc` directory.
3.  Run `npm install` or `yarn install` to install dependencies.

## Running the Application

1.  Create a `.env` file in the root of the `nest-basic-poc` directory with the necessary environment variables.
2.  Run `npm run start:dev` or `yarn start:dev` to start the development server.

## Project Structure

* `src/`: Contains the application source code.
    * `modules/`: Feature-specific modules (auth, users, ping, etc.).
    * `common/`: Shared utilities, DTOs, interfaces, etc.
    * `config/`: Configuration files.
    * `main.ts`: Application entry point.
* `test/`: End-to-end tests.
* `Dockerfile`: Docker configuration.