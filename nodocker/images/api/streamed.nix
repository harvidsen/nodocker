{ pkgs }:

let
  python = pkgs.python3.withPackages (ps: with ps; [ uvicorn fastapi ]);

in
pkgs.dockerTools.streamLayeredImage {
  name = "api";
  tag = "latest";

  config = {
    Cmd = [ "${python}/bin/python" "${./main.py}" ];
  };

  copyToRoot = with pkgs.dockerTools; [
    binSh
    pkgs.cowsay
  ];
}
