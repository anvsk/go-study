* 为什么客户端在timewait阶段要等待2msl
msl:最大数据段周期,2msl=服务端fin+客户端ack周期

* 两次握手为什么不行
如果服务端发的丢包,客户端没有收到,这种情况客户端不会给服务端发数据

* 为什么挥手比握手多一次
服务端 第一次挥手 直接回应ack 进入close wait
然后等待数据传输完成
完成后第二次挥手 发送fin
客户端ack
(被关闭方分两次回复的,传输前和传输完成后)

* tcp三次握手:
c->s:syn M ,seq N
s->c:ack N+1,seq O
c->s:ack O+1
四次挥手:
c->s:Fin seq M c:fin_wait1 
s-c:ack M+1 seq N s:close_wait c:fin_wait2
s->c:ack M+1 seq O  Fin s:last_ack
c->s:ack O+1 seq M+1 s:closed 



