# Medcare

Medcare is a revolutionary solution aimed at modernizing hospital operations and improving patient care delivery, It is an innovative online-based platform designed to streamline healthcare services by connecting patients and doctors seamlessly

## How does it work?
### Patient-Doctor Interaction:
<h4>Appointment Scheduling:</h4>h4> Patients can log in to the UHMS portal and schedule appointments with their preferred doctors at their convenience. They can view available time slots, select a suitable appointment time, and receive confirmation of their booking.<br><br>
Remote Consultations: Through the platform, patients can connect with their doctors via online meetings. These remote consultations allow for real-time communication, enabling patients to discuss their health concerns, receive medical advice, and seek prescriptions from the comfort of their homes.<br>
### Doctor Verification:
Authentication Process: UHMS implements a doctor verification mechanism to ensure that only legitimate healthcare professionals join the platform. Doctors are required to undergo a verification process where their credentials and qualifications are authenticated before they are approved to interact with patients.<br><br>
Access to Medical Records:
Secure Access: Patients have access to their medical records and reports through the UHMS portal. They can securely view their past medical history, test results, prescriptions, and other relevant information, empowering them to take control of their healthcare journey.
### Prescription Management:
Digital Prescriptions: Doctors can generate digital prescriptions for their patients directly through the UHMS platform. They can prescribe medications, dosage instructions, and other relevant details, which are then accessible to patients for fulfillment at pharmacies.
### Seamless Communication:
Efficient Communication Channels: UHMS provides efficient communication channels for patients and doctors to exchange information, ask questions, and address concerns. Whether it's through text-based messaging or video consultations, the platform ensures that communication is smooth and effective.
### Online-Based Platform:
Accessibility: 
Being an online-based platform, UHMS offers accessibility to patients and doctors from anywhere with an internet connection. This eliminates geographical barriers, allowing individuals to receive quality healthcare services regardless of their location.<br>
### Centralized Management:
Centralized Database: UHMS centralizes patient records, appointment schedules, and communication logs, making it easier for healthcare providers to manage and track patient interactions. This centralized approach enhances efficiency and reduces the likelihood of errors or miscommunications.

## How do GB's virtual branches differ from Git branches?

The branches that we know and love in Git are separate universes, and switching between them is a full context switch. GitButler allows you to work with multiple branches in parallel in the same working directory. This effectively means having the content of multiple branches available at the same time.

GitButler is aware of changes before they are committed. This allows it to keep a record of which virtual branch each individual diff belongs to. Effectively, this means that you can separate out individual branches with their content at any time to push them to a remote or to unapply them from your working directory.

And finally, while in Git it is preferable that you create your desired branch ahead of time, using GitButler you can move changes between virtual branches at any point during development.

## Why GitButler?

We love Git. Our own [@schacon](https://github.com/schacon) has even published the [Pro Git](https://git-scm.com/book/en/v2) book. At the same time, Git's user interface hasn't been fundamentally changed for 15 years. While it was written for Linux kernel devs sending patches to each other over mailing lists, most developers today have different workflows and needs.

Instead of trying to fit the semantics of the Git CLI into a graphical interface, we are starting with the developer workflow and mapping it back to Git.

## Tech

GitButler is a [Tauri](https://tauri.app/)-based application. Its UI is written in [Svelte](https://svelte.dev/) using [TypeScript](https://www.typescriptlang.org) and its backend is written in [Rust](https://www.rust-lang.org/).

## Main Features

- **Virtual Branches**
  - Organize work on multiple branches simultaneously, rather than constantly switching branches
  - Automatically create new branches when needed
- **Easy Commit Management**
  - Undo, Amend and Squash commits by dragging and dropping
- **GitHub Integration**
  - Authenticate to GitHub to open Pull Requests, list branches and statuses and more
- **Easy SSH Key Management**
  - GitButler can generate an SSH key to upload to GitHub automatically
- **AI Tooling**
  - Automatically write commit messages based on your work in progress
  - Automatically create descriptive branch names
- **Commit Signing**
  - Easy commit signing with our generated SSH key

## Example uses

### Fixing a bug while working on a feature

> Say that while developing a feature, you encounter a bug that you wish to fix. It's often desirable that you ship the fix as a separate contribution (Pull request).

Using Git you can stash your changes and switch to another branch, where you can commit, and push your fix.

*With GitButler* you simply assign your fix to a separate virtual branch, which you can individually push (or directly create a PR). An additional benefit is that you can retain the fix in your working directory while waiting for CI and/or code review.

### Trying someone else's branch together with my work in progress

> Say you want to test a branch from someone else for the purpose of code review.

Using Git trying out someone else's branch is a full context switch away from your own work.
*With GitButler* you can apply and unapply (add / remove) any remote branch directly into your working directory.

## Documentation

You can find our end user documentation at: https://docs.gitbutler.com

## Bugs and Feature Requests

If you have a bug or feature request, feel free to open an [issue](https://github.com/gitbutlerapp/gitbutler/issues/new),
or [join our Discord server](https://discord.gg/wDKZCPEjXC).

## AI Commit Message generation

Commit message generation is an opt-in feature. You can enable it while adding your repository for the first time or later in the project settings.

Currently GitButler uses OpenAI's API for diff summarization, which means that if enabled, code diffs would be sent to OpenAI's servers.

Our goal is to make this feature more modular such that in the future you can modify the prompt as well as plug a different LLM endpoints (including a local ones).

## Contributing

So you want to help out? Please check out the [CONTRIBUTING.md](CONTRIBUTING.md)
document.

If you want to skip right to getting the code to actually compile, take a look
at the [DEVELOPMENT.md](DEVELOPMENT.md) file.

Want to show your support? Add a GitButler badge to your project's README:
```md
[![GitButler](https://img.shields.io/badge/GitButler-%23B9F4F2?logo=data%3Aimage%2Fsvg%2Bxml%3Bbase64%2CPHN2ZyB3aWR0aD0iMzkiIGhlaWdodD0iMjgiIHZpZXdCb3g9IjAgMCAzOSAyOCIgZmlsbD0ibm9uZSIgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIj4KPHBhdGggZD0iTTI1LjIxNDUgMTIuMTk5N0wyLjg3MTA3IDEuMzg5MTJDMS41NDI5NSAwLjc0NjUzMiAwIDEuNzE0MDYgMCAzLjE4OTQ3VjI0LjgxMDVDMCAyNi4yODU5IDEuNTQyOTUgMjcuMjUzNSAyLjg3MTA3IDI2LjYxMDlMMjUuMjE0NSAxNS44MDAzQzI2LjcxOTcgMTUuMDcyMSAyNi43MTk3IDEyLjkyNzkgMjUuMjE0NSAxMi4xOTk3WiIgZmlsbD0iYmxhY2siLz4KPHBhdGggZD0iTTEzLjc4NTUgMTIuMTk5N0wzNi4xMjg5IDEuMzg5MTJDMzcuNDU3MSAwLjc0NjUzMiAzOSAxLjcxNDA2IDM5IDMuMTg5NDdWMjQuODEwNUMzOSAyNi4yODU5IDM3LjQ1NzEgMjcuMjUzNSAzNi4xMjg5IDI2LjYxMDlMMTMuNzg1NSAxNS44MDAzQzEyLjI4MDMgMTUuMDcyMSAxMi4yODAzIDEyLjkyNzkgMTMuNzg1NSAxMi4xOTk3WiIgZmlsbD0idXJsKCNwYWludDBfcmFkaWFsXzMxMF8xMjkpIi8%2BCjxkZWZzPgo8cmFkaWFsR3JhZGllbnQgaWQ9InBhaW50MF9yYWRpYWxfMzEwXzEyOSIgY3g9IjAiIGN5PSIwIiByPSIxIiBncmFkaWVudFVuaXRzPSJ1c2VyU3BhY2VPblVzZSIgZ3JhZGllbnRUcmFuc2Zvcm09InRyYW5zbGF0ZSgxNi41NzAxIDE0KSBzY2FsZSgxOS44NjQxIDE5LjgzODMpIj4KPHN0b3Agb2Zmc2V0PSIwLjMwMTA1NiIgc3RvcC1vcGFjaXR5PSIwIi8%2BCjxzdG9wIG9mZnNldD0iMSIvPgo8L3JhZGlhbEdyYWRpZW50Pgo8L2RlZnM%2BCjwvc3ZnPgo%3D
)](https://gitbutler.com/)
```
[![BADGE][s6]][l6]
