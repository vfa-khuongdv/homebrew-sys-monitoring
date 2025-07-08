class SysMonitoring < Formula
  desc "Real-time system monitoring tool with beautiful terminal UI"
  homepage "https://github.com/vfa-khuongdv/homebrew-sys-monitoring"
  url "https://github.com/vfa-khuongdv/homebrew-sys-monitoring/releases/download/v1.0.0/sys-monitoring-v1.0.0-binary.tar.gz"
  sha256 "c4b541d5a9142856e727f9ed35ab26414d3de9c5df86a9adefe3354014999f64"  # This will be replaced with actual SHA256 when you create a release
  license "MIT"
  head "https://github.com/vfa-khuongdv/homebrew-sys-monitoring.git", branch: "main"

  def install
    # Install the pre-built binary
    bin.install "sys-monitoring"
  end

  test do
    # Test that the binary was installed correctly
    assert_match "📊", shell_output("#{bin}/sys-monitoring --help 2>&1", 1)
  end
end
