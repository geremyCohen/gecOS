# gecOS
OS Written in ASM, based on MikeOS

## Requirements
<li>nasm
<li>qemu
<li>[MikeOS](http://mikeos.sourceforge.net/#downloads)

### Compile
```
nasm -f bin -o myfirst.bin myfirst.asm
```

### Write boot sector using dd
```
dd status=noxfer conv=notrunc if=myfirst.bin of=myfirst.flp
```

### Boot floppy image
```
qemu-system-i386 -fda myfirst.flp
```

