"".foo STEXT size=256 args=0x10 locals=0x58 funcid=0x0
	0x0000 00000 (pro_2.go:3)	TEXT	"".foo(SB), ABIInternal, $96-16
	0x0000 00000 (pro_2.go:3)	MOVD	16(g), R1
	0x0004 00004 (pro_2.go:3)	PCDATA	$0, $-2
	0x0004 00004 (pro_2.go:3)	MOVD	RSP, R2
	0x0008 00008 (pro_2.go:3)	CMP	R1, R2
	0x000c 00012 (pro_2.go:3)	BLS	244
	0x0010 00016 (pro_2.go:3)	PCDATA	$0, $-1
	0x0010 00016 (pro_2.go:3)	MOVD.W	R30, -96(RSP)
	0x0014 00020 (pro_2.go:3)	MOVD	R29, -8(RSP)
	0x0018 00024 (pro_2.go:3)	SUB	$8, RSP, R29
	0x001c 00028 (pro_2.go:3)	FUNCDATA	ZR, gclocals·2589ca35330fc0fce83503f4569854a0(SB)
	0x001c 00028 (pro_2.go:3)	FUNCDATA	$1, gclocals·9fb7f0986f647f17cb53dda1484e0f7a(SB)
	0x001c 00028 (pro_2.go:3)	FUNCDATA	$5, "".foo.arginfo1(SB)
	0x001c 00028 (pro_2.go:5)	MOVD	$11, R0
	0x0020 00032 (pro_2.go:5)	MOVD	R0, "".foo_val1-16(SP)
	0x0024 00036 (pro_2.go:6)	MOVD	$12, R0
	0x0028 00040 (pro_2.go:6)	MOVD	R0, "".foo_val2-24(SP)
	0x002c 00044 (pro_2.go:7)	MOVD	$type.int(SB), R0
	0x0034 00052 (pro_2.go:7)	MOVD	R0, 8(RSP)
	0x0038 00056 (pro_2.go:7)	PCDATA	$1, ZR
	0x0038 00056 (pro_2.go:7)	CALL	runtime.newobject(SB)
	0x003c 00060 (pro_2.go:7)	MOVD	16(RSP), R0
	0x0040 00064 (pro_2.go:7)	MOVD	R0, "".&foo_val3-8(SP)
	0x0044 00068 (pro_2.go:7)	MOVD	$13, R1
	0x0048 00072 (pro_2.go:7)	MOVD	R1, (R0)
	0x004c 00076 (pro_2.go:8)	MOVD	$14, R1
	0x0050 00080 (pro_2.go:8)	MOVD	R1, "".foo_val4-32(SP)
	0x0054 00084 (pro_2.go:9)	MOVD	$15, R1
	0x0058 00088 (pro_2.go:9)	MOVD	R1, "".foo_val5-40(SP)
	0x005c 00092 (pro_2.go:9)	MOVD	ZR, R1
	0x0060 00096 (pro_2.go:14)	JMP	220
	0x0064 00100 (pro_2.go:14)	MOVD	R1, "".i-48(SP)
	0x0068 00104 (pro_2.go:15)	PCDATA	$1, $1
	0x0068 00104 (pro_2.go:15)	CALL	runtime.printlock(SB)
	0x006c 00108 (pro_2.go:15)	MOVD	$"".arg_val(FP), R0
	0x0070 00112 (pro_2.go:15)	MOVD	R0, 8(RSP)
	0x0074 00116 (pro_2.go:15)	CALL	runtime.printpointer(SB)
	0x0078 00120 (pro_2.go:15)	CALL	runtime.printsp(SB)
	0x007c 00124 (pro_2.go:15)	MOVD	$"".foo_val1-16(SP), R0
	0x0080 00128 (pro_2.go:15)	MOVD	R0, 8(RSP)
	0x0084 00132 (pro_2.go:15)	CALL	runtime.printpointer(SB)
	0x0088 00136 (pro_2.go:15)	CALL	runtime.printsp(SB)
	0x008c 00140 (pro_2.go:15)	MOVD	$"".foo_val2-24(SP), R0
	0x0090 00144 (pro_2.go:15)	MOVD	R0, 8(RSP)
	0x0094 00148 (pro_2.go:15)	CALL	runtime.printpointer(SB)
	0x0098 00152 (pro_2.go:15)	CALL	runtime.printsp(SB)
	0x009c 00156 (pro_2.go:15)	MOVD	"".&foo_val3-8(SP), R0
	0x00a0 00160 (pro_2.go:15)	MOVD	R0, 8(RSP)
	0x00a4 00164 (pro_2.go:15)	CALL	runtime.printpointer(SB)
	0x00a8 00168 (pro_2.go:15)	CALL	runtime.printsp(SB)
	0x00ac 00172 (pro_2.go:15)	MOVD	$"".foo_val4-32(SP), R0
	0x00b0 00176 (pro_2.go:15)	MOVD	R0, 8(RSP)
	0x00b4 00180 (pro_2.go:15)	CALL	runtime.printpointer(SB)
	0x00b8 00184 (pro_2.go:15)	CALL	runtime.printsp(SB)
	0x00bc 00188 (pro_2.go:15)	MOVD	$"".foo_val5-40(SP), R0
	0x00c0 00192 (pro_2.go:15)	MOVD	R0, 8(RSP)
	0x00c4 00196 (pro_2.go:15)	CALL	runtime.printpointer(SB)
	0x00c8 00200 (pro_2.go:15)	CALL	runtime.printnl(SB)
	0x00cc 00204 (pro_2.go:15)	CALL	runtime.printunlock(SB)
	0x00d0 00208 (pro_2.go:14)	MOVD	"".i-48(SP), R0
	0x00d4 00212 (pro_2.go:14)	ADD	$1, R0, R1
	0x00d8 00216 (pro_2.go:19)	MOVD	"".&foo_val3-8(SP), R0
	0x00dc 00220 (pro_2.go:14)	CMP	$5, R1
	0x00e0 00224 (pro_2.go:14)	BLT	100
	0x00e4 00228 (pro_2.go:19)	MOVD	R0, "".~r1+8(FP)
	0x00e8 00232 (pro_2.go:19)	MOVD	-8(RSP), R29
	0x00ec 00236 (pro_2.go:19)	MOVD.P	96(RSP), R30
	0x00f0 00240 (pro_2.go:19)	RET	(R30)
	0x00f4 00244 (pro_2.go:19)	NOP
	0x00f4 00244 (pro_2.go:3)	PCDATA	$1, $-1
	0x00f4 00244 (pro_2.go:3)	PCDATA	$0, $-2
	0x00f4 00244 (pro_2.go:3)	MOVD	R30, R3
	0x00f8 00248 (pro_2.go:3)	CALL	runtime.morestack_noctxt(SB)
	0x00fc 00252 (pro_2.go:3)	PCDATA	$0, $-1
	0x00fc 00252 (pro_2.go:3)	JMP	0
	0x0000 81 0b 40 f9 e2 03 00 91 5f 00 01 eb 49 07 00 54  ..@....._...I..T
	0x0010 fe 0f 1a f8 fd 83 1f f8 fd 23 00 d1 60 01 80 d2  .........#..`...
	0x0020 e0 27 00 f9 e0 07 7e b2 e0 23 00 f9 00 00 00 90  .'....~..#......
	0x0030 00 00 00 91 e0 07 00 f9 00 00 00 94 e0 0b 40 f9  ..............@.
	0x0040 e0 2b 00 f9 a1 01 80 d2 01 00 00 f9 e1 0b 7f b2  .+..............
	0x0050 e1 1f 00 f9 e1 0f 40 b2 e1 1b 00 f9 01 00 80 d2  ......@.........
	0x0060 1f 00 00 14 e1 17 00 f9 00 00 00 94 e0 a3 01 91  ................
	0x0070 e0 07 00 f9 00 00 00 94 00 00 00 94 e0 23 01 91  .............#..
	0x0080 e0 07 00 f9 00 00 00 94 00 00 00 94 e0 03 01 91  ................
	0x0090 e0 07 00 f9 00 00 00 94 00 00 00 94 e0 2b 40 f9  .............+@.
	0x00a0 e0 07 00 f9 00 00 00 94 00 00 00 94 e0 e3 00 91  ................
	0x00b0 e0 07 00 f9 00 00 00 94 00 00 00 94 e0 c3 00 91  ................
	0x00c0 e0 07 00 f9 00 00 00 94 00 00 00 94 00 00 00 94  ................
	0x00d0 e0 17 40 f9 01 04 00 91 e0 2b 40 f9 3f 14 00 f1  ..@......+@.?...
	0x00e0 2b fc ff 54 e0 3b 00 f9 fd 83 5f f8 fe 07 46 f8  +..T.;...._...F.
	0x00f0 c0 03 5f d6 e3 03 1e aa 00 00 00 94 c1 ff ff 17  .._.............
	rel 44+8 t=3 type.int+0
	rel 56+4 t=9 runtime.newobject+0
	rel 104+4 t=9 runtime.printlock+0
	rel 116+4 t=9 runtime.printpointer+0
	rel 120+4 t=9 runtime.printsp+0
	rel 132+4 t=9 runtime.printpointer+0
	rel 136+4 t=9 runtime.printsp+0
	rel 148+4 t=9 runtime.printpointer+0
	rel 152+4 t=9 runtime.printsp+0
	rel 164+4 t=9 runtime.printpointer+0
	rel 168+4 t=9 runtime.printsp+0
	rel 180+4 t=9 runtime.printpointer+0
	rel 184+4 t=9 runtime.printsp+0
	rel 196+4 t=9 runtime.printpointer+0
	rel 200+4 t=9 runtime.printnl+0
	rel 204+4 t=9 runtime.printunlock+0
	rel 248+4 t=9 runtime.morestack_noctxt+0
"".main STEXT size=288 args=0x0 locals=0x58 funcid=0x0
	0x0000 00000 (pro_2.go:22)	TEXT	"".main(SB), ABIInternal, $96-0
	0x0000 00000 (pro_2.go:22)	MOVD	16(g), R1
	0x0004 00004 (pro_2.go:22)	PCDATA	$0, $-2
	0x0004 00004 (pro_2.go:22)	MOVD	RSP, R2
	0x0008 00008 (pro_2.go:22)	CMP	R1, R2
	0x000c 00012 (pro_2.go:22)	BLS	268
	0x0010 00016 (pro_2.go:22)	PCDATA	$0, $-1
	0x0010 00016 (pro_2.go:22)	MOVD.W	R30, -96(RSP)
	0x0014 00020 (pro_2.go:22)	MOVD	R29, -8(RSP)
	0x0018 00024 (pro_2.go:22)	SUB	$8, RSP, R29
	0x001c 00028 (pro_2.go:22)	FUNCDATA	ZR, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001c 00028 (pro_2.go:22)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001c 00028 (pro_2.go:32)	MOVD	$666, R0
	0x0020 00032 (pro_2.go:32)	MOVD	R0, "".arg_val-16(SP)
	0x0024 00036 (<unknown line number>)	NOP
	0x0024 00036 (pro_2.go:5)	MOVD	$11, R0
	0x0028 00040 (pro_2.go:5)	MOVD	R0, "".foo_val1-24(SP)
	0x002c 00044 (pro_2.go:6)	MOVD	$12, R0
	0x0030 00048 (pro_2.go:6)	MOVD	R0, "".foo_val2-32(SP)
	0x0034 00052 (pro_2.go:7)	MOVD	$13, R0
	0x0038 00056 (pro_2.go:7)	MOVD	R0, "".foo_val3-40(SP)
	0x003c 00060 (pro_2.go:8)	MOVD	$14, R0
	0x0040 00064 (pro_2.go:8)	MOVD	R0, "".foo_val4-48(SP)
	0x0044 00068 (pro_2.go:9)	MOVD	$15, R0
	0x0048 00072 (pro_2.go:9)	MOVD	R0, "".foo_val5-56(SP)
	0x004c 00076 (pro_2.go:9)	MOVD	ZR, R0
	0x0050 00080 (pro_2.go:14)	JMP	200
	0x0054 00084 (pro_2.go:14)	MOVD	R0, "".i-64(SP)
	0x0058 00088 (pro_2.go:15)	PCDATA	$1, ZR
	0x0058 00088 (pro_2.go:15)	CALL	runtime.printlock(SB)
	0x005c 00092 (pro_2.go:15)	MOVD	$"".arg_val-16(SP), R0
	0x0060 00096 (pro_2.go:15)	MOVD	R0, 8(RSP)
	0x0064 00100 (pro_2.go:15)	CALL	runtime.printpointer(SB)
	0x0068 00104 (pro_2.go:15)	CALL	runtime.printsp(SB)
	0x006c 00108 (pro_2.go:15)	MOVD	$"".foo_val1-24(SP), R0
	0x0070 00112 (pro_2.go:15)	MOVD	R0, 8(RSP)
	0x0074 00116 (pro_2.go:15)	CALL	runtime.printpointer(SB)
	0x0078 00120 (pro_2.go:15)	CALL	runtime.printsp(SB)
	0x007c 00124 (pro_2.go:15)	MOVD	$"".foo_val2-32(SP), R0
	0x0080 00128 (pro_2.go:15)	MOVD	R0, 8(RSP)
	0x0084 00132 (pro_2.go:15)	CALL	runtime.printpointer(SB)
	0x0088 00136 (pro_2.go:15)	CALL	runtime.printsp(SB)
	0x008c 00140 (pro_2.go:15)	MOVD	$"".foo_val3-40(SP), R0
	0x0090 00144 (pro_2.go:15)	MOVD	R0, 8(RSP)
	0x0094 00148 (pro_2.go:15)	CALL	runtime.printpointer(SB)
	0x0098 00152 (pro_2.go:15)	CALL	runtime.printsp(SB)
	0x009c 00156 (pro_2.go:15)	MOVD	$"".foo_val4-48(SP), R0
	0x00a0 00160 (pro_2.go:15)	MOVD	R0, 8(RSP)
	0x00a4 00164 (pro_2.go:15)	CALL	runtime.printpointer(SB)
	0x00a8 00168 (pro_2.go:15)	CALL	runtime.printsp(SB)
	0x00ac 00172 (pro_2.go:15)	MOVD	$"".foo_val5-56(SP), R0
	0x00b0 00176 (pro_2.go:15)	MOVD	R0, 8(RSP)
	0x00b4 00180 (pro_2.go:15)	CALL	runtime.printpointer(SB)
	0x00b8 00184 (pro_2.go:15)	CALL	runtime.printnl(SB)
	0x00bc 00188 (pro_2.go:15)	CALL	runtime.printunlock(SB)
	0x00c0 00192 (pro_2.go:14)	MOVD	"".i-64(SP), R0
	0x00c4 00196 (pro_2.go:14)	ADD	$1, R0, R0
	0x00c8 00200 (pro_2.go:14)	CMP	$5, R0
	0x00cc 00204 (pro_2.go:14)	BLT	84
	0x00d0 00208 (pro_2.go:34)	MOVD	"".foo_val3-40(SP), R0
	0x00d4 00212 (pro_2.go:34)	MOVD	R0, ""..autotmp_16-8(SP)
	0x00d8 00216 (pro_2.go:34)	CALL	runtime.printlock(SB)
	0x00dc 00220 (pro_2.go:34)	MOVD	""..autotmp_16-8(SP), R0
	0x00e0 00224 (pro_2.go:34)	MOVD	R0, 8(RSP)
	0x00e4 00228 (pro_2.go:34)	CALL	runtime.printint(SB)
	0x00e8 00232 (pro_2.go:34)	CALL	runtime.printsp(SB)
	0x00ec 00236 (pro_2.go:34)	MOVD	$"".foo_val3-40(SP), R0
	0x00f0 00240 (pro_2.go:34)	MOVD	R0, 8(RSP)
	0x00f4 00244 (pro_2.go:34)	CALL	runtime.printpointer(SB)
	0x00f8 00248 (pro_2.go:34)	CALL	runtime.printnl(SB)
	0x00fc 00252 (pro_2.go:34)	CALL	runtime.printunlock(SB)
	0x0100 00256 (pro_2.go:37)	MOVD	-8(RSP), R29
	0x0104 00260 (pro_2.go:37)	MOVD.P	96(RSP), R30
	0x0108 00264 (pro_2.go:37)	RET	(R30)
	0x010c 00268 (pro_2.go:37)	NOP
	0x010c 00268 (pro_2.go:22)	PCDATA	$1, $-1
	0x010c 00268 (pro_2.go:22)	PCDATA	$0, $-2
	0x010c 00268 (pro_2.go:22)	MOVD	R30, R3
	0x0110 00272 (pro_2.go:22)	CALL	runtime.morestack_noctxt(SB)
	0x0114 00276 (pro_2.go:22)	PCDATA	$0, $-1
	0x0114 00276 (pro_2.go:22)	JMP	0
	0x0000 81 0b 40 f9 e2 03 00 91 5f 00 01 eb 09 08 00 54  ..@....._......T
	0x0010 fe 0f 1a f8 fd 83 1f f8 fd 23 00 d1 40 53 80 d2  .........#..@S..
	0x0020 e0 27 00 f9 60 01 80 d2 e0 23 00 f9 e0 07 7e b2  .'..`....#....~.
	0x0030 e0 1f 00 f9 a0 01 80 d2 e0 1b 00 f9 e0 0b 7f b2  ................
	0x0040 e0 17 00 f9 e0 0f 40 b2 e0 13 00 f9 00 00 80 d2  ......@.........
	0x0050 1e 00 00 14 e0 0f 00 f9 00 00 00 94 e0 23 01 91  .............#..
	0x0060 e0 07 00 f9 00 00 00 94 00 00 00 94 e0 03 01 91  ................
	0x0070 e0 07 00 f9 00 00 00 94 00 00 00 94 e0 e3 00 91  ................
	0x0080 e0 07 00 f9 00 00 00 94 00 00 00 94 e0 c3 00 91  ................
	0x0090 e0 07 00 f9 00 00 00 94 00 00 00 94 e0 a3 00 91  ................
	0x00a0 e0 07 00 f9 00 00 00 94 00 00 00 94 e0 83 00 91  ................
	0x00b0 e0 07 00 f9 00 00 00 94 00 00 00 94 00 00 00 94  ................
	0x00c0 e0 0f 40 f9 00 04 00 91 1f 14 00 f1 4b fc ff 54  ..@.........K..T
	0x00d0 e0 1b 40 f9 e0 2b 00 f9 00 00 00 94 e0 2b 40 f9  ..@..+.......+@.
	0x00e0 e0 07 00 f9 00 00 00 94 00 00 00 94 e0 c3 00 91  ................
	0x00f0 e0 07 00 f9 00 00 00 94 00 00 00 94 00 00 00 94  ................
	0x0100 fd 83 5f f8 fe 07 46 f8 c0 03 5f d6 e3 03 1e aa  .._...F..._.....
	0x0110 00 00 00 94 bb ff ff 17 00 00 00 00 00 00 00 00  ................
	rel 88+4 t=9 runtime.printlock+0
	rel 100+4 t=9 runtime.printpointer+0
	rel 104+4 t=9 runtime.printsp+0
	rel 116+4 t=9 runtime.printpointer+0
	rel 120+4 t=9 runtime.printsp+0
	rel 132+4 t=9 runtime.printpointer+0
	rel 136+4 t=9 runtime.printsp+0
	rel 148+4 t=9 runtime.printpointer+0
	rel 152+4 t=9 runtime.printsp+0
	rel 164+4 t=9 runtime.printpointer+0
	rel 168+4 t=9 runtime.printsp+0
	rel 180+4 t=9 runtime.printpointer+0
	rel 184+4 t=9 runtime.printnl+0
	rel 188+4 t=9 runtime.printunlock+0
	rel 216+4 t=9 runtime.printlock+0
	rel 228+4 t=9 runtime.printint+0
	rel 232+4 t=9 runtime.printsp+0
	rel 244+4 t=9 runtime.printpointer+0
	rel 248+4 t=9 runtime.printnl+0
	rel 252+4 t=9 runtime.printunlock+0
	rel 272+4 t=9 runtime.morestack_noctxt+0
go.cuinfo.packagename. SDWARFCUINFO dupok size=0
	0x0000 6d 61 69 6e                                      main
""..inittask SNOPTRDATA size=24
	0x0000 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0010 00 00 00 00 00 00 00 00                          ........
go.info."".foo$abstract SDWARFABSFCN dupok size=106
	0x0000 04 2e 66 6f 6f 00 01 01 11 61 72 67 5f 76 61 6c  ..foo....arg_val
	0x0010 00 00 00 00 00 00 0c 66 6f 6f 5f 76 61 6c 31 00  .......foo_val1.
	0x0020 05 00 00 00 00 0c 66 6f 6f 5f 76 61 6c 32 00 06  ......foo_val2..
	0x0030 00 00 00 00 0c 66 6f 6f 5f 76 61 6c 33 00 07 00  .....foo_val3...
	0x0040 00 00 00 0c 66 6f 6f 5f 76 61 6c 34 00 08 00 00  ....foo_val4....
	0x0050 00 00 0c 66 6f 6f 5f 76 61 6c 35 00 09 00 00 00  ...foo_val5.....
	0x0060 00 0c 69 00 0e 00 00 00 00 00                    ..i.......
	rel 0+0 t=23 type.*int+0
	rel 0+0 t=23 type.int+0
	rel 18+4 t=31 go.info.int+0
	rel 33+4 t=31 go.info.int+0
	rel 48+4 t=31 go.info.int+0
	rel 63+4 t=31 go.info.int+0
	rel 78+4 t=31 go.info.int+0
	rel 93+4 t=31 go.info.int+0
	rel 101+4 t=31 go.info.int+0
gclocals·2589ca35330fc0fce83503f4569854a0 SRODATA dupok size=10
	0x0000 02 00 00 00 02 00 00 00 00 00                    ..........
gclocals·9fb7f0986f647f17cb53dda1484e0f7a SRODATA dupok size=10
	0x0000 02 00 00 00 01 00 00 00 00 01                    ..........
"".foo.arginfo1 SRODATA static dupok size=3
	0x0000 00 08 ff                                         ...
gclocals·33cdeccccebe80329f1fdbee7f5874cb SRODATA dupok size=8
	0x0000 01 00 00 00 00 00 00 00                          ........
