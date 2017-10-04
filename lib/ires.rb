require 'ires/service'
require 'ires/view_helper'
require 'ires/util'
ActiveSupport.on_load(:action_view) do
  include Ires::ViewHelper
end