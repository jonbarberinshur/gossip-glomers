# gossip-glomers
Skeleton project for the [fly.io distributed systems challenges](https://fly.io/dist-sys/)

## Setting up
We use devbox to make sure the correct tools are available

1. Use the [Determinate Nix installer](https://determinate.systems/posts/determinate-nix-installer/) to install nix
2. Install [devbox](https://www.jetpack.io/devbox/docs/installing_devbox/)
3. *Optional* install [direnv](https://direnv.net/docs/installation.html) so that cding into this directory will set up your shell
4. Clone this repository 
5. Check everything is working so far. In a new shell in the root of this project run 
```shell
devbox run help
```
You should see a nicely formatted version of this doc. If not, pls ask for help

Note that the first time you run devbox it can take some time to fetch tooling.

6. Download maelstrom. In a new shell cd into this directory and run
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

## Hints & tips

If you want to print diagnostics or log lines you must write to stdout - the standard `log` package does this by default. More details in the [protocol](https://github.com/jepsen-io/maelstrom/blob/main/doc/protocol.md#protocol) docs

You can write your own message types & handlers for the nodes to communicate with each other, in addition to the ones that maelstrom sends.

The `sync` package provides some useful synchronization primitives - good summary by [Teiva Harsanyi](https://teivah.medium.com/a-closer-look-at-go-sync-package-9f4e4a28c35a)

There are always 2 setup messages sent by maelstrom ahead of the main workload, and you can add handlers for them if you wish. In order, they are

* `init` has [this body](https://github.com/jepsen-io/maelstrom/blob/main/demo/go/node.go#L332-L336) - more [details](https://github.com/jepsen-io/maelstrom/blob/main/doc/protocol.md#initialization)
* `topology` - more [details](https://github.com/jepsen-io/maelstrom/blob/main/doc/workloads.md#rpc-topology) 

You don't have to use go for this - there are [example node implementations](https://github.com/jepsen-io/maelstrom/tree/main/demo) in different languages. (The python one has an implementation for [raft](https://github.com/jepsen-io/maelstrom/blob/main/demo/python/raft.py)).

The [maelstrom project docs](https://github.com/jepsen-io/maelstrom?tab=readme-ov-file#documentation) are really handy - there are walkthroughs of some of the challenges and descriptions of supporting services. Definitely read the page on interpreting [results](https://github.com/jepsen-io/maelstrom/blob/main/doc/results.md) as a start
