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

## Dataset  

To use this search engine, you need to download the Wikipedia abstract dump. You can get the latest dump from:  
[Wikipedia Dumps - Abstracts](https://dumps.wikimedia.org/enwiki/latest/enwiki-latest-abstract1.xml.gz)  

## Screenshot  

Here’s a performance screenshot of the search engine in action:  

![image](https://github.com/user-attachments/assets/bada1ce2-075d-44d5-868a-4f6a85458142)



## License  

GoKage is licensed under the MIT License. See the LICENSE file for more information.  
