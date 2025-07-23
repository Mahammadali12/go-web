workflow-automation/
├── cmd/
│   └── api/
│       └── main.go           # Entry point for the Go application
├── internal/
│   ├── handlers/            # HTTP handlers for API and HTML endpoints
│   │   ├── auth.go          # Authentication handlers (register, login)
│   │   ├── workflows.go     # Workflow and task CRUD handlers
│   │   └── templates.go     # Handlers for rendering HTML with HTMX
│   ├── models/              # Data models (User, Workflow, Task)
│   │   ├── user.go
│   │   ├── workflow.go
│   │   └── task.go
│   ├── database/            # Database connection and queries
│       └── db.go
│   ├── scheduler/           # Task scheduling logic
│       └── scheduler.go
├── templates/               # HTML templates with HTMX
│   ├── base.html            # Base template with layout
│   ├── login.html           # Login page
│   ├── register.html        # Registration page
│   ├── workflows.html       # Workflow list page
│   └── workflow_form.html   # Form for creating/editing workflows
├── static/                  # Static assets (CSS, HTMX)
│   ├── styles.css           # Basic CSS for styling
│   └── htmx.min.js          # HTMX library
├── tests/                   # Unit and integration tests
│   ├── auth_test.go
│   ├── workflows_test.go
│   └── scheduler_test.go
├── docker/
│   ├── Dockerfile           # Dockerfile for the Go application
│   └── docker-compose.yml   # Docker Compose for local development
├── db/
│   └── init.sql             # SQL script for database initialization
├── .env                     # Environment variables
├── go.mod                   # Go module file
├── go.sum                   # Go dependencies checksum
└── README.md                # Project documentation