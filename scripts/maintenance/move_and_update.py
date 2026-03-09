import os
import shutil
import re

# Mapping of files to their destination packages
file_mapping = {
    # Drive operations
    'files.go': 'drive',
    'folders.go': 'drive',
    'drives.go': 'drive',
    'permissions.go': 'drive',
    'changes.go': 'drive',
    'labels.go': 'drive',
    
    # Workspace
    'sheets.go': 'workspace',
    'docs.go': 'workspace',
    'slides.go': 'workspace',
    'forms.go': 'workspace',
    'appscript.go': 'workspace',
    
    # Admin
    'admin.go': 'admin',
    'iamadmin.go': 'admin',
    'groups.go': 'admin',
    
    # Other services
    'chat.go': 'chat',
    'gmail.go': 'gmail',
    'people.go': 'people',
    'calendar.go': 'calendar',
    'tasks.go': 'calendar',
    'meet.go': 'calendar',
    'activity.go': 'activity',
    'sync.go': 'sync',
    
    # System
    'auth.go': 'system',
    'config.go': 'system',
    'about.go': 'system',
    'completion.go': 'system',
    'ai.go': 'system',
    'cloudlogging.go': 'system',
    'monitoring.go': 'system',
    
    # Test files
    'sheets_test.go': 'workspace',
    'docs_test.go': 'workspace',
    'slides_test.go': 'workspace',
    'admin_test.go': 'admin',
    'auth_error_test.go': 'system',
    'parse_test.go': 'system',
    'config_parse_bool_test.go': 'system',
}

cli_dir = 'internal/cli'

# Move files and update package declarations
for filename, pkg in file_mapping.items():
    src = os.path.join(cli_dir, filename)
    dst_dir = os.path.join(cli_dir, pkg)
    dst = os.path.join(dst_dir, filename)
    
    if os.path.exists(src):
        # Ensure destination directory exists
        os.makedirs(dst_dir, exist_ok=True)
        
        # Read file content
        with open(src, 'r') as f:
            content = f.read()
        
        # Update package declaration
        content = re.sub(r'^package cli$', f'package {pkg}', content, flags=re.MULTILINE)
        
        # Add import for base package
        if 'import (' in content:
            content = content.replace(
                'import (',
                'import (\n\t"github.com/dl-alexandre/gdrv/internal/cli/base"'
            )
        else:
            # Single line import
            content = re.sub(
                r'^import "([^"]+)"$',
                r'import (\n\t"github.com/dl-alexandre/gdrv/internal/cli/base"\n\t"\1"\n)',
                content,
                flags=re.MULTILINE
            )
        
        # Update references to shared utilities
        content = content.replace('NewOutputWriter(', 'base.NewOutputWriter(')
        content = content.replace('handleCLIError(', 'base.HandleCLIError(')
        content = content.replace('convertDriveFile(', 'base.ConvertDriveFile(')
        content = content.replace('truncate(', 'base.Truncate(')
        content = content.replace('formatSize(', 'base.FormatSize(')
        
        # Write to destination
        with open(dst, 'w') as f:
            f.write(content)
        
        # Remove original
        os.remove(src)
        print(f"Moved: {filename} -> {pkg}/")

print("Done moving files!")
