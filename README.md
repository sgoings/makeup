makeup
======

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
  makeup add github.com/sgoings/makeup-kit-info
  ```

Demo
----

[![asciicast](https://asciinema.org/a/31535.png)](https://asciinema.org/a/31535)

Why the name?
-------------

makeup helps you create prettier makefiles
