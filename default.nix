{
  pkgs,
  lib,
}: let
  go-blueprint = pkgs.buildGoModule {
    pname = "go-blueprint";
    version = "0.5.14";
    src = ./.;

    # Needs to be updated each time dependencies change
    # Use lib.fakeHash and run 'nix build .#' to obtain the new hash
    # vendorHash = lib.fakeHash; # uncomment this and comment the definition below out to obtain the new hash
    vendorHash = "sha256-WBzToupC1/O70OYHbKk7S73OEe7XRLAAbY5NoLL7xvw=";

    meta = with lib; {
      description = "The ultimate golang blueprint library";
      homepage = "https://github.com/Melkeydev/go-blueprint";
      licence = licenses.mit;
    };
  };
in
  # Output a shell script that runs go-blueprint with go as a runtime dependency
  pkgs.writeShellApplication {
    name = "go-blueprint";
    runtimeInputs = with pkgs; [go];
    text = ''
      "${go-blueprint}/bin/go-blueprint" "$@"
    '';
  }
