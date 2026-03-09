import os
import re

packages = {
    'drive': 'internal/cli/drive',
    'workspace': 'internal/cli/workspace',
    'admin': 'internal/cli/admin',
    'chat': 'internal/cli/chat',
    'gmail': 'internal/cli/gmail',
    'people': 'internal/cli/people',
    'calendar': 'internal/cli/calendar',
    'activity': 'internal/cli/activity',
    'sync': 'internal/cli/sync',
    'system': 'internal/cli/system',
}

for pkg_name, pkg_path in packages.items():
    if not os.path.exists(pkg_path):
        continue
    
    for filename in os.listdir(pkg_path):
        if not filename.endswith('.go'):
            continue
            
        filepath = os.path.join(pkg_path, filename)
        with open(filepath, 'r') as f:
            content = f.read()
        
        # Fix package declaration
        content = re.sub(r'^package \w+', f'package {pkg_name}', content)
        
        # Ensure base import is present
        if 'github.com/dl-alexandre/gdrv/internal/cli/base' not in content:
            if 'import (' in content:
                content = content.replace(
                    'import (',
                    'import (\n\t"github.com/dl-alexandre/gdrv/internal/cli/base"'
                )
            else:
                # Add single import
                content = re.sub(
                    r'^import "([^"]+)"',
                    r'import (\n\t"github.com/dl-alexandre/gdrv/internal/cli/base"\n\t"\1"\n)',
                    content,
                    flags=re.MULTILINE
                )
        
        with open(filepath, 'w') as f:
            f.write(content)
        
        print(f"Fixed: {filepath}")

print("Done!")
