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

## Dataset  

To use this search engine, you need to download the Wikipedia abstract dump. You can get the latest dump from:  
[Wikipedia Dumps - Abstracts](https://dumps.wikimedia.org/enwiki/latest/enwiki-latest-abstract1.xml.gz)  

## Screenshot  

Here’s a performance screenshot of the search engine in action:  

![GoKage Performance](![image](https://github.com/user-attachments/assets/f3713529-fb25-4ba3-abd6-2effbe4beb31)
)  

*(Replace `path-to-your-screenshot.png` with the actual path or URL of your screenshot.)*  

## License  

GoKage is licensed under the MIT License. See the LICENSE file for more information.  
