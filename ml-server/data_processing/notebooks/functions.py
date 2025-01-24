# functions.py
import pandas as pd
import matplotlib.pyplot as plt

def analyze_correlation(df, metric1, metric2, show_plot=False):
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
