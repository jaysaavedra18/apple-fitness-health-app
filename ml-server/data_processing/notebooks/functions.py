# functions.py
import pandas as pd
import matplotlib.pyplot as plt
from sklearn.model_selection import train_test_split
from sklearn.linear_model import LinearRegression
from sklearn.metrics import mean_squared_error

# Function to analyze correlation between two metrics
def analyze_metric_correlation(df, metric1, metric2, show_plot=False):
    # Function implementation
    data1 = df[df['metric'] == metric1]
    data2 = df[df['metric'] == metric2]
    merged_df = pd.merge(data1[['date', 'value']], data2[['date', 'value']], on='date', suffixes=('_' + metric1, '_' + metric2))
    correlation = merged_df['value_' + metric1].corr(merged_df['value_' + metric2])
    print(f"Correlation between {metric1} and {metric2}: {correlation}")
    
    if show_plot:
        plt.figure(figsize=(8, 5))
        plt.scatter(merged_df['value_' + metric2], merged_df['value_' + metric1], alpha=0.7)
        plt.title(f"Correlation between {metric1} and {metric2}")
        plt.xlabel(f"{metric2} value")
        plt.ylabel(f"{metric1} value")
        plt.grid(True)
        plt.show()

def analyze_workout_correlation(df, column1, column2, show_plot=False):
    """
    Analyze the correlation between two numeric columns in the DataFrame.
    
    Args:
        df (pd.DataFrame): The workouts DataFrame.
        column1 (str): First column to analyze.
        column2 (str): Second column to analyze.
        show_plot (bool): Whether to display a scatter plot of the correlation.
    """
    # Ensure the columns exist
    if column1 not in df.columns or column2 not in df.columns:
        print(f"Error: {column1} or {column2} not found in DataFrame.")
        return
    
    # Calculate correlation
    correlation = df[column1].corr(df[column2])
    print(f"Correlation between {column1} and {column2}: {correlation}")
    
    # Optionally show scatter plot
    if show_plot:
        plt.figure(figsize=(8, 5))
        plt.scatter(df[column1], df[column2], alpha=0.7)
        plt.title(f"Correlation between {column1} and {column2}")
        plt.xlabel(f"{column2} values")
        plt.ylabel(f"{column1} values")
        plt.grid(True)
        plt.show()


# Function to plot trends for a given metric
def plot_metric_trends(group, metric):
    plt.figure(figsize=(10, 5))
    plt.plot(group['date'], group['value'], marker='o', label=metric)
    plt.title(f"Trend for {metric}")
    plt.xlabel("Date")
    plt.ylabel(f"Value ({group['units'].iloc[0]})")
    plt.legend()
    plt.grid(True)
    plt.show()

# Function to plot trends for a given metric
def plot_workout_trends(workouts_group, metric, ylabel="Value"):
    plt.figure(figsize=(10, 5))
    plt.plot(workouts_group['start'], workouts_group[metric], marker='o', label=metric)
    plt.title(f"Trend for {metric}")
    plt.xlabel("Date")
    plt.ylabel(ylabel)
    plt.legend()
    plt.grid(True)
    plt.show()


def linear_regression(pivoted_df, predictor, target):
    # Calculate correlations
    # correlation_matrix = pivoted_df.corr()
    # print(correlation_matrix)

    # Prepare data
    X = pivoted_df[[predictor]].dropna()  # Predictor
    y = pivoted_df[target].dropna()     # Target
    X, y = X.align(y, join='inner', axis=0)  # Align datasets to avoid NaNs

    # Split data
    X_train, X_test, y_train, y_test = train_test_split(X, y, test_size=0.2, random_state=42)

    # Train model
    model = LinearRegression()
    model.fit(X_train, y_train)

    # Predictions
    y_pred = model.predict(X_test)

    # Evaluate model
    print("Mean Squared Error:", mean_squared_error(y_test, y_pred))
    print("Model Coefficients:", model.coef_)

    # Plot predictions
    plt.scatter(X_test, y_test, color='blue', label='Actual')
    plt.scatter(X_test, y_pred, color='red', label='Predicted')
    plt.title("Linear Regression: Exercise Time vs Swimming Distance")
    plt.xlabel("Exercise Time (seconds)")
    plt.ylabel("Swimming Distance (yd)")
    plt.legend()
    plt.show()