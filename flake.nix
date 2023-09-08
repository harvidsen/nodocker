{
  description = "Build and deploy NixOS servers";

  nixConfig.flake-registry = "https://raw.githubusercontent.com/fornybar/registry/nixos-23.05/registry.json";

  outputs = {
    self,
    nixpkgs,
    azure
  }@inputs:
  let
    system = "x86_64-linux";
    pkgs = import nixpkgs { inherit system; };

    # publishDocker = pkgs.writeShellApplication {
    #   name = "publishDocker"; runtimeInputs = with pkgs; [ docker-client ];
    #   text = with self.packages.${system}; ''
    #     docker load < ${docker}
    #     docker tag ${docker.imageName}:${docker.imageTag} "ghcr.io/harvidsen/nodocker:${docker.imageTag}"
    #     docker tag ${docker.imageName}:${docker.imageTag} "ghcr.io/harvidsen/nodocker:latest"
    #     docker push "ghcr.io/harvidsen/nodocker:${docker.imageTag}"
    #     docker push "ghcr.io/harvidsen/nodocker:latest"
    #   '';  # Needs to be logged in "docker login ghcr.io"
    # };

  in {
    packages.${system}.images = {
      simple = import ./images/simple.nix { inherit pkgs; };
      webserver = import ./images/webserver { inherit pkgs; };
    };
  };
}
