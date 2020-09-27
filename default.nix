{ stdenv
, lib
, buildGoModule
, gitignoreSource
, nixFilter
}:
buildGoModule rec {
  pname = "packer-builder-delete-vultr";
  version = "1.0.14";
  src = nixFilter (gitignoreSource ./.);
  vendorSha256 = "0d6qgz4s6la24nd50c8hzbmk0gi63rhy77rdrqvgrn07wrab2clh";
  subPackages = [ "cmd/packer-builder-delete-vultr" ];
  postInstall = ''
    ln -sf $out/bin/packer-builder-delete-vultr \
      $out/bin/packer-builder-read-vultr
  '';
  meta = with stdenv.lib; {
    description = "Packer builder to delete vultr images";
    license = licenses.mpl20;
    platforms = platforms.linux;
  };
}
