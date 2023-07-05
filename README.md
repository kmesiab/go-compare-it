# Go Compare It
![Golang](https://img.shields.io/badge/Go-00add8.svg?labelColor=171e21&style=for-the-badge&logo=go)
![Postgres](https://img.shields.io/badge/postgres-%23316192.svg?style=for-the-badge&logo=postgresql&logoColor=white)
![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white)
![Terraform](https://img.shields.io/badge/terraform-%235835CC.svg?style=for-the-badge&logo=terraform&logoColor=white)


Go Compare It is an AI-based comparison engine that allows users to build lists of products that are analyzed for a variety of "value weights," resulting in an opinionated product review.

## 📐Current Release

| v.01 | Alpha |
| --- | --- |
| v.1.0 | TBD |

## 🎁 Features

- 🤖 AI-powered product comparison
- 📚 GPT-4 Large Language Model integration
- 📊 Proprietary product analysis graph
- 🐳 Docker Compose environment for easy installation

## 🔨 Installation

To get started with Go Compare It, follow these steps:

1. ⚙️ Start the Docker Compose environment:

    ```
    docker-compose -f docker-compose.yml up
    ```

2. ✅ Verify that the application is running by hitting the health check endpoint:

    ```
    <http://localhost:8080/healthcheck>
    ```
3. 🌐 The base URL is `/api/v1/`


## This application is…

…written in [Go](https://go.dev/), 
 backed by [PostgreSQL](https://www.postgresql.org/), 
 powered by [dGraph](https://github.com/dgraph-io/dgraph),
 made sentient by [OpenLLaMA](https://github.com/yxuansu/OpenAlpaca)...

## Graph Database
Go Compare It uses a graph database to maintain high dimensional relationships 
between products, prices, customer sentiment, reliability analysis, and a myriad 
of other informative bits of data. For an introduction to graph databases, check out 
Developing a Small-Scale Graph Database: A Ten Step Learning Guide for Beginners.

## PostgreSQL
Go Compare It uses a relational database to store and manage the ticky tacky 
administrative minutiae of a web platform.


# Author
Kevin Mesiab <kmesiab@gmail.com>
