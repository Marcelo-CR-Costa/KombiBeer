FROM sensedia/linux-base:1.4

USER appexecute
ENV HOST=postgres
# Run Image
COPY bin/application application
EXPOSE 8080
ENTRYPOINT ["./application"]