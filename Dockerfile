# gf docker -t go-gf-blog
FROM loads/alpine:3.8

LABEL maintainer="wmnihaoya@qq.com"

###############################################################################
#                                INSTALLATION
###############################################################################

# 设置固定的项目路径
ENV WORKDIR /usr/local/docker/go-gf-blog

# 添加应用可执行文件，并设置执行权限
ADD ./main   $WORKDIR/main
RUN chmod +x $WORKDIR/main

# 添加I18N多语言文件、静态文件、配置文件、模板文件
ADD public   $WORKDIR/config

###############################################################################
#                                   START
###############################################################################
WORKDIR $WORKDIR
CMD ./main
