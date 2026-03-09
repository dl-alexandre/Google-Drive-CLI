#!/bin/bash

# Fix imports in drive package files
for f in internal/cli/drive/*.go; do
    # Add cli import after the package declaration and import opening
    sed -i '' '/^import (/a\
\tcli "github.com/dl-alexandre/gdrv/internal/cli"' "$f"
done

# Fix references in drive package - replace Globals with cli.Globals, etc.
for f in internal/cli/drive/*.go; do
    # Skip test files for now
    if [[ "$f" == *_test.go ]]; then continue; fi
    
    # Replace function calls that need cli. prefix
    sed -i '' 's/func (cmd \*\w*Cmd) Run(globals \*Globals)/func (cmd *\1) Run(globals *cli.Globals)/g' "$f"
    sed -i '' 's/globals\.ToGlobalFlags()/globals.ToGlobalFlags()/g' "$f"
    sed -i '' 's/NewOutputWriter/cli.NewOutputWriter/g' "$f"
    sed -i '' 's/ResolveFileID/cli.ResolveFileID/g' "$f"
    sed -i '' 's/GetGlobalFlags/cli.GetGlobalFlags/g' "$f"
    sed -i '' 's/GetLogger/cli.GetLogger/g' "$f"
    sed -i '' 's/GetPathResolver/cli.GetPathResolver/g' "$f"
    sed -i '' 's/GetResolveOptions/cli.GetResolveOptions/g' "$f"
    sed -i '' 's/handleCLIError/cli.HandleCLIError/g' "$f"
    sed -i '' 's/convertDriveFile/cli.ConvertDriveFile/g' "$f"
done

echo "Drive imports updated"
