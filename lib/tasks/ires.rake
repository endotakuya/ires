require 'ires/os'

desc 'Build shared object'
namespace :ires do
  task :build do
    os = Ires::Os.current
    return if os.nil?
    exec "CGO_ENABLED=1 GOOS=\"#{os}\" go build -v -buildmode=c-shared -o shared/\"#{os}\"/ires.so ext/main.go"
  end
end
