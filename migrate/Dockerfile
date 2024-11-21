FROM postgres:14

COPY migrate/create_users.sql migrate/migrate.sh /migrate/
RUN chmod +x migrate/migrate.sh

ENTRYPOINT ["/migrate/migrate.sh"]