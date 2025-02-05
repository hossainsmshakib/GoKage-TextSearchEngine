# GoKage: A Full-Text Search Engine

GoKage is a lightweight and efficient full-text search engine built in Go. It processes large Wikipedia abstract dumps and returns relevant documents based on user queries. GoKage uses an inverted index, tokenization, stopword removal, stemming, and set theory for optimal search performance.

---

## Features

- **Tokenization**: Breaks down text into individual tokens (words).
- **Stopword Removal**: Removes common, insignificant words to improve search accuracy.
- **Stemming**: Reduces words to their root forms to enhance matching.
- **Inverted Index**: Efficiently maps words to document IDs for quick search retrieval.
- **Set Union**: Uses set theory (union) to combine results from multiple search tokens.
- **Efficient Search**: Finds relevant documents in large datasets quickly.

---

## Installation

### Prerequisites

- Go 1.16+ is required.

### Clone the Repository

```bash
git clone https://github.com/hossainsmshakib/GoKage.git
cd GoKage
