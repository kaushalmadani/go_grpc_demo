FROM envoyproxy/envoy:v1.12.2

COPY ./envoy.yaml /etc/envoy/envoy.yaml

CMD /usr/local/bin/envoy -c /etc/envoy/envoy.yaml


# docker build -t my-envoy:1.0 .
# docker run --network=host grpc-medium-envoy:1.0