#!/bin/bash
###
 # @Description: 
 # @Version: 1.0
 # @Autor: Sean
 # @Date: 2023-07-21 15:33:00
 # @LastEditors: Sean
 # @LastEditTime: 2023-07-21 15:35:19
### 
# chkconfig: 2345 80 90
# description: wechatbot service
# processname: wechatbot
# Source function library.
. /etc/rc.d/init.d/functions
# Source networking configuration.
. /etc/sysconfig/network
# Check that networking is up.
[ "$NETWORKING" = "no" ] && exit 0
RETVAL=0
prog="wechatbot"
start() {
        echo -n $"Starting $prog: "
        nohup /home/ec2-user/wechatbot > /home/ec2-user/wechatbot.log 2>&1 &
        RETVAL=$?
        echo
        [ $RETVAL = 0 ] && touch /var/lock/subsys/$prog
        return $RETVAL
}
stop() {
        echo -n $"Stopping $prog: "
        killproc $prog
        RETVAL=$?
        echo
        [ $RETVAL = 0 ] && rm -f /var/lock/subsys/$prog
        return $RETVAL
}
restart() {
        stop
        start
}
case "$1" in
  start)
        start
        ;;
  stop)
        stop
        ;;
  restart)
        restart
        ;;
  status)
        status $prog
        RETVAL=$?
        ;;
  *)
        echo $"Usage: $0 {start|stop|restart|status}"
        RETVAL=1
esac
exit $RETVAL
```
