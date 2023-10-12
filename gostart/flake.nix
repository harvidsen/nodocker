{
  description = "GO develop environment";

  nixConfig.flake-registry = "https://raw.githubusercontent.com/fornybar/registry/nixos-23.05/registry.json";

  outputs = { self, nixpkgs }@inputs:

  let
    system = "x86_64-linux";
    pkgs = import nixpkgs { inherit system; };

  in
  {
    devShells.${system}.default = pkgs.mkShell {
      buildInputs = with pkgs; [ go gopls ];
    };
  };
}
