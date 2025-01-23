# Stride - Apple Fitness Data Analysis

Stride is a comprehensive full-stack application for managing, analyzing, and visualizing Apple Health & Fitness data. The application now consists of multiple interconnected components designed to provide powerful insights into your fitness journey.

## Architecture

The Stride application is built with a modern, modular architecture:

- **Server (Go)**: RESTful API for data management and processing
- **Frontend (React)**: Interactive data visualization interface
- **Machine Learning (Python)**: Neural network for advanced data analysis and predictions

## Key Features

- **Automated Data Imports**: Leverages the iOS app, _Health Auto Export_, to import data from Apple Health into iCloud Drive
- **Comprehensive Data Processing**:
  - Local data caching to reduce redundant calls
  - Advanced data aggregation and analysis
- **Flexible Visualization**: Interactive web interface for exploring fitness metrics
- **Machine Learning Insights**:
  - Neural network-powered predictions
  - Data clustering and pattern recognition
  - Advanced fitness trend analysis

## Components

### Backend (Go)

A robust RESTful API built in Go that handles:

- Data import from iCloud Drive
- Data caching and management
- Expose endpoints for data retrieval and processing

#### Prerequisites

- Go 1.20+
- Configured iCloud Drive with Apple Health data exports

#### Backend Setup

```bash
# Clone the repository
git clone https://github.com/jaysaavedra18/apple-fitness-health-app.git

# Navigate to go_server directory
cd apple-fitness-health-app/go_server

# Build the application
go build

# Run the server
./fitness
```

### Frontend (React)

A modern, responsive web application built with:

- TypeScript
- React
- TailwindCSS

#### Frontend Setup

```bash
# Navigate to frontend directory
cd ../frontend

# Install dependencies
npm install

# Start development server
npm start
```

### Machine Learning Backend (Python)

A sophisticated data analysis component using:

- Python
- PyTorch
- Neural network models for:
  - Predictive analytics
  - Workout pattern recognition
  - Performance trend analysis

#### Machine Learning Setup

```bash
# Navigate to ML Backend directory
cd ../ml-backend

# Create virtual environment
python -m venv venv
source venv/bin/activate

# Install dependencies
pip install -r requirements.txt

# Run the ML server
python server.py
```

## How It Works

1. **Data Export**

   - Use _Health Auto Export_ iOS app to define export settings
   - Configure data points, format (CSV/JSON), and export frequency

2. **Data Import**

   - Go server automatically imports data from iCloud Drive
   - Caches data locally for efficient processing

3. **Data Visualization**

   - React frontend provides interactive dashboards
   - Explore detailed fitness metrics and trends

4. **Advanced Analytics**
   - Python backend with PyTorch applies machine learning techniques
   - Generate predictions and insights from your fitness data

## Contributions

Contributions are welcome! Please:

- Open issues for bugs or feature requests
- Submit pull requests with improvements
- Follow project coding standards

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contact

For questions or support, please open an issue on the GitHub repository.
