# holmcoin
Blockchain Implementation in Go

This repository contains my implementation of a blockchain by modifying material from a Udemy course on blockchain development

## Overview

The project includes the implementation of basic blockchain concepts such as:

- Creating blocks
- Validating the blockchain
- Implementing Proof of Work
- Creating transactions
- Generating wallets

## Setup and Running

This project uses a Dev Container for a consistent development environment. To set up and run the project:

1. Install [Docker](https://www.docker.com/products/docker-desktop) on your machine.

2. Install the [Remote - Containers](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers) extension in Visual Studio Code.

3. Clone this repository:

    ```bash
    git clone https://github.com/holmen1/holmcoin.git
    ```

4. Open the project in Visual Studio Code, then choose the "Reopen in Container" option. This will start the Dev Container and open the project inside it.

5. Once the Dev Container is running, you can run the main file:

    ```bash
    go run main.go
    ```


## Dev Containers

The configuration for the Dev Container is defined in the `devcontainer.json` file.


## Disclaimer
This is a learning project and should not be used for real-world transactions or production use-cases.

## Acknowledgements
This project is based on a Udemy course on blockchain  https://www.udemy.com/share/105QXQ/ I would like to thank the course instructors for their valuable insights and guidance.