# thanks

feross made a [really cool thing](https://github.com/feross/thanks/)! So I stole the idea for the Go community!

It's not as fancy looking, I wrote it quickly on Saturday morning.

## Usage

```bash
$ go get github.com/Southclaws/thanks
```

```bash
$ cd some/project/that/uses/dep
$ thanks
You depend on:
- https://github.com/Masterminds/semver
- https://github.com/docker/docker
Go buy em a beer!
```

## Adding Your Donation Link

Submit a PR to the `lists.go` page, add your Go package import path followed by whatever donation page you use (Patreon, OpenCollective, etc.)
