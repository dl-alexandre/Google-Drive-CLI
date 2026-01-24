# Admin SDK Workflows

This guide covers Google Workspace Admin SDK operations using the Google Drive CLI.

## Table of Contents

- [Service Account Setup Guide](#service-account-setup-guide)
- [User Provisioning](#user-provisioning)
- [Group Management](#group-management)
- [Bulk Operations](#bulk-operations)

## Service Account Setup Guide

### Prerequisites

Admin SDK operations require **service account authentication with domain-wide delegation**. OAuth user authentication is **not supported** for admin operations.

### Step 1: Create Service Account

1. Go to [Google Cloud Console](https://console.cloud.google.com/)
2. Select your project (or create a new one)
3. Navigate to **IAM & Admin** > **Service Accounts**
4. Click **Create Service Account**
5. Enter a name and description
6. Click **Create and Continue**
7. Skip role assignment (not needed for Admin SDK)
8. Click **Done**

### Step 2: Create Service Account Key

1. Click on the newly created service account
2. Go to the **Keys** tab
3. Click **Add Key** > **Create new key**
4. Select **JSON** format
5. Click **Create** (key file downloads automatically)
6. **Store this file securely** - it contains sensitive credentials

### Step 3: Enable Domain-Wide Delegation

1. In the service account details, check **Enable Google Workspace Domain-wide Delegation**
2. Note the **Client ID** (you'll need this)

### Step 4: Authorize Scopes in Google Workspace Admin Console

1. Go to [Google Workspace Admin Console](https://admin.google.com/)
2. Navigate to **Security** > **API Controls** > **Domain-wide Delegation**
3. Click **Add new**
4. Enter the **Client ID** from Step 3
5. Add the following OAuth scopes:
   - `https://www.googleapis.com/auth/admin.directory.user`
   - `https://www.googleapis.com/auth/admin.directory.group`
   - `https://www.googleapis.com/auth/admin.directory.group.member`
6. Click **Authorize**

### Step 5: Authenticate with CLI

```bash
# Authenticate with service account
gdrive auth service-account ./path/to/service-account-key.json \
  --impersonate-user admin@yourdomain.com \
  --scopes https://www.googleapis.com/auth/admin.directory.user,https://www.googleapis.com/auth/admin.directory.group
```

**Important Notes:**
- The `--impersonate-user` must be a super admin or have appropriate admin privileges
- The service account key file must be kept secure
- Domain-wide delegation must be enabled and authorized

### Step 6: Verify Authentication

```bash
# Check auth status
gdrive auth status

# Test with a simple operation
gdrive admin users list --domain yourdomain.com --limit 1 --json
```

## User Provisioning

### List Users

List users in your domain:

```bash
# List all users in domain
gdrive admin users list --domain yourdomain.com --json

# List with query filter
gdrive admin users list --domain yourdomain.com \
  --query "isSuspended=true" \
  --json

# List with pagination
gdrive admin users list --domain yourdomain.com \
  --paginate \
  --json

# Limit results per page
gdrive admin users list --domain yourdomain.com \
  --limit 50 \
  --json
```

**Common Query Filters:**
- `isSuspended=true` - Suspended users
- `isAdmin=true` - Admin users
- `orgUnitPath=/Engineering` - Users in specific OU
- `email:admin*` - Users with email starting with "admin"

### Get User Details

Retrieve details for a specific user:

```bash
# By email
gdrive admin users get user@yourdomain.com --json

# By user ID
gdrive admin users get 12345678901234567890 --json
```

### Create User

Create a new user account:

```bash
gdrive admin users create newuser@yourdomain.com \
  --given-name "John" \
  --family-name "Doe" \
  --password "TempPass123!" \
  --json
```

**Password Requirements:**
- Minimum 8 characters
- Must contain letters and numbers
- Consider using a password generator for secure temporary passwords

### Update User

Update user properties:

```bash
# Update name
gdrive admin users update user@yourdomain.com \
  --given-name "Jane" \
  --family-name "Smith" \
  --json

# Suspend user
gdrive admin users update user@yourdomain.com \
  --suspended true \
  --json

# Move to organizational unit
gdrive admin users update user@yourdomain.com \
  --org-unit-path "/Departments/Engineering" \
  --json
```

### Suspend/Unsuspend User

```bash
# Suspend user
gdrive admin users suspend user@yourdomain.com --json

# Unsuspend user
gdrive admin users unsuspend user@yourdomain.com --json
```

### Delete User

**Warning:** This permanently deletes the user account.

```bash
gdrive admin users delete user@yourdomain.com --json
```

## Group Management

### List Groups

List groups in your domain:

```bash
# List all groups
gdrive admin groups list --domain yourdomain.com --json

# List with query
gdrive admin groups list --domain yourdomain.com \
  --query "name:Engineering*" \
  --json

# Paginate results
gdrive admin groups list --domain yourdomain.com \
  --paginate \
  --json
```

### Get Group Details

```bash
gdrive admin groups get group@yourdomain.com --json
```

### Create Group

Create a new Google Workspace group:

```bash
gdrive admin groups create engineering@yourdomain.com \
  "Engineering Team" \
  --description "All engineering team members" \
  --json
```

### Update Group

```bash
# Update group name
gdrive admin groups update group@yourdomain.com \
  --name "New Group Name" \
  --json

# Update description
gdrive admin groups update group@yourdomain.com \
  --description "Updated description" \
  --json
```

### Delete Group

**Warning:** This permanently deletes the group.

```bash
gdrive admin groups delete group@yourdomain.com --json
```

## Group Membership Management

### List Group Members

List all members of a group:

```bash
# List all members
gdrive admin members list group@yourdomain.com --json

# Filter by role
gdrive admin members list group@yourdomain.com \
  --roles OWNER \
  --json

# Paginate results
gdrive admin members list group@yourdomain.com \
  --paginate \
  --json
```

**Member Roles:**
- `OWNER` - Full control of the group
- `MANAGER` - Can manage members and settings
- `MEMBER` - Regular member

### Add Member to Group

```bash
# Add as regular member
gdrive admin members add group@yourdomain.com user@yourdomain.com \
  --role MEMBER \
  --json

# Add as manager
gdrive admin members add group@yourdomain.com user@yourdomain.com \
  --role MANAGER \
  --json

# Add as owner
gdrive admin members add group@yourdomain.com user@yourdomain.com \
  --role OWNER \
  --json
```

### Remove Member from Group

```bash
gdrive admin members remove group@yourdomain.com user@yourdomain.com --json
```

## Bulk Operations

### Bulk User Creation

Create multiple users from a CSV file:

```bash
#!/bin/bash
# CSV format: email,given_name,family_name
while IFS=, read -r EMAIL GIVEN FAMILY; do
  # Skip header row
  [[ "$EMAIL" == "email" ]] && continue
  
  # Generate temporary password
  PASSWORD=$(openssl rand -base64 12 | tr -d "=+/" | cut -c1-12)
  
  echo "Creating user: $EMAIL"
  gdrive admin users create "$EMAIL" \
    --given-name "$GIVEN" \
    --family-name "$FAMILY" \
    --password "$PASSWORD" \
    --json
  
  # Log credentials (store securely!)
  echo "$EMAIL,$PASSWORD" >> new_users.csv
  
done < users.csv
```

**Example `users.csv`:**
```csv
email,given_name,family_name
alice@yourdomain.com,Alice,Smith
bob@yourdomain.com,Bob,Jones
charlie@yourdomain.com,Charlie,Brown
```

### Bulk Group Membership

Add multiple users to a group:

```bash
#!/bin/bash
GROUP="engineering@yourdomain.com"

# Read emails from file (one per line)
while read -r EMAIL; do
  echo "Adding $EMAIL to $GROUP"
  gdrive admin members add "$GROUP" "$EMAIL" --role MEMBER --json
done < members.txt
```

**Example `members.txt`:**
```
alice@yourdomain.com
bob@yourdomain.com
charlie@yourdomain.com
```

### Bulk User Suspension

Suspend multiple users:

```bash
#!/bin/bash
# Read user emails from file
while read -r EMAIL; do
  echo "Suspending $EMAIL"
  gdrive admin users suspend "$EMAIL" --json
done < users_to_suspend.txt
```

### Bulk User Updates

Update multiple users with different properties:

```bash
#!/bin/bash
# CSV format: email,given_name,family_name,org_unit
while IFS=, read -r EMAIL GIVEN FAMILY ORG_UNIT; do
  [[ "$EMAIL" == "email" ]] && continue
  
  echo "Updating $EMAIL"
  
  # Build update command
  CMD="gdrive admin users update $EMAIL"
  
  [[ -n "$GIVEN" ]] && CMD="$CMD --given-name \"$GIVEN\""
  [[ -n "$FAMILY" ]] && CMD="$CMD --family-name \"$FAMILY\""
  [[ -n "$ORG_UNIT" ]] && CMD="$CMD --org-unit-path \"$ORG_UNIT\""
  
  CMD="$CMD --json"
  eval $CMD
  
done < user_updates.csv
```

### Automated Onboarding Workflow

Complete onboarding workflow:

```bash
#!/bin/bash
NEW_USER="newuser@yourdomain.com"
FIRST_NAME="John"
LAST_NAME="Doe"
DEPARTMENT="Engineering"

# 1. Create user account
PASSWORD=$(openssl rand -base64 12 | tr -d "=+/" | cut -c1-12)
gdrive admin users create "$NEW_USER" \
  --given-name "$FIRST_NAME" \
  --family-name "$LAST_NAME" \
  --password "$PASSWORD" \
  --json

# 2. Add to department group
gdrive admin members add "${DEPARTMENT,,}@yourdomain.com" "$NEW_USER" \
  --role MEMBER \
  --json

# 3. Add to all-employees group
gdrive admin members add "all-employees@yourdomain.com" "$NEW_USER" \
  --role MEMBER \
  --json

# 4. Move to department OU
gdrive admin users update "$NEW_USER" \
  --org-unit-path "/Departments/$DEPARTMENT" \
  --json

# 5. Send welcome email (external script)
# send_welcome_email "$NEW_USER" "$PASSWORD"

echo "User $NEW_USER onboarded successfully"
echo "Temporary password: $PASSWORD"
```

### Compliance Reporting

Generate compliance reports:

```bash
#!/bin/bash
# Generate suspended users report
gdrive admin users list --domain yourdomain.com \
  --query "isSuspended=true" \
  --paginate \
  --json > suspended_users.json

# Generate admin users report
gdrive admin users list --domain yourdomain.com \
  --query "isAdmin=true" \
  --paginate \
  --json > admin_users.json

# Generate users by OU
gdrive admin users list --domain yourdomain.com \
  --query "orgUnitPath=/Departments/Engineering" \
  --paginate \
  --json > engineering_users.json
```

### Group Sync from External Source

Sync group membership from external system:

```bash
#!/bin/bash
GROUP="engineering@yourdomain.com"

# Get current members
CURRENT_MEMBERS=$(gdrive admin members list "$GROUP" --paginate --json | \
  jq -r '.[].email' | sort)

# Get desired members from external source
DESIRED_MEMBERS=$(get_external_members | sort)

# Find users to add
TO_ADD=$(comm -13 <(echo "$CURRENT_MEMBERS") <(echo "$DESIRED_MEMBERS"))

# Find users to remove
TO_REMOVE=$(comm -23 <(echo "$CURRENT_MEMBERS") <(echo "$DESIRED_MEMBERS"))

# Add new members
for EMAIL in $TO_ADD; do
  echo "Adding $EMAIL"
  gdrive admin members add "$GROUP" "$EMAIL" --role MEMBER --json
done

# Remove old members
for EMAIL in $TO_REMOVE; do
  echo "Removing $EMAIL"
  gdrive admin members remove "$GROUP" "$EMAIL" --json
done
```

## Tips and Best Practices

1. **Security**: Never commit service account keys to version control. Use environment variables or secure secret management.

2. **Impersonation**: Always use a dedicated admin account for impersonation, not your personal account.

3. **Scope minimization**: Only request the scopes you actually need.

4. **Error handling**: Always check for errors in bulk operations and log failures.

5. **Rate limits**: Be aware of API rate limits when performing bulk operations. Add delays if needed.

6. **Testing**: Test operations on a small subset before running on all users/groups.

7. **Audit logging**: Log all admin operations for compliance and troubleshooting.

8. **Password security**: Generate strong temporary passwords and require users to change them on first login.

9. **Dry run**: Consider implementing a dry-run mode for destructive operations.

10. **Backup**: Before bulk deletions, ensure you have backups or can recover data.

## Common Error Scenarios

### Authentication Errors

**Error:** "Admin SDK requires service account authentication"

**Solution:** Ensure you're using service account authentication, not OAuth user flow.

### Permission Errors

**Error:** "Insufficient permissions"

**Solution:** 
- Verify domain-wide delegation is enabled
- Check that scopes are authorized in Admin Console
- Ensure impersonated user has admin privileges

### User Already Exists

**Error:** "User already exists"

**Solution:** Check if user exists first, or use update instead of create.

## See Also

- [Google Workspace Admin SDK Documentation](https://developers.google.com/admin-sdk/directory)
- [Domain-Wide Delegation Guide](https://developers.google.com/identity/protocols/oauth2/service-account#delegatingauthority)
- Main CLI [README](../../README.md)
