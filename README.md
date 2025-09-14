# MindEase

MindEase is a web application that connects clients with mental health providers. Clients can search for providers, book appointments, and manage their mental health journey. Providers can create profiles, manage their availability, and connect with clients.

## Features

*   **User Authentication:** Secure registration and login for both clients and providers.
*   **Provider Profiles:** Providers can create and manage detailed profiles, including their specialty, bio, and years of experience.
*   **Client Profiles:** Clients can create and manage their profiles.
*   **Search and Filter:** Clients can search for providers based on their needs and preferences.
*   **Appointment Booking:** Clients can book appointments with available providers.
*   **Notifications:** Users receive notifications about their appointments and other important events.
*   **Prescription Management:** Providers can create and manage prescriptions for their clients.

## Tech Stack

*   **Backend:** Go
*   **Frontend:** HTML, CSS, JavaScript
*   **Database:** PostgreSQL
*   **Routing:** go-chi/chi
*   **Session Management:** alexedwards/scs
*   **CSRF Protection:** justinas/nosurf

## Getting Started

### Prerequisites

*   PostgreSQL
*   Go

### Installation

1.  **Clone the repository:**
    ```bash
    git clone https://github.com/jofosuware/mindease.git
    cd mindease
    ```

2.  **Install dependencies:**
    ```bash
    go mod tidy
    ```

### Database Setup

1.  **Create a PostgreSQL database:**
    ```sql
    CREATE DATABASE mindease;
    ```

2.  **Connect to the database** and run the SQL commands in the `migrations/schema.sql` file to create the necessary tables.

    Alternatively, you can run the `up` migration files in the `migrations` directory.

### Running the Application

1.  **Run the web server:**
    ```bash
    go run ./cmd/web
    ```
    The application will be available at `http://localhost:8080`.

2.  **Run the API server:**
    ```bash
    go run ./cmd/api
    ```
    The API will be available at `http://localhost:8081`.