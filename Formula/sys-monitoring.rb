class SysMonitoring < Formula
  desc "Real-time system monitoring tool with beautiful terminal UI"
  homepage "https://github.com/vfa-khuongdv/sys-monitoring"
  url "https://github.com/vfa-khuongdv/sys-monitoring/archive/refs/tags/v1.0.0.tar.gz"
  sha256 "4e65ebde706263646071f94ab9c2783dcbaf911580c0183e659347dabb5f7b0d"  # This will be replaced with actual SHA256 when you create a release
  license "MIT"
  head "https://github.com/vfa-khuongdv/sys-monitoring.git", branch: "main"

  depends_on "go" => :build

  def install
    # Build the binary
    system "go", "build", *std_go_args(ldflags: "-s -w"), "./..."
    
    # Install the binary
    bin.install "sys-monitoring"
  end

  test do
    # Test that the binary was installed correctly
    assert_match "📊", shell_output("#{bin}/sys-monitoring --help 2>&1", 1)
  end
end
