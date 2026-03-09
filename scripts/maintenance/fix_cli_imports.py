import os
import re

# Directories to process
packages = [
    'internal/cli/drive',
    'internal/cli/workspace', 
    'internal/cli/admin',
    'internal/cli/chat',
    'internal/cli/gmail',
    'internal/cli/people',
    'internal/cli/calendar',
    'internal/cli/activity',
    'internal/cli/sync',
    'internal/cli/system'
]

def fix_file(filepath):
    with open(filepath, 'r') as f:
        content = f.read()
    
    # Skip if not a .go file (excluding _test.go for now)
    if not filepath.endswith('.go') or filepath.endswith('_test.go'):
        return
    
    # Fix import block - ensure cli import is there and formatted correctly
    if 'cli "github.com/dl-alexandre/gdrv/internal/cli"' not in content:
        # Add import after the import (
        content = re.sub(
            r'import \(',
            'import (\n\tcli "github.com/dl-alexandre/gdrv/internal/cli"',
            content
        )
    
    # Fix malformed import that may have been added
    content = re.sub(
        r'tcli "github.com/dl-alexandre/gdrv/internal/cli"',
        '\tcli "github.com/dl-alexandre/gdrv/internal/cli"',
        content
    )
    content = re.sub(
        r'\ntcli ',
        '\n\tcli ',
        content
    )
    
    # Fix function signatures: *Globals -> *cli.Globals
    content = re.sub(
        r'func \(cmd \*(\w+)\) Run\(globals \*Globals\)',
        r'func (cmd *\1) Run(globals *cli.Globals)',
        content
    )
    
    # Fix shared function calls
    replacements = [
        ('NewOutputWriter(', 'cli.NewOutputWriter('),
        ('ResolveFileID(', 'cli.ResolveFileID('),
        ('GetGlobalFlags()', 'cli.GetGlobalFlags()'),
        ('GetLogger()', 'cli.GetLogger()'),
        ('GetPathResolver(', 'cli.GetPathResolver('),
        ('GetResolveOptions(', 'cli.GetResolveOptions('),
        ('handleCLIError(', 'cli.HandleCLIError('),
        ('convertDriveFile(', 'cli.ConvertDriveFile('),
        ('isPath(', 'cli.IsPath('),
        ('getConfigDir()', 'cli.GetConfigDir()'),
        ('resolveTimeRange(', 'cli.ResolveTimeRange('),
        ('splitCSV(', 'cli.SplitCSV('),
        ('scopesForPreset(', 'cli.ScopesForPreset('),
        ('resolveAuthScopes(', 'cli.ResolveAuthScopes('),
        ('validateAdminScopesRequireImpersonation(', 'cli.ValidateAdminScopesRequireImpersonation('),
        ('buildOAuthClientError(', 'cli.BuildOAuthClientError('),
        ('oauthClientSource', 'cli.OAuthClientSource'),
        ('oauthClientSourceFlags', 'cli.OAuthClientSourceFlags'),
        ('oauthClientSourceEnv', 'cli.OAuthClientSourceEnv'),
        ('oauthClientSourceConfig', 'cli.OAuthClientSourceConfig'),
        ('oauthClientSourceBundled', 'cli.OAuthClientSourceBundled'),
        ('isTruthyEnv(', 'cli.IsTruthyEnv('),
        ('buildAuthFlowError(', 'cli.BuildAuthFlowError('),
        ('oauthClientSecretHint(', 'cli.OAuthClientSecretHint('),
        ('openBrowser(', 'cli.OpenBrowser('),
    ]
    
    for old, new in replacements:
        content = content.replace(old, new)
    
    with open(filepath, 'w') as f:
        f.write(content)
    
    print(f"Fixed: {filepath}")

# Process all packages
for pkg in packages:
    if os.path.exists(pkg):
        for filename in os.listdir(pkg):
            if filename.endswith('.go') and not filename.endswith('_test.go'):
                filepath = os.path.join(pkg, filename)
                fix_file(filepath)

print("Done fixing imports!")
