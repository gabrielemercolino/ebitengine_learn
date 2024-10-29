{
  description = "template dev flake";
  inputs.nixpkgs.url = "github:NixOS/nixpkgs/nixpkgs-unstable";
  inputs.flake-utils.url = "github:numtide/flake-utils";

  outputs = { self, nixpkgs, flake-utils, ... }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs { inherit system; };
        myGoApp = pkgs.callPackage ./package.nix { };
      in {
        devShells.default = pkgs.mkShell {
          nativeBuildInputs = with pkgs; [
            go
            libGL
            xorg.libX11
            xorg.libXrandr
            xorg.libXcursor
            xorg.libXinerama
            xorg.libXi
            xorg.libXxf86vm
            pkg-config
          ];
          shellHook = ''
            export LD_LIBRARY_PATH=${pkgs.wayland}/lib:${
              pkgs.lib.getLib pkgs.libGL
            }/lib:$LD_LIBRARY_PATH
            echo "Welcome!" | ${pkgs.cowsay}/bin/cowsay | ${pkgs.lolcat}/bin/lolcat
          '';
        };

        packages.abstract = myGoApp;
        apps.${system}.abstract = {
          type = "app";
          package = self.packages.${system}.abstract;
        };
      });
}

