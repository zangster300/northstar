
{
  description = "Northstar - Go web application development environment";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = nixpkgs.legacyPackages.${system};
      in
      {
        devShells.default = pkgs.mkShell {
          buildInputs = with pkgs; [
            # Go toolchain
            go_1_25
            # Task runner (optionally use nix package instead of `go tool task`)
            go-task
            # Tailwind CSS (standalone CLI)
            tailwindcss_4
            # For any custom JS/TS libraries needed to run
            pnpm
            # Git (usually available but explicit is good)
            git
            # Watcher dependency
            watchman
            # Optional: useful Go tools
            gopls
            gotools
            go-tools
            delve
          ];

          shellHook = ''
            echo "🚀 Northstar dev environment loaded."
            echo ""
            echo "Available tools:"
            echo "  - Go $(go version | cut -d' ' -f3)"
            echo "  - Task $(task --version)"
            echo "  - Tailwind CSS $(tailwindcss --help | head -n1 || echo 'installed')"
            echo "  - pnpm $(pnpm --version)"
            echo ""
            echo "Run 'task live' to start the development server"
            echo ""

            # Set up Go environment
            export GOPATH="$HOME/go"
            export PATH="$GOPATH/bin:$PATH"
            
            # Prefer static binaries (avoid dynamically linked binaries that fail on NixOS)
            export PREFER_STATIC_BINARIES=true
            export TAILWIND_BIN=tailwindcss
          '';
          };
        }
      );
}
