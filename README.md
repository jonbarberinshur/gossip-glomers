# gossip-glomers
Skeleton project for the [fly.io distributed systems challenges](https://fly.io/dist-sys/)

## Setting up
We use devbox to make sure the correct tools are available

1. Use the [Determinate Nix installer](https://determinate.systems/posts/determinate-nix-installer/) to install nix
2. Install [devbox](https://www.jetpack.io/devbox/docs/installing_devbox/)
3. *Optional* install [direnv](https://direnv.net/docs/installation.html) so that cding into this directory will set up your shell
4. Check everything is working so far. In a new shell run
```shell
devbox run help
```
You should see a nicely formatted version of this doc. If not, pls ask for help

5. Download maelstrom. In a new shell cd into this directory and run
```shell
devbox run install
```

## Tackling the challenges.
One way to tackle the challenges is to clone the way the first [echo](https://fly.io/dist-sys/1/) challenge has been done for you.

To see this in action run
```shell
devbox run echo
```
You should see maelstrom run the test against the code in the `echo` directory.

To follow this approach, clone the `echo` script entry in [devbox.json](devbox.json) and modify to point to your new directory and change the parameters accordingly.

This has already been done for the [unique id generation](https://fly.io/dist-sys/2/) challenge, but you have to write the code :-)

## Interesting things

The [maelstrom project docs](https://github.com/jepsen-io/maelstrom?tab=readme-ov-file#documentation) are really handy - there are walkthroughs of some of the challenges and descriptions of supporting services
