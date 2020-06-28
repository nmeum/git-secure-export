# git-secure-export contrib

Scripts for using `git-secure-export` a git remote-helper.

## Architecture

The client-side git remote helper (`git-remote-secure`) connects using
`ssh(1)` to the git server and spawns the `git-secure-receive` command
on this server. Afterwards, remote helper commands are piped to it. Data
send to the server is encrypted using `git-secure-export`, data received
from the server is decrypted using `git-secure-import`.

## Installation

This setup requires install one script on the server and the client.

### Server

Copy `git-secure-receive` to your `$PATH`, make sure you can invoke it
from `ssh(1)` as `ssh <host> git-secure-receive`.

### Client

Copy `git-remote-secure` to your `$PATH`. Start using
`git-remote-secure` for your repository by adding a secure remote using:

	git remote add secure://<git server>:<path to git repo on server>

Afterwards, you can push/pull your repository as you would normally
while data is transparently encrypted/decrypted by the remote helper.
You can verify that the plaintext is not stored on the server by cloning
your encrypted repository over ssh without using `git-remote-secure`.
