{ pkgs }:

pkgs.dockerTools.buildImage {
  name = "simple";

  copyToRoot = with pkgs.dockerTools; [
    binSh
  ];
}
