# VIMANIAC
### Collaborative Text Editor in Go

Welcome to our collaborative text editor repository! This project is a simple yet powerful text editor written in Go, designed for collaborative editing. With this editor, multiple users can work on the same document simultaneously, making it a great tool for team collaboration.

## Setup

To get started, follow these simple steps:

1. Clone the repository:

   ```bash
   git clone https://github.com/your-username/collaborative-text-editor.git
   ```

2. Navigate to the project directory:

   ```bash
   cd collaborative-text-editor
   ```

3. For client use, you have two modes: collaborative mode and normal mode.

   - **Collaborative Mode:**
   
     Use the following flags to connect to a collaborative session:

     ```bash
     go run main.go --client --connect "ip"
     ```

     Replace "ip" with the actual IP address of the server you want to connect to.

   - **Normal Mode:**

     Run the editor without any flags for normal mode:

     ```bash
     go run main.go
     ```

     Or, if you want to open a specific file:

     ```bash
     go run main.go filename
     ```

## Usage

Once the server is running and clients have connected, you can collaboratively edit documents in real-time. In normal mode, you can open, edit, and save files as you would in a regular text editor.

## Features

- **Real-time Collaboration:** Collaborate with team members on the same document in real-time.
- **Flexible Modes:** Choose between collaborative mode for joint editing or normal mode for individual use.
- **Simple Setup:** Clone the repository, run the server, and connect clients easily.

Feel free to explore and contribute to make this collaborative text editor even more powerful! If you encounter any issues or have suggestions, please create an issue in the repository. Happy collaborating!
