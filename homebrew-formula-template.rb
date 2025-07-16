class HelmRelease < Formula
  desc "CLI tool that performs Helm releases using standard Kubernetes resource YAML files"
  homepage "https://github.com/meithon/helm-release"
  url "https://github.com/meithon/helm-release/archive/refs/tags/v0.1.0.tar.gz"
  sha256 "cc86e974b6f3e9cf03b9d4e7722228ba5f4b68004b5b4ff8a2d9f304fa66cfd0"
  license "MIT"

  depends_on "go" => :build

  def install
    system "go", "build", *std_go_args
  end

  test do
    system "#{bin}/helm-release", "--help"
  end
end
