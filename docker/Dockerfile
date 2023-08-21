FROM mysql:8.0.34-debian@sha256:966645be9470598ff5b5b8d5289ed2dcd760dccaf0f6cd84481adf9058d1cbfc

COPY my.cnf /etc/mysql/conf.d/my.cnf

# 日本語環境を追加
RUN apt-get update && apt-get install -y locales \
    && sed -i -e 's/# \(ja_JP.UTF-8\)/\1/' /etc/locale.gen \
    && locale-gen \
    && update-locale LANG=ja_JP.UTF-8
# ここまで

ENV LC_ALL ja_JP.UTF-8

ENV TZ Asia/Tokyo
ENV LANG=ja_JP.UTF-8

CMD ["mysqld"]

EXPOSE 3306