 let
   npins = import ./npins;
   system = builtins.currentSystem;
   pkgs = import npins.nixpkgs { inherit system; config = {}; overlays = []; };
   build = pkgs.callPackage ./build.nix {};
   HOME = builtins.getEnv "HOME";
   PROJECT_ROOT = builtins.toString ./.;
 in
 {
   build = build;
   shell = pkgs.mkShell {
     inputsFrom = [ build ];
     packages = [
      
        pkgs.git
      
        pkgs.nixpkgs-fmt
      
     ];
     
     hardeningDisable = [ "fortify" ];
     
     shellHook = ''
                                goversion=$(go version)
                                goversion=''${goversion// /_}
                                export GOPATH="$HOME/gohome/$goversion/go"
                                export GOCACHE="$HOME/gohome/$goversion/cache"
                                export GOENV="$HOME/gohome/$goversion/env"
                                export PATH=$GOPATH/bin:$PATH
                                '';
     
   };
 }