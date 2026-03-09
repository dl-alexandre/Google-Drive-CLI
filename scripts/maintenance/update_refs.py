import os
import re

packages = ['drive', 'workspace', 'admin', 'chat', 'gmail', 'people', 'calendar', 'activity', 'sync', 'system']

replacements = [
    (r'\bNewOutputWriter\(', 'base.NewOutputWriter('),
    (r'\bhandleCLIError\(', 'base.HandleCLIError('),
    (r'\bconvertDriveFile\(', 'base.ConvertDriveFile('),
    (r'\btruncate\(', 'base.Truncate('),
    (r'\bformatSize\(', 'base.FormatSize('),
    (r'\bisPath\(', 'cli.IsPath('),  # These are in root cli
    (r'\bgetConfigDir\(\)', 'cli.GetConfigDir()'),
    (r'\bresolveTimeRange\(', 'cli.ResolveTimeRange('),
    (r'\bsplitCSV\(', 'cli.SplitCSV('),
    (r'\bscopesForPreset\(', 'cli.ScopesForPreset('),
    (r'\bresolveAuthScopes\(', 'cli.ResolveAuthScopes('),
    (r'\bvalidateAdminScopesRequireImpersonation\(', 'cli.ValidateAdminScopesRequireImpersonation('),
    (r'\bbuildOAuthClientError\(', 'cli.BuildOAuthClientError('),
    (r'\boauthClientSource\b', 'cli.OAuthClientSource'),
    (r'\boauthClientSourceFlags\b', 'cli.OAuthClientSourceFlags'),
    (r'\boauthClientSourceEnv\b', 'cli.OAuthClientSourceEnv'),
    (r'\boauthClientSourceConfig\b', 'cli.OAuthClientSourceConfig'),
    (r'\boauthClientSourceBundled\b', 'cli.OAuthClientSourceBundled'),
    (r'\bisTruthyEnv\(', 'cli.IsTruthyEnv('),
    (r'\bbuildAuthFlowError\(', 'cli.BuildAuthFlowError('),
    (r'\boauthClientSecretHint\(', 'cli.OAuthClientSecretHint('),
    (r'\bopenBrowser\(', 'cli.OpenBrowser('),
]

for pkg in packages:
    pkg_path = f'internal/cli/{pkg}'
    if not os.path.exists(pkg_path):
        continue
    
    for filename in os.listdir(pkg_path):
        if not filename.endswith('.go'):
            continue
            
        filepath = os.path.join(pkg_path, filename)
        with open(filepath, 'r') as f:
            content = f.read()
        
        # Apply replacements
        for pattern, replacement in replacements:
            content = re.sub(pattern, replacement, content)
        
        # Fix function signatures to use *cli.Globals
        content = re.sub(
            r'func \(cmd \*(\w+)\) Run\(globals \*Globals\)',
            r'func (cmd *\1) Run(globals *cli.Globals)',
            content
        )
        
        # Add cli import if needed
        if 'cli.Globals' in content or 'cli.' in content:
            if '"github.com/dl-alexandre/gdrv/internal/cli"' not in content:
                if 'import (' in content:
                    content = content.replace(
                        'import (',
                        'import (\n\tcli "github.com/dl-alexandre/gdrv/internal/cli"'
                    )
        
        with open(filepath, 'w') as f:
            f.write(content)
        
        print(f"Updated: {filepath}")

print("Done!")
