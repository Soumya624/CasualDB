# CasualDB 🗄️

A simple, lightweight database implementation built in Go that demonstrates fundamental database storage concepts through an intuitive block-and-page architecture.

## 🎯 Core Concept

CasualDB breaks files into **multiple blocks**, each with individual identifiers, and uses **pages** as an intermediary layer for efficient data handling.

### How It Works

**Writing Process:**
1. Data is first written to a **page** (in-memory buffer)
2. Page content is then copied to the target **block** in the file
3. Each block maintains its unique identity for easy retrieval

**Reading Process:**
1. Initialize an empty **page**
2. Copy content from the specified **block** to the page
3. User reads data directly from the page

This design provides a clean abstraction between memory operations (pages) and disk storage (blocks).

## 🏗️ Current Architecture

- **Blocks**: Fixed-size storage units in files, each with unique IDs
- **Pages**: In-memory buffers for data manipulation
- **File Controller**: Manages file I/O operations and block access
- **Page Controller**: Handles in-memory data operations
- **Web Interface**: Simple form-based interaction for testing

## 🚀 Getting Started

```bash
go run main.go
```

Visit `http://localhost:8080` to interact with the database through the web interface.

---

## 🧩 Future Vision: Causal Database

**Next Evolution**: Transform CasualDB into a database that stores not just *what happened*, but also *why it happened*.

### 💡 The Idea
Each record won't just be state, it will include **causal metadata** that tracks:
- **Why** the change was made
- **Who** made the change  
- **What** triggered the change
- **Dependencies** and relationships

### 🛠️ Planned Enhancements

**Enhanced Data Models**
- Add causal fields to existing Page/Block structures

**Extended Controllers** 
- Causal-aware read/write operations
- Metadata handling alongside data operations

**Causal Indexing**
- Index causal metadata within blocks

**Metadata Serialization**
- Efficient storage of causal information
- JSON-based metadata alongside binary data
