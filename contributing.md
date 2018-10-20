# CONTRIBUTING

go-lcss project is [Apache licensed](license.md) and accepts contributions via GitHub pull requests.
This document outlines some of the conventions on development workflow, commit message formatting,
contact points, and other resources to make it easier to get your contribution accepted.

## Support Channels

The official support channels, for both users and contributors, are:

* GitHub [issues](https://github.com/vmarkovtsev/go-lcss/issues)\*
* E-mail: see `git log` for Vadim's email address.

\*Before opening a new issue or submitting a new pull request, it's helpful to search the project -
it's likely that another user has already reported the issue you're facing, or it's a known issue
that we're already aware of.

## How to Contribute

Pull Requests \(PRs\) are the main and exclusive way to contribute to the official go-lcss project.
In order for a PR to be accepted it needs to pass a list of requirements:

* Code coverage does not decrease.
* All the current tests pass.
* `go vet` and `golint` stay silent.
* * If the PR is a bug fix, it has to include a new unit test that fails before the patch is merged.
* If the PR is a new feature, it has to come with a suite of unit tests, that tests the new functionality.

### Format of the commit message

The commit summary must start with a capital letter and with a verb in present tense. No dot in the end.

```text
Add a feature
Remove unused code
Fix a bug
```

Every commit details should describe what was changed, under which context and, if applicable,
the GitHub issue it relates to.
