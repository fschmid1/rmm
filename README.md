
# FES_RMM
## Description
A remote monitoring and management software written in Go. It provides a set of API endpoints to remotely monitor and manage the system.

## Installation
1. Clone the repository to your local machine
2. Install Bazel
    ```
    brew install bazel 
    ```
## Deploy Backend
1. Build Backend
    ```
    make build.backend.linux
    ```
    This will create a binary at: "./bazel-bin/projects/backend/RMM_API_LINUX_/RMM_API_LINUX"
2. Transfer the binary to a server
3. Set the environment variables
    ```
    export DATABASE_URI="<GORM ORM format>"
    export PUSHOVER_APITOKEN="<Pushover API token>"
    export JWT_SECRET="<JWT secret>"
    export SOCKET_JWT_SECRET="<Socket JWT secret>"
    ```
4. Run binary

## Deploy Client
1. Build Client
    ```
    make build.client.linux
    ```
    This will create a binary at: "./bazel-bin/projects/client/RMM_CLIENT_LINUX_/RMM_CLIENT_LINUX"
2. Transfer the binary to a server
3. Run binary as root
4. Create a Device Token in the UI
5. Setup config
    ```
    {
        "secure": false,
        "token": "<Token you just created>",
        "host": "<ip or url of server>",
        "path": "<Path which the api runs on>",
        "port": "<Port of the api>",
        "allow": {
            "run": true,
            "shutdown": true,
            "reboot": true,
            "ProcessList": true,
            "ServiceList": true,
            "ServiceLogs": true,
            "ServiceStop": true,
            "ServiceStart": true,
            "ServiceStatus": true,
            "ServiceRestart": true,
            "Kill": true
        }
    }
    ```
6. Run client again. This should register the device.
7. Restart the client one more time. Now the Device should be visible in the UI.
