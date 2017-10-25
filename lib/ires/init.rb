require 'ffi'
require 'ires/util'

module Ires
  module Init
    extend FFI::Library
    extend Ires::Util

    os = current_os
    return if os.nil?

    # NOTE: ires.so is golang object
    ffi_lib File.expand_path("../../shared/#{os}/ires.so", File.dirname(__FILE__))
    
    # resize func
    attach_function :iresImagePath, [:string, :int, :int, :string, :string, :string], :string
  end
end
