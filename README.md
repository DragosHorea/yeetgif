# yeetgif

## Get it

```sh
go get -u github.com/sgreben/yeetgif/cmd/gif
```

Or [download the binary](https://github.com/sgreben/yeetgif/releases/latest) from the releases page.

```sh
# Linux
curl -LO https://github.com/sgreben/yeetgif/releases/download/1.0.5/gif_1.0.5_linux_x86_64.tar.gz | tar xz

# OS X
curl -LO https://github.com/sgreben/yeetgif/releases/download/1.0.5/gif_1.0.5_osx_x86_64.tar.gz | tar xz

# Windows
curl -LO https://github.com/sgreben/yeetgif/releases/download/1.0.5/gif_1.0.5_windows_x86_64.zip
unzip gif_1.0.5_windows_x86_64.zip
```

## Use it

```sh
<doc/yeet.png | gif fried | gif wobble | gif crop | gif optimize >doc/yeet.gif
```
![before](doc/yeet.png)
![after](doc/yeet.gif)


## Usage

```text

Usage: gif [OPTIONS] COMMAND [arg...]

v1.0.0-32f76fdf8817692838c799f5c2f766cc47f5e70d

Options:
  -n                 Duplicate a single input image this many times (default 20)
  -q, --quiet        Disable all log output (stderr)
  -d, --delay-ms     Frame delay in milliseconds (default 20)
  -p, --pad          Pad images (default true)
      --write-meta   Write command line options into output GIF metadata (default true)

Commands:
  roll               (☭ ͜ʖ ☭)
  wobble             🍆( ͡° ͜ʖ ͡°)🍆
  pulse              ( ͡◉ ͜ʖ ͡◉)
  zoom               (⌐▀͡ ̯ʖ▀)
  shake              ˵(˵ ͡⚆ ͜ʖ ͡⚆˵)˵
  woke               💯  W O K E F L A R E S ( ͡ 🅱️ ͜ʖ ͡ 🅱️ ) 💯
  fried              fr͍͈i̗̟̲̻e͕̗d̬ m̷͔͊e̶̪̿m̷̙̈́é̵̤s̷̺͒
  hue                ( ͡☆ ͜ʖ ͡☆)
  tint               🎨༼ຈل͜ຈ༽
  resize             (° ͜ʖ°)¯\_( ͡☉ ͜ʖ ͡☉)_/¯
  crop               ʖ ͡☉)
  optimize           👌( ͡ᵔ ͜ʖ ͡ᵔ )👌
  meta               (🧠 ͡ಠ ʖ̯ ͡ಠ)┌

Run 'gif COMMAND --help' for more information on a command.
```

## Licensing

- [Modified copy](pkg/imaging) of `github.com/disintegration/imaging`: [MIT License](pkg/imaging/LICENSE)
- `yeetgif` itself: [MIT License](LICENSE)
