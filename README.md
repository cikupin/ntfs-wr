# ntfs-wr
Simple CLI to mount and enable write on NTFS file system easily on MacOS.

Tested on:
1. Macbook Pro 14", M1 Pro, 16GB ram, MacOS Monterey 12.6.1

## Requirements

## Installation

```bash
$ go install github.com/cikupin/ntfs-wr 
```

## Usage

1. See list of commands
   ```bash
   $ ntfs-wr --help
   ```

2. List external drive with NTFS file system
   ```bash
   $ ntfs-wr list
   ```

3. Mount external NTFS partition (you need to run as root)
   ```bash
   $ sudo ntfs-wr mount
   ```

4. Open NTFS mounted volume. You must do copy-paste on this directory.
   ```bash
   $ ntfs-wr open <volume dir name>
   ```

4. Unmount mounted volume (you need to run as root)
   ```bash
   $ sudo ntfs-wr unmount <volume dir name>
   ```

## To do plan

- [ ] include partition size on partition list info
- [ ] persist mounting configuration on local
- [ ] unmount volume using selection menu
