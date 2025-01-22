import os
import subprocess
import sys

# Define the backend directory
backend_dir = './backend'

# List of available Go commands
commands = {
    'build': 'go build',
    'run': 'go run main.go',
    'tidy': 'go mod tidy'
}

# Check if the argument is provided and valid
if len(sys.argv) < 2 or sys.argv[1] not in commands:
    print("Usage: python run_backend.py <build|run|tidy> [flags]")
    sys.exit(1)

# Get the command from the first argument
command = commands[sys.argv[1]]

# For the 'run' command, append any flags passed as arguments
if sys.argv[1] == 'run' and len(sys.argv) > 2:
    flags = ' '.join(sys.argv[2:])  # Combine all remaining arguments
    command = f"{command} {flags}"

# Change to the backend directory
os.chdir(backend_dir)

# Run the Go command
# print(f"Running: {commands[sys.argv[1]]}")
subprocess.run(command, shell=True, check=True)
