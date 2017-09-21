require 'test_helper'

class ToppageControllerTest < ActionDispatch::IntegrationTest
  test "should get index" do
    get toppage_index_url
    assert_response :success
  end

end
