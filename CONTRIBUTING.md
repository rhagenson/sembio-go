Contributing
============

It's good to hear that you want to contribute to `sembio/go`!

There are a number of ways to contribute to `sembio/go`, but currently all contributions are triaged through [Issues](https://github.com/sembio/go/issues). (This is likely to change as more people require more from `sembio/go`)

- [Contributing](#contributing)
  - [Feature request](#feature-request)
  - [Bug report](#bug-report)
  - [How to contribute](#how-to-contribute)
  - [Pull request](#pull-request)
  - [Documentation Formatting](#documentation-formatting)
  - [Code Formatting](#code-formatting)

Feature request
---------------
For any feature requests or enhancements to `sembio/go`, please open an issue and label it as an `enhancement`.

If you submit a pull request to implement a new feature before a corresponding issue exists you will be asked to open an issue describing the proposed enchancement first.

Bug report
----------
First of all please [search existing issues](https://github.com/sembio/go/issues) to make sure the bug has not been reported before. If you cannot find a suitable issue — [create a new one](https://github.com/sembio/go/issues/new).

Provide the following details:

- short summary of what you was trying to achieve,
- a code causing the bug,
- expected result,
- actual results and
- environment details: at least operating system and library version

If possible, try to isolate the problem and provide just enough code to demonstrate it. Add any related information which might help to fix the issue.

How to contribute
-----------------
We use a fairly standard GitHub pull request workflow. If you have already contributed to a project via GitHub pull request, you can skip this section and proceed to the [specific details of what we ask for in a pull request](#pull-request). If this is your first time contributing to a project via GitHub, read on.

Here is the basic GitHub workflow:

1. Fork the `sembio/go` repo. you can do this via the GitHub website. This will result in you having your own copy of the `sembio/go` repo under your GitHub account.
2. Clone your `sembio/go` repo to your local machine
3. Make a branch for your change
4. Make your change on that branch
5. Push your change to your repo
6. Use the GitHub UI to open a PR

Some things to note that aren't immediately obvious to folks just starting out:

1. Your fork doesn't automatically stay up to date with change in the main repo.
2. Any changes you make on your branch that you used for the PR will automatically appear in the PR so if you have more than 1 PR, be sure to always create different branches for them.
3. Weird things happen with commit history if you don't create your PR branches off of master so always make sure you have the master branch checked out before creating a branch for a PR.

You can get help using GitHub via [the official documentation](https://help.github.com/). Some highlights include:

- [Fork A Repo](https://help.github.com/articles/fork-a-repo/)
- [Creating a pull request](https://help.github.com/articles/creating-a-pull-request/)
- [Syncing a fork](https://help.github.com/articles/syncing-a-fork/)

Pull request
------------
Before issuing a pull request we ask that you squash all your commits into a single logical commit. While your PR is in review, we may ask for additional changes, please do not squash those commits while the review is underway. Once everything is good, we'll then ask you to further squash those commits before merging. We ask that you not squash while a review is underway as it can make it hard to follow what is going on. Additionally, we ask that you:

- [Write a good commit message](http://chris.beams.io/posts/git-commit/)
- Issue 1 Pull Request per feature. Don't lump unrelated changes together.

If you aren't sure how to squash multiple commits into one, Steve Klabnik wrote [a handy guide](http://blog.steveklabnik.com/posts/2012-11-08-how-to-squash-commits-in-a-github-pull-request) that you can refer to.

Ensure that the ticket title is appropriate - the title will be used as the summary of the change, so it should be appropriately formatted, including a ticket reference if the PR is a fix to an existing bug ticket.
For example, an appropriate title for a PR that fixes a bug reported in issue ticket #98 might look like: *Fixed compiler crash related to tuple recovery (issue #98)*

Documentation Formatting
---------------
When contributing to documentation, try to keep the following style guidelines in mind:

- Wherever possible, try to match the style of surrounding documentation.
- Use semantic line wrapping (line break after each sentence) combined with our editors soft-wrapping. Avoid hard-wrapping lines within paragraphs (using line breaks in the middle of or between sentences to make lines shorter than a certain length).
- Apply code highlighting to identifier names, such as types, fields, variables and parameters via the customary markdown syntax of wrapping the name in backticks.  

Code Formatting
---------------
For code formatting, please use `gofmt`.
