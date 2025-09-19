# Network Event Data Generator

This repository contains a high-volume data generator for network events, designed as an interview assessment project.

## Challenge Overview

The current implementation generates correct event data but **is intentionally unoptimized** and contains performance
bottlenecks and potential security vulnerabilities.

**Your mission**: Refactor and optimize this codebase to achieve production-grade performance, security, and
maintainability while handling millions of events efficiently.

## Prerequisites

- **Go 1.24+** (check with `go version`)
- **Make** (for running build commands)
- Basic understanding of high-performance data processing
- graphviz (optional for generating profiler reports)

## Getting Started

1. **Clone and explore** the repository structure
2. **Run the current implementation**: `make test_1M`
3. **Analyze performance bottlenecks** and identify optimization opportunities
4. **Implement your improvements** following the guidelines below

## Available Resources

You have access to all standard development tools and resources:

- Any IDE (GoLand, VS Code, Cursor, etc.)
- Search engines and documentation
- AI assistants (GitHub Copilot, ChatGPT, etc.)
- Feel free to ask the interviewers questions

## Code Modification Guidelines

You are **free to modify any code** within this repository to improve performance, fix vulnerabilities, and enhance
maintainability. The following constraints apply:

* **Event Data Integrity**: The generated events must preserve their semantic meaning and all required data fields
* **Data Structure Flexibility**: You may propose changes to the data structure format (e.g., JSON, CSV, binary, etc.)
  as long as no data is lost and the events remain loadable into a database
* **Field Requirements**: All event fields listed in the Event Structure section must be preserved, though you may
  optimize their representation or storage format

The key principle is: **optimize the implementation while preserving the data's meaning and completeness**.

# Description

## Task Background

The goal of this project is to develop a data generator for **network events**. These events simulate those captured by
a mediation system, which processes hundreds of thousands of events per second. The mediation system's sole purpose is
to persist these events on disk for further processing.

### Objective:

Design a mechanism to generate millions of events and persist them on disk, mimicking the behavior of the mediation
system for internal testing purposes.

---

## System Description

### Data Generator

The **data generator** is responsible for generating a specified number of **random events**. Each event's content is
randomly generated, and while uniqueness is not enforced, constants or identical values across events should be avoided.

### Event Type Distribution

The `event_type` field is **critical** for downstream rating systems, as different types trigger operations with varying
computational costs:

| Event Type | Probability | Processing Cost | Business Impact     |
|------------|-------------|-----------------|---------------------|
| **1**      | 15%         | Low             | Standard calls      |
| **2**      | 20%         | Medium          | Premium services    |
| **3**      | 20%         | Medium          | International calls |
| **5**      | 45%         | High            | Complex routing     |

**Important**: This distribution must be maintained precisely, as it reflects real-world traffic patterns and affects
downstream system resource planning.

The generated events are persisted to the local file system in a format optimized for high-throughput processing by
downstream rating systems.

---

### Event Structure

The data generator creates events with the following structure:

- **event_source**: Unique identifier of the client
- **event_ref**: Unique identifier of the event
- **event_type**: Integer (1, 2, 3, or 5) with specific probability distribution
- **event_date**: Timestamp of the event
- **calling_number**: Source phone number (BIGINT)
- **called_number**: Destination phone number (BIGINT)
- **location**: Geographic location identifier
- **duration_seconds**: Call duration in seconds (BIGINT)
- **attr_1** through **attr_8**: Optional text attributes for additional data

---

## Performance Goals & Validation

### Target Performance

Below is actual timing & optimized time gathered on MacBook Pro 16 with M4 Max

| Number of Events | Initial Time   | Optimized Time | Improvement |
|------------------|----------------|----------------|-------------|
| 10k              | 75.437791ms    | 5.574375ms     | +1253%      |
| 100k             | 528.785833ms   | 52.891875ms    | +900%       |
| 1M               | 5.015893833s   | 445.898958ms   | +1025%      |
| 10M              | 1m5.344727125s | 8.048328959s   | +712%       |

## Need Help?

- Review the existing codebase to understand current implementation
- Run `make help` to see available commands
- Ask questions during the interview - we're here to help!

*Good luck with the optimization challenge!* 🚀