FROM public.ecr.aws/docker/library/golang:1.20.4-alpine
COPY ./saintspace /
EXPOSE 3001
CMD [ "/saintspace" ]
