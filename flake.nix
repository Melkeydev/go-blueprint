{
  description = "Go-blueprint allows users to spin up a quick Go project using a popular framework";

  inputs = {
    nixvim.url = "github:nix-community/nixvim";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = {
    nixpkgs,
    flake-utils,
    ...
  }:
    flake-utils.lib.eachDefaultSystem (system: let
      pkgs = import nixpkgs {inherit system;};
    in {
      packages = {
        default = let
          module = pkgs.buildGoModule {
            name = "go-blueprint";
            # version = "0.5.14";

            src = ./.;

            # src = pkgs.fetchFromGitHub {
            #   owner = "Melkeydev";
            #   repo = "go-blueprint";
            #   rev = "v${version}";
            #   hash = "sha256-vQ7LXC70MpqdHxR4JpM5iRQ4mTO0MVtZF94G/f8sR6A=";
            # };

            vendorHash = "sha256-WBzToupC1/O70OYHbKk7S73OEe7XRLAAbY5NoLL7xvw=";

            meta = with pkgs.lib; {
              description = "Go-blueprint allows users to spin up a quick Go project using a popular framework";
              homepage = "https://github.com/Melkeydev/go-blueprint";
              licence = licenses.mit;
            };
          };
        in
          pkgs.symlinkJoin rec {
            name = "go-blueprint";
            # Go needed in PATH to generate go.mod
            paths = [module] ++ [pkgs.go];
            buildInputs = [pkgs.makeWrapper];
            postBuild = "wrapProgram $out/bin/${name} --prefix PATH : $out/bin";
          };
      };
    });
}
