![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white) ![AmazonDynamoDB](https://img.shields.io/badge/Amazon%20DynamoDB-4053D6?style=for-the-badge&logo=Amazon%20DynamoDB&logoColor=white) ![OpenTelemetry](https://img.shields.io/badge/OpenTelemetry-FFFFFF?&style=for-the-badge&logo=opentelemetry&logoColor=black) ![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white) ![GitHub Actions](https://img.shields.io/badge/github%20actions-%232671E5.svg?style=for-the-badge&logo=githubactions&logoColor=white)
# Go URL Shortener 

A scalable URL shortener API built from scratch using **Golang**. 

The goal of this project is to explore modern backend engineering concepts without relying on "magic libraries" for the core logic. It focuses on performance, clean architecture, and observability.

## 🚀 Tech Stack

* **Language:** Golang (`1.25`)
* **Database:** Amazon DynamoDB
* **Containerization:** Docker & Docker Compose
* **Observability:** OpenTelemetry (Traces & Metrics)
* **CI/CD:** GitHub Actions

## 🧠 Core Logic & Architecture

This project explicitly implements **two different algorithms** to demonstrate engineering trade-offs between performance and mathematical theory:

### 1. The Performance Approach (Direct Base62)
* **Logic:** Generates a random Base62 string directly from the alphabet characters.
* **Storage:** The generated string is stored as the Primary Key (PK).
* **Why:** This is the "production-ready" approach. It skips the computational overhead of mathematical conversion (`int` -> `string`) and is optimized for high-throughput scenarios.

### 2. The Educational Approach (Numeric Conversion)
* **Logic:** Generates a cryptographically secure random Integer (uint64), converts it to Base62 using a custom mathematical algorithm (Modulo/Division), and then persists data.
* **Storage:** The original numeric ID is stored in the database.
* **Why:** To master the underlying algorithms of data representation and base conversion manually, implementing the logic from scratch without external libraries.
