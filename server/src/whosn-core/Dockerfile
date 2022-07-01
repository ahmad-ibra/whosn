FROM postgres
ENV POSTGRES_USER dev
ENV POSTGRES_PASSWORD pass
ENV POSTGRES_DB whosn
COPY schema.sql /docker-entrypoint-initdb.d/
