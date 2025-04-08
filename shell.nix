{ pkgs ? import (builtins.fetchTarball {
    name = "nixpkgs-stable-24.11";
    url = "https://github.com/NixOS/nixpkgs/archive/refs/tags/24.11.tar.gz";
    sha256 = "1gx0hihb7kcddv5h0k7dysp2xhf1ny0aalxhjbpj2lmvj7h9g80a";
  }) {}
}:
pkgs.mkShell {
  name = "go-env";

  buildInputs = [
    pkgs.git
    pkgs.go_1_23
    pkgs.golangci-lint
  ];
}
