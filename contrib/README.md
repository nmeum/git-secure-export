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

## Usage

Example initialization of a new repository on both client and server:

	$ mkdir testrepo && cd testrepo
	$ git init
	Initialized empty Git repository in /tmp/testrepo
	$ git secure-init
	Initialized symmetric key in .git/git-secure-key
	$ ssh example.org 'mkdir -p testrepo && git -C testrepo init'
	Initialized empty Git repository in /home/user/testrepo/.git/
	$ git remote add secure://example.org:repos/testrepo

Afterwards files can be committed and pushed as usual.
