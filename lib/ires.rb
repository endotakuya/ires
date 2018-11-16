require 'ires/view_helper'

ActiveSupport.on_load(:action_view) do
  include Ires::ViewHelper
end
