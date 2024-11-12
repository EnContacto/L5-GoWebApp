# L5-GoWebApp
Web application developed in Go that shows the elapsed hours and how many Monday days have passed from the beginning of the year to the current day.

## Description
This is a simple web application built with Go. 
  - Displays the current time
  - shows elapsed hours 
  - Displays how many Monday days have passed since the beginning of this year.

The application is designed to run on port `5000` by default.

## Prerequisites
To run this application, ensure that the following are installed on your system:
- **Ruby** (version 3.3.6 or later)
- **Sinatra** (web framework)
- **Rackup**
- **Webrick**
- **Bundler**

## How to Run the Application

### Running Locally
To run the application on your local machine:
1. Clone or download this repository to your computer.
   ```bash
   git clone https://github.com/EnContacto/L5-GoWebApp.git
   cd L5-GoWebApp.git
2. Open a terminal in the root directory of the project.
3. Start app with:
   ```bash
   go run main.go
Open your web browser and visit `http://localhost:5000` to see the application at work.
 
### Running with Docker
You can also run this application using Docker for a containerized environment.
1. In the project’s root directory, open a terminal and run:
   ```bash
   docker build -t go-webapp .

2. After the image is built, run the container using:
   ```bash
   docker run -p 5000:5000 go-webapp
The application should now be accessible at `http://localhost:5000` in your web browser.

## Troubleshooting
  - Ensure Docker is installed and running correctly if using the Docker setup.
  - Make sure no other application is using port `5000` before running the server.
