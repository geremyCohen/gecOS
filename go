nasm -f bin -o myfirst.bin myfirst.asm
dd status=noxfer conv=notrunc if=myfirst.bin of=myfirst.flp
