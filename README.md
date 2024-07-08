# WASMATE: Unleash Innovation, Secure the Future
![build](https://github.com/wasmate/wasmate-runtime/actions/workflows/build.yml/badge.svg)
![test](https://github.com/wasmate/wasmate-runtime/actions/workflows/test.yml/badge.svg)

"WASMATE" is an innovative project revolutionizing application runtime environments by bridging Web2.0 and Web3.0. It offers advanced WASM runtimes, enabling easy access to cloud-native environments and microservice middleware with low resource utilization and simplified distributed systems. For Web3.0, WASMATE integrates decentralized storage, trusted computing, AI model interaction, decentralized identity verification, and cross-chain computing, making it a leader in the blockchain ecosystem. Join WASMATE to explore the future of secure, efficient, and innovative web applications, driving technological innovation and shaping the web's next era.

<p align="center">
<img 
    src="https://github.com/wasmate/wasmate-runtime/assets/34047788/06697deb-1523-49ce-a8c3-d851adfe646c" 
     alt="wasmate">
</p>


# Runtime capabilities

<p align="center">
<img 
    src="https://github.com/wasmate/wasmate-runtime/assets/34047788/c8a5e96d-6806-4ebc-bfa2-e349b26c347c" 
     alt="wasmate">
</p>

## Web2.0 Domain

| Capability                           | Description                                                                            |
|---------------------------------------|----------------------------------------------------------------------------------------|
| WASM Runtime Environment              | Provides a secure default WASM runtime environment, facilitating easy deployment of applications. |
| Resource Efficiency                   | Stands out with low resource utilization, offering an efficient application runtime environment. |
| Function-level Service Invocation     | Allows communication between services through function-level calls, promoting loosely coupled microservices architecture. |
| State Management                      | Supports effective tracking and handling of application state.                           |
| Publish-Subscribe                     | Implements a publish-subscribe pattern for event-driven communication and collaboration. |
| Trigger Management                    | Enables users to manage and configure triggers for corresponding actions based on events. |
| Actor Concurrency Management          | Supports the Actor model, making concurrent programming easier with advanced concurrency control. |
| Confidential Storage and Management    | Provides features for secure storage and management of sensitive information.            |
| Remote Configuration                   | Allows remote configuration, enabling dynamic adjustments to application settings without redeployment. |
| Distributed Networking                 | Supports distributed networking, facilitating collaborative work and communication in multi-node systems. |
| Distributed Messaging Network          | Offers a distributed messaging network for reliable communication in distributed applications. |
| Distributed Lock                       | Supports distributed locks to ensure synchronization and coordination of concurrent access across multiple nodes. |
| Simplified Distributed System Complexity | Design philosophy aimed at simplifying the complexity of distributed systems, allowing developers to focus on business logic rather than low-level details. |
| Suitable for Cloud-native Environments and Microservices Middleware | Provides a solution for users to enter cloud-native environments or use microservices middleware without extensive application refactoring. |

## Web3.0 Domain

| Capability                           | Description                                                                            |
|---------------------------------------|----------------------------------------------------------------------------------------|
| Decentralized Storage                 | Supports storage in a decentralized environment, ensuring data decentralization and security. |
| Trusted Computing                     | Provides support for trusted computing to ensure the security and reliability of computational processes. |
| AI Big Model Interfaces               | Supports interfaces with large AI models, offering powerful artificial intelligence capabilities. |
| Decentralized Identity                | Provides a decentralized identity solution, ensuring the security and privacy of user identities. |
| Blockchain Contracts                  | Capable of interacting with blockchain contracts, supporting the execution and management of smart contracts. |
| Cross-chain Computing                 | Supports cross-chain computing, enabling interoperability and collaboration of computations across different blockchain networks. |
| Blockchain Data Interaction           | Provides interaction with blockchain networks for reading and writing data on the blockchain. |
| Decentralized App Integration         | Allows integration and execution of decentralized applications, supporting the ecosystem of applications in the Web3.0 environment. |
| Decentralized Finance (DeFi) Support  | Offers functionality to support decentralized finance (DeFi) applications, including smart contracts and digital asset management. |
| Decentralized Identity Verification   | Supports decentralized identity verification, ensuring secure and trusted user identity in the Web3.0 environment. |
| Cross-chain Asset Transfer            | Supports asset transfer between different blockchain networks, facilitating interconnectivity of assets. |
| Blockchain Governance Support         | Provides functionality to support blockchain governance, including decentralized mechanisms for voting, proposing, and decision-making. |
| Decentralized File Storage            | Supports decentralized file storage solutions, ensuring distributed storage and sharing of files. |


# Quick start

## 1. Clone the repository
   ```bash
   git clone https://github.com/wasmate/wasmate-runtime.git
   ```

## 2. Build
   ```bash
   cd wasmate-runtime
   make
   ```

## 3. Adjust configuration file
   ```yaml
app-type: "wasmate-runtime-WORKER"

# The network model of HTTP handle ,NetPoll(gin) RAWEPOLL(fiber)
net-model: "NETPOLL"

# Process inflow traffic network configuration
NetWork:
  bind-network: "TCP" #Network transport layer type: TCP | UDP 
  protocol-type: "HTTP" #Application layer network protocol：HTTP | RESP | QUIC
  bind-address: "127.0.0.1:28080" #Network listening address

#Runtime debug option
debug:
  enable: false
  pprof-bind-addr: "127.0.0.1:19090"

wasm-modules-files:
  enable: false
  path:
    - "hello.wasm"

wasm-modules-ipfs:
  enable: false
  lassie-net:
    scheme: "http"
    host: "x.x.x.x"
    port: xxxx
  cids:
    - "QmeDsaLTc8dAfPrQ5duC4j5KqPdGbcinEo5htDqSgU8u8Z"

wasm-modules-starknet:
  enable: true
  rpc-address: "https://starknet-sepolia.public.blastapi.io"
  smart-contract: "0x01016993aa219f246d39ec6c25e1eef4920fe1e650179957bff9c0a08e09ed89"
  contract-method: "get_wasm_cid"
  lassie-net:
    scheme: "http"
    host: "x.x.x.x"
    port: xxxx
  wasm-func-names:
    - "sayhello"
   ```

## 4. Load configuration and run
 ```shell
   wasmate-runtime -c wmr_worker.yaml
 ```

## 5. Testing WASMATE Runtime

```shell
$ curl -d "wasmate-runtime" "http://localhost:28080"
👋 Hello wasmate-runtime%
```

# Performance Testing

```shell
$ hey -n 1000000 -c 100 -m POST \
-d 'wasmate-runtime' \
"http://127.0.0.1:28080"

Summary:
  Total:        29.0599 secs
  Slowest:      0.0522 secs
  Fastest:      0.0001 secs
  Average:      0.0029 secs
  Requests/sec: 34411.6948
  
  Total data:   26000000 bytes
  Size/request: 26 bytes

Response time histogram:
  0.000 [1]     |
  0.005 [985746]        |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.010 [13559] |■
  0.016 [132]   |
  0.021 [108]   |
  0.026 [147]   |
  0.031 [91]    |
  0.037 [124]   |
  0.042 [19]    |
  0.047 [1]     |
  0.052 [72]    |


Latency distribution:
  10% in 0.0023 secs
  25% in 0.0026 secs
  50% in 0.0029 secs
  75% in 0.0033 secs
  90% in 0.0036 secs
  95% in 0.0039 secs
  99% in 0.0057 secs

Details (average, fastest, slowest):
  DNS+dialup:   0.0000 secs, 0.0001 secs, 0.0522 secs
  DNS-lookup:   0.0000 secs, 0.0000 secs, 0.0000 secs
  req write:    0.0000 secs, 0.0000 secs, 0.0318 secs
  resp wait:    0.0028 secs, 0.0001 secs, 0.0521 secs
  resp read:    0.0000 secs, 0.0000 secs, 0.0260 secs

Status code distribution:
  [200] 1000000 responses

```

# Join the Revolution

**WASMATE** is more than just a technological platform; it's a visionary project leading towards future innovation. Join us in shaping a secure, efficient, and innovative web application ecosystem. Feel free to explore the project further on our [official website](https://www.wasmate.xyz/).

Let's revolutionize the way we build and deploy web applications with **WASMATE**!

We welcome contributions from the community. To contribute to this project:

1. Fork the repository.
2. Create a new branch (`git checkout -b feature/your-feature`).
3. Make your changes and commit them (`git commit -am 'Add new feature'`).
4. Push your changes to the branch (`git push origin feature/your-feature`).
5. Create a new Pull Request.

# Acknowledgements
* **IPFS and Filecoin** - For decentralized storage and retrieval.
* **StarkNet** - For providing the zk-rollup-based Layer 2 solution.
* **Extism** - For wasm plugin management.
* **wazero** - For WebAssembly virtual machine capabilities.
* **Fiber** - For HTTP server functionalities.

# License

This project is dual-licensed under Apache 2.0 and MIT terms.

## Disclaimers

When you use this software, you have agreed and stated that the author, maintainer and contributor of this software are not responsible for any risks, costs or problems you encounter. If you find a software defect or BUG, please submit a patch to help improve it!
