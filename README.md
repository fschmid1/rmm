
# FES_RMM
## Description
A remote monitoring and management software written in Go.

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
    This will create a binary at: "bazel-bin/projects/backend/RMM_API_LINUX_/RMM_API_LINUX"
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
    This will create a binary at: "bazel-bin/projects/client/RMM_CLIENT_LINUX_/RMM_CLIENT_LINUX"
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
## Deploy Frontend
1. Change api base to backend url
    ``` ts
     // projects/frontend/src/vars.prod.ts
    export const apiBase = '/';
    export const wsBase = 'ws://';
    ```
2. Install dependencies
    ```
    npm i (in root folder of this repo)
    ```
3. Build frontend
    ```
    make build.backend (in root folder of this repo)
    ```
    This will create the static files under: "projects/frontend/dist"
