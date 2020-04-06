Toast-Chest API
===============

This repository contains the source code for the **Toast Chest** API. The Toast Chest is a database of dinner/drinking toasts to get a toast on the go. An Alexa Skill is also available to bring the Toast Chest to your home (see link).

This application serves as a microservice API, connecting to Postgres. It provides a simple REST API in Go. See the API documentation for all endpoints.



Environment Setup
-----------------
1. Copy **.env-example** to **.env** and adjust variables as needed
2. Install Docker and Docker Compose
3. Build the images:
    ```
    make build
    ```

3. Start the containers:
    ```
    make start
    ```



License
-------
The material in this repository is released under a GNU General Public License v2.0.

Copyright (c) 2020.
