FROM progrium/busybox
COPY ./main /server
CMD ["/server"]
