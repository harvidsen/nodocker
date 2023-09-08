{ pkgs }:

pkgs.dockerTools.buildImage {
  name = "webserver";
  tag = "latest";

  fromImage = pkgs.dockerTools.pullImage {
    imageName = "nginx";
    finalImageTag = "1.25.2";
    imageDigest = "sha256:9504f3f64a3f16f0eaf9adca3542ff8b2a6880e6abfb13e478cca23f6380080a";
    sha256 = "sha256-SYvUdHIl0sN0hF1JXxgqOky3po7YQ9UKlgGCr4oCStY="; # ""; ## Trick for finding sha
  };

  runAsRoot = ''
    #!${pkgs.runtimeShell}
    cp ${./index.html} /usr/share/nginx/html/index.html
  '';

  config = {
    Cmd = [ "nginx" "-g" "daemon off;" ];
    Volumes = { # nginx -s reload ## to reload
      "/usr/share/nginx/html" = { };
    };
  };
}
