makeup
======

[![Build Status](https://travis-ci.org/sgoings/makeup.svg?branch=master)](https://travis-ci.org/sgoings/makeup)

Quickstart
----------

1. Clone this repository

  ```
  git clone https://github.com/sgoings/makeup.git
  ```

2. Install `makeup` into your PATH

  ```
  make install
  ```

3. Change to the directory that holds the project you want to add `makeup` into

  ```
  cd <project directory> # this project must already be using git for version control
  ```

4. Run the init command

  ```
  makeup init
  ```

5. Add makeup kits using the `add` command (uses similar syntax as `go get`)

  ```
  makeup add github.com/sgoings/makeup-bag-deis
  ```
  
6. In your Makefile, you'll see:

  ```
  # makeup-managed:begin
  include makeup.mk
  # makeup-managed:end
  ```
  
  Add the following below that section:
  
  ```
  include .makeup/makeup-bag-deis/info.mk
  ```

7. Run `make info`!

Demo
----

[![asciicast](https://asciinema.org/a/31535.png)](https://asciinema.org/a/31535)

Why the name?
-------------

makeup helps you create prettier makefiles
