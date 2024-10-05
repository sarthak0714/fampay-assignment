# Backend Assignment (Intern) | FamPay 📹

An API to fetch latest videos sorted in reverse chronological order of their publishing date-time from YouTube for a given tag/search query in a paginated response.

## Tech Stack
- Echo (Golang 1.22)
- SQLite
- HTMX
- Templ
- Tailwind CSS

## Setup
1. Clone the repository:
   ```
   git clone https://github.com/sarthak0714/fampay-assignment.git
   cd fampay-assignment
   ```
2. Create a `.env` file with the following content:
   ```
   DB_PATH=./test.db
   YOUTUBE_API_KEYS=your_youtube_api_key1,your_youtube_api_key2,...
   SEARCH_QUERY=your_search_query
   ```
    By default search query is set to `Marvel`

3. Install dependencies:
   ```
   go mod download
   ```

4. Run the Server
    ```
    make run
    ```


## API 

1. `GET /api/video`
    - Query Parameters:
        - `page`: Page number (default: 1)
        - `size`: Number of videos per page (default: 9)
    - Example: `http://localhost:8080/api/video?page=1&size=9`

## Frontend 

```http://localhost:8080/```

- The video page displays fetched videos with pagination.
- The api will return RAW JSON API.