# Backend Server Technical Specs

## Business Goal:

A language learning school wants to build a prototype of learning portal which will act as three things:
- Inventory of possible vocabulary that can be learned
- Act as a  Learning record store (LRS), providing correct and wrong score on practice vocabulary
- A unified launchpad to launch different learning apps

## Technical Requirements

- The backend will be built using Go
- The database will be SQLite3
- The API will be built using Gin
- The API will always return JSON
- There will no authentication or authorization
- Everything be treated as a single user

## Database Schema

Our database will be a single sqlite database called `words.db` that will be in the root of the project folder of `backend_go`.

We have the following tables:
- words - stored vocabulary words
  - id integer
  - bahasa_indonesia string
  - english string
  - parts json
- words_groups - join table for words and groups many-to-many
  - id integer
  - word_id integer
  - group_id integer
- groups - thematic groups of words
  - id integer
  - name string
- study_sessions - records of study sessions grouping word_review_items
  - id integer
  - group_id integer
  - created_at datetime
  - study_activity_id integer
- study_activities - a specific study activity, linking a study session to group
  - id integer
  - study_session_id integer
  - group_id integer
  - created_at datetime
- word_review_items - a record of word practice, determining if the word was correct or not
  - word_id integer
  - study_session_id integer
  - correct boolean
  - created_at datetime

## API Endpoints
### GET /api/dashboard/last_study_session
#### JSON Response
```json
{
  "id": 1,
  "activity_name": "Vocabulary Practice",
  "last_used": "2025-02-15T10:00:00Z",
  "correct_count": 10,
  "wrong_count": 2,
  "group_id": 1,
  "group_name": "Basic Vocabulary"
}
```


### GET /api/dashboard/study_progress
#### JSON Response
```json
{
  "total_words_studied": 124,
  "mastery_progress": 0
}
```

### GET /api/dashboard/quick-stats
#### JSON Response
```json
{
  "success_rate": 80,
  "total_study_sessions": 4,
  "total_active_groups": 3,
  "study_streak": 4
}
```

### GET /api/study_activities/:id
#### JSON Response
```json
{
  "id": 1,
  "name": "Vocabulary Practice",
  "thumbnail_url": "url_to_thumbnail",
  "description": "Practice basic vocabulary",
}
```

### GET /api/study_activities/:id/study_sessions
  - pagination with 100 per page
#### JSON Response
```json
{
  "items": [
    {
      "id": 1,
      "activity_name": "Vocabulary Practice",
      "group_name": "Basic Vocabulary",
      "start_time": "2025-02-15T10:00:00Z",
      "end_time": "2025-02-15T10:30:00Z",
      "number_of_review_items": 12
    }
  ],
  "pagination": {
    "page": 1,
    "per_page": 100,
    "total_pages": 1,
    "total_items": 1
  }
}
```

### POST /api/study_activities
#### Request Params
  - group_id integer
  - study_activity_id integer

### JSON Response
```json
{
  "id": 1,
  "group_id": 1
}
```

### GET /api/words
  - pagination with 100 items per page
#### JSON response
```json
{
  "items": [
    {
      "id": 1,
      "bahasa_indonesia": "kata",
      "english": "word",
      "correct_count": 10,
      "wrong_count": 2
    }
  ],
  "pagination": {
    "page": 1,
    "per_page": 100,
    "total_pages": 1,
    "total_items": 1
  }
}
```

### GET /api/words/:id
#### JSON Response
```json
{
  "id": 1,
  "bahasa_indonesia": "kata",
  "english": "word",
  "correct_count": 10,
  "wrong_count": 2,
  "groups": [
    {
      "id": 1,
      "name": "Basic Vocabulary"
    }
  ]
}
```

### GET /api/groups
  - pagination with 100 items per page
#### JSON Response
```json
{
  "items": [
    {
      "id": 1,
      "name": "Basic Vocabulary",
      "word_count": 100
    }
  ],
  "pagination": {
    "page": 1,
    "per_page": 100,
    "total_pages": 1,
    "total_items": 1
  }
}
```

### GET /api/groups/:id
#### JSON Response
```json
{
  "id": 1,
  "name": "Basic Vocabulary",
  "word_count": 100
}
```

### GET /api/groups/:id/words
  - pagination with 100 items per page
#### JSON Response
```json
{
  "items": [
    {
      "id": 1,
      "bahasa_indonesia": "kata",
      "english": "word",
      "correct_count": 10,
      "wrong_count": 2
    }
  ],
  "pagination": {
    "page": 1,
    "per_page": 100,
    "total_pages": 1,
    "total_items": 1
  }
}
```

### GET /api/groups/:id/study_sessions
  - pagination with 100 items per page
#### JSON Response
```json
{
  "items": [
    {
      "id": 1,
      "activity_name": "Vocabulary Practice",
      "group_name": "Basic Vocabulary",
      "start_time": "2025-02-15T10:00:00Z",
      "end_time": "2025-02-15T10:30:00Z",
      "number_of_review_items": 12
    }
  ],
  "pagination": {
    "page": 1,
    "per_page": 100,
    "total_pages": 1,
    "total_items": 1
  }
}
```

### GET /api/study_sessions
  - pagination with 100 items per page
#### JSON Response
```json
{
  "items": [
    {
      "id": 1,
      "activity_name": "Vocabulary Practice",
      "group_name": "Basic Vocabulary",
      "start_time": "2025-02-15T10:00:00Z",
      "end_time": "2025-02-15T10:30:00Z",
      "number_of_review_items": 12
    }
  ],
  "pagination": {
    "page": 1,
    "per_page": 100,
    "total_pages": 1,
    "total_items": 1
  }
}
```

### GET /api/study_sessions/:id
#### JSON Response
```json
{
  "id": 1,
  "activity_name": "Vocabulary Practice",
  "group_name": "Basic Vocabulary",
  "start_time": "2025-02-15T10:00:00Z",
  "end_time": "2025-02-15T10:30:00Z",
  "number_of_review_items": 12
}
```

### GET /api/study_sessions/:id/words
#### JSON Response
```json
{
  "items": [
    {
      "id": 1,
      "bahasa_indonesia": "kata",
      "english": "word",
      "correct": true,
      "created_at": "2025-02-15T10:05:00Z"
    }
  ],
  "pagination": {
    "page": 1,
    "per_page": 100,
    "total_pages": 1,
    "total_items": 1
  }
}
```

### POST /api/reset_history
#### JSON Response
```json
{
  "success": true,
  "message": "History reset successfully"
}
```

### POST /api/full_reset
#### JSON Response
```json
{
  "success": true
  "message": "Full reset successfully"
}
```

### POST /api/study_sessions/:id/words/:word_id/review
#### Request Params
  - id (study_session_id) integer
  - word_id integer
  - correct boolean

#### Request Payload
```json
{
  "correct": true
}
```

### JSON Response
```json
{
  "id": 1,
  "word_id": 1,
  "study_session_id": 1,
  "correct": true,
  "created_at": "2025-02-15T10:05:00Z"
}
```

## Mage Tasks

Mage is a task runner for Go.
Lets list out possible tasks we need for our lang portal.

### Initialize Database
This task will initialize the sqlite database called `words.db`.

### Migrate Database
This task will run a series of migrations sql files on the database.

Migrations live in the `migrations` folder.
The migration files will be run in order of their file name.
The file names should look like this:
```sql
0001_init.sql
0002_create_words_table.sql
```

### Seed Data
This task will import json files and transform them into target data for our database.

All seed files live in the `seeds` folder.
All seed files should be loaded.

In our task we should have DSL to specific each seed file and its expected group word name.
```json
[
  {
    "bahasa_indonesia": "kata",
    "english": "word"
  },
  ...
]
```
