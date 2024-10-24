# hashring [![Go](https://github.com/marcsantiago/go-hashring/actions/workflows/go.yml/badge.svg)](https://github.com/marcsantiago/go-hashring/actions/workflows/go.yml)

A simple consistent hashing ring implementation in Go.

A hash ring, also known as consistent hashing, is a technique used in distributed systems to distribute data across a cluster of nodes in a way that minimizes reorganization when nodes are added or removed. It is particularly useful for load balancing and distributed caching.

### Key Concepts:
1. **Hash Function**: Maps data to a position on the ring.
2. **Ring Structure**: The hash space is treated as a circular ring.
3. **Nodes**: Each node in the cluster is assigned a position on the ring.
4. **Data Distribution**: Data is assigned to nodes based on their position on the ring.

### How It Works:
1. **Hashing Nodes**: Each node is hashed to a position on the ring.
2. **Hashing Data**: Each piece of data is hashed to a position on the ring.
3. **Data Assignment**: Data is assigned to the first node that is encountered when moving clockwise around the ring from the data's position.

### Benefits:
- **Scalability**: Easily add or remove nodes with minimal data movement.
- **Load Balancing**: Evenly distributes data across nodes.
- **Fault Tolerance**: If a node fails, its data can be redistributed to other nodes.

### Example:
Consider a ring with nodes A, B, and C. Data items are hashed and placed on the ring. Each data item is stored in the first node encountered moving clockwise from its position.

This technique ensures that when a node is added or removed, only a small portion of the data needs to be reassigned, making the system more efficient and resilient. items are hashed and placed on the ring. Each data item is stored in the first node encountered moving clockwise from its position.  This technique ensures that when a node is added or removed, only a small portion of the data needs to be reassigned, making the system more efficient and resilient.