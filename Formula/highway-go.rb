# typed: false
# frozen_string_literal: true

# This file was generated by GoReleaser. DO NOT EDIT.
class HighwayGo < Formula
  desc "The Official Sonr Highway CLI tool for building and deploying services on the Sonr network/blockchain."
  homepage "https://sonr.io"
  version "0.0.2"
  license "GPLv3"

  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/sonr-io/highway-go/releases/download/v0.0.2/highwayd-0.0.2-darwin-arm64.tar.gz"
      sha256 "c6456883c139e7ed3f7909c2500074b36bea2be3917d7deac264b7f0c1c80775"

      def install
        bin.install "highwayd"
      end
    end
    if Hardware::CPU.intel?
      url "https://github.com/sonr-io/highway-go/releases/download/v0.0.2/highwayd-0.0.2-darwin-amd64.tar.gz"
      sha256 "aeae4027e81230ee1882855834a95d8ff1d52d079edd0c6832fd9be2ed803c50"

      def install
        bin.install "highwayd"
      end
    end
  end

  on_linux do
    if Hardware::CPU.arm? && Hardware::CPU.is_64_bit?
      url "https://github.com/sonr-io/highway-go/releases/download/v0.0.2/highwayd-0.0.2-linux-arm64.tar.gz"
      sha256 "17170cdf86f020a2dd10e33d5c995ca9275ef912a65483bcf352acc3b3ccf349"

      def install
        bin.install "highwayd"
      end
    end
    if Hardware::CPU.intel?
      url "https://github.com/sonr-io/highway-go/releases/download/v0.0.2/highwayd-0.0.2-linux-amd64.tar.gz"
      sha256 "dc6a8ef094b3104afc9a52116f2ee8d0740add290eefe8f97001abf81b26cd6f"

      def install
        bin.install "highwayd"
      end
    end
  end

  def caveats; <<~EOS
    Run `brew info sonr` for more information, or run `sonr --help`.
  EOS
  end
end
