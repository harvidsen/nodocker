# Nodocker

Example for creating docker images with Nix.

Useful documentation
https://nixos.wiki/wiki/Docker#Creating_images
https://nixos.org/manual/nixpkgs/stable/#sec-pkgs-dockerTools

## Guide

### Enable docker on NixOS

NixOS config
```nix
{
  virtualisation.docker.enable = true; # Enable docker daemon
  users.users.<myuser>.extraGroups = [ "docker" ]; # Docker without sudo, but more access!
}
```

### pkgs.dockerTools.buildImage

Simple image
```
docker load < $(nix build .#images.simple --print-out-paths)
docker container run simple:<commit-sha> /bin/sh -c "echo hei"
```


- pkgs.dockerTools.buildLayeredImage
- pkgs.dockerTools.streamLayeredImage
- Translate docker things
  - FROM, fromImage
  - RUN, runAsRoot
  - COPY,  copyToRoot
  - config
    - Cmd, Entrypoint, Env, Volumes
