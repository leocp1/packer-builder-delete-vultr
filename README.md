# `read-vultr` and `delete-vultr` Builders

Type: `read-vultr`

The `read-vultr` builder accepts the same configuration as the
[`vultr`](https://github.com/vultr/packer-builder-vultr) builder. It finds all
images that could have been built with this configuration, then returns an
[Artifact](https://godoc.org/github.com/hashicorp/packer/packer#Artifact) with
the `Id()` of the most recent such image if it exists, and a `String()` equal to
the number of images found.

Type: `delete-vultr`

Same as `read-vultr`, but it also deletes the images found.

## License

All Go source files (files with extension `.go`) in this repository are licensed
under the Mozilla Public License 2.0.
[![License: MPL 2.0](https://img.shields.io/badge/License-MPL%202.0-brightgreen.svg)](https://opensource.org/licenses/MPL-2.0)

All Nix files (files with extension `.nix`) in this repository are licensed
under the MIT License.
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
