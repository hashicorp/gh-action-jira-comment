# gh-action-jira-comment

Use [GitHub actions](https://docs.github.com/en/actions) to comment on Jira issues.

## Authentication

To provide a URL and credentials you can use the [`gajira-login`](https://github.com/atlassian/gajira-login) action, which will write a config file this action can read.
Alternatively, you can set some environment variables:

- `JIRA_BASE_URL` - e.g. `https://my-org.atlassian.net`. The URL for your Jira instance.
- `JIRA_API_TOKEN` - e.g. `iaJGSyaXqn95kqYvq3rcEGu884TCbMkU`. An access token.
- `JIRA_USER_EMAIL` - e.g. `user@example.com`. The email address for the access token.

## Inputs

- `issue` (required) - The issue key to comment on, e.g. `'TEST-23'`
- `comment` (required) - The comment to make, e.g. `'This one's important'`

## Outputs

None.

## Examples

The following examples are valid `steps` for a particular job in a workflow; to see how to integrate them into a fully worked example, refer to the [documentation](https://docs.github.com/en/actions/configuring-and-managing-workflows/configuring-a-workflow).

Using `atlassian/gajira-login` and [GitHub secrets](https://docs.github.com/en/actions/configuring-and-managing-workflows/creating-and-storing-encrypted-secrets) for authentication:

```yaml
- name: Login
  uses: atlassian/gajira-login@v2.0.0
  env:
    JIRA_BASE_URL: ${{ secrets.JIRA_BASE_URL }}
    JIRA_USER_EMAIL: ${{ secrets.JIRA_USER_EMAIL }}
    JIRA_API_TOKEN: ${{ secrets.JIRA_API_TOKEN }}

- name: Comment
  uses: tomhjp/gh-action-jira-comment@v0.1.3
  with:
    issue: TEST-23
    comment: "This is an automated comment"
```
