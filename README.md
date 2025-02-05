# GoKage: Full Text Search Engine

GoKage is a simple, efficient, and scalable full-text search engine built in Go. It is designed to index and search Wikipedia abstract dumps using an inverted index and standard text analysis techniques such as tokenization, stopword removal, and stemming. 

The search engine is capable of querying large datasets with minimal latency and supports efficient intersection-based searching for relevant documents.

## Table of Contents

- [Project Overview](#project-overview)
- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
  - [Running the Search Engine](#running-the-search-engine)
  - [Query Example](#query-example)
- [Code Structure](#code-structure)
- [Contributing](#contributing)
- [License](#license)

## Project Overview

GoKage loads Wikipedia abstract dumps, processes them using text analysis techniques, and provides an efficient search mechanism through an inverted index. The system indexes documents and supports text-based queries to return relevant results in a fraction of the time.

### Key Components:
1. **Tokenizer**: Splits the text into tokens (words) based on letters and numbers.
2. **Analyzer**: Processes the text by converting it to lowercase, removing stopwords, and applying stemming.
3. **Indexing**: Documents are indexed using an inverted index that maps tokens to document IDs.
4. **Search**: A query is processed by tokenizing and analyzing it. The system returns matching document IDs based on token intersections.

## Features

- **Document Indexing**: Efficiently indexes documents with text-based analysis.
- **Full-Text Search**: Queries can be run against the indexed documents, with results returned by intersecting token matches.
- **Time Efficiency**: Logs the time taken for each step (loading, indexing, and searching), allowing you to monitor performance.
- **Handling Large Dumps**: Supports large datasets such as Wikipedia abstract dumps in gzip format.

## Installation

### Prerequisites

- **Go**: Ensure that Go is installed on your machine. You can download it from [Go's official website](https://golang.org/dl/).
- **Wikipedia Dump**: Download the Wikipedia abstract dump you want to use. For example, the file `enwiki-latest-abstract1.xml.gz` can be found at [Wikipedia Dumps](https://dumps.wikimedia.org/enwiki/latest/enwiki-latest-abstract1.xml.gz).

### Clone the Repository

Clone the repository using Git:

## Usage

### Running the Search Engine

To run the search engine, use the following command:

```bash
go run main.go -p <path-to-wikipedia-dump> -q <search-query>

