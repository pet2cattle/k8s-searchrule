import yaml
import sys

# Read in the YAML file
data = yaml.safe_load(sys.stdin)

# Define the desired values
api_groups = sys.argv[1]
resources = sys.argv[2]
verbs = sys.argv[3]

# Find the rules that contain the desired values
try:
    matching_rules = []
    for rule in data['rules']:
        if api_groups in rule.get('apiGroups', []) :
            if resources in rule.get('resources', []) or '*' in rule.get('resources', []):
                if verbs in rule.get('verbs', []) or '*' in rule.get('verbs', []):
                    matching_rules.append(rule)

    # Print name of the resource
    if matching_rules:
        print(data['kind'].lower()+"/"+data['metadata']['name'])
except:
    pass
