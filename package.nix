{
  buildGoModule,
  libGL,
  pkg-config,
  xorg,

  makeWrapper,
}:

buildGoModule {
  pname = "abstract";
  version = "0.1.0";

  src = ./.;

  vendorHash = null;

  subPackages = [ "cmd/abstract" ];

  buildInputs = [
    libGL
    xorg.libX11
    xorg.libXcursor
    xorg.libXext
    xorg.libXi
    xorg.libXinerama
    xorg.libXrandr
    xorg.libXxf86vm
  ];

  nativeBuildInputs = [
    pkg-config
    makeWrapper
  ];

  postInstall = ''
    wrapProgram "$out/bin/abstract" \
      --set LD_LIBRARY_PATH "${libGL}/lib:${libGL}/lib:${xorg.libX11}/lib"
  '';

  doCheck = false;
}

