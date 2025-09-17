# GitHub Setup

## Push to GitHub

```bash
# Add your GitHub remote (replace with your repo URL)
git remote add origin https://github.com/your-username/glui.git

# Push to GitHub
git push -u origin main
```

## GitHub Secrets (Optional)

For integration tests to run on main branch pushes, add this secret in GitHub:

1. Go to your repo → Settings → Secrets and variables → Actions
2. Add repository secret:
   - **Name**: `GITLAB_TOKEN`
   - **Value**: Your GitLab personal access token

**Note**: Integration tests will be skipped if the secret isn't set - this is safe.

## CI/CD Features

The workflow will automatically:
- ✅ Run unit tests
- ✅ Run E2E tests  
- ✅ Generate test coverage
- ✅ Security audit
- ✅ Code linting
- ✅ Integration tests (if token provided)
- ✅ Upload coverage to Codecov (optional)

## Status Badges

Add to your README.md:

```markdown
[![CI](https://github.com/your-username/glui/workflows/CI/badge.svg)](https://github.com/your-username/glui/actions)
[![codecov](https://codecov.io/gh/your-username/glui/branch/main/graph/badge.svg)](https://codecov.io/gh/your-username/glui)
```
