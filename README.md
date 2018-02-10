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

## Footnotes

I never figured out exactly how feross was grabbing patreon/OC links for the "where to donate" - I looked on quite a few Go packages and noticed that none of them even took donations, only the larger projects that were using GitCoin, OpenCollective, Linux Foundation, etc.

PR's welcome though!
