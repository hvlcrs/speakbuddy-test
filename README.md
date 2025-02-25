# Speakbuddy

This application provides audio file transcoding services. It allows users to upload audio files and retrieve them in different formats.

## Dependencies

This applications use following applications

* PostgreSQL
* ffmpeg

## Docker Build and Run

1. Clone the repository:

    ```sh
    git clone https://github.com/yourusername/speakbuddy.git
    ```

2. Navigate to the project directory:

    ```sh
    cd speakbuddy
    ```

3. Build and run the docker image

    ```sh
    docker-compose up
    ```

4. By default the application will run on port 8081

5. Access the API documentation at `http://localhost:8081/swagger/index.html`

## Local Installation

1. Clone the repository:

    ```sh
    git clone https://github.com/yourusername/speakbuddy.git
    ```

2. Navigate to the project directory:

    ```sh
    cd speakbuddy
    ```

3. Install dependencies:

    ```sh
    go mod tidy
    ```

4. Build the project

    ```sh
    go build speakbuddy
    ```

## Local Usage

1. Install the dependencies

2. Start the application:

    ```sh
    go run speakbuddy
    ```

3. By default the application will run on port 8081

4. Access the API documentation at `http://localhost:8081/swagger/index.html`

## Configurations

This application can be dynamically configured during startup time by supplying the configuration through
environment variables.

All configurations default value can be found within `./pkg/configs` package

### Application Configuration

* `FILE_SAVE_PATH` default file upload directory
* `FILE_MAX_SIZE` maximum file upload size
* `FILE_ALLOW_EXTS` allowed uploaded file type in strings separated by comma (ex: .mp4a,.mp3,.wav)
* `FILE_TARGET_EXT` target file type in strings (ex: .wav)

### Server Configuration

* `HTTP_PORT` server http port
* `READ_TIMEOUT` server read timeout
* `WRITE_TIMEOUT` server write timeout

### Database Configuration

* `DB_HOST` database hostname
* `DB_PORT` database connection port
* `DB_USER` database connection username
* `DB_PASSWORD` database connection password
* `DB_NAME` default database name
* `MAX_IDLE_CONNECTION`maximum idle connection
* `MAX_OPEN_CONNECTION` maximum open connection
