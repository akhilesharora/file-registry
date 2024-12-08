# File Registry Service

This project allows you to upload files to IPFS and store their corresponding CIDs on an Ethereum-based smart contract. It includes:

- A **Solidity smart contract** for on-chain storage of file-to-CID mappings.
- A **Golang API service** for handling file uploads, interacting with IPFS, and storing/retrieving data on-chain.
- A **Docker Compose environment** that sets up everything, including a Hardhat Ethereum test node and an IPFS node.

## Overview

**Workflow:**
1. **Upload a file** via the `POST /v1/files` endpoint.
2. The file is **added to IPFS**, which returns a CID.
3. The returned CID and the given file path are **stored on the Ethereum blockchain** using a deployed `FileRegistry` smart contract.
4. Later, you can **retrieve the CID** for a given file path via `GET /v1/files?filePath=<your_file_path>`.

**Key Components:**
- **Smart Contract** (`FileRegistry.sol`):
    - `save(string filePath, string cid)`: Store the CID for the given file path on-chain.
    - `get(string filePath)`: Retrieve the CID associated with a given file path.
- **Golang API Service**:
    - Endpoints:
        - `POST /v1/files`: Uploads the file to IPFS and saves CID on-chain.
        - `GET /v1/files?filePath=<path>`: Retrieves the CID from the contract.
- **IPFS Node**: Runs as a Docker container, enabling file uploads and CID generation.
- **Hardhat Node**: Local Ethereum test environment for deploying and interacting with the contract.

## Architectural Design: Hexagonal Architecture (Ports and Adapters)

This project implements a Hexagonal Architecture (Ports and Adapters) to achieve high modularity, testability, and decoupling of components. This pattern allows us to:

- **Isolate Core Business Logic:** The domain logic remains independent of external concerns like frameworks, or blockchain interactions.
- **Minimize Dependencies:** External systems can be swapped without modifying core application logic.
- **Improve Testability:** Core logic can be tested in isolation, without dependencies on external system.

### Architectural Components

- **Ports:** Interfaces that define contracts for interactions
    - Inbound Ports: Define how external systems can interact with core logic
    - Outbound Ports: Define how core logic can interact with external systems

- **Adapters:** Implement ports and translate between external systems and core logic
    - Primary (Driving) Adapters: Handle incoming requests (e.g., HTTP handlers)
    - Secondary (Driven) Adapters: Handle external system interactions (e.g., IPFS, Ethereum)

## Project Structure

```
.
├─ config/                    # Configuration for Go service
├─ cmd/
│  └─ file-registry           # Main Golang service entrypoint
├─ contracts/
│  ├─ fileregistry/           # Generated Go bindings for FileRegistry contract
│  └─ solidity/               # Solidity source & Hardhat environment
│      ├─ hardhat/
│      │   ├─ contracts/FileRegistry.sol
│      │   ├─ scripts/deploy.js
│      │   ├─ Dockerfile.hardhat
│      │   └─ ... other Hardhat files
├─ internal/
│  ├─ ports/                  # Interfaces defining contracts
│  │   ├─ storage/            # Storage port interfaces
│  │   └─ service/            # Service port interfaces
│  ├─ adapters/               # Implementations of ports
│  │   ├─ transport/          # HTTP/API adapters (Primary adapters)
│  │   └─ storage/            # IPFS & Ethereum adapters (Secondary adapters)
│  ├─ core/                   # Core domain logic & services
│  │   ├─ domain/             # Domain models
│  │   └─ services/           # Business logic implementations
│  └─ handlers/               # HTTP request handlers
├─ static/                    # Static files 
├─ Dockerfile                 # Dockerfile for the Go API service
├─ docker-compose.yml         # Docker setup for API, IPFS, and Hardhat
└─ Makefile                   # Build, test, and run commands
```

## Prerequisites

- **Docker & Docker Compose**
- **make** (optional but recommended)
- **Go 1.22+** (if you plan to run locally outside Docker)
- **Node.js & npm** (if deploying/contracts without Docker)

## Setup & Running

### Quick Start with Docker

1. **Environment:**  
   Copy `.env.template` to `.env` and adjust as needed:
   ```bash
   cp .env.template .env
   ```
   If `CONTRACT_ADDRESS` is empty, it will be automatically set after deployment via `make docker-up`.

2. **Run the stack:**
   ```bash
   make docker-up
   ```
   This will:
    - Start a Hardhat node and deploy the `FileRegistry` contract automatically.
    - Start an IPFS node.
    - Start the Golang API service with all required environment variables set.

3. **Check services:**
    - API: [http://localhost:8090](http://localhost:8090)
    - IPFS: [http://localhost:5001/webui](http://localhost:5001/webui)
    - Ethereum (Hardhat): http://localhost:8545

### Manual Setup (Without Docker)

1. **Run Hardhat Node:**
   In `contracts/solidity/hardhat`:
   ```bash
   npx hardhat node
   ```

2. **Deploy the Contract:**
   In a separate terminal:
   ```bash
   cd contracts/solidity/hardhat
   npx hardhat compile
   npx hardhat run scripts/deploy.js --network localhost
   ```
   This will print out the deployed contract address. Update your `.env` file with this `CONTRACT_ADDRESS`.

3. **Run IPFS Node:**
   Follow IPFS setup instructions (e.g., `ipfs daemon`), or run a local IPFS node separately.

4. **Run the Go API:**
   With `CONTRACT_ADDRESS` and `IPFS_NODE_URL` set in `.env`, run:
   ```bash
   make build
   make run
   ```

   The server should now be accessible at [http://127.0.0.1:8090](http://127.0.0.1:8090).

## Usage

- **Upload a file:**
  ```bash
  curl -F "file=@/path/to/yourfile.txt" \
       -F "filePath=/myfile.txt" \
       http://localhost:8090/v1/files
  ```
  This returns a JSON response with the CID and path.

- **Retrieve a CID:**
  ```bash
  curl "http://localhost:8090/v1/files?filePath=/myfile.txt"
  ```
  This returns the CID associated with `/myfile.txt`.

## Testing

- **Go Tests:**
  ```bash
  make test
  ```

## Cleanup

- Stop Docker containers:
  ```bash
  make docker-down
  ```

- Remove build artifacts:
  ```bash
  make clean
  ```

## Notes

- `.env` and `docker-compose.yml` handle the environment setup.
- `Makefile` offers build/test/run shortcuts.

## License

This project is licensed under the [MIT License](LICENSE).
