# Medical Chatbot Project

## Overview
This Medical Chatbot is designed to provide answers to medical-related questions using OpenAI's GPT-3 model. It features a simple Go server for the backend and a basic HTML/JavaScript frontend.

## Prerequisites
- [Go](https://golang.org/dl/) (version 1.16 or later)
- An API key from [OpenAI](https://openai.com/)

## Setup
1. **Clone the Repository**

2. **OpenAI API Key:**
- Obtain an API key from OpenAI by signing up at their [developer platform](https://platform.openai.com/signup).
- Set your OpenAI API key in the `main.go` file:
  ```go
  const openAIKey = "YOUR_OPENAI_API_KEY" // Replace with your actual API key
  ```

## Running the Server
1. **Start the Server:**
- In the project directory, run:
  ```
  go run main.go
  ```
- This will start the server on `localhost` at port `8080`.

2. **Accessing the Chatbot:**
- Open your web browser and go to `http://localhost:8080`.
- You should see a simple chatbot interface.

3. **Using the Chatbot:**
- Type your medical-related question in the input box and click "Ask".
- The response will be displayed below the input box.

## Project Structure
- `main.go`: The Go server handling the backend logic.
- `index.html`: The frontend HTML file for the chatbot interface.
- `script.js`: JavaScript file handling frontend interactions.

## Features
- Interactive chat interface for asking questions.
- Integration with OpenAI's GPT-3 for processing and answering queries.

## Limitations and Notes
- This chatbot is a basic demonstration and should not be used as a substitute for professional medical advice.
- Ensure that the API key is kept secure and not exposed in the frontend code.
- The chatbot's accuracy and reliability depend on the model used and the input provided.

## Contributing
Feel free to fork the repository and submit pull requests. For major changes, please open an issue first to discuss what you would like to change.

