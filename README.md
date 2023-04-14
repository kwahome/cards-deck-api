# cards-deck-api

![img.png](img.png)

This is a simple golang application that implements a REST API to simulate a deck of cards.

## Getting Started

### 1. Clone the source

To clone the source code, use:

    git clone https://github.com/kwahome/cards-deck-api.git

or

    git clone git@github.com:kwahome/cards-deck-api.git

### 2. Run tests

To run all tests, use:

> go test ./tests/...

```text

cards-deck-api % go test ./tests/... 
?       github.com/kwahome/cards-deck-api/tests/integration     [no test files]
ok      github.com/kwahome/cards-deck-api/tests/integration/healthcheck (cached)
?       github.com/kwahome/cards-deck-api/tests/mocks   [no test files]
?       github.com/kwahome/cards-deck-api/tests/unit    [no test files]
ok      github.com/kwahome/cards-deck-api/tests/integration/v1  (cached)
ok      github.com/kwahome/cards-deck-api/tests/unit/internal/api/healthcheck   (cached)
ok      github.com/kwahome/cards-deck-api/tests/unit/internal/api/v1/handlers   (cached)
```

To run all integration tests, use:

> go test ./tests/integration/...

```text
cards-deck-api % go test ./tests/integration/...
?       github.com/kwahome/cards-deck-api/tests/integration     [no test files]
ok      github.com/kwahome/cards-deck-api/tests/integration/healthcheck (cached)
ok      github.com/kwahome/cards-deck-api/tests/integration/v1  (cached)

```

To run all unit tests, use:

> go test ./tests/unit/...

```text
cards-deck-api % go test ./tests/unit/...
?       github.com/kwahome/cards-deck-api/tests/unit    [no test files]
ok      github.com/kwahome/cards-deck-api/tests/unit/internal/api/healthcheck   (cached)
ok      github.com/kwahome/cards-deck-api/tests/unit/internal/api/v1/handlers   (cached)

```

### 3. Start web server:

Use:

> go run cmd/main.go

to start a web server running on port `8080` or the port configured via the env variable `APP_PORT`:

```text
cards-deck-api % go run cmd/main.go
INFO[0000] using the config file: %s/Users/wahome/Code/toggl/cards-deck-api/.app.yaml 
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /healthcheck              --> github.com/kwahome/cards-deck-api/internal/api/healthcheck.(*CheckStatusHandler).Ping-fm (5 handlers)
[GIN-debug] POST   /api/v1/decks             --> github.com/kwahome/cards-deck-api/internal/api/v1/handlers.(*CreateDeckHandler).CreateDeck-fm (6 handlers)
[GIN-debug] GET    /api/v1/decks/:id         --> github.com/kwahome/cards-deck-api/internal/api/v1/handlers.(*GetDeckHandler).OpenDeck-fm (6 handlers)
[GIN-debug] GET    /api/v1/decks/:id/draw    --> github.com/kwahome/cards-deck-api/internal/api/v1/handlers.(*DrawCardsHandler).DrawCards-fm (6 handlers)

```

You can also start the web server as a `docker` container:

> docker-compose up

```text
cards-deck-api % docker-compose up
[+] Running 1/0
 ⠿ Container cards-deck-api-app-1  Created                                                                                                                                                                 0.0s
Attaching to cards-deck-api-app-1
cards-deck-api-app-1  | [GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.
cards-deck-api-app-1  | 
cards-deck-api-app-1  | [GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
cards-deck-api-app-1  |  - using env:   export GIN_MODE=release
cards-deck-api-app-1  |  - using code:  gin.SetMode(gin.ReleaseMode)
cards-deck-api-app-1  | 
cards-deck-api-app-1  | [GIN-debug] GET    /healthcheck              --> github.com/kwahome/cards-deck-api/internal/api/healthcheck.(*CheckStatusHandler).Ping-fm (5 handlers)
cards-deck-api-app-1  | [GIN-debug] POST   /api/v1/decks             --> github.com/kwahome/cards-deck-api/internal/api/v1/handlers.(*CreateDeckHandler).CreateDeck-fm (6 handlers)
cards-deck-api-app-1  | [GIN-debug] GET    /api/v1/decks/:id         --> github.com/kwahome/cards-deck-api/internal/api/v1/handlers.(*GetDeckHandler).OpenDeck-fm (6 handlers)
cards-deck-api-app-1  | [GIN-debug] GET    /api/v1/decks/:id/draw    --> github.com/kwahome/cards-deck-api/internal/api/v1/handlers.(*DrawCardsHandler).DrawCards-fm (6 handlers)
cards-deck-api-app-1  | time="2023-04-13T22:09:06Z" level=info msg="using the config file: /app/.app.yaml"
cards-deck-api-app-1  | time="2023-04-13T22:09:06Z" level=info msg="HTTP server is listening on port: 8080"

```

### 4. Interacting with the APIs

- Health Check:

    ```
    curl --location --request GET 'localhost:8080/health'
  
    {
        "status": "alive",
        "timestamp": "Thursday, 13-Apr-23 23:22:58 EAT"
    }
    ```
        

- Create Deck:

    ```
    curl --location --request POST 'localhost:8080/api/v1/decks?shuffle=true&cards=AC%2CKH%2C2S' --header 'AuthToken: ab38bf18-6f87-41a7-9aed-c1eb0db64b9c'

    {
        "deck_id": "60e87195-25b2-44eb-9d67-3004ad664e5c",
        "shuffled": true,
        "remaining": 3
    }
    ```

- Open Deck:
    
    ```
    curl --location 'localhost:8080/api/v1/decks/60e87195-25b2-44eb-9d67-3004ad664e5c' --header 'AuthToken: ab38bf18-6f87-41a7-9aed-c1eb0db64b9c'

    {
        "deck_id": "60e87195-25b2-44eb-9d67-3004ad664e5c",
        "shuffled": true,
        "remaining": 3,
        "cards": [
            {
                "value": "King",
                "suite": "Heart",
                "code": "KH"
            },
            {
                "value": "Ace",
                "suite": "Club",
                "code": "AC"
            },
            {
                "value": "2",
                "suite": "Spade",
                "code": "2S"
            }
        ]
    }
    ```

- Draw Cards:

    ```
    curl --location 'localhost:8080/api/v1/decks/60e87195-25b2-44eb-9d67-3004ad664e5c/draw?count=3' --header 'AuthToken: ab38bf18-6f87-41a7-9aed-c1eb0db64b9c'

    [
        {
            "value": "King",
            "suite": "Heart",
            "code": "KH"
        },
        {
            "value": "Ace",
            "suite": "Club",
            "code": "AC"
        },
        {
            "value": "2",
            "suite": "Spade",
            "code": "2S"
        }
    ]
    ```

## Code Structure
The backend service is implemented using clean architecture principles. 
It is written in Go with complete Dependency Injection along with Mocking for testing, following SOLID principles. 
Domain Driven Design and layered architecture have been used in its design and thus the code structure follows 
elements of these architecture patterns.

The adopted architecture pattern and code organization emphasizes on separating concerns and isolating layers even
at a code level. Interaction between the layers is facilitated by interfaces and mapping/transformation of objects 
between the interacting layers. For instance, infrastructure concerns are isolated from domain logic which is 
reflected in the code organization. Their interaction is through dependency injection via composition/embedding.

A key benefit of layered architecture is that maintainability and future evolutions are seamless because modification,
refactoring, extensions are contained and transparent. For instance, we can transparently switch out technologies, 
libraries, implementations, e.t.c., in one layer without any other layers being made aware or required to make changes.

Below is a breakdown of how the code is structured:

### 1. /bin

The bin directory contains commands, scripts and any other executable files.

### 2. /cmd

This directory contains the main application entry point files for the application processes.

### 3. /config

This directory contains configuration objects and functions to load them.

### 4. /internal

This package holds the private library code used in the service. 
It is specific to the function of the service and not shared with other services.

Below are the enclosed packages:
- `api/` contains the REST API code. The nested directory structure includes versioning in line with the 
    API version which provides a clean way of isolating objects of the different versions that will typically 
    have same names. For example, Data Transfer Objects of `v1` and a future `v2` of the API will have the 
    same name but differ in schema or structure hence the two versions. We can thus have multiple version of 
    the same endpoint with an added benefit of having them side by side.

- `domain/` contains code for the business logic associated with the problem domain; i.e., deck of cards domain. 
    In typical Domain Driven microservices architecture, a single microservice encapsulates exactly one domain. 
    The code in this directory is agnostic of infrastructure and data transfer primitives as the domain logic is 
    not specific to any particular technology.

### 5. /pkg

This directory contains code which is OK for other services to consume, this may include API clients, 
or utility functions which may be handy for other projects but don’t justify their own project.

### 6. /tests

This directory contains the test files for the code in the project. This way of organizing tests was chosen over
having test files live side by side with the files they are testing because it provides an easy way to group 
various types of tests. The importance of grouping various kinds of tests is that their runtime in a build 
pipeline can be optimized. For instance, because typically unit tests mock out dependencies, they are cheap and 
quick to run tests while integration tests that must spin up such dependencies (database, message bus, etc) are 
more complex and take a little longer to run. In addition, generated mock files are contained in one place as 
opposed to having them scattered across the various packages. It makes for a cleaner structure.

It encloses the following subdirectories:

- `integration/` - contains integration tests.
- `mocks/` - contains autogenerated mocks for the various interfaces.
- `unit/` - contains unit tests. The structure of the sub-directories mirrors that of the code being tested.
