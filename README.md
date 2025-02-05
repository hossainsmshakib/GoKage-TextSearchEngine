# GoKage: Full Text Search Engine

GoKage is a lightweight and efficient full-text search engine built in Go. It is designed to index and search large datasets, such as Wikipedia abstract dumps, using an inverted index and basic text analysis techniques like tokenization, stopword removal, and stemming.

## Features

- **Efficient Document Indexing**: GoKage indexes documents with text analysis methods for quick and efficient retrieval of results.
- **Full-Text Search**: Users can run full-text search queries against indexed documents, retrieving relevant results based on token matches.
- **Tokenization and Analysis**: Text is broken down into tokens (words), normalized to lowercase, and processed to remove unnecessary stopwords.
- **Inverted Indexing**: An inverted index maps terms to their document locations, enabling fast search results.
- **Query Support**: Users can input a search query and the engine will return matching document IDs based on token intersections.
- **Optimized for Large Datasets**: Handles large Wikipedia abstract dumps, making it scalable and practical for large-scale search applications.
- **Time Efficiency Logging**: Logs the time taken for each step (loading, indexing, searching), providing insights into the performance of the search process.
- **Simple Command-Line Interface**: Easily run the search engine via command-line inputs specifying the path to a Wikipedia dump and the query term.

## How It Works

GoKage processes a Wikipedia abstract dump by loading, analyzing, and indexing its content. When a user searches for a term, the system retrieves relevant documents based on indexed tokens.

### Workflow:

1. **Loading Data**

   - The search engine reads a Wikipedia abstract dump (`.xml.gz` file).
   - Each document is parsed and extracted from the dataset.

2. **Text Processing**

   - The text is split into tokens (words).
   - Stopwords are removed (common words like "the", "is", "and").
   - Stemming is applied (converting words to their root form, e.g., "running" → "run").

3. **Indexing**

   - An inverted index is built where each token is mapped to the documents it appears in.
   - The index allows fast retrieval of relevant documents.

4. **Querying**

   - A user inputs a search query via the command line.
   - The query is tokenized, analyzed, and matched against the index.
   - The engine finds documents containing the search tokens and returns them.

### Process Diagram

```
            +----------------------+
            |  Wikipedia Dump File  |
            +----------------------+
                      |
                      v
            +----------------------+
            |  Load & Parse Data    |
            +----------------------+
                      |
                      v
            +----------------------+
            |  Tokenization & NLP   |
            +----------------------+
                      |
                      v
            +----------------------+
            |  Inverted Indexing    |
            +----------------------+
                      |
                      v
      +----------------------------------+
      |  User Input: Search Query       |
      +----------------------------------+
                      |
                      v
            +----------------------+
            |  Query Processing     |
            +----------------------+
                      |
                      v
            +----------------------+
            |  Retrieve Documents   |
            +----------------------+
                      |
                      v
            +----------------------+
            |  Display Results      |
            +----------------------+
```

## Dataset

To use this search engine, you need to download the Wikipedia abstract dump. You can get the latest dump from:\
[Wikipedia Dumps - Abstracts](https://dumps.wikimedia.org/enwiki/latest/enwiki-latest-abstract1.xml.gz)

## Usage

Run the search engine with the following command:

```bash
go run main.go -p <path-to-wikipedia-dump> -q "<search-query>"
```

Example:

```bash
go run main.go -p enwiki-latest-abstract1.xml.gz -q "Small wild cat"
```

This command will:

- Load the documents from the Wikipedia dump file.
- Index the documents.
- Search for the term "Small wild cat" and return any matching documents.
- Output the matched documents along with the time taken for the indexing and search processes.

## Code Structure

The project is organized as follows:

```
/GoKage
├── main.go             # Main entry point for running the search engine
├── /utils              # Utility functions (tokenization, indexing, etc.)
│   ├── analyze.go      # Text analysis functions (tokenization, stopword removal, etc.)
│   ├── index.go        # Inverted index functions (adding documents, searching)
│   ├── load.go         # Functions for loading the Wikipedia dump
├── /data               # Data directory (can contain Wikipedia dumps)
├── go.mod              # Go module dependencies
└── README.md           # Project documentation
```

- **main.go**: The main program that loads documents, creates an index, and runs the search query.
- **/utils**: A package containing utility functions used throughout the search engine.
  - **analyze.go**: Functions for tokenization, stopword removal, and stemming.
  - **index.go**: Functions for indexing documents and performing search queries.
  - **load.go**: Functions for loading and parsing Wikipedia dump files.

## Screenshot

Here’s a performance screenshot of the search engine in action:



## Contributing

We welcome contributions! If you'd like to contribute to GoKage, follow these steps:

1. **Fork the Repository**: Create a personal copy of the repository.
2. **Create a New Branch**: Work on your feature or fix on a new branch.

```bash
git checkout -b feature/my-new-feature
```

3. **Make Changes**: Implement your changes and commit them.

```bash
git commit -am "Add feature or fix bug"
```

4. **Push to Your Fork**: Push your changes to your forked repository.

```bash
git push origin feature/my-new-feature
```

5. **Create a Pull Request**: Open a pull request to merge your changes into the main repository.

Please ensure that your code is well-documented, follows Go's best practices, and adheres to the standard Go coding conventions. Tests are encouraged for new features.

## License

GoKage is licensed under the MIT License. See the LICENSE file for more information.

make the abouve plain text as a readme text format

