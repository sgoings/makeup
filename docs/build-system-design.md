Existing Problems
====

- Makefile logic duplicated across all deis components
  - including new helm related projects
- no sense of consistency in expected "lifecycle" tasks
  - build
  - test
  - push vs publish vs release
  - deploy
- prerequisites have varying quality of setup
  - kubectl
  - docker-machine
  - helm
- there's generally an opinion that we have a "higher-than-we-should" barrier to entry

Leading Questions
====

1. Wouldn't it be nice if you didn't have to figure out how to setup a local docker registry on each project in the Makefile?

Visionquest
====

### Step 1: "Everyone" wants a containerized dev environment

1. Take Kent's `docker-go-dev` project to new heights
2. Enable reduced maintenance by pulling in a product that already exists in this space

##### wercker

"Containerized build and deployment pipelines at your fingertips"
- same yaml you use locally on your laptop is executed in CI

Design:
  - `wercker dev` - your quick dev/test local workflow (subset of build)
  - `wercker build` - what you run when you want to create a deployable
  - `wercker deploy` - how you send your docker image out into the wild

Show (wercker.yml) and Demo Wercker

Things I Liked:
  - tools with minimal user interface (build, deploy, dev targets)
    - people tend to use flexible lifecycle build tools like children engage in fingerpainting
  - everything! in containers
    - with "internal" steps built by wercker to enable you to write steps that interact with the docker engine + workflow
  - "steps" as plugins (url + tar.gz)

Things I Didn't Like:
  - wercker cli (and other related components) are closed source :-(
  - half baked plugin infrastructure (you can include steps via url + tar.gz)
  - ability to run subset of build cycle (glide install tends to make the 0 -> build timeline a bit longer than comfortable)
    - it'd be neat if there was a way to see if the steps are up to date and then skip them smartly

### Step 2: We need a system that gives us the niceties of wercker but with more flexibility

If only wercker was open source... :-(

#### Gradle

- Amazing plugin system (move logic from build file to buildSrc, move buildSrc to separate project, publish project as a jar, include that jar via)

```
plugins {
  id "engineyard:go-conventions"
  id "engineyard:docker-conventions"
}
```

or

```
plugins {
  id "engineyard:deis-component"
}
```

Show Gradle build file and experience

If only Gradle wasn't Java...

### Step 3: We need a system that gives us the niceties of wercker, extensibility of Gradle (or Rake, etc. etc), but without the requirement of Java

Why are makefile plugins _still_ not a thing!?

What if we could have Makefiles for Deis components that looked like:

```
include makeup.mk
include $(MAKEUP_DIR)/conventions/deis-component/main.mk
```

Which would then import:
  - targets
  - reusable functions
  - environment vars/Make vars

Essentially what we're building here is a way to share common makefile logic using tools we already are using.

Let me show you something...

[makeup](https://github.com/sgoings/makeup) - when you want to have beautiful makefiles

- Show makeup backends (main.mk)
- Show what it looks like integrated with a project
- Show what it looks like to bootstrap a project

### My End Game

- get some really good + usable conventions for the Deis + Helm project:
  - building a simple go project with glide and our supporting tools (lint, gofmt, etc)
  - publishing this to bintray
  - publishing Docker images to multiple sources
  - containerized build environment options for ^^ (so you can choose if you want, but you're not forced to)
- these conventions + plugins help us transform our prototype repo from a "model" to a "foundation"
- have this catch on as a tool that the Go community is drawn to as a way to handle nontrivial project builds
  - compelling: git + make = only prerequisites
    - bootstrapping + git submodule/repo tool could (and should) be written in Go

### Potential Questions

- how do we inject things into makefiles before some of the auto-created targets (we can create lifecycle hooks)
- how can someone start building their own plugins? We're not there yet, but my idea is that the makeup utility could handle the versioning for you + multiple git submodule management for you... or we might find that we should be building a solution like helm offers
-
