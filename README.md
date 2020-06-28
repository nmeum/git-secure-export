# git-secure-export

Experimental tooling for encrypting the `git-fast-export(1)` output.

## Description

This repository provides `git-secure-export`, a postprocessor for
`git-fast-export(1)` which encrypts file and commit message data of
exported git repositories. Additionally, `git-secure-import` is provided
which acts as a preprocessor for `git-fast-import(1)` and allows
importing a previously encrypted `git-fast-export(1)` output.

Based on these two programs, two very hacky shell scripts were written
which implement an ssh-based git remote helper (refer to
`gitremote-helpers(7)`) for encrypting repository on the remote. The
setup requires access to the remote server for installing a custom
script and thus does not work with GitHub or other hosted Git solutions.

## Status

Proof of concept, largely untested and very buggy.

## Security

The code uses [secretbox][secretbox doc] for symmetric encryption and
authentication of file contents and commit messages. The symmetric key
is stored in `.git/git-secure-key`, the file must be created explicitly
using `git-secure-init`. Encryption of file names is also being
considered but would likely require a separate deterministic encryption
scheme or some kind of local database.

## Installation

To install run:

	$ go get github.com/nmeum/git-secure-export/cmd/...

If you want to use the remote helper also install the scripts from the
`contrib/` directory. Refer to `contrib/README.md` for more information
on these scripts.

## Usage

The software requires the creation of a symmetric key this key must be
created explicitly by invoking `git secure-init` in an existing git
repository. Afterwards `git-secure-export` can be used in combination
with `git-fast-export(1)` as follows:

	$ git fast-export <options> | git secure-export | \
		git secure-import | git fast-import

Of cause it would be more meaningful to write the output of
`git-secure-export` to a file, using an output redirection, and passing
it to `git-secure-import`, using an input redirection, on a different
computer. Though this use-case would require copying the symmetric key.

## See also

Existing tooling which encrypts single files in a git repository:

* https://github.com/elasticdog/transcrypt
* https://github.com/StackExchange/blackbox
* https://github.com/AGWA/git-crypt

Existing tooling which encrypts entire git repositories:

* https://github.com/spwhitton/git-remote-gcrypt
* https://github.com/rovaughn/git-remote-grave

## License

This program is free software: you can redistribute it and/or modify it
under the terms of the GNU General Public License as published by the
Free Software Foundation, either version 3 of the License, or (at your
option) any later version.

This program is distributed in the hope that it will be useful, but
WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General
Public License for more details.

You should have received a copy of the GNU General Public License along
with this program. If not, see <http://www.gnu.org/licenses/>.

[secretbox doc]: https://pkg.go.dev/golang.org/x/crypto/nacl/secretbox
