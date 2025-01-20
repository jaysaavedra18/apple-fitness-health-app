import json
from typing import Dict, Any, Union
import sys

def get_go_type(value: Any, parent_name: str = "", field_name: str = "") -> str:
    """Determine the Go type based on the Python value type."""
    if isinstance(value, bool):
        return "bool"
    elif isinstance(value, int):
        return "int64"
    elif isinstance(value, float):
        return "float64"
    elif isinstance(value, str):
        return "string"
    elif isinstance(value, list):
        if value:  # If list is not empty
            if isinstance(value[0], dict):
                # Create a new struct type for array elements
                struct_name = f"{parent_name}{field_name}"
                return f"[]{struct_name}"
            return f"[]{ get_go_type(value[0], parent_name, field_name) }"
        return "[]interface{}"
    elif isinstance(value, dict):
        return "struct"
    elif value is None:
        return "interface{}"
    return "interface{}"

def generate_struct(name: str, data: Union[Dict, list], indent: str = "") -> str:
    """Generate Go struct definition from dictionary or list."""
    output = []
    
    # If data is a list and contains dictionaries, use the first item
    if isinstance(data, list) and data and isinstance(data[0], dict):
        data = data[0]
    
    if not isinstance(data, dict):
        return ""

    output.append(f"{indent}type {name} struct {{")
    
    for key, value in data.items():
        field_name = key.title()  # Capitalize first letter for Go export
        go_type = get_go_type(value, name, field_name)
        
        # Handle nested structures
        if isinstance(value, dict):
            nested_name = f"{name}{field_name}"
            nested_struct = generate_struct(nested_name, value, indent)
            if nested_struct:
                output.insert(0, nested_struct + "\n")
            go_type = nested_name
        
        # Handle arrays of objects
        elif isinstance(value, list) and value and isinstance(value[0], dict):
            nested_name = f"{name}{field_name}"
            nested_struct = generate_struct(nested_name, value[0], indent)
            if nested_struct:
                output.insert(0, nested_struct + "\n")
        
        output.append(f'{indent}\t{field_name} {go_type} `json:"{key}"`')
    
    output.append(indent + "}")
    return "\n".join(output)

def main():
    if len(sys.argv) != 2:
        print("Usage: python script.py <json_file>")
        sys.exit(1)
        
    json_file = sys.argv[1]
    
    try:
        with open(json_file, 'r') as f:
            data = json.load(f)
            
        # Generate the root struct
        if isinstance(data, dict):
            print(generate_struct("Root", data))
        elif isinstance(data, list) and data:
            print(generate_struct("Root", data[0]))
        else:
            print("Error: JSON must be an object or non-empty array")
            sys.exit(1)
            
    except json.JSONDecodeError:
        print("Error: Invalid JSON file")
        sys.exit(1)
    except FileNotFoundError:
        print(f"Error: File '{json_file}' not found")
        sys.exit(1)

if __name__ == "__main__":
    main()