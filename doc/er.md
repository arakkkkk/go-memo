```mermaid
erDiagram
  users {
    string name
    string email
    integer age
  }

  articles {
    string title
    text text
  }
  users ||--o{ articles: ""
```

```mermaid
erDiagram
  SearchHistory {
    string query
    string created_at
  }

  SearchRecord {
    foreign SearchHistoryId
    string created_at
  }
  SearchHistory ||--o{ SearchRecord: ""

  Users {
    string name
    string password
    string created_at
  }

  Record {
    string key
    text description
  }

  LikeRecord {
    string RecordId
    string UserId
    string created_at
  }
  Record ||--o{ LikeRecord: ""
  Users ||--o{ LikeRecord: ""

  Bookmarks {
    string RecordId
    string UserId
    string created_at
  }
  SearchHistory ||--o{ SearchRecord: ""
  Record ||--o{ Bookmarks: ""
  Users ||--o{ Bookmarks: ""
```
