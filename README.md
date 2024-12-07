# LogiLog: Logistics Delivery App

LogiLog is a microservices-based logistics delivery application that aims to provide comprehensive solutions for e-commerce businesses.

## Table of Contents

1. [Project Overview](#project-overview)
2. [Features](#features)
3. [How it Works](#how-it-works)
4. [Requirements](#requirements)
5. [Project Structure](#project-structure)
6. [Setup Instructions](#setup-instructions)
   - [Development Environment](#development-environment)
7. [Environment Variables](#environment-variables)
8. [Running the Application](#running-the-application)
9. [GraphQL Interaction](#graphql-interaction)
10. [Database Management](#database-management)
11. [Troubleshooting](#troubleshooting)
12. [Contributing](#contributing)

## Project Overview

LogiLog is a comprehensive logistics delivery platform that helps e-commerce businesses streamline their operational and logistical processes. It is built using a microservices architecture and includes the following components:

1. **Account Service**: Handles user authentication and account management.
2. **Shopify Service**: Integrates with the Shopify API for order management.
3. **Shipment Service**: Manages shipment processing and courier integrations.
4. **Payment Service**: Handles payment processing and wallet management.
5. **GraphQL Gateway**: Acts as a single access point for querying and managing data from the microservices.

Each service is containerized and uses PostgreSQL as its database.

## Features

LogiLog offers the following key features:

1. **Courier Recommendations**: Intelligent courier recommendations based on factors such as cost, delivery times, and customer preferences.

2. **Order Tracking**: Comprehensive order tracking facilities to keep customers informed about the status of their deliveries.
3. **Dedicated Support**: Provides dedicated support to help customers with their logistics needs.
4. **Buyer Intent**: Analyzes buyer behavior and patterns to optimize logistics and delivery processes.
5. **Pay-as-you-go Model**: No subscription or integration fees, customers only pay for the services they use.

## How it Works

LogiLog consolidates multiple shipping companies and automates various shipping processes to provide a seamless logistics solution for e-commerce businesses. It integrates with popular 3PL (third-party logistics) providers such as Delhivery, XpressBees, FedEx, Ecomm, EKart, and Bluedart to offer a wide range of delivery options.

The platform also provides a knowledge base with information, guides, and solutions to help customers better understand and utilize the platform.

## Requirements

Make sure the following tools are installed:

- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)
- [PostgreSQL](https://www.postgresql.org/)
- [Go](https://golang.org/) (for backend services and GraphQL gateway)
- [Git](https://git-scm.com/)

## Project Structure

```plaintext
root/
├── account/       # User management service
├── shopify/       # Shopify API integration service
├── shipment/      # Shipment processing service
├── payment/       # Payment and wallet service
├── graphql/       # GraphQL gateway for microservices
```

## Setup Instructions

### Development Environment

1. **Clone the repository**:
    ```bash
    git clone <repository-url>
    cd <repository-folder>
    ```

2. **Environment Variables**:
    Create `.env.development` files in the root directories of each microservice (`account`, `shopify`, `shipment`, `payment`, `graphql`). Refer to the [Environment Variables](#environment-variables) section for details.

3. **Start PostgreSQL Containers**:
    ```bash
    docker-compose -f docker-compose.dev.yml up -d
    ```

4. **Run Each Service**:
    - For Go-based services:
      ```bash
      cd <service-folder>
      go run main.go
      ```
    - For the GraphQL gateway:
      ```bash
      cd graphql
      npm install
      npm run dev
      ```

## Environment Variables

Each service requires specific environment variables. Refer to the [README.md](README.md) file in the respective service directories for details.

## Running the Application

1. **Start PostgreSQL Containers**:
    ```bash
    docker-compose -f docker-compose.dev.yml up -d
    ```

2. **Run Each Service**:
    Follow the steps in [Setup Instructions](#setup-instructions).

3. **Access the Services**:
    - Account: `http://localhost:3001`
    - Shopify: `http://localhost:3002`
    - Shipment: `http://localhost:3003`
    - Payment: `http://localhost:3004`
    - GraphQL Gateway: `http://localhost:4000`

## GraphQL Interaction

The GraphQL gateway serves as a single access point for interacting with all microservices. Refer to the [README.md](README.md) file for examples of GraphQL queries and mutations.

## Database Management

- Development: `postgres://dev_user:dev_password@localhost:5432/{service_db}`

Use tools like [pgAdmin](https://www.pgadmin.org/) or `psql` for database operations.

## Troubleshooting

Refer to the [README.md](README.md) file for troubleshooting steps related to database issues, service connectivity, and GraphQL gateway errors.

## Contributing

We welcome contributions! Please follow the steps outlined in the [README.md](README.md) file.