# Ires

[![Gem Version](https://badge.fury.io/rb/ires.svg)](https://badge.fury.io/rb/ires)


`Ires` is image resizer gem.

## Usage

### View

```erb
<!-- Usually -->
<%= ires_tag( path: "image_01.jpg", width: 90, height: 120 ) %>

<!-- Using image_tag options -->
<%= ires_tag( path: "http://example.com/image_02.jpg", width: 200, height: 200, Ires::Mode::CROP, alt: "example image" ) %>
```

### Get resize path

```ruby
Ires::Service.path(path: '<FULL IMAGE PATH>', width: 400, height: 300)
=> /ires/<resize image path>
```

### Select mode

| info                       |     　　        mode       　　　  |
|:---------------------------|:---------------------------------|
| Resize                     | Ires::Mode::RESIZE (default)     |
| Cropping                   | Ires::Mode::CROP                 |
| Rsize after Cropping       | Ires::Mode::RESIZE_TO_CROP       |

### Select type

Filter of resize image.

| info                       |     　　       type        　　　  |
|:---------------------------|:---------------------------------|
| All                        | Ires::Type::ALL (default)        |
| Smaller than               | Ires::Type::SMALLER              |
| Larger than                | Ires::Type::LARGER               |

### Specify cache expiration

Default: **30days**

```erb
<%= ires_tag( path: '/image.jpg', width: 400, height: 300, expire: 7.days ) %>
```

### Saved directory

####  Target image is local

```
public
├── image.jpg
└── ires
    ├── crop
    │   ├── 20171019_image_120x90_crop.jpg
    │   ├── 20171117_image_200x200_crop.jpg
    │   └── 20171117_image_400x300_crop.jpg
    ├── resize
    │   ├── 20171019_image_120x90_resize.jpg
    │   ├── 20171117_image_200x200_resize.jpg
    │   └── 20171117_image_400x300_resize.jpg
    └── resize_to_crop
        ├── 20171019_image_120x90_resize_to_crop.jpg
        ├── 20171117_image_200x200_resize_to_crop.jpg
        └── 20171117_image_400x300_resize_to_crop.jpg
```

#### Target image is http

Parse URL & Create directory by parse URL.


## Installation
Add this line to your application's Gemfile:

```ruby
gem 'ires'
```

And then execute:
```bash
$ bundle
```

Or install it yourself as:
```bash
$ gem install ires
```

## Caution

- It works only with `linux` and `darwin` now.
- Can build only linux（.so）in this docker.


## Development

Docker environment.

```shell
$ docker build -t ires:v1 .

# Into the ires container.
$ docker run -it -v $(pwd):/go/src/github.com/endotakuya/ires -p 3000:3000 ires:v1 /bin/bash
```

## Gem test

Working in ires container.

### 1. Go（Create a shared object）

Package manager is [dep](https://github.com/golang/dep).

```shell
# Dependent resolution
$ dep ensure

# Output to a shared object.
$ CGO_ENABLED=1 GOOS=linux go build -v -buildmode=c-shared -o shared/linux/ires.so ext/main.go
```
※ In the current Docker, you can build only linux environment.
※ If you want to build in other environments, add GCC or install Go on the host side.🙇

### 2. Start rails server

```shell
$ test/dummy/bin/rails s -b 0.0.0.0
```

## License
The gem is available as open source under the terms of the [MIT License](http://opensource.org/licenses/MIT).
